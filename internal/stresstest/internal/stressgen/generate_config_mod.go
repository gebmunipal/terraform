package stressgen

import (
	"github.com/hashicorp/terraform/internal/stresstest/internal/stressaddr"
)

// GenerateModifiedConfig produces a new configuration which is a valid
// modification of the reciever, using the given modification address as
// a random seed for deciding what to change.
func (c *Config) GenerateModifiedConfig(modAddr stressaddr.ModConfig) *Config {
	rnd := newRand(modAddr.RandomSeed())
	addr := c.Addr.NewMod(modAddr)

	objs := make([]ConfigObject, 0, len(c.Objects))
	for _, obj := range c.Objects {
		new := obj.GenerateModified(rnd)
		if new == nil {
			// This represents removing the object altogether.
			continue
		}
		objs = append(objs, new)

		// TODO: With a relatively low likelihood, potentially generate
		// new blocks too.
	}

	return &Config{
		Addr:    addr,
		Objects: objs,
	}
}
