module Vulnsearch
  class DeinitHandler < Kemal::Handler
    def call(context)
      context.response.headers.delete "X-Powered-By"
      call_next context
    end
  end
end

add_handler Vulnsearch::DeinitHandler.new
