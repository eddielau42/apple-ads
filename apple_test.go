package appleads

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

const (
	envPath = "./.env"
)

var (
	conf map[string]string
	appleEngine *Engine
)

func setup(t *testing.T) func(t *testing.T) {
	// Setup...

	conf, _ = godotenv.Read(envPath)

	appleEngine = New(
		WithClientID(conf["APPLE_CLIENT_ID"]),
		WithTeamID(conf["APPLE_TEAM_ID"]),
		WithKeyID(conf["APPLE_KEY_ID"]),
		WithPrivateKey(conf["APPLE_PRIVATE_KEY"]),
	)

    return func(t *testing.T) {
        // Teardown...
    }
}

func TestAppleEngine(t *testing.T) {
	teardown := setup(t)
    defer teardown(t)

	t.Logf("env-config: %+v\n", conf)

	t.Logf("apple-engine: %+v\n", appleEngine)

	assert.NotEmpty(t, appleEngine)
}

func TestAppleDetail(t *testing.T) {
	teardown := setup(t)
    defer teardown(t)

	appleEngine.Auth()
	t.Logf("apple-engine: %+v\n", conf)

	// ACL
	aclResp, err := appleEngine.UserAcl()
	t.Logf("user-ACL: %+v, error: %+v\n", aclResp, err)
	assert.NotEmpty(t, aclResp)

	// Me detail
	meResp, err := appleEngine.Me()
	t.Logf("me-detail: %+v, error: %+v\n", meResp, err)
	assert.NotEmpty(t, meResp)
}

