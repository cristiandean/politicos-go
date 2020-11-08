// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package collector

import (
	"github.com/labstack/gommon/log"
	"github.com/olhoneles/politicos-go/db"
	"github.com/olhoneles/politicos-go/politicos"
	"go.mongodb.org/mongo-driver/bson"
)

func ProcessAllPoliticalParties() error {
	log.Debug("[Collector] ProcessAllPoliticalParties")

	dbInstance, err := db.NewMongoSession()
	if err != nil {
		return err
	}

	opts := db.UniqueOptions{
		IDs: bson.D{
			{"siglum", "$sg_partido"},
			{"name", "$nm_partido"},
			{"tseNumber", "$nr_partido"},
		},
	}
	results, err := dbInstance.GetUnique(
		&politicos.Candidatures{},
		&politicos.PoliticalParty{},
		opts,
	)
	if err != nil {
		return err
	}

	// FIXME
	politicalParties := []interface{}{}
	for _, p := range results {
		politicalParties = append(politicalParties, p)
	}

	_, err = dbInstance.InsertMany(&politicos.PoliticalParty{}, politicalParties)
	if err != nil {
		return err
	}

	return nil
}
