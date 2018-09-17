-- +micrate Up
CREATE TABLE IF NOT EXISTS cves (
  id            text PRIMARY KEY,
  description   text NOT NULL DEFAULT "",
  cwe_id        text NOT NULL DEFAULT "",
  vendor        text NOT NULL DEFAULT "",
  product       text NOT NULL DEFAULT "",
  cvss_v2_score text NOT NULL DEFAULT "",
  cvss_v3_score text NOT NULL DEFAULT "",
  published     text NOT NULL DEFAULT "",
  last_modified text NOT NULL DEFAULT ""
);

-- +micrate Down
DROP TABLE IF EXISTS cves;
