package storage

import "status-links/internal/models"

type tempStorage struct {
	sets map[int]models.LinksAnswer
}

func NewTempStorage() *tempStorage {
	return &tempStorage{
		sets: make(map[int]models.LinksAnswer, 0),
	}
}
