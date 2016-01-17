package middlewares

import(
  "net/http"
  "github.com/labstack/echo"
  "github.com/unsign3d/provago/models"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)


func Welcome(c *echo.Context) error {

  session, err := mgo.Dial("localhost")
  if err != nil {
         panic(err)
  }

  defer session.Close()

  mongo := session.DB("welcome").C("welcomes")

  wm :=  models.Welcome{
    Id: "123456",
    Message: "Hello world",
  }

  err = mongo.Insert(&wm)

  if err != nil {
    panic(err)
  }

  result := models.Welcome{}

  err = mongo.Find(bson.M{"id":"123456"}).One(&result)

  if err != nil {
    panic(err)
  }

	return c.String(http.StatusOK, result.Message)
}
