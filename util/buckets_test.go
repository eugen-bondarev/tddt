package util

import (
	"testing"
)

func TestGetBucketName(t *testing.T) {
	var bucket, path string
	bucket, path = GetBucketName("test/test.txt")
	if bucket != "test" {
		t.Errorf("expected test, got %s", bucket)
	}
	if path != "test.txt" {
		t.Errorf("expected test.txt, got %s", path)
	}

	bucket, path = GetBucketName("foo/bar/baz/test.txt")
	if bucket != "foo" {
		t.Errorf("expected foo, got %s", bucket)
	}

	if path != "bar/baz/test.txt" {
		t.Errorf("expected bar/baz/test.txt, got %s", path)
	}
}
