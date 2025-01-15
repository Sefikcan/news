package couchbase

import (
	"github.com/couchbase/gocb/v2"
	"log"
	"time"
)

type Client struct {
	Bucket *gocb.Bucket
}

func NewCouchbaseClient(host, userName, password, bucketName string) *Client {
	cluster, err := gocb.Connect("couchbase://"+host, gocb.ClusterOptions{
		Username: userName,
		Password: password,
	})
	if err != nil {
		log.Fatalf("Connection failed: %v", err)
	}
	bucket := cluster.Bucket(bucketName)
	if err := bucket.WaitUntilReady(10*time.Second, nil); err != nil {
		log.Fatalf("Bucket not ready: %v", err)
	}

	return &Client{Bucket: bucket}
}

func (cb *Client) Collection() *gocb.Collection {
	return cb.Bucket.DefaultCollection()
}
