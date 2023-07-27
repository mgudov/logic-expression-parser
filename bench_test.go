package lep

import (
	"testing"
)

var (
	benchSmallQuery  = `a=1000 || b="foo"`
	benchMediumQuery = `a>1000 && b<5000 || c="foo" && d="bar" || e!="test" || e starts_with "some"`
	benchLargeQuery  = `(a=false) && b>=c && (d<1000 || e>=2000 || (g!=5000 && g>=1000 && h="foo")) || j in [1,2,3,4,5] && k>dt:"2020-01-01" || m starts_with "foo" && m ends_with n`
)

func BenchmarkSmallQuery(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := ParseExpression(benchSmallQuery); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkMediumQuery(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := ParseExpression(benchMediumQuery); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkLargeQuery(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := ParseExpression(benchLargeQuery); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkSmallQueryWithMemo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := ParseExpression(benchSmallQuery, Memoize(true)); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkMediumQueryWithMemo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := ParseExpression(benchMediumQuery, Memoize(true)); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkLargeQueryWithMemo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := ParseExpression(benchLargeQuery, Memoize(true)); err != nil {
			b.Error(err)
		}
	}
}
