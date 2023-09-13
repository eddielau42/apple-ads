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

func TestAd(t *testing.T) {
	teardown := setup(t)
    defer teardown(t)

	campaignID := 1440199010
	adGroupID := 1440263406

	adsResp, err := appleEngine.GetAllAds(int64(campaignID), int64(adGroupID))
	t.Logf("get-all-ads: %+v, error: %+v\n", adsResp, err)
	assert.NotEmpty(t, adsResp)
}

func TestReporting(t *testing.T) {
	teardown := setup(t)
    defer teardown(t)

	repReq := &ReportingRequest{
		StartTime: "2023-08-01",
		EndTime: "2023-08-30",
		TimeZone: "UTC",
		Granularity: "DAILY", // HOURLY|DAILY|WEEKLY|MONTHLY
		ReturnGrandTotals: true,
		ReturnRecordsWithNoMetrics: true,
		ReturnRowTotals: true,
		GroupBy: []string{"countryOrRegion"},
		Selector: &Selector{
			OrderBy: []Sorting{
				{
					Field: "countryOrRegion",
					SortOrder: "ASCENDING",
				},
			},
			Pagination: Pagination{
				Offset: 0,
				Limit: 1,
			},
		},
	}
	repResp, err := appleEngine.GetCampaignReports(repReq)
	// t.Logf("get-campaign-reporting: %+v, error: %+v\n", repResp, err)
	assert.Empty(t, err)
	assert.NotEmpty(t, repResp)

	campaignID := 1440199010
	adGroupID := 1440263406

	repResp, err = appleEngine.GetAdGroupReports(int64(campaignID), repReq)
	// t.Logf("get-adgroup-reporting: %+v, error: %+v\n", repResp, err)
	assert.Empty(t, err)
	assert.NotEmpty(t, repResp)
	
	repResp, err = appleEngine.GetKeywordReports(int64(campaignID), repReq)
	// t.Logf("get-keyword-reporting: %+v, error: %+v\n", repResp, err)
	assert.Empty(t, err)
	assert.NotEmpty(t, repResp)
	
	repResp, err = appleEngine.GetKeywordReportsWithinAdGroup(int64(campaignID), int64(adGroupID), repReq)
	// t.Logf("get-keyword-within-adgroup-reporting: %+v, error: %+v\n", repResp, err)
	assert.Empty(t, err)
	assert.NotEmpty(t, repResp)
	
	repResp, err = appleEngine.GetSearchTermReports(int64(campaignID), repReq)
	// t.Logf("get-search-term-reporting: %+v, error: %+v\n", repResp, err)
	assert.Empty(t, err)
	assert.NotEmpty(t, repResp)
	
	repResp, err = appleEngine.GetSearchTermReportsWithinAdGroup(int64(campaignID), int64(adGroupID), repReq)
	// t.Logf("get-search-term-within-adgroup-reporting: %+v, error: %+v\n", repResp, err)
	assert.Empty(t, err)
	assert.NotEmpty(t, repResp)

	repResp, err = appleEngine.GetAdReports(int64(campaignID), repReq)
	// t.Logf("get-ad-reporting: %+v, error: %+v\n", repResp, err)
	assert.Empty(t, err)
	assert.NotEmpty(t, repResp)
}

func fakeAuth() {
	appleEngine.SetAccessToken(&AccessToken{
		TokenType: "Bearer",
		AccessToken: "eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2R0NNIiwia2lkIjpudWxsfQ..jepIOKda49YXGGWX.aX6NNNFls2_yfDtXYbbpue9KpNwdKOQceFaF-Zjs6_Ln6fc9jHspkgHhdJRSBU878fBA_eeigKTcJX1rHgy-CDjdQxLYr9k4PthFHcbOgjwPWrsXJn3La6fiZcW8jCFHKHhZ_NabdpY_XCmAPlgQf0PoXJqMxEIPTXfLrdpi4_8isYg4XcPKhet0niPpY4YilClFJIj8LvYJvrN8J3vAgFG0NG6j28v-1fL-T-JfTh113dvenepZb2SLx4pbw44x_J6ksiwf8vogVcz4_SK6Vx8.62SfKJDzUTAFKJGsPJsoEg",
		ExpiresIn: 3600,
	})
	appleEngine.SetJwt("eyJhbGciOiJFUzI1NiIsImtpZCI6Ijc4YWRhMWFkLTk5OTItNDVkNi04OTE0LTdiODE1YmY5Njk2MyIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJodHRwczovL2FwcGxlaWQuYXBwbGUuY29tIiwiZXhwIjoxNzA5MDIwMTY3LCJpYXQiOjE2OTM0NjgxNjcsImlzcyI6IlNFQVJDSEFEUy5mYWU1OGI3MS00MjRlLTRkNDItYjg2Zi1hYmIwMzRjYmY4MmUiLCJzdWIiOiJTRUFSQ0hBRFMuZmFlNThiNzEtNDI0ZS00ZDQyLWI4NmYtYWJiMDM0Y2JmODJlIn0.INy5xbEj68_NtnZT_oDmgkyT8BWxbkT5oPd6f0p0qcDLLRv8kNP2w_SNurm7DLiwgDl6RnSnecw9C1jGofVb1Q")
}