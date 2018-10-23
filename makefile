build:
	docker build -t testrestapi:latest .
run:
	docker run -d --rm --name testrestapi -p 8080:8080 testrestapi:latest
clean:
	docker stop testrestapi || true
	docker rm testrestapi || true
	docker rmi testrestapi:latest
