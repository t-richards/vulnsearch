-- CVEs from the NVD JSON 2.0 feeds. Scores are NULL when NVD has no metric for
-- that CVSS version, which is distinct from a score of 0.0.
CREATE TABLE IF NOT EXISTS cves (
    id            TEXT     NOT NULL PRIMARY KEY,
    description   TEXT     NOT NULL,
    cwe_ids       TEXT     NOT NULL,
    cvss_v4_score REAL,
    cvss_v3_score REAL,
    cvss_v2_score REAL,
    published     DATETIME NOT NULL,
    last_modified DATETIME NOT NULL
);

-- One row per vulnerable CPE match criterion of a CVE. A criterion either names
-- an exact version, or a wildcard version ("*") optionally bounded by a range.
-- Ranges are evaluated at query time, so any version can be looked up, even
-- one that NVD has never listed.
CREATE TABLE IF NOT EXISTS cpe_matches (
    cve_id                  TEXT NOT NULL REFERENCES cves (id) ON DELETE CASCADE,
    vendor                  TEXT NOT NULL COLLATE NOCASE,
    product                 TEXT NOT NULL COLLATE NOCASE,
    version                 TEXT NOT NULL,
    version_start_including TEXT NOT NULL DEFAULT '',
    version_start_excluding TEXT NOT NULL DEFAULT '',
    version_end_including   TEXT NOT NULL DEFAULT '',
    version_end_excluding   TEXT NOT NULL DEFAULT ''
);

CREATE INDEX IF NOT EXISTS cpe_matches_vendor_product ON cpe_matches (vendor, product);
CREATE INDEX IF NOT EXISTS cpe_matches_cve_id ON cpe_matches (cve_id);

-- Search suggestions, derived from cpe_matches after each load.
CREATE TABLE IF NOT EXISTS products (
    vendor  TEXT NOT NULL COLLATE NOCASE,
    product TEXT NOT NULL COLLATE NOCASE,

    PRIMARY KEY (vendor, product)
) WITHOUT ROWID;

CREATE TABLE IF NOT EXISTS product_versions (
    vendor  TEXT NOT NULL COLLATE NOCASE,
    product TEXT NOT NULL COLLATE NOCASE,
    version TEXT NOT NULL,

    PRIMARY KEY (vendor, product, version)
) WITHOUT ROWID;
