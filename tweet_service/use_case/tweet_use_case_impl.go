package use_case

import (
	"github.com/rohmanseo/golang-clean-arch/model"
	"github.com/rohmanseo/golang-clean-arch/repository"
)

type tweetUseCaseImpl struct {
	tweetRepository repository.ITweetRepository
}

func (t tweetUseCaseImpl) AddTweet(request model.AddTweetRequest) (model.AddTweetResponse, error) {
	res, err := t.tweetRepository.AddTweet(request.UserId, request.TweetContent)

	if err != nil {
		return model.AddTweetResponse{}, err
	}
	return model.AddTweetResponse{
		TweetId:      res.Id,
		UserId:       res.UserId,
		TweetContent: res.TweetContent,
	}, nil
}

func NewTweetUseCaseImpl(tweetRepo *repository.ITweetRepository) ITweetUseCase {
	return &tweetUseCaseImpl{
		tweetRepository: *tweetRepo,
	}
}
