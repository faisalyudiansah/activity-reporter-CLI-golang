package entity

import (
	"time"

	"activity-reporter-cli/variable"
)

type Photo struct {
	Like           []*User
	TimeOfLastLike time.Time
}

func NewUploadPhoto() *Photo {
	return &Photo{
		Like:           []*User{},
		TimeOfLastLike: time.Now(),
	}
}

func (p *Photo) AddLike(user *User) error {
	for _, v := range p.Like {
		if v == user {
			return variable.ErrorAlreadyLIkePhoto
		}
	}
	p.UpdateTimeOfLastLike()
	p.Like = append(p.Like, user)
	return nil
}

func (p *Photo) UpdateTimeOfLastLike() {
	p.TimeOfLastLike = time.Now()
}
