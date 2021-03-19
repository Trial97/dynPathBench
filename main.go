package main

import (
	"math/rand"

	"github.com/cgrates/cgrates/utils"
)

func init() {
	// gen = generateRandomTemplate(noGen)
	gen = []pathInfo{
		{
			strPath: "Field1[*raw][0].Field2[0].Field3[*new]",
			data:    "1001",
		},
	}
	for i := range gen {
		gen[i].fp = utils.NewFullPath(gen[i].strPath, ".")
		gen[i].pathItems = gen[i].fp.PathItems
		for _, pt := range gen[i].fp.PathItems {
			gen[i].path = append(gen[i].path, pt.Field)
			gen[i].path = append(gen[i].path, pt.Index...)
		}
	}
}

const noGen = 10000

var generator = rand.New(rand.NewSource(42))
var gen []pathInfo

type pathInfo struct {
	pathItems utils.PathItems
	path      []string
	strPath   string
	data      string
	fp        *utils.FullPath
}

func generateRandomPath() (out []string) {
	size := generator.Intn(16) + 1
	out = make([]string, size)
	for i := 0; i < size; i++ {
		out[i] = utils.Sha1(utils.GenUUID())
	}
	return
}
func generateRandomTemplate(size int) (out []pathInfo) {
	out = make([]pathInfo, size)
	for i := 0; i < size; i++ {
		out[i].path = generateRandomPath()
		out[i].data = utils.UUIDSha1Prefix()
		out[i].pathItems = utils.NewPathItems(out[i].path)
		out[i].strPath = out[i].pathItems.String()
		// out[i].pathItems[len(out[i].pathItems)-1].Index = IntPointer(0)
		out[i].fp = &utils.FullPath{PathItems: out[i].pathItems, Path: out[i].strPath}
	}
	return
}
