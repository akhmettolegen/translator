// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/translation": {
            "post": {
                "description": "Translate a text",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "translation"
                ],
                "summary": "Translate",
                "operationId": "do-translate",
                "parameters": [
                    {
                        "description": "Set up translation",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.doTranslateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Translation"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/translation/history": {
            "get": {
                "description": "Show all translation history",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "translation"
                ],
                "summary": "Show history",
                "operationId": "history",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.historyResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Translation": {
            "type": "object",
            "properties": {
                "destination": {
                    "type": "string",
                    "example": "en"
                },
                "original": {
                    "type": "string",
                    "example": "текст для перевода"
                },
                "source": {
                    "type": "string",
                    "example": "auto"
                },
                "translation": {
                    "type": "string",
                    "example": "text for translation"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "v1.doTranslateRequest": {
            "type": "object",
            "required": [
                "destination",
                "original",
                "source"
            ],
            "properties": {
                "destination": {
                    "type": "string",
                    "example": "en"
                },
                "original": {
                    "type": "string",
                    "example": "текст для перевода"
                },
                "source": {
                    "type": "string",
                    "example": "auto"
                }
            }
        },
        "v1.historyResponse": {
            "type": "object",
            "properties": {
                "history": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Translation"
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/v1",
	Schemes:     []string{},
	Title:       "Translator API",
	Description: "Translation service",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
