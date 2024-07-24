package processor

import (
	"dev03/internal/parser"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Processor interface {
	Sort(p *parser.CurrentParser)
	removeDuplicates(data []string) []string
	reverse(data []string)
}

func NewProcessor() Processor {
	return &CurrentProcessor{}
}

type CurrentProcessor struct {
}

func (cp *CurrentProcessor) Sort(p *parser.CurrentParser) {
	if p.Flags.U {
		cp.removeDuplicates(p.Data)
	}

	sort.Slice(p.Data, func(i, j int) bool {
		lineA := p.Data[i]
		lineB := p.Data[j]

		if p.Flags.K > 0 && p.Flags.K <= len(strings.Fields(lineA)) && p.Flags.K <= len(strings.Fields(lineB)) {
			tmpA := strings.Fields(lineA)[p.Flags.K-1]
			tmpB := strings.Fields(lineB)[p.Flags.K-1]

			if p.Flags.N {
				NumA, err := strconv.Atoi(tmpA)
				if err != nil {
					log.Fatalf("Some errors in Processor's sorting: %s", err)
				}
				NumB, err := strconv.Atoi(tmpB)
				if err != nil {
					log.Fatalf("Some errors in Processor's sorting: %s", err)
				}
				return NumA < NumB
			}

		}

		return lineA < lineB

	})

	if p.Flags.R {
		cp.reverse(p.Data)
	}
}

func (cp *CurrentProcessor) removeDuplicates(data []string) []string {
	uniqueData := make(map[string]struct{})
	result := make([]string, 0)

	for _, elem := range data {
		if _, ok := uniqueData[elem]; !ok {
			uniqueData[elem] = struct{}{}
			result = append(result, elem)
		}
	}

	return result
}

func (cp *CurrentProcessor) reverse(data []string) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
