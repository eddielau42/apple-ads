package appleads

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/wangyong321/gogorequest"
)

type (
	ReportingRequest struct {
		EndTime string `json:"endTime"`
		Granularity string `json:"granularity,omitempty"`
		GroupBy []string `json:"groupBy,omitempty"`
		ReturnGrandTotals bool `json:"returnGrandTotals,omitempty"`
		ReturnRecordsWithNoMetrics bool `json:"returnRecordsWithNoMetrics,omitempty"`
		ReturnRowTotals bool `json:"returnRowTotals,omitempty"`
		StartTime string `json:"startTime"`
		TimeZone string `json:"timeZone"`
		Selector *Selector `json:"selector"`
	}
)

// GetCampaignReports 获取推广活动报告
func (engine *Engine)GetCampaignReports(reportRequest *ReportingRequest) (*ReportingResponse, error) {
	var (
		err error
		repResp *ReportingResponse
	)

	link := ApiBaseURL() + "/reports/campaigns"

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	reqBody, err := json.Marshal(reportRequest)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败: %w", err)
	}

	resp := reqCli.Visit("POST", link, reqHeaders, string(reqBody), apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("获取推广活动报告失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &repResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return repResp, nil
}

// GetAdGroupReports 获取推广活动广告组报告
func (engine *Engine)GetAdGroupReports(campaignID int64, reportRequest *ReportingRequest) (*ReportingResponse, error) {
	var (
		err error
		repResp *ReportingResponse
	)

	link := ApiBaseURL() + "/reports/campaigns/" + strconv.FormatInt(campaignID, 10) + "/adgroups"

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	reqBody, err := json.Marshal(reportRequest)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败: %w", err)
	}

	resp := reqCli.Visit("POST", link, reqHeaders, string(reqBody), apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("获取推广活动广告组报告失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &repResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return repResp, nil
}

// GetKeywordReports 获取营销活动中定位关键字的报告
func (engine *Engine)GetKeywordReports(campaignID int64, reportRequest *ReportingRequest) (*ReportingResponse, error) {
	var (
		err error
		repResp *ReportingResponse
	)

	link := ApiBaseURL() + "/reports/campaigns/" + strconv.FormatInt(campaignID, 10) + "/keywords"

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	reqBody, err := json.Marshal(reportRequest)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败: %w", err)
	}

	resp := reqCli.Visit("POST", link, reqHeaders, string(reqBody), apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("获取营销活动中定位关键字的报告失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &repResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return repResp, nil
}

// GetKeywordReportsWithinAdGroup 获取广告组中定位关键字的报告
func (engine *Engine)GetKeywordReportsWithinAdGroup(campaignID, adGroupID int64, reportRequest *ReportingRequest) (*ReportingResponse, error) {
	var (
		err error
		repResp *ReportingResponse
	)

	link := ApiBaseURL() + "/reports/campaigns/" + strconv.FormatInt(campaignID, 10) + "/adgroups/" + strconv.FormatInt(adGroupID, 10) + "/keywords"

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	reqBody, err := json.Marshal(reportRequest)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败: %w", err)
	}

	resp := reqCli.Visit("POST", link, reqHeaders, string(reqBody), apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("获取广告组中定位关键字的报告失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &repResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return repResp, nil
}

// GetSearchTermReports 获取活动中搜索词的报告
func (engine *Engine)GetSearchTermReports(campaignID int64, reportRequest *ReportingRequest) (*ReportingResponse, error) {
	var (
		err error
		repResp *ReportingResponse
	)

	link := ApiBaseURL() + "/reports/campaigns/" + strconv.FormatInt(campaignID, 10) + "/searchterms"

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	reqBody, err := json.Marshal(reportRequest)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败: %w", err)
	}

	resp := reqCli.Visit("POST", link, reqHeaders, string(reqBody), apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("获取活动中搜索词的报告失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &repResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return repResp, nil
}

// GetSearchTermReports 获取广告组内搜索字词的报告
func (engine *Engine)GetSearchTermReportsWithinAdGroup(campaignID, adGroupID int64, reportRequest *ReportingRequest) (*ReportingResponse, error) {
	var (
		err error
		repResp *ReportingResponse
	)

	link := ApiBaseURL() + "/reports/campaigns/" + strconv.FormatInt(campaignID, 10) + "/adgroups/" + strconv.FormatInt(adGroupID, 10) + "/searchterms"

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	reqBody, err := json.Marshal(reportRequest)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败: %w", err)
	}

	resp := reqCli.Visit("POST", link, reqHeaders, string(reqBody), apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("获取广告组内搜索字词的报告失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &repResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return repResp, nil
}

// GetAdReports 获取活动中的广告效果数据
func (engine *Engine)GetAdReports(campaignID int64, reportRequest *ReportingRequest) (*ReportingResponse, error) {
	var (
		err error
		repResp *ReportingResponse
	)

	link := ApiBaseURL() + "/reports/campaigns/" + strconv.FormatInt(campaignID, 10) + "/ads"

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	reqBody, err := json.Marshal(reportRequest)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败: %w", err)
	}

	resp := reqCli.Visit("POST", link, reqHeaders, string(reqBody), apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("获取活动中的广告效果数据失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &repResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return repResp, nil
}