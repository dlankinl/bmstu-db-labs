-- 1. Хранимая процедура с параметрами
-- Добавить предпринимателя
create or replace procedure insert_entrepreneur(
    first_name text,
    last_name text,
    age int,
    gender boolean,
    married boolean,
    net_worth int,
    birth_date date
) as '
    insert into lab.enterpreneurs(first_name, last_name, age, gender, married, net_worth, birth_date)
    values (first_name, last_name, age, gender, married, net_worth, birth_date);
' language sql;

call insert_entrepreneur(
    'Dmitry',
    'Lankin',
    20,
    true,
    false,
    1000,
    '2003-09-29');

select *
from lab.enterpreneurs
where first_name = 'Dmitry'
    and last_name = 'Lankin';

-- 2. Рекурсивная хранимая процедура или хранимую процедур с рекурсивным ОТВ
-- Вычисляются числа Фибоначчи до max
create or replace procedure fibP
    (
        max int,
        res inout int,
        cur int default 1,
        next int default 1
    ) as '
    declare
        sum int;
    begin
        sum = cur + next;
        if sum <= max then
            res = sum;
            raise info ''%'', cur;
            call fibP(max, res, next, sum);
        end if;
    end' language plpgsql;

call fibP(24,null);

-- 3. Хранимая процедура с курсором

select *
into tmp_cursor
from lab.skilldescription;

select *
from tmp_cursor;

create or replace procedure proc_update_cursor
(
    old_skill_descr text,
    new_skill_descr text
)
as '
    declare
        curs cursor for
            select *
            from tmp_cursor
            where description = old_skill_descr;
        tmp tmp_cursor;
    begin
        open curs;

        loop
            fetch curs
                into tmp;
            exit when not found;

            update tmp_cursor
            set description = new_skill_descr
            where tmp_cursor.skill_id = tmp.skill_id;

            raise notice ''Elem =  %'', tmp;
        end loop;

        close curs;
    end;
'language  plpgsql;

call proc_update_cursor('He spoke of the happiness that was now certainly theirs, of the folly of not breaking sooner out of that magnificent prison of latter-day life, of the old romantic days that had passed from the world for ever.', 'new description!');

select *
from tmp_cursor
where description = 'new description!';
select *
from tmp_cursor
where description = 'He spoke of the happiness that was now certainly theirs, of the folly of not breaking sooner out of that magnificent prison of latter-day life, of the old romantic days that had passed from the world for ever.';

-- 4. Хранимая процедура доступа к метаданным
create or replace procedure meta_data(name text)
as '
    declare
        elem record;
    begin
        for elem in (select column_name, data_type
                     from information_schema.columns
                     where table_name = name)
        loop
            raise info ''column = %'', elem;
        end loop;
    end;
' language plpgsql;

call meta_data('enterpreneurs');