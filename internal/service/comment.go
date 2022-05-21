package service

import "github.com/FlyInThesky10/TikTok-Fly/internal/model"

type CommentRequest struct {
	UserId      uint32 `form:"user_id" binding:"required"`
	Token       string `form:"token" binding:"required"`
	VideoId     uint32 `form:"video_id" binding:"required"`
	ActionType  uint8  `form:"action_type" binding:"required"`
	CommentText string `form:"comment_text"`
	CommentId   uint32 `form:"comment_id"`
}

type CommentListRequest struct {
	UserId  uint32 `form:"user_id" binding:"required"`
	Token   string `form:"token" binding:"required"`
	VideoId uint32 `form:"video_id" binding:"required"`
}

func (svc *Service) CreateComment(param *CommentRequest) error {
	return nil
}

func (svc *Service) GetCommentList(param *CommentListRequest) ([]*model.Comment, error) {
	return nil, nil
}
