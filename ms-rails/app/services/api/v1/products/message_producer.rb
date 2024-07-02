module Api
  module V1
    module Products
      class MessageProducer
        def self.send_product_update(product)
          message = {
            id: product.id,
            name: product.name,
            price: product.price,
            stock: product.stock
          }.to_json

          # código para enviar a mensagem para o Kafka
          KafkaProducer.produce(message, topic: 'rails-to-go')
        end
      end
    end
  end
end
