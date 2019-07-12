#!/bin/bash

function=CloudStorageTriggerTest
bucket=ez2on

gcloud functions deploy ${function} --runtime go111 --trigger-resource ${bucket} --trigger-event google.storage.object.finalize
