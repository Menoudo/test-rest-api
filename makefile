build:
	docker build -t vanapagan/test-rest-api:latest .
run:
	docker run -d --rm --name vanapagan/test-rest-api -p 8080:8080 vanapagan/test-rest-api:latest
clean:
	docker stop vanapagan/test-rest-api || true
	docker rm vanapagan/test-rest-api || true
	docker rmi vanapagan/test-rest-api:latest
