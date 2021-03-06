// Copyright 2019 Eryx <evorui аt gmail dοt com>, All rights reserved.
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
	"fmt"

	"github.com/sysinner/incore/config"
	"github.com/sysinner/incore/inapi"
)

type hostJoinCommand struct {
	cmd        *baseCommand
	argRequest config.HostJoinRequest
}

func NewHostJoinCommand() *baseCommand {

	c := &hostJoinCommand{
		cmd: &baseCommand{
			Use:   "host-join",
			Short: "initialize current host to join the zone cluster",
			Long:  `Perform one-time-only initialization of a host to join the zone cluster`,
		},
	}

	c.cmd.Flags().StringVar(&c.argRequest.ZoneAddr, "zone-addr",
		"",
		`one of the zone main node address (format: <ip>[:port])
the ip must be a LAN ip in range of:
   10.0.0.0 ~ 10.255.255.255,
   172.16.0.0 ~ 172.31.255.255,
   192.168.0.0 ~ 192.168.255.25.
if the port number is left unspecified, it defaults to 9529
		`)

	c.cmd.Flags().StringVar(&c.argRequest.HostAddr, "host-addr",
		"",
		`cuurrent local host address (format <ip>[:port])
the ip must be a LAN ip in range of:
   10.0.0.0 ~ 10.255.255.255,
   172.16.0.0 ~ 172.31.255.255,
   192.168.0.0 ~ 192.168.255.25.
if the port number is left unspecified, it defaults to 9529
		`)

	c.cmd.Flags().StringVar(&c.argRequest.CellId, "cell-id",
		"g1",
		"the name of host group")

	c.cmd.RunE = c.run

	return c.cmd
}

func (it *hostJoinCommand) run(cmd *baseCommand, args []string) error {

	if err := rootAllow(); err != nil {
		return err
	}

	var rep inapi.WebServiceReply
	if err := localApiCommand("config/host-join", it.argRequest, &rep); err != nil {
		return err
	}

	if rep.Kind != "OK" {
		return fmt.Errorf("Fail: %s\n", rep.Message)
	}

	fmt.Println("host successfully initialized")
	return nil
}
