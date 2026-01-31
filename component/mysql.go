//go:build mysql || full || mini
// +build mysql full mini

package build

import (
	_ "github.com/catcursor/trojan-go/statistic/mysql"
)
