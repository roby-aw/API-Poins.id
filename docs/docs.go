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
        "/account": {
            "put": {
                "description": "Updatecustomer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Updatecustomer",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Register",
                        "name": "Registercustomer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/customermitra.UpdateCustomer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/customermitra.UpdateCustomer"
                        }
                    }
                }
            }
        },
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
                    "Admin"
                ],
                "summary": "Create admin",
                "parameters": [
                    {
                        "description": "admin",
                        "name": "admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.RegisterAdmin"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/admin.RegisterAdmin"
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
        "/admin/approve/{transactionid}": {
            "post": {
                "description": "Approve Transaction",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Approve Transaction",
                "parameters": [
                    {
                        "type": "string",
                        "description": "transaction_id",
                        "name": "transactionid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/admin/login": {
            "post": {
                "description": "Login admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Login admin",
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
        "/cashout": {
            "post": {
                "description": "Redeem Emoney customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Redeem"
                ],
                "summary": "Redeem Cashout",
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
        "/emoney": {
            "post": {
                "description": "Redeem Emoney customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Redeem"
                ],
                "summary": "Redeem Emoney",
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
        "/login": {
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
        "/paketdata": {
            "post": {
                "description": "Redeem PaketData customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Redeem"
                ],
                "summary": "Redeem PaketData",
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
        "/pulsa": {
            "post": {
                "description": "Redeem Pulsa customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Redeem"
                ],
                "summary": "Redeem Pulsa",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "inputdataemoney",
                        "name": "InputDataCashout",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/customermitra.RedeemPulsaData"
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
                    "Customer"
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
        }
    },
    "definitions": {
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
        "admin.RegisterAdmin": {
            "type": "object",
            "required": [
                "email",
                "fullname",
                "no_hp",
                "password"
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
        "customermitra.RedeemPulsaData": {
            "type": "object",
            "required": [
                "amount",
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
