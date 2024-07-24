package processor

import (
	"sort"
	"strings"
)

type Processor interface {
	SetInfo(data []string)
	GetInfo() map[string][]string
}

func NewProcessor() Processor {
	return &CurrentProcessor{
		Info: make(map[string][]string),
	}
}

type CurrentProcessor struct {
	Info map[string][]string
}

func (cp *CurrentProcessor) SetInfo(data []string) {
	for _, str := range data {
		cp.CheckAnagram(str)
	}

	for k := range cp.Info {
		sort.Slice(cp.Info[k], func(i, j int) bool {
			return cp.Info[k][i] < cp.Info[k][j]
		})
	}

	for k := range cp.Info {
		if len(cp.Info[k]) < 2 {
			delete(cp.Info, k)
		}
	}

}

func (cp *CurrentProcessor) GetInfo() map[string][]string {
	return cp.Info
}

func (cp *CurrentProcessor) CheckAnagram(str string) {
	tmpStr := sortString(strings.ToLower(str))
	flag := false
	for k := range cp.Info {
		if tmpStr == sortString(k) && !checkExists(str, cp.Info[k]) {
			cp.Info[k] = append(cp.Info[k], str)
			flag = true
			break
		}
	}

	if !flag {
		cp.Info[str] = append(cp.Info[str], str)
	}
}

func sortString(str string) string {
	sortedRunes := []rune(str)
	sort.Slice(sortedRunes, func(i, j int) bool {
		return sortedRunes[i] < sortedRunes[j]
	})
	return string(sortedRunes)
}

func checkExists(value string, data []string) bool {
	for _, v := range data {
		if v == value {
			return true
		}
	}
	return false
}
