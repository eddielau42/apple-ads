package appleads

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/wangyong321/gogorequest"
)

// CreateAdGroup 创建活动广告组
func (engine *Engine)CreateAdGroup(campaignID int64, adgroup AdGroup) (*AdGroupResponse, error) {
	var (
		err error
		adGroupResp *AdGroupResponse
	)

	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/adgroups"

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	reqBody, err := json.Marshal(adgroup)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败: %w", err)
	}

	resp := reqCli.Visit("POST", link, reqHeaders, reqBody, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("创建活动广告组失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &adGroupResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return adGroupResp, nil
}

// GetAllAdGroups 获取指定活动下所有广告组
func (engine *Engine)GetAllAdGroups(campaignID int64, limit, offset int32) (*AdGroupListResponse, error) {
	var (
		err error
		adGroupsResp *AdGroupListResponse
	)

	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/adgroups"

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
		return nil, fmt.Errorf("获取指定活动下所有广告组失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &adGroupsResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return adGroupsResp, nil
}

// GetAdGroup 获取指定活动广告组
func (engine *Engine)GetAdGroup(campaignID, adGroupID int64) (*AdGroupResponse, error) {
	var (
		err error
		adGroupResp *AdGroupResponse
	)

	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/adgroups/" + strconv.FormatInt(adGroupID, 10)

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	resp := reqCli.Visit("GET", link, reqHeaders, nil, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("获取指定活动广告组失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &adGroupResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return adGroupResp, nil

	return adGroupResp, nil
}

// FindAdGroup 查询获取活动中的广告组
func (engine *Engine)FindAdGroup(campaignID int64, selector Selector) (*AdGroupListResponse, error) {
	var (
		err error
		adGroupsResp *AdGroupListResponse
	)

	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/adgroups/find"

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

	resp := reqCli.Visit("POST", link, reqHeaders, reqBody, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("查询获取活动广告组失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &adGroupsResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return adGroupsResp, nil
}

// GetAdGroup 更新指定活动广告组
func (engine *Engine)UpdateAdGroup(campaignID, adGroupID int64, updateData AdGroup) (*AdGroupResponse, error) {
	var (
		err error
		adGroupResp *AdGroupResponse
	)

	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/adgroups" + strconv.FormatInt(adGroupID, 10)

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	reqBody, err := json.Marshal(updateData)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败: %w", err)
	}

	resp := reqCli.Visit("PUT", link, reqHeaders, reqBody, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("更新广告组失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &adGroupResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return adGroupResp, nil
}

// DeleteAdGroup 删除广告组
func (engine *Engine)DeleteAdGroup(campaignID, adGroupID int64) (*VoidResponse, error) {
	var (
		err error
		voidResp *VoidResponse
	)

	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/adgroups" + strconv.FormatInt(adGroupID, 10)

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	resp := reqCli.Visit("DELETE", link, reqHeaders, nil, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("删除广告组失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &voidResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}
	
	return voidResp, nil
}