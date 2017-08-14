// sender implement a uri sender
// each call of UriSender, it will return a unique string
// the unique string represents a url which one-to-one with each other

package sender

var (
	CHARS       = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	index int64 = 0
	start int64 = 0
	// m     sync.RWMutex
)

func New() (int64, string) {
	// m.RLock()
	// defer m.RUnlock()
	start++
	return start, GetByteByIndex(start, CHARS)
}

func GetByteByIndex(index int64, chars []byte) string {
	baseLen := int64(len(chars))
	var container []byte
	for ; index != 0; index = index / baseLen {
		m := index % baseLen
		container = append(container, chars[m])
	}

	return string(container)
}
