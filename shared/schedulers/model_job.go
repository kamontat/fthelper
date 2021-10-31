package schedulers

import "context"

// Job use for doing task in schedulers
type Job func(ctx context.Context)
