package kit

// Binder ...
type Binder interface {
	Marshal(interface{}) ([]byte, error)
	Unmarshal([]byte, interface{}) error
}
