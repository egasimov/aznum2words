package api

// This is an example of how to add a prefix to the name of the generated Client struct
// See https://github.com/deepmap/oapi-codegen/issues/785 for why this might be necessary

//go:generate oapi-codegen -config models.cfg.yaml ./open-api-spec.yaml

//go:generate oapi-codegen -config converter-server.cfg.yaml ./open-api-spec.yaml

//go:generate oapi-codegen -config health-server.cfg.yaml ./open-api-spec.yaml
