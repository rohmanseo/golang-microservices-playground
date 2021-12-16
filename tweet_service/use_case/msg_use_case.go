package use_case

type IMsgUseCase interface {
	TestMessageReceived(message string)
	GetTotalTweet() int
}
