package schedulers

import "context"

type Job func(ctx context.Context)
