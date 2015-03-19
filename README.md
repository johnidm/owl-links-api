#API CRUD Owl Links

###Written in [Go lang](https://golang.org/)

#### Setup project

You need to define Go root path.
```bash
export GOPATH=$HOME/go/owl-links-api
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
```
> You can to choose directory your preference. In my case I choose `$HOME/go/owl-links-api`.

Create folders in Go root
```bash
cd $GOPATH
mkdir pkg bin src
```

####Dependences

```bash
go get "gopkg.in/pg.v3"
go get "github.com/julienschmidt/httprouter"
go get "gopkg.in/mgo.v2"
go get "github.com/johnidm/owl-links-api/"
```

####Run project
```bash
go run src/github.com/johnidm/owl-links-api/service.go
```

Access `http://localhost:8000`

Heroku `http://owl-links-api.herokuapp.com/`

####Methods

**Note:** You need send field **api-key** in HTTP Header.

| HTTP Verb | Path (URL)| Description
|-----------|----------------------|----------------------------|
| GET       | /links               | Get all links              |
| GET       | /link/{id}           | Get specific link          |
| POST      | /link                | Insert new link            |
| PUT       | /link                | Update specific link       |
| DELETE    | /link/{id}           | Delete specific link       |
| GET       | /contacts            | Get all contacts           |
| GET       | /contact/{id}        | Get specific contact       |
| DELETE    | /contact/{id}        | Delete specific contact    |
| GET       | /collectlinks        | Get all collectlinks       |
| DELETE    | /collectlink/{id}    | Delete specific collectlink|
| GET       | /newslatters         | Get all newslatters        |
| DELETE    | /newslatter/{id}     | Delete specific newslatter |

####Example methods

Structure result body method **HTTP - GET** all links (return one or many records)

`http://owl-links-api.herokuapp.com/links`

```json
[  
   {  
      "_id":"p5o4ia0e4b024839a0458yu",
      "url":"www.douglas.com",
      "title":"E-Commerce",
      "description":"My E-Commerce",
      "signedup":"2015-03-02 23:16:25.846983",
      "tags" : 
		["PHP", "MySQL", "Go", "Java"],
      "notifynews" : "S"

   },
   {  
      "_id":"54fb0ea0e4b024839a01be80",
      "url":"www.johni.com",
      "title":"Blog",
      "description":"My Blog",
      "signedup":"2015-03-03 16:37:26.432123",
      "tags" : 
		["Linux", "Delphi", "HTML"],
      "notifynews" : "S"
   }
]
```

Structure result body method **HTTP - GET** specific links

`http://owl-links-api.herokuapp.com/link/{_id}`

```json
{  
   "_id":"54fb0ea0e4b024839a01be80",
   "url":"www.douglas.com",
   "title":"E-Commerce",
   "description":"Meu E-Commerce",
   "signedup":"2015-03-02 23:16:25.846983",
   "tags": [
        "PHP",
        "Scala",
        "Java"
    ],
   "notifynews" : "S"
}
```

Structure send body method **HTTP - POST**  

`http://owl-links-api.herokuapp.com/link`

```json
{  
   "url":"www.douglas.com",
   "title":"E-Commerce",
   "description":"Meu E-Commerce",
   "tags": [
        "Go",
        "MySQL",
        "Delphi"
    ]
   "notifynews" : "N"
}
```
Structure send body method **HTTP - PUT**

`http://owl-links-api.herokuapp.com/link`

```json
{  
   "_id":"54fb0ea0e4b024839a01be80",
   "url":"www.douglas.com",
   "title":"E-Commerce",
   "description":"Meu E-Commerce",
   "tags": [
        "PHP",
        "Scala",
        "Java"
    ]
   "notifynews" : "N"
}
```

Structure method **HTTP - DELETE** specific links `http://owl-links-api.herokuapp.com/link/{_id}`





