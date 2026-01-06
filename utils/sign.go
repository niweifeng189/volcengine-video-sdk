package utils

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/base64"
    "fmt"
    "net/url"
    "sort"
    "strings"
    "time"
    
    "volcengine-video-sdk/config"
)

// Sign 生成火山引擎API签名
func Sign(cfg *config.Config, params url.Values) (string, error) {
    // 1. 添加公共参数
    params.Set("AccessKey", cfg.AccessKey)
    params.Set("Region", cfg.Region)
    params.Set("Timestamp", fmt.Sprintf("%d", time.Now().Unix()))
    params.Set("Nonce", fmt.Sprintf("%d", time.Now().UnixNano()))

    // 2. 排序参数（火山签名要求）
    keys := make([]string, 0, len(params))
    for k := range params {
        keys = append(keys, k)
    }
    sort.Strings(keys)

    // 3. 拼接参数字符串
    var paramStr strings.Builder
    for _, k := range keys {
        paramStr.WriteString(k)
        paramStr.WriteString(params.Get(k))
    }

    // 4. HMAC-SHA256签名
    h := hmac.New(sha256.New, []byte(cfg.SecretKey))
    h.Write([]byte(paramStr.String()))
    signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

    // 5. 拼接最终请求URL
    u, err := url.Parse(cfg.Endpoint)
    if err != nil {
        return "", err
    }
    u.RawQuery = params.Encode() + "&Signature=" + url.QueryEscape(signature)

    return u.String(), nil
}
