#!/bin/sh
/server &
go run app/consumers/kafka_consumer.go