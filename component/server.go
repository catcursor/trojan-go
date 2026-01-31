//go:build server || full || mini
// +build server full mini

package build

import (
	_ "github.com/catcursor/trojan-go/proxy/server"
)
