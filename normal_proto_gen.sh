#!/usr/bin/env sh
protoc model/normal_proto/analytics.proto --go_out=.


echo "normal_proto generated"