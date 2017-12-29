#!/bin/bash

function=IncludedependenciesTest

gcloud functions deploy $function --runtime go111 --trigger-http --memory=128