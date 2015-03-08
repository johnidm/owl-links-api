package controllers

import (
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"

)

type Link struct {
  Id    bson.ObjectId `json:"id" bson:"_id"`
  Url   string        `json:"url" bson:"url"`
  Title string        `json:"title" bson:"title"`
  Description string  `json:"description" bson:"description"`
  Tags  []string      `json:"tags" bson:"tags"`
}

type Links []*Link

func Connect() (*mgo.Session, error) {
    uri := "mongodb://owl:owl@ds051851.mongolab.com:51851/owl-links"
    return mgo.Dial(uri)
} 

func CreateLink(link *Link, C *mgo.Collection) error {
    
    link.Id = bson.NewObjectId()

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
  
    err := C.UpdateId( link.Id, link )
    if err != nil {
      return err       
    }
   
    return nil
}


func GetLink( id string, C *mgo.Collection) (*Link, error) {
   
    link := &Link{}
    
    err := C.FindId(bson.ObjectIdHex(id)).One(&link)

    if err != nil {
      return nil, err 
    } 
    return link, nil
    
}

func GetLinks(C *mgo.Collection) (Links, error) {

    var links Links

    err := C.Find(bson.M{}).All(&links)
    if err != nil {
      return nil, err 
    }  
 
    return links, nil
}
