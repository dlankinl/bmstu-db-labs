copy lab.Enterpreneurs(first_name, last_name, age, gender, married, net_worth, birth_date)  FROM 'csv/enterpreneurs.csv' DELIMITER ',' CSV HEADER;
copy lab.Financials(revenue, profit, total_assets, taxes)  FROM 'csv/financials.csv' DELIMITER ',' CSV HEADER;
copy lab.Cities(name, population)  FROM 'csv/cities.csv' DELIMITER ',' CSV HEADER;
copy lab.SkillDescription(description) FROM 'csv/skills_descr.csv' DELIMITER ',' CSV HEADER;
copy lab.EnterpreneurSkill(enterpreneur_id, name, skill_id)  FROM 'csv/skills_name.csv' DELIMITER ',' CSV HEADER;
copy lab.Companies(name, owner_id, city_id, financials_id) FROM 'csv/companies.csv' DELIMITER ',' CSV HEADER;