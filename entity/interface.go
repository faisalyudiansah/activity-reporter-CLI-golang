package entity

type Publisher interface {
	RegisteredUser(Observer)
	NotifyForLike(*User)
	NotifyForUpload()
}

type Observer interface {
	UpdateLike(*User, *User)
	UpdateUpload(*User)
}
