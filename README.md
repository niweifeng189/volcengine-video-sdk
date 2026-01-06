# 火山引擎视频发布SDK

一个用于火山引擎视频服务的Go SDK，封装了视频发布等核心功能。

## 项目结构

```
volcengine-video-sdk/
├── client/          // 客户端核心（初始化、请求封装）
│   └── client.go
├── api/             // 发布相关API接口
│   └── publish.go
├── config/          // 配置项（AK/SK、域名等）
│   └── config.go
├── utils/           // 工具函数（签名、JSON处理）
│   └── sign.go
├── example/         // 使用示例
│   └── main.go
├── go.mod           // 模块声明
└── README.md        // 使用文档
```

## 功能特性

- 支持火山引擎视频发布API
- 自动生成API签名
- 支持自定义服务域名
- 支持多种地域配置
- 完整的错误处理

## 快速开始

### 安装

```bash
go get github.com/yourusername/volcengine-video-sdk
```

### 使用示例

```go
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
```

## 配置说明

| 配置项 | 类型 | 说明 |
|--------|------|------|
| AccessKey | string | 火山引擎访问密钥AK |
| SecretKey | string | 火山引擎访问密钥SK |
| Endpoint | string | 火山引擎视频服务域名，如 https://vod.volcengineapi.com |
| Region | string | 地域，如 cn-north-1 |

## API说明

### PublishVideo

发布视频到火山引擎视频服务。

**请求参数**：

| 参数名 | 类型 | 说明 |
|--------|------|------|
| VideoID | string | 上传后的视频ID |
| Title | string | 视频标题 |
| Status | int | 发布状态：1-发布，0-草稿 |
| CategoryID | int | 分类ID |

**响应参数**：

| 参数名 | 类型 | 说明 |
|--------|------|------|
| ResponseMetadata.RequestID | string | 请求ID |
| ResponseMetadata.Error | struct | 错误信息，成功时为nil |
| Result.PublishID | string | 发布任务ID |

## 错误处理

SDK会返回详细的错误信息，包括错误码和错误描述，可以通过解析返回的错误对象获取。

## 版本历史

- v1.0.0: 初始版本，支持视频发布功能

## 许可证

MIT
