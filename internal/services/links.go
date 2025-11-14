package services

import "status-links/internal/models"

type LinksService struct {
}

func NewLinksService() *LinksService {
	return &LinksService{}
}

func (l *LinksService) UploadAllUnfinishedWork() *models.AllUnfinishedWork {
	return nil
}
func (l *LinksService) UploadAll(set models.SetLinksGet) *models.ProcessedLinks {
	return nil
}

func (l *LinksService) AddLinkSet(set models.SetLinksGet) *models.ProcessedLinks {
	return nil
}

func (l *LinksService) GiveLinkAnswer(list []models.SetNumsOfLinksGet) *models.ListOfProcessedLinks {
	return nil
}
