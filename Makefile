swagger_codegen_version := "v0.19.0"

ifeq (${platform},Darwin)
swagger_binary := "swagger_darwin_amd64"
else
swagger_binary := "swagger_linux_amd64"
endif

install-swagger:
	@sudo curl -o /usr/local/bin/swagger -L'#' https://github.com/go-swagger/go-swagger/releases/download/${swagger_codegen_version}/${swagger_binary} && chmod +x /usr/local/bin/swagger; \

download-swagger:
	@mkdir swagger || true
	curl -s http://api-docs.form3.tech/assets/form3-swagger.yaml -o swagger/form3-swagger-raw.yaml
	yq r swagger/form3-swagger-raw.yaml paths | grep -v "^   .*" | tr '\n' 'X' | sed -e 's/\(\/[^X]*\):X *\([^:]*\):/paths.\1.\2.OperationId: TODO/g' | tr 'X' '\n' > swagger/paths.txt


modify-swagger-file: download-swagger
	yq w swagger/form3-swagger-raw.yaml -s operationNames.txt > swagger/form3-swagger-updated.yaml
	yq d -i swagger/form3-swagger-updated.yaml definitions.ReportLinks
	yq d -i swagger/form3-swagger-updated.yaml definitions.ReportDetailsResponse.properties.links

generate-client: modify-swagger-file
	@rm -r pkg/generated || true
	@mkdir pkg/generated
	swagger generate client -f swagger/form3-swagger-updated.yaml -t pkg/generated/ -T templates -C templates/layout.yaml

GOFMT_FILES?=$$(find ./ -name '*.go' | grep -v vendor)

goimports:
	goimports -w $(GOFMT_FILES)

test:
	go test ./tests