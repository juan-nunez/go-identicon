package main

import "go-identicon"

func main() {
	icon := identicon.NewIdenticon("example@gmail.com")
	icon.ToImage()
}