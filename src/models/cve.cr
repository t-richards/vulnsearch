class Cve
  include Core::Schema
  include Core::Query
  include Core::Validation

  schema :cves do
    primary_key :id, String
    field :summary, String
    field :cwe_id, String
    field :published, Time
    field :last_modified, Time
  end
end
