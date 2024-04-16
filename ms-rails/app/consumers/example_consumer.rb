# frozen_string_literal: true
class ExampleConsumer < ApplicationConsumer
  def consume
    messages.each { |message| puts message.payload.with_indifferent_access }
  end
end
