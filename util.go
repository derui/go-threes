package main

import (
	"github.com/derui/go-threes/libthrees"
)

func GetAllCombPos() (ret []libthrees.Pos) {

	const size = libthrees.BASEMENT_SIZE
	
	for i := 0; i < size; i++ {
		for j := 0; j < size;j++ {
			ret = append(ret, libthrees.GetPos(i, j))
		}
	}

	return
}
