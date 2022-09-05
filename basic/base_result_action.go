package basic

import "fmt"

type IResultAction interface {
	OnFirstHash(info ResultInfo)
	OnHashChanged(info ResultInfo)
	OnHashUnchanged(info ResultInfo)
	OnError(info ResultInfo, err error)
}

type ResultInfo struct {
	Filepath string
	Hash     string
	Url      string
}

type DebugResultAction struct{}

func (DebugResultAction) OnFirstHash(info ResultInfo) {
	fmt.Printf("Saved initial hash \"%s\" for %s\n at %q", info.Hash, info.Url, info.Filepath)
}

func (DebugResultAction) OnHashChanged(info ResultInfo) {
	fmt.Printf("There is changes at %s\n", info.Url)
}

func (DebugResultAction) OnHashUnchanged(info ResultInfo) {
	fmt.Printf("There is no changes at %s\n", info.Url)
}

func (DebugResultAction) OnError(info ResultInfo, err error) {
	fmt.Printf("Caught error while handling %s: %v\n", info.Url, err)
}
