module go-snark

go 1.16

replace github.com/filecoin-project/filecoin-ffi => ./extern/filecoin-ffi

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/filecoin-project/filecoin-ffi v0.30.4-0.20200910194244-f640612a1a1f
	github.com/filecoin-project/go-address v1.1.0
	github.com/filecoin-project/go-state-types v0.11.2-0.20230712101859-8f37624fa540
	github.com/gin-gonic/gin v1.7.2
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529
	google.golang.org/genproto v0.0.0-20210715145939-324b959e9c22 // indirect
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.27.1
)
