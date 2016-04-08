package main

import (
	"os"

	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/tscolari/bosh-c3pi/cpi"
)

func main() {
	cpiClient := yourimplementation.New()
	logger := boshlog.NewWriterLogger(boshlog.LevelDebug, os.Stderr, os.Stderr)
	runner := cpi.NewRunner(cpiClient, logger)

	runner.Run(os.Stdin, os.Stderr)
}
