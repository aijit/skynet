package config

import (
	_ "embed"
	libs "github.com/vadv/gopher-lua-libs"
	lua "github.com/yuin/gopher-lua"
	"os"
)

//go:embed include.lua
var load_config string

type SkynetEnv struct {
	L *lua.LState
}

func NewSkynetEnv() (e *SkynetEnv) {
	e = &SkynetEnv{
		L: lua.NewState(),
	}
	libs.Preload(e.L)
	return
}

func (e *SkynetEnv) Load(from string) (err error) {
	L := e.L
	if err = L.DoString(load_config); err != nil {
		return
	}
	// 加载 lua 文件/代码
	if fd, fault := os.Stat(from); fault == nil && !fd.IsDir() {
		if err = L.DoFile(from); err != nil {
			return
		}
	} else if err = L.DoString(from); err != nil {
		return
	}
	return
}

func (e *SkynetEnv) Close() {
	if e.L != nil {
		e.L.Close()
		e.L = nil
	}
}

func (e *SkynetEnv) String(k string, opt ...string) string {
	v := e.L.GetGlobal(k)
	if ret, ok := v.(lua.LString); ok {
		return string(ret)
	}
	if len(opt) > 0 && len(opt[0]) > 0 {
		return opt[0]
	}
	return ""
}

func (e *SkynetEnv) Bool(k string, opt ...bool) bool {
	v := e.L.GetGlobal(k)
	if ret, ok := v.(lua.LBool); ok {
		return bool(ret)
	}
	return len(opt) > 0 && opt[0]
}

func (e *SkynetEnv) Int(k string, opt ...int) int {
	v := e.L.GetGlobal(k)
	if ret, ok := v.(lua.LNumber); ok {
		return int(ret)
	}
	if len(opt) > 0 {
		return opt[0]
	}
	return 0
}
