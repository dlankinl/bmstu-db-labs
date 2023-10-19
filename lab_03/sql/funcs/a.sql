-- 1. Скалярная функция
-- Возвращает максимальный капитал предпринимателя
create or replace function get_max_net_worth()
    returns int as '
        select max(net_worth)
        from lab.enterpreneurs;'
    language sql;

select get_max_net_worth() as max_net_worth;

-- 2. Подставляемая табличная функция
-- Возвращает таблицу с id, name компаний, расположенных в данном городе
create or replace function get_companies_in_city(city text)
    returns table(
        id int,
        name text
                 ) as '
    select com.id, com.name
    from lab.companies com join lab.cities cit
                                on com.city_id = cit.id
    where cit.name = city;
' language sql;

select *
from get_companies_in_city('Baldock');

-- 3. Многооператорная табличная функция
-- Возвращает таблицу с id, name, age компаний, расположенных в данном городе,
-- владельцы которых младше данного возраста
create or replace function get_companies_in_city_with_enterpreneur_le_age(city text, age_ int)
    returns table(
        id int,
        name text,
        age int
                 ) as '
    select com.id, com.name, e.age
    from lab.companies com
        join lab.cities cit
            on com.city_id = cit.id
        join lab.enterpreneurs e
            on com.owner_id = e.id
            where cit.name = city
                and e.age <= age_
    order by e.age desc;
    ' language sql;

select *
from get_companies_in_city_with_enterpreneur_le_age('Baldock', 30);

-- 4. Рекурсивная функция или функция с рекурсивным ОТВ
create or replace function fib(first int, second int,max int)
    returns table (
        cur int,
        next int,
        sum int
                  ) as '
    begin
        return query
            select first, second, first + second;
        if second <= max then
            return query
                select *
                from fib(second, first + second, max);
        end if;
    end' language plpgsql;

SELECT *
FROM fib(1,1, 13);