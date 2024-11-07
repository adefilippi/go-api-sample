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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/models": {
            "get": {
                "description": "get all models",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Model"
                ],
                "summary": "Show all models",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Model"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    }
                }
            },
            "post": {
                "description": "Add by json model",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Model"
                ],
                "summary": "Add a model",
                "parameters": [
                    {
                        "description": "Add account",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Model"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Model"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    }
                }
            }
        },
        "/models/{id}": {
            "get": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Model"
                ],
                "summary": "Show an account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Model ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Model"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete by model ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Model"
                ],
                "summary": "Delete a model",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Model ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/entity.Model"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update by json Model",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Model"
                ],
                "summary": "Update a Model",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Model ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update model",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Model"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Model"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    }
                }
            }
        },
        "/models/{id}/image": {
            "post": {
                "description": "Post by model image ID",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Model - image"
                ],
                "summary": "Post a model image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Model ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Image file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Tag for the image",
                        "name": "tag",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.MediaObject"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    }
                }
            }
        },
        "/models/{id}/image/{image-id}": {
            "get": {
                "description": "Get by model image ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "Model - image"
                ],
                "summary": "Get a model image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Model ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Model Image ID",
                        "name": "image-id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Image or file",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete by  model image ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Model - image"
                ],
                "summary": "Delete a model image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Model ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Model Image ID",
                        "name": "image-id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/entity.MediaObject"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ApiError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.ApiError": {
            "type": "object",
            "properties": {
                "detail": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "entity.MediaObject": {
            "type": "object",
            "properties": {
                "association": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "dimensions": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "id": {
                    "description": "Define custom UUID ID",
                    "type": "string"
                },
                "mime_type": {
                    "type": "string"
                },
                "model_id": {
                    "description": "Foreign key for Model",
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "original_name": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "size": {
                    "type": "integer"
                },
                "tag": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "entity.Model": {
            "type": "object",
            "properties": {
                "alt_image": {
                    "type": "string"
                },
                "body_type": {
                    "type": "string"
                },
                "cargo_volume": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "displayable": {
                    "type": "boolean"
                },
                "emission_co2": {
                    "type": "string"
                },
                "encrypt": {
                    "type": "string"
                },
                "engine_type": {
                    "type": "string"
                },
                "fuel_consumption": {
                    "type": "string"
                },
                "fuel_type": {
                    "type": "string"
                },
                "hybrid_system": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "images": {
                    "description": "Correct association",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.MediaObject"
                    }
                },
                "is_new": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "position": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "random": {
                    "type": "string"
                },
                "seat_capacity": {
                    "type": "integer"
                },
                "settings_link": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                },
                "sub_title": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "vehicle_engine": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "Authentication with Bearer Token (JWT)",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BearerAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Commercial Info API",
	Description:      "Authentication with Bearer Token (JWT)",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
