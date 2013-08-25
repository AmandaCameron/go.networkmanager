package nm

func bytesToString(src []interface{}) string {
	var tmp []byte
	for _, b := range src {
		tmp = append(tmp, b.(byte))
	}

	return string(tmp)
}
