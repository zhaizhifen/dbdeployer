// Copyright © 2017-2018 Giuseppe Maxia
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
	//"fmt"

	"github.com/datacharmer/dbdeployer/sandbox"
	"github.com/spf13/cobra"
)

func MultipleSandbox(cmd *cobra.Command, args []string) {
	var sd sandbox.SandboxDef
	flags := cmd.Flags()
	sd = FillSdef(cmd, args)
	nodes, _ := flags.GetInt("nodes")
	sandbox.CreateMultipleSandbox(sd, args[0], nodes)
}

// multipleCmd represents the multiple command
var multipleCmd = &cobra.Command{
	Use:   "multiple MySQL-Version",
	Args:  cobra.ExactArgs(1),
	Short: "create multiple sandbox",
	Long:  ``,
	Run:   MultipleSandbox,
	Example: `
	$ dbdeployer multiple 5.7.21
	Installing and starting node 1
	Database installed in /Users/gmax/sandboxes/multi_msb_5_7_21/node1
	. sandbox server started
	Installing and starting node 2
	Database installed in /Users/gmax/sandboxes/multi_msb_5_7_21/node2
	. sandbox server started
	Installing and starting node 3
	Database installed in /Users/gmax/sandboxes/multi_msb_5_7_21/node3
	. sandbox server started
	Multiple directory installed in /Users/gmax/sandboxes/multi_msb_5_7_21
	`,
}

func init() {
	rootCmd.AddCommand(multipleCmd)
	multipleCmd.PersistentFlags().IntP("nodes", "n", 3, "How many nodes will be installed")
}