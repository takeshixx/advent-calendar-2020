package main

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
)

func runCode(code string, w http.ResponseWriter) {
	vm := goja.New()
	myPrinter := console.PrinterFunc(func(s string) { fmt.Fprintln(w, s) })
	registry := require.NewRegistry()
	registry.Enable(vm)

	registry.RegisterNativeModule("console", console.RequireWithPrinter(myPrinter))
	vm.Set("console", require.Require(vm, "console"))

	registry.RegisterNativeModule("OS", requireOS())
	vm.Set("OS", require.Require(vm, "OS"))

	registry.RegisterNativeModule("secrets", requireSecrets())
	vm.Set("secrets", require.Require(vm, "secrets"))

	vm.Set("xmas", func() goja.Value { return vm.ToValue(xmas()) })

	time.AfterFunc(200*time.Millisecond, func() {
		vm.Interrupt("halt")
	})

	_, err := vm.RunString(code)
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}
}

func requireOS() require.ModuleLoader {
	return func(runtime *goja.Runtime, module *goja.Object) {
		osFuncs := &osFunctions{
			vm: runtime,
		}

		o := module.Get("exports").(*goja.Object)
		o.Set("exec", osFuncs.exec)
		o.Set("ls", osFuncs.ls)
		o.Set("readFile", osFuncs.readFile)
		o.Set("writeFile", osFuncs.writeFile)
	}
}

type osFunctions struct {
	vm *goja.Runtime
}

func (o osFunctions) exec(call goja.FunctionCall) goja.Value {
	argLen := len(call.Arguments)
	if argLen < 1 {
		return o.vm.NewGoError(fmt.Errorf("exec requires cmd argument"))
	}
	args := make([]string, argLen-1)
	for i := 0; i < argLen-1; i++ {
		args[i] = call.Arguments[i+1].String()
	}
	cmd := exec.Command(call.Arguments[0].String(), args...)

	stdoutStderr, err := cmd.CombinedOutput()

	var result string
	if err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if ok {
			result = string(stdoutStderr)
			result = fmt.Sprintf("%s\nProcess failed with exit code: %d", result, exitErr.ExitCode())
		} else {
			return o.vm.NewGoError(err)
		}
	} else {
		result = string(stdoutStderr)
	}

	return o.vm.ToValue(result)
}

func (o osFunctions) ls(call goja.FunctionCall) goja.Value {
	argLen := len(call.Arguments)
	if argLen < 1 {
		return o.vm.NewGoError(fmt.Errorf("ls requires path argument"))
	}

	files, err := ioutil.ReadDir(call.Arguments[0].String())
	if err != nil {
		return o.vm.NewGoError(err)
	}

	var b strings.Builder
	for _, f := range files {
		fmt.Fprintln(&b, f.Name())
	}

	return o.vm.ToValue(b.String())
}

func (o osFunctions) readFile(call goja.FunctionCall) goja.Value {
	argLen := len(call.Arguments)
	if argLen < 1 {
		return o.vm.NewGoError(fmt.Errorf("readFile requires path argument"))
	}

	content, err := ioutil.ReadFile(call.Arguments[0].String())
	if err != nil {
		return o.vm.NewGoError(err)
	}

	var result string
	if len(call.Arguments) > 1 && call.Arguments[1].ToBoolean() {
		result = b64.StdEncoding.EncodeToString([]byte(content))
	} else {
		result = string(content)
	}
	return o.vm.ToValue(result)
}

func (o osFunctions) writeFile(call goja.FunctionCall) goja.Value {
	argLen := len(call.Arguments)
	if argLen < 2 {
		return o.vm.NewGoError(fmt.Errorf("readFile requires path and content argument"))
	}

	err := ioutil.WriteFile(call.Arguments[0].String(), []byte(call.Arguments[1].String()), 0644)
	if err != nil {
		return o.vm.NewGoError(err)
	}

	return nil
}

func xmas() string {
	const xmas = "🎅🎄☃️🦌🤶🎁🎍🧣"
	return xmas
}

var secretStore = map[string]string{
	"cryptorPassword": "🎅🎄🦌🤶🎁🎍🧣🎅",
	"rootPassword":    "toor",
}

func requireSecrets() require.ModuleLoader {
	return func(runtime *goja.Runtime, module *goja.Object) {
		secFuncs := &secretsFunctions{
			vm: runtime,
		}

		o := module.Get("exports").(*goja.Object)
		o.Set("unsealed", runtime.ToValue(true))
		o.Set("get", secFuncs.get)
		o.Set("list", secFuncs.list)
	}
}

type secretsFunctions struct {
	vm *goja.Runtime
}

func (s secretsFunctions) list() goja.Value {
	keys := make([]string, 0, len(secretStore))
	for key := range secretStore {
		keys = append(keys, key)
	}
	return s.vm.ToValue(keys)
}

func (s secretsFunctions) get(call goja.FunctionCall) goja.Value {
	argLen := len(call.Arguments)
	if argLen < 1 {
		return s.vm.NewGoError(fmt.Errorf("get requires name argument"))
	}
	name := call.Argument(0).String()
	value, ok := secretStore[name]
	if !ok {
		return s.vm.NewGoError(fmt.Errorf("could find secret with name %s", name))
	}
	return s.vm.ToValue(value)
}
