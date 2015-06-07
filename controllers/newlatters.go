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

type Newsletters []*Newsletter

func (newsletters *Newsletters) New() interface{} {
	u := &Newsletter{}
	*newsletters = append(*newsletters, u)
	return u
}

func GetNewsletters(db *pg.DB) (Newsletters, error) {
	var newsletters Newsletters

	_, err := db.Query(&newsletters, `SELECT id, name, email, subscribe, signedup FROM newsletter`)

	if err != nil {
		return nil, err
	}
	return newsletters, nil
}

func DeleteNewsletter(db *pg.DB, id int64) error {
	_, err := db.ExecOne("DELETE FROM newsletter WHERE id = ?", id)
	return err
}
