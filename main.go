package main

import (
	"os"
	"path"
	"path/filepath"

	"github.com/curtisnewbie/gocommon/common"
	"github.com/curtisnewbie/gocommon/server"
	"github.com/gin-gonic/gin"
)

func main() {
	common.DefaultReadConfig(os.Args)
	server.AddRoutesRegistar(func(ctx *gin.Engine) {
		ctx.NoRoute(func(c *gin.Context) {
			dir, file := path.Split(c.Request.RequestURI)
			ext := filepath.Ext(file)
			if file == "" || ext == "" {
				c.File("./dist/index.html")
			} else {
				c.File("./dist/" + path.Join(dir, file))
			}
		})
	})
	server.BootstrapServer()
}
