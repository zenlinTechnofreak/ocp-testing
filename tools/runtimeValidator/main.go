// Copyright 2015 Huawei Inc. All Rights Reserved.
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

package main

import (
	"github.com/codegangsta/cli"
)

var Runtime string

func main() {
	app := cli.NewApp()
	app.Name = "oci-runtimeValidator"
	app.Version = "0.0.1"
	app.Usage = "Utilities for OCI runtime validation"
	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "runtime",
			Value: "runc",
			Usage: "runtime to be validated",
		},
	}
	app.Action = func(c *cli.Context) {
		Runtime = c.String("runtime")
		validate()
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
