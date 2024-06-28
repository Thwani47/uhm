package cmd

import (
	"fmt"
	"os/exec"

	"github.com/fatih/color"
)

func addAzureCommand() {
	// az webapp deploy -g derivco-dictionary-rg --name derivcodictionary --src-path web.zipç
	// check if azure cli is installed
	_, err := exec.Command("az", "--version").Output()

	if err != nil {
		color.Yellow(fmt.Sprintf("[WARN] error confirming azure cli installation: %v\n", err.Error()))
		return
	}

	// TODO get commands from azure
	// TODO search for RGs, and allow user to preselect values fetched from Azure
	// TODO allow user to create new RG if DNE and create a command to create the RG
	// TODO do we include all the commands? No. Maybe allow the user to create custom commands? How in the world would we do this???

}
