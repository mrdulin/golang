#!/usr/bin/env bash

gcloud functions deploy TestGolangLog --runtime go111 --trigger-http --memory=128 &
gcloud functions deploy TestGolangLogWithNewlineSymbol --runtime go111 --trigger-http --memory=128 &
gcloud functions deploy TestGolangLogUsingLog --runtime go111 --trigger-http --memory=128 &
gcloud functions deploy TestGolangLogUsingCloudLogging --runtime go111 --trigger-http --memory=128 