#!/usr/bin/env bash

PLATFORM=$1
ARCH=$2
KUBECTL_VERSION=$3

if [ "${PLATFORM}" == 'windows' ]; then
  wget -O "cmd/kubectl.exe" "https://storage.googleapis.com/kubernetes-release/release/${KUBECTL_VERSION}/bin/windows/amd64/kubectl.exe"
  chmod +x "cmd/kubectl.exe"
else
  wget -O "cmd/kubectl" "https://storage.googleapis.com/kubernetes-release/release/${KUBECTL_VERSION}/bin/${PLATFORM}/${ARCH}/kubectl"
  chmod +x "cmd/kubectl"
fi

exit 0
