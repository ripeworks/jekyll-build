jekyll-build
============

Jekyll build server written in go.

Requirements
----

* git
* Ruby
* Jekyll `gem install jekyll`
* rsync _Optional. For publishing to remote server_
* s3cmd _Optional. For publishing to Amazon S3_

Install
----

* Make sure all requirements are installed
* Move `bin/jekyll-build` somewhere
* Run the app using foreman/forego
* Configure nginx to point to it.

Usage
----

Send a `POST` request to the web app at the proper path to trigger a build.
