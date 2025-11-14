package storage

type reliableStorage struct {
	nameFile string
}

func NewReliableStorage(fileName string) *reliableStorage {
	return &reliableStorage{
		nameFile: fileName,
	}
}
