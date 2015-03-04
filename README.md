#API CRUD Owl Links

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
go get github.com/johnidm/owl-links-api/
```

####Run project
```bash
go run src/github.com/johnidm/owl-links-api/service.go
```

Access `http://localhost:8000


####Methods

| HTTP Verb | Path (URL)| Description
|-----------|------------------|-----------------
| GET       | /links           | Get all links|
| GET       | /link/{id}      |Get specific link|
| POST     | /link            | Insert new link|
| PUT       | /link            | Update specific link|
| DELETE  | /link/{id}      | Delete specific link |

**Note:** You need send field **api-key** in HTTP Header.

Structure send body method **HTTP - POST**
```json
{  
   "url":"www.douglas.com",
   "title":"E-Commerce",
   "description":"Meu E-Commerce"
}
```
Structure send body method **HTTP - PUT**
```json
{  
   "id":1,
   "url":"www.douglas.com",
   "title":"E-Commerce",
   "description":"Meu E-Commerce",

}
```

Structure result body method **HTTP - GET** all links (return one or many records)
```json
[  
   {  
      "id":1,
      "url":"www.douglas.com",
      "title":"E-Commerce",
      "description":"My E-Commerce",
      "signedup":"2015-03-02 23:16:25.846983"
   },
   {  
      "id":2,
      "url":"www.johni.com",
      "title":"Blog",
      "description":"My Blog",
      "signedup":"2015-03-03 16:37:26.432123"
   }
]
```

Structure result body method **HTTP - GET** specific links

```json
{  
   "id":1,
   "url":"www.douglas.com",
   "title":"E-Commerce",
   "description":"Meu E-Commerce",
   "signedup":"2015-03-02 23:16:25.846983"
}
```





