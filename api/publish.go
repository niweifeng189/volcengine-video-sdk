package api

// PublishVideoRequest 发布视频的请求参数
type PublishVideoRequest struct {
    VideoID    string `json:"VideoId"`    // 上传后的视频ID
    Title      string `json:"Title"`      // 视频标题
    Status     int    `json:"Status"`     // 发布状态：1-发布，0-草稿
    CategoryID int    `json:"CategoryId"` // 分类ID
}

// PublishVideoResponse 发布视频的响应结果
type PublishVideoResponse struct {
    ResponseMetadata struct {
        RequestID string `json:"RequestId"`
        Error     *struct {
            Code    string `json:"Code"`
            Message string `json:"Message"`
        } `json:"Error"`
    } `json:"ResponseMetadata"`
    Result struct {
        PublishID string `json:"PublishId"`
    } `json:"Result"`
}
