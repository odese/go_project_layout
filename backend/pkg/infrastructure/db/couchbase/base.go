package couchbase

import (
	// "strings"
	// "time"

	// ".../backend/pkg/infrastructure/config"
	// ".../backend/pkg/utils"
	"github.com/couchbase/gocb/v2"

	// log ".../backend/pkg/infrastructure/logger"
)

// Cluster Instance, represents a connection to Couchbase cluster.
var cluster *gocb.Cluster

// Bucket instance
var bucket *gocb.Bucket

// Scope instance
var scope *gocb.Scope

// collections, map of collection instances.
var collections map[string]*gocb.Collection

// // Init, initializes couchbase connections.
// func Init() {
// 	initClusterInstance()
// 	initBucketInstance()
// 	initScopeInstance()
// 	initCollectionInstances()
// }

// // initClusterInstance, connects to couchbase server and initiate cluster instance.
// func initClusterInstance() {
// 	var err error

// 	server := config.Call().GetString("couchbase.server")
// 	port := config.Call().GetString("couchbase.port")
// 	username := config.Call().GetString("couchbase.username")
// 	password := config.Call().GetString("couchbase.password")

// 	cluster, err = ConnectToCluster(server, username, password)
// 	if err != nil {
// 		log.Fatal().Err(err).Msg("Error on connecting to cluster.")
// 	}

// 	err = cluster.WaitUntilReady(5*time.Second, nil)
// 	if err != nil {
// 		log.Fatal().Err(err).Send()
// 	}

// 	log.Info().Str("Couchbase", "http://"+server+port).Send()
// }

// // initBucketInstance, initiates & assigns project's Bucket instance.
// func initBucketInstance() {
// 	var err error
// 	bucket, err = ConnectToBucket(utils.CouchbaseBucketName)
// 	if err != nil {
// 		log.Fatal().Err(err).Str("Bucket Name", utils.CouchbaseBucketName).Send()
// 	}
// 	log.Info().Str("Bucket Name", bucket.Name()).Send()
// }

// // initScopeInstance, initiates & assigns project's Scope instance according to WORK_ENV.
// func initScopeInstance() {
// 	scope = ConnectToScope(bucket, strings.ReplaceAll(utils.WorkEnvironment, "LOCAL_", ""))
// 	bucket.Collections()
// 	log.Info().Str("Scope Initiated", scope.Name()).Send()
// }

// // initCollectionInstances, initiates & assigns project's Collection instances into a map.
// func initCollectionInstances() {
// 	collections = make(map[string]*gocb.Collection)
// 	for i := 0; i < len(utils.CouchbaseCollectionNames); i++ {
// 		collectionName := utils.CouchbaseCollectionNames[i]
// 		collections[collectionName] = ConnectToCollection(scope, collectionName)
// 		log.Info().Str("Collection Loaded", collections[collectionName].Name()).Send()
// 	}
// }
