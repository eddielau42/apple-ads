package appleads

import (
	"encoding/json"
	"fmt"
	"net/url"
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

	CustomReportRequest struct {
		DateRange string `json:"dateRange,omitempty"`
		StartTime string `json:"startTime,omitempty"`
		EndTime string `json:"endTime,omitempty"`
		Granularity string `json:"granularity,omitempty"`
		Name string `json:"name"`
		Selector *Selector `json:"selector,omitempty"`
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

// GetImpressionShareReport 获取报告ID
func (engine *Engine)GetImpressionShareReport(reportRequest *CustomReportRequest) (*CustomReportResponse, error) {
	var (
		err error
		repResp *CustomReportResponse
	)

	link := ApiBaseURL() + "/custom-reports"

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
		return nil, fmt.Errorf("获取报告ID失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &repResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return repResp, nil
}

// GetSingleImpressionShareReport 获取包含指标和元数据的单个印象份额报告
func (engine *Engine)GetSingleImpressionShareReport(reportID int64) (*CustomReportResponse, error) {
	var (
		err error
		repResp *CustomReportResponse
	)

	link := ApiBaseURL() + "/custom-reports/" + strconv.FormatInt(reportID, 10)

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	resp := reqCli.Visit("GET", link, reqHeaders, nil, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("获取包含指标和元数据的单个印象份额报告失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &repResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return repResp, nil
}
// GetAllImpressionShareReports 获取包含指标和元数据的所有印象份额报告
func (engine *Engine)GetAllImpressionShareReports(field, sort string, limit, offset int32) (*CustomReportResponse, error) {
	var (
		err error
		repResp *CustomReportResponse
	)

	link := ApiBaseURL() + "/custom-reports"

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	query := url.Values{}
	if field != "" {
		query.Add("field", field)
	}
	if sort != "" {
		query.Add("sortOrder", sort)
	}
	if limit < 0 {
		limit = 20
	} else if limit > 50 {
		limit = 50
	}
	query.Add("limit", strconv.Itoa(int(limit)))
	if offset < 0 {
		limit = 0
	}
	query.Add("offset", strconv.Itoa(int(offset)))

	if len(query) > 0{
		link += "?" + query.Encode()
	}

	resp := reqCli.Visit("GET", link, reqHeaders, nil, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("获取包含指标和元数据的所有印象份额报告失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &repResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return repResp, nil
}