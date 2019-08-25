class CveRange < ApplicationRecord
  DB.mapping({
    severity:    String,
    score_range: String,
    count:       Int32,
  })

  def critical_cves
    query = <<-EOT
      SELECT 'Critical' as severity, '9.0 - 10.0' as score_range, COUNT(*) as count FROM cves
      JOIN cves_products ON cves.id = cves_products.cve_id
      WHERE cves_products.product_id = #{@product_id}
      AND cves.cvss_v3_score >= 9.0
      AND cves.cvss_v3_score <= 10.0;
    EOT
    from_rs(db.query(query))
  end

  def high_cves
    query = <<-EOT
      SELECT 'High' as severity, '7.0 - 8.9' as score_range, COUNT(*) as count FROM cves
      JOIN cves_products ON cves.id = cves_products.cve_id
      WHERE cves_products.product_id = #{@product_id}
      AND cves.cvss_v3_score >= 7.0
      AND cves.cvss_v3_score <= 8.9;
    EOT
    from_rs(db.query(query))
  end

  def medium_cves
    query = <<-EOT
      SELECT 'Medium' as severity, '4.0 - 6.9' as score_range, COUNT(*) as count FROM cves
      JOIN cves_products ON cves.id = cves_products.cve_id
      WHERE cves_products.product_id = #{@product_id}
      AND cves.cvss_v3_score >= 4.0
      AND cves.cvss_v3_score <= 6.9;
    EOT
    from_rs(db.query(query))
  end

  def low_cves
    query = <<-EOT
      SELECT 'Low' as severity, '0.1 - 3.9' as score_range, COUNT(*) as count FROM cves
      JOIN cves_products ON cves.id = cves_products.cve_id
      WHERE cves_products.product_id = #{@product_id}
      AND cves.cvss_v3_score >= 0.1
      AND cves.cvss_v3_score <= 3.9;
    EOT
    from_rs(db.query(query))
  end
end
