#!/bin/bash
k6 run --out influxdb=http://localhost:8086/server-call server-stream.k6.js