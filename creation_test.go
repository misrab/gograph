package graph

import (
	"testing"
)


const (
	TEST_EDGE_LIST="./testing/edge_list.csv"
)

func TestReadEdgeList(t *testing.T) {
	ReadEdgeList(TEST_EDGE_LIST)
}
