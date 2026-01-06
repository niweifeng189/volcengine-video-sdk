package config

// Config 火山引擎视频服务配置
type Config struct {
    AccessKey  string // 访问密钥AK
    SecretKey  string // 访问密钥SK
    Endpoint   string // 火山引擎视频服务域名（可配置）
    Region     string // 地域，如cn-north-1
}
