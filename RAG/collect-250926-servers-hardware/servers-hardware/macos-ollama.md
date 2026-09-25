---
id: collect-250926-servers-hardware/servers-hardware/macos-ollama
title: "macos-ollama"
domain: servers-hardware
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["gpu"]
source: docs/RAG/clean4/macos-ollama.md
source_anchor: ""
source_lines: [1, 20]
sha256: c9a86e9d4d0e65ebbfa2d359b9fe6af5a28e3edb43375a65432d535a413e629e
---

# macos-ollama

## System Requirements

- MacOS Sonoma (v14) or newer
- Apple M series (CPU and GPU support) or x86 (CPU only)

## Filesystem Requirements

The preferred method of installation is to mount the`ollama.dmg` and drag-and-drop the Ollama application to the system-wide `Applications` folder.  Upon startup, the Ollama app will verify the `ollama` CLI is present in your PATH, and if not detected, will prompt for permission to create a link in `/usr/local/bin`
Once you’ve installed Ollama, you’ll need additional space for storing the Large Language models, which can be tens to hundreds of GB in size.  If your home directory doesn’t have enough space, you can change where the binaries are installed, and where the models are stored.
### Changing Install Location

To install the Ollama application somewhere other than`Applications`, place the Ollama application in the desired location, and ensure the CLI `Ollama.app/Contents/Resources/ollama` or a sym-link to the CLI can be found in your path.  Upon first start decline the “Move to Applications?” request.
## Troubleshooting

Ollama on MacOS stores files in a few different locations.
- `~/.ollama` contains models and configuration
- `~/.ollama/logs` contains logs
  - *app.log* contains most recent logs from the GUI application
  - *server.log* contains the most recent server logs
- `<install location>/Ollama.app/Contents/Resources/ollama` the CLI binary
