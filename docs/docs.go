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
        "/cart/add-product": {
            "post": {
                "description": "Adds a Product to Customer's Cart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "parameters": [
                    {
                        "description": "AddProductPayload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.AddProductPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "{\\\"error\\\": \\\"Internal Server Error\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/cart/edit-product": {
            "post": {
                "description": "Edits a Product from Customer's Cart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "parameters": [
                    {
                        "description": "EditProductPayload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.EditProductPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "{\\\"error\\\": \\\"Internal Server Error\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/cart/overview": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "parameters": [
                    {
                        "description": "GetCartPayload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.GetCartPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/cart/remove-product": {
            "post": {
                "description": "Removes a Product from Customer's Cart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "parameters": [
                    {
                        "description": "RemoveProductPayload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.RemoveProductPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/customer": {
            "get": {
                "description": "Overview all customer's list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Overview customer list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/customer.Customer"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Create customer",
                "parameters": [
                    {
                        "description": "CustomerPayload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.CustomerPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "customer.Customer",
                        "schema": {
                            "$ref": "#/definitions/customer.Customer"
                        }
                    },
                    "400": {
                        "description": "{\\\"error\\\": \\\"Invalid CPF\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/customer/{cpf}": {
            "get": {
                "description": "Overview a customer by cpf",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Overview customer by cpf",
                "parameters": [
                    {
                        "type": "string",
                        "description": "customer cpf",
                        "name": "cpf",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/customer.Customer"
                        }
                    },
                    "404": {
                        "description": "Customer not found"
                    }
                }
            }
        },
        "/customer/{id}": {
            "put": {
                "description": "Update a customer by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Update customer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the customer",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "CustomerPayload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.CustomerPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "customer.Customer",
                        "schema": {
                            "$ref": "#/definitions/customer.Customer"
                        }
                    },
                    "400": {
                        "description": "{\\\"error\\\": \\\"Invalid CPF\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a customer by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Delete customer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "123456789",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Invalid identifier informed"
                    }
                }
            }
        },
        "/orders": {
            "get": {
                "description": "Get all order's list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Get order list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/order.Order"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "Create order from Cart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Create order from Cart",
                "parameters": [
                    {
                        "description": "CreateOrderPayload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.CreateOrderPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/order.Order"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    }
                }
            }
        },
        "/orders/{id}": {
            "get": {
                "description": "Get an order by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Get order by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the order",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/order.Order"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update by json an Order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Update an Order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the Order",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "UpdateOrderPayload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.UpdateOrderPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/order.Order"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.HTTPError"
                        }
                    }
                }
            }
        },
        "/product": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Search products with given criteria",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of Product",
                        "name": "search",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "description": "Category of Product",
                        "name": "category",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "product.Product",
                        "schema": {
                            "$ref": "#/definitions/product.Product"
                        }
                    },
                    "500": {
                        "description": "{\\\"error\\\": \\\"something went bad :(\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/product/{id}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Delete a Product with given ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of Product",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "500": {
                        "description": "{\\\"error\\\": \\\"something went bad :(\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.AddProductPayload": {
            "type": "object",
            "required": [
                "client_id",
                "product_id",
                "quantity"
            ],
            "properties": {
                "client_id": {
                    "type": "string",
                    "format": "uuid"
                },
                "comments": {
                    "type": "string"
                },
                "product_id": {
                    "type": "string",
                    "format": "uuid"
                },
                "quantity": {
                    "type": "integer",
                    "minimum": 1,
                    "example": 1
                }
            }
        },
        "controller.CreateOrderPayload": {
            "type": "object",
            "required": [
                "client_id"
            ],
            "properties": {
                "client_id": {
                    "type": "string",
                    "format": "uuid"
                }
            }
        },
        "controller.CustomerPayload": {
            "type": "object",
            "required": [
                "cpf",
                "email",
                "name"
            ],
            "properties": {
                "age": {
                    "type": "integer",
                    "maximum": 120,
                    "minimum": 18
                },
                "cpf": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 5
                }
            }
        },
        "controller.EditProductPayload": {
            "type": "object",
            "required": [
                "client_id",
                "product_id",
                "quantity"
            ],
            "properties": {
                "client_id": {
                    "type": "string",
                    "format": "uuid"
                },
                "comments": {
                    "type": "string"
                },
                "product_id": {
                    "type": "string",
                    "format": "uuid"
                },
                "quantity": {
                    "type": "integer",
                    "example": 2
                }
            }
        },
        "controller.GetCartPayload": {
            "type": "object",
            "required": [
                "client_id"
            ],
            "properties": {
                "client_id": {
                    "type": "string",
                    "format": "uuid"
                }
            }
        },
        "controller.HTTPError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "controller.RemoveProductPayload": {
            "type": "object",
            "required": [
                "client_id",
                "product_id"
            ],
            "properties": {
                "client_id": {
                    "type": "string",
                    "format": "uuid"
                },
                "product_id": {
                    "type": "string",
                    "format": "uuid"
                }
            }
        },
        "controller.UpdateOrderPayload": {
            "type": "object",
            "required": [
                "status"
            ],
            "properties": {
                "status": {
                    "type": "string",
                    "example": "paid"
                }
            }
        },
        "customer.Customer": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "cpf": {
                    "type": "string"
                },
                "email": {
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
        "order.Order": {
            "type": "object",
            "properties": {
                "client_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "total_amount": {
                    "type": "number"
                }
            }
        },
        "product.Product": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Tech Challenge Food API",
	Description:      "Fast Food API for FIAP Tech course",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
