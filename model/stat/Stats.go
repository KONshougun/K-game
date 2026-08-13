package stat

import (
	"strconv"
	"strings"

	"gioco/game/comm"
)

type Stats [StatCount]float32

func (s Stats) String() string {
	var str strings.Builder
	for i := range StatCount {
		str.WriteString("\n")
		str.WriteString(comm.Yellow.Write(StatName.String(i)))
		str.WriteString(" : ")
		str.WriteString(strconv.FormatFloat(float64(s[i]), 'f', -1, 32))
	}
	return str.String()
}

func (s Stats) Len() uint8 {
	var count uint8 = 0
	for _, v := range s {
		if v != 0 {
			count++
		}
	}
	return count
}

func (s Stats) ForEach(f func(StatName, float32)) {
	for i, v := range s {
		if v != 0 {
			f(StatName(i), v)
		}
	}
}
