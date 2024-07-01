class AddStockToProducts < ActiveRecord::Migration[7.1]
  def change
    add_column :products, :stock, :integer
  end
end
