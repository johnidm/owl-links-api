package controllers

import (
	"gopkg.in/pg.v3"
)

type Collectlink struct {
	Id int64 `json:"id"`
	Link string `json:"link"` 	
	Signedup string `json:"signedup"` 
}

type Collectlinks []*Collectlink

func (collectlinks *Collectlinks) New() interface{} {
	u := &Collectlink{}
	*collectlinks = append(*collectlinks, u)
	return u
}


func GetCollectlinks(db *pg.DB) (Collectlinks, error) { 
	var collectlinks Collectlinks 

	_, err := db.Query(&collectlinks, `SELECT id, link, signedup FROM collectlink`)

	if err != nil {
		return nil, err
	}
	return collectlinks, nil
}

func DeleteCollectlink(db *pg.DB, id int64) error {
	_, err := db.ExecOne("DELETE FROM collectlink WHERE id = ?", id)
	return err
}


