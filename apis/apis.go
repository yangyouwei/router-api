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

	caddr,_ := c.GetPostForm("ipaddr")
	addr := fmt.Sprint(caddr)
	fmt.Println(addr)

	cport ,_ := c.GetPostForm("port")
	port, err := strconv.Atoi(cport)
	ccomment, _ := c.GetPostForm("comment")
	comment := fmt.Sprint(ccomment)
	l := Line{Id: id,Ipaddr: addr,Port:port,Comment:comment}

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
	
	aline := line.Ipaddr + ":" + strconv.Itoa(id)
	fmt.Fprintln(gin.DefaultWriter, "[GIN] applay line ",aline)
	fmt.Println(aline)
	fmt.Println("reboot vpn redirect service.")
	fmt.Println("clear iptables rules.")
	fmt.Println("imoport new iptables rules for new line.")
        msg := fmt.Sprintf("line will be applay and restart redirect service, please waite secoends.")
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func RebootApi(c *gin.Context)  {
        fmt.Println("reboot system")
	msg := fmt.Sprintf("system will be reboot, please waite secoends.")
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func UseLineApi(c *gin.Context)  {
	l := getline()
	c.JSON(http.StatusOK, gin.H{
		"lines": l,
	})
}

func getline() (l Line) {
	l.Ipaddr = "192.168.2.55"
	l.Port = 443
	return 
}
