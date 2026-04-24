# BlueFrame SDK Extension

A minimal VS Code extension that:

1. Extracts the bundled `field-sdk.tgz` (`tar -xzvf`)
2. Runs `./blueframe-sdk-manager.sh -e -d sdk-output-dir`
3. `cd`s into `sdk-output-dir`
4. Builds and launches the devcontainer via the Dev Containers extension

## Setup

1. Drop your `field-sdk.tgz` into this folder (next to `extension.js`).
2. Package: `npm install && npm run package`
3. Install the resulting `.vsix` in VS Code: **Extensions → ⋯ → Install from VSIX…**

## Usage

Open the Command Palette (`Ctrl/Cmd+Shift+P`) and run:

```
BlueFrame SDK: Launch
```

All output streams to a dedicated **BlueFrame SDK** terminal panel.

## Requirements

- The **Dev Containers** extension (`ms-vscode-remote.remote-containers`) must be installed for automatic devcontainer launch.
- Docker must be running.
- The extension falls back to opening `sdk-output-dir` as a plain folder if Dev Containers is absent.
