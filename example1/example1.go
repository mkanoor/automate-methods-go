package main

import (
	"flag"
	"fmt"
	"github.com/mkanoor/manageiq-api-client-go/manageiq/api/client"
	"log"
)

func parseArgs(params *client.ConnectionParameters_t, guid *string) {
	guidPtr := flag.String("guid", "", "Automate Workspace GUID")
	urlPtr := flag.String("url", "http://localhost:4000/api/", "Automate API URL")
	tokenPtr := flag.String("token", "", "Automate Token")
	userPtr := flag.String("username", "admin", "User")
	passwordPtr := flag.String("password", "smartvm", "Password")
	groupPtr := flag.String("group", "", "MIQ Group")

	flag.Parse()
	if len(*guidPtr) == 0 {
		log.Fatal("GUID is a required Parameter")
	}

	params.Username = *userPtr
	params.Password = *passwordPtr
	params.BaseUrl = *urlPtr
	params.MIQToken = *tokenPtr
	params.Group = *groupPtr
	*guid = *guidPtr
	return
}

func main() {
	var params client.ConnectionParameters_t
	var guid string
	parseArgs(&params, &guid)

	workspace := client.NewWorkspace(&params, guid)

	err := workspace.Fetch()
	if err != nil {
		log.Fatal(err)
	}

	root, _ := workspace.GetObject("root")
	malayalam_rune_slice := []rune{0x0D2E, 0x0D27, 0x0D41}
	root.SetAttribute("my_name", string(malayalam_rune_slice))

	vm := root.GetAttribute("vm").(*client.VMDB_Object)

	// Start the VM
	response, err := vm.Action("start", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Response from action", response)
	vm.AddCustomAttribute(root.GetAttribute("cattr").(string), root.GetAttribute("cvalue").(string))

	workspace.Update()
}
