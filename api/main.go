package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

// Topos obj
type Topos struct {
	Specs map[string]interface{}
}

var topoi *mgo.Collection

func main() {
	// Connect to mongo
	session, err := mgo.Dial("mongo:27017")
	if err != nil {
		log.Fatalln(err)
		log.Fatalln("mongo err")
		os.Exit(1)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// Get topoi collection
	topoi = session.DB("app").C("topoi")

	// Set up routes
	r := gin.Default()
	r.POST("/topoi", createTopos)
	r.GET("/topoi", readTopoi)

	r.Run(":8080")
}

func createTopos(ctx *gin.Context) {
	// Read body
	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not found!"})
		return
	}

	// Read topos
	var topos interface{}
	err = json.Unmarshal([]byte(data), &topos)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not found!"})
		return
	}

	// Insert new topos
	if err := topoi.Insert(topos); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Not found !"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "topos": topos})

	return
}

func readTopoi(ctx *gin.Context) {
	var result []interface{}
	if err := topoi.Find(nil).Sort("-created_at").All(&result); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No topos found!"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": result})
	}
}
