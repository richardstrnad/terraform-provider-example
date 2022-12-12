#!/bin/bash

VERSION=$(cat version)
echo "building terraform-provider-filr_${VERSION}"
go build -o terraform-provider-filr_${VERSION}
