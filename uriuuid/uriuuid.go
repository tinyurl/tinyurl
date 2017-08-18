package uriuuid

// UriUUID interface for extension
type UriUUID interface {
	New() string
	NewLen(int) string
	NewLenChars(int, []byte) string
}
