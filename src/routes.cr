require "kemal"

get "/" do
  render "src/views/home.html.ecr", "src/views/layouts/application.html.ecr"
end

get "/product/:id" do |env|
  id = env.params.url["id"].to_i
  product = ::Product.find(id)
  full_name = "#{product.vendor} #{product.name} #{product.version}"
  cves = ::Cve.find_by_product_id(id)

  render "src/views/product.html.ecr", "src/views/layouts/application.html.ecr"
end

post "/product" do |env|
  products = ::Product.names(
    vendor: env.params.json["vendor"].as(String),
    name: env.params.json["name"].as(String)
  )

  { products: products }.to_json
end

post "/search" do |env|
  product = ::Product.find_by_parts(
    vendor: env.params.json["vendor"].as(String),
    name: env.params.json["name"].as(String),
    version: env.params.json["version"].as(String)
  )

  env.redirect "/product/#{product.id}"
end

post "/vendor" do |env|
  query = env.params.url["query"]
  vendors = ::Product.vendors(query)

  { vendors: vendors }.to_json
end

post "/version" do |env|
  versions = ::Product.versions(
    vendor: env.params.json["vendor"].as(String),
    name: env.params.json["name"].as(String),
    version: env.params.json["version"].as(String)
  )

  { versions: versions }.to_json
end
