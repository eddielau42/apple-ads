package appleads

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/wangyong321/gogorequest"
)

type (
	AdCreateRequest Ad
	AdUpdateRequest Ad
)

// CreateAd 创建广告
func (engine *Engine)CreateAd(campaignID, adGroupID int64, adCreateData *AdCreateRequest) (*AdResponse, error) {
	var (
		err error
		adResp *AdResponse
	)

	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/adgroups/" + strconv.FormatInt(adGroupID, 10) + "/ads"

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	reqBody, err := json.Marshal(adCreateData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败: %w", err)
	}

	resp := reqCli.Visit("POST", link, reqHeaders, string(reqBody), apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("创建广告失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &adResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return adResp, nil
}

// FindAds 查找广告活动广告
func (engine *Engine)FindAds(campaignID int64, selector Selector) (*AdListResponse, error) {
	var (
		err error
		adsResp *AdListResponse
	)

	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/ads/find"

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
		return nil, fmt.Errorf("查找广告活动广告失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &adsResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return adsResp, nil
}

// GetAd 获取广告
func (engine *Engine)GetAd(campaignID, adGroupID, adID int64) (*AdResponse, error) {
	var (
		err error
		adResp *AdResponse
	)

	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/adgroups/" + strconv.FormatInt(adGroupID, 10) + "/ads/" + strconv.FormatInt(adID, 10)

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	resp := reqCli.Visit("GET", link, reqHeaders, nil, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("获取广告失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &adResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return adResp, nil
}

// GetAllAds 获取活动中所有广告
func (engine *Engine)GetAllAds(campaignID, adGroupID int64) (*AdListResponse, error) {
	var (
		err error
		adsResp *AdListResponse
	)

	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/adgroups/" + strconv.FormatInt(adGroupID, 10) + "/ads/"

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	resp := reqCli.Visit("GET", link, reqHeaders, nil, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("获取活动中所有广告失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &adsResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return adsResp, nil
}

// UpdateAd 更新广告
func (engine *Engine)UpdateAd(campaignID, adGroupID, adID int64, adUpdateData *AdUpdateRequest) (*AdResponse, error) {
	var (
		err error
		adResp *AdResponse
	)

	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/adgroups/" + strconv.FormatInt(adGroupID, 10) + "/ads/" + strconv.FormatInt(adID, 10)

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	reqBody, err := json.Marshal(adUpdateData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败: %w", err)
	}

	resp := reqCli.Visit("PUT", link, reqHeaders, reqBody, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("更新广告失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &adResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return adResp, nil
}

// DeleteAd 删除广告
func (engine *Engine)DeleteAd(campaignID, adGroupID, adID int64) (*VoidResponse, error) {
	var (
		err error
		voidResp *VoidResponse
	)

	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/adgroups/" + strconv.FormatInt(adGroupID, 10) + "/ads/" + strconv.FormatInt(adID, 10)

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	resp := reqCli.Visit("DELETE", link, reqHeaders, nil, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("删除广告失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &voidResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return voidResp, nil
}