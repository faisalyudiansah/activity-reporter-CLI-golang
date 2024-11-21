package variable

import (
	"errors"
	"fmt"
)

var (
	ErrorInvalidKeyword       = errors.New("invalid keyword")
	ErrorInvalidMenu          = errors.New("invalid menu")
	ErrorAlreadyFollow        = errors.New("you already followed the user")
	ErrorUserFollowThemselves = errors.New("a user cannot follow themselves")
	ErrorAlreadyUploadPhoto   = errors.New("you cannot upload more than once")
	ErrorAlreadyLIkePhoto     = errors.New("you already liked the photo")
	ErrorUserDoesNotHavePhoto = errors.New("you don't have a photo")
)

var (
	ErrorAnotherUserDoesNotHavePhoto = errors.New("{inputted-user-name} doesn't have a photo")
	ErrorUnknownUserForTesting       = errors.New("unknown user Faisal")
)

func SetErrorUnkownUser(nameOfUser string) error {
	errorMsg := fmt.Sprintf("unknown user %v", nameOfUser)
	return errors.New(errorMsg)
}

func SetErrorDoesNotHavePhoto(nameOfUser string) error {
	errorMsg := fmt.Sprintf("%v doesn't have a photo", nameOfUser)
	return errors.New(errorMsg)
}

func SetErrorUnableLikePhoto(nameOfUser string) error {
	errorMsg := fmt.Sprintf("unable to like %v's photo", nameOfUser)
	return errors.New(errorMsg)
}
