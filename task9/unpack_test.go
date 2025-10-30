package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	input := "a4bc2d5e"

	res, err := UnpackString(input)
	assert.NoError(t, err)
	assert.Equal(t, "aaaabccddddde", res)
}

func TestSame(t *testing.T) {
	input := "abcde"

	res, err := UnpackString(input)
	assert.NoError(t, err)
	assert.Equal(t, "abcde", res)
}

func TestDigidsErr(t *testing.T) {
	input := "45"

	_, err := UnpackString(input)
	assert.ErrorContains(t, err, "digid without prev char")
}

func TestEmpty(t *testing.T) {
	input := ""

	res, err := UnpackString(input)
	assert.NoError(t, err)
	assert.Equal(t, "", res)
}

func TestDigidsEsc(t *testing.T) {
	input := "qwe\\4\\5"

	res, err := UnpackString(input)
	assert.NoError(t, err)
	assert.Equal(t, "qwe45", res)
}
