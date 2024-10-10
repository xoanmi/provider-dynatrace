/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/xoanmi/provider-dynatrace/config/alerting"
	"github.com/xoanmi/provider-dynatrace/config/dashboard"
	"github.com/xoanmi/provider-dynatrace/config/event"
	"github.com/xoanmi/provider-dynatrace/config/notification"
)

const (
	resourcePrefix = "dynatrace"
	modulePath     = "github.com/xoanmi/provider-dynatrace"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		alerting.Configure,
		event.Configure,
		notification.Configure,
		dashboard.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
