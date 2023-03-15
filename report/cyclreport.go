package report

import (
	"fmt"
	"os"
	"sort"

	"strconv"
	"strings"
	"trident/utility"
)

func GenerateSortedCReport() {
	c := readOriginalReport()
	writeSortedReport(c)
}

func readOriginalReport() map[int][]string {
	lines := utility.ReadLines("./reportfile/1")
	lines = filterContentOut(lines, "a total cyclomatic complexity")
	return stringToMapUsingCycls(lines)
}

func stringToMapUsingCycls(lines []string) map[int][]string {
	cycls := make(map[int][]string)

	for _, v := range lines {
		cyclIndex := strings.LastIndex(v, " ") + 1

		cycl, _ := strconv.Atoi(v[cyclIndex : len(v)-1])
		cycls[cycl] = append(cycls[cycl], v)
	}

	return cycls
}

func filterContentOut(c []string, filter string) []string {
	for i, v := range c {
		if strings.Contains(v, filter) {
			c[i] = c[len(c)-1]
			c = c[:len(c)-1]
		}
	}
	return c
}

func writeSortedReport(cycls map[int][]string) {
	keys := make([]int, 0, len(cycls))
	for k := range cycls {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	f, err := os.Create("./reportfile/2")
	if err != nil {
		fmt.Println(err)
		panic("err create file")
	}
	defer f.Close()

	for _, k := range keys {
		for _, l := range cycls[k] {
			fmt.Fprintln(f, l)
		}
	}
}
