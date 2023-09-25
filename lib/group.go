package tracing

import "context"

func Group(ctx context.Context) context.Context {
	g, ok := groupFromContext(ctx)
	if ok {
		logger.Debugw("group already presented in context", "group_uuid", g)
		return ctx
	}

	logger.Debugw("creating group")

	newGroup, _, err := api.DefaultApi.CreateGroup(ctx).Execute()
	if err != nil {
		logger.Errorw("cant create group", "error", err)
		return ctx
	}

	return GroupToContext(ctx, newGroup.Uuid)
}
