// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package collector

import (
	"github.com/olhoneles/politicos-go/db"
	"github.com/olhoneles/politicos-go/politicos"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

const ItemsPerInsert = 1000

func ProcessAllCandidates() error {
	log.Debug("[Collector] ProcessAllCandidates")

	dbInstance, err := db.NewMongoSession()
	if err != nil {
		return err
	}

	log.Debugf("[Collector] Getting all unique CPF...")
	opts := db.UniqueOptions{IDs: bson.D{{"cpf", "$nr_cpf_candidato"}}}
	// FIXME: paginate?
	cpfs, err := dbInstance.GetUnique(
		&politicos.Candidatures{},
		&politicos.Candidate{},
		opts,
	)
	if err != nil {
		return err
	}

	log.Debugf("[Collector] Total of CPFs: %d", len(cpfs))

	candidates := []interface{}{}
	for _, candidate := range cpfs {
		c := candidate.(*politicos.Candidate)
		log.Debugf("[Collector] Grouping %s data...", c.CPF)
		opts := db.UniqueOptions{
			Match: bson.D{{"nr_cpf_candidato", c.CPF}},
			IDs:   bson.D{{"cpf", "$nr_cpf_candidato"}},
			Data:  bson.D{{"$push", "$$ROOT"}},
			Sort:  bson.D{{"ano_eleicao", -1}, {"nr_turno", -1}},
		}
		results, err := dbInstance.GetUnique(
			&politicos.Candidatures{},
			&politicos.Candidate{},
			opts,
		)
		if err != nil {
			return err
		}

		// FIXME
		item := results[0].(*politicos.Candidate)
		candidature := item.Candidatures[0]
		item.CPF = candidature.NR_CPF_CANDIDATO
		item.CD_COR_RACA = candidature.CD_COR_RACA
		item.CD_ESTADO_CIVIL = candidature.CD_ESTADO_CIVIL
		item.CD_MUNICIPIO_NASCIMENTO = candidature.CD_MUNICIPIO_NASCIMENTO
		item.CD_NACIONALIDADE = candidature.CD_NACIONALIDADE
		item.DS_COR_RACA = candidature.DS_COR_RACA
		item.DS_ESTADO_CIVIL = candidature.DS_ESTADO_CIVIL
		item.DS_NACIONALIDADE = candidature.DS_NACIONALIDADE
		item.DT_NASCIMENTO = candidature.DT_NASCIMENTO
		item.NM_CANDIDATO = candidature.NM_CANDIDATO
		item.NM_EMAIL = candidature.NM_EMAIL
		item.NM_MUNICIPIO_NASCIMENTO = candidature.NM_MUNICIPIO_NASCIMENTO
		item.NM_SOCIAL_CANDIDATO = candidature.NM_SOCIAL_CANDIDATO
		item.NM_URNA_CANDIDATO = candidature.NM_URNA_CANDIDATO
		item.NR_TITULO_ELEITORAL_CANDIDATO = candidature.NR_TITULO_ELEITORAL_CANDIDATO
		item.SG_UF_NASCIMENTO = candidature.SG_UF_NASCIMENTO

		candidates = append(candidates, item)

		if len(candidates) == ItemsPerInsert {
			log.Debugf("[Collector] Inserting %d...", ItemsPerInsert)
			_, err := dbInstance.InsertMany(&politicos.Candidate{}, candidates)
			if err != nil {
				return err
			}
			candidates = []interface{}{}
		}
	}

	if len(candidates) > 0 {
		log.Debugf("[Collector] Inserting %d...", len(candidates))
		_, err = dbInstance.InsertMany(&politicos.Candidate{}, candidates)
		if err != nil {
			return err
		}
	}

	return nil
}
