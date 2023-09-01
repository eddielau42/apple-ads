package appleads

import (
	"encoding/json"
	"fmt"
	"net/url"
	_ "net/url"
	"strconv"

	"github.com/wangyong321/gogorequest"
)

// CreateKeywords 在广告组中创建目标关键字
func (engine *Engine)CreateTargetKeywords(campaignID, adGroupID int64, keywords []KeywordDetail) (*KeywordListResponse, error) {
	var (
		err error
		keywordsResp *KeywordListResponse
	)
	
	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/adgroups/" + strconv.FormatInt(adGroupID, 10) + "/targetingkeywords/bulk"
	
	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	reqBody, err := json.Marshal(keywords)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败: %w", err)
	}

	resp := reqCli.Visit("POST", link, reqHeaders, reqBody, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("创建广告组目标关键字失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &keywordsResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return keywordsResp, nil
}

// FindKeyword 查询获取广告组目标关键字
func (engine *Engine)FindTargetKeywords(campaignID int64, selector Selector) (*KeywordListResponse, error) {
	var (
		err error
		keywordsResp *KeywordListResponse
	)

	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/adgroups/targetingkeywords/find"

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
		return nil, fmt.Errorf("查询获取广告组目标关键字失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &keywordsResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return keywordsResp, nil
}

// GetKeyword 获取活动广告组目标关键字
func (engine *Engine)GetTargetKeyword(campaignID, adGroupID, keywordID int64) (*KeywordResponse, error) {
	var (
		err error
		keywordResp *KeywordResponse
	)

	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/adgroups/" + strconv.FormatInt(adGroupID, 10) + "/targetingkeywords/" + strconv.FormatInt(keywordID, 10)

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	resp := reqCli.Visit("GET", link, reqHeaders, nil, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("获取活动广告组目标关键字失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &keywordResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}
	
	return keywordResp, nil
}

// GetAllKeywords 获取活动广告组所有目标关键字
func (engine *Engine)GetAllTargetKeywords(campaignID, adGroupID int64) (*KeywordListResponse, error) {
	var (
		err error
		keywordsResp *KeywordListResponse
	)
	
	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/adgroups/" + strconv.FormatInt(adGroupID, 10) + "/targetingkeywords"
	
	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	resp := reqCli.Visit("GET", link, reqHeaders, nil, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("获取活动广告组所有目标关键字失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &keywordsResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return keywordsResp, nil
}

// UpdateKeyword 更新广告组目标关键字
func (engine *Engine)UpdateTargetKeywords(campaignID, adGroupID int64, keywords []Keyword) (*KeywordListResponse, error) {
	var (
		err error
		keywordsResp *KeywordListResponse
	)

	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/adgroups/" + strconv.FormatInt(adGroupID, 10) + "/targetingkeywords/bulk"
	
	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	reqBody, err := json.Marshal(keywords)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败: %w", err)
	}

	resp := reqCli.Visit("PUT", link, reqHeaders, reqBody, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("更新广告组目标关键字失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &keywordsResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}
	
	return keywordsResp, nil
}

// DeleteKeyword 删除活动广告组目标关键字
func (engine *Engine)DeleteTargetKeyword(campaignID, adGroupID, keywordID int64) (*VoidResponse, error) {
	var (
		err error
		voidResp *VoidResponse
	)

	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/adgroups/" + strconv.FormatInt(adGroupID, 10) + "/targetingkeywords/" + strconv.FormatInt(keywordID, 10)

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	resp := reqCli.Visit("DELETE", link, reqHeaders, nil, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("删除活动广告组目标关键字失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &voidResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return voidResp, nil
}

// DeleteKeyword 删除活动广告组所有目标关键字
func (engine *Engine)DeleteAllTargetKeywords(campaignID, adGroupID int64, keywordIDs []int64) (*IntegerResponse, error) {
	var (
		err error
		intResp *IntegerResponse
	)

	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/adgroups/" + strconv.FormatInt(adGroupID, 10) + "/targetingkeywords/delete/bulk"

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	reqBody, err := json.Marshal(keywordIDs)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败: %w", err)
	}

	resp := reqCli.Visit("POST", link, reqHeaders, reqBody, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("删除活动广告组目标关键字失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &intResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return intResp, nil
}

// CreateCampaignNegativeKeyword 创建活动否定关键字
func (engine *Engine)CreateCampaignNegativeKeywords(campaignID int64, keywords []NegativeKeyword) (*NegativeKeywordListResponse, error) {
	var (
		err error
		keywordsResp *NegativeKeywordListResponse
	)
	
	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/negativekeywords/bulk"
	
	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	reqBody, err := json.Marshal(keywords)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败: %w", err)
	}

	resp := reqCli.Visit("POST", link, reqHeaders, reqBody, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("创建否定关键字失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &keywordsResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return keywordsResp, nil
}

// FindCampaignNegativeKeyword 查询获取活动否定关键字
func (engine *Engine)FindCampaignNegativeKeywords(campaignID int64, selector Selector) (*NegativeKeywordListResponse, error) {
	var (
		err error
		keywordsResp *NegativeKeywordListResponse
	)

	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/negativekeywords/find"

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
		return nil, fmt.Errorf("查询获取活动否定关键字失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &keywordsResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return keywordsResp, nil
}

// GetCampaignNegativeKeyword 获取活动否定关键字
func (engine *Engine)GetCampaignNegativeKeyword(campaignID, keywordID int64) (*NegativeKeywordResponse, error) {
	var (
		err error
		keywordResp *NegativeKeywordResponse
	)

	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/negativekeywords/" + strconv.FormatInt(keywordID, 10)

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	resp := reqCli.Visit("GET", link, reqHeaders, nil, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("获取活动否定关键字失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &keywordResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return keywordResp, nil
}

// GetCampaignNegativeKeyword 获取活动所有否定关键字
func (engine *Engine)GetAllCampaignNegativeKeywords(campaignID int64, limit, offset int32) (*NegativeKeywordListResponse, error) {
	var (
		err error
		keywordsResp *NegativeKeywordListResponse
	)

	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/negativekeywords"

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
		return nil, fmt.Errorf("获取活动所有否定关键字失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &keywordsResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return keywordsResp, nil
}

// UpdateNegativeKeyword 更新活动否定关键字
func (engine *Engine)UpdateNegativeKeywords(campaignID int64, keywords []NegativeKeyword) (*NegativeKeywordListResponse, error) {
	var (
		err error
		keywordsResp *NegativeKeywordListResponse
	)
	
	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/negativekeywords/bulk"
	
	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	reqBody, err := json.Marshal(keywords)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败: %w", err)
	}

	resp := reqCli.Visit("PUT", link, reqHeaders, reqBody, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("更新活动否定关键字失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &keywordsResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return keywordsResp, nil
}

// DeleteCampaignNegativeKeywords 删除活动否定关键字
func (engine *Engine)DeleteCampaignNegativeKeywords(campaignID int64, keywordIDs []int64) (*IntegerResponse, error) {
	var (
		err error
		intResp *IntegerResponse
	)

	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/negativekeywords/delete/bulk"
	
	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	reqBody, err := json.Marshal(keywordIDs)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败: %w", err)
	}

	resp := reqCli.Visit("POST", link, reqHeaders, reqBody, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("删除活动否定关键字失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &intResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return intResp, nil
}

// CreateAdGroupNegativeKeywords 创建广告组否定关键字
func (engine *Engine)CreateAdGroupNegativeKeywords(campaignID, adGroupID int64, keywords []NegativeKeyword) (*NegativeKeywordListResponse, error) {
	var (
		err error
		keywordsResp *NegativeKeywordListResponse
	)
	
	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/adgroups/" + strconv.FormatInt(adGroupID, 10) + "/negativekeywords/bulk"
	
	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	reqBody, err := json.Marshal(keywords)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败: %w", err)
	}

	resp := reqCli.Visit("POST", link, reqHeaders, reqBody, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("创建广告组否定关键字失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &keywordsResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return keywordsResp, nil
}

// FindAdGroupNegativeKeyword 查询获取广告组否定关键字
func (engine *Engine)FindAdGroupNegativeKeywords(campaignID int64, selector Selector) (*NegativeKeywordListResponse, error) {
	var (
		err error
		keywordsResp *NegativeKeywordListResponse
	)

	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/adgroups/negativekeywords/find"

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
		return nil, fmt.Errorf("查询获取广告组否定关键字失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &keywordsResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return keywordsResp, nil
}

// GetAdGroupNegativeKeyword 获取广告组否定关键字
func (engine *Engine)GetAdGroupNegativeKeyword(campaignID, adGroupID, keywordID int64) (*NegativeKeywordResponse, error) {
	var (
		err error
		keywordResp *NegativeKeywordResponse
	)

	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/adgroups/" + strconv.FormatInt(adGroupID, 10) + "/negativekeywords/" + strconv.FormatInt(keywordID, 10)

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	resp := reqCli.Visit("GET", link, reqHeaders, nil, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("获取广告组否定关键字失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &keywordResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return keywordResp, nil
}

// GetCampaignNegativeKeyword 获取广告组所有否定关键字
func (engine *Engine)GetAllAdGroupNegativeKeywords(campaignID, adGroupID int64, limit, offset int32) (*NegativeKeywordListResponse, error) {
	var (
		err error
		keywordsResp *NegativeKeywordListResponse
	)

	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/adgroups/" + strconv.FormatInt(adGroupID, 10) + "/negativekeywords"

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
		return nil, fmt.Errorf("获取广告组所有否定关键字失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &keywordsResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return keywordsResp, nil
}

// UpdateNegativeKeyword 更新广告组否定关键字
func (engine *Engine)UpdateAdGroupNegativeKeywords(campaignID, adGroupID int64, keywords []NegativeKeyword) (*NegativeKeywordListResponse, error) {
	var (
		err error
		keywordsResp *NegativeKeywordListResponse
	)
	
	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/adgroups/" + strconv.FormatInt(adGroupID, 10) + "/negativekeywords/bulk"
	
	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	reqBody, err := json.Marshal(keywords)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败: %w", err)
	}

	resp := reqCli.Visit("PUT", link, reqHeaders, reqBody, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("更新广告组否定关键字失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &keywordsResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return keywordsResp, nil
}

// DeleteCampaignNegativeKeywords 删除广告组否定关键字
func (engine *Engine)DeleteAdGroupNegativeKeywords(campaignID, adGroupID int64, keywordIDs []int64) (*IntegerResponse, error) {
	var (
		err error
		intResp *IntegerResponse
	)

	link := ApiBaseURL() + "/campaigns/" + strconv.FormatInt(campaignID, 10) + "/adgroups/" + strconv.FormatInt(adGroupID, 10) + "/negativekeywords/delete/bulk"
	
	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	reqBody, err := json.Marshal(keywordIDs)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败: %w", err)
	}

	resp := reqCli.Visit("POST", link, reqHeaders, reqBody, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("删除广告组否定关键字失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &intResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return intResp, nil
}