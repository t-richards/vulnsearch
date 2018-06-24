require "ecr/macros"

macro render_string(filename)
  io = IO::Memory.new
  ECR.embed({{filename}}, io)
  io.to_s
end

macro render(filename)
  ECR.embed({{filename}}, context.response)
end

macro render(filename, layout)
  __content_filename__ = {{filename}}
  content = render_string {{filename}}
  render {{layout}}
end

macro render_default(filename)
  render "src/views/#{{{filename}}}.ecr", "src/views/layouts/default.ecr"
end
