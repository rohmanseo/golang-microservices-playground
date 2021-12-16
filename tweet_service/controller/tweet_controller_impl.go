package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rohmanseo/golang-clean-arch/model"
	"github.com/rohmanseo/golang-clean-arch/security"
	"github.com/rohmanseo/golang-clean-arch/use_case"
	"net/http"
	"strconv"
	"strings"
)

func NewTweetController(tweetUseCase *use_case.ITweetUseCase, jwtmgr security.IJwtToken) ITweetController {
	return &tweetControllerImpl{
		TweetUseCase: *tweetUseCase,
		jwtMgr:       jwtmgr,
	}
}

//TODO: input validation, req&res mapper, response helper

type tweetControllerImpl struct {
	TweetUseCase use_case.ITweetUseCase
	jwtMgr       security.IJwtToken
}

func (t *tweetControllerImpl) Add(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	token = strings.Split(token, " ")[1]
	userId, _ := t.jwtMgr.GetValue(token, "user_id")
	id, _ := strconv.ParseInt(userId, 10, 64)

	res, err := t.TweetUseCase.AddTweet(model.AddTweetRequest{
		UserId:       id,
		TweetContent: ctx.PostForm("tweet_content"),
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ApiResponse{
			Code:   http.StatusBadRequest,
			Status: "Error",
			Data:   nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, model.ApiResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	})
	return
}
