// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package politicos

type Candidate struct {
	CPF                           string
	Data                          []Candidatures `bson:"data"`
	CD_COR_RACA                   string
	CD_ESTADO_CIVIL               string
	CD_MUNICIPIO_NASCIMENTO       string
	CD_NACIONALIDADE              string
	DS_COR_RACA                   string
	DS_ESTADO_CIVIL               string
	DS_NACIONALIDADE              string
	DT_NASCIMENTO                 string
	NM_CANDIDATO                  string
	NM_EMAIL                      string
	NM_MUNICIPIO_NASCIMENTO       string
	NM_SOCIAL_CANDIDATO           string
	NM_URNA_CANDIDATO             string
	NR_TITULO_ELEITORAL_CANDIDATO string
	SG_UF_NASCIMENTO              string
}

func (c Candidate) GetCollectionName() string {
	return "candidates"
}

func (c Candidate) Cast() Queryable {
	return &c
}
