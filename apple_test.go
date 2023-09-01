package appleads

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

const (
	envPath = "./.env"

	orgID = 7836740
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

	campsResp, err := appleEngine.AllCampaign(1, 0)
	t.Logf("all-campaign: %+v, error: %+v\n", campsResp, err)
	assert.NotEmpty(t, campsResp)

	campaignID := 1435939707
	campResp, err := appleEngine.GetCampaign(int64(campaignID))
	t.Logf("get-campaign: %+v, error: %+v\n", campResp, err)
	assert.NotEmpty(t, campResp)
}

func TestAdGroups(t *testing.T) {
	teardown := setup(t)
    defer teardown(t)

	campaignID := 1440199010

	limit := 1
	offset := 0
	adGroupsResp, err := appleEngine.GetAllAdGroups(int64(campaignID), int32(limit), int32(offset))
	t.Logf("get-ad-groups: %+v, error: %+v\n", adGroupsResp, err)
	assert.NotEmpty(t, adGroupsResp)

	adGroupID := 1440263406
	adGroupResp, err := appleEngine.GetAdGroup(int64(campaignID), int64(adGroupID))
	t.Logf("get-ad-groups: %+v, error: %+v\n", adGroupResp, err)
	assert.NotEmpty(t, adGroupResp)
}

func fakeAuth() {
	appleEngine.SetAccessToken(&AccessToken{
		TokenType: "Bearer",
		AccessToken: "eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2R0NNIiwia2lkIjpudWxsfQ..pFBLKzgzWJocS-T6.cObIsmdw9SlkppYaunpp6l-QU8FJBtRctJOI8FSsA2GEDbPbK-MS_R6_jfFNa4c2SRWs4zWZ_cAzAGuMjymNZ3J9bsWkUYdJptL9n3xHmZA4BkvjkZ3EhMwPfwDPbtTA062ElHQK_L3V5Fq2rXiYXBfYRqRNMEQJSywEeB7aMqtMf6rG3L7gs77JjEkSbQzP8IvBB3IeHiYHdZgy6WVYcHa9EakCvtnaSJpCAjMII4r4PKEjdk7JpcHgmAz_e1uILo7I4BlX_yW8xtWMIAvKI-0.QFuBw8nQ3oOfueKgom8VkA",
		ExpiresIn: 3600,
	})
	appleEngine.SetJwt("eyJhbGciOiJFUzI1NiIsImtpZCI6Ijc4YWRhMWFkLTk5OTItNDVkNi04OTE0LTdiODE1YmY5Njk2MyIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJodHRwczovL2FwcGxlaWQuYXBwbGUuY29tIiwiZXhwIjoxNzA5MTExOTI4LCJpYXQiOjE2OTM1NTk5MjgsImlzcyI6IlNFQVJDSEFEUy5mYWU1OGI3MS00MjRlLTRkNDItYjg2Zi1hYmIwMzRjYmY4MmUiLCJzdWIiOiJTRUFSQ0hBRFMuZmFlNThiNzEtNDI0ZS00ZDQyLWI4NmYtYWJiMDM0Y2JmODJlIn0.jzsrp-BvM0ulyqtAFPYY4xlZFrYF5RZO2nLz3WGA3RlEUPk5B20i9OtNxBpq-l_IlzDJOTe0AuQGySk_sTw_Sg")
}