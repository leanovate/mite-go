package domain_test

import (
	"github.com/leanovate/mite-go/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestToday(t *testing.T) {
	expected := time.Now().Local().Format("2006-01-02")
	actual := domain.Today().String()

	assert.Equal(t, expected, actual)
}

func TestParseLocalDate(t *testing.T) {
	expected := domain.NewLocalDate(time.Date(1970, time.January, 1, 0, 0, 0, 0, time.Local))
	actual, err := domain.ParseLocalDate("1970-01-01")

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)

	_, err = domain.ParseLocalDate("1970-01-01T00:00:00Z")

	assert.IsType(t, &time.ParseError{}, err)
}

func TestLocalDate_Add(t *testing.T) {
	expected := domain.NewLocalDate(time.Date(1971, time.February, 2, 0, 0, 0, 0, time.Local))
	actual := domain.
		NewLocalDate(time.Date(1970, time.January, 1, 0, 0, 0, 0, time.Local)).
		Add(1, 1, 1)

	assert.Equal(t, expected, actual)
}

func TestLocalDate_String(t *testing.T) {
	expected := "1970-01-01"
	actual := domain.NewLocalDate(time.Date(1970, time.January, 1, 0, 0, 0, 0, time.Local)).String()

	assert.Equal(t, expected, actual)
}

func Test_ParseMinutes(t *testing.T) {
	expected := domain.NewMinutes(23)
	actual, err := domain.ParseMinutes("23m")

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)

	actual, err = domain.ParseMinutes("22m33s")

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)

	_, err = domain.ParseMinutes("1970-01-01")
	assert.NotNil(t, err)
}

func TestMinutes_Value(t *testing.T) {
	expected := 23
	actual := domain.NewMinutes(23).Value()

	assert.Equal(t, expected, actual)
}

func TestMinutes_String(t *testing.T) {
	expected := "23m"
	actual := domain.NewMinutes(23).String()

	assert.Equal(t, expected, actual)
}
