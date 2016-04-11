package main

import (
	"encoding/json"
	"os"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/scaleway/scaleway-cli/pkg/api"
	"github.com/tscolari/bosh-c3pi/cpi"
	"github.com/tscolari/bosh-scaleway-cpi/scaleway"
)

func main() {
	var config scaleway.Config
	json.Unmarshal([]byte(os.Getenv("CONFIG")), &config)

	scalewayApi, _ := api.NewScalewayAPI(config.Organization, config.Token, config.UserAgent)

	cpiClient := scaleway.New(scalewayApi)

	logger := boshlog.NewWriterLogger(boshlog.LevelDebug, os.Stderr, os.Stderr)
	runner := cpi.NewRunner(cpiClient, logger)

	runner.Run(os.Stdin, os.Stdout)
}
