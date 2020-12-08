package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

type XmasView struct {
	vecty.Core
}

const (
	token = "1502022722252a011e7f76386100343d083470012f000f6b6f7e29071a192c18"
)

func (x *XmasView) GetToken() string {
	output := ""
	K := "XM"
	fun, _ := hex.DecodeString(token)
	kk := "S"
	KKKKKK := "A"
	K = K + KKKKKK + kk
	for i := 0; i < len(fun); i++ {
		output += string(fun[i] ^ K[i%len(K)])
	}
	return output
}

// xmasfunforeveryone
func (x *XmasView) CheckLogin(input string) bool {
	if len(input) == 18 && input[7] == 'f' && input[len(input)-1] == 'e' && input[0] == 'x' && input[11] == 'v' && input[1] == 'm' && strings.Contains(input, "fun") && strings.HasPrefix(input, "xmas") && input[10] == 'e' && strings.HasSuffix(input, "one") {
		return true
	}
	return false
}

var LoggedIn = false
var Status = ""
var Messages = []string{
	"NO NO NO",
	"WHAAAT?!",
	"YOU ARE NOT SANTA!!!!",
	"HAHAHAHA",
	"try again....",
	"Rudolf is best",
	"funfunfun... but not for you",
	"Wrong password",
	"WHAT ARE YOU DOING?!",
	"Maybe next time",
	"Not your day, huh",
}

func (x *XmasView) Render() vecty.ComponentOrHTML {
	var login vecty.ComponentOrHTML
	if LoggedIn {
		fmt.Printf("Logged in rendering")
		login = elem.Div(
			vecty.Markup(
				vecty.Style("border", "25px dashed #235E6F"),
				vecty.Style("border-radius", "50%"),
				vecty.Style("font-size", "100px"),
			),
			elem.Paragraph(
				vecty.Text("You found the password, here is your token! "+x.GetToken()),
			),
		)
	} else {
		login = elem.Div(
			elem.Paragraph(
				vecty.Markup(
					vecty.Style("font-size", "66px"),
					vecty.Style("color", "yellow"),
				),
				vecty.Text(
					Status,
				),
			),
		)
	}

	return elem.Body(
		vecty.Markup(
			vecty.Style("background-color", "#bb2528"),
			vecty.Style("color", "white"),
			vecty.Style("text-align", "center"),
			vecty.Style("font-family", "Cursive"),
			vecty.Style("font-weight", "bolder"),
		),
		elem.Div(
			vecty.Markup(
				prop.ID("content"),
				vecty.Style("border", "33px solid #34A65F"),
			),
			elem.Paragraph(
				vecty.Markup(
					vecty.Style("font-size", "65px"),
				),
				vecty.Text("Welcome back, Santa!\n\nPlease provide your password, to get access to the secret XMAS tokens:"),
			),
			elem.Form(
				elem.Input(
					vecty.Markup(
						event.Input(func(event *vecty.Event) {
							val := event.Target.Get("value").String()
							if x.CheckLogin(val) {
								fmt.Println("Login successful")
								LoggedIn = true
								Status = ""
							} else {
								LoggedIn = false
								rand.Seed(time.Now().Unix())
								Status = Messages[rand.Intn(len(Messages))]
							}
							vecty.Rerender(x)
							fmt.Printf("Got value: %s\n", val)
						}),
					),
				),
			),
			login,
		),
	)
}

func main() {
	vecty.SetTitle("XMAS 2020")
	main := &XmasView{}
	vecty.RenderBody(main)
}
