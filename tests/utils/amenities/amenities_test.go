package amenities_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"ascenda_assessment/utils/amenities"
	amen "ascenda_assessment/utils/amenities"
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
		l := amen.AmenityList{}
		for _, a := range tc.list {
			l.Add(a)
		}
		amenities.CleanGeneralListDuplicatedItem(l)
		_, isPoolExist := l[amen.Pool]
		ass.Equal(tc.expect, isPoolExist)
	}
}
