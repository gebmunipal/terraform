package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/hashicorp/terraform/internal/stresstest/internal/stressaddr"
	"github.com/hashicorp/terraform/internal/stresstest/internal/stressgen"
	"github.com/mitchellh/cli"
)

// graphCommand implements the "stresstest graph" command, which is the
// main index for the category of commands related to graph-based stress
// testing.
type graphCommand struct {
}

var _ cli.Command = (*graphCommand)(nil)

func (c *graphCommand) Run(args []string) int {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	seriesAddr := stressaddr.RandomConfigSeries(rnd)

	fmt.Printf("Series %s\n", seriesAddr)

	series := stressgen.GenerateConfigSeries(seriesAddr)

	for _, config := range series.Steps {
		fmt.Printf("# Config %s\n%s\n", config.Addr, config.GenerateConfigFile())
	}
	return 0
}

func (c *graphCommand) Synopsis() string {
	return "Stress-test the graph build and walk"
}

func (c *graphCommand) Help() string {
	return strings.TrimSpace(`
Usage: stresstest graph [subcommand]

...
`)
}
