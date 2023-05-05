/*
Copyright 2021 Upbound Inc.
*/

package advancedcluster

import (
	ujconfig "github.com/upbound/upjet/pkg/config"
)

// Configure configures the null group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("mongodbatlas_advanced_cluster", func(r *ujconfig.Resource) {
		// r.Kind = "Resource"
		// And other overrides.
	})
}
