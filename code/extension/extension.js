const vscode = require('vscode');
const path = require('path');
const fs = require('fs');

function activate(context) {
  const workDir = () => vscode.workspace.workspaceFolders
    ? vscode.workspace.workspaceFolders[0].uri.fsPath
    : require('os').homedir();

  // Command 1: Extract and run the SDK manager
  context.subscriptions.push(
    vscode.commands.registerCommand('blueframe-sdk.launch', () => {
      const tgzPath = path.join(context.extensionPath, 'field-sdk.tgz');

      if (!fs.existsSync(tgzPath)) {
        vscode.window.showErrorMessage(`field-sdk.tgz not found at ${tgzPath}. Re-install the extension.`);
        return;
      }

      const cwd = workDir();
      const scriptPath = path.join(cwd, 'blueframe-sdk-manager.sh');
      const terminal = vscode.window.createTerminal({ name: 'BlueFrame SDK', cwd });
      terminal.show();

      terminal.sendText(`tar -xzvf "${tgzPath}"`);
      terminal.sendText(`chmod +x "${scriptPath}" && "${scriptPath}" -e -d sdk-output-dir`);
    })
  );

  // Command 2: Open the devcontainer
  context.subscriptions.push(
    vscode.commands.registerCommand('blueframe-sdk.openDevcontainer', () => {
      const outputDir = path.join(workDir(), 'sdk-output-dir');

      if (!fs.existsSync(outputDir)) {
        vscode.window.showErrorMessage(`sdk-output-dir not found at ${outputDir}. Run "BlueFrame SDK: Launch" first.`);
        return;
      }

      const folderUri = vscode.Uri.file(outputDir);

      if (vscode.extensions.getExtension('ms-vscode-remote.remote-containers')) {
        vscode.commands.executeCommand('remote-containers.openFolder', folderUri);
      } else {
        vscode.commands.executeCommand('vscode.openFolder', folderUri, { forceNewWindow: false });
        vscode.window.showWarningMessage('Dev Containers extension not found. Opened sdk-output-dir as a plain folder.');
      }
    })
  );
}

function deactivate() {}

module.exports = { activate, deactivate };
