#!/usr/bin/env bash

PLATFORM=$1
ARCH=$2
KOMPOSE_VERSION=$3

if [ "${PLATFORM}" == 'windows' ]; then
  wget -O "cmd/kompose.exe" "https://github.com/kubernetes/kompose/releases/download/${KOMPOSE_VERSION}/kompose-windows-amd64.exe"
  chmod +x "cmd/kompose.exe"
else
  wget -O "cmd/kompose" "https://github.com/kubernetes/kompose/releases/download/${KOMPOSE_VERSION}/kompose-${PLATFORM}-${ARCH}"
  chmod +x "cmd/kompose"
fi

exit 0
