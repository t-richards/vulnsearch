-- +micrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE cves (id TEXT PRIMARY KEY, summary TEXT, published TEXT, last_modified TEXT);

-- +micrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE cves;
