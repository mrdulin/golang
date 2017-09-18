#!/usr/bin/env bash

cd "$(dirname "$0")"
go mod vendor && gcloud functions deploy GetCampaignReport --runtime=go111 --trigger-http