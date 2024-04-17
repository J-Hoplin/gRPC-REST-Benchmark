#!/bin/bash
k6 run --out influxdb=http://localhost:8086/ping ping.k6.js