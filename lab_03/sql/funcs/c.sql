-- 1. Триггер AFTER
select *
into financials_tmp
from lab.financials;

select *
into companies_tmp
from lab.companies;

-- drop table companies_tmp;
-- drop table financials_tmp;

create or replace function update_id()
    returns trigger
as '
    begin
        update companies_tmp
        set
            financials_id = new.id
        where companies_tmp.financials_id = old.id;


        return new;
    end;
' language plpgsql;

create trigger update_trigger
    after update on financials_tmp
    for each row
execute procedure update_id();

update financials_tmp
set id = 3000
where id = 2;

select *
from financials_tmp;

select *
from companies_tmp
order by financials_id desc;

select *
from companies_tmp
where financials_id = 2;

-- 2. Триггер INSTEAD OF
-- Вместо удаления описания оно заменяется на "there is no description:("
create view tmp_descr as
select *
from lab.skilldescription
where skill_id between 1000 and 1400;

create or replace function del_skill_descr_func()
    returns trigger
as '
    begin
        update tmp_descr
        set
            description = ''there is no description:(''
        where tmp_descr.skill_id = old.skill_id;

        return new;
    end;
' language plpgsql;

create trigger del_descr_trigger
    instead of delete on tmp_descr
    for each row
execute procedure del_skill_descr_func();

delete from tmp_descr
where skill_id = 1221;

select *
from tmp_descr
where skill_id = 1221;


-- Если нет предприятий в городе, то необходимо удалить этот город из таблицы.
create or replace function check_city_reference()
    returns trigger
as '
    begin
        delete from lab.cities
        where id in (select c.id
                     from lab.Cities c
                          left join lab.Companies co on c.id = co.city_id
                     where co.id is null);
        return new;
    end;
' language plpgsql;

create trigger check_city_reference_trigger
    before insert or update on lab.companies
    for each row
execute function check_city_reference();

insert into lab.companies(name, owner_id, city_id, financials_id)
values ('Heyor', 1, 1, 1);

insert into lab.cities(name, population)
values ('SPb', 654654);
