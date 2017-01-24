package main

import (
	"fmt"
	"log"
	"github.com/jroimartin/gocui"
	"strings"
)

var gui *gocui.Gui

func RunUi() {
	var err error
	gui, err = gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatalln(err)
	}
	defer gui.Close()

	gui.Cursor = true
	gui.Highlight = false
	gui.BgColor = gocui.ColorDefault

	gui.SetManagerFunc(layout)

	if err := initKeybindings(gui); err != nil {
		log.Fatalln(err)
	}

	if err := gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatalln(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("help", maxX - 23, 0, maxX - 1, 3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "↑ ↓: Seek input")
		fmt.Fprintln(v, "^W: Exit")
		v.Title = "KEY BINDINGS"
	}

	if v, err := g.SetView("main", 24, 0, maxX - 24, maxY - 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.Title = "main"
	}

	if v, err := g.SetView("stdin", 26, maxY - 2, maxX - 24, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if _, err := g.SetCurrentView("stdin"); err != nil {
			return err
		}
		v.Frame = false
		v.Wrap = true
		v.Autoscroll = true
		v.Editable = true
	}
	if v, err := g.SetView("stdin-prefix", 24, maxY - 2, 26, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if _, err := g.SetCurrentView("stdin"); err != nil {
			return err
		}
		fmt.Fprint(v, "$")
		v.Frame = false
	}

	return nil
}

func initKeybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("stdin", gocui.KeyCtrlW, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("stdin", gocui.KeyEnter, gocui.ModNone, input); err != nil {
		return err
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func input(g *gocui.Gui, v *gocui.View) error {
	str := v.Buffer()
	v.Clear()
	v.SetCursor(0, 0)
	vi, err := g.View("main");
	if err != nil {
		return err
	}
	g.Execute(func(g *gocui.Gui) error {
		vi.Autoscroll = true
		str = strings.TrimSpace(str)
		sp := strings.SplitN(str, " ", 2)
		if len(sp) == 1 {
			sp = append(sp, "")
		}
		err := exec(strings.TrimSpace(sp[0]), strings.TrimSpace(sp[1]))
		if err != nil {
			return err
		}
		return nil
	})
	return nil
}