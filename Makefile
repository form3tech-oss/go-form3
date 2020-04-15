.DEFAULT_GOAL := generate-client

swagger_codegen_version := "v0.19.0"

platform := $(shell uname)
ifeq (${platform},Darwin)
swagger_binary := "swagger_darwin_amd64"
else
swagger_binary := "swagger_linux_amd64"
endif

install-swagger:
	@curl -o /usr/local/bin/swagger -L'#' https://github.com/go-swagger/go-swagger/releases/download/${swagger_codegen_version}/${swagger_binary} && chmod +x /usr/local/bin/swagger

download-swagger:
	@mkdir -p swagger
	curl -s https://api-docs.form3.tech/assets/form3-swagger.yaml -o swagger/form3-swagger-raw.yaml
	./scripts/extract_paths.py swagger/form3-swagger-raw.yaml swagger/paths.yaml


modify-swagger-file: download-swagger
	# Add an operation name (operationId) to each endpoint
	yq w swagger/form3-swagger-raw.yaml -s operationNames.yaml > swagger/form3-swagger-updated.yaml
	# Delete the ReportDetailsResponse links property (and its definition)
	yq d -i swagger/form3-swagger-updated.yaml definitions.ReportLinks
	yq d -i swagger/form3-swagger-updated.yaml definitions.ReportDetailsResponse.properties.links
	# Links.properties.Self appears to be a duplicate of Links.properties.self
	yq d -i swagger/form3-swagger-updated.yaml definitions.Links.properties.Self
	# Delete anything to do with /participants/[open_banking_id]/open-banking/v3.1/accounts/name-verification
	yq d -i swagger/form3-swagger-updated.yaml paths./participants/*
	yq d -i swagger/form3-swagger-updated.yaml responses.OBNameVerification1
	yq d -i swagger/form3-swagger-updated.yaml parameters.OBNameVerificationRequest1
	yq d -i swagger/form3-swagger-updated.yaml parameters.Authorization
	yq d -i swagger/form3-swagger-updated.yaml definitions.Meta

generate-client: modify-swagger-file
	@rm -rf pkg/generated
	@mkdir pkg/generated
	swagger generate client -f swagger/form3-swagger-updated.yaml -t pkg/generated/ -T templates -C templates/layout.yaml


GOFMT_FILES?=$$(find ./ -name '*.go' | grep -v vendor)

goimports:
	goimports -w $(GOFMT_FILES)

test:
	go test -v -race -cover ./pkg/form3
