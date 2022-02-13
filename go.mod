module study-goland

go 1.16

require (
	github.com/360EntSecGroup-Skylar/excelize/v2 v2.4.0
	github.com/Shopify/sarama v1.27.2
	github.com/astaxie/beego v1.12.3
	github.com/bmizerany/pat v0.0.0-20170815010413-6226ea591a40
	github.com/bmizerany/pq v0.0.0-20131128184720-da2b95e392c1
	github.com/bwmarrin/snowflake v0.3.0
	github.com/coreos/bbolt v1.3.6 // indirect
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/gin-gonic/gin v1.7.7 // indirect
	github.com/go-redis/redis v6.14.2+incompatible
	github.com/go-redis/redis/v8 v8.8.2
	github.com/go-sql-driver/mysql v1.6.0
	github.com/go-zookeeper/zk v1.0.2
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/btree v1.0.1 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/pkg/profile v1.6.0 // indirect
	github.com/soheilhy/cmux v0.1.5 // indirect
	github.com/sony/sonyflake v1.0.0
	github.com/tealeg/xlsx v1.0.5
	github.com/tmc/grpc-websocket-proxy v0.0.0-20201229170055-e5319fda7802 // indirect
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
	go.etcd.io/etcd v3.3.25+incompatible
	go.uber.org/ratelimit v0.2.0
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/net v0.0.0-20210805182204-aaa1db679c0d
	golang.org/x/time v0.0.0-20210611083556-38a9dc6acbc6 // indirect
	golang.org/x/tools v0.1.6 // indirect
	google.golang.org/grpc v1.37.1 // indirect
	gopl.io v0.0.0-20200323155855-65c318dde95e
	gorm.io/datatypes v1.0.1
	gorm.io/driver/mysql v1.1.1
	gorm.io/gorm v1.21.9
	sigs.k8s.io/yaml v1.2.0 // indirect
	xorm.io/xorm v1.0.5
)

replace (
	github.com/coreos/bbolt v1.3.6 => go.etcd.io/bbolt v1.3.6
	google.golang.org/grpc v1.37.1 => google.golang.org/grpc v1.26.0
)

//module declares its path as: go.etcd.io/bbolt
//but was required as: github.com/coreos/bbolt
