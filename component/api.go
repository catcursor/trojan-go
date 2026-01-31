//go:build api || full
// +build api full

package build

import (
	_ "github.com/catcursor/trojan-go/api/control"
	_ "github.com/catcursor/trojan-go/api/service"
)
