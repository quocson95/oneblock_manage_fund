package thirdparty

import (
	"oneblock_manage_fund/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserAsset(t *testing.T) {
	name := "TestUserAsset"
	config.ParseFromPath("/home/sondq/oneblock_manage_fund/be/config.yaml")
	t.Run(name, func(t *testing.T) {
		got, err := UserAsset()
		assert.NoError(t, err)
		assert.NotNil(t, got)
	})
}

func TestUserBalance(t *testing.T) {
	name := "TestUserBalance"
	config.ParseFromPath("/home/sondq/oneblock_manage_fund/be/config.yaml")
	t.Run(name, func(t *testing.T) {
		got, err := UserBalance()
		assert.NoError(t, err)
		assert.NotNil(t, got)
	})
}
