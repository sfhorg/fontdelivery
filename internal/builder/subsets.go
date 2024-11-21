package builder

import (
	"fmt"
	"os"
	"strings"
)

type UnicodeRange struct {
	Start, End rune
}

var subsetRanges = map[string][]UnicodeRange{
	"latin": {
		{0x0000, 0x00FF},
		{0x0131, 0x0131},
		{0x0152, 0x0153},
		{0x02BB, 0x02BC},
		{0x02C6, 0x02C6},
		{0x02DA, 0x02DA},
		{0x02DC, 0x02DC},
		{0x0304, 0x0304},
		{0x0308, 0x0308},
		{0x0329, 0x0329},
		{0x2000, 0x206F},
		{0x2074, 0x2074},
		{0x20AC, 0x20AC},
		{0x2122, 0x2122},
		{0x2191, 0x2191},
		{0x2193, 0x2193},
		{0x2212, 0x2212},
		{0x2215, 0x2215},
		{0xFEFF, 0xFEFF},
		{0xFFFD, 0xFFFD},
	},
	"latin-ext": {
		{0x0100, 0x02BA},
		{0x02BD, 0x02C5},
		{0x02C7, 0x02CC},
		{0x02CE, 0x02D7},
		{0x02DD, 0x02FF},
		{0x0304, 0x0304},
		{0x0308, 0x0308},
		{0x0329, 0x0329},
		{0x1D00, 0x1DBF},
		{0x1E00, 0x1E9F},
		{0x1EF2, 0x1EFF},
		{0x2020, 0x2020},
		{0x20A0, 0x20AB},
		{0x20AD, 0x20C0},
		{0x2113, 0x2113},
		{0x2C60, 0x2C7F},
		{0xA720, 0xA7FF},
	},
	"vietnamese": {
		{0x0102, 0x0103},
		{0x0110, 0x0111},
		{0x0128, 0x0129},
		{0x0168, 0x0169},
		{0x01A0, 0x01A1},
		{0x01AF, 0x01B0},
		{0x0300, 0x0301},
		{0x0303, 0x0304},
		{0x0308, 0x0309},
		{0x0323, 0x0323},
		{0x0329, 0x0329},
		{0x1EA0, 0x1EF9},
		{0x20AB, 0x20AB},
	},
}

// Takes a slice of Unicode ranges and writes a file on this format:
// 0000-00FF
// 0131
// 0152-0153
// ...
func WriteHarfbuzzFile(fileName string, unicodeRanges []UnicodeRange) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, r := range unicodeRanges {
		var line string
		if r.Start == r.End {
			line = fmt.Sprintf("%04X\n", r.Start)
		} else {
			line = fmt.Sprintf("%04X-%04X\n", r.Start, r.End)
		}
		if _, err := file.WriteString(line); err != nil {
			return err
		}
	}
	return nil
}

// Takes a slice of Unicode ranges and creates a string on CSS format:
// "U+0000-00FF, U+0131, U+0152-0153, [...]"
func WriteCSSRangeString(unicodeRanges []UnicodeRange) string {
	var cssRanges []string
	for _, r := range unicodeRanges {
		if r.Start == r.End {
			cssRanges = append(cssRanges, fmt.Sprintf("U+%04X", r.Start))
		} else {
			cssRanges = append(cssRanges, fmt.Sprintf("U+%04X-%04X", r.Start, r.End))
		}
	}
	return strings.Join(cssRanges, ", ")
}