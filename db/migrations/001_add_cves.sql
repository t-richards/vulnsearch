-- +micrate Up
CREATE TABLE IF NOT EXISTS cves (
  id            VARCHAR(64) PRIMARY KEY,
  description   TEXT        NOT NULL,
  cwe_id        VARCHAR(32) NOT NULL,
  cvss_v2_score REAL        NOT NULL,
  cvss_v3_score REAL        NOT NULL,
  published     DATETIME    NOT NULL,
  last_modified DATETIME    NOT NULL
);

-- +micrate Down
DROP TABLE IF EXISTS cves;
