// Code generated by hertz generator.

package ApiServer

import (
	"bytes"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cold-runner/simpleTikTok/apiServer/biz/handler/response"
	"github.com/cold-runner/simpleTikTok/apiServer/biz/model/ApiServer"
	"github.com/cold-runner/simpleTikTok/apiServer/rpc"
	"github.com/cold-runner/simpleTikTok/kitex_gen/VideoService"
	"github.com/cold-runner/simpleTikTok/pkg/errno"
	"github.com/cold-runner/simpleTikTok/pkg/log"
	mw "github.com/cold-runner/simpleTikTok/pkg/middleware"
	"io"
	"strconv"
)

// Feed .
// @router /douyin/feed/ [GET]
func Feed(ctx context.Context, c *app.RequestContext) {
	var err error
	var uid int64
	latestTimeString := c.Query("latest_time")
	latestTime, err := strconv.Atoi(latestTimeString)
	v, _ := c.Get(mw.IdentityKey)
	if v == nil {
		uid = 0
	} else {
		uid = v.(*ApiServer.User).Id
	}
	videoFeedResp, err := rpc.VideoFeed(context.Background(),
		&VideoService.VideoFeedRequest{
			LatestTime: int64(latestTime),
			UserId:     uid,
		})
	if err != nil {
		log.Errorw("rpc get video feed failed", "err", err)
	}

	response.SendFeedResponse(c, videoFeedResp.VideoList, videoFeedResp.NextTime,
		errno.OK)
}

// PublishAction .
// @router /douyin/publish/action/ [POST]
func PublishAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ApiServer.PublishActionRequest
	log.Debugw("start publish action")
	// 获取请求参数
	req.Title = c.PostForm("title")
	req.Token = c.PostForm("token")
	fileHeader, err := c.Request.FormFile("data")
	if err != nil {
		log.Errorw("get file failed", "err", err)
		response.SendPublishActionResponse(c, err)
		return
	}
	// 打开文件
	file, err := fileHeader.Open()
	if err != nil {
		log.Errorw("open file failed", "err", err)
		response.SendPublishActionResponse(c, err)
		return
	}
	defer file.Close()

	// 读取文件
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, file)
	if err != nil {
		log.Errorw("read file failed", "err", err)
		response.SendPublishActionResponse(c, err)
	}
	v, _ := c.Get(mw.IdentityKey)
	request := &VideoService.VideoPublishActionRequest{
		UserId: v.(*ApiServer.User).Id,
		Data:   buf.Bytes(),
		Title:  req.Title,
	}
	resp, err := rpc.PublishVideo(context.Background(), request)
	if err != nil {
		log.Errorw("rpc publish video failed", "err", err)
		response.SendPublishActionResponse(c, err)
		return
	}
	log.Debugw("rpc publish video success", "resp", resp)
	response.SendPublishActionResponse(c, errno.OK)
}

// PublishList .
// @router /douyin/publish/list/ [GET]
func PublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req ApiServer.PublishListRequest
	v, _ := c.Get(mw.IdentityKey)
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	// 获取请求参数
	videoPublishListResp, err := rpc.VideoPublishList(context.Background(),
		&VideoService.VideoPublishListRequest{
			UserId:   v.(*ApiServer.User).Id,
			ToUserId: req.UserId,
		})
	if err != nil {
		log.Errorw("rpc get video publish list failed", "err", err)
		response.SendVideoPublishListResponse(c, nil, err)
		return
	}

	response.SendVideoPublishListResponse(c, videoPublishListResp.VideoList,
		errno.OK)
}
