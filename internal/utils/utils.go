package utils

import (
	"fmt"
	"os/exec"

	"github.com/koyo-os/murocami/pkg/logger"
)

func CloneRepo(url,dir string, logger *logger.Logger) error {
	logger.Infof("cloning repo, url: %s", url)

	cmd := exec.Command("git", "clone", url, dir)
	output,err := cmd.CombinedOutput()
	if err != nil{
		return fmt.Errorf("cant clone repo: %v, output: %s", err, output)
	}

	return nil
}

func DockerBuild(tag, path, tempDir string, logger *logger.Logger) error {
	command := fmt.Sprintf("docker build -t %s %s", tag, path)

	cmd := exec.Command(command)
	cmd.Dir = tempDir
	output, err := cmd.CombinedOutput()
	if err != nil{
		return fmt.Errorf("cant build docker: %v, output %s", err, output)
	}

	return nil
}