package repositories

type Repo interface {
	PushMessage(msg string) error
}
