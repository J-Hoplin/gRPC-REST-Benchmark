#!/bin/bash
k6 run --out influxdb=http://localhost:8086/rest-call rest.k6.js