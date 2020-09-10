package config

var GetAccountNumber = func() *string {
	n := "789"
	return &n
}

func GetConfigName() string {
	var configName string
	accountNumber := GetAccountNumber()
	switch *accountNumber {
	case "123":
		configName = "test1"
	case "456":
		configName = "test2"
	}
	return configName
}
