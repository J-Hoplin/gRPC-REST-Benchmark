#!/bin/bash
k6 run --out influxdb=http://localhost:8086/client-call client-stream.k6.js