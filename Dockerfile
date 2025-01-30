FROM golang:latest
RUN mkdir /build
WORKDIR /build
RUN cd /build && rm -rf rest-api-demo
RUN cd /build && git clone https://github.com/Damon42Dev/rest-api-demo.git
RUN cd /build/rest-api-demo && go build
EXPOSE 8080
ENTRYPOINT [ "/build/rest-api-demo/rest-api-demo" ]