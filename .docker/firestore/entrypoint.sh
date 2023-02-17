#!/bin/bash

set +xe

gcloud emulators firestore start --host-port="127.0.0.1:${PORT}"
