// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-07-02 20:26:30.351614976 +0800 CST m=+0.044448705

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "bulrush web library template",
        "title": "bulrush-template api",
        "termsOfService": "https://github.com/2637309949/bulrush",
        "contact": {
            "name": "bulrush-template",
            "url": "https://github.com/2637309949/bulrush",
            "email": "2637309949@qq.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "1.0"
    },
    "host": "127.0.0.1:8080",
    "basePath": "/api/v1",
    "paths": {
        "/chache": {
            "get": {
                "description": "测试缓存路由",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cache"
                ],
                "summary": "缓存路由",
                "parameters": [
                    {
                        "type": "string",
                        "description": "令牌",
                        "name": "accessToken",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\": \"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"message\": \"failure\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/mgoTest": {
            "get": {
                "description": "Task测试",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "TASK"
                ],
                "summary": "Task测试",
                "parameters": [
                    {
                        "type": "string",
                        "description": "令牌",
                        "name": "accessToken",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\": \"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"message\": \"failure\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/testsql": {
            "get": {
                "description": "MGO测试",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GORM"
                ],
                "summary": "MGO测试",
                "parameters": [
                    {
                        "type": "string",
                        "description": "令牌",
                        "name": "accessToken",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\": \"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "{\"message\": \"failure\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
