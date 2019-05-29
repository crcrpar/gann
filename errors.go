package gann

import "github.com/pkg/errors"

var (
	errDimensionMismatch         = errors.New("dimension mismatch")
	errInvalidIndex              = errors.New("invalid index")
	errInvalidKeyVector          = errors.New("invalid key vector")
	errItemNotFoundOnGivenItemID = errors.New("item not found for given item id")
)
