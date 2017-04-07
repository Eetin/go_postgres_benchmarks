#!/bin/sh
protoc CoreRPC.proto --gofast_out=plugins=grpc:./
