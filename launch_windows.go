package main

import (
	"os/exec"
)

func launch(url string) error {
	return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Run()
}
