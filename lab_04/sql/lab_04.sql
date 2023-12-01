CREATE EXTENSION plpython3u;
SELECT * FROM pg_language;

-- • Определяемую пользователем скалярную функцию CLR
-- Выводит название компании по ее идентификатору
create or replace function get_company_name(id_ int) returns varchar
as $$
    res = plpy.execute(f" \
        SELECT name \
        FROM lab.companies  \
        WHERE financials_id = {id_};")

    return res[0]["name"]
$$ language plpython3u;

select get_company_name(15);

-- • Пользовательскую агрегатную функцию CLR
-- Выводит средний капитал представителей пола
create or replace function get_avg_net_worth(gender bool)
    returns decimal(10, 2)
AS $$
    res = plpy.execute(f" \
        select AVG(net_worth) \
        from lab.enterpreneurs      \
        where gender = {gender};")
    return res[0]['avg']
$$ language plpython3u;

select get_avg_net_worth(false);

-- • Определяемую пользователем табличную функцию CLR
-- Выводит название компании, ее выручку и ФИ основателя
create or replace function get_cities_with_companies(revenue_min int)
    returns table
        (
            name text,
            revenue decimal(15, 2),
            owner_first_name text,
            owner_last_name text
        )
as $$
    table = plpy.execute(f"\
        select tmp.name, revenue, e.first_name, e.last_name \
        from \
            (select c.name as name, c.owner_id as owner_id, f.revenue as revenue \
            from lab.companies c join lab.financials f \
            on c.financials_id = f.id \
            where f.revenue > {revenue_min}) as tmp join lab.enterpreneurs e \
            on tmp.owner_id = e.id \
        order by revenue;")
    res = []

    for elem in table:
        res.append((elem["name"], elem["revenue"], elem["first_name"], elem["last_name"]))

    return res
$$ language plpython3u;

select *
from get_cities_with_companies(9000000);

-- • Хранимую процедуру CLR
-- Изменяет статус каждого мужчины с "холост" на "женат"
select *
into enterpreneurs_tmp
from lab.enterpreneurs;

create or replace procedure marry_every_man()
as $$
    plpy.execute(f"\
        update enterpreneurs_tmp \
        set \
            married = true \
        where gender = true; \
    ")
$$ language plpython3u;

select *
from enterpreneurs_tmp
where gender = true
order by married;

call marry_every_man();

-- • Триггер CLR
-- Вместо удаления умения по идентификатору изменяет его описание
create view tmp_descr as
select *
from lab.skilldescription
where skill_id between 1000 and 1400;

create or replace function del_skill_descr_func()
    returns trigger language plpython3u
as $$
    plpy.execute(f'''
        update tmp_descr
        set
            description = 'here is no description:('
        where tmp_descr.skill_id = {TD["old"]["skill_id"]};
        ''')
    return "MODIFY"
$$;

create trigger del_descr_trigger
    instead of delete on tmp_descr
    for each row
execute procedure del_skill_descr_func();

delete from tmp_descr
where skill_id = 1323;

select *
from tmp_descr
where skill_id = 1323;

-- • Определяемый пользователем тип данных CLR
-- Тип "финансы", в котором есть доп поле - рентабельность (в п.п.)
drop type Financials cascade ;
create type Financials AS
(
    id int,
    revenue decimal(15, 2),
    profit decimal(15, 2),
    profit_margin decimal(5, 2)
);

create or replace function GetFinancials()
    RETURNS setof Financials language plpython3u
as $$
    res = plpy.cursor(f"\
        select id, revenue, profit \
        from lab.financials \
        where revenue > profit \
        order by revenue")

    for elem in map(lambda elem: (elem['id'], elem['revenue'], elem['profit'], elem['profit'] / elem['revenue'] * 100), res):
        yield elem
$$;

select *
from GetFinancials();


-- Реализовать функцию, получающую в качестве параметров метаданные
select *
from pg_constraint;

create or replace function metadata(conrelid oid)
    returns table (
                      oid int,
                      conname text,
                      contype char
                  )
as $$
    table = plpy.execute(f"\
        select oid, conname, contype \
        from pg_constraint \
        where conrelid = {conrelid}")
    res = []

    for elem in table:
        res.append((elem["oid"], elem["conname"], elem["contype"]))

    return res
$$ language plpython3u;

select *
from metadata(16447);
