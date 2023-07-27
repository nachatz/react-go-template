package api

// API consts
const (
	BaseURL        = "http://localhost:8080" // Base URL of the API
	EndpointSuffix = "/v1"                   // Endpoint suffix for version 1 of the API
	Endpoint       = BaseURL + EndpointSuffix
)

// API endpoints
const (
	EndpointAuthToken = EndpointSuffix + "/jwt"
	EndpointHello     = EndpointSuffix + "/hello"
)
