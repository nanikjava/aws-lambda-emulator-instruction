#!/usr/bin/env bash

set -o nounset
set -o errexit
set -o pipefail
# -v  /home/nanik/GolandProjects/aws-lambda-runtime-interface-emulator/cmd/aws-lambda-rie:/aws-lambda \
#     -v /home/nanik/Downloads:/aws-lambda \

docker run \
     -v /home/nanik/GolandProjects/aws-lambda-runtime-interface-emulator/bin:/aws-lambda \
    --publish 9000:8080 \
    --entrypoint /aws-lambda/aws-lambda-rie \
    --env HEADER_KEY \
    --env HEADER_VALUE_PARAMETER \
    --env PRINCIPAL_ID \
    golambda:latest  /main