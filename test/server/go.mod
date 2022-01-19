module flow/test/server

go 1.16

require (
	flow/test/message v0.0.0
	flow/test/proto v0.0.0
	google.golang.org/grpc v1.43.0
)

replace (
	flow/test/message => ../message
	flow/test/proto => ../proto
)
