package material_test

import (
	"gioui.org/doc"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

var (
	_ material.ButtonLayoutStyle
	_ material.ButtonStyle
)

func ExampleScreenshot() {
	var clickable widget.Clickable
	doc.Screenshot(func(gtx layout.Context, theme *material.Theme) layout.Dimensions {
		return material.Button(theme, "Click me!").Layout(gtx, &clickable)
	})
}

/*
Button Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.

Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
*/

func ExampleBasicButton() {
	doc.Example(func(run doc.Run) {
		var clickable widget.Clickable

		run(func(gtx layout.Context, theme *material.Theme) layout.Dimensions {
			return material.Button(theme, "Click me!").Layout(gtx, &clickable)
		})
	})
}

/*
## Styling

You can customize the button by modifying `material.ButtonLayoutStyle`.

Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
*/

func ExampleButtonWithStyle() {
	doc.Example(func(run doc.Run) {
		var clickable widget.Clickable

		run(func(gtx layout.Context, theme *material.Theme) layout.Dimensions {
			style := material.Button(theme, "Click me!")
			style.CornerRadius = unit.Dp(8)
			return style.Layout(gtx, &clickable)
		})
	})
}
