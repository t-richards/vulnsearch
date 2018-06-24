struct Actions::Home::Index < Prism::Action
  include Params

  params do
    param :q, String?
  end

  def call
    query = params[:q]
    render_default "home/index"
  end
end
