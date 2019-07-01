#!/usr/bin/env bash

projectId=$(gcloud config get-value project 2> /dev/null)
app=includedeps
gcloud beta run deploy --allow-unauthenticated --image=gcr.io/${projectId}/includedeps
