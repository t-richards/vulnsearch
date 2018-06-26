-- +micrate Up
CREATE TABLE IF NOT EXISTS cves (
  id                   text PRIMARY KEY,
  description          text NOT NULL,
  cwe_id               text,
  severity             text,
  exploitability_score real,
  impact_score         real,
  published            text,
  last_modified        text
);

-- +micrate Down
DROP TABLE IF EXISTS cves;
