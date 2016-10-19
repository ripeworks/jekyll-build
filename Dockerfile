FROM ruby:2.3.0
MAINTAINER mike@ripeworks.com
ENV LANG C.UTF-8

RUN apt-get update -qq && apt-get install -y build-essential s3cmd

# install nodejs
RUN curl -sL https://deb.nodesource.com/setup_6.x | bash -
RUN apt-get install -y nodejs

# install gems
RUN gem install bundler github-pages jekyll-fridge

# install deployment method (surge)
RUN npm install -g surge

# Disable strict host checking
RUN mkdir /root/.ssh && echo "StrictHostKeyChecking no " > /root/.ssh/config

WORKDIR /app
COPY jekyll-build /app/
EXPOSE 80
ENTRYPOINT ["./jekyll-build"]
