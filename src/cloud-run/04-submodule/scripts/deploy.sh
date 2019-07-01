#!/usr/bin/env bash

projectId=$(gcloud config get-value project 2> /dev/null)
app=submodule
gcloud beta run deploy --allow-unauthenticated --image=gcr.io/${projectId}/submodule
