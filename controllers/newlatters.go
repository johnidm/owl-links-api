package controllers

import (
	"gopkg.in/pg.v3"
)

type Newslatter struct {
	Id int64 `json:"id"`
	Name string `json:"name"` 
	Email string `json:"email"`  	 
	Subscribe string `json:"subscribe"`  	 
	Signedup string `json:"signedup"` 
}

type Newslatters []*Newslatter

func (newslatters *Newslatters) New() interface{} {
	u := &Newslatter{}
	*newslatters = append(*newslatters, u)
	return u
}


func GetNewslatters(db *pg.DB) (Newslatters, error) { 
	var newslatters Newslatters 

	_, err := db.Query(&newslatters, `SELECT id, name, email, subscribe, signedup FROM newslatter`)

	if err != nil {
		return nil, err
	}
	return newslatters, nil
}

func DeleteNewslatter(db *pg.DB, id int64) error {
	_, err := db.ExecOne("DELETE FROM newslatter WHERE id = ?", id)
	return err
}


