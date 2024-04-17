#!/bin/sh

# Wait until influx bootstraped and generate database used in test. It'll stop at least one failed
# If your system's influx container bootstrap is too slow, modify it longer than 5s

sleep 2

set -e

influx -execute 'CREATE DATABASE "ping"'
influx -execute 'CREATE DATABASE "rest-call"'
influx -execute 'CREATE DATABASE "unary-call"'
influx -execute 'CREATE DATABASE "client-call"'
influx -execute 'CREATE DATABASE "server-call"'
influx -execute 'CREATE DATABASE "bi-call"'
