FROM golang:latest
RUN mkdir /build
WORKDIR /build
RUN cd /build && rm -rf REST-API-DEMO
RUN cd /build && git clone https://github.com/Damon42Dev/rest-api-demo.git
RUN cd /build/REST-API-DEMO && go build
EXPOSE 8080
ENTRYPOINT [ "/build/REST-API-DEMO/REST-API-DEMO" ]