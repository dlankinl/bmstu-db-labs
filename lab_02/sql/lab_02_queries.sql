/* 1. Инструкция SELECT, использующая предикат сравнения */
SELECT first_name, last_name, age, net_worth FROM lab.enterpreneurs
WHERE net_worth > 8013012
  AND married = false
  AND gender = true
ORDER BY net_worth desc;

/* 2. Инструкция SELECT, использующая предикат BETWEEN */
SELECT first_name, last_name, age, net_worth FROM lab.enterpreneurs
WHERE net_worth BETWEEN 5013012 AND 8013012
  AND married = false
  AND gender = true
ORDER BY net_worth desc;

/* 3. Инструкция SELECT, использующая предикат LIKE */
SELECT first_name, last_name, age, net_worth FROM lab.enterpreneurs
WHERE net_worth > 8013012
  AND married = false
  AND first_name LIKE '%am'
ORDER BY net_worth desc;

/* 4. Инструкция SELECT, использующая предикат IN со вложенным подзапросом */
SELECT first_name, last_name, age, net_worth FROM lab.enterpreneurs
WHERE net_worth > 8013012
  AND married = false
  AND first_name IN(SELECT first_name
                    FROM lab.enterpreneurs
                    WHERE gender = true)
ORDER BY net_worth desc;

/* 5. Инструкция SELECT, использующая предикат EXISTS со вложенным подзапросом */
SELECT * from lab.companies
WHERE NOT EXISTS(SELECT * FROM lab.financials
                 WHERE lab.financials.id = lab.companies.id);

/* 6. Инструкция SELECT, использующая предикат сравнения с квантором */
SELECT id, profit FROM lab.financials
WHERE profit > ALL(SELECT taxes
                   FROM lab.financials
                   WHERE taxes > 100000)
ORDER by profit;

/* 7. Инструкция SELECT, использующая агрегатные функции в выражениях столбцов */
SELECT avg_profit AS avg_profit,
       sum_profit AS sum_profit
FROM (
         SELECT AVG(profit) as avg_profit,
                SUM(profit) as sum_profit
         FROM lab.financials
     ) AS AVGS_PROFIT;

/* 8. Инструкция SELECT, использующая скалярные подзапросы в выражениях столбцов */


/* 9. Инструкция SELECT, использующая простое выражение CASE */
SELECT lab.companies.name, lab.financials.revenue, lab.financials.profit,
       CASE
           WHEN lab.financials.profit < 100000 THEN 'Low profit'
           WHEN lab.financials.profit BETWEEN 100000 AND 300000 THEN 'Mid profit'
           WHEN lab.financials.profit BETWEEN 300001 AND 600000 THEN 'High profit'
           WHEN lab.financials.profit > 600000 THEN 'Extra high profit'
       END AS Mark
FROM lab.financials JOIN lab.companies ON financials.id = companies.financials_id;

/* 10. Инструкция SELECT, использующая поисковое выражение CASE */
SELECT lab.enterpreneurs.first_name, lab.enterpreneurs.last_name, lab.enterpreneurs.net_worth, companies_res.profit
    FROM (SELECT lab.companies.owner_id, lab.financials.profit,
               CASE
                   WHEN lab.financials.profit < 100000 THEN 'Low profit'
                   WHEN lab.financials.profit BETWEEN 100000 AND 300000 THEN 'Mid profit'
                   WHEN lab.financials.profit BETWEEN 300001 AND 600000 THEN 'High profit'
                   WHEN lab.financials.profit > 600000 THEN 'Extra high profit'
                   END
        FROM lab.financials JOIN lab.companies ON lab.financials.id = lab.companies.id) AS companies_res
        JOIN lab.enterpreneurs ON companies_res.owner_id = lab.enterpreneurs.id ORDER BY net_worth desc;


/* 11. Создание новой временной локальной таблицы из резальтирующего набора данных инструкции SELECT */
SELECT lab.enterpreneurs.age
INTO temp_best
FROM lab.enterpreneurs
WHERE lab.enterpreneurs.net_worth > 10000000
GROUP BY lab.enterpreneurs.age;

SELECT * FROM temp_best;

/* 12. Инструкция SELECT, использующая вложенные коррелированные подзапросы в качестве производных таблиц в предложении FROM */
SELECT f_n, l_n, age, net_worth FROM
    (SELECT lab.enterpreneurs.first_name AS f_n, lab.enterpreneurs.last_name AS l_n, lab.enterpreneurs.age, lab.enterpreneurs.net_worth
     FROM lab.enterpreneurs e JOIN lab.companies c JOIN lab.companies c on e.id = c.owner_id
     WHERE (net_worth > (SELECT AVG(net_worth)
                         FROM lab.enterpreneurs e JOIN lab.companies c ON e.id = c.owner_id
                         WHERE e.id = e.id)))
        AS innerTable;

SELECT f_n, l_n, age, net_worth from
    (SELECT lab.enterpreneurs.first_name as f_n, lab.enterpreneurs.last_name as l_n, lab.enterpreneurs.age, lab.enterpreneurs.net_worth
     from lab.enterpreneurs e inner join lab.companies c on c.owner_id = e.id
     where (net_worth >
            (SELECT avg(net_worth)
             from lab.enterpreneurs e inner join lab.companies c on c.owner_id = e.id
             where e.first_name = e.first_name )))
        as inner_table;

/* 13. Инстркуция SELECT, использующая вложенные подзапросы с уровнем вложенности 3 */
/*/1* */
/* */

/* 14. Инстркуция SELECT, консолидирующая данные с помощью предложения GROUP BY, но без предложения HAVING */
SELECT married, avg(net_worth) AS avg_net_worth
    FROM lab.enterpreneurs GROUP BY married;

/* 15. Инстркуция SELECT, консолидирующая данные с помощью предложения GROUP BY и предложения HAVING */
SELECT married, avg(net_worth) as avg_net_worth
    FROM lab.enterpreneurs GROUP BY married
    HAVING avg(net_worth) < (SELECT avg(net_worth) FROM lab.enterpreneurs);

/* 16. Однострочная INSERT, выполняющая вставку в таблицу одной строки значений */
INSERT INTO lab.skilldescription (description) VALUES ('That`s and amazing skill, bro!');
SELECT * FROM lab.skilldescription ORDER BY skill_id desc LIMIT 10;

/* 17. Многострочная инструкция INSERT, выполняющая вставку в таблицу результирующего набора данных вложенного подзапроса */
CREATE TEMPORARY TABLE IF NOT EXISTS the_richest(
    id serial,
    first_name text,
    last_name text,
    net_worth int
);

ALTER TABLE the_richest ADD company_name text;

INSERT INTO the_richest (first_name, last_name, net_worth, company_name)
    SELECT e.first_name, e.last_name, e.net_worth, c.name as name
        FROM lab.enterpreneurs e JOIN lab.companies c ON e.id = c.owner_id;

SELECT * FROM the_richest;

/* 18. Простая инструкция UPDATE */
UPDATE lab.enterpreneurs
    SET
        married = true
    WHERE age > 50;

SELECT * FROM lab.enterpreneurs
    WHERE age > 50;

/* 19. Инструкция UPDATE со скалярным подзапросом в предложении SET */
UPDATE lab.financials
    SET
        taxes = (SELECT AVG(Taxes)
                 FROM lab.financials
                 WHERE lab.financials.taxes BETWEEN 10000 AND 100000)
    WHERE lab.financials.taxes BETWEEN 10000 AND 100000;

SELECT * FROM lab.financials;

/* 20. Простая инструкция DELETE */
DELETE FROM lab.skilldescription
WHERE skill_id > 1500;

/* 21. Инструкция DELETE со вложенным коррелированным подзапросом в предложении WHERE */
DELETE FROM lab.companies
    WHERE owner_id IN (SELECT id FROM lab.enterpreneurs
                                 WHERE net_worth < 10600000);

SELECT id FROM lab.enterpreneurs WHERE net_worth < 10600000;
SELECT * FROM lab.enterpreneurs WHERE net_worth < 10600000;

/* 22. Инструкция SELECT, использующая простое обобщённое табличное выражение */
WITH BestEnterpreneurs (id, first_name, last_name, net_worth)
         AS (
        SELECT id, first_name, last_name, net_worth
        FROM lab.enterpreneurs
        GROUP BY id
        HAVING net_worth BETWEEN 75600000 AND 90000000
    )
SELECT first_name, last_name, net_worth AS MaxNetWorth
FROM lab.companies c JOIN BestEnterpreneurs be
    ON c.owner_id = be.id;

/* 23. Инструкция SELECT, использующая рекурсивное обобщённое табличное выражение */
/*/1* */
/* */

/* 24. Оконные функции. Использование конструкция MIN/MAX/AVG/OVER() */
WITH BestCompanies (id, name, owner_id, financials_id) AS (
        SELECT id, name, owner_id, financials_id
        FROM lab.companies
        GROUP BY id
    )
SELECT name, revenue, MAX(profit) OVER (PARTITION BY name) AS Profit
FROM lab.financials f JOIN BestCompanies bc
    ON f.id = bc.financials_id
    WHERE revenue > 5000000 AND taxes < 32432;


/* 25. Оконные функции для устранения дублей */
/*/1* */
/* */
