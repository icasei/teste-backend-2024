class ApplicationController < ActionController::API

  def default_render(json: {}, status: nil)
    status = json[:status].to_i if json.is_a?(Hash) && json[:status].present?

    status = 200 if status.to_i.zero?

    render json: { data: json }, status: status
  end
end
