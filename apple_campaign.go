package appleads

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/wangyong321/gogorequest"
)

type (
	UpdateCampaign struct {
		BudgetAmount Money `json:"budgetAmount"`
		BudgetOrders []int64 `json:"budgetOrders"`
		CountriesOrRegions []string `json:"countriesOrRegions"`
		DailyBudgetAmount Money `json:"dailyBudgetAmount"`
		LocInvoiceDetails LOCInvoiceDetails `json:"locInvoiceDetails"`
	}

	UpdateCampaignRequest struct {
		Campaign UpdateCampaign `json:"campaign"`
		Clearable bool `json:"clearGeoTargetingOnCountryOrRegionChange"`
	}
)

// CreateCampaign 创建推广活动
func (engine *Engine)CreateCampaign(campaign Campaign) (*CampaignResponse, error) {
	var (
		err error
		campaignResp *CampaignResponse
	)

	link := ApiBaseURL() + "/campaigns"

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	reqBody, err := json.Marshal(campaign)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败: %w", err)
	}

	resp := reqCli.Visit("POST", link, reqHeaders, string(reqBody), apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("创建推广活动失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &campaignResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return campaignResp, nil
}

// FindCampaign 查找推广活动
func (engine *Engine)FindCampaign(selector *Selector) (*CampaignListResponse, error) {
	var (
		err error
		campaignListResp *CampaignListResponse
	)

	link := ApiBaseURL() + "/campaigns/find"

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	reqBody, err := json.Marshal(selector)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败: %w", err)
	}

	resp := reqCli.Visit("POST", link, reqHeaders, string(reqBody), apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("推广活动失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &campaignListResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return campaignListResp, nil
}

// GetCampaign 通过活动ID获取推广活动信息
func (engine *Engine)GetCampaign(campaignID int64) (*CampaignResponse, error) {
	var (
		err error
		campaignResp *CampaignResponse
	)

	link := ApiBaseURL() + "/campaigns" + "/" + strconv.FormatInt(campaignID, 10)

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	resp := reqCli.Visit("GET", link, reqHeaders, nil, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("创建推广活动失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &campaignResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return campaignResp, nil
}

// AllCampaign 获取所有推广活动
func (engine *Engine)AllCampaign(limit, offset int32) (*CampaignListResponse, error) {
	var (
		err error
		campaignListResp *CampaignListResponse
	)

	link := ApiBaseURL() + "/campaigns"

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	query := url.Values{}
	query.Add("limit", strconv.Itoa(int(limit)))
	query.Add("offset", strconv.Itoa(int(offset)))
	link += "?" + query.Encode()

	resp := reqCli.Visit("GET", link, reqHeaders, nil, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("获取推广活动失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &campaignListResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return campaignListResp, nil
}

// UpdateCampaign 更新推广活动
func (engine *Engine)UpdateCampaign(campaignID int64, updateReq *UpdateCampaignRequest) (*CampaignResponse, error) {
	var (
		err error
		campaignResp *CampaignResponse
	)

	link := ApiBaseURL() + "/campaigns" + "/" + strconv.FormatInt(campaignID, 10)

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	reqBody, err := json.Marshal(updateReq)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败: %w", err)
	}

	resp := reqCli.Visit("PUT", link, reqHeaders, string(reqBody), apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("更新推广活动失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &campaignResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return campaignResp, nil
}

// DeleteCampaign 删除推广活动
func (engine *Engine)DeleteCampaign(campaignID int64) (*VoidResponse, error) {
	var (
		err error
		voidResp *VoidResponse
	)

	link := ApiBaseURL() + "/campaigns" + "/" + strconv.FormatInt(campaignID, 10)

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	resp := reqCli.Visit("DELETE", link, reqHeaders, nil, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("删除推广活动失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &voidResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return voidResp, nil
}