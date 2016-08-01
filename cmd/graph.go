// Copyright © 2016 Asteris, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"fmt"
	"log"

	"golang.org/x/net/context"

	"github.com/asteris-llc/converge/load"
	"github.com/asteris-llc/converge/prettyprinters"
	"github.com/asteris-llc/converge/prettyprinters/graphviz"
	"github.com/asteris-llc/converge/prettyprinters/graphviz/providers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// graphCmd represents the graph command
var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "graph the execution of a module",
	Long: `graphing is a convenient way to visualize how your graph and
dependencies are structured.

You can pipe the output directly to the 'dot' command, for example:

		converge graph myFile.hcl | dot -Tpng -o myFile.png`,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Need one module filename as argument, got 0")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fname := args[0]

		ctx, cancel := context.WithCancel(context.Background())
		GracefulExit(cancel)

		graph, err := load.Load(ctx, fname)
		if err != nil {
			log.Fatalf("[FATAL] %s: could not parse file: %s\n", fname, err)
		}

		provider := providers.ResourceProvider{}
		provider.ShowParams = viper.GetBool("show-params")
		dotPrinter := graphviz.New(graphviz.DefaultOptions(), provider)
		printer := prettyprinters.New(dotPrinter)
		dotCode, err := printer.Show(ctx, graph)
		if err != nil {
			log.Fatalf("[FATAL] %s: could not generate dot output: %s", fname, err)
		}
		fmt.Println(dotCode)

	},
}

func init() {
	graphCmd.Flags().Bool("show-params", false, "also graph param dependencies")
	viperBindPFlags(graphCmd.Flags())

	RootCmd.AddCommand(graphCmd)
}