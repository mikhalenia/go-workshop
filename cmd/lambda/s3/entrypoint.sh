#!/bin/sh

/usr/bin/minio server --address "0.0.0.0:${MINIO_PORT:-8080}" /data