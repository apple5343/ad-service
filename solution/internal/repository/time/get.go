package time

import "context"

func (r *timeRepository) Get(ctx context.Context) (int, error) {
	return r.rdb.Get(ctx, "time").Int()
}
