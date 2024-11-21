package controller_test

import (
	"testing"

	"activity-reporter-cli/controller"
	"activity-reporter-cli/variable"

	"github.com/stretchr/testify/assert"
)

func TestActivityFollow(t *testing.T) {
	t.Run("successful in following people when the user inputs their account name and the account name of their follow target", func(t *testing.T) {
		socialGraph := controller.NewSocialGraph()
		errMsg := socialGraph.FollowUser("Alice", "Bob")
		assert.Nil(t, errMsg)
	})

	t.Run("should fail to follow someone when the user has already followed that user", func(t *testing.T) {
		socialGraph := controller.NewSocialGraph()
		socialGraph.FollowUser("Alice", "Bob")
		errMsgFollowBack := socialGraph.FollowUser("Bob", "Alice")
		errMsg := socialGraph.FollowUser("Alice", "Bob")
		assert.Nil(t, errMsgFollowBack)
		assert.NotNil(t, errMsg)
		assert.Equal(t, variable.ErrorAlreadyFollow, errMsg)
	})

	t.Run("should fail to follow someone when the user only inputs spaces into their name account", func(t *testing.T) {
		socialGraph := controller.NewSocialGraph()
		errMsg := socialGraph.FollowUser(" ", "Ronaldinho")
		assert.NotNil(t, errMsg)
		assert.Equal(t, variable.ErrorInvalidKeyword, errMsg)
	})

	t.Run("should fail to follow someone when the user only inputs spaces into their target account", func(t *testing.T) {
		socialGraph := controller.NewSocialGraph()
		errMsg := socialGraph.FollowUser("Ronaldo", " ")
		assert.NotNil(t, errMsg)
		assert.Equal(t, variable.ErrorInvalidKeyword, errMsg)
	})

	t.Run("should fail to follow someone when the user contains numbers in their input", func(t *testing.T) {
		socialGraph := controller.NewSocialGraph()
		errMsg := socialGraph.FollowUser("Ronaldo7", "Messi")
		assert.NotNil(t, errMsg)
		assert.Equal(t, variable.ErrorInvalidKeyword, errMsg)
	})

	t.Run("should fail to follow when the user wants to follow their own account", func(t *testing.T) {
		socialGraph := controller.NewSocialGraph()
		errMsg := socialGraph.FollowUser("Marcelo", "Marcelo")
		assert.NotNil(t, errMsg)
		assert.Equal(t, variable.ErrorUserFollowThemselves, errMsg)
	})
}

func TestActivityUpload(t *testing.T) {
	t.Run("should successfully upload a photo when the user has not uploaded a photo before", func(t *testing.T) {
		socialGraph := controller.NewSocialGraph()
		socialGraph.FollowUser("Faisal", "Marcelo")
		errMsg := socialGraph.UploadPhoto("Faisal")
		assert.Nil(t, errMsg)
	})

	t.Run("should fail upload a photo when the user has uploaded a photo before", func(t *testing.T) {
		socialGraph := controller.NewSocialGraph()
		socialGraph.FollowUser("Faisal", "Marcelo")
		socialGraph.UploadPhoto("Faisal")
		errMsg := socialGraph.UploadPhoto("Faisal")
		assert.NotNil(t, errMsg)
		assert.Equal(t, variable.ErrorAlreadyUploadPhoto, errMsg)
	})

	t.Run("should fail upload a photo when the user input an unregistered username", func(t *testing.T) {
		socialGraph := controller.NewSocialGraph()
		username := "Faisal"
		errMsg := socialGraph.UploadPhoto(username)
		assert.NotNil(t, errMsg)
		assert.Equal(t, variable.ErrorUnknownUserForTesting, errMsg)
	})
}

func TestActivityLikePhoto(t *testing.T) {
	t.Run("should successfully like a photo when the user's photo target is registered", func(t *testing.T) {
		socialGraph := controller.NewSocialGraph()
		socialGraph.FollowUser("Alice", "Bob")
		socialGraph.UploadPhoto("Bob")
		errMsg := socialGraph.LikePhoto("Alice", "Bob")
		assert.Nil(t, errMsg)
	})

	t.Run("should successfully like a photo when the user like their own photo", func(t *testing.T) {
		socialGraph := controller.NewSocialGraph()
		socialGraph.FollowUser("Alice", "Bob")
		socialGraph.UploadPhoto("Alice")
		errMsg := socialGraph.LikePhoto("Alice", "Alice")
		assert.Nil(t, errMsg)
	})

	t.Run("should failed like a photo when the user's target does not have a photo", func(t *testing.T) {
		socialGraph := controller.NewSocialGraph()
		socialGraph.FollowUser("Alice", "Bob")
		socialGraph.UploadPhoto("Alice")
		errMsg := socialGraph.LikePhoto("Alice", "Bob")
		assert.NotNil(t, errMsg)
		assert.Equal(t, variable.SetErrorDoesNotHavePhoto("Bob"), errMsg)
	})

	t.Run("should failed like a photo when the user's target is not registered", func(t *testing.T) {
		socialGraph := controller.NewSocialGraph()
		socialGraph.FollowUser("Alice", "Bob")
		socialGraph.UploadPhoto("Alice")
		errMsg := socialGraph.LikePhoto("Alice", "Kiki")
		assert.NotNil(t, errMsg)
		assert.Equal(t, variable.SetErrorUnkownUser("Kiki"), errMsg)
	})

	t.Run("should failed like a photo when the user's who want to follow is not registered", func(t *testing.T) {
		socialGraph := controller.NewSocialGraph()
		socialGraph.FollowUser("Alice", "Bob")
		socialGraph.UploadPhoto("Alice")
		errMsg := socialGraph.LikePhoto("Ronaldo", "Kiki")
		assert.NotNil(t, errMsg)
		assert.Equal(t, variable.SetErrorUnkownUser("Ronaldo"), errMsg)
	})

	t.Run("should failed like their own photo when the user's who want to follow is not registered", func(t *testing.T) {
		socialGraph := controller.NewSocialGraph()
		socialGraph.FollowUser("Alice", "Ronaldo")
		socialGraph.UploadPhoto("Bob")
		errMsg := socialGraph.LikePhoto("Ronaldo", "Ronaldo")
		assert.NotNil(t, errMsg)
		assert.Equal(t, variable.ErrorUserDoesNotHavePhoto, errMsg)
	})

	t.Run("should failed like photo when the user not follow their target user", func(t *testing.T) {
		socialGraph := controller.NewSocialGraph()
		socialGraph.FollowUser("Alice", "Bob")
		socialGraph.FollowUser("Alice", "Ronaldo")
		socialGraph.UploadPhoto("Ronaldo")
		errMsg := socialGraph.LikePhoto("Bob", "Ronaldo")
		assert.NotNil(t, errMsg)
		assert.Equal(t, variable.SetErrorUnableLikePhoto("Ronaldo"), errMsg)
	})

	t.Run("should failed like photo when the user already like the photo", func(t *testing.T) {
		socialGraph := controller.NewSocialGraph()
		socialGraph.FollowUser("Alice", "Bob")
		socialGraph.FollowUser("Alice", "Ronaldo")
		socialGraph.UploadPhoto("Ronaldo")
		socialGraph.LikePhoto("Alice", "Ronaldo")
		errMsg := socialGraph.LikePhoto("Alice", "Ronaldo")
		assert.NotNil(t, errMsg)
		assert.Equal(t, variable.ErrorAlreadyLIkePhoto, errMsg)
	})

	t.Run("should failed like photo when the user input is not valid", func(t *testing.T) {
		socialGraph := controller.NewSocialGraph()
		socialGraph.FollowUser("Alice", "Ronaldo")
		socialGraph.UploadPhoto("Ronaldo")
		errMsg := socialGraph.LikePhoto(" ", " ")
		assert.NotNil(t, errMsg)
		assert.Equal(t, variable.ErrorInvalidKeyword, errMsg)
	})

}

func TestActivityList(t *testing.T) {
	t.Run("should successfully provide list activity when the user inputs who wants to see the list", func(t *testing.T) {
		socialGraph := controller.NewSocialGraph()

		socialGraph.FollowUser("Alice", "Bob")
		socialGraph.FollowUser("Alice", "Bill")
		socialGraph.FollowUser("John", "Bob")
		socialGraph.FollowUser("Bob", "Alice")
		socialGraph.FollowUser("Bob", "Bill")
		socialGraph.FollowUser("John", "Alice")

		socialGraph.UploadPhoto("Alice")
		socialGraph.LikePhoto("Bob", "Alice")
		socialGraph.UploadPhoto("Bill")
		socialGraph.LikePhoto("Bob", "Bill")
		socialGraph.LikePhoto("Bill", "Bill")
		socialGraph.LikePhoto("Alice", "Bill")
		getErrUnable := socialGraph.LikePhoto("Bill", "Alice")

		result, err := socialGraph.ActivityUser("Bob")
		expectedResult := []string{"Alice uploaded photo", "You liked Alice's photo", "Bill uploaded photo", "You liked Bill's photo", "Bill liked Bill's photo", "Alice liked Bill's photo"}

		assert.Nil(t, err)
		assert.NotNil(t, getErrUnable)
		assert.NotNil(t, result)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("should successfully provide an empty list if a user does not get any activity", func(t *testing.T) {
		socialGraph := controller.NewSocialGraph()
		socialGraph.FollowUser("Alice", "Bob")
		socialGraph.FollowUser("Ronaldo", "Alice")
		socialGraph.UploadPhoto("Alice")
		socialGraph.LikePhoto("Ronaldo", "Alice")
		result, err := socialGraph.ActivityUser("eBob")

		var expectedResult = variable.SetErrorUnkownUser("eBob")

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, expectedResult, err)
	})

}

func TestActivityTopTrending(t *testing.T) {
	t.Run("successfully produce a top trending list when the user requests a trending list", func(t *testing.T) {
		socialGraph := controller.NewSocialGraph()

		socialGraph.FollowUser("Alice", "Bob")
		socialGraph.FollowUser("Alice", "Bill")
		socialGraph.FollowUser("John", "Bob")
		socialGraph.FollowUser("Bob", "Alice")
		socialGraph.FollowUser("Bob", "Bill")
		socialGraph.FollowUser("Bill", "Bob")
		socialGraph.FollowUser("John", "Alice")
		socialGraph.FollowUser("Bob", "John")

		socialGraph.UploadPhoto("Alice")
		socialGraph.UploadPhoto("Bill")
		socialGraph.UploadPhoto("Bob")
		socialGraph.UploadPhoto("John")

		socialGraph.LikePhoto("Bob", "Bob")
		socialGraph.LikePhoto("Bill", "Bob")
		socialGraph.LikePhoto("Alice", "Bob")

		socialGraph.LikePhoto("John", "Alice")
		socialGraph.LikePhoto("Bob", "Alice")
		socialGraph.LikePhoto("Alice", "Alice")

		resultList := socialGraph.TrendingPhotos()

		expectedResult := []string{"1. Bob photo got 3 likes", "2. Alice photo got 3 likes"}
		assert.Equal(t, expectedResult, resultList)
	})

	t.Run("successfully produce a top trending list when the user requests a trending list", func(t *testing.T) {
		socialGraph := controller.NewSocialGraph()

		socialGraph.FollowUser("Alice", "Bob")
		socialGraph.FollowUser("Alice", "Bill")
		socialGraph.FollowUser("John", "Bob")
		socialGraph.FollowUser("Bob", "Alice")
		socialGraph.FollowUser("Bob", "Bill")
		socialGraph.FollowUser("John", "Alice")

		socialGraph.UploadPhoto("Alice")
		socialGraph.LikePhoto("Bob", "Alice")
		socialGraph.UploadPhoto("Bill")

		socialGraph.LikePhoto("Bob", "Bill")
		socialGraph.LikePhoto("Bill", "Bill")
		socialGraph.LikePhoto("Alice", "Bill")

		resultList := socialGraph.TrendingPhotos()

		expectedResult := []string{"1. Bill photo got 3 likes", "2. Alice photo got 1 like"}
		assert.NotNil(t, resultList)
		assert.Equal(t, expectedResult, resultList)
	})
}
