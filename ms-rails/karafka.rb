# frozen_string_literal: true

class KarafkaApp < Karafka::App
  setup do |config|
    config.initial_offset = 'latest'
    config.kafka = { 'bootstrap.servers': 'kafka:29092' }
    config.client_id = 'ms-rails'
    config.consumer_persistence = !Rails.env.development?


    config.admin do |admin|
      admin.kafka["allow.auto.create.topics"] = true
    end
  end

  Karafka.monitor.subscribe(Karafka::Instrumentation::LoggerListener.new)

  Karafka.producer.monitor.subscribe(
    WaterDrop::Instrumentation::LoggerListener.new(
      Karafka.logger,
      log_messages: false
    )
  )

  routes.draw do
    topic :'go-to-rails' do
      consumer ProductConsumer
    end
  end
end
