#!/bin/bash
k6 run --out influxdb=http://localhost:8086/bi-call bi-directional-stream.k6.js