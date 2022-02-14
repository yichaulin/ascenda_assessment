package amenities_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"ascenda_assessment/utils/amenities"
	amen "ascenda_assessment/utils/amenities"
	"ascenda_assessment/utils/string_list"
)

func TestCleanGeneralListDuplicatedItem(t *testing.T) {
	ass := assert.New(t)
	tests := []struct {
		list   []string
		expect bool
	}{
		{
			list:   []string{amen.Pool, amen.OutdoorPool, amen.IndoorPool},
			expect: false,
		},
		{
			list:   []string{amen.Pool, amen.IndoorPool},
			expect: false,
		},
		{
			list:   []string{amen.Pool, amen.OutdoorPool},
			expect: false,
		},
		{
			list:   []string{amen.Pool},
			expect: true,
		},
	}

	for _, tc := range tests {
		l := string_list.New(tc.list...)
		amenities.CleanGeneralListDuplicatedItem(l)
		_, isPoolExist := l[amen.Pool]
		ass.Equal(tc.expect, isPoolExist)
	}
}
