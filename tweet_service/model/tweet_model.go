package model

type AddTweetRequest struct {
	UserId       int64  `json:"user_id"`
	TweetContent string `json:"tweet_content"`
}
type AddTweetResponse struct {
	TweetId      int64  `json:"tweet_id"`
	UserId       int64  `json:"user_id"`
	TweetContent string `json:"tweet_content"`
}
