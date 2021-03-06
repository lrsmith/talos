/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

// nolint: dupl,golint
package cmd

import (
	"log"

	"github.com/autonomy/talos/internal/app/osinstall/internal/userdata"
	"github.com/autonomy/talos/internal/pkg/install"
	udata "github.com/autonomy/talos/internal/pkg/userdata"
	"github.com/spf13/cobra"
)

// installCmd reads in a userData file and attempts to parse it
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install Talos to a specified disk",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var ud *udata.UserData
		var err error
		if userdataFile != "" {
			ud, err = userdata.UserData(userdataFile)
			if err != nil {
				log.Fatal(err)
			}

			err = install.Prepare(ud)
			if err != nil {
				log.Fatal(err)
			}
		}
		err = install.Mount()
		if err != nil {
			log.Fatal(err)
		}

		err = install.Install(ud)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	installCmd.Flags().StringVarP(&userdataFile, "userdata", "u", "", "path or url of userdata file")
	rootCmd.AddCommand(installCmd)
}
