TAG := lownie/services:xatosiz-frontend-$(shell date +%Y%m%d%H%M%S)

docker-push:
	@docker build --tag $(TAG)  .
	@docker push $(TAG)
	@echo Service pushed with tag: $(TAG)