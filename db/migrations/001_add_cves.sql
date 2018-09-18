-- +micrate Up
CREATE TABLE IF NOT EXISTS cves (
  id            VARCHAR(64) PRIMARY KEY,
  description   TEXT        NOT NULL DEFAULT "",
  cwe_id        VARCHAR(32) NOT NULL DEFAULT "",
  cvss_v2_score REAL        NOT NULL DEFAULT 0.0,
  cvss_v3_score REAL        NOT NULL DEFAULT 0.0,
  published     DATETIME    NOT NULL DEFAULT current_timestamp,
  last_modified DATETIME    NOT NULL DEFAULT current_timestamp
);

-- +micrate Down
DROP TABLE IF EXISTS cves;
