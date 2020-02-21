.DEFAULT_GOAL := generate-client

swagger_codegen_version := "v0.21.0"

platform := $(shell uname)
ifeq (${platform},Darwin)
swagger_binary := "swagger_darwin_amd64"
else
swagger_binary := "swagger_linux_amd64"
endif

install-swagger:
	@curl -o /usr/local/bin/swagger -L'#' https://github.com/go-swagger/go-swagger/releases/download/${swagger_codegen_version}/${swagger_binary} && chmod +x /usr/local/bin/swagger

install-yq:
	GO111MODULE=on go get github.com/mikefarah/yq/v2

download-swagger:
	@mkdir swagger || true
	curl -s https://api-docs.form3.tech/assets/form3-swagger.yaml -o swagger/form3-swagger-raw.yaml
	yq r swagger/form3-swagger-raw.yaml paths | grep -v "^   .*" | ./scripts/determine-endpoints.py > swagger/endpoints.txt

modify-swagger-file: download-swagger
	yq w swagger/form3-swagger-raw.yaml -s operation-names.txt > swagger/form3-swagger-updated.yaml
	yq d -i swagger/form3-swagger-updated.yaml definitions.ReportLinks
	yq d -i swagger/form3-swagger-updated.yaml definitions.ReportDetailsResponse.properties.links

generate-client: modify-swagger-file
	@rm -r pkg/generated || true
	@mkdir pkg/generated
	swagger generate client -f swagger/form3-swagger-updated.yaml -t pkg/generated/ -T templates -C templates/layout.yaml

determine-missing-operations:
	@bash -c "comm -13 <(cut -d: -f1 operation-names.txt | sed 's/$$/: TODO/' | sort) <(sort swagger/endpoints.txt) >> operation-names.txt"

GOFMT_FILES?=$$(find ./ -name '*.go' | grep -v vendor)

goimports:
	goimports -w $(GOFMT_FILES)

test:
	go test -v -race -cover ./pkg/form3