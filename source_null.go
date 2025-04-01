package representation

import (
	"context"
	"io"

	"github.com/whosonfirst/go-whosonfirst-uri"
)

type NullSource struct {
	Source
}

func init() {

	err := RegisterSource(context.Background(), "null", NewNullSource)

	if err != nil {
		panic(err)
	}
}

func NewNullSource(ctx context.Context, uri string) (Source, error) {
	s := &NullSource{}
	return s, nil
}

func (s *NullSource) GetFeature(ctx context.Context, id int64, uri_args *uri.URIArgs) (io.ReadSeekCloser, error) {
	return nil, ErrNotFound
}
