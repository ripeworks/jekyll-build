package main

import (
  "fmt"
  "log"
  "path"
  "os/exec"
)

/*
 * Publish to remove server with rsync
 * Must provide remote destination in CNAME file
 */
func PublishRsync(dir string) error {
  info("Publish using rsync")

  remote, err := exec.Command("sh", "-c", fmt.Sprintf("cd %s && tr -d '\r\n' < CNAME", dir)).Output()
  if err != nil {
    return fail("No location provided in CNAME", err)
  }

  // Sync files
  info("Syncing files to remote location " + string(remote))
  out, err := exec.Command(
    "rsync",
    "-avz",
    "--delete",
    path.Join(dir) + "/",
    string(remote),
  ).CombinedOutput()
  log.Printf("%s\n", out)
  if err != nil {
    return fail("Problem syncing", err)
  }

  return nil
}
