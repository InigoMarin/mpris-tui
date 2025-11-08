# MPRIS TUI

A simple terminal application (TUI) written in Go to control MPRIS-compatible media players on Linux via `playerctl`.

## Features

- **Automatic Player Detection:** Detects and lists all available media players on startup.
- **Player Selection:** Allows you to select a player from the list to control it.
- **Playback Controls:** Offers basic controls like Play/Pause, Stop, Next, and Previous.
- **Intuitive Interface:** Designed to be easy to use, even on small screens or through an SSH connection from a mobile device.
- **Lightweight and Fast:** Being written in Go, the application is a single binary with hardly any external dependencies, beyond `playerctl`.

## Requirements

- **Go:** Required to compile the application.
- **playerctl:** Required for player detection and control. Make sure it is installed on your system.

## Installation and Usage

1.  **Clone or download the repository.**
2.  **Navigate to the project directory:**
    ```bash
    cd mpris-tui
    ```
3.  **Run the application directly:**
    ```bash
    go run .
    ```
4.  **(Optional) Compile the binary:**
    To create a single executable file, you can compile the application:
    ```bash
    go build .
    ```
    This will generate a binary named `mpris-tui` that you can move to any directory in your `$PATH` for global access (e.g., `/usr/local/bin`).

## Makefile Usage

This project includes a `Makefile` to simplify common tasks.

- **`make build`**: Compiles the application and creates the `mpris-tui` binary.
- **`make install`**: Compiles the application and installs the binary to `/usr/local/bin`. This may require `sudo`.
- **`make clean`**: Removes the compiled binary.
- **`make uninstall`**: Removes the installed binary from `/usr/local/bin`. This may require `sudo`.

## Controls

### Player Selection View

- **Up/Down Arrows:** Navigate through the list of players.
- **Enter:** Select a player and switch to the control view.
- **q / Ctrl+c:** Exit the application.

### Control View

- **p:** Play / Pause
- **s:** Stop
- **n:** Next
- **v:** Previous
- **b:** Go back to the player selection list.
- **q / Ctrl+c:** Exit the application.

## Contributing

Contributions are welcome! If you have any ideas, suggestions, or bug reports, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.