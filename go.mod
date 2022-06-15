module go-snark

go 1.16

replace github.com/filecoin-project/filecoin-ffi => ./extern/filecoin-ffi

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/filecoin-project/filecoin-ffi v0.0.0-00010101000000-000000000000
	github.com/filecoin-project/go-address v0.0.5
	github.com/filecoin-project/go-state-types v0.1.3
	github.com/gin-gonic/gin v1.7.2
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	google.golang.org/genproto v0.0.0-20210715145939-324b959e9c22 // indirect
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1
)
