package controllers

import (
	"gopkg.in/pg.v3"
)

type Newsletter struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Subscribe string `json:"subscribe"`
	Signedup  string `json:"signedup"`
}

type Newsletters struct {
	C []Newsletter
}

var _ pg.Collection = &Newsletters{}

func (newsletters *Newsletters) NewRecord() interface{} {
	newsletters.C = append(newsletters.C, Newsletter{})
	return &newsletters.C[len(newsletters.C)-1]
}

func GetNewsletters(db *pg.DB) ([]Newsletter, error) {
	var newsletters Newsletters

	_, err := db.Query(&newsletters, `SELECT id, name, email, subscribe, signedup FROM newsletter order by id desc`)

	if err != nil {
		return nil, err
	}
	return newsletters.C, nil
}

func DeleteNewsletter(db *pg.DB, id int64) error {
	_, err := db.ExecOne("DELETE FROM newsletter WHERE id = ?", id)
	return err
}
