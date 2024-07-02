class Api::V1::ProductsController < ApplicationController
  def index
    default_render json: ::Services::Api::V1::Products::ListAll.new(product_params, request).execute, status: 200
  end

  def show
    default_render json: ::Services::Api::V1::Products::Details.new(product_params, request).execute, status: 200
  end

  def create
    default_render json: ::Services::Api::V1::Products::Upsert.new(product_params.merge({is_api: true}), request).execute, status: 201
  end

  def update
    default_render json: ::Services::Api::V1::Products::Upsert.new(product_params.merge({is_api: true}), request).execute, status: 200
  end

  private
  def product_params
    params&.permit(
      [
        :id,
        :name,
        :brand,
        :price,
        :description,
        :stock
      ]
    )
  end
end
