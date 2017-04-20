-- +micrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE VIRTUAL TABLE cves USING fts5(id, summary, cwe_id, published, last_modified);

-- +micrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE cves;
