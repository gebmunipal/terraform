package stressgen

import (
	"github.com/hashicorp/hcl/v2/hclwrite"

	"github.com/hashicorp/terraform/internal/stresstest/internal/stressaddr"
)

// Config represents a generated configuration.
//
// It only directly refers to the generated root module, but that module might
// in turn contain references to child modules via module call objects.
//
// This type and most of its descendents have exported fields just because
// this package is aimed at testing use-cases and having them exported tends
// to make debugging easier. With that said, external callers should generally
// not modify any data in those exported fields, and should instead prefer to
// use the methods on these types that know how to derive new objects while
// keeping all of the expected invariants maintained.
//
// The top-level object representing a test case is ConfigSeries, which is a
// sequence of Config instances that will be planned, applied, and verified in
// order. Config therefore represents only a single step in a test case.
type Config struct {
	// Addr is an identifier for this particular generated configuration, which
	// a caller can use to rebuild the same configuration as long as nothing
	// in the config generator code has changed in the meantime.
	Addr stressaddr.Config

	// A generated configuration is made from a series of "objects", each of
	// which typically corresponds to one configuration block when we serialize
	// the configuration into normal Terraform language input.
	//
	// Some ConfigObjects also know how to verify that a final state contains
	// the results they expect, which is part of our definition of success
	// or failure when we're verifying test results.
	Objects []ConfigObject
}

// GenerateConfigFile generates the potential content of a single configuration
// (.tf) file which declares all of the given configuration objects.
//
// It's the caller's responsibility to make sure that the given objects all
// make sense to be together in a single module, including making sure they all
// together meet any uniqueness constraints and that any objects that refer
// to other objects are given along with the objects they refer to.
func (c *Config) GenerateConfigFile() []byte {
	f := hclwrite.NewEmptyFile()
	body := f.Body()
	for _, obj := range c.Objects {
		obj.AppendConfig(body)
	}
	return f.Bytes()
}
