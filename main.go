package main

import (
  "net/http"
  "os"
)

func JekyllBuild(rw http.ResponseWriter, r *http.Request) {
  // webhook must be POST
  if r.Method != "POST" {
    rw.WriteHeader(http.StatusNotFound)
  }

  // Determine folder
  // Run cmd
  // git clone repo dir
  // jekyll build -s dir -d www_dir
  // rm -Rf dir

  // publish to S3
  // https://github.com/developmentseed/jekyll-hook/blob/master/scripts/publish-s3.sh

}

func main() {
  http.HandleFunc("/fridgehook", JekyllBuild)
  port := os.Getenv("PORT")
  if port == "" {
      port = "8080"
  }
  http.ListenAndServe(":"+port, nil)
}

