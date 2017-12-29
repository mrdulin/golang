#!/usr/bin/env bash

cd "$(dirname "$0")"
export GO111MODULE=on
go mod vendor && gcloud functions deploy GetAdPerformanceReport --runtime=go111 --trigger-topic get-ad-performance-report --env-vars-file .env.yaml