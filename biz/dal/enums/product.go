package enums

func ReturnProductValues(key int64) (values string) {
	switch key {
	case 1:
		values = "潜在"
	case 2:
		values = "正式"
	default:
		values = "状态异常"
	}
	return
}
func ReturnPropertyValues(key int64) (values string) {
	switch key {
	case 1:
		values = "潜在"
	case 2:
		values = "正式"
	default:
		values = "状态异常"
	}
	return
}
