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
SELECT lab.companies.name, lab.financials.revenue, lab.financials.profit,
CASE
    WHEN lab.financials.profit < 100000 THEN 'Low profit'
    WHEN lab.financials.profit BETWEEN 100000 AND 300000 THEN 'Mid profit'
    WHEN lab.financials.profit BETWEEN 300001 AND 600000 THEN 'High profit'
    WHEN lab.financials.profit > 600000 THEN 'Extra high profit'
END
FROM lab.financials JOIN lab.companies ON financials.id = companies.financials_id;

/* 9. Инструкция SELECT, использующая простое выражение CASE */
SELECT

/* 10. Инструкция SELECT, использующая поисковое выражение CASE */
/*/1* */
/* */

/* 11. Создание новой временной локальной таблицы из резальтирующего набора данных инструкции SELECT */
/*/1* */
/* */

/* 12. Инструкция SELECT, использующая вложенные коррелированные подзапросы в качестве производных таблиц в предложении FROM */
/*/1* */
/* */

/* 13. Инстркуция SELECT, использующая вложенные подзапросы с уровнем вложенности 3 */
/*/1* */
/* */

/* 14. Инстркуция SELECT, консолидирующая данные с помощью предложения GROUP BY, но без предложения HAVING */
/*/1* */
/* */

/* 15. Инстркуция SELECT, консолидирующая данные с помощью предложения GROUP BY и предложения HAVING */
/*/1* */
/* */

/* 16. Однострочная INSERT, выполняющая вставку в таблицу одной строки значений */
/*/1* */
/* */

/* 17. Многострочная инструкция INSERT, выполняющая вставку в таблицу результирующего набора данных вложенного подзапроса */
/*/1* */
/* */

/* 18. Простая инструкция UPDATE */
/*/1* */
/* */

/* 19. Инструкция UPDATE со скалярным подзапросом в предложении SET */
/*/1* */
/* */

/* 20. Простая инструкция DELETE */
/*/1* */
/* */

/* 21. Инструкция DELETE со вложенным коррелированным подзапросом в предложении WHERE */
/*/1* */
/* */

/* 22. Инструкция SELECT, использующая простое обобщённое табличное выражение */
/*/1* */
/* */

/* 23. Инструкция SELECT, использующая рекурсивное обобщённое табличное выражение */
/*/1* */
/* */

/* 24. Оконные функции. Использование конструкция MIN/MAX/AVG/OVER() */
/*/1* */
/* */

/* 25. Оконные функции для устранения дублей */
/*/1* */
/* */
