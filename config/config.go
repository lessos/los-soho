// Copyright 2015 Authors, All rights reserved.
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

package config

import (
	"code.hooto.com/lessos/iam/iamapi"
	"github.com/lessos/lessgo/crypto/idhash"
	"github.com/lessos/lessgo/types"
)

var (
	version    = "0.0.1.dev"
	InstanceId = idhash.HashToHexString([]byte("los-soho"), 16)
)

func IamAppInstance() iamapi.AppInstance {

	return iamapi.AppInstance{
		Meta: types.InnerObjectMeta{
			ID:   InstanceId,
			User: "sysadmin",
		},
		Version:  version,
		AppID:    "los-soho",
		AppTitle: "lessOS for SOHO",
		Status:   1,
		Url:      "",
		Privileges: []iamapi.AppPrivilege{
			{
				Privilege: "los.admin",
				Roles:     []uint32{1},
				Desc:      "lessOS Management",
			},
		},
	}
}
