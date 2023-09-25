.PHONY: .generate-swagger
.generate-swagger:
	docker run --rm \
      -v ${PWD}:/local openapitools/openapi-generator-cli generate \
      -i /local/api/swagger.yml \
      -g go \
      -o /local/openapi/ \
      --additional-properties=outputAsLibrary=true,sourceFolder=.

.PHONY: goimports
goimports:
	goimports -w .

.PHONY: generate-swagger
generate-swagger: .generate-swagger goimports