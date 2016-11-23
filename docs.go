/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

// CmdDocs subcommand
import (
	"github.com/urfave/cli"
	"github.com/skratchdot/open-golang/open"
)

// CmdDocs: Open docs in the default browser
var CmdDocs = cli.Command{
	Name:      "docs",
	Usage:     "Open docs in the default browser.",
	ArgsUsage: " ",
	Description: `Open docs in the default browser.

   Example:
    $ ernest docs
	`,
	Action: func(c *cli.Context) error {
		open.Run("http://ernest.io/documentation/")
		return nil
	},
}
