package utils

/**
 * @Description 添加qq群467890197 交流学习
 * @Author 熊猫
 * @Date 2022/1/13 17:16
 **/

func DdmKey(data string) string {
	if len(data) < 6 {
		return data
	}
	return data[:3] + "****" + data[len(data)-3:]
}

func IsDdmKey(data string) bool {
	return len(data) > 6 && data[3:len(data)-3] == "****"
}

func DdmMail(data string) string {
	return data[:3] + "****" + data[len(data)-8:]
}

func ISDdmMail(data string) bool {
	return len(data) > 11 && data[3:len(data)-8] == "****"
}

func DdmPassword(data string) string {
	return "******"
}

func IsDdmPassword(data string) bool {
	return data == "******"
}
