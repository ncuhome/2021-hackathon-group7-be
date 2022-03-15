// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "email": "316851756@qq.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/activity": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "活动系统"
                ],
                "summary": "获取活动详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "活动id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/auth/activity": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "活动系统"
                ],
                "summary": "社团账号修改活动",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "活动id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": " ",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Activity"
                        }
                    }
                ],
                "responses": {}
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "活动系统"
                ],
                "summary": "社团账号发布活动",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": " ",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Activity"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "活动系统"
                ],
                "summary": "社团账号删除活动",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "活动id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/auth/org/ended-activity": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "活动系统"
                ],
                "summary": "分页获取组织历史活动列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "上一次调用本接口得到的活动列表的最后一个活动的结束时间戳，第一次调用用当前时间戳",
                        "name": "pre",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/auth/org/not-ended-activity": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "活动系统"
                ],
                "summary": "分页获取组织非历史活动列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "上一次调用本接口得到的活动列表的最后一个活动的结束时间戳，第一次调用用当前时间戳",
                        "name": "pre",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/auth/organization": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户系统"
                ],
                "summary": "获取组织的负责人管理的组织名称及是否激活",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户系统"
                ],
                "summary": "社团账号激活（注册）、修改",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": " ",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.OrgInfo"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/auth/picture": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "其它"
                ],
                "summary": "图片上传接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "file字段放图片数据",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/auth/token": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户系统"
                ],
                "summary": "检验token是否有效，若有效返回用户角色和新token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户令牌",
                        "name": "Token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/during-activity": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "活动系统"
                ],
                "summary": "分页获取正在进行的活动列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "上一次调用本接口得到的活动列表的最后一个活动的开始时间戳，第一次调用用当前时间戳",
                        "name": "pre",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "当前时间戳",
                        "name": "now",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/ended-activity": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "活动系统"
                ],
                "summary": "分页获取历史活动列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "上一次调用本接口得到的活动列表的最后一个活动的结束时间戳，第一次调用用当前时间戳",
                        "name": "pre",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户系统"
                ],
                "summary": "云家园账号或社团账号登录",
                "parameters": [
                    {
                        "description": " ",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Login"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/login/ncuos-token": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户系统"
                ],
                "summary": "云家园账号token登录",
                "parameters": [
                    {
                        "description": " ",
                        "name": "JSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Token"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/not-start-activity": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "活动系统"
                ],
                "summary": "分页获取未开始的活动列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "上一次调用本接口得到的活动列表的最后一个活动的开始时间戳，第一次调用用当前时间戳",
                        "name": "pre",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/recommend-activity": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "活动系统"
                ],
                "summary": "分页获取推荐活动列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "上一次调用本接口得到的活动列表的最后一个活动的开始时间戳，第一次调用用当前时间戳",
                        "name": "pre",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/user-info": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户系统"
                ],
                "summary": "获取用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "dto.Activity": {
            "type": "object",
            "required": [
                "content",
                "end_time",
                "place",
                "start_time",
                "title"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "end_time": {
                    "type": "string"
                },
                "place": {
                    "type": "string"
                },
                "start_time": {
                    "description": "Time：自1970年1月1日00:00:00 UTC以来经过的毫秒数",
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dto.Login": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.OrgInfo": {
            "type": "object",
            "properties": {
                "logo_url": {
                    "type": "string"
                },
                "password": {
                    "description": "长度大于8",
                    "type": "string"
                }
            }
        },
        "dto.Token": {
            "type": "object",
            "required": [
                "token"
            ],
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "tudo-api-test.ncuos.com",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Swagger API",
	Description: "给出了请求方法，点击Model可以查看请求体模型及备注。\n查看响应体需要打开浏览器开发者工具，在页面接口初Try it out，填写数据然后Execute。",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
