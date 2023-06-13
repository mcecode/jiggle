# Jiggle

An application that automatically moves your mouse around which _should_\* keep your machine awake.

![A demo showing Jiggle being started, moving the mouse, and finally being stopped by a keypress.](assets/demo.webp)

\*Should because I haven't actually tested if it can keep a machine awake long term because I don't really need to do that, and the friend that I made this for ended up buying a USB [mouse jiggler](https://en.wikipedia.org/wiki/Mouse_jiggler) before I could finish this, so she wasn't able to test it either. So yeah, it was a good learning experience in Go though.

## Packaging

You'll need the following dependencies to compile and package Jiggle:

- [Go](https://go.dev), preferably version 16 or later
- [GCC](https://gcc.gnu.org), compiled for your OS
- [Fyne](https://developer.fyne.io), you can install it by running `go install fyne.io/fyne/v2/cmd/fyne@latest`
- Any other OS-specific requirements that [Robotgo](https://github.com/go-vgo/robotgo#requirements) and [Fyne](https://developer.fyne.io/started#prerequisites) need as they're the main dependencies of the project

To package the application, you can run the [`fyne package`](https://developer.fyne.io/started/packaging) command. You can also use the [`build.sh`](build.sh) helper that I made. It automatically detects the OS that it's running under, adds flags and commands for that OS, and moves the packaged application to a `build` directory.

## Usage

Press the start button to start it, and press any keyboard key to stop it. That's it.

## Contributing

I don't really have anything planned for this project since I have no need for it, but suggestions for improvements are still welcome. I'll probably implement them if I can since it was pretty fun working on this.

## Third-party assets

- [`cursor.png`](assets/cursor.png) and [`cursor.svg`](assets/cursor.svg) are modified versions of the [Arrow, cursor, interface icon](https://www.iconfinder.com/icons/6820524/arrow_cursor_interface_ui_ux_icon) from [panca okta](https://www.iconfinder.com/pancaza)'s [Interface Solid icon pack](https://www.iconfinder.com/iconsets/interface-solid-8)

## Similar projects

- [Automatic-Mouse-Mover](https://github.com/prashantgupta24/automatic-mouse-mover), if you want a mouse jiggler for Mac, this should be your go-to; it's battle-tested and has better and more features than Jiggle.

## License

Copyright 2022-present Matthew Espino

This project is licensed under the [version 3 of the GPL or any later version](LICENSE).
