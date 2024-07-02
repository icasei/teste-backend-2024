# frozen_string_literal: true

class ProductConsumer < ApplicationConsumer
    def consume
      messages.each do |message|
        process_message(message.payload.with_indifferent_access)
      end
    end
  
    private
  
    def process_message(params)
      product = Product.find_or_initialize_by(id: params[:id])
  
      if product.new_record? || product.updated_at <= params[:updated_at]
        product.assign_attributes(
          name: params[:name],
          brand: params[:brand],
          price: params[:price],
          stock: params[:stock],
          description: params[:description],
          updated_at: params[:updated_at]
        )
        product.save
      end
    end
  end