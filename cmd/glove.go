// Copyright © 2017 Makoto Ito
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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ynqa/word-embedding/config"
)

var GloVeCmd = &cobra.Command{
	Use:   "glove",
	Short: "Embed words using glove",
	Long:  "Embed words using glove",
	PreRun: func(cmd *cobra.Command, args []string) {
		configBind(cmd)
		gloveBind(cmd)
	},
	Run: func(cmd *cobra.Command, args []string) {
		runGloVe()
	},
}

func init() {
	GloVeCmd.Flags().AddFlagSet(ConfigFlagSet())
	GloVeCmd.Flags().Int(config.Iteration.String(), config.DefaultIteration,
		"Set the iteration")
	GloVeCmd.Flags().Float64(config.Alpha.String(), config.DefaultAlpha,
		"Set alpha")
	GloVeCmd.Flags().Int(config.XMax.String(), config.DefaultXMax,
		"Set xmax")
}

func gloveBind(cmd *cobra.Command) {
	viper.BindPFlag(config.Alpha.String(), cmd.Flags().Lookup(config.Alpha.String()))
}

func runGloVe() error {
	return nil
}
