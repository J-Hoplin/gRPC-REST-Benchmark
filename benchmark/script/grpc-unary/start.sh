#!/bin/bash
k6 run --out influxdb=http://localhost:8086/unary-call unary.k6.js