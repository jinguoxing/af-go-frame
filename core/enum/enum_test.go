package enum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type FlowchartClassEnum Object

var (
	FlowchartClassConfigStatusNormal      = New[FlowchartClassEnum](1, "normal")
	FlowchartClassConfigStatusMissingRole = New[FlowchartClassEnum](2, "missingRole")
)

type SystemRoleClassEnum Object

var (
	SystemRoleClassEnumConfigStatusNormal      = New[SystemRoleClassEnum](1, "normals")
	SystemRoleClassEnumConfigStatusMissingRole = New[SystemRoleClassEnum](2, "missingRoles")
)

func TestToInteger(t *testing.T) {
	i := ToInteger[FlowchartClassEnum]("normal").Int()
	assert.Equal(t, i, 1)
}

func TestToString(t *testing.T) {
	s := ToString[FlowchartClassEnum](1)
	assert.Equal(t, s, "normal")
}