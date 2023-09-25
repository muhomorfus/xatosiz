package tracing

import "context"

const (
	groupCtxKey = "xatosiz_group"
	traceCtxKey = "xatosiz_trace"
)

func GroupToContext(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, groupCtxKey, id)
}

func groupFromContext(ctx context.Context) (string, bool) {
	g, ok := ctx.Value(groupCtxKey).(string)

	return g, ok
}

func TraceToContext(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, traceCtxKey, id)
}

func traceFromContext(ctx context.Context) (string, bool) {
	t, ok := ctx.Value(traceCtxKey).(string)

	return t, ok
}
