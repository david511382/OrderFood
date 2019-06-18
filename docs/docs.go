// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-06-18 12:05:03.0718557 +0800 CST m=+0.133015601

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "訂餐系統",
        "title": "Order Food API",
        "contact": {},
        "license": {},
        "version": "1.0"
    },
    "host": "{{.Host}}",
    "basePath": "/api/",
    "paths": {
        "/auth/register": {
            "post": {
                "description": "註冊",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "註冊",
                "parameters": [
                    {
                        "type": "string",
                        "description": "稱號",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "帳號",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密碼",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu": {
            "get": {
                "description": "取得菜單",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shop"
                ],
                "summary": "取得菜單",
                "responses": {
                    "200": {
                        "description": "菜單",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu/item": {
            "post": {
                "description": "新增商品",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "新增商品",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "商店",
                        "name": "shopID",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "商名",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "價格",
                        "name": "price",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "菜單",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/resp.Item"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu/item/{id}": {
            "put": {
                "description": "修改商品",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "修改商品",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "編號",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "品名",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "價格",
                        "name": "price",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "結果",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "刪除商品",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "刪除商品",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "編號",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu/item/{shopID}": {
            "get": {
                "description": "取得商品",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "取得商品",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "商店編號",
                        "name": "shopID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "選單編號",
                        "name": "optionID",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "菜單",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/resp.Item"
                            }
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu/itemOption": {
            "post": {
                "description": "商品加入選單",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "商品加入選單",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "商品編號",
                        "name": "itemID",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "選單編號",
                        "name": "optionID",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "結果",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/resp.ItemOption"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu/itemOption/{id}": {
            "delete": {
                "description": "刪除選單的商品",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "刪除選單的商品",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu/menu/{shop}": {
            "get": {
                "description": "取得菜單",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "取得菜單",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商店",
                        "name": "shop",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "菜單",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/resp.ShopMenu"
                            }
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu/option": {
            "post": {
                "description": "新增商品選單",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "新增商品選單",
                "parameters": [
                    {
                        "type": "string",
                        "description": "選單選項",
                        "name": "selectionName",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "必選數",
                        "name": "selectNum",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "商品選單",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/resp.Option"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu/option/{id}": {
            "get": {
                "description": "刪除商品選單",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "刪除商品選單",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "修改商品選單",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "修改商品選單",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "商品選單",
                        "name": "selectNum",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu/selection": {
            "post": {
                "description": "新增選單選項",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "新增選單選項",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "選單編號",
                        "name": "optionID",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "名稱",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "價格",
                        "name": "price",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "菜單",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/resp.MenuSelection"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu/selection/{id}": {
            "get": {
                "description": "刪除選單選項",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "刪除選單選項",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "修改選單選項",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "修改選單選項",
                "parameters": [
                    {
                        "type": "string",
                        "description": "名稱",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "價格",
                        "name": "price",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "結果",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu/shop": {
            "get": {
                "description": "取得商店",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "取得商店",
                "parameters": [
                    {
                        "type": "string",
                        "description": "編號",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "商名",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "菜單",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/resp.Shop"
                            }
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "新增商店",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "新增商店",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商名",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "菜單",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/resp.Shop"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu/shop/{id}": {
            "put": {
                "description": "修改商店",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "修改商店",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "編號",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "店名",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "結果",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "刪除商店",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "刪除商店",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "編號",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu/shopmenu": {
            "get": {
                "description": "取得菜單",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "取得菜單",
                "responses": {
                    "200": {
                        "description": "菜單",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/resp.ShopMenu"
                            }
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/order": {
            "get": {
                "description": "取得訂單",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "取得訂單",
                "responses": {
                    "200": {
                        "description": "餐點",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "訂餐",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "訂餐",
                "parameters": [
                    {
                        "type": "string",
                        "description": "餐點",
                        "name": "orders",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "餐點",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/order/all": {
            "get": {
                "description": "取得所有訂單",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "取得所有訂單",
                "responses": {
                    "200": {
                        "description": "餐點",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/shop": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "更改商店",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "manager"
                ],
                "summary": "更改商店",
                "parameters": [
                    {
                        "type": "string",
                        "description": "商店",
                        "name": "view",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "商店",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "取得用戶名稱",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "取得用戶名稱",
                "responses": {
                    "200": {
                        "description": "用戶名稱",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "修改用戶名稱和密碼",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "修改用戶資訊",
                "parameters": [
                    {
                        "description": "修改用戶資訊",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/reqs.ModifyUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "用戶資訊",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Member"
                        }
                    },
                    "500": {
                        "description": "内部错误",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Member": {
            "type": "object",
            "properties": {
                "ID": {
                    "type": "integer"
                },
                "Password": {
                    "type": "string"
                },
                "Username": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "reqs.ModifyUser": {
            "type": "object",
            "properties": {
                "Name": {
                    "type": "string"
                },
                "Password": {
                    "type": "string"
                },
                "Username": {
                    "type": "string"
                }
            }
        },
        "resp.Item": {
            "type": "object",
            "properties": {
                "ID": {
                    "type": "integer"
                },
                "Name": {
                    "type": "string"
                },
                "Options": {
                    "type": "string"
                },
                "Price": {
                    "type": "integer"
                }
            }
        },
        "resp.ItemOption": {
            "type": "object",
            "properties": {
                "ID": {
                    "type": "integer"
                },
                "ItemID": {
                    "type": "integer"
                },
                "OptionID": {
                    "type": "integer"
                }
            }
        },
        "resp.MenuOption": {
            "type": "object",
            "properties": {
                "Items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/resp.Item"
                    }
                },
                "Name": {
                    "type": "string"
                },
                "Option": {
                    "type": "object",
                    "$ref": "#/definitions/resp.Option"
                },
                "Selections": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/resp.MenuSelection"
                    }
                }
            }
        },
        "resp.MenuSelection": {
            "type": "object",
            "properties": {
                "ID": {
                    "type": "integer"
                },
                "Name": {
                    "type": "string"
                },
                "Price": {
                    "type": "integer"
                }
            }
        },
        "resp.Option": {
            "type": "object",
            "properties": {
                "ID": {
                    "type": "integer"
                },
                "SelectNum": {
                    "type": "integer"
                }
            }
        },
        "resp.Shop": {
            "type": "object",
            "properties": {
                "ID": {
                    "type": "integer"
                },
                "Name": {
                    "type": "string"
                }
            }
        },
        "resp.ShopMenu": {
            "type": "object",
            "properties": {
                "Options": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/resp.MenuOption"
                    }
                },
                "Shop": {
                    "type": "object",
                    "$ref": "#/definitions/resp.Shop"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
