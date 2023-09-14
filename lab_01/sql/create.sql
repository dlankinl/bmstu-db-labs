-- CREATE DATABASE companiesdb;

DROP SCHEMA IF EXISTS lab CASCADE;

CREATE SCHEMA lab;

CREATE TABLE lab.Enterpreneurs(
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    age INT NOT NULL,
    gender BOOLEAN NOT NULL,
    married BOOLEAN NOT NULL,
    net_worth INT NOT NULL,
    birth_date DATE
);

CREATE TABLE lab.Cities(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL DEFAULT 'Moscow',
    population INT NOT NULL
);

CREATE TABLE lab.Financials(
    id SERIAL PRIMARY KEY,
    revenue DECIMAL(15, 2),
    profit DECIMAL(15, 2),
    total_assets DECIMAL(15, 2),
    taxes DECIMAL(15, 2)
);

CREATE TABLE lab.EnterpreneurSkill(
    enterpreneur_id INT,
    name TEXT NOT NULL,
    skill_id INT
);

CREATE TABLE lab.SkillDescription(
    skill_id SERIAL PRIMARY KEY,
    description TEXT
);

CREATE TABLE lab.Companies(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    owner_id INT,
    city_id INT,
    financials_id INT
);
