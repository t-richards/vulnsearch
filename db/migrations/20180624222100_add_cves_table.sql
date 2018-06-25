-- +micrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS cves (id, description, cwe_id, severity, exploitability_score, impact_score, published, last_modified);

-- +micrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS cves;
