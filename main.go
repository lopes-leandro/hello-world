package main

import "github.com/gin-gonic/gin"
import "encoding/xml"

type Person struct {
	XMLName xml.Name `xml:"person"`
	FirstName string `xml:"firstName,attr"`
	LastName string `xml:"lastName,attr"`
}

func IndexHandler(c *gin.Context) {
	c.XML(200, Person{FirstName: "Leandro", LastName: "Lopes"})
}

func main() {
	router := gin.Default()
	router.GET("/", IndexHandler)
	router.Run(":5000")
}
