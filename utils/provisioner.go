package utils

import (
	"fmt"
	"log"
	"os/exec"
)

// Provisioner is an interface for running infrastructure commands.
type Provisioner interface {
	Init(directory string, args ...string) (string, error)
	Apply(directory string, args ...string) (string, error)
	Output(directory string) (string, error)
	Destroy(directory string, args ...string) (string, error)
}

const TOFUCOMMAND = "tofu"
const TERRAFORMCOMMAND = "terraform"

// TerraformProvisioner implements Provisioner for terraform.
type TerraformProvisioner struct{}

func (p *TerraformProvisioner) Init(directory string, args ...string) (string, error) {
	cmdArgs := append([]string{"init", "-input=false"}, args...)
	return runCommand(TERRAFORMCOMMAND, directory, cmdArgs...)
}

func (p *TerraformProvisioner) Apply(directory string, args ...string) (string, error) {
	cmdArgs := append([]string{"apply", "-input=false", "-auto-approve"}, args...)
	return runCommand(TERRAFORMCOMMAND, directory, cmdArgs...)
}

func (p *TerraformProvisioner) Output(directory string) (string, error) {
	cmdArgs := []string{"output", "-json"}
	return runCommand(TERRAFORMCOMMAND, directory, cmdArgs...)
}

func (p *TerraformProvisioner) Destroy(directory string, args ...string) (string, error) {
	cmdArgs := append([]string{"destroy", "-input=false", "-auto-approve"}, args...)
	return runCommand(TERRAFORMCOMMAND, directory, cmdArgs...)
}

type TofuProvisioner struct{}

func (p *TofuProvisioner) Init(directory string, args ...string) (string, error) {
	cmdArgs := append([]string{"init", "-input=false"}, args...)
	return runCommand(TOFUCOMMAND, directory, cmdArgs...)
}

func (p *TofuProvisioner) Apply(directory string, args ...string) (string, error) {
	cmdArgs := append([]string{"apply", "-input=false", "-auto-approve"}, args...)
	return runCommand(TOFUCOMMAND, directory, cmdArgs...)
}

func (p *TofuProvisioner) Output(directory string) (string, error) {
	cmdArgs := []string{"output", "-json"}
	return runCommand(TOFUCOMMAND, directory, cmdArgs...)
}

func (p *TofuProvisioner) Destroy(directory string, args ...string) (string, error) {
	cmdArgs := append([]string{"destroy", "-input=false", "-auto-approve"}, args...)
	return runCommand(TOFUCOMMAND, directory, cmdArgs...)
}

// runCommand is a helper function to run a command in a given directory with arguments.
func runCommand(cmdName, directory string, args ...string) (string, error) {
	log.Printf("Executing the following command %s on directory %s \n", cmdName, directory)
	cmd := exec.Command(cmdName, args...)
	cmd.Dir = directory
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error running command %s %v: %w - output: %s", cmdName, args, err, string(output))
	}
	return string(output), nil
}
