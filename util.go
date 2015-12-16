package main

import (
  "log"
  "errors"
)

func info(msg string) {
  log.Println("-----> " + msg)
}

func fail(msg string, err error) error {
  log.Printf("%s\n", err)
  return errors.New(msg)
}
