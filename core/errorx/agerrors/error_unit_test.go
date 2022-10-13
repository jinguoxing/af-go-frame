package agerrors

import (
    "github.com/jinguoxing/af-go-frame/core/errorx/agcodes"
    "github.com/stretchr/testify/assert"
    "testing"
)

func Test_New(t *testing.T) {

    t.Run("new", func(t *testing.T) {
        err := New("1")
        assert.NotEqual(t, err, nil)
        assert.Equal(t, err.Error(), "1")
    })

    t.Run("newF", func(t *testing.T) {
        err := NewF("%d", 1)
        assert.NotEqual(t, err, nil)
        assert.Equal(t, err.Error(), "1")
    })

}

func Test_Code(t *testing.T) {

    t.Run("newCode", func(t *testing.T) {
        err := NewCode(agcodes.CodeUnknown, "111")
        assert.Equal(t, Code(err), agcodes.CodeUnknown)
        assert.Equal(t, err.Error(), "111")
    })

}
