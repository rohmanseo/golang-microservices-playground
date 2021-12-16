package use_case

import "github.com/rohmanseo/golang-clean-arch/model"

type ITweetUseCase interface {
	AddTweet(request model.AddTweetRequest) (model.AddTweetResponse, error)
}
