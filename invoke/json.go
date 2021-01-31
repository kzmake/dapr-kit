package util

import (
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var (
	// ContentType ...
	ContentType = "application/json"
	// Marshal ...
	Marshal = json.Marshal
	// Unmarshal ...
	Unmarshal = json.Unmarshal
)
