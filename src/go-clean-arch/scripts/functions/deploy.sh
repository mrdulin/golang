#!/usr/bin/env bash

source=./infrastructure/functions

export GO111MODULE=on

${source}/getAdPerformanceReport/deploy.sh
