package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/hashicorp/terraform-provider-tfe/tfe"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: tfe.Provider})
}
