package repositories_test

import (
	"testing"

	"github.com/ThanaponBunchot/go-unit-test/repositories"
	"github.com/stretchr/testify/assert"
)

func TestGrade(t *testing.T) {

	type testCase struct {
		expect      string
		score       int
		description string
	}
	cases := []testCase{
		{score: 80, expect: "A", description: "grade A test"},
		{score: 70, expect: "B", description: "grade B test"},
		{score: 60, expect: "C", description: "grade C test"},
		{score: 50, expect: "D", description: "grade D test"},
		{score: 40, expect: "F", description: "grade F test"},
	}
	for _, v := range cases {
		t.Run(v.description, func(t *testing.T) {
			actual := repositories.CheckGrade(v.score)
			expect := v.expect
			assert.Equal(t, expect, actual, "The two words should be the same.")
		})
	}

}
