package entity_test

import (
	"testing"

	"activity-reporter-cli/controller"
	"activity-reporter-cli/variable"

	"github.com/stretchr/testify/assert"
)

func TestEntityPhoto(t *testing.T) {
	t.Run("should successful when the user wants to upload/add a photo", func(t *testing.T) {
		socialGraph := controller.NewSocialGraph()
		socialGraph.FollowUser("Alice", "Bob")
		getErrUploadPhoto := socialGraph.UploadPhoto("Bob")
		getErrLikePhoto := socialGraph.LikePhoto("Alice", "Bob")

		assert.Nil(t, getErrUploadPhoto)
		assert.Nil(t, getErrLikePhoto)
	})

	t.Run("should fail like photo when the user already like the photo", func(t *testing.T) {
		socialGraph := controller.NewSocialGraph()
		socialGraph.FollowUser("Alice", "Bob")
		getErrUploadPhoto := socialGraph.UploadPhoto("Bob")
		socialGraph.LikePhoto("Alice", "Bob")
		getErrLikePhoto := socialGraph.LikePhoto("Alice", "Bob")

		assert.Nil(t, getErrUploadPhoto)
		assert.NotNil(t, getErrLikePhoto)
		assert.Equal(t, variable.ErrorAlreadyLIkePhoto, getErrLikePhoto)
	})
}
