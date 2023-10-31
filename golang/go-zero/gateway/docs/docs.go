package docs

import _ "embed"

//go:embed swagger-ui.html
var SwaggerUI []byte

//go:embed openapi.yaml
var OpenAPI []byte
