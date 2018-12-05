package main

import (
	// "github.com/gin-gonic/gin"
	"github.com/2637309949/bulrush"
	"github.com/2637309949/bulrush/utils"
	"github.com/2637309949/bulrush_template/routes"
	"github.com/2637309949/bulrush_template/models"
	"github.com/2637309949/bulrush_template/middles"
	"path"
	// "time"
	"fmt"
	"os"
)

// GINMODE APP ENV
var GINMODE 	= utils.Some(os.Getenv("GIN_MODE"), "local")
// CONFIGPATH PATH
var CONFIGPATH  = path.Join(fmt.Sprintf("conf/%s.yaml", GINMODE))

func main() {
	app := bulrush.Default()
	app.LoadConfig(CONFIGPATH)
	app.Inject(&middles.Middles{}, &routes.Routes{}, &models.Model{})
	app.Run()
}
