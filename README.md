#  test-rest-api

A simple golang TEST ReST API that utilizes gorilla/mux and returns the host or
container name and IP address to aid in load balancer verification and debugging.

This TEST Rest API is also useful for Docker/Kubernetes teaching demonstrations
allowing the user to demonstrate that each instance of the container has a
separate and distinct host name and IP address.

This API uses gorilla/mux, a powerful URL router and dispatcher for golang.
See: http://www.gorillatoolkit.org/pkg/mux for more information.

This container is built using the official golang Alpine container as a base.

## Build the Application

To run the application, first build the application with the following command:

`go build`

Please insure that you have gorilla/mux installed in your golang environment.
Refer to the golang documentation for more information on installing golang
modules.

## Run the Application

Then execute the build with the following command:

`./test-rest-api`

## Test the Application

To test out the container listening on port 8080, use these URLs:
```sh
curl http://localhost:8080/about
curl http://localhost:8080/test
curl http://localhost:8080/hostname
curl http://localhost:8080/hello/Vanapagan
```

The random value for "Code" will change with each execution to help you insure
you are not seeing cached results in your browser.

## Build the Docker Image

To build the image with the application run:

`docker build -f Dockerfile . -t testrestapi:latest`

## Run the Docker image

To run the image, execute the following commands to create the containers:

`docker run --rm -d --name testrest8080 -p 8080:8080 testrestapi:latest`

`docker run --rm -d --name testrest8081 -p 8081:8080 testrestapi:latest`

## Test the Docker containers

To test out the container listening on port 8080, use these URLs:
```sh
curl http://localhost:8080/about
curl http://localhost:8080/test
curl http://localhost:8080/hostname
curl http://localhost:8080/json?name=test
curl http://localhost:8080/hello/Vanapagan
```

To test out the container listening on port 8081, use these URLs:
```sh
curl http://localhost:8081/about
curl http://localhost:8081/test
curl http://localhost:8081/hostname
curl http://localhost:8081/hello/Vanapagan
```

test_br fix