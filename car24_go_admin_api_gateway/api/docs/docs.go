// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/v1/brand": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for getting all brands",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "brand"
                ],
                "summary": "Get brands",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 10,
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/car24_car_service.BrandListModel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for creating brand",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "brand"
                ],
                "summary": "Create Brand",
                "parameters": [
                    {
                        "type": "string",
                        "default": "7d4a4c38-dd84-4902-b744-0488b80a4c01",
                        "description": "platform-id",
                        "name": "Platform-Id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "brand",
                        "name": "Brand",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/car24_car_service.CreateBrand"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseOK"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/brand/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for getting brand",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "brand"
                ],
                "summary": "Get Brand",
                "parameters": [
                    {
                        "type": "string",
                        "default": "7d4a4c38-dd84-4902-b744-0488b80a4c01",
                        "description": "platform-id",
                        "name": "Platform-Id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/car24_car_service.Brand"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
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
                "description": "API for updating brand",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "brand"
                ],
                "summary": "Update Brand",
                "parameters": [
                    {
                        "type": "string",
                        "default": "7d4a4c38-dd84-4902-b744-0488b80a4c01",
                        "description": "platform-id",
                        "name": "Platform-Id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "car",
                        "name": "Brand",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/car24_car_service.UpdateBrand"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseOK"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for deleting brand",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "brand"
                ],
                "summary": "Delete Brand",
                "parameters": [
                    {
                        "type": "string",
                        "default": "7d4a4c38-dd84-4902-b744-0488b80a4c01",
                        "description": "platform-id",
                        "name": "Platform-Id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseOK"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/car": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for getting all cars",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "car"
                ],
                "summary": "Get cars",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 10,
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "model_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/car24_car_service.CarListModel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for creating car",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "car"
                ],
                "summary": "Create Car",
                "parameters": [
                    {
                        "type": "string",
                        "default": "7d4a4c38-dd84-4902-b744-0488b80a4c01",
                        "description": "platform-id",
                        "name": "Platform-Id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "car",
                        "name": "Car",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/car24_car_service.CreateCarModel"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseOK"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/car/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for getting car",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "car"
                ],
                "summary": "Get Car",
                "parameters": [
                    {
                        "type": "string",
                        "default": "7d4a4c38-dd84-4902-b744-0488b80a4c01",
                        "description": "platform-id",
                        "name": "Platform-Id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/car24_car_service.CarModel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
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
                "description": "API for updating car",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "car"
                ],
                "summary": "Update Car",
                "parameters": [
                    {
                        "type": "string",
                        "default": "7d4a4c38-dd84-4902-b744-0488b80a4c01",
                        "description": "platform-id",
                        "name": "Platform-Id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "car",
                        "name": "Car",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/car24_car_service.UpdateCarModel"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseOK"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for deleting car",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "car"
                ],
                "summary": "Delete Car",
                "parameters": [
                    {
                        "type": "string",
                        "default": "7d4a4c38-dd84-4902-b744-0488b80a4c01",
                        "description": "platform-id",
                        "name": "Platform-Id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseOK"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/model": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for getting all models",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "model"
                ],
                "summary": "Get models",
                "parameters": [
                    {
                        "type": "string",
                        "name": "brand_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/car24_car_service.ModelListModel"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for creating model",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "model"
                ],
                "summary": "Create Model",
                "parameters": [
                    {
                        "type": "string",
                        "default": "7d4a4c38-dd84-4902-b744-0488b80a4c01",
                        "description": "platform-id",
                        "name": "Platform-Id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "model",
                        "name": "Model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/car24_car_service.CreateModel"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseOK"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
                        }
                    }
                }
            }
        },
        "/v1/model/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for getting model",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "model"
                ],
                "summary": "Get Model",
                "parameters": [
                    {
                        "type": "string",
                        "default": "7d4a4c38-dd84-4902-b744-0488b80a4c01",
                        "description": "platform-id",
                        "name": "Platform-Id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/car24_car_service.Model"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
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
                "description": "API for updating model",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "model"
                ],
                "summary": "Update Model",
                "parameters": [
                    {
                        "type": "string",
                        "default": "7d4a4c38-dd84-4902-b744-0488b80a4c01",
                        "description": "platform-id",
                        "name": "Platform-Id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "car",
                        "name": "Model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/car24_car_service.UpdateModel"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseOK"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "API for deleting model",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "model"
                ],
                "summary": "Delete Model",
                "parameters": [
                    {
                        "type": "string",
                        "default": "7d4a4c38-dd84-4902-b744-0488b80a4c01",
                        "description": "platform-id",
                        "name": "Platform-Id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseOK"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "car24_car_service.Brand": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "car24_car_service.BrandListModel": {
            "type": "object",
            "properties": {
                "cars": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/car24_car_service.Brand"
                    }
                },
                "count": {
                    "type": "integer"
                }
            }
        },
        "car24_car_service.CarListModel": {
            "type": "object",
            "properties": {
                "cars": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/car24_car_service.CarModel"
                    }
                },
                "count": {
                    "type": "integer"
                }
            }
        },
        "car24_car_service.CarModel": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "model_data": {
                    "$ref": "#/definitions/car24_car_service.Model"
                },
                "model_id": {
                    "type": "string"
                },
                "state_number": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "car24_car_service.CreateBrand": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "car24_car_service.CreateCarModel": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "model_id": {
                    "type": "string"
                },
                "state_number": {
                    "type": "string"
                }
            }
        },
        "car24_car_service.CreateModel": {
            "type": "object",
            "properties": {
                "brand_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "car24_car_service.Model": {
            "type": "object",
            "properties": {
                "brand_data": {
                    "$ref": "#/definitions/car24_car_service.Brand"
                },
                "brand_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "car24_car_service.ModelListModel": {
            "type": "object",
            "properties": {
                "cars": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/car24_car_service.Model"
                    }
                },
                "count": {
                    "type": "integer"
                }
            }
        },
        "car24_car_service.UpdateBrand": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "car24_car_service.UpdateCarModel": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "model_id": {
                    "type": "string"
                },
                "state_number": {
                    "type": "string"
                }
            }
        },
        "car24_car_service.UpdateModel": {
            "type": "object",
            "properties": {
                "brand_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "response.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "response.ResponseError": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/response.Error"
                }
            }
        },
        "response.ResponseOK": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
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
