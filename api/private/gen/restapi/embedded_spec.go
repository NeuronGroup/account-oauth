// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Oauth Private API",
    "contact": {
      "name": "mars"
    },
    "version": "v1"
  },
  "basePath": "/private-api/v1/oauth",
  "paths": {
    "/authorize": {
      "post": {
        "operationId": "Authorize",
        "parameters": [
          {
            "type": "string",
            "name": "jwt",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "response_type",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "client_id",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "redirect_uri",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "scope",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "state",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/AuthorizationCode"
            }
          },
          "default": {
            "description": "Error response",
            "schema": {
              "$ref": "#/definitions/authorizeDefaultBody"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "AuthorizationCode": {
      "type": "object",
      "properties": {
        "code": {
          "type": "string"
        },
        "expiresSeconds": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "authorizeDefaultBody": {
      "type": "object",
      "properties": {
        "code": {
          "description": "Error code",
          "type": "string"
        },
        "errors": {
          "$ref": "#/definitions/authorizeDefaultBodyErrors"
        },
        "message": {
          "description": "Error message",
          "type": "string"
        },
        "status": {
          "type": "string",
          "format": "int32",
          "default": "Http status"
        }
      },
      "x-go-gen-location": "operations"
    },
    "authorizeDefaultBodyErrors": {
      "description": "Errors",
      "type": "array",
      "items": {
        "$ref": "#/definitions/authorizeDefaultBodyErrorsItems"
      },
      "x-go-gen-location": "operations"
    },
    "authorizeDefaultBodyErrorsItems": {
      "type": "object",
      "properties": {
        "code": {
          "description": "error code",
          "type": "string"
        },
        "field": {
          "description": "field name",
          "type": "string"
        },
        "message": {
          "description": "error message",
          "type": "string"
        }
      },
      "x-go-gen-location": "operations"
    }
  },
  "responses": {
    "ErrorResponse": {
      "description": "Error response",
      "schema": {
        "type": "object",
        "properties": {
          "code": {
            "description": "Error code",
            "type": "string"
          },
          "errors": {
            "$ref": "#/definitions/authorizeDefaultBodyErrors"
          },
          "message": {
            "description": "Error message",
            "type": "string"
          },
          "status": {
            "type": "string",
            "format": "int32",
            "default": "Http status"
          }
        }
      }
    }
  },
  "securityDefinitions": {
    "Basic": {
      "type": "basic"
    }
  }
}`))
}
