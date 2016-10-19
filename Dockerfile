FROM ruby:2.3.0
MAINTAINER mike@ripeworks.com

RUN apt-get update -qq && apt-get install -y build-essential s3cmd

# install nodejs
RUN curl -sL https://deb.nodesource.com/setup_6.x | bash -
RUN apt-get install -y nodejs

# install gems
RUN gem install bundler therubyracer github-pages jekyll-fridge

# install deployment method (surge)
RUN npm install -g surge

WORKDIR /app
COPY jekyll-build /app/
EXPOSE 8080
ENTRYPOINT ["./jekyll-build"]
