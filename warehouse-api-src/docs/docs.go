// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/articles": {
            "get": {
                "tags": [
                    "articles"
                ],
                "summary": "Returns all articles",
                "responses": {}
            },
            "patch": {
                "tags": [
                    "articles"
                ],
                "summary": "Updates inventory (articles)",
                "parameters": [
                    {
                        "description": "articles",
                        "name": "articles",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/products": {
            "get": {
                "tags": [
                    "products"
                ],
                "summary": "Returns all products",
                "responses": {}
            },
            "patch": {
                "tags": [
                    "products"
                ],
                "summary": "Updates product definitions",
                "parameters": [
                    {
                        "description": "products",
                        "name": "products",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/products/{id}/sale": {
            "post": {
                "tags": [
                    "products"
                ],
                "summary": "Sells one specified product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Warehouse API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
