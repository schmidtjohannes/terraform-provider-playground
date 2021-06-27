package main

import (
	"context"
	"flag"
	"github.com/schmidtjohannes/terraform-provider-playground/playground"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{ProviderFunc: playground.Provider}

	if debugMode {
		// TODO: update this string with the full name of your provider as used in your configs
		err := plugin.Debug(context.Background(), "registry.terraform.io/schmidtjohannes.com/terraform-provider-playground", opts)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	plugin.Serve(opts)
}

