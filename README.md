# Lotus Programming Language

A simple and intuitive programming language.

## About

Lotus is an interpreted programming language that combines syntactic simplicity with powerful features. Designed to be easy to learn and use while maintaining the ability to express complex logic.

Inspired by languages like Go and Javascript, Lotus aims to combine the best of different paradigms in an accessible and expressive syntax.

## Features

- **Clean and intuitive syntax**
- **Dynamic typing**
- **First-class functions**
- **Native data structures**
- **Built-in functions**

## Installation

### Linux: Ubuntu 24.4.

To install Lotus on Linux, run the following command in your terminal:

```sh
curl -fsSL https://raw.githubusercontent.com/mateusprt/lotus/main/install.sh | bash
```

After installation, restart your terminal. . To check the version installed run ```lotus --version```.

You can also download the binary directly:

- [Download for Linux](https://github.com/mateusprt/lotus/releases/download/1/lotus)

### Windows 10

1. Open PowerShell as administrator.
2. Run the following commands:

```powershell
Invoke-WebRequest -Uri "https://github.com/mateusprt/lotus/releases/download/1/lotus.exe" -OutFile "$HOME\lotus.exe"
$lotusBin = "$HOME\lotus"
New-Item -ItemType Directory -Force -Path $lotusBin
Move-Item "$HOME\lotus.exe" "$lotusBin\lotus.exe" -Force
[Environment]::SetEnvironmentVariable("Path", $env:Path + ";$lotusBin", [EnvironmentVariableTarget]::User)
```

After installation, restart your terminal. To check the version installed run ```lotus.exe --version```.

You can also download the binary directly:

- [Download for Windows](https://github.com/mateusprt/lotus/releases/download/1/lotus.exe)

Move the binary to a folder of your choice (e.g., `C:\lotus`) and add this folder to your system PATH to use Lotus from any terminal.

## Documentation

Comprehensive documentation and language reference are available at:
[Lotus Documentation](https://mateusprt.github.io/lotus-documentation/)
