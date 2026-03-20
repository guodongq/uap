package auth

// Anonymous
const (
	SchemeIDAnonymous = "monkey.api#noAuth"
)

// HTTP auth schemes
const (
	SchemeIDHTTPBasic  = "monkey.api#httpBasicAuth"
	SchemeIDHTTPDigest = "monkey.api#httpDigestAuth"
	SchemeIDHTTPBearer = "monkey.api#httpBearerAuth"
	SchemeIDHTTPAPIKey = "monkey.api#httpApiKeyAuth"
)
