/**
 * @author [author]
 * @email [example@mail.com]
 * @create date 2019-01-16 20:40:52
 * @modify date 2019-01-16 20:40:52
 * @desc [description]
 */

package main

import (
	"errors"
	"fmt"
	"net/http"
	"path"

	captcha "github.com/2637309949/bulrush-captcha"
	"github.com/mojocn/base64Captcha"

	"github.com/2637309949/bulrush"
	delivery "github.com/2637309949/bulrush-delivery"
	identify "github.com/2637309949/bulrush-identify"
	logger "github.com/2637309949/bulrush-logger"
	proxy "github.com/2637309949/bulrush-proxy"
	role "github.com/2637309949/bulrush-role"
	"github.com/2637309949/bulrush-template/addition"
	"github.com/2637309949/bulrush-template/binds"
	"github.com/2637309949/bulrush-template/models"
	"github.com/2637309949/bulrush-template/routes"
	upload "github.com/2637309949/bulrush-upload"
	"github.com/gin-gonic/gin"
)

func appUsePlugins(app bulrush.Bulrush) {
	app.Use(&proxy.Proxy{
		Host:  "https://xxx.com",
		Match: "^/api/v1/proxyTest",
		Map: func(reqPath string) string {
			return reqPath
		},
	})
	app.Use(
		&delivery.Delivery{
			URLPrefix: "/public",
			Path:      path.Join("assets/public", ""),
		},
		&upload.Upload{
			URLPrefix: "/public/upload",
			AssetPath: path.Join("assets/public/upload", ""),
		},
		&logger.Logger{},
		&captcha.Captcha{
			URLPrefix: "/captcha",
			Secret:    "7658388",
			Config: base64Captcha.ConfigDigit{
				Height:     80,
				Width:      240,
				MaxSkew:    0.7,
				DotCount:   80,
				CaptchaLen: 5,
			},
		},
	)
	app.Use(&identify.Identify{
		Auth: func(ctx *gin.Context) (interface{}, error) {
			var login binds.Login
			// captcha := ctx.GetString("captcha")
			if err := ctx.ShouldBind(&login); err != nil {
				return nil, err
			}
			if login.Password == "xx" && login.UserName == "xx" {
				return map[string]interface{}{
					"id":       "3e4r56u80a55",
					"username": login.UserName,
				}, nil
			}
			return nil, errors.New("user authentication failed")
		},
		Tokens: &identify.RedisTokensGroup{
			Redis: addition.Redis,
		},
		FakeTokens: []interface{}{"DEBUG"},
		FakeURLs:   []interface{}{`^/api/v1/ignore$`, `^/api/v1/docs/*`, `^/public/*`, `^/api/v1/ptest$`},
	})
	app.Use(&role.Role{
		FailureHandler: func(c *gin.Context, action string) {
		},
		RoleHandler: func(c *gin.Context, action string) bool {
			actions := role.TransformAction(action)
			fmt.Printf("actions: %s\n", actions)
			return true
		},
	})
	app.Use(&models.Model{}, &routes.Route{})
	app.Use(bulrush.PNQuick(func(testInject string, role *role.Role, router *gin.RouterGroup) {
		router.GET("/bulrushApp", role.Can("r1,r2@p1,p3,p4;r4"), func(c *gin.Context) {
			addition.Logger.Info("from bulrushApp %s", "info")
			addition.Logger.Error("from bulrushApp %s", "error")
			c.JSON(http.StatusOK, gin.H{
				"message": testInject,
			})
		})
	}))
}
