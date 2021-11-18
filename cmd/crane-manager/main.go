package main

import (
	"fmt"
	"os"

	"github.com/gocrane-io/crane/cmd/crane-manager/app"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/component-base/logs"

	"github.com/gocrane-io/crane/pkg/utils/clogs"
)

// crane-manager main.
func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	clogs.InitLogs("crane-manager")

	ctx := genericapiserver.SetupSignalContext()

	if err := app.NewManagerCommand(ctx).Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}