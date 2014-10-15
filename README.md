jekyll-build
============

Jekyll build server written in go. Uses a `POST` http request to trigger builds, making it useful for webhooks.

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
* `foreman start` or `forego start`
* Configure nginx to point to it.

Usage
----

Send a `POST` request to the web app at the proper path to trigger a build.
