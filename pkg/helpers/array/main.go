package array

import "strconv"

// Проверка, находится элемент в списке или нет
func Includes(list []interface{}, key interface{}) bool {
	for _, value := range list {
		if value == key {
			return true
		}
	}
	return false
}

// Проверка, находится элемент в списке или нет
func ConvertStringToInt(list []string) ([]int, error) {
	result := []int{}
	for _, value := range list {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return result, err
		}
		result = append(result, intValue)
	}
	return result, nil
}
