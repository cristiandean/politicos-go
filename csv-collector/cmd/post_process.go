// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package cmd

import (
	"errors"
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/olhoneles/politicos-go/csv-collector/collector"
	"github.com/spf13/cobra"
)

var postProcessCmd = &cobra.Command{
	Use:   "post-process",
	Short: "Post-processes the imported data",
	Run: func(cmd *cobra.Command, args []string) {
		subCmd := args[0]
		log.Debugf("Post-processing '%s'...", subCmd)
		err := runPostProcess(subCmd)
		if err != nil {
			log.Fatal(err)
		}
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a subcommand")
		}
		return nil
	},
}

func runPostProcess(step string) error {
	var postProcessSteps = map[string]func() error{
		"party":            collector.ProcessAllPoliticalParties,
		"political-office": collector.ProcessAllPoliticalOffices,
		"candidacy-status": collector.ProcessAllCandidaciesStatus,
		"education":        collector.ProcessAllEducations,
	}
	runAll := step == "all"
	postProcessFunc, stepExists := postProcessSteps[step]

	if !runAll && !stepExists {
		return fmt.Errorf("step '%s' unrecognized.", step)
	}
	if !runAll {
		return postProcessFunc()
	}

	var err error
	for runningStep, postProcessFunc := range postProcessSteps {
		log.Debugf("Running '%s'...	", runningStep)
		err = postProcessFunc()
		if err != nil {
			return err
		}
	}

	return nil
}
