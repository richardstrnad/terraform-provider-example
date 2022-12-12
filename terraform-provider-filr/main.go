package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/richardstrnad/terraform-provider-filr/filr"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: filr.Provider,
	})
}
