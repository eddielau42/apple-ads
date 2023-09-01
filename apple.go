package appleads

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/wangyong321/gogorequest"
)

const (
	audienceURL = "https://appleid.apple.com"
	oauth2URL = "https://appleid.apple.com/auth/oauth2/token"

	apiURL = "https://api.searchads.apple.com/api"
	ver = "v4"

	apiTimeout = 10
)

type(
	Engine struct {
		clientID string
		teamID string
		keyID string
		privateKEY string

		accessToken *AccessToken

		jwtToken string

		orgID int64
	}

	EngineOption func(engien *Engine)

	// AccessToken
	AccessToken struct{
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int64  `json:"expires_in"`
	}
)

type (
	SelectorCondition struct {
		Field string `json:"field"`
		Operator string `json:"operator"`
		Values []string `json:"values"`
	}
	Pagination struct {
		Limit int32 `json:"limit"`
		Offset int32 `json:"offset"`
	}
	Sorting struct {
		Field string `json:"field"`
		SortOrder string `json:"sortOrder"`
	}
)

// New 获取并返回实例
func New(options ...EngineOption) *Engine {
	engine := new(Engine)
	for _, fn := range options {
		fn(engine)
	}
	return engine
}

// WithClientID
func WithClientID(val string) EngineOption {
	return func (engine *Engine)  {
		engine.clientID = val
	}
}
// WithTeamID
func WithTeamID(val string) EngineOption {
	return func (engine *Engine)  {
		engine.teamID = val
	}
}
// WithKeyID
func WithKeyID(val string) EngineOption {
	return func (engine *Engine)  {
		engine.keyID = val
	}
}
// PrivateKey
func WithPrivateKey(val string) EngineOption {
	return func (engine *Engine)  {
		engine.privateKEY = val
	}
}

// JwtTTL JWT有效时长
func JwtTTL() int64 {
	return 86400 * 180
}

// ApiBaseURL
func ApiBaseURL() string {
	return apiURL + "/" + ver
}

// JWT 使用公钥、私钥、keyid、teamid、clientid生成本地JWT token
func (engine *Engine) JWT() (string, error) {
	var (
		token string
		err error
	)

	issueAt := time.Now().Unix()
	expiresAt := issueAt + JwtTTL()

	c := jwt.StandardClaims{
		Subject: engine.clientID,
		Audience: audienceURL,
		IssuedAt: issueAt,
		ExpiresAt: expiresAt,
		Issuer: engine.teamID,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodES256, c)
	t.Header["kid"] = engine.keyID
	privateKey, err := jwt.ParseECPrivateKeyFromPEM([]byte(engine.privateKEY))
	if err != nil {
		return "", err
	}

	token, err = t.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

// AccessToken 携带 JWT token 获取苹果的 access_token
func (engine *Engine) AccessToken(jwtToken string) (*AccessToken, error) {
	var (
		err error
		authResp *AuthorizationResponse
	)

	const timeout = 20

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	reqBody := fmt.Sprintf("grant_type=client_credentials&client_id=%v&client_secret=%v&scope=searchadsorg", engine.clientID, jwtToken)
	
	resp := reqCli.Visit("POST", oauth2URL, reqHeaders, reqBody, timeout, "", nil)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("access_token请求失败, 响应结果: %w", resp.Error)
	}
	
	err = json.Unmarshal([]byte(resp.Text), &authResp)
	if err != nil {
		return nil, fmt.Errorf("解析access_token响应数据失败: %w", err)
	}
	return &authResp.AccessToken, nil
}

// Auth 授权
func (engine *Engine) Auth() error {
	// var err error
	if engine.accessToken == nil || engine.accessToken.AccessToken == "" {
		if engine.jwtToken == "" {
			jwtStr, err := engine.JWT()
			if err != nil {
				return fmt.Errorf("engine.Auth->JWT_error: %w", err)
			}
			engine.SetJwt(jwtStr)
		}

		ascTkn, err := engine.AccessToken(engine.jwtToken)
		if err != nil {
			return fmt.Errorf("engine.Auth->AccessToken_error: %w", err)
		}
		engine.SetAccessToken(ascTkn)
	}

	return nil
}

// SetAccessToken 设置 access-token
func (engine *Engine) SetAccessToken(accessToken *AccessToken) {
	engine.accessToken = accessToken
}

// SetJwt 设置 jwt-token
func (engine *Engine) SetJwt(jwtToken string) {
	engine.jwtToken = jwtToken
}

// SetOrgID 设置 orgID
func (engine *Engine) SetOrgID(orgID int64) {
	engine.orgID = orgID
}

// AuthInfo
func (engine *Engine) AuthInfo() map[string]string {
	return map[string]string{
		"jwt": engine.jwtToken,
		"accessToken": engine.accessToken.AccessToken,
		"accessTokenType": engine.accessToken.TokenType,
		"accessTokenExpiresIn": strconv.FormatInt(engine.accessToken.ExpiresIn, 10),
	}
}

// UserAcl 获取API可访问的角色和组织
func (engine *Engine) UserAcl() (*UserAclListResponse, error) {
	var (
		err error
		aclResp *UserAclListResponse
	)

	link := ApiBaseURL() + "/acls"

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
	}

	resp := reqCli.Visit("GET", link, reqHeaders, nil, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("获取API可访问的角色和组织失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &aclResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return aclResp, nil
}

// Me 获取API调用者的详细信息
func (engine *Engine) Me() (*MeDetailResponse, error) {
	var (
		err error
		meResp *MeDetailResponse
	)

	link := ApiBaseURL() + "/me"

	reqCli := gogorequest.NewSyncEngine()
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
		"Authorization": engine.accessToken.TokenType + " " + engine.accessToken.AccessToken,
	}

	resp := reqCli.Visit("GET", link, reqHeaders, nil, apiTimeout, "", nil)
	if resp.Error != nil {
		return nil, fmt.Errorf("请求广告数据失败, 响应结果: %w", resp.Error)
	}

	err = json.Unmarshal([]byte(resp.Text), &meResp)
	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	return meResp, nil
}