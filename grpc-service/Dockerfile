FROM golang:1.21-bullseye

WORKDIR /app
COPY . .

# Install protoc & Download go dependencies
RUN apt update\
&& apt install -y protobuf-compiler\
&& go install google.golang.org/protobuf/cmd/protoc-gen-go@latest\
&& go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest\
&& go mod download\
&& make service

EXPOSE 8082

CMD [ "-c", "./bin/server" ]
ENTRYPOINT [ "/bin/bash" ]
