package appleads

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

	// ErrorResponse
	ErrorResponse struct {
		Errors []ErrorItem `json:"errors"`
	}

	// ApiBaseResponse
	ApiBaseResponse struct {
		Pagination PageDetail `json:"pagination,omitempty"`
		Error ErrorResponse `json:"error"`
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
		ApiBaseResponse
		Data []UserAcl `json:"data"`
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