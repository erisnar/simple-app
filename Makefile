all: build deploy forward

build:
	@eval $$(minikube docker-env) ;\
	docker build -t erik/go-web-app go/
deploy:
	kubectl apply -f k8s/redis-master-deployment.yaml
	kubectl apply -f k8s/redis-master-service.yaml
	kubectl apply -f k8s/webserver-deployment.yaml
	kubectl apply -f k8s/webserver-service.yaml
forward:
	kubectl port-forward service/webserver 8080:webserver