.PHONY: compile
PROTOC_GEN_GO := $(GOPATH)/bin/protoc-gen-go

# If $GOPATH/bin/protoc-gen-go does not exist, we'll run this command to install
# it.
$(PROTOC_GEN_GO):
	go get -u github.com/golang/protobuf/protoc-gen-go

rssService.pb.go: ./service/rss/grpc/rssService.proto | $(PROTOC_GEN_GO)
	cd ./service/rss/grpc && \
	protoc --go_out=. ./rssService.proto

# This is a "phony" target - an alias for the above command, so "make compile"
# still works.
compile: rssService.pb.go
