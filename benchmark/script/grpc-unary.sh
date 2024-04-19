#!/bin/bash
k6 run --out influxdb=http://localhost:8086/unary-call -e ENDPOINT="http://localhost:8080/grpc/unary?from=1&to=100" script.k6.js