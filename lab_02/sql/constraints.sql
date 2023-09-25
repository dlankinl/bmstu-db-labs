ALTER TABLE lab.Financials
    ADD CONSTRAINT check_revenue_positive CHECK (revenue >= 0);
ALTER TABLE lab.Financials
    ADD CONSTRAINT check_total_assets_positive CHECK (total_assets >= 0);
ALTER TABLE lab.Financials
    ADD CONSTRAINT check_taxes_positive CHECK (taxes >= 0);

ALTER TABLE lab.Cities
    ADD CONSTRAINT check_population_positive CHECK (population >= 0);

ALTER TABLE lab.Enterpreneurs
    ADD CONSTRAINT check_age_positive CHECK (age >= 0);
ALTER TABLE lab.Enterpreneurs
    ADD CONSTRAINT check_net_worth_positive CHECK (net_worth >= 0);
ALTER TABLE lab.Enterpreneurs
    ADD CONSTRAINT check_birth_date CHECK (birth_date >= '1800-01-01'::date AND birth_date <= current_date);

ALTER TABLE lab.Companies
    ADD CONSTRAINT fk_owner FOREIGN KEY (owner_id) REFERENCES lab.Enterpreneurs(id);
ALTER TABLE lab.Companies
    ADD CONSTRAINT fk_loc FOREIGN KEY (city_id) REFERENCES lab.Cities(id);
ALTER TABLE lab.Companies
    ADD CONSTRAINT fk_fin FOREIGN KEY (financials_id) REFERENCES lab.Financials(id);

ALTER TABLE lab.EnterpreneurSkill
    ADD CONSTRAINT fk_enterpreneur FOREIGN KEY (enterpreneur_id) REFERENCES lab.Enterpreneurs(id);
ALTER TABLE lab.EnterpreneurSkill
    ADD CONSTRAINT fk_skill FOREIGN KEY (skill_id) REFERENCES lab.SkillDescription(skill_id);
