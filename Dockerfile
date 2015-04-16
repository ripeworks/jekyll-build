FROM ubuntu:14.04

MAINTAINER mike@ripeworks.com

RUN apt-get update
RUN locale-gen en_US.UTF-8
ENV LANG en_US.UTF-8
ENV LC_ALL en_US.UTF-8

RUN apt-get -y install build-essential zlib1g-dev git ruby ruby-dev golang s3cmd ImageMagick

RUN gem install bundler therubyracer github-pages jekyll-fridge

WORKDIR /jekyll-build
ADD main.go /jekyll-build/main.go

EXPOSE 8080
CMD go build && ./jekyll-build
