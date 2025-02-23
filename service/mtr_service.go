package service

import (
	"fmt"
	"io"
	"net/http"
)

const (
	MTR_API_BASE_URL = "https://rt.data.gov.hk/v1/transport/mtr/getschedule.php"
)

type MTRResponse struct {
	ResultCode int       `json:"resultCode"`
	Timestamp  string    `json:"timestamp"`
	Status     int       `json:"status"`
	Message    string    `json:"message"`
	Error      MTRError  `json:"error,omitempty"`
}

type MTRError struct {
	ErrorCode string `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
}

type MTRService struct {}

func NewMTRService() *MTRService {
	return &MTRService{}
}

func (s *MTRService) HandleMTRSchedule(w http.ResponseWriter, r *http.Request) {
	// 获取 line 和 sta 参数
	line := r.URL.Query().Get("line")
	sta := r.URL.Query().Get("sta")
	
	if line == "" || sta == "" {
		http.Error(w, "缺少必要参数 line 或 sta", http.StatusBadRequest)
		return
	}

	// 构建 MTR API 请求
	mtrURL := fmt.Sprintf("%s?line=%s&sta=%s", MTR_API_BASE_URL, line, sta)
	
	// 发送请求到 MTR API
	resp, err := http.Get(mtrURL)
	if err != nil {
		http.Error(w, "请求 MTR API 失败", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "读取响应失败", http.StatusInternalServerError)
		return
	}

	// 设置响应头
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	
	// 返回响应内容
	w.Write(body)
} 