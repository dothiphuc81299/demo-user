package zookeeper

import (
	"demo-user/config"
	"os"

	"github.com/samuel/go-zookeeper/zk"

	"fmt"
	"time"
)

var conn *zk.Conn

// Connect ...
func Connect() {
	var (
		uri     = os.Getenv("ZOOKEEPER_URI")
		envVars = config.GetEnv()
	)
	conn, _, err := zk.Connect([]string{uri}, time.Second*30)
	if err != nil {
		fmt.Println("ZookeeperURI:", uri)
		panic(err)
	}
	fmt.Println("Zookeeper Connected to", uri)

	// Get env key
	// App port
	appUserPort, _, _ := conn.Get("/app/port/user")
	envVars.AppUserPort = string(appUserPort)

	// Database
	databaseURI, _, _ := conn.Get("/database/uri")
	envVars.Database.URI = string(databaseURI)
	databaseUserName, _, _ := conn.Get("/database/name/user")
	envVars.Database.UserName = string(databaseUserName)
	databaseTestName, _, _ := conn.Get("/database/test/name")
	envVars.Database.TestName = string(databaseTestName)

	// GRPC
	grpcURI, _, _ := conn.Get("/grpc/uri")
	envVars.GRPCUri = string(grpcURI)
}
