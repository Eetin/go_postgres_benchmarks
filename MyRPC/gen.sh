#!/bin/sh
protoc MyRPC.proto --gofast_out=plugins=grpc:./
