package main

import (
  "fmt"
  "log"
  "os/exec"
)

/*
 * Publish site to surge.sh
 * Must have surge installed and set up using SURGE_LOGIN and SURGE_TOKEN
 * Provide CNAME file to specify target domain name
 */
func PublishSurge(dir string) error {
  info("Publish to surge.sh")

  out, err := exec.Command("sh", "-c", fmt.Sprintf("cd %s && surge .", dir)).Output()
  log.Printf("%s\n", out)
  if err != nil {
    return fail("Unable to deploy using surge", err)
  }

  return nil
}
