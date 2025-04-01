package representation

import (
	"context"
	"io"
	"net/url"

	"github.com/whosonfirst/go-reader"
	"github.com/whosonfirst/go-whosonfirst-uri"
)

type ReaderSource struct {
	Source
	reader reader.Reader
}

func init() {

	err := RegisterSource(context.Background(), "reader", NewReaderSource)

	if err != nil {
		panic(err)
	}
}

func NewReaderSource(ctx context.Context, uri string) (Source, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	q := u.Query()

	reader_uri := q.Get("reader-uri")

	r, err := reader.NewReader(ctx, reader_uri)

	if err != nil {
		return nil, err
	}

	s := &ReaderSource{
		reader: r,
	}
	return s, nil
}

func (s *ReaderSource) GetFeature(ctx context.Context, id int64, uri_args *uri.URIArgs) (io.ReadSeekCloser, error) {

	rel_path, err := uri.Id2RelPath(id, uri_args)

	if err != nil {
		return nil, err
	}

	return s.reader.Read(ctx, rel_path)
}
