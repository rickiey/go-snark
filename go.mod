module go-snark

go 1.16

replace github.com/filecoin-project/filecoin-ffi => ./extern/filecoin-ffi

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/filecoin-project/filecoin-ffi v0.0.0-00010101000000-000000000000
	github.com/filecoin-project/go-address v0.0.6
	github.com/filecoin-project/go-state-types v0.1.8
	github.com/gin-gonic/gin v1.7.2
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529
	google.golang.org/genproto v0.0.0-20210715145939-324b959e9c22 // indirect
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1
)
