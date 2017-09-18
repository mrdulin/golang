#!/usr/bin/env bash

source=./infra/functions

export GO111MODULE=on

${source}/getAdPerformanceReport/deploy.sh &
${source}/getCampaignReport/deploy.sh