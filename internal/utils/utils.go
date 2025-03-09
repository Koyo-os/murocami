package utils

import (
	"fmt"
	"os/exec"
)

func CloneRepo(url,dir string) error {
	cmd := exec.Command("git", "clone", url, dir)
	output,err := cmd.CombinedOutput()
	if err != nil{
		return fmt.Errorf("cant clone repo: %v, output: %s", err, output)
	}

	return nil
}