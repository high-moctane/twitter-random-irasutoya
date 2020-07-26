FROM ruby:2.7.1-buster

COPY [ "./Gemfile", "/twitter-random-irasutoya/Gemfile" ]
COPY [ "./Gemfile.lock", "/twitter-random-irasutoya/Gemfile.lock" ]

WORKDIR /twitter-random-irasutoya

RUN bundle install

COPY [ ".", "/twitter-random-irasutoya" ]

CMD [ "ruby", "main.rb" ]
