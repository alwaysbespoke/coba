codegen:
	@bash code-generator/code-generator.sh

sip-server-image:
	@docker build -f apps/sip-server/deploy/local/sip-server/Dockerfile -t coba .

sip-server-k8-local:
	@kubectl apply -f apps/sip-server/deploy/local/namespaces
	@kubectl apply -f apps/sip-server/deploy/local/sip-server

sip-server-deploy-local: sip-server-image sip-server-k8-local