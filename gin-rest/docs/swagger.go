package docs

import "github.com/swaggo/swag"

// @title My API
// @version 1.0
// @description This is a sample API with Swagger documentation.

// @host localhost:8080
// @BasePath /api/v1
func init() {
	swag.Register(swag.Name, &APIInfo{})
}

// APIInfo represents the API information.
type APIInfo struct {
}
