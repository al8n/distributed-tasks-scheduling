package service

func GetSum(inputs ...int64) (rst int64) {
	for _, v := range inputs {
		rst += v
	}
	return
}
