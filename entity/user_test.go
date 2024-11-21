package entity_test

import (
	"testing"

	"activity-reporter-cli/controller"

	"github.com/stretchr/testify/assert"
)

func TestEntityUser(t *testing.T) {
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
}
