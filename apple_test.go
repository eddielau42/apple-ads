package appleads

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

const (
	envPath = "./.env"

	orgID = 1234567
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

	fakeAuth()
	appleEngine.SetOrgID(orgID)

    return func(t *testing.T) {
        // Teardown...
    }
}

func TestAppleDetail(t *testing.T) {
	teardown := setup(t)
    defer teardown(t)

	// appleEngine.Auth()
	// t.Logf("auth-info: %+v\n", appleEngine.AuthInfo())

	// ACL
	aclResp, err := appleEngine.UserAcl()
	t.Logf("user-ACL: %+v, error: %+v\n", aclResp, err)
	assert.NotEmpty(t, aclResp)

	// Me detail
	meResp, err := appleEngine.Me()
	t.Logf("me-detail: %+v, error: %+v\n", meResp, err)
	assert.NotEmpty(t, meResp)
}

func TestCampaigns(t *testing.T) {
	teardown := setup(t)
    defer teardown(t)

	resp, err := appleEngine.AllCampaign(1, 0)
	t.Logf("all-campaign: %+v, error: %+v\n", resp, err)
	assert.NotEmpty(t, resp)
}


func fakeAuth() {
	appleEngine.SetAccessToken(&AccessToken{
		TokenType: "Bearer",
		AccessToken: "eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2R0NNIiwia2lkIjpudWxsfQ..BTfyJjgkmNn1ViMt.SzvponRxmYCF0XOTHOS3qx6Q3BVl5BSsIWRqvThOmvlSXKliubO3xM9C1mgSEXR4416A3ht7PPv-jtjQlq_OJrYum96vPOxj6J1icx8ATreg9pb-qAgxVp0QOPC5tM9goT1lPBu8u1nF4hsy7Jh7CMwOdpK0WC5kfhlOXEmilb6AZkABkSQefBHPi28fSDJjbFtGnIUHmJSRXHHntngommeNHUymDPb-rAXMvSjWmBDcDrWdSAWVrL4BGHwsAEZnFnyyPJC0kp_ddm4dB7jYpO4.sVAiKeOz4AEXejdSLSBbfA",
		ExpiresIn: 3600,
	})
	appleEngine.SetJwt("eyJhbGciOiJFUzI1NiIsImtpZCI6Ijc4YWRhMWFkLTk5OTItNDVkNi04OTE0LTdiODE1YmY5Njk2MyIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJodHRwczovL2FwcGxlaWQuYXBwbGUuY29tIiwiZXhwIjoxNzA5MTA2Mzg5LCJpYXQiOjE2OTM1NTQzODksImlzcyI6IlNFQVJDSEFEUy5mYWU1OGI3MS00MjRlLTRkNDItYjg2Zi1hYmIwMzRjYmY4MmUiLCJzdWIiOiJTRUFSQ0hBRFMuZmFlNThiNzEtNDI0ZS00ZDQyLWI4NmYtYWJiMDM0Y2JmODJlIn0.W9VFut2PKu74CWVVaST45qnFEDI65N3Rpfagr_LtJQh9ZKwWNGeemNKen0meyiNBRpC75fKHy9PORj42qen0zg")
}