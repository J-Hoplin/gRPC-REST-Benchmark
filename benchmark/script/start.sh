#!/bin/bash
k6 run --out influxdb=http://localhost:8086/rest-call ping.k6.js