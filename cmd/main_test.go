package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getting_user_secrets(t *testing.T) {
	err := AddUserSecretsIfApplicable()

	assert.Error(t, err)
	t.Log(err)
}
