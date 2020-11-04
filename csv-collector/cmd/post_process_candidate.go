// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package cmd

import (
	"github.com/olhoneles/politicos-go/csv-collector/collector"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var postProcessCandidateCmd = &cobra.Command{
	Use:   "post-process-candidate",
	Short: "Post process candidate data",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("Processing candidate...")
		if err := collector.ProcessAllCandidates(); err != nil {
			log.Fatalf("Couldn't process candidate! %v", err)
		}
	},
}
