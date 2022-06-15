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
        "/admin": {
            "post": {
                "description": "create admin with data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Create admin",
                "parameters": [
                    {
                        "description": "admin",
                        "name": "admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.AdminSwagger"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/admin.Admin"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/admin/login": {
            "post": {
                "description": "Get token for admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Get Token",
                "parameters": [
                    {
                        "description": "admin",
                        "name": "admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.AuthLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": ""
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Register customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "Register",
                        "name": "Registercustomer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/customermitra.RegisterCustomer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/customermitra.RegisterCustomer"
                        }
                    }
                }
            }
        },
        "/v1/login": {
            "post": {
                "description": "Login Customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Customer",
                        "name": "Customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/customermitra.AuthLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Login"
                        }
                    }
                }
            }
        },
        "/v1/order/cashout": {
            "post": {
                "description": "Emoney customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customerOrder"
                ],
                "summary": "Order Emoney/Cashout",
                "parameters": [
                    {
                        "description": "inputdataemoney",
                        "name": "InputDataCashout",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/customermitra.InputTransactionBankEmoney"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/customermitra.InputTransactionBankEmoney"
                        }
                    }
                }
            }
        },
        "/v1/order/emoney": {
            "post": {
                "description": "Emoney customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customerOrder"
                ],
                "summary": "Order Emoney/Cashout",
                "parameters": [
                    {
                        "description": "inputdataemoney",
                        "name": "InputDataCashout",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/customermitra.InputTransactionBankEmoney"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/customermitra.InputTransactionBankEmoney"
                        }
                    }
                }
            }
        },
        "/v1/order/paketdata": {
            "post": {
                "description": "PaketData customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customerOrder"
                ],
                "summary": "Order PaketData",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/v1/order/pulsa": {
            "post": {
                "description": "Pulsa customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customerOrder"
                ],
                "summary": "Order Pulsa",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
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
        "admin.Admin": {
            "type": "object",
            "required": [
                "email",
                "fullname",
                "no_hp",
                "password"
            ],
            "properties": {
                "createdat": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "no_hp": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "updatedat": {
                    "type": "string"
                }
            }
        },
        "admin.AdminSwagger": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "admin.AuthLogin": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "customermitra.AuthLogin": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "customermitra.InputTransactionBankEmoney": {
            "type": "object",
            "required": [
                "amount",
                "an_rekening",
                "bank_provider",
                "customer_id",
                "nomor",
                "poin_account",
                "poin_redeem"
            ],
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "an_rekening": {
                    "type": "string"
                },
                "bank_provider": {
                    "type": "string"
                },
                "customer_id": {
                    "type": "integer"
                },
                "nomor": {
                    "type": "string"
                },
                "poin_account": {
                    "type": "integer"
                },
                "poin_redeem": {
                    "type": "integer"
                }
            }
        },
        "customermitra.RegisterCustomer": {
            "type": "object",
            "required": [
                "email",
                "fullname",
                "no_hp",
                "password",
                "pin"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "no_hp": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "pin": {
                    "type": "integer"
                }
            }
        },
        "customermitra.ResponseLogin": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "pin": {
                    "type": "integer"
                },
                "poin": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "customermitra.UpdateCustomer": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "no_hp": {
                    "type": "string"
                }
            }
        },
        "response.Login": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "messages": {
                    "type": "string"
                },
                "results": {
                    "$ref": "#/definitions/customermitra.ResponseLogin"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "api-poins-id.herokuapp.com/v1",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "API Jasa Pengiriman",
	Description:      "Berikut API Jasa Pengiriman",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
