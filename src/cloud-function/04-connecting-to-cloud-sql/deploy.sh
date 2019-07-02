#!/bin/bash

function=PostgresDemo
gcloud functions deploy ${function} --runtime go111 --trigger-http --memory=128 --env-vars-file .env.yaml

