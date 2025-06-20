package handlers

import (
	_ "api-gateway/docs"
	httpSwagger "github.com/ducknes/http-swagger"
	"net/http"
)

func SwaggerHandler() http.Handler {
	return httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"))
}
