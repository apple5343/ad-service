package time

import (
	"context"
)

func (r *timeRepository) Set(ctx context.Context, day int) error {
	return r.rdb.Set(ctx, "time", day, 0).Err()

}
