package storage

type Storage interface {
	Push(localFileName, remoteFileName string) error
}
