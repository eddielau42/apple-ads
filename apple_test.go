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


func TestKeyword(t *testing.T) {
	teardown := setup(t)
    defer teardown(t)

	campaignID := 1440199010
	adGroupID := 1440263406

	// keywordsResp, err := appleEngine.GetAllTargetKeywords(int64(campaignID), int64(adGroupID))
	// t.Logf("get-keywords: %+v, error: %+v\n", keywordsResp, err)
	// assert.NotEmpty(t, keywordsResp)

	keywordID := 1440294362
	keywordResp, err := appleEngine.GetTargetKeyword(int64(campaignID), int64(adGroupID), int64(keywordID))
	t.Logf("get-keyword: %+v, error: %+v\n", keywordResp, err)
	assert.NotEmpty(t, keywordResp)

	// limit := 1
	// offset := 0
	// nkeywordsResp, err := appleEngine.GetAllCampaignNegativeKeywords(int64(campaignID), int32(limit), int32(offset))
	// t.Logf("get-negative-keywords: %+v, error: %+v\n", nkeywordsResp, err)
	// assert.NotEmpty(t, nkeywordsResp)

}

func TestSearch(t *testing.T) {
	teardown := setup(t)
    defer teardown(t)

	limit, offset := 10, 0

	searchWords, owned := "筷子", false
	appsResp, err := appleEngine.SearchApps(searchWords, owned, int32(limit), int32(offset))
	t.Logf("search-apps: %+v, error: %+v\n", appsResp, err)
	assert.NotEmpty(t, appsResp)

	countrycode, entity := "PH", ""
	searchWords = "Manila"
	geosResp, err := appleEngine.SearchGeo(countrycode, entity, searchWords, int32(limit), int32(offset))
	t.Logf("search-geos: %+v, error: %+v\n", geosResp, err)
	assert.NotEmpty(t, geosResp)

	cond := make([]GeoRequest, 0)
	cond = append(cond, GeoRequest{
		ID: "US|CA|Cupertino",
		Entity: "locality",
	})
	geosResp2, err := appleEngine.GetGeoList(cond, int32(limit), int32(offset))
	t.Logf("get-geos: %+v, error: %+v\n", geosResp2, err)
	assert.NotEmpty(t, geosResp2)
}


func fakeAuth() {
	appleEngine.SetAccessToken(&AccessToken{
		TokenType: "Bearer",
		AccessToken: "eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2R0NNIiwia2lkIjpudWxsfQ..A9kMAoFQ2CxwTdwu.Ft4Xj27R1FVf_LmalFNaegntD-VMCfyb1bNz9-8rCVhdE7S8mIv3uUeK3kIUPRZFkBR_FxT8h_SJhbltDWXdy3FbyzQfmPowcRGrKuHDcuI5wScckiSw7lQlJBQlCmXWlRJMb7lb3_HCPMvc89lOeM__QAciEqY7ywVIz9cFQNE4XjyGcflia_wFxzVFKHsQme9L4xY6MITHfvQIjMv5WDOGiROyx60R_PTtLBmtHgDDgE2GIZHcOEQc3LqD3fX77b2y9F3gTO7Z6BhB9Y16VeI.vzffORXqBVUZ9xUpWbMueA",
		ExpiresIn: 3600,
	})
	appleEngine.SetJwt("eyJhbGciOiJFUzI1NiIsImtpZCI6Ijc4YWRhMWFkLTk5OTItNDVkNi04OTE0LTdiODE1YmY5Njk2MyIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJodHRwczovL2FwcGxlaWQuYXBwbGUuY29tIiwiZXhwIjoxNzA5MDIwMTY3LCJpYXQiOjE2OTM0NjgxNjcsImlzcyI6IlNFQVJDSEFEUy5mYWU1OGI3MS00MjRlLTRkNDItYjg2Zi1hYmIwMzRjYmY4MmUiLCJzdWIiOiJTRUFSQ0hBRFMuZmFlNThiNzEtNDI0ZS00ZDQyLWI4NmYtYWJiMDM0Y2JmODJlIn0.INy5xbEj68_NtnZT_oDmgkyT8BWxbkT5oPd6f0p0qcDLLRv8kNP2w_SNurm7DLiwgDl6RnSnecw9C1jGofVb1Q")
}