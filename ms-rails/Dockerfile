FROM ruby:3.0.0-alpine3.13

LABEL maintainer="Matheus M<mmacedo@icasei.com.br>"

RUN apk --update --upgrade add \
      bash\
      sqlite-dev\
      build-base\
      tzdata

ENV INSTALL_PATH /ms-rails

RUN mkdir -p $INSTALL_PATH

WORKDIR $INSTALL_PATH

COPY Gemfile $INSTALL_PATH/

RUN bundle install

ADD . $INSTALL_PATH

RUN mkdir -p $INSTALL_PATH/tmp/pids

RUN chmod +x run.sh

CMD ["./run.sh"]