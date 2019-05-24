package limit

import (
	"fmt"
)

//创建限时缓存key
func CreateLimiKey(key string, publicsh_model string, publish_type string) (string) {
	return fmt.Sprintf("zldz:sms:%s:%s:%s", publish_type, publicsh_model, key)
}