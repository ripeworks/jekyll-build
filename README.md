jekyll-build
============

Jekyll build server written in go. Uses a `POST` http request to trigger builds, making it useful for webhooks. A simple `Dockerfile` has been provided with all necessary dependencies to run `jekyll build`.

Requirements
----

* Something that can run `jekyll build`. See `Dockerfile`
* git
* Go

Deploy Techniques
-----------------

Using a `DEPLOY` file you can specify how you would like jekyll-build to publish your site. By default it will attempt to use Amazon S3.

* __Amazon S3__
* __surge.sh__

Usage
----

Send a POST request to `/build` to get started. See below for the format to map a request to a repository url.

`POST /build/:host/:user/:repo`

* __host__ - Host where repository is located. (i.e. github.com)
* __user__ - User/Organization that owns repository.
* __repo__ - Name of repository

### Example

```
POST /build/github.com/tamagokun/tamagokun.github.com

# Maps to

git@github.com:tamagokun/tamagokun.github.com.git
```

### S3 Bucket

By default successfully built sites will be synced to an Amazon S3 Bucket. Specify the bucket location by creating a `BUCKET` file at the root of your project.

### Permissions

The user running jekyll-build must have proper permissions to access the repository specified, as well as a configured `s3cmd` client.
