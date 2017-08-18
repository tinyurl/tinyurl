package uriuuid

var (
	// DefaultChars chars to consists of random string
	DefaultChars = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	// DefaultLen default random string length
	DefaultLen = 6
)

// UriUUID interface for extension
type UriUUID interface {
	New() string
	NewLen(int) string
	NewLenChars(int, []byte) string
}
