// uriuuid implement a uri sender
// each call of UriSender, it will return a unique string
// the unique string represents a url which one-to-one with each other

package uriuuid

type SenderURIUUID struct {
	start int64
}

// New new uri with default length and []byte
func (sender *SenderURIUUID) New() string {
	// m.RLock()
	// defer m.RUnlock()
	sender.start++
	return GetByteByIndex(sender.start, DefaultChars)
}

// New new uri with default length and []byte
func (sender *SenderURIUUID) NewLen(length int) string {
	// m.RLock()
	// defer m.RUnlock()
	sender.start++
	return GetByteByIndex(sender.start, DefaultChars)
}

// New new uri with default length and []byte
func (sender *SenderURIUUID) NewLenDefaultChars(length int, DefaultChars []byte) string {
	// m.RLock()
	// defer m.RUnlock()
	sender.start++
	return GetByteByIndex(sender.start, DefaultChars)
}

func GetByteByIndex(index int64, DefaultChars []byte) string {
	baseLen := int64(len(DefaultChars))
	var container []byte
	for ; index != 0; index = index / baseLen {
		m := index % baseLen
		container = append(container, DefaultChars[m])
	}

	return string(container)
}
