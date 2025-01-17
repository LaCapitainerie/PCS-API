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
        "/api/Admin": {
            "get": {
                "description": "Récupère tous les Admin",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "administration"
                ],
                "summary": "Admin",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Admin"
                            }
                        }
                    }
                }
            }
        },
        "/api/Property": {
            "get": {
                "description": "Récupère tous les Property",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "administration"
                ],
                "summary": "Property",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Property"
                            }
                        }
                    }
                }
            }
        },
        "/api/Property_image": {
            "get": {
                "description": "Récupère tous les PropertyImage",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "PropertyImage"
                ],
                "summary": "PropertyImage",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.PropertyImage"
                            }
                        }
                    }
                }
            }
        },
        "/api/Traveler": {
            "get": {
                "description": "Récupère tous les Traveler",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "administration"
                ],
                "summary": "Traveler",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Traveler"
                            }
                        }
                    }
                }
            }
        },
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
        },
        "/api/user/login": {
            "post": {
                "description": "Se connecte à un utilisateur",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Connexion"
                ],
                "summary": "User",
                "parameters": [
                    {
                        "description": "Mail de l'utilisateur",
                        "name": "mail",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Mot de passe de l'utilisateur",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Retourne un token de connexion",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Requête incorrecte - données invalides",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "Conflit - L'email ou le mot de passe existe déjà",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/user/register": {
            "post": {
                "description": "Crée un nouvel utilisateur",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Création"
                ],
                "summary": "User",
                "parameters": [
                    {
                        "description": "User to create",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UsersDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Retourne l'utilisateur crée",
                        "schema": {
                            "$ref": "#/definitions/models.UsersDTO"
                        }
                    },
                    "400": {
                        "description": "Requête incorrecte - données invalides",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "Conflit - L'email ou le mot de passe existe déjà",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Admin": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "site": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "models.Property": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "administrationValidation": {
                    "type": "boolean"
                },
                "bathroom": {
                    "type": "integer"
                },
                "city": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "garage": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "lessorId": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "room": {
                    "type": "integer"
                },
                "surface": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                },
                "zipCode": {
                    "type": "string"
                }
            }
        },
        "models.PropertyImage": {
            "type": "object",
            "properties": {
                "ID": {
                    "type": "string"
                },
                "Path": {
                    "type": "string"
                },
                "property_id": {
                    "type": "string"
                }
            }
        },
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
        },
        "models.Traveler": {
            "type": "object",
            "properties": {
                "Id": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "models.UsersDTO": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lastConnectionDate": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "mail": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "registerDate": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
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
