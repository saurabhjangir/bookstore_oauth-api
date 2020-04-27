package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/saurabhjangir/utils-lib-golang/errors"
	"os"
)

const (
	db_host = "DBHOST"
)

var (
	cluster *gocql.ClusterConfig
	dbhost  = os.Getenv(db_host)
)

func init() {
	cluster = gocql.NewCluster(dbhost)
	cluster.Keyspace = "oauth"
	session, err := cluster.CreateSession()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	defer session.Close()
	fmt.Println("casandra connection initialized")
}

func GetSession() (*gocql.Session, *errors.RestErr) {
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, errors.NewRestErrInteralServer(err.Error())
	}
	return session, nil
}
