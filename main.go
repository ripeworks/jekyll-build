package main

import (
  "net/http"
  "os"
  "os/exec"
  "fmt"
  "strings"
)

/*
 * POST /build/:host/:user/:name
 */
func JekyllBuild(rw http.ResponseWriter, r *http.Request) {
  // webhook must be POST
  if r.Method != "POST" {
    rw.WriteHeader(http.StatusNotFound)
    return
  }

  url := strings.Split(r.URL.Path, "/")
  host, user, name := url[2], url[3], url[4]
  tmp := "/tmp"

  dir  := tmp + "/src/" + name
  dest := tmp + "/build/" + name
  repo := "git@" + host + ":" + user + "/" + name + ".git"

  cmd := []string{
    "git clone %[2]s %[1]s",
    "cd %[1]s",
    // "git checkout master",
    "jekyll build -s %[1]s -d %[3]s",
    "rm -Rf %[1]s"}

  fmt.Println("-----> Cloning " + repo)
  fmt.Println("-----> Building Jekyll site ...")
  out, err := exec.Command("sh", "-c", fmt.Sprintf(strings.Join(cmd, " && "), dir, repo, dest)).Output()
  if err != nil {
    fmt.Printf("%s", err)
    rw.WriteHeader(http.StatusInternalServerError)
    return
  }
  fmt.Printf("%s", out)
  fmt.Println("-----> Jekyll site built successfully.")

  status, message := JekyllPublish(dest)
  fmt.Println(message)

  rw.WriteHeader(status)
}

func JekyllPublish(dir string) (int, string) {

  // Find S3 Bucket location from the BUCKET file
  bucket, err := exec.Command("sh", "-c", fmt.Sprintf("cd %s && cat BUCKET", dir)).Output()
  if err != nil {
    return 500, "No bucket specified"
  }

  // Sync files to S3
  fmt.Printf("-----> Publishing to Amazon S3 Bucket %s...\n", bucket)
  out, err := exec.Command("s3cmd", "sync", dir + "/", "s3://" + string(bucket), "--delete-removed", "--acl-public").Output()
  fmt.Printf("%s", out)
  if err != nil {
    fmt.Printf("%s", err)
    return 500, "Problem syncing"
  }

  // TODO Optional - Use rsync to remote web server
  // rsync -avz --delete _site/ location:/var/www/html/jekyll

  // remove build
  os.Remove(dir)

  return 200, "Success"
}

func main() {
  http.HandleFunc("/build/", JekyllBuild)
  port := os.Getenv("PORT")
  if port == "" {
      port = "8080"
  }
  http.ListenAndServe(":"+port, nil)
}

