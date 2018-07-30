package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	expectedConnectionString := "user=postgres password=postgres port=5432 dbname=comet_test sslmode=disable"

	assert.Equal(t, expectedConnectionString, ConnectionString())
}
