codegen:
	@bash code-generator/code-generator.sh

deploy-local:
	@kubectl apply -f deploy/local/namespaces

sbc-service-image:
	@docker build -f apps/sbc-service/deploy/local/sbc-service/Dockerfile -t sbc-service .

sbc-service-k8-local:
	@kubectl apply -f apps/sbc-service/deploy/local/sbc-service

sbc-service-deploy-local: sbc-service-image sbc-service-k8-local

sip-server-image:
	@docker build -f apps/sip-server/deploy/local/sip-server/Dockerfile -t sip-server .

sip-server-k8-local:
	@kubectl apply -f apps/sip-server/deploy/local/sip-server

sip-server-deploy-local: sip-server-image sip-server-k8-local

up-local: deploy-local sip-server-deploy-local sbc-service-deploy-local