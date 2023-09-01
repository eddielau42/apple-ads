package appleads

type (
	// AuthorizationResponse 授权返回响应数据
	AuthorizationResponse struct {
		AccessToken
		Error       string `json:"error"`
	}

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

	// ErrorResponse
	ErrorResponse struct {
		Errors []ErrorItem `json:"errors"`
	}

	// ApiBaseResponse
	ApiBaseResponse struct {
		Pagination PageDetail `json:"pagination,omitempty"`
		Error ErrorResponse `json:"error"`
	}

	VoidResponse struct {
		Data interface{} `json:"data"`
		ApiBaseResponse
	}

	Money struct {
		Amount string `json:"amount"`
		Currency string `json:"currency"`
	}

	LOCInvoiceDetails struct {
		BillingContactEmail string `json:"billingContactEmail"`
		BuyerEmail string `json:"buyerEmail"`
		BuyerName string `json:"buyerName"`
		ClientName string `json:"clientName"`
		OrderNumber string `json:"orderNumber"`
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

type (
	AdGroup struct {
		AutomatedKeywordsOptIn bool `json:"automatedKeywordsOptIn"`
		CpaGoal Money `json:"cpaGoal"`
		DefaultBidAmount Money `json:"defaultBidAmount"`
		EndTime string `json:"endTime"`
		Name string `json:"name"`
		StartTime string `json:"startTime"`
		Status string `json:"status"`
		TargetingDimensions interface{} `json:"targetingDimensions"`
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