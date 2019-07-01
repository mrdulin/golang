#!/bin/bash

projectId=$(gcloud config get-value project 2> /dev/null)
app=connecting-to-cloud-sql
gcloud beta run deploy --image=gcr.io/${projectId}/${app}

