module github.com/ops-cn/go-devops

go 1.14

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/bwmarrin/snowflake v0.3.0
	github.com/casbin/casbin/v2 v2.7.2
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fatih/camelcase v1.0.0 // indirect
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis v6.15.8+incompatible
	github.com/go-redis/redis_rate v6.5.0+incompatible
	github.com/golang/protobuf v1.4.2
	github.com/google/gops v0.3.10
	github.com/google/uuid v1.1.1
	github.com/google/wire v0.4.0
	github.com/jinzhu/copier v0.0.0-20190924061706-b57f9002281a
	github.com/jinzhu/gorm v1.9.14
	github.com/json-iterator/go v1.1.10
	github.com/koding/multiconfig v0.0.0-20171124222453-69c27309b2d7
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.0
	github.com/ops-cn/admin v0.0.0-20200703083910-07aa72edbc6e
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.6.0
	github.com/stretchr/testify v1.6.1
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.7
	github.com/tidwall/buntdb v1.1.2
	go.mongodb.org/mongo-driver v1.3.4
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1
	google.golang.org/protobuf v1.24.0
	gopkg.in/yaml.v2 v2.3.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
