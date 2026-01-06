package main

import (
    "fmt"
    
    "volcengine-video-sdk/api"
    "volcengine-video-sdk/client"
    "volcengine-video-sdk/config"
)

func main() {
    // 1. 配置AK/SK和域名
    cfg := &config.Config{
        AccessKey: "你的火山AK",
        SecretKey: "你的火山SK",
        Endpoint:  "火山引擎视频服务域名",
        Region:    "cn-north-1",
    }

    // 2. 初始化客户端
    cli := client.NewVideoClient(cfg)

    // 3. 发布视频
    req := &api.PublishVideoRequest{
        VideoID:    "上传后的视频ID",
        Title:      "测试视频",
        Status:     1, // 发布
        CategoryID: 100,
    }
    resp, err := cli.PublishVideo(req)
    if err != nil {
        fmt.Printf("publish failed: %v\n", err)
        return
    }
    fmt.Printf("publish success, PublishID: %s\n", resp.Result.PublishID)
}
