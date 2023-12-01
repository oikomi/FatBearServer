// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/api/v1/dev/login": {
            "post": {
                "description": "DevLogin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "super_token",
                        "in": "header"
                    },
                    {
                        "description": "dev login",
                        "name": "DevLoginReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dev.DevLoginReq"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/dev/order": {
            "get": {
                "description": "OrderList",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "super_token",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "model 名字",
                        "name": "send_user",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dev.Order"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "send Order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "super_token",
                        "in": "header"
                    },
                    {
                        "description": "send order",
                        "name": "OrderReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dev.OrderReq"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/room/create": {
            "post": {
                "description": "CreateRoom",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "super_token",
                        "in": "header"
                    },
                    {
                        "description": "create room",
                        "name": "createRoom",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/room.CreateRoomReq"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/room/list": {
            "get": {
                "description": "GetRoomList",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "super_token",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/room.Room"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/room/msg": {
            "get": {
                "description": "GetRoomMsg",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "super_token",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/room.RoomMsg"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "GetRoomMsg",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "super_token",
                        "in": "header"
                    },
                    {
                        "description": "send room msg",
                        "name": "SendRoomMsgReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/room.SendRoomMsgReq"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/user/addToken": {
            "post": {
                "description": "user add token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "super_token",
                        "in": "header"
                    },
                    {
                        "description": "AddToken",
                        "name": "AddTokenReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.AddTokenReq"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/v1/user/getToken": {
            "get": {
                "description": "user add token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "super_token",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "用户名称",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "token",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/api/v1/user/login": {
            "post": {
                "description": "user login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authentication header",
                        "name": "super_token",
                        "in": "header"
                    },
                    {
                        "description": "user login",
                        "name": "LoginReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.LoginReq"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/health": {
            "get": {
                "description": "HealthCheck",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "HealthCheck",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dev.DevLoginReq": {
            "type": "object",
            "properties": {
                "dev_name": {
                    "type": "string",
                    "format": "string",
                    "example": "设备名称"
                },
                "model_name": {
                    "type": "string",
                    "format": "string",
                    "example": "设备主播账号名称"
                },
                "password": {
                    "type": "string",
                    "format": "string",
                    "example": "设备密码"
                }
            }
        },
        "dev.Order": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "dev_name": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "model_name": {
                    "type": "string"
                },
                "send_user": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "token": {
                    "type": "integer"
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                },
                "vibration": {
                    "type": "string"
                }
            }
        },
        "dev.OrderReq": {
            "type": "object",
            "properties": {
                "dev_name": {
                    "type": "string",
                    "format": "string",
                    "example": "设备名称"
                },
                "duration": {
                    "type": "integer",
                    "format": "int",
                    "example": 10
                },
                "model_name": {
                    "type": "string",
                    "format": "string",
                    "example": "主播账号名称"
                },
                "send_user": {
                    "type": "string",
                    "format": "string",
                    "example": "观众名称"
                },
                "token": {
                    "type": "integer",
                    "format": "int",
                    "example": 50
                },
                "vibration": {
                    "type": "string",
                    "format": "string",
                    "example": "震动强度, 比如 Medium"
                }
            }
        },
        "room.CreateRoomReq": {
            "type": "object",
            "properties": {
                "creator": {
                    "type": "string",
                    "format": "string",
                    "example": "房间创建者，就是主播账号名称"
                },
                "name": {
                    "type": "string",
                    "format": "string",
                    "example": "房间名称"
                }
            }
        },
        "room.Room": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "creator": {
                    "type": "string",
                    "format": "string",
                    "example": "房间创建者，就是主播账号名称"
                },
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "room_name": {
                    "type": "string",
                    "format": "string",
                    "example": "房间名称"
                },
                "room_status": {
                    "type": "integer",
                    "format": "int",
                    "example": 0
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        },
        "room.RoomMsg": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                },
                "room_name": {
                    "type": "string"
                },
                "send_user": {
                    "type": "string"
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        },
        "room.SendRoomMsgReq": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "send_user": {
                    "type": "string"
                }
            }
        },
        "user.AddTokenReq": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "format": "string",
                    "example": "充值的用户名称"
                },
                "token": {
                    "type": "integer",
                    "format": "int",
                    "example": 1000
                }
            }
        },
        "user.LoginReq": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server Petstore server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
