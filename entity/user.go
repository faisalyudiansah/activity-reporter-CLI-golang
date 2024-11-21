package entity

import "fmt"

type User struct {
	Name         string
	Following    []*User
	Followers    []*User
	Photo        *Photo
	observerList []Observer
	Activity     []string
}

func NewUser(name string) *User {
	return &User{
		Name:         name,
		Following:    []*User{},
		Followers:    []*User{},
		Photo:        nil,
		observerList: []Observer{},
		Activity:     []string{},
	}
}

func (u *User) UploadPhotoToUser(newUpload *Photo) {
	u.Photo = newUpload
}

func (u *User) IsFollowers(user *User) bool {
	if u.Name == user.Name {
		return true
	}
	for _, v := range u.Followers {
		if v == user {
			return true
		}
	}
	return false
}

func (u *User) RegisteredUser(observer Observer) {
	u.observerList = append(u.observerList, observer)
}

func (u *User) UpdateLike(userWantToLike, userTarget *User) {
	var userName string = fmt.Sprintf("%v's", userTarget.Name)
	if u.Name == userTarget.Name {
		userName = "your"
	}
	u.Activity = append(u.Activity, fmt.Sprintf("%v liked %v photo", userWantToLike.Name, userName))
}

func (u *User) UpdateUpload(user *User) {
	u.Activity = append(u.Activity, fmt.Sprintf("%v uploaded photo", user.Name))
}

func (u *User) NotifyForUpload() {
	for _, observer := range u.observerList {
		observer.UpdateUpload(u)
	}
}

func (u *User) NotifyForLike(targetUser *User) {
	for _, observer := range u.observerList {
		observer.UpdateLike(u, targetUser)
	}
}
