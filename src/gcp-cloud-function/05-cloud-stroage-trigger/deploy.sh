#!/bin/bash

function=CloudStorageTriggerTest
bucket=ez2on

go mod vendor && gcloud functions deploy ${function} --runtime go111 --trigger-resource ${bucket} --trigger-event google.storage.object.finalize --set-env-vars bucket=ez2on
