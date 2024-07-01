class Product < ApplicationRecord
  validates :name,  length: {minimum:4}
  validates :price, presence: true, format: { with: /\A\d+(?:\.\d{2})?\z/ }, numericality: { greater_than: 0, less_than: 1000000}
  validates :brand, :description, presence: true,
  validates :stock, numericality: { only_integer: true, greater_than_or_equal_to: 0 }
end
