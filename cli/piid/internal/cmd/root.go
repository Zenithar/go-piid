/*
 * Copyright 2019 Thibault NORMAND
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	iconfig "go.zenithar.org/piid/cli/piid/internal/config"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.zenithar.org/pkg/config"
	cmdcfg "go.zenithar.org/pkg/config/cmd"
	"go.zenithar.org/pkg/flags/feature"
	"go.zenithar.org/pkg/log"
	defaults "gopkg.in/mcuadros/go-defaults.v1"
)

// -----------------------------------------------------------------------------

// RootCmd describes root command of the tool
var mainCmd = &cobra.Command{
	Use:   "piid",
	Short: "Personal identifiable information database",
}

func init() {
	mainCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (config.toml)")
	mainCmd.AddCommand(versionCmd)
	mainCmd.AddCommand(cmdcfg.NewConfigCommand(conf, "PIID"))
}

// -----------------------------------------------------------------------------

// Execute main command
func Execute() error {
	feature.DefaultMutableGate.AddFlag(mainCmd.Flags())

	runtimeConf := conf.Instrumentation.Runtime.Config
	defaults.SetDefaults(&runtimeConf)
	conf.Instrumentation.Runtime.Config = runtimeConf

	initConfig()

	return mainCmd.Execute()
}

// -----------------------------------------------------------------------------

var (
	cfgFile string
	conf    = &iconfig.Configuration{}
)

// -----------------------------------------------------------------------------

func initConfig() {
	if err := config.Load(conf, "PIID", cfgFile); err != nil {
		log.Bg().Fatal("Unable load config", zap.Error(err))
	}
}
