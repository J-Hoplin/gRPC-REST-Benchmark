#!/bin/bash

# Intialization jobs

serviceInitialzie(){
    local directory=$1
    cd $directory
    go mod download # This will install dependencies from go.mod file
    make service
    go mod tidy # This will install golang additional dependencies from grpc generated codes
}

# Directories
restAPIDir="$(pwd)/user-rest-handler"
grpcServiceDir="$(pwd)/grpc-service"
restServiceDir="$(pwd)/rest-service"

initializeOrder=("$grpcServiceDir" "$restServiceDir" "$restAPIDir")
for dir in "${initializeOrder[@]}"; do
    serviceInitialzie "$dir"
done