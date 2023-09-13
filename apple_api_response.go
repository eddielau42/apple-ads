package appleads

// 基础数据结构
type (
	// PageDetail
	PageDetail struct {
		ItemsPerPage int32 `json:"itemsPerPage"`
		StartIndex int32 `json:"startIndex"`
		TotalResults int64 `json:"totalResults"`
	}
	// ErrorItem
	ErrorItem struct {
		Field string `json:"field"`
		Message string `json:"message"`
		Code string `json:"messageCode"`
	}
	// Money
	Money struct {
		Amount string `json:"amount"`
		Currency string `json:"currency"`
	}
	// LOCInvoiceDetails
	LOCInvoiceDetails struct {
		BillingContactEmail string `json:"billingContactEmail"`
		BuyerEmail string `json:"buyerEmail"`
		BuyerName string `json:"buyerName"`
		ClientName string `json:"clientName"`
		OrderNumber string `json:"orderNumber"`
	}
	// AgeRange
	AgeRange struct {
		MaxAge int32 `json:"maxAge"`
		MinAge int32 `json:"minAge"`
	}
	// TargetingDimensions
	TargetingDimensions struct {
		AdminArea *struct{
			Included []string `json:"included"`
		} `json:"adminArea"`
		
		Age *struct{
			Included []AgeRange `json:"included"`
		} `json:"age"`
		
		AppCategories struct{
			Excluded []int `json:"excluded"`
			Included []int `json:"included"`
		} `json:"appCategories"`
		
		AppDownloaders struct{
			Excluded []string `json:"excluded"`
			Included []string `json:"included"`
		} `json:"appDownloaders"`

		Country *struct{
			Included []string `json:"included"`
		} `json:"country"`

		Daypart *struct{
			UserTime struct{
				Included []int32 `json:"included"`
			} `json:"userTime"`
		} `json:"daypart"`

		DeviceClass *struct{
			Included []string `json:"included"`
		} `json:"deviceClass"`

		Gender *struct{
			Included []string `json:"included"`
		} `json:"gender"`

		Locality *struct{
			Included []string `json:"included"`
		} `json:"locality"`
	}
)

// 基本响应结构
type (
	// AuthorizationResponse 授权返回响应数据
	AuthorizationResponse struct {
		AccessToken
		Error       string `json:"error"`
	}
	// ErrorResponse
	ErrorResponse struct {
		Errors []ErrorItem `json:"errors"`
	}
	// ApiBaseResponse
	ApiBaseResponse struct {
		Pagination PageDetail `json:"pagination,omitempty"`
		Error ErrorResponse `json:"error"`
	}
	// VoidResponse
	VoidResponse struct {
		Data interface{} `json:"data"`
		ApiBaseResponse
	}
	IntegerResponse struct {
		Data int32 `json:"data"`
		ApiBaseResponse
	}
)

type (
	UserAcl struct {
		Currency string `json:"currency"`
		OrgID int64 `json:"orgId"`
		OrgName string `json:"orgName"`
		ParentOrgID string `json:"parentOrgId"`
		PaymentModel string `json:"paymentModel"`
		RoleNames []string `json:"roleNames"`
		TimeZone string `json:"timeZone"`
		DisplayName string `json:"displayName,omitempty"`
	}
	UserAclListResponse struct {
		Data []UserAcl `json:"data"`
		ApiBaseResponse
	}
)

type (
	MeDetail struct {
		ParentOrgID int64 `json:"parentOrgId"`
		UserID int64 `json:"userId"`
	}
	MeDetailResponse struct {
		Data MeDetail `json:"data"`
	}
)

// 活动相关
type (
	Campaign struct {
		AdminID int64 `json:"adamId"`
		AdChannelType string `json:"adChannelType"`
		BillingEvent string `json:"billingEvent"`
		BudgetAmount Money `json:"budgetAmount"`
		BudgetOrders []int64 `json:"budgetOrders"`
		CountriesOrRegions []string `json:"countriesOrRegions"`
		CountryOrRegionServingStateReasons map[string][]string `json:"countryOrRegionServingStateReasons"`
		CreationTime string `json:"creationTime"`
		DailyBudgetAmount Money `json:"dailyBudgetAmount"`
		Deleted bool `json:"deleted"`
		DisplayStatus string `json:"displayStatus"`
		EndTime string `json:"endTime"`
		ID int64 `json:"id"`
		LocInvoiceDetails LOCInvoiceDetails `json:"locInvoiceDetails"`
		ModificationTime string `json:"modificationTime"`
		Name string `json:"name"`
		OrgId int64 `json:"orgId"`
		PaymentModel string `json:"paymentModel"`
		ServingStateReasons []string `json:"servingStateReasons"`
		ServingStatus string `json:"servingStatus"`
		StartTime string `json:"startTime"`
		Status string `json:"status"`
		SupplySources []string `json:"supplySources"`
	}
	CampaignResponse struct {
		Data Campaign `json:"data"`
		ApiBaseResponse
	}
	CampaignListResponse struct {
		Data []Campaign `json:"data"`
		ApiBaseResponse
	}
)

// 广告组相关
type (
	AdGroup struct {
		AutomatedKeywordsOptIn bool `json:"automatedKeywordsOptIn"`
		CpaGoal Money `json:"cpaGoal"`
		DefaultBidAmount Money `json:"defaultBidAmount"`
		EndTime string `json:"endTime"`
		Name string `json:"name"`
		StartTime string `json:"startTime"`
		Status string `json:"status"`
		TargetingDimensions TargetingDimensions `json:"targetingDimensions"`
	}
	AdGroupResponse struct {
		Data AdGroup `json:"data"`
		ApiBaseResponse
	}
	AdGroupListResponse struct {
		Data []AdGroup `json:"data"`
		ApiBaseResponse
	}
)

// 关键字相关
type (
	Keyword struct {
		AdGroupID int64 `json:"adGroupId"`
		BidAmount Money `json:"bidAmount"`
		Deleted bool `json:"deleted"`
		ID int64 `json:"id"`
		MatchType string `json:"matchType"`
		ModificationTime string `json:"modificationTime"`
		Status string `json:"status"`
		Text string `json:"text"`
	}
	KeywordDetail struct {
		CampaignId int64 `json:"campaignId"`
		CreationTime string `json:"creationTime"`
		Keyword
	}
	KeywordResponse struct {
		Data KeywordDetail `json:"data"`
		ApiBaseResponse
	}
	KeywordListResponse struct {
		Data []KeywordDetail `json:"data"`
		ApiBaseResponse
	}

	NegativeKeyword struct {
		AdGroupID int64 `json:"adGroupId"`
		CampaignId int64 `json:"campaignId"`
		Deleted bool `json:"deleted"`
		ID int64 `json:"id"`
		MatchType string `json:"matchType"`
		ModificationTime string `json:"modificationTime"`
		Status string `json:"status"`
		Text string `json:"text"`
	}
	NegativeKeywordResponse struct {
		Data NegativeKeyword `json:"data"`
		ApiBaseResponse
	}
	NegativeKeywordListResponse struct {
		Data []NegativeKeyword `json:"data"`
		ApiBaseResponse
	}
)

// 搜索相关
type (
	AppInfo struct {
		AdamId int64 `json:"adamId"`
		AppName string `json:"appName"`
		DeveloperName string `json:"developerName"`
		CountryOrRegionCodes []string `json:"countryOrRegionCodes"`
	}
	AppInfoListResponse struct {
		Data []AppInfo `json:"data"`
		ApiBaseResponse
	}

	SearchEntity struct {
		AdminArea string `json:"adminArea"`
		CountryOrRegion string `json:"countryOrRegion"`
		DisplayName string `json:"displayName"`
		Entity string `json:"entity"`
		ID string `json:"id"`
		Locality string `json:"locality"`
	}
	SearchEntityListResponse struct {
		Data []SearchEntity `json:"data"`
		ApiBaseResponse
	}
)

// 广告相关
type (
	Ad struct {
		AdGroupID int64 `json:"adGroupId"`
		CampaignID int64 `json:"campaignId"`
		CreationTime string `json:"creationTime"`
		CreativeID int64 `json:"creativeId"`
		CreativeType string `json:"creativeType"`
		Deleted bool `json:"deleted"`
		ID int64 `json:"id"`
		ModificationTime string `json:"modificationTime"`
		Name string `json:"name"`
		OrgId int64 `json:"orgId"`
		ServingStateReasons []string `json:"servingStateReasons"`
		ServingStatus string `json:"servingStatus"`
		Status string `json:"status"`
	}
	AdResponse struct {
		Data Ad `json:"data"`
	}
	AdListResponse struct {
		Data []Ad `json:"data"`
		ApiBaseResponse
	}
)

// 报告相关
type (
	SpendRow struct {
		ConversionRate float64 `json:"conversionRate"`
		Impressions int64 `json:"impressions"`
		Installs int64 `json:"installs"`
		LatOffInstalls int64 `json:"latOffInstalls"`
		LatOnInstalls int64 `json:"latOnInstalls"`
		NewDownloads int64 `json:"newDownloads"`
		Redownloads int64 `json:"redownloads"`
		Taps int64 `json:"taps"`
		TTR float64 `json:"ttr"`
		AvgCPA Money `json:"avgCPA"`
		AvgCPM Money `json:"avgCPM"`
		AvgCPT Money `json:"avgCPT"`
		LocalSpend Money `json:"localSpend"`
	}
	ExtendedSpendRow struct {
		Date string `json:"date"`
		SpendRow
	}
	GrandTotals struct {
		Other bool `json:"other"`
		Total SpendRow `json:"total"`
	}
	ReportingRow struct {
		GrandTotals
		Insights interface{} `json:"insights,omitempty"`
		Granularity []ExtendedSpendRow `json:"granularity,omitempty"`
		MetaData interface{} `json:"metadata,omitempty"`
	}
	ReportingDataResponse struct {
		Row []ReportingRow `json:"row"`
		GrandTotals GrandTotals `json:"grandTotals"`
	}
	ReportingResponse struct {
		Data struct{
			ReportingDataResponse `json:"reportingDataResponse"`
		} `json:"data"`
		ApiBaseResponse
	}
)