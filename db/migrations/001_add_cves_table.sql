-- +micrate Up
CREATE TABLE IF NOT EXISTS cves (
  id                   text PRIMARY KEY,
  summary              text NOT NULL DEFAULT "",
  cwe_id               text NOT NULL DEFAULT "",
  vendor               text NOT NULL DEFAULT "",
  product              text NOT NULL DEFAULT "",
  severity             text NOT NULL DEFAULT "",
  exploit_score        real NOT NULL DEFAULT 0,
  impact_score         real NOT NULL DEFAULT 0,
  published            text NOT NULL DEFAULT "",
  last_modified        text NOT NULL DEFAULT ""
);

-- +micrate Down
DROP TABLE IF EXISTS cves;
