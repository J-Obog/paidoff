package server

import (
	"fmt"
	"io"

	"github.com/J-Obog/paidoff/data"
	"github.com/J-Obog/paidoff/rest"
	"github.com/gin-gonic/gin"
)

type GinServer struct {
	BaseServer
	eng *gin.Engine
}

func (g *GinServer) Start(port int) error {
	g.eng.GET("/account", ginHandler(g.accountManager.GetByRequest))
	g.eng.PUT("/account", ginHandler(g.accountManager.UpdateByRequest))
	g.eng.DELETE("/account", ginHandler(g.accountManager.DeleteByRequest))

	g.eng.GET("/category", ginHandler(g.categoryManager.GetAllByRequest))
	g.eng.GET("/category/:id", ginHandler(g.categoryManager.GetByRequest))
	g.eng.POST("/category", ginHandler(g.categoryManager.CreateByRequest))
	g.eng.PUT("/category/:id", ginHandler(g.categoryManager.UpdateByRequest))
	g.eng.DELETE("/category/:id", ginHandler(g.accountManager.DeleteByRequest))

	g.eng.GET("/transaction", ginHandler(g.transactionManager.GetAllByRequest))
	g.eng.GET("/transaction/:id", ginHandler(g.transactionManager.GetByRequest))
	g.eng.POST("/transaction", ginHandler(g.transactionManager.CreateByRequest))
	g.eng.PUT("/transaction/:id", ginHandler(g.transactionManager.UpdateByRequest))
	g.eng.DELETE("/transaction/:id", ginHandler(g.transactionManager.DeleteByRequest))

	g.eng.GET("/budget", ginHandler(g.budgetManager.GetAllByRequest))
	g.eng.GET("/budget/:id", ginHandler(g.budgetManager.GetByRequest))
	g.eng.POST("/budget", ginHandler(g.budgetManager.CreateByRequest))
	g.eng.PUT("/budget/:id", ginHandler(g.budgetManager.UpdateByRequest))
	g.eng.DELETE("/budget/:id", ginHandler(g.budgetManager.DeleteByRequest))

	return g.eng.Run(fmt.Sprintf(":%d", port))
}

func (g *GinServer) Stop() error {
	return nil
}

func ginCtxToRequest(c *gin.Context) *rest.Request {
	//TODO: handle error when reading body

	b, _ := io.ReadAll(c.Request.Body)

	// TODO: remove dummy account for auth
	req := &rest.Request{
		Account: &data.Account{
			Id: "some-account-id",
		},

		Url:   c.Request.URL.String(),
		Query: c.Request.URL.Query(),
		Body:  b,
	}

	if id, inRequest := c.Params.Get("id"); inRequest {
		req.ResourceId = id
	}

	return req
}

func ginHandler(rh routeHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := ginCtxToRequest(c)

		res := rh(req)

		respb, code := res.ToJSON()

		if code == 500 {
			//log error
		}

		c.JSON(code, respb)
	}
}
