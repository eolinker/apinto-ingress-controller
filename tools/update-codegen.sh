#!/usr/bin/env bash

# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")
PROJECT_ROOT="$SCRIPT_ROOT/.."

#CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${PROJECT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}
CODEGEN_PKG=${SCRIPT_ROOT}
PKG_NAME="github.com/eolinker/apinto-ingress-controller"

GENERATED_ROOT="$PROJECT_ROOT/.generated"
rm -rf "$GENERATED_ROOT"

# 生成基本的client,informer,lister文件
bash "${CODEGEN_PKG}"/generate-groups.sh "deepcopy,client,informer,lister" \
  "${PKG_NAME}"/pkg/kube/apinto/client "${PKG_NAME}"/pkg/kube/apinto/configs \
  apinto:v1 ${PKG_NAME} \
  --output-base "$GENERATED_ROOT" \
  --go-header-file "${SCRIPT_ROOT}"/boilerplate.go.txt \
  "$@"

bash "${CODEGEN_PKG}"/generate-groups.sh "deepcopy" \
  "${PKG_NAME}"/pkg/types "${PKG_NAME}"/pkg/types \
  apinto:v1 ${PKG_NAME} \
  --output-base "$GENERATED_ROOT" \
  --go-header-file "${SCRIPT_ROOT}"/boilerplate.go.txt \
  "$@"

# 生成register文件
bash "${CODEGEN_PKG}"/generate-groups.sh "register" \
  ${PKG_NAME}/pkg/kube/apinto/configs ${PKG_NAME}/pkg/kube/apinto/configs \
  apinto:v1 ${PKG_NAME} \
  --output-base "$GENERATED_ROOT" \
  --go-header-file "${SCRIPT_ROOT}"/boilerplate.go.txt \
  "$@"

cp -r "$GENERATED_ROOT/${PKG_NAME}/"** "$PROJECT_ROOT"
rm -rf "$GENERATED_ROOT"
