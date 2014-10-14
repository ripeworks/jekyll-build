package main

import (
  "net/http"
  "os"
  "os/exec"
  "fmt"
  "strings"
)

func JekyllBuild(rw http.ResponseWriter, r *http.Request) {
  // webhook must be POST
  if r.Method != "POST" {
    rw.WriteHeader(http.StatusNotFound)
    return
  }

  dir := "/Users/mkruk/Desktop/jekyll_test"   // 1
  repo := "git@github.com:tamagokun/tamagokun.github.com.git"   // 2
  dest := "/Users/mkruk/Desktop/jekyll_build"   // 3

  cmd := []string{
    "git clone %[2]s %[1]s",
    "cd %[1]s",
    // "git checkout master",
    "jekyll build -s %[1]s -d %[3]s",
    "rm -Rf %[1]s"}

  fmt.Println("-----> Building Jekyll site ...")
  out, err := exec.Command("sh", "-c", fmt.Sprintf(strings.Join(cmd, " && "), dir, repo, dest)).Output()
  if err != nil {
    fmt.Printf("%s", err)
    rw.WriteHeader(http.StatusInternalServerError)
    return
  }
  fmt.Printf("%s", out)
  fmt.Println("-----> Jekyll site built successfully.")

  JekyllPublish(dest)

  rw.WriteHeader(http.StatusOK)
}

func JekyllPublish(dir string) {

  // rsync to remote
  // rsync -avz --delete _site/ location:/var/www/html/jekyll

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

