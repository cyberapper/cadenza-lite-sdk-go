// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"

	"github.com/cyberapper/cadenza-lite-sdk-go/internal/apijson"
	"github.com/cyberapper/cadenza-lite-sdk-go/internal/requestconfig"
)

type Offset[T any] struct {
	Data   []T        `json:"data"`
	Offset int64      `json:"offset,nullable"`
	Limit  int64      `json:"limit,nullable"`
	JSON   offsetJSON `json:"-"`
	cfg    *requestconfig.RequestConfig
	res    *http.Response
}

// offsetJSON contains the JSON metadata for the struct [Offset[T]]
type offsetJSON struct {
	Data        apijson.Field
	Offset      apijson.Field
	Limit       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Offset[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r offsetJSON) RawJSON() string {
	return r.raw
}

// NextPage returns the next page as defined by this pagination style. When there
// is no next page, this function will return a 'nil' for the page value, but will
// not return an error
func (r *Offset[T]) GetNextPage() (res *Offset[T], err error) {
	// This page represents a response that isn't actually paginated at the API level
	// so there will never be a next page.
	cfg := (*requestconfig.RequestConfig)(nil)
	if cfg == nil {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *Offset[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &Offset[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type OffsetAutoPager[T any] struct {
	page *Offset[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewOffsetAutoPager[T any](page *Offset[T], err error) *OffsetAutoPager[T] {
	return &OffsetAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *OffsetAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *OffsetAutoPager[T]) Current() T {
	return r.cur
}

func (r *OffsetAutoPager[T]) Err() error {
	return r.err
}

func (r *OffsetAutoPager[T]) Index() int {
	return r.run
}
