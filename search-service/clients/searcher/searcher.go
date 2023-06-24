package clients

import (
	"context"
)

type Searcher interface {
	Search(ctx context.Context) error
}
