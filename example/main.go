package main

import "go-identicon"

func main() {
	icon := identicon.NewIdenticon("Hello world, how are you")
	icon.ToImage()
}