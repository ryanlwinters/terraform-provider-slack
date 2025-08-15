package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/ryanlwinters/terraform-provider-slack/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}