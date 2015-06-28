package controllers

import (
	"gopkg.in/pg.v3"
)

type Collectlink struct {
	Id       int64  `json:"id"`
	Link     string `json:"link"`
	Signedup string `json:"signedup"`
}

type Collectlinks struct {
	C []Collectlink
}

var _ pg.Collection = &Collectlinks{}

func (collectlinks *Collectlinks) NewRecord() interface{} {
	collectlinks.C = append(collectlinks.C, Collectlink{})
	return &collectlinks.C[len(collectlinks.C)-1]
}

func GetCollectlinks(db *pg.DB) ([]Collectlink, error) {
	var collectlinks Collectlinks

	_, err := db.Query(&collectlinks, `SELECT id, link, signedup FROM collectlink order by id desc`)

	return collectlinks.C, err
}

func DeleteCollectlink(db *pg.DB, id int64) error {
	_, err := db.ExecOne("DELETE FROM collectlink WHERE id = ?", id)
	return err
}
