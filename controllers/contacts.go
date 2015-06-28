package controllers

import (
	"gopkg.in/pg.v3"
)

type ContactFull struct {
	Id        int64  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Site      string `json:"site"`
	Message   string `json:"message"`
	Signedup  string `json:"signedup"`
}

type ContactClean struct {
	Id        int64  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

type ContactsClean struct {
	C []ContactClean
}

var _ pg.Collection = &ContactsClean{}

func (contactsclean *ContactsClean) NewRecord() interface{} {
	contactsclean.C = append(contactsclean.C, ContactClean{})
	return &contactsclean.C[len(contactsclean.C)-1]
}

func GetContacts(db *pg.DB) ([]ContactClean, error) {
	var contacts ContactsClean

	_, err := db.Query(&contacts, `SELECT id, firstname, lastname, email FROM contact order by id desc`)

	if err != nil {
		return nil, err
	}
	return contacts.C, nil
}

func GetContact(db *pg.DB, id int64) (*ContactFull, error) {

	contact := &ContactFull{}
	_, err := db.QueryOne(contact, `SELECT id, firstname, lastname, email, site, message, signedup FROM contact WHERE id = ?`, id)

	return contact, err
}

func DeleteContact(db *pg.DB, id int64) error {
	_, err := db.ExecOne("DELETE FROM contact WHERE id = ?", id)
	return err
}
