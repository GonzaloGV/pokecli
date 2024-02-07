package main

import (
	"os"
)

func commandExit(c *context, args ...string) error {
	os.Exit(0)

	return nil
}
