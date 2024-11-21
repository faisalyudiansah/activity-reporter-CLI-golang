package controller

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"activity-reporter-cli/entity"
	"activity-reporter-cli/variable"
)

type Activity struct {
	userRegistered []*entity.User
}

func NewSocialGraph() *Activity {
	return &Activity{
		userRegistered: []*entity.User{},
	}
}

func (a *Activity) FollowUser(userA string, userB string) error {
	if err := a.isValidInputUser(userA, userB); err != nil {
		return err
	}
	if userA == userB {
		return variable.ErrorUserFollowThemselves
	}
	userWhoWillFollow := a.validationOrAddUser(userA)
	userToFollow := a.validationOrAddUser(userB)
	if a.isAlreadyFollow(userWhoWillFollow, userToFollow) {
		return variable.ErrorAlreadyFollow
	}
	a.setFollowingAndFollower(userWhoWillFollow, userToFollow)
	return nil
}

func (a *Activity) setFollowingAndFollower(userWhoWillFollow *entity.User, userToFollow *entity.User) {
	userWhoWillFollow.Following = append(userWhoWillFollow.Following, userToFollow)
	userToFollow.Followers = append(userToFollow.Followers, userWhoWillFollow)
	if !a.isListedRegister(userWhoWillFollow) {
		a.setListUserRegistered(userWhoWillFollow)
	}
	if !a.isListedRegister(userToFollow) {
		a.setListUserRegistered(userToFollow)
	}
	userToFollow.RegisteredUser(userWhoWillFollow)
}

func (a *Activity) setListUserRegistered(user *entity.User) {
	a.userRegistered = append(a.userRegistered, user)
}

func (a *Activity) isListedRegister(user *entity.User) bool {
	for _, v := range a.userRegistered {
		if v == user {
			return true
		}
	}
	return false
}

func (a *Activity) isAlreadyFollow(userWhoWillFollow, userToFollow *entity.User) bool {
	for _, v := range userWhoWillFollow.Following {
		if v.Name == userToFollow.Name {
			return true
		}
	}
	return false
}

func (a *Activity) getDataUser(nameOfUser string) (bool, *entity.User) {
	for _, v := range a.userRegistered {
		if v.Name == nameOfUser {
			return true, v
		}
	}
	return false, nil
}

func (a *Activity) validationOrAddUser(nameOfUser string) *entity.User {
	isExist, getUser := a.getDataUser(nameOfUser)
	if isExist {
		return getUser
	}
	newUser := entity.NewUser(nameOfUser)
	return newUser
}

func (a *Activity) isNameFromInputValid(nameOfUser string) bool {
	nameOfUser = strings.Trim(nameOfUser, " ")
	return len(nameOfUser) < 1
}

func (a *Activity) isThereHasNumber(nameOfUser string) bool {
	for _, v := range nameOfUser {
		if a.checkNumber(string(v)) {
			return true
		}
	}
	return false
}

func (a *Activity) checkNumber(nameOfUser string) bool {
	_, err := strconv.Atoi(nameOfUser)
	return err == nil
}

//=====================================================================================================

func (a *Activity) UploadPhoto(nameOfUser string) error {
	_, user := a.getDataUser(nameOfUser)
	if user == nil {
		return variable.SetErrorUnkownUser(nameOfUser)
	}
	if user.Photo != nil {
		return variable.ErrorAlreadyUploadPhoto
	}
	newUpload := entity.NewUploadPhoto()
	user.UploadPhotoToUser(newUpload)
	user.Activity = append(user.Activity, "You uploaded photo")
	user.NotifyForUpload()
	return nil
}

//=====================================================================================================

func (a *Activity) LikePhoto(userA string, userB string) error {
	if err := a.isValidInputUser(userA, userB); err != nil {
		return err
	}
	if err := a.isUserUnknown(userA, userB); err != nil {
		return err
	}
	_, getUserA := a.getDataUser(userA)
	_, getUserB := a.getDataUser(userB)
	if !getUserB.IsFollowers(getUserA) {
		return variable.SetErrorUnableLikePhoto(userB)
	}
	if err := a.hasPhoto(getUserB, userA, userB); err != nil {
		return err
	}
	if err := getUserB.Photo.AddLike(getUserA); err != nil {
		return err
	}
	a.setActivityNotify(userA, userB, getUserA, getUserB)
	return nil
}

func (a *Activity) setActivityNotify(userA string, userB string, getUserA *entity.User, getUserB *entity.User) {
	var userName string = fmt.Sprintf("%v's", getUserB.Name)
	if userA == userB {
		userName = "your"
	}
	getUserA.Activity = append(getUserA.Activity, fmt.Sprintf("You liked %v photo", userName))
	if !getUserA.IsFollowers(getUserB) {
		getUserB.Activity = append(getUserB.Activity, fmt.Sprintf("%v liked your photo", getUserA.Name))
	}
	getUserA.NotifyForLike(getUserB)
}

func (a *Activity) hasPhoto(getUserB *entity.User, userA string, userB string) error {
	if getUserB.Photo == nil && userA == userB {
		return variable.ErrorUserDoesNotHavePhoto
	}
	if getUserB.Photo == nil {
		return variable.SetErrorDoesNotHavePhoto(userB)
	}
	return nil
}

func (a *Activity) isUserUnknown(userA, userB string) error {
	okUserA, _ := a.getDataUser(userA)
	okUserB, _ := a.getDataUser(userB)
	if !okUserA {
		return variable.SetErrorUnkownUser(userA)
	}
	if !okUserB {
		return variable.SetErrorUnkownUser(userB)
	}
	return nil
}

func (a *Activity) isValidInputUser(userA, userB string) error {
	if a.isNameFromInputValid(userA) || a.isNameFromInputValid(userB) || a.isThereHasNumber(userA) || a.isThereHasNumber(userB) {
		return variable.ErrorInvalidKeyword
	}
	return nil
}

//=====================================================================================================

func (a *Activity) ActivityUser(nameOfUser string) ([]string, error) {
	_, user := a.getDataUser(nameOfUser)
	if user == nil {
		return nil, variable.SetErrorUnkownUser(nameOfUser)
	}
	for _, v := range a.userRegistered {
		if v.Name == nameOfUser {
			return v.Activity, nil
		}
	}
	return nil, variable.SetErrorUnkownUser(nameOfUser)
}

//=====================================================================================================

func (a *Activity) TrendingPhotos() []string {
	a.sortingForPhotoUsers()
	var result []string
	for idx, v := range a.userRegistered {
		if v.Photo != nil && len(v.Photo.Like) > 0 {
			result = append(result, a.setInfoTextTrending(v, idx))
		}
	}
	if len(a.selectionListTrending(result)) <= 0 {
		return result
	}
	return a.selectionListTrending(result)
}

func (a *Activity) selectionListTrending(listTopTrending []string) []string {
	cutIdx := 3
	if len(listTopTrending) < 3 {
		cutIdx = len(listTopTrending)
	}
	listTopTrending = listTopTrending[0:cutIdx]
	return listTopTrending
}

func (a *Activity) setInfoTextTrending(v *entity.User, idx int) string {
	setTextTimes := "like"
	if len(v.Photo.Like) > 1 {
		setTextTimes = "likes"
	}
	informationText := fmt.Sprintf("%d. %v photo got %d %v", idx+1, v.Name, len(v.Photo.Like), setTextTimes)
	return informationText
}

func (a *Activity) sortingForPhotoUsers() {
	sort.Slice(a.userRegistered, func(i, j int) bool {
		if a.userRegistered[i].Photo == nil {
			return false
		}
		if a.userRegistered[j].Photo == nil {
			return true
		}
		if len(a.userRegistered[i].Photo.Like) > len(a.userRegistered[j].Photo.Like) {
			return true
		}
		if len(a.userRegistered[i].Photo.Like) < len(a.userRegistered[j].Photo.Like) {
			return false
		}
		x := a.userRegistered[i].Photo.TimeOfLastLike
		y := a.userRegistered[j].Photo.TimeOfLastLike
		return x.Before(y)
	})
}
