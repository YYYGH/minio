package storage

import (
	"bytes"
	"math/rand"
	"strconv"

	. "gopkg.in/check.v1"
)

func APITestSuite(c *C, create func() Storage) {
	testCreateBucket(c, create)
	testMultipleObjectCreation(c, create)
	testPaging(c, create)
	testObjectOverwriteFails(c, create)
	testNonExistantBucketOperations(c, create)
}

func testCreateBucket(c *C, create func() Storage) {
	// test create bucket
	// test bucket exists
	// test no objects exist
	// 2x
}

func testMultipleObjectCreation(c *C, create func() Storage) {
	objects := make(map[string][]byte)
	storage := create()
	storage.StoreBucket("bucket")
	for i := 0; i < 10; i++ {
		randomPerm := rand.Perm(10)
		randomString := ""
		for _, num := range randomPerm {
			randomString = randomString + strconv.Itoa(num)
		}
		key := "obj" + strconv.Itoa(i)
		objects[key] = []byte(randomString)
		err := storage.StoreObject("bucket", key, bytes.NewBufferString(randomString))
		c.Assert(err, IsNil)
	}

	// ensure no duplicates
	etags := make(map[string]string)
	for key, value := range objects {
		var byteBuffer bytes.Buffer
		storage.CopyObjectToWriter(&byteBuffer, "bucket", key)
		c.Assert(bytes.Equal(value, byteBuffer.Bytes()), Equals, true)

		metadata, err := storage.GetObjectMetadata("bucket", key)
		c.Assert(err, IsNil)
		c.Assert(metadata.Size, Equals, len(value))

		_, ok := etags[metadata.ETag]
		c.Assert(ok, Equals, false)
		etags[metadata.ETag] = metadata.ETag
	}
}

func testPaging(c *C, create func() Storage) {
}

func testObjectOverwriteFails(c *C, create func() Storage) {
	// test overwriting object fails
}
func testNonExistantBucketOperations(c *C, create func() Storage) {
	// test writing object in non-existant bucket fails
}
