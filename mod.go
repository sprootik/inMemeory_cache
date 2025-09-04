package cache

import "time"

// modifires parameters for modification
type modifires struct {
	ttl int64
}

// ChangeTTL change ttl of element
func ChangeTTL(t time.Duration) modifires {
	return modifires{ttl: t.Nanoseconds()}
}
