package controllers

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Link struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Url         string        `json:"url" bson:"url"`
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description"`
	Tags        []string      `json:"tags" bson:"tags"`
	Notifynews  string        `json:"notifynews" bson:"notifynews"`
	CreatedAt 	string 		  `json:"created_at" bson:"created_at"`
	UpadateAt 	string 		  `json:"upadate_at" bson:"upadate_at"`

}

type Links []*Link

type Tags []string

func Connect() (*mgo.Session, error) {
	uri := "mongodb://owl:owl@ds051851.mongolab.com:51851/owl-links"
	return mgo.Dial(uri)
}

func getNowDate() string {
	return time.Now().Format(time.RFC3339)
}

func Exists(C *mgo.Collection, url string) bool {
	count, _ := C.Find(bson.M{"url" : url}).Count()

	return count > 0
}

func CreateLink(link *Link, C *mgo.Collection) error {

	link.Id = bson.NewObjectId()

	link.CreatedAt = getNowDate()
	link.UpadateAt = getNowDate()

	err := C.Insert(link)
	if err != nil {
		return err
	}

	return nil

}

func DeleteLink(id string, C *mgo.Collection) error {

	err := C.RemoveId(bson.ObjectIdHex(id))

	if err != nil {
		return err
	}

	return nil
}

func UpdateLink(link *Link, C *mgo.Collection) error {

	link.UpadateAt = getNowDate()

	err := C.UpdateId(link.Id, link)
	if err != nil {
		return err
	}

	return nil
}

func GetLink(id string, C *mgo.Collection) (*Link, error) {

	link := &Link{}

	err := C.FindId(bson.ObjectIdHex(id)).One(&link)

	if err != nil {
		return nil, err
	}
	return link, nil

}

func GetLinks(C *mgo.Collection) (Links, error) {

	var links Links

	err := C.Find(bson.M{}).Sort("-upadate_at").All(&links)
	if err != nil {
		return nil, err
	}

	return links, nil
}

func GetLinksTags(C *mgo.Collection) (Tags, error) {

	var tags Tags

	err := C.Find(bson.M{}).Distinct("tags", &tags)

	if err != nil {
		return nil, err
	}

	return tags, nil
}
