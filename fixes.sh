#!/bin/bash

# Fixes vendoring issues: https://github.com/moby/moby/issues/29362
rm -rf $GOPATH/src/github.com/docker/docker/vendor
rm -rf $GOPATH/src/github.com/docker/cli/vendor
rm -rf $GOPATH/src/github.com/docker/distribution/vendor