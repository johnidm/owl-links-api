package controllers

import (
	"gopkg.in/pg.v3"
)

type ContactFull struct {
	Id int64 `json:"id"`
	Firstname string `json:"firstname"` 
	Lastname string `json:"lastname"` 
	Email string `json:"email"` 
	Site string `json:"site"` 
	Message string `json:"message"` 		 
	Signedup string `json:"signedup"` 
}

type ContactClean struct {
	Id int64 `json:"id"`
	Firstname string `json:"firstname"` 
	Lastname string `json:"lastname"` 
	Email string `json:"email"` 	
}

type ContactsClean []*ContactClean

func (contacts *ContactsClean) New() interface{} {
	u := &ContactClean{}
	*contacts = append(*contacts, u)
	return u
}

func GetContacts(db *pg.DB) (ContactsClean, error) { 
	var contacts ContactsClean 

	_, err := db.Query(&contacts, `SELECT id, firstname, lastname, email FROM contact`)

	if err != nil {
		return nil, err
	}
	return contacts, nil
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
