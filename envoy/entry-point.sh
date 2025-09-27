#!/bin/sh
set -e

exec envoy -c /etc/envoy/envoy.yaml

