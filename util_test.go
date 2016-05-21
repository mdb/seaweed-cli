package main

import (
	"flag"
	"os"
	"testing"
	"time"

	"github.com/codegangsta/cli"
	"github.com/mdb/seaweed"
)

func TestConcat(t *testing.T) {
	joined := concat([]string{
		"foo",
		"bar",
	})

	if joined != "foobar" {
		t.Error("concat should properly concatenate strings")
	}
}

func TestCacheAgeWithEnvVar(t *testing.T) {
	os.Setenv("MAGIC_SEAWEED_CACHE_AGE", "10m")
	age, _ := time.ParseDuration("10m")

	if cacheAge() != age {
		t.Error("cacheAge should properly use the $MAGIC_SEAWEED_CACHE_AGE value")
	}
}

func TestCacheAgeWithNoEnvVar(t *testing.T) {
	os.Setenv("MAGIC_SEAWEED_CACHE_AGE", "")
	age, _ := time.ParseDuration("5m")

	if cacheAge() != age {
		t.Error("cacheAge should properly default to 5m when no $MAGIC_SEAWEED_CACHE_AGE is set")
	}
}

func TestCacheDirWithEnvVar(t *testing.T) {
	os.Setenv("MAGIC_SEAWEED_CACHE_DIR", "/tmp")

	if cacheDir() != "/tmp" {
		t.Error("cacheDir should properly use the $MAGIC_SEAWEED_CACHE_DIR value")
	}
}

func TestCacheDirWithNoEnvVar(t *testing.T) {
	os.Setenv("MAGIC_SEAWEED_CACHE_DIR", "")

	if cacheDir() != os.TempDir() {
		t.Error("cacheDir should properly default to os.TempDir() when no $MAGIC_SEAWEED_CACHE_DIR is set")
	}
}

func TestClientWithDefaults(t *testing.T) {
	os.Setenv("MAGIC_SEAWEED_API_KEY", "123")

	c := getTestClient()
	cache, _ := time.ParseDuration("5m")

	if c.APIKey != "123" {
		t.Error("client.APIKey should be properly set")
	}
	if c.CacheAge != cache {
		t.Error("client.CacheAge should default to 5m")
	}
	if c.CacheDir != os.TempDir() {
		t.Error("client.CacheDir should default to os.TempDir()")
	}
}

func TestClientWithEnvVars(t *testing.T) {
	os.Setenv("MAGIC_SEAWEED_API_KEY", "123")
	os.Setenv("MAGIC_SEAWEED_CACHE_DIR", "/tmp")
	os.Setenv("MAGIC_SEAWEED_CACHE_AGE", "10m")

	c := getTestClient()
	cache, _ := time.ParseDuration("10m")

	if c.APIKey != "123" {
		t.Error("client.APIKey should be properly set")
	}
	if c.CacheAge != cache {
		t.Error("client.CacheAge should be set to $MAGIC_SEAWEED_CACHE_AGE")
	}
	if c.CacheDir != "/tmp" {
		t.Error("client.CacheDir should be set to $MAGIC_SEAWEED_CACHE_DIR")
	}
}

func getTestClient() *seaweed.Client {
	set := flag.NewFlagSet("test", 0)

	set.Int("myflag", 12, "doc")
	set.Float64("myflag64", float64(17), "doc")
	globalSet := flag.NewFlagSet("test", 0)
	globalSet.Int("myflag", 42, "doc")
	globalSet.Float64("myflag64", float64(47), "doc")

	globalCtx := cli.NewContext(nil, globalSet, nil)
	context := cli.NewContext(nil, set, globalCtx)

	return client(context)
}
