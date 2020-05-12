require "sass"

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
