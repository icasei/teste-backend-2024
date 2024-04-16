module Services
  module Api 
    module V1
      module Products
        class Details
          attr_accessor :params, :request

          def initialize(params, request)
            @params = params
            @request = request
          end

          def execute
            raise ActiveRecord::RecordInvalid, "Missing parameter"  if params[:id].to_i.zero?

            raise ActiveRecord::RecordNotFound, "Product Not Found" if product.blank?

            product
          end

          private
          def product
            @product ||= Product.find_by(id: params[:id].to_i)
          end
        end
      end
    end
  end
end