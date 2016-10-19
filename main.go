package main

import (
  "net/http"
  "os"
  "os/exec"
  "log"
  "fmt"
  "strings"
)

const (
  Route = "/build/"
  DefaultPort = "80"
)

var queue = make(chan string, 100)

func JekyllBuild(path string) {
  url := strings.Split(path, "/")
  host, user, name := url[2], url[3], url[4]
  tmp := "/tmp"

  dir  := tmp + "/src/" + name
  dest := tmp + "/build/" + name
  repo := "git@" + host + ":" + user + "/" + name + ".git"
  // repo := "https://" + host + "/" + user + "/" + name + ".git"

  cmd := []string{
    "git clone %[2]s %[1]s &&",
    "cd %[1]s;",
    // "git checkout master",
    "[ -f Gemfile ] && bundle install;",
    "jekyll build -s %[1]s -d %[3]s;",
    "rm -Rf %[1]s;"}

  info("Cloning " + repo)
  info("Building Jekyll site ...")
  out, err := exec.Command("sh", "-c", fmt.Sprintf(strings.Join(cmd, " "), dir, repo, dest)).Output()
  log.Printf("%s\n", out)
  if err != nil {
    fail("Jekyll Error", err)
    return
  }
  info("Jekyll site built successfully.")

  JekyllPublish(dest)
}

func JekyllPublish(dir string) {

  // Determine deployment method from DEPLOY file
  out, err := exec.Command("sh", "-c", fmt.Sprintf("cd %s && tr -d '\r\n' < DEPLOY", dir)).Output()
  method := string(out)

  if err != nil {
    method = "amazon" // default to Amazon S3
  }

  if method == "amazon" {
    err = PublishAmazon(dir)
  }

  if method == "surge" {
    err = PublishSurge(dir)
  }

  // if method == "rsync"
  // TODO Optional - Use rsync to remote web server
  // rsync -avz --delete _site/ location:/var/www/html/jekyll

  // remove build
  os.RemoveAll(dir)

  if err == nil {
    info("Success! Jekyll site published.")
  }
}

/*
 * POST /build/:host/:user/:name
 */
func routeHandler(rw http.ResponseWriter, r *http.Request) {
  // must be POST
  if r.Method != "POST" {
    rw.WriteHeader(http.StatusNotFound)
    return
  }

  queue <- r.URL.Path

  rw.WriteHeader(http.StatusCreated)
  return
}

// listen for url paths from queue channel
// and process builds one at a time.
func workOnQueue() {
  go func() {
    for {
      select {
      case path := <-queue:
        info("Processing POST " + path)
        JekyllBuild(path)
      }
    }
  }()
}

func main() {
  // start the queue
  workOnQueue()

  // listen for requests
  http.HandleFunc(Route, routeHandler)
  port := os.Getenv("PORT")
  if port == "" { port = DefaultPort }

  http.ListenAndServe(":"+port, nil)
}

