#!/bin/bash
k6 run --out influxdb=http://localhost:8086/server-call -e ENDPOINT="http://localhost:8080/grpc/stream/server?from=1&to=100" script.k6.js