#! /bin/bash

docker run  \
  --rm \
  -w /root/go/src/$GO_PATH \
  --mount type=bind,source="$(pwd)",target=/root/go/src/$GO_PATH \
  eignhpants/image-builder
