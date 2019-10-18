package apis

import (
	"net/http"
	"log"
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
	. "github.com/yangyouwei/router-api/models"
)


func AddlineApi(c *gin.Context) {
	ipaddr := c.Request.FormValue("ipaddr")
	port,_ := strconv.Atoi(c.Request.FormValue("port"))
	comment := c.Request.FormValue("comment")

	l := Line{Ipaddr:ipaddr, Port: port,Comment:comment}

	ra, err := l.AddLines()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func GetlinesApi(c *gin.Context) {
	var l Line
	lines, err := l.GetLines()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"lines": lines,
	})
}

func GetlineApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	l := Line{Id: id}
	line, err := l.GetLine()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"lines": line,
	})

}

func ModlineApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	l := Line{Id: id}
        l.Ipaddr = c.Params.ByName("ipaddr")

	err = c.Bind(&l)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err := l.ModLine()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("Update line %d successful %d", l.Id, ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func DellineApi(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	p := Line{Id: id}
	ra, err := p.DelLine()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("Delete line %d successful %d", id, ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func ApplayApi(c *gin.Context)  {
	fmt.Println("applay line")
}

func RebootApi(c *gin.Context)  {
	fmt.Println("reboot system")
}
