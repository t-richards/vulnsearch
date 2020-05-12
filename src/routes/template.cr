def render_with_layout(io, &block)
  ECR.embed("#{__DIR__}/../layouts/application.html.ecr", io)
end

module Onyx::HTTP::View
  macro html_template(template)
    define_type_renderer(render_to_text_html, "text/html", {"text/html"}) do
      render_with_layout(io) do
        ECR.embed("#{__DIR__}/#{{{template}}}", io)
        nil
      end
    end
  end
end
