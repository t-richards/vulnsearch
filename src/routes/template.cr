require "sass"

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

def compile_scss(path)
  css = Sass.compile_file(
    path,
    include_path: "node_modules/bootstrap/scss",
    output_style: Sass::OutputStyle::EXPANDED,
    precision: 8,
    output_path: "/dist",
    source_map_file: "/dist/application.css.map"
  )
  base = File.basename(path, ".scss")
  out_path = File.join("public/dist", base) + ".css"
  File.write(out_path, css)

  css
end
