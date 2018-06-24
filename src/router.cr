class Router
  def self.new
    Prism::Router.new do
      get "/", Actions::Home::Index
      get "/search/:query", Actions::Search::Index
    end
  end
end
