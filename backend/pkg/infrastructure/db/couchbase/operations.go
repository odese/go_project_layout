package couchbase

import (
	"github.com/couchbase/gocb/v2"
)

// GetCollection, returns collection instance by name.
func GetCollection(name string) *gocb.Collection {
	return collections[name]
}

// CouchbaseCollection, represents collection instance, "wrapper" for gocb.Collection.
type CouchbaseCollection struct {
	*gocb.Collection
}

// Collection, returns "wrapped" collection instance by name.
func Collection(name string) *CouchbaseCollection {
	var col CouchbaseCollection
	col.Collection = GetCollection(name)
	return &col
}

// Key-Document Operations

// GetDocument, gets the document, associated with the key.
func (collection *CouchbaseCollection) GetDocument(key string, obj interface{}) (err error) {
	getResult, err := collection.Get(key, nil)
	if err != nil {
		return err
	}
	err = getResult.Content(&obj)
	if err != nil {
		return err
	}
	return err
}

// InsertDocument, inserts a document with the key and object.
func (collection *CouchbaseCollection) InsertDocument(key string, obj interface{}) (err error) {
	_, err = collection.Insert(key, obj, nil)
	return err
}

// UpsertDocument, upserts a document with the key and object.
func (collection *CouchbaseCollection) UpsertDocument(key string, obj interface{}) (err error) {
	_, err = collection.Upsert(key, obj, nil)
	return err
}

// DeleteDocument, deletes a document associated with the key.
func (collection *CouchbaseCollection) DeleteDocument(key string) (err error) {
	_, err = collection.Remove(key, nil)
	return err
}

// Sub-Document Operations

// LookupInDocument, gets the specified fields from the document.
func (collection *CouchbaseCollection) LookupInDocument(key string, obj interface{}, fields ...string) (err error) {
	ops := make([]gocb.LookupInSpec, 0)
	for _, field := range fields {
		ops = append(ops, gocb.GetSpec(field, nil))
	}

	getResult, err := collection.LookupIn(key, ops, nil)
	if err != nil {
		return err
	}
	err = getResult.ContentAt(0, &obj)
	if err != nil {
		return err
	}
	return err
}

// MutateInDocument, updates the specified fields in a document.
func (collection *CouchbaseCollection) MutateInDocument(key string, obj interface{}, m map[string]interface{}) (err error) {
	ops := make([]gocb.MutateInSpec, 0)
	for field, value := range m {
		ops = append(ops, gocb.UpsertSpec(field, value, nil))
	}

	_, err = collection.MutateIn(key, ops, nil)
	if err != nil {
		return err
	}
	return err
}

// ArrayAppendInDocument, appends the value to specified array field in a document.
func (collection *CouchbaseCollection) ArrayAppendInDocument(key string, obj interface{}, m map[string]interface{}) (err error) {
	ops := make([]gocb.MutateInSpec, 0)
	for field, value := range m {
		ops = append(ops, gocb.ArrayAppendSpec(field, value, nil))
	}

	_, err = collection.MutateIn(key, ops, nil)
	if err != nil {
		return err
	}
	return err
}

// ArrayPrependInDocument, prepends the value to specified array field in a document.
func (collection *CouchbaseCollection) ArrayPrependInDocument(key string, obj interface{}, m map[string]interface{}) (err error) {
	ops := make([]gocb.MutateInSpec, 0)
	for field, value := range m {
		ops = append(ops, gocb.ArrayPrependSpec(field, value, nil))
	}

	_, err = collection.MutateIn(key, ops, nil)
	if err != nil {
		return err
	}
	return err
}

// Query, executes the query.
func Query(query string) (results *gocb.QueryResult, err error) {
	results, err = cluster.Query(query, nil)
	if err != nil {
		return results, err
	}
	return results, err
}

// TODO: Decide which form of wrappers will be used.
// example usage:
// err = couchbase.Insert2(couchbase.GetCollection(utils.UserType), user.ID, user)
// func Insert2(collection *gocb.Collection, key string, obj interface{}) (err error) {
// 	_, err = collection.Insert(key, obj, nil)
// 	return err
// }

// TODO: ADD Full Text Search in couchbase.
