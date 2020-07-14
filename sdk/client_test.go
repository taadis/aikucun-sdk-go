package sdk

import (
	"flag"
	"testing"
)

var (
	appId      = flag.String("appId", "", "appId")
	appSecret  = flag.String("appSecret", "", "appSecret")
	erp        = flag.String("erp", "", "erp")
	erpVersion = flag.String("erpVersion", "", "erpVersion")
)

// parseArgs
func parseArgs(t *testing.T) {
	if !flag.Parsed() {
		flag.Parse()
	}
	t.Logf("-args -appId %s -appSecret %s", *appId, *appSecret)
}
