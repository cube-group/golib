package dataFormat

func FirstNotZeroInt(val ...int) int {
	for _, v := range val {
		if v != 0 {
			return v
		}
	}
	return 0
}

func FirstNotZeroInt64(val ...int64) int64 {
	for _, v := range val {
		if v != 0 {
			return v
		}
	}
	return 0
}

func FirstNotZeroFloat64(val ...float64) float64 {
	for _, v := range val {
		if v != 0 {
			return v
		}
	}
	return 0
}

func FirstNotNullString(val ...string) string {
	for _, v := range val {
		if v != "" {
			return v
		}
	}
	return ""
}

func FirstNotZeroString(val ...string) string {
	for _, v := range val {
		if v != "" && v != "0" {
			return v
		}
	}
	return "0"
}
