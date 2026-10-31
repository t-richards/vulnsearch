-- name: UpsertCVE :exec
INSERT INTO cves (id, description, cwe_ids, cvss_v4_score, cvss_v3_score, cvss_v2_score, published, last_modified)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
ON CONFLICT (id) DO UPDATE SET
    description   = excluded.description,
    cwe_ids       = excluded.cwe_ids,
    cvss_v4_score = excluded.cvss_v4_score,
    cvss_v3_score = excluded.cvss_v3_score,
    cvss_v2_score = excluded.cvss_v2_score,
    published     = excluded.published,
    last_modified = excluded.last_modified;

-- name: DeleteCVE :exec
DELETE FROM cves WHERE id = ?;

-- name: DeleteMatchesForCVE :exec
DELETE FROM cpe_matches WHERE cve_id = ?;

-- name: CreateMatch :exec
INSERT INTO cpe_matches (
    cve_id, vendor, product, version,
    version_start_including, version_start_excluding,
    version_end_including, version_end_excluding
) VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: ClearProducts :exec
DELETE FROM products;

-- name: ClearProductVersions :exec
DELETE FROM product_versions;

-- name: RebuildProducts :exec
INSERT INTO products (vendor, product)
SELECT DISTINCT vendor, product FROM cpe_matches;

-- name: RebuildProductVersions :exec
-- Exact versions and range bounds are the only concrete versions the CVE feeds
-- mention, so they make up the version suggestions.
INSERT OR IGNORE INTO product_versions (vendor, product, version)
SELECT vendor, product, version FROM cpe_matches WHERE version NOT IN ('*', '-', '')
UNION SELECT vendor, product, version_start_including FROM cpe_matches WHERE version_start_including != ''
UNION SELECT vendor, product, version_start_excluding FROM cpe_matches WHERE version_start_excluding != ''
UNION SELECT vendor, product, version_end_including FROM cpe_matches WHERE version_end_including != ''
UNION SELECT vendor, product, version_end_excluding FROM cpe_matches WHERE version_end_excluding != '';

-- name: SearchVendors :many
SELECT DISTINCT vendor FROM products
WHERE vendor LIKE ? ESCAPE '\'
ORDER BY vendor
LIMIT 100;

-- name: SearchProducts :many
SELECT product FROM products
WHERE vendor = ? AND (product LIKE ? ESCAPE '\')
ORDER BY product
LIMIT 100;

-- name: SearchVersions :many
SELECT version FROM product_versions
WHERE vendor = ? AND product = ? AND (version LIKE ? ESCAPE '\');

-- name: ListMatchesForProduct :many
SELECT sqlc.embed(cves), sqlc.embed(cpe_matches)
FROM cpe_matches
JOIN cves ON cves.id = cpe_matches.cve_id
WHERE cpe_matches.vendor = @vendor AND cpe_matches.product = @product
ORDER BY cves.published DESC, cves.id;
