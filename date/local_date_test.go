package date_test

import (
	"github.com/leanovate/mite-go/date"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDate_Today(t *testing.T) {
	expected := time.Now().Local().Format("2006-01-02")
	actual := date.Today().String()

	assert.Equal(t, expected, actual)
}

func TestDate_Parse(t *testing.T) {
	expected := date.From(time.Date(1970, time.January, 1, 0, 0, 0, 0, time.Local))
	actual, err := date.Parse("1970-01-01")

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)

	_, err = date.Parse("1970-01-01T00:00:00Z")

	assert.IsType(t, &time.ParseError{}, err)
}

func TestDate_Add(t *testing.T) {
	expected := date.From(time.Date(1971, time.February, 2, 0, 0, 0, 0, time.Local))
	actual := date.
		From(time.Date(1970, time.January, 1, 0, 0, 0, 0, time.Local)).
		Add(1, 1, 1)

	assert.Equal(t, expected, actual)
}