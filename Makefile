.DEFAULT_GOAL := generate-client

GOFMT_FILES?=$$(find ./ -name '*.go' | grep -v vendor)

swagger_codegen_version := "v0.27.0"

platform := $(shell uname)
ifeq (${platform},Darwin)
swagger_binary := "swagger_darwin_amd64"
else
swagger_binary := "swagger_linux_amd64"
endif

PHONY: install-swagger
install-swagger:
	@curl -o /usr/local/bin/swagger -L'#' https://github.com/go-swagger/go-swagger/releases/download/${swagger_codegen_version}/${swagger_binary} && chmod +x /usr/local/bin/swagger

PHONY: download-swagger
download-swagger:
	@mkdir -p swagger
	curl -s https://api-docs.form3.tech/assets/form3-swagger.yaml -o swagger/form3-swagger-raw.yaml
	./scripts/extract_paths.py swagger/form3-swagger-raw.yaml swagger/paths.yaml

PHONY: modify-swagger-file
modify-swagger-file: download-swagger
	# Add an operation name (operationId) to each endpoint
	yq eval-all 'select(fi==0) * select(fi==1)' swagger/form3-swagger-raw.yaml operationNames.yaml > swagger/form3-swagger-updated.yaml
	# Delete the ReportDetailsResponse links property (and its definition)
	yq eval 'del(.definitions.ReportLinks)' -i swagger/form3-swagger-updated.yaml
	yq eval 'del(.definitions.ReportDetailsResponse.properties.links)' -i swagger/form3-swagger-updated.yaml
	# Links.properties.Self appears to be a duplicate of Links.properties.self
	yq eval 'del(.definitions.Links.properties.Self)' -i swagger/form3-swagger-updated.yaml
	# remove addressbook/health as codgenerator generates invalid code for it
	yq eval 'del(.paths./addressbook/health)' -i swagger/form3-swagger-updated.yaml
	# Delete anything to do with /participants/[open_banking_id]/open-banking/v3.1/accounts/name-verification
	yq eval 'del(.paths./openbanking/accounts/name-verification)' -i swagger/form3-swagger-updated.yaml
	yq eval 'del(.definitions.Meta)' -i swagger/form3-swagger-updated.yaml
	yq eval 'del(.responses.OBNameVerification1.schema.properties.Meta)' -i swagger/form3-swagger-updated.yaml

	# remove paths that are missing a `response` property
	yq eval 'del(.paths./transaction/mandates/health.get)' -i swagger/form3-swagger-updated.yaml
	yq eval 'del(.paths./transaction/payments/health.get)' -i swagger/form3-swagger-updated.yaml
	yq eval 'del(.paths./notification/subscriptions/health.get)' -i swagger/form3-swagger-updated.yaml
	yq eval 'del(.paths./transaction/directdebits/health.get)' -i swagger/form3-swagger-updated.yaml
	yq eval 'del(.paths./security/users/health.get)' -i swagger/form3-swagger-updated.yaml
	yq eval 'del(.paths./audit/entries/health.get)' -i swagger/form3-swagger-updated.yaml
	yq eval 'del(.paths./transaction/claims/health.get)' -i swagger/form3-swagger-updated.yaml
	yq eval 'del(.paths./organisation/units/health.get)' -i swagger/form3-swagger-updated.yaml


PHONY: generate-client
generate-client: modify-swagger-file
	@rm -rf pkg/generated
	@mkdir pkg/generated
	swagger generate client -f swagger/form3-swagger-updated.yaml -t pkg/generated/ -T templates -C templates/layout.yaml

PHONY: goimports
goimports:
	goimports -w $(GOFMT_FILES)

PHONY: test
test:
	go test -v -race -cover ./pkg/form3
