#!/bin/bash
# set -x

apk add --no-cache \
        python3 \
        py3-pip \
    && pip3 install --upgrade pip \
    && pip3 install --no-cache-dir \
        awscli \
    && rm -rf /var/cache/apk/*

k8s_suffix=$CI_KUBE_ENVIRONMENT
namespace_suffix=$CI_KUBE_NAMESPACE

k8s_env="CI_KUBE_ENVIRONMENT"
k8s_server="CI_KUBE_SERVER"
k8s_user="CI_KUBE_USER_K8S"
k8s_token="CI_KUBE_TOKEN"

version1=$(shuf -i 0-100 -n 1)
version2=$(shuf -i 0-100 -n 1)
version3=$(shuf -i 0-100 -n 1)

eval k8s_server=\$$k8s_server
eval k8s_user=\$$k8s_user
eval k8s_token=\$$k8s_token

echo $k8s_server
echo $k8s_user
echo $k8s_suffix

deployment="deploy.yaml"

docker info
docker login -u gitlab-ci-token -p $CI_BUILD_TOKEN github.com/diploma:latest

docker run \
    -v $(pwd)/$deployment:/kube/deployment.yaml \
    github.com/diploma/devops:latest \
    $k8s_server $namespace_suffix $k8s_user $k8s_token \
    "apply -f /kube/deployment.yaml"
