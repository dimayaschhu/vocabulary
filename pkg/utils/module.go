package utils

import "go.uber.org/fx"

var Module = fx.Provide(
	NewObjectMatcher,
	NewPlaceholderValidator,
)
