package main

import (
  "fmt"
  "log"
  "os/exec"
)

/*
 * Publish site to an Amazon S3 bucket.
 * Must provide a BUCKET file with the bucket name and path
 * s3cmd must be present and configured
 */
func PublishAmazon(dir string) error {
  // Find S3 Bucket location from the BUCKET file
  bucket, err := exec.Command("sh", "-c", fmt.Sprintf("cd %s && tr -d '\r\n' < BUCKET", dir)).Output()
  if err != nil {
    return fail("No bucket specified", err)
  }

  // Sync files to S3
  info("Publishing to Amazon S3 Bucket " + string(bucket))
  out, err := exec.Command(
    "s3cmd",
    "sync",
    dir + "/",
    "s3://" + string(bucket),
    "--delete-removed",
    "--acl-public",
    "--add-header=Cache-Control:max-age=60",
  ).CombinedOutput()
  log.Printf("%s\n", out)
  if err != nil {
    return fail("Problem syncing", err)
  }

  return nil
}
