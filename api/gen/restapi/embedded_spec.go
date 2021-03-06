// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Oauth API",
    "contact": {
      "name": "mars"
    },
    "version": "v1"
  },
  "basePath": "/api/v1/oauth",
  "paths": {
    "/me": {
      "get": {
        "operationId": "Me",
        "parameters": [
          {
            "type": "string",
            "name": "access_token",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/token": {
      "post": {
        "security": [
          {
            "Basic": []
          }
        ],
        "operationId": "Token",
        "parameters": [
          {
            "type": "string",
            "name": "grant_type",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "code",
            "in": "query"
          },
          {
            "type": "string",
            "name": "response_type",
            "in": "query"
          },
          {
            "type": "string",
            "name": "redirect_uri",
            "in": "query"
          },
          {
            "type": "string",
            "name": "state",
            "in": "query"
          },
          {
            "type": "string",
            "name": "client_id",
            "in": "query"
          },
          {
            "type": "string",
            "name": "refresh_token",
            "in": "query"
          },
          {
            "type": "string",
            "name": "scope",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/AccessToken"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "AccessToken": {
      "type": "object",
      "properties": {
        "access_token": {
          "type": "string"
        },
        "expires_in": {
          "type": "integer",
          "format": "int64"
        },
        "refresh_token": {
          "type": "string"
        },
        "scope": {
          "type": "string"
        },
        "token_type": {
          "type": "string"
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
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Oauth API",
    "contact": {
      "name": "mars"
    },
    "version": "v1"
  },
  "basePath": "/api/v1/oauth",
  "paths": {
    "/me": {
      "get": {
        "operationId": "Me",
        "parameters": [
          {
            "type": "string",
            "name": "access_token",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/token": {
      "post": {
        "security": [
          {
            "Basic": []
          }
        ],
        "operationId": "Token",
        "parameters": [
          {
            "type": "string",
            "name": "grant_type",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "code",
            "in": "query"
          },
          {
            "type": "string",
            "name": "response_type",
            "in": "query"
          },
          {
            "type": "string",
            "name": "redirect_uri",
            "in": "query"
          },
          {
            "type": "string",
            "name": "state",
            "in": "query"
          },
          {
            "type": "string",
            "name": "client_id",
            "in": "query"
          },
          {
            "type": "string",
            "name": "refresh_token",
            "in": "query"
          },
          {
            "type": "string",
            "name": "scope",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/AccessToken"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "AccessToken": {
      "type": "object",
      "properties": {
        "access_token": {
          "type": "string"
        },
        "expires_in": {
          "type": "integer",
          "format": "int64"
        },
        "refresh_token": {
          "type": "string"
        },
        "scope": {
          "type": "string"
        },
        "token_type": {
          "type": "string"
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
