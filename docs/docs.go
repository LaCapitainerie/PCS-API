// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/sidebar": {
            "get": {
                "description": "Récupère tous les Sidebar",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "administration"
                ],
                "summary": "Sidebar",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Sidebar"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Sidebar": {
            "type": "object",
            "properties": {
                "Hover": {
                    "type": "string"
                },
                "Href": {
                    "type": "string"
                },
                "Icon": {
                    "type": "string"
                },
                "Id_tab": {
                    "type": "string"
                },
                "Permission": {
                    "type": "integer"
                }
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
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
