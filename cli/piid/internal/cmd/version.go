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
	"fmt"

	"github.com/spf13/cobra"
	"go.zenithar.org/piid/internal/version"
)

// -----------------------------------------------------------------------------

var displayAsJSON bool

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display service version",
	Run: func(cmd *cobra.Command, args []string) {
		if displayAsJSON {
			fmt.Printf("%s", version.JSON())
		} else {
			fmt.Printf("%s", version.Full())
		}
	},
}

func init() {
	versionCmd.Flags().BoolVar(&displayAsJSON, "json", false, "Display build info as json")
}
