package main

import (
	"os/exec"
)

func launch(url string) error {
	return exec.Command("open", url).Run()
}
