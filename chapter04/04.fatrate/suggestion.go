package main

func GetFatRateSuggestion() *fatRateSuggestion {
	return &fatRateSuggestion{
		suggArr: [][][]int{

			{ //	男

				{ // 18-39

				},
				{ //

				},
			},

			{ // 女

			},
		},
	}
}

type fatRateSuggestion struct {
	suggArr [][][]int
}

func (fatRateSuggestion) GetSuggestion(person *Person) string {
	//todo
	return ""
}

// 性别
func (s *fatRateSuggestion) getIndexOfSex(sex string) int {
	if sex == "男" {
		return 0
	}
	return 1
}

// 年龄
func (s *fatRateSuggestion) getIndexOfAge(age int) int {
	switch {
	case age >= 18 && age <= 39:
		return 0
	case age >= 40 && age <= 59:
		return 1
	case age >= 60:
		return 2
	default:
		return -1
	}
}

// 结论
func (s *fatRateSuggestion) translateResult(idx int) string {
	switch idx {
	case 0:
		return "偏瘦"
	case 1:
		return "标准"
	case 2:
		return "偏重"
	case 3:
		return "肥胖"
	case 4:
		return "过度肥胖"
	default:
		return "无法判断"
	}
}
