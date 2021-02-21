package util

func SnakeToCamel(snake string) (camel string) {
	var camelBytes []byte
	b := []byte(snake)
	toUpper := false
	for _, v := range b {
		b2 := v
		if v == byte('_') {
			toUpper = true
			continue
		}
		if toUpper {
			b2 = v - 32
			toUpper = false
		}
		camelBytes = append(camelBytes, b2)
	}
	if len(camelBytes) > 0 && camelBytes[0] >= 97 && camelBytes[0] <= 122 {
		camelBytes[0] = camelBytes[0] - 32
	}
	camel = string(camelBytes)
	return
}
