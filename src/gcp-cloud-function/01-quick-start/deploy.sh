#!/bin/bash

function=HelloGet

gcloud functions deploy $function --runtime go111 --trigger-http --memory=128