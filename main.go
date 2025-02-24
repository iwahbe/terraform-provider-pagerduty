package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tf5server"
	"github.com/hashicorp/terraform-plugin-mux/tf5muxserver"

	"github.com/terraform-providers/terraform-provider-pagerduty/pagerduty"
	pagerdutyplugin "github.com/terraform-providers/terraform-provider-pagerduty/pagerdutyplugin"
)

func main() {
	Serve()
}

func Serve() {
	ctx := context.Background()

	muxServer, err := tf5muxserver.NewMuxServer(
		ctx,
		// terraform-plugin-framework
		providerserver.NewProtocol5(pagerdutyplugin.New()),
		// terraform-plugin-sdk
		pagerduty.Provider().GRPCProvider,
	)
	if err != nil {
		log.Fatal(err)
	}

	var serveOpts []tf5server.ServeOpt

	address := "registry.terraform.io/pagerduty/pagerduty"
	err = tf5server.Serve(address, muxServer.ProviderServer, serveOpts...)
	if err != nil {
		log.Fatal(err)
	}
}
