package couchbase

import (
	"time"

	"github.com/couchbase/gocb/v2"
)

// ConnectToCluster, wrapper for Connect Cluster, creates and returns a Cluster instance by the provided options
// TODO: Customize ClusterOptions, there are plenty of options e.g. 'RetryStrategy', 'Timeouts'
func ConnectToCluster(server, username, password string) (clusterInstance *gocb.Cluster, err error) {
	clusterInstance, err = gocb.Connect(server,
		gocb.ClusterOptions{
			Authenticator: gocb.PasswordAuthenticator{
				Username: username,
				Password: password,
			},
			TimeoutsConfig: gocb.TimeoutsConfig{
				ConnectTimeout: 1 * time.Second,
			},
		},
	)
	return clusterInstance, err
}

// ConnectToBucket, wrapper for Connecting to Bucket, returns a Bucket instance.
func ConnectToBucket(name string) (bucketInstance *gocb.Bucket, err error) {
	bucketInstance = cluster.Bucket(name)
	err = bucketInstance.WaitUntilReady(5*time.Second, nil)
	return bucketInstance, err
}

// ConnectToScope, wrapper for Connecting to Scope, returns a Scope instance.
func ConnectToScope(bucketInstance *gocb.Bucket, name string) (scopeInstance *gocb.Scope) {
	scopeInstance = bucketInstance.Scope(name)
	return scopeInstance
}

// ConnectToCollection, wrapper for Connecting to Collection, returns a Collection instance.
func ConnectToCollection(scopeInstance *gocb.Scope, name string) (collectionInstance *gocb.Collection) {
	collectionInstance = scopeInstance.Collection(name)
	return collectionInstance
}
