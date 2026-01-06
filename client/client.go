package client

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
    "time"
    
    "github.com/niweifeng189/volcengine-video-sdk/api"
    "github.com/niweifeng189/volcengine-video-sdk/config"
    "github.com/niweifeng189/volcengine-video-sdk/utils"
)

// VideoClient 火山引擎视频服务客户端
type VideoClient struct {
    config  *config.Config
    httpCli *http.Client
}

// NewVideoClient 初始化视频服务客户端
func NewVideoClient(cfg *config.Config) *VideoClient {
    return &VideoClient{
        config: cfg,
        httpCli: &http.Client{
            Timeout: 30 * time.Second,
        },
    }
}

// PublishVideo 发布视频方法
func (c *VideoClient) PublishVideo(req *api.PublishVideoRequest) (*api.PublishVideoResponse, error) {
    // 1. 组装请求参数
    params := url.Values{}
    params.Set("Action", "PublishVideo")
    params.Set("Version", "2020-08-01") // 火山引擎API版本
    params.Set("VideoId", req.VideoID)
    params.Set("Title", req.Title)
    params.Set("Status", fmt.Sprintf("%d", req.Status))
    params.Set("CategoryId", fmt.Sprintf("%d", req.CategoryID))

    // 2. 生成火山引擎API签名
    signedURL, err := utils.Sign(c.config, params)
    if err != nil {
        return nil, fmt.Errorf("sign failed: %v", err)
    }

    // 3. 发送POST请求
    resp, err := c.httpCli.Post(signedURL, "application/x-www-form-urlencoded", bytes.NewBufferString(params.Encode()))
    if err != nil {
        return nil, fmt.Errorf("request failed: %v", err)
    }
    defer resp.Body.Close()

    // 4. 解析响应
    var publishResp api.PublishVideoResponse
    if err := json.NewDecoder(resp.Body).Decode(&publishResp); err != nil {
        return nil, fmt.Errorf("decode response failed: %v", err)
    }

    // 5. 处理错误
    if publishResp.ResponseMetadata.Error != nil {
        return nil, fmt.Errorf("api error: %s - %s",
            publishResp.ResponseMetadata.Error.Code,
            publishResp.ResponseMetadata.Error.Message)
    }

    return &publishResp, nil
}
