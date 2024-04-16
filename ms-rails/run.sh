#! /bin/ash
export RAILS_MASTER_KEY=2c8e51116b8f319c236cbbe890720420
export NEW_RELIC_AGENT_ENABLED=false 
export RAILS_ENV=development 
export RAILS_MAX_THREADS=5
bundle exec puma -C config/puma.rb
