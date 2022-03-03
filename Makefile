
all:
	rm -rf *pb
	buf generate internal
	mv gen/github.com/microbun/protobuf/* .
	rm -rf gen