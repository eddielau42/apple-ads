package appleads

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/wangyong321/gogorequest"
)

type (
	GeoRequest struct {
		Entity string `json:"entity"`
		ID string `json:"id"`
	}
)

// SearchApps 搜索活动中推广的 iOS 应用程序
func (engine *Engine)SearchApps(searchWords string, owned bool, limit, offset int32) (*AppInfoListResponse, error) {
	const (
		_limit = 20
		_offset = 0
	)
	var (
		err error
		appsResp *AppInfoListResponse
	)

	link := ApiBaseURL() + "/search/apps"
	
	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	query := url.Values{}
	query.Add("query", searchWords)
	query.Add("returnOwnedApps", strconv.FormatBool(owned))

	if limit <= 0 {
		limit = _limit
	} else if limit > 1000 {
		limit = 1000
	}
	query.Add("limit", strconv.Itoa(int(limit)))

	if offset < _offset {
		offset = _offset
	}
	query.Add("offset", strconv.Itoa(int(offset)))

	link += "?" + query.Encode()
	resp := reqCli.Visit("GET", link, reqHeaders, nil, apiTimeout, "", nil)
	if resp.Error != nil {
		// 搜索要在营销活动中推广的 iOS 应用程序
		return nil, fmt.Errorf("获取广告组所有否定关键字失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &appsResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return appsResp, nil
}

// SearchGeo 获取用于定位的地理位置列表
func (engine *Engine) SearchGeo(countryCode, entity, searchWords string, limit, offset int32) (*SearchEntityListResponse, error) {
	const (
		_limit = 20
		_offset = 0
	)
	var (
		err error
		entitiesResp *SearchEntityListResponse
	)

	link := ApiBaseURL() + "/search/geo"

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	query := url.Values{}
	if countryCode != "" {
		query.Add("countrycode", countryCode)
	}
	if entity != "" {
		query.Add("entity", entity)
	}
	if searchWords != "" {
		query.Add("query", searchWords)
	}

	if limit <= 0 {
		limit = _limit
	} else if limit > 1000 {
		limit = 1000
	}
	query.Add("limit", strconv.Itoa(int(limit)))

	if offset < _offset {
		offset = _offset
	}
	query.Add("offset", strconv.Itoa(int(offset)))

	link += "?" + query.Encode()
	resp := reqCli.Visit("GET", link, reqHeaders, nil, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("获取用于定位的地理位置列表失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &entitiesResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return entitiesResp, nil
}

// GetGeoList 使用地理标识符获取地理位置详细信息
func (engine *Engine) GetGeoList(condition []GeoRequest, limit, offset int32) (*SearchEntityListResponse, error) {
	const (
		_limit = 20
		_offset = 0
	)
	var (
		err error
		entitiesResp *SearchEntityListResponse
	)

	link := ApiBaseURL() + "/search/geo"

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
		"X-AP-Context": fmt.Sprintf("orgId=%v", engine.orgID),
	}

	query := url.Values{}
	if limit <= 0 {
		limit = _offset
	}
	query.Add("limit", strconv.Itoa(int(limit)))

	if offset < 0 {
		offset = _offset
	}
	query.Add("offset", strconv.Itoa(int(offset)))
	if len(query) > 0 {
		link += "?" + query.Encode()
	}

	reqBody, err := json.Marshal(condition)
	if err != nil {
		return nil, fmt.Errorf("序列化请求参数失败: %w", err)
	}
	
	resp := reqCli.Visit("POST", link, reqHeaders, string(reqBody), apiTimeout, "", nil)
	if resp.Error != nil {
		// 搜索要在营销活动中推广的 iOS 应用程序
		return nil, fmt.Errorf("获取地理位置详细信息失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &entitiesResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return entitiesResp, nil
}