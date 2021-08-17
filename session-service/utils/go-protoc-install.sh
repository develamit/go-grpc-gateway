#!/bin/bash
#===============================================
# go install
# Ref: https://golang.org/doc/install
#===============================================
GO_VERSION="1.16.7" # go version
ARCH="amd64" # go archicture
mkdir -p ~/Downloads
cd ~/Downloads
curl -O -L https://golang.org/dl/go${GO_VERSION}.linux-${ARCH}.tar.gz
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go${GO_VERSION}.linux-${ARCH}.tar.gz


#===============================================
# protoc install
# Ref: https://github.com/protocolbuffers/protobuf/releases/tag/v3.17.3
#      https://grpc.io/docs/protoc-installation/
#===============================================
PROTO_VERSION="3.17.3"
PB_REL="https://github.com/protocolbuffers/protobuf/releases"
mkdir ~/Downloads
cd ~/Downloads
curl -LO $PB_REL/download/v${PROTO_VERSION}/protoc-${PROTO_VERSION}-linux-x86_64.zip
mkdir $HOME/.local
sudo apt-get install -y unzip
unzip protoc-${PROTO_VERSION}-linux-x86_64.zip -d $HOME/.local

#===============================================
# Ref: https://grpc.io/docs/languages/go/quickstart/
#      https://github.com/grpc-ecosystem/grpc-gateway
# Go plugins for the protocol compiler:
# protoc-gen-go install
# export PATH="$PATH:$(go env GOPATH)/bin"
#===============================================
echo "export PATH=$PATH:/usr/local/go/bin:$HOME/.local/bin:$HOME/go/bin" >> ~/.bash_profile
source ~/.bash_profile
go install \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

#==============================================
# clone googleapis
# this will be needed for gRPC-gateway for REST over gRPC calls
#==============================================
#git clone https://github.com/grpc-ecosystem/grpc-gateway.git $(go env GOPATH)/test
git clone https://github.com/googleapis/googleapis.git $(go env GOPATH)/googleapis

#===============================================
# checks
#===============================================
go version
protoc --version

