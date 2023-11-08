codegen:
	@bash code-generator/code-generator.sh

image-local:
	@docker build -f deploy/local/coba/Dockerfile -t coba .

deploy-local:
	@kubectl apply -f deploy/local/namespaces
	@kubectl apply -f deploy/local/coba

up-local: image-local deploy-local