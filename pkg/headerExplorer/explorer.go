package headerExplorer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type Explorer interface {
	Get() interface{}
}

func NewGinExplorer(c *gin.Context) Explorer {
	return &ginExplorer{
		c: c,
	}
}

type ginExplorer struct {
	c *gin.Context
}

func (g *ginExplorer) Get() interface{} {
	body, _ := ioutil.ReadAll(g.c.Request.Body)
	response := RequestContent{
		RequestHeaders: g.c.Request.Header,
		RequestMethod:  g.c.Request.Method,
		Body:           fmt.Sprintf("%s", body),
	}

	return response
}
