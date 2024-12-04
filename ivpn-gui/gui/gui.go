package gui

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"git.ana/dorbmon/ivpn-gui/engine"
	"git.ana/dorbmon/ivpn-gui/log"
	imgui "github.com/AllenDang/cimgui-go"
)

var settingWindowOpened = true
var logWindowOpened = true
var runningStatus = false
var key = new(engine.Key)
var insideSchool = true

func getCFGFilePath() string {
	p, err := os.Executable()
	if err != nil {
		log.Fatalf(err.Error())
	}
	dir := filepath.Dir(p)
	cfgPath := path.Join(dir, "config.json")
	return cfgPath
}
func init() {
	key.Socks5Addr = ":1234"
	key.Device = "xjtuana"
	key.LogLevel = "info"
	key.HttpProxyAddr = ":1080"
	cfgPath := getCFGFilePath()
	cf, err := os.Open(cfgPath)
	if err != nil {
		cf, err = os.Create(cfgPath)
		if err != nil {
			panic(err)
		}
		str, _ := json.Marshal(*key)
		cf.Write(str)
		cf.Seek(0, io.SeekStart)
	}
	defer cf.Close()
	b, err := io.ReadAll(cf)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if err := json.Unmarshal(b, key); err != nil {
		log.Fatalf(err.Error())
	}
}
func shutDownEngine() {
	if runningStatus {
		engine.Stop()
		runningStatus = false
	}
}
func beginEngine() {
	if insideSchool {
		key.Proxy = fmt.Sprintf("wss://%s@ana.xjtu.edu.cn/ivp/", key.Token)
	} else {
		key.Proxy = fmt.Sprintf("wss://%s@xjtuana.com/ivp/", key.Token)
	}
	engine.Insert(key)
	engine.Start()
	runningStatus = true
}
func applyConfig() {
	cfgPath := getCFGFilePath()
	fmt.Println(cfgPath)
	os.Remove(cfgPath)
	f, err := os.Create(cfgPath)
	fmt.Println(key.Token)
	if err != nil {
		log.Errorf("[GUI] SAVE CONFIG FAILED: %s", err.Error())
	} else {
		defer f.Close()
		b, _ := json.Marshal(*key)
		_, _ = f.Write(b)
	}
	shutDownEngine()
	beginEngine()
}
func loop() {
	imgui.BeginV("Setting", &settingWindowOpened, 0)
	imgui.Text("Status: ")
	imgui.SameLine()
	if runningStatus {
		imgui.TextColored(imgui.NewVec4(0, 1, 0, 1), "Running")
	} else {
		imgui.TextColored(imgui.NewVec4(1, 0, 0, 1), "Not Running")
	}
	imgui.InputTextWithHint("token", "TOKEN", &key.Token, 0, nil)
	imgui.InputTextWithHint("Device", "xjtuana", &key.Device, 0, nil)
	imgui.Checkbox("Inside School", &insideSchool)
	if imgui.Checkbox("Enable Http Proxy", &key.EnableHttpProxy) {
		if !key.EnableSocks5 {
			key.EnableHttpProxy = false
		}
	}
	imgui.PushItemWidth(90)
	imgui.InputTextWithHint("Http Proxy Listennig Port", ":1080", &key.HttpProxyAddr, 0, nil)
	imgui.Checkbox("Enable Socks5 Proxy", &key.EnableSocks5)
	imgui.InputTextWithHint("Socks5 Listennig Address(like :1234)", ":1234", &key.Socks5Addr, 0, nil)
	if imgui.Button("Apply And Run") {
		go applyConfig()
	}
	imgui.SameLine()
	if imgui.Button("Stop") {
		if runningStatus {
			go shutDownEngine()
		}
	}
	imgui.End()
	imgui.SetNextWindowSizeV(imgui.Vec2{X: 100, Y: 500}, imgui.CondOnce)
	imgui.BeginV("Log", &logWindowOpened, 0)
	//imgui.BeginChildStr("Scrolling")
	for _, item := range log.CurrentLogs {
		imgui.Text(item)
	}
	imgui.End()
}
func GUIMain() {
	if !amAdmin() {
		runMeElevated()
		return
	}
	backend := imgui.CreateBackend()
	backend.CreateWindow("XJTUANA IVPN", 1200, 900, 0)
	backend.Run(loop)
}
