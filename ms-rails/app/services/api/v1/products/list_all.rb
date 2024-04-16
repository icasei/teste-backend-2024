module Services
  module Api 
    module V1
      module Products
        class ListAll
          attr_accessor :params, :request

          def initialize(params, request)
            @params = params
            @request = request
          end

          def execute
            list_all
          end

          private
          def list_all
            @list_all ||= Product.all
          end
        end
      end
    end
  end
end