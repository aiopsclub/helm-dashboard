package initialize

import (
	"fmt"
	"log"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/gates"
)

func debug(format string, v ...interface{}) {
	if settings.Debug {
		format = fmt.Sprintf("[debug] %s\n", format)
		log.Output(2, fmt.Sprintf(format, v...))
	}
}

const FeatureGateOCI = gates.Gate("HELM_EXPERIMENTAL_OCI")

var settings = cli.New()

var ActionConfig *action.Configuration = new(action.Configuration)

func init() {
	helmDriver := os.Getenv("HELM_DRIVER")
	if err := ActionConfig.Init(settings.RESTClientGetter(), settings.Namespace(), helmDriver, debug); err != nil {
		log.Fatal(err)
	}
}
