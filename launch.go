// +build !darwin,!linux,!windows

package main

import (
	"fmt"
)

func launch(url string) error {
	_, err := fmt.Println("Found:", url)
	return err
}
