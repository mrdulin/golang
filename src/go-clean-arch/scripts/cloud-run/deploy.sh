#!/usr/bin/env bash

projectId=$(gcloud config get-value project 2> /dev/null)
app=sls-fns-go
gcloud beta run deploy --allow-unauthenticated --image=gcr.io/${projectId}/sls-fns-go
