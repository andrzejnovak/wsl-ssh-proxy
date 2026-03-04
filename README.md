# WSL SSH Proxy

A tiny Go program that proxies SSH calls through WSL, compiled as a native Windows `.exe`. This is needed because VS Code / Antigravity's Remote SSH extension uses Node.js `spawn()` which cannot execute `.bat` or `.cmd` script files — only real `.exe` executables.

## Why?

Using WSL's SSH client (instead of Windows native OpenSSH) gives you access to:
- Complex `ProxyCommand` / `ProxyJump` configurations using Unix tools
- SSH agent forwarding via Unix sockets
- Your Linux SSH config (`~/.ssh/config`) with all its host entries

## How It Works

The executable simply forwards all arguments to `wsl.exe ssh`:

```
ssh.exe -v -T submit  →  wsl.exe ssh -v -T submit
```

Key build flags:
- **`-H windowsgui`** — GUI subsystem, so the exe itself doesn't create a console window
- **`CREATE_NO_WINDOW`** (0x08000000) — prevents `wsl.exe` from spawning a visible console window either

## Build

Requires Go (e.g. `pixi global install go`). Cross-compile from Linux/WSL:

```bash
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -H windowsgui" -o ssh.exe main.go
```

## Install

1. Copy `ssh.exe` to `C:\Users\<username>\.ssh\ssh.exe`

2. In your Antigravity / VS Code `settings.json`:
   ```json
   "remote.SSH.path": "C:\\Users\\<username>\\.ssh\\ssh.exe",
   ```

3. Verify from PowerShell:
   ```powershell
   & "C:\Users\<username>\.ssh\ssh.exe" -V
   # Should print the WSL OpenSSH version
   ```

4. Reload the editor window and connect to your remote host.
