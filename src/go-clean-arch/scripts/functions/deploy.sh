#!/usr/bin/env bash

PACKAGE=serverlessfns
FUNCTION=GetAdPerformanceReport

#gcloud functions deploy ${PACKAGE} --entry-point ${FUNCTION} --runtime go111 --trigger-topic get-ad-performance-report
gcloud functions deploy ${FUNCTION} --runtime go111 --trigger-topic get-ad-performance-report