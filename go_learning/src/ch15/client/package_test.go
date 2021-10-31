package client

import (
	"ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log("TestPackage")
	t.Log(series.GetFibonacciSeries(10))
}
