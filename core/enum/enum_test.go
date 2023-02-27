package enum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type FlowchartClassEnum MapObject

var (
	FlowchartClassConfigStatusNormal      = New[FlowchartClassEnum](1, "normal")
	FlowchartClassConfigStatusMissingRole = New[FlowchartClassEnum](2, "missingRole")
)

type SystemRoleClassEnum MapObject

var (
	SystemRoleClassEnumConfigStatusNormal      = New[SystemRoleClassEnum](1, "normals")
	SystemRoleClassEnumConfigStatusMissingRole = New[SystemRoleClassEnum](2, "missingRoles")
)

func TestToInteger(t *testing.T) {
	i := ToInteger[FlowchartClassEnum]("normal")
	assert.Equal(t, i, 1)
}

func TestToString(t *testing.T) {
	s := ToString[FlowchartClassEnum](1)
	assert.Equal(t, s, "normal")
}

func TestGet(t *testing.T) {
	s := Get[SystemRoleClassEnum](uint8(1))
	t.Log(s.Integer, s.String)
	assert.NotNil(t, s)
}
