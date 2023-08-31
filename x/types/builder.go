package types

func NewOpenAPIData(title, version string) *OpenAPI {
	return &OpenAPI{
		OpenAPI: "3.0.0",
		Info: Info{
			title, version,
		},
		Paths: make(PathMap),
	}
}
