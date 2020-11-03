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
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "produces": [
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "From the todo list tutorial on goswagger.io",
    "title": "A Todo list application",
    "version": "1.0.0"
  },
  "paths": {
    "/": {
      "get": {
        "tags": [
          "todos"
        ],
        "operationId": "readTodos",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Список дел",
            "schema": {
              "$ref": "#/definitions/itemlist"
            }
          },
          "default": {
            "description": "Ошибка",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "todos"
        ],
        "operationId": "createOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/item"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Создано",
            "schema": {
              "$ref": "#/definitions/item"
            }
          },
          "default": {
            "description": "Ошибка",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/{id}": {
      "put": {
        "tags": [
          "todos"
        ],
        "operationId": "updateOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/item"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Ок",
            "schema": {
              "$ref": "#/definitions/item"
            }
          },
          "default": {
            "description": "Ошибка",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "todos"
        ],
        "operationId": "deleteOne",
        "responses": {
          "204": {
            "description": "Удалено"
          },
          "default": {
            "description": "Ошибка",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "item": {
      "type": "object",
      "required": [
        "description"
      ],
      "properties": {
        "completed": {
          "type": "boolean",
          "default": false
        },
        "description": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        }
      }
    },
    "itemlist": {
      "type": "object",
      "required": [
        "items",
        "total"
      ],
      "properties": {
        "items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/item"
          }
        },
        "total": {
          "type": "integer",
          "format": "int64"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "produces": [
    "application/io.goswagger.examples.todo-list.v1+json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "From the todo list tutorial on goswagger.io",
    "title": "A Todo list application",
    "version": "1.0.0"
  },
  "paths": {
    "/": {
      "get": {
        "tags": [
          "todos"
        ],
        "operationId": "readTodos",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "since",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Список дел",
            "schema": {
              "$ref": "#/definitions/itemlist"
            }
          },
          "default": {
            "description": "Ошибка",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "todos"
        ],
        "operationId": "createOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/item"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Создано",
            "schema": {
              "$ref": "#/definitions/item"
            }
          },
          "default": {
            "description": "Ошибка",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/{id}": {
      "put": {
        "tags": [
          "todos"
        ],
        "operationId": "updateOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/item"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Ок",
            "schema": {
              "$ref": "#/definitions/item"
            }
          },
          "default": {
            "description": "Ошибка",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "todos"
        ],
        "operationId": "deleteOne",
        "responses": {
          "204": {
            "description": "Удалено"
          },
          "default": {
            "description": "Ошибка",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "item": {
      "type": "object",
      "required": [
        "description"
      ],
      "properties": {
        "completed": {
          "type": "boolean",
          "default": false
        },
        "description": {
          "type": "string",
          "minLength": 1
        },
        "id": {
          "type": "integer",
          "format": "int64",
          "readOnly": true
        }
      }
    },
    "itemlist": {
      "type": "object",
      "required": [
        "items",
        "total"
      ],
      "properties": {
        "items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/item"
          }
        },
        "total": {
          "type": "integer",
          "format": "int64"
        }
      }
    }
  }
}`))
}
