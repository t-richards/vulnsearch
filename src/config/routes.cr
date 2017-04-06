require "kemal"
require "../controllers/*"

get "/" { |env| Vulnsearch::HomeController.index(env) }
get "/search" { |env| Vulnsearch::HomeController.search(env) }
