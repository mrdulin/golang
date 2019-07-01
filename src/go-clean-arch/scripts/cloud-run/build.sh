#!/usr/bin/env bash

projectId=$(gcloud config get-value project 2> /dev/null)
gcloud builds submit --tag gcr.io/${projectId}/sls-fns-go
