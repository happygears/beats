// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package netspyglass

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "netspyglass", asset.ModuleFieldsPri, AssetNetspyglass); err != nil {
		panic(err)
	}
}

// AssetNetspyglass returns asset data.
// This is the base64 encoded gzipped contents of module/netspyglass.
func AssetNetspyglass() string {
	return "eJyskz9u8zAMxXef4iF7fAAP3/INBQqkS3MBImYUIbIkSExa376Q3SSOqvwxWo4k9d4PFLnEnvsGliX6XhmKsQJEi+EGi0l2UQEtx03QXrSzDf5VAPDG8u77l9SBlWsPhitgq9m0sRkalrDUcW6QQnrPDVRwB/+dKehfi00FjVPnXEnspuAYU3DjFIy2HOtJS+575c1HNleVEwEZTTGreJLdwFuX3nVaBRoJJRy46NdxjKR4pmP51S2/ovHGdd5ZtlK03nP/4UKb1e4MPcX/kyRkx8PodUw+2ipsg+vqMshkcf4A4pWONGrOopBdYMqtfoGxHvQG8XkcgTblbcgv4AkIYJ3koO2ZIR1DnTWW7gEPfujReJ6i+/FbMuLemRMe3s6FTPgz3+4nsYDVqA1PQeC2F7q6+goAAP//eUlZeg=="
}