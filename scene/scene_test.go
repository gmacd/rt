package scene

import (
	"testing"

	"github.com/gmacd/rt/test"
)

func TestCreateSimpleScene(t *testing.T) {
	s := CreateCornellBoxOfSpheres()
	test.NotNil(t, s, "Scene not built")
}
