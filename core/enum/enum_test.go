package enum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type FlowchartClassEnum Object

var (
	FlowchartClassConfigStatusNormal      = New[FlowchartClassEnum](1, "normal", "正常")
	FlowchartClassConfigStatusMissingRole = New[FlowchartClassEnum](2, "missingRole", "缺失角色")
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
	s1 := ToString[FlowchartClassEnum](1)
	assert.Equal(t, s1, "normal")

	a := int8(2)
	s2 := ToString[FlowchartClassEnum](a)
	assert.Equal(t, s2, "missingRole")
}

func TestIs(t *testing.T) {
	s1 := Is[FlowchartClassEnum](4)
	assert.Equal(t, s1, false)

	s2 := Is[FlowchartClassEnum]("xxx")
	assert.Equal(t, s2, false)
}

func TestDisplay(t *testing.T) {
	assert.Equal(t, FlowchartClassConfigStatusNormal.Display, "正常")
	assert.Equal(t, Get[FlowchartClassEnum]("normal").Display, "正常")
	assert.Equal(t, len(List[FlowchartClassEnum]()), 2)
}
