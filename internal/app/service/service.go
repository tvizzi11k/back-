package service

import (
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) DaysLeft() int64 {
	data := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)

	dur := time.Until(data)

	return int64(int64(dur.Hours()) / 24)
}
