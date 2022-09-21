module study-goland

go 1.17

require (
	github.com/360EntSecGroup-Skylar/excelize/v2 v2.4.0
	github.com/Shopify/sarama v1.36.0
	github.com/astaxie/beego v1.12.3
	github.com/bmizerany/pat v0.0.0-20170815010413-6226ea591a40
	github.com/bmizerany/pq v0.0.0-20131128184720-da2b95e392c1
	github.com/bwmarrin/snowflake v0.3.0
	github.com/go-redis/redis v6.14.2+incompatible
	github.com/go-redis/redis/v8 v8.11.5
	github.com/go-sql-driver/mysql v1.6.0
	github.com/go-zookeeper/zk v1.0.2
	github.com/google/uuid v1.2.0
	github.com/pkg/profile v1.6.0
	github.com/sony/sonyflake v1.0.0
	go.etcd.io/etcd/client/v3 v3.5.5
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.34.0
	go.opentelemetry.io/otel v1.10.0
	go.opentelemetry.io/otel/exporters/prometheus v0.31.0
	go.opentelemetry.io/otel/metric v0.32.0
	go.opentelemetry.io/otel/sdk/metric v0.31.0
	go.uber.org/ratelimit v0.2.0
	golang.org/x/net v0.0.0-20220809184613-07c6da5e1ced
	gopl.io v0.0.0-20200323155855-65c318dde95e
	gorm.io/datatypes v1.0.1
	gorm.io/driver/mysql v1.1.1
	gorm.io/gorm v1.21.9
	xorm.io/xorm v1.0.5
)

require (
	github.com/andres-erbsen/clock v0.0.0-20160526145045-9e14626cd129 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/eapache/go-resiliency v1.3.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.3 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.2 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/prometheus/client_golang v1.12.2 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.32.1 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/richardlehane/mscfb v1.0.3 // indirect
	github.com/richardlehane/msoleps v1.0.1 // indirect
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/syndtr/goleveldb v1.0.0 // indirect
	github.com/xuri/efp v0.0.0-20210322160811-ab561f5b45e3 // indirect
	go.etcd.io/etcd/api/v3 v3.5.5 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.5 // indirect
	go.opentelemetry.io/otel/sdk v1.8.0 // indirect
	go.opentelemetry.io/otel/trace v1.10.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.17.0 // indirect
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
	golang.org/x/sys v0.0.0-20220728004956-3c1f35247d10 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c // indirect
	google.golang.org/grpc v1.41.0 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	xorm.io/builder v0.3.7 // indirect
)

//github.com/coreos/bbolt v1.3.6 => go.etcd.io/bbolt v1.3.6
replace google.golang.org/grpc v1.37.1 => google.golang.org/grpc v1.26.0

//module declares its path as: go.etcd.io/bbolt
//but was required as: github.com/coreos/bbolt
