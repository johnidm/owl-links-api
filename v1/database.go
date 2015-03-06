package database

import (
	"gopkg.in/pg.v3"
	"fmt"
)

const (
	DB_USER     = "jtbrooyyxtkpsz"
	DB_PASSWORD = "kS_f5apof0j0MiHVUC4WP74Rrm"
	DB_NAME     = "d9o4h18p276q2s"
	DB_DATABASE = "ec2-107-21-102-69.compute-1.amazonaws.com"
	DB_PORT     = "5432"
)

type Link struct {
	Id          int64  `json:"id"`
	Url         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Signedup	string `json:"signedup"`
}

type Links []*Link

func (links *Links) New() interface{} {
	u := &Link{}
	*links = append(*links, u)
	return u
}

func CreateLink(db *pg.DB, link *Link) error {

	fmt.Println(link)

	_, err := db.ExecOne(`INSERT INTO link(id, url, title, description, signedup) VALUES (nextval('link_seq'), ?url, ?title, ?description, current_timestamp)`, link)
	return err
}

func DeleteLink(db *pg.DB, id int64) error {
	_, err := db.ExecOne("DELETE FROM link WHERE id = ?", id)
	return err
}

func UpdateLink(db *pg.DB, link *Link) error {
	_, err := db.ExecOne("UPDATE link set url = ?url, title = ?title, description = ?description where id = ?id", link)
	return err
}

func GetLinks(db *pg.DB) (Links, error) {
	var links Links
	_, err := db.Query(&links, `SELECT id, url, title, description, signedup FROM link`)
	if err != nil {
		return nil, err
	}
	return links, nil
}

func GetLink(db *pg.DB, id int64) (*Link, error) {
	link := &Link{}
	_, err := db.QueryOne(link, `SELECT id, url, title, description, signedup FROM link WHERE id = ?`, id)

	return link, err
}

func OpenDB() *pg.DB {
	return pg.Connect(&pg.Options{
		Host:     DB_DATABASE,
		Database: DB_NAME,
		User:     DB_USER,
		Port:     DB_PORT,
		Password: DB_PASSWORD,
		SSL:      true,
	})
}
