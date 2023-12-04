-- 1. извлечь данные в JSON
\copy (select json_agg(row_to_json(e))::text from lab.enterpreneurs e) to '/json_data/entrepreneurs.json';
\copy (select json_agg(row_to_json(f))::text from lab.financials f) to '/json_data/financials.json';
\copy (select json_agg(row_to_json(e))::text from lab.enterpreneurskill e) to '/json_data/skills_name.json';
\copy (select json_agg(row_to_json(s))::text from lab.skilldescription s) to '/json_data/skills_descr.json';
\copy (select json_agg(row_to_json(c))::text from lab.cities c) to '/json_data/cities.json';
\copy (select json_agg(row_to_json(c))::text from lab.companies c) to '/json_data/companies.json';

-- 2. Выполнить загрузку и сохранение JSON файла в таблицу.

create table if not exists tmp_data (
    data jsonb
);

create table if not exists lab.entrepreneurs_copy(
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    age INT NOT NULL,
    gender BOOLEAN NOT NULL,
    married BOOLEAN NOT NULL,
    net_worth INT NOT NULL,
    birth_date DATE
);

copy tmp_data(data) from '/json_data/entrepreneurs.json';
select * from tmp_data;

select jsonb_array_elements(data::jsonb) as json_data from tmp_data;

insert into lab.entrepreneurs_copy(first_name, last_name, age, gender, married, net_worth, birth_date)
select dt.data->>'first_name', dt.data->>'last_name', (dt.data->>'age')::int, (dt.data->>'gender')::bool, (dt.data->>'married')::bool, (dt.data->>'net_worth')::int, (dt.data->>'birth_date')::date from (select jsonb_array_elements(data::jsonb) as data from tmp_data)
as dt;

select *
from lab.entrepreneurs_copy;

select * from lab.enterpreneurs except (select * from lab.entrepreneurs_copy);

--3. Создать таблицу, в которой будет атрибут с типом JSON, или добавить к уже существующей таблице.
create table if not exists lab.json_table (
    id serial primary key,
    name text not null,
    data json not null
);

insert into lab.json_table(name, data) values
    ('Samsung', '{"model": "Galaxy S23", "rom": "256"}'::json),
    ('Apple', '{"model": "14 Pro Max", "rom": "512"}'::json),
    ('Xiaomi', '{"model": "Note 13 Pro", "rom": "256"}'::json);

select * from lab.json_table;
--4.1. Извлечь JSON фрагмент из JSON документа
create table if not exists lab.phone_description (
    model text,
    rom int
);

select * from lab.json_table, json_populate_record(null::lab.phone_description, data);

select
    data->'model' as model,
    data->'rom' as rom
from lab.json_table;

--4.2. Извлечь значения конкретных узлов или атрибутов JSON документа
create table if not exists lab.companies_tmp (
    data jsonb
);

insert into lab.companies_tmp (data) values
        ('{
        "name": "Oracle",
        "city": "Austin",
        "owner":  {
                  "name": "Larry Ellison",
                  "age": 79
                  }
        }'),
        ('{
          "name": "Pied Piper",
          "city": "Palo Alto",
          "owner":  {
            "name": "Richard Hendricks",
            "age": 35
          }
        }'),
        ('{
          "name": "Aviato",
          "city": "Palo Alto",
          "owner":  {
            "name": "Erlich Bachman",
            "age": 42
          }
        }');

truncate table lab.companies_tmp;

select
    data->'owner'->'name' as owner_name,
    data->'name' as name
from lab.companies_tmp;

--4.3. Выполнить проверку существования узла или атрибута

select exists(
    select *
    from lab.companies_tmp
    where data::jsonb ? 'nft');

--4.4. Изменить JSON документ
update lab.companies_tmp
    set data = data || '{"city": "Palo Alto"}'::jsonb
where (data->>'city')::text = 'CA';

select * from lab.companies_tmp;

--4.5. Разделить JSON документ на несколько строк по узлам

create table if not exists lab.entrepreneurs_tmp(
    data jsonb
);

insert into lab.entrepreneurs_tmp values
    ('[{"name": "Richard Hendricks", "age": 35, "company_name": "Pied Piper"},
       {"name": "Erlich Bachman", "age": 42, "company_name": "Aviato"},
       {"name": "Peter Gregory", "age": 48, "company_name": "Investor"}]'::jsonb);

SELECT jsonb_array_elements(data::jsonb)
FROM lab.entrepreneurs_tmp;