package capsolver

import "context"

type contextkey string

const ContextKey contextkey = "data"

// WithContext wraps the given session with a context.
func WithContext(ctx context.Context, s *Session) context.Context {
	return context.WithValue(ctx, ContextKey, s)
}

// FromContext retrieves the session from the context.
// This will return nil if the context does not contain a session.
func FromContext(ctx context.Context) *Session {
	if s, ok := ctx.Value(ContextKey).(*Session); ok {
		return s
	}
	return nil
}
