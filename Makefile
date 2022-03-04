
all:
	rm -rf *pb
	buf generate
	mv gen/github.com/microbun/protobuf/* .
	rm -rf gen