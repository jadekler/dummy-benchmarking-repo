#!/bin/bash
set -ex

echo "Hello World"

env

echo $KOKORO_BLAZE_DIR
ls $KOKORO_BLAZE_DIR
ls $KOKORO_BLAZE_DIR/build_ls
ls $KOKORO_BLAZE_DIR/build_ls/blaze-bin
ls $KOKORO_BLAZE_DIR/build_ls/blaze-bin/devrel
ls $KOKORO_BLAZE_DIR/build_ls/blaze-bin/devrel/cloud
ls $KOKORO_BLAZE_DIR/build_ls/blaze-bin/devrel/cloud/client_libraries
ls $KOKORO_BLAZE_DIR/build_ls/blaze-bin/devrel/cloud/client_libraries/benchmarking
$KOKORO_BLAZE_DIR/build_ls/blaze-bin/devrel/cloud/client_libraries/benchmarking/ls

pwd
