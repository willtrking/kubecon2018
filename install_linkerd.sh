#!/bin/bash

# This will install linkerd2 with an admission controller, which will automatically inject it into all new deployments
# At time of writing, this only works with Deployment workloads
# See https://linkerd.io/2/proxy-injection/ for details

linkerd install --tls=optional --proxy-auto-inject | kubectl apply -f -