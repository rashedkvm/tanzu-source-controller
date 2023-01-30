//go:build tools
// +build tools

/*
Copyright 2022 VMware, Inc.

This product is licensed to you under the Apache License, V2.0 (the "License"). You may not use this product except in compliance with the License.

This product may include a number of subcomponents with separate copyright notices and license terms. Your use of these subcomponents is subject to the terms and conditions of the subcomponent's license, as noted in the LICENSE file.
*/

// This package imports things required by build scripts, to force `go mod` to see them as dependencies
package tools

import (
	_ "github.com/get-woke/woke"
)
