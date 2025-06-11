package feishu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Webglhost-QA-Backend/backend/config"
	"github.com/Webglhost-QA-Backend/backend/internal/app/models"
	"io"
	"log"
	"net/http"
	"time"
)

type FeishuClient struct {
	config     config.FeishuConfig
	apiURL     string
	userToken  string
	GameList   []models.Game
	httpClient *http.Client
}

func NewFeishuClient(config config.FeishuConfig) (*FeishuClient, error) {
	if config.FORMAPI == "" || config.SHEETTOKEN == "" || config.TAPSHEETID == "" {
		return nil, fmt.Errorf("飞书配置不完整")
	}

	apiURL := fmt.Sprintf("%s/%s/values/%s!%s", config.FORMAPI, config.SHEETTOKEN, config.TAPSHEETID, config.RANGE)

	return &FeishuClient{
		config: config,
		apiURL: apiURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				ResponseHeaderTimeout: 5 * time.Second,
			},
		},
	}, nil
}

func (f *FeishuClient) PostRequest(url string, headers map[string]string, data interface{}) (map[string]interface{}, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("JSON序列化失败: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := f.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求执行失败: %w", err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			log.Printf("关闭响应体失败: %v", closeErr)
		}
	}()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API返回错误状态码: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("响应解析失败: %w", err)
	}

	return result, nil
}

func (f *FeishuClient) GetRequest(url string, headers map[string]string) (map[string]interface{}, error) {
	// 创建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建GET请求失败: %w", err)
	}

	// 设置请求头
	req.Header.Set("Accept", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// 发送请求
	resp, err := f.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GET请求执行失败: %w", err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			log.Printf("关闭响应体失败: %v", closeErr)
		}
	}()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应体失败: %w", err)
	}

	// 检查状态码
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API返回错误状态码: %d, 响应体: %s", resp.StatusCode, string(body))
	}

	// 解析JSON响应
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("响应解析失败: %w, 原始响应: %s", err, string(body))
	}

	return result, nil
}

func (f *FeishuClient) GetUserToken() error {
	requestData := map[string]string{
		"app_id":     f.config.APPID,
		"app_secret": f.config.APPSECRET,
	}

	headers := map[string]string{
		"Accept": "application/json",
	}
	response, err := f.PostRequest(f.config.USERTOKENAPI, headers, requestData)
	if err != nil {
		return fmt.Errorf("fail to get user token: %w", err)
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("fail marshal response data: %w", err)
	}

	var resultMap map[string]interface{}
	if err := json.Unmarshal(jsonData, &resultMap); err != nil {
		return fmt.Errorf("JSON parsing faild: %w", err)
	}
	if resultMap["msg"] == "ok" {
		token := resultMap["app_access_token"].(string)
		log.Printf("successful get user token %s", token)
		f.userToken = token
	}

	log.Printf("JSON response data is : %v\n", resultMap)
	return nil
}

func (f *FeishuClient) GetFeishuSheetData() error {
	headers := map[string]string{
		"Authorization": "Bearer " + f.userToken,
		"Content-Type":  "application/json",
	}
	log.Println(f.apiURL)
	log.Println(headers)
	response, err := f.GetRequest(f.apiURL, headers)
	if err != nil {
		log.Printf("fail to get feishu sheet data %v", err)
		return err
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("fail marshal response data: %w", err)
	}
	var resultMap map[string]interface{}
	if err := json.Unmarshal(jsonData, &resultMap); err != nil {
		return fmt.Errorf("JSON parsing faild: %w", err)
	}
	data := resultMap["data"].(map[string]interface{})
	valueRange := data["valueRange"].(map[string]interface{})
	values := valueRange["values"].([]interface{})

	for i, value := range values {
		check := value.([]interface{})
		if check[0] == nil || check[1] == nil || check[3] == nil || check[16] == nil {
			continue
		}
		urlSlice, ok := check[4].([]interface{})
		if !ok || len(urlSlice) == 0 {
			continue
		}

		urlMap, ok2 := urlSlice[0].(map[string]interface{})

		if !ok2 {
			continue
		}

		typeVal, typeExists := urlMap["type"]
		if typeExists && typeVal != "url" {
			_, ok = typeVal.(string)
			if !ok {
				continue
			}
		}

		textVal, textExists := urlMap["text"]
		if textExists && textVal != nil {
			_, ok = textVal.(string)
			if !ok {
				continue
			}
		} else {
			continue
		}

		Package := "com.u3d.webglhost"
		Type := "unity 引擎小游戏"
		gameEngine := check[3].(string)
		gameUrl := urlMap["text"].(string)
		gameName := check[0].(string)
		gameType := check[1].(string)
		CaseType := make([]string, 2)
		CaseType[0] = check[16].(string)
		gameId := i
		gameRecord := models.Game{
			Package:    Package,
			Type:       Type,
			GameEngine: gameEngine,
			GameUrl:    gameUrl,
			GameName:   gameName,
			GameType:   gameType,
			GameId:     gameId,
			CaseType:   CaseType,
		}
		f.GameList = append(f.GameList, gameRecord)
	}

	return nil
}
