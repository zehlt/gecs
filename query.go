package gecs

import (
	"github.com/zehlt/gecs/signature"
)

type Access []interface{}

type Exclude []interface{}

type Query struct {
	w            World
	access_sign  signature.Signature
	exclude_sign signature.Signature
}
