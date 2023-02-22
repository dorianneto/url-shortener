#!/bin/bash

set +xe

gcloud emulators firestore start --host-port="0.0.0.0:${PORT}"
