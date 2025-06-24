package agent

import "encoding/hex"

const (
	Support_Working     = "Working/Tested"
	Support_Untested    = "Untested (use at own risk)"
	Support_Unsupported = "Unsupported (for now)"
)

func invertHexBytes(hexStr string) string {
	bytes, _ := hex.DecodeString(hexStr)
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return hex.EncodeToString(bytes)
}
