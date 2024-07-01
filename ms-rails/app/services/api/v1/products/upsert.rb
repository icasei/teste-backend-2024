module Services
  module Api 
    module V1
      module Products
        class Upsert
          attr_accessor :params, :request

          def initialize(params, request)
            @params = params
            @request = request
          end

          def execute
            raise ActiveRecord::RecordNotFound, "Product Not Found" if product.blank?

            ActiveRecord::Base.transaction do
              product.id        ||= params[:id]           if params[:id].present?
              product.name        = params[:name]         if params[:name].present?
              product.brand       = params[:brand]        if params[:brand].present?
              product.price       = params[:price]        if params[:price].present?
              product.description = params[:description]  if params[:description].present?
              product.stock       = params[:stock]        if params[:stock].present?

              product.save!

              Karafka.producer.produce_sync(topic: 'rails-to-go', payload: product.to_json) if !!params[:is_api]
            end
            
            product
          end

          private
          def product
            @product ||= Product.find_by(id: params[:id]).presence || Product.new
          end
        end
      end
    end
  end
end