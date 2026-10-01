---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/aeon-7-unifi-ai-network-management-blob-head-readme-md-63fc2239-1
title: "OpenClaw target"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents"]
source: docs/RAG/collect-261001-unifi-ubiquiti/aeon-7-unifi-ai-network-management-blob-head-readme-md-63fc2239.md
source_anchor: ""
source_lines: [1, 130]
sha256: 91d54e57264884bff1bb393caa158f87e925fde3f9d2150dadce96d1f98840a4
---

# OpenClaw target

A safety-first UniFi API skill and tooling package for AI agents that can inspect, troubleshoot, document, and carefully optimize UniFi networks.
This package gives an AI agent a practical UniFi runbook, a deterministic API helper, owner-only shell tooling, secure .env handling, backup/restore commands, and break-glass controls for revoking local access quickly.
Repository name: unifi-ai-network-management.
This repo is designed to be installed as a complete package: the skill folder is copied into the agent skills directory, helper scripts are installed into the user's local bin directory, and credentials are stored in a private env file outside the repository.
This repository packages three things:
| Component | Location | Purpose | 
|---|---|---|
| Agent skill | skill/unifi-api/ | SKILL.md package for OpenClaw (~/.openclaw/workspace/skills/unifi-api ) or Hermes (~/.hermes/skills/unifi-api ) with UniFi API operations, safety tiers, endpoint selection, and troubleshooting. | 
| API helper | skill/unifi-api/scripts/unifi_api.py | Zero-dependency Python helper for Site Manager, official local Network API, and legacy controller API. | 
| Operational scripts | scripts/ | Setup, status, break-glass disable/enable, secure backup, and dry-run/apply restore tooling. | 
The goal is not to let an AI randomly mutate your network. The goal is to give a trusted local agent enough structure to:
- inventory sites, controllers, devices, clients, SSIDs, VLANs, and firmware state
- diagnose Wi-Fi and client issues from real controller data
- summarize health and configuration drift
- identify risky states before they become outages
- prepare and verify changes with backups and rollback paths
- act on explicit, narrow instructions when write access is intentionally granted
UniFi has multiple APIs. This package deliberately separates them instead of pretending there is one universal endpoint.
| Surface | Base | Auth | Best for | 
|---|---|---|---|
| Site Manager API | https://api.ui.com/v1 | X-API-Key | Cloud/fleet/site/host overview across UniFi sites. | 
| Local Network Integration API | https://<controller>/proxy/network/integration/v1 | X-API-Key | Official local UniFi Network data and supported actions. | 
| Legacy private Network Application API | https://<controller>/api/... or/proxy/network/api/... | username/password session cookie | Compatibility actions not yet exposed by the official API. Use sparingly. | 
Official documentation:
- Ubiquiti: Getting Started with the Official UniFi API: https://help.ui.com/hc/en-us/articles/30076656117655-Getting-Started-with-UniFi-API
- UniFi Site Manager API: https://developer.ui.com/site-manager-api/
- Ubiquiti Developer portal: https://developer.ui.com/
Ubiquiti states that localized Network API documentation is available inside UniFi Network at Settings > Control Plane > Integrations. Always prefer the docs from your deployed controller version when route details differ.
Clone the repo on the machine where your agent or gateway runs.
For OpenClaw:
git clone https://github.com/AEON-7/unifi-ai-network-management.git
cd unifi-ai-network-management
./setup.sh --target openclaw
source ~/.bashrc
unifi-setup-openclaw
unifi-status-openclaw
For Hermes Agent:
git clone https://github.com/AEON-7/unifi-ai-network-management.git
cd unifi-ai-network-management
./setup.sh --target hermes
source ~/.bashrc
unifi-setup-hermes
unifi-status-hermes
hermes doctor
If your shell already had an older unifi-* alias block, open a new shell or run source ~/.bashrc after setup. The installer is safe to re-run; it backs up an existing skill directory before replacing it and preserves the existing target env file unless --force-env is used.
setup.sh installs:
# OpenClaw target
~/.openclaw/workspace/skills/unifi-api/
~/.openclaw/unifi.env
# Hermes target
~/.hermes/skills/unifi-api/
~/.hermes/unifi.env
# Both targets
~/.local/bin/unifi-api-status
~/.local/bin/unifi-api-configure
~/.local/bin/unifi-api-disable
~/.local/bin/unifi-api-enable
~/.local/bin/unifi-config-backup
~/.local/bin/unifi-config-restore
It installs explicit target commands into ~/.local/bin and ensures that directory is on your shell PATH:
# OpenClaw target
unifi-status-openclaw
unifi-setup-openclaw
unifi-off-openclaw
unifi-on-openclaw
unifi-backup-openclaw
unifi-restore-openclaw
# Hermes target
unifi-status-hermes
unifi-setup-hermes
unifi-off-hermes
unifi-on-hermes
unifi-backup-hermes
unifi-restore-hermes
Use ./setup.sh --target hermes for Hermes Agent, ./setup.sh --target openclaw for OpenClaw, ./setup.sh --no-aliases if you want files installed without modifying your shell rc/PATH, and ./setup.sh --force-env only if you intentionally want to replace the existing env file with env.example after creating a timestamped backup.
If you do not use OpenClaw or Hermes, you can still use the helper script and docs. Install the skill folder wherever your agent framework expects tool instructions.
The installer copies the canonical skill directory from:
./skill/unifi-api/
to the selected agent skill path:
~/.openclaw/workspace/skills/unifi-api/  # OpenClaw
~/.hermes/skills/unifi-api/              # Hermes
It preserves this layout because the SKILL.md references references/ and scripts/unifi_api.py by relative path:
unifi-api/
|-- SKILL.md
|-- scripts/unifi_api.py
|-- references/
`-- agents/openai.yaml
For other agent frameworks, copy the entire skill/unifi-api directory into that framework's skills/tools directory. Do not copy only SKILL.md; the helper script and reference files are part of the package.
See AGENTS.md for detailed agent-framework installation and behavior guidance.
Hermes Agent uses SKILL.md skills under ~/.hermes/skills/. The Hermes target installs this package as:
~/.hermes/skills/unifi-api/
|-- SKILL.md
|-- scripts/unifi_api.py
|-- references/
`-- agents/openai.yaml
The private credential file is stored separately at:
~/.hermes/unifi.env
After setup, open a new shell or run source ~/.bashrc, then run unifi-setup-hermes. Start a new Hermes session so it reloads the skill catalog. If Hermes is installed, hermes doctor should still pass.
This repository intentionally does not commit a real .env file. The common and safer convention is:
- Commit env.example as a template.
- Copy it to a private runtime env file.
- Keep real .env /unifi.env files ignored by git.
setup.sh handles this automatically by copying env.example to the target-private env path (~/.openclaw/unifi.env or ~/.hermes/unifi.env) when one does not already exist.
The private env file lives at ~/.openclaw/unifi.env for OpenClaw or ~/.hermes/unifi.env for Hermes.
It should be readable only by your user:
chmod 600 ~/.openclaw/unifi.env  # OpenClaw
chmod 600 ~/.hermes/unifi.env    # Hermes
Template:
# Site Manager cloud API
UNIFI_SITE_MANAGER_API_KEY=
UNIFI_SITE_MANAGER_BASE_URL=https://api.ui.com/v1
# Local UniFi Network API
UNIFI_NETWORK_BASE_URL=https://192.168.1.1
UNIFI_NETWORK_API_KEY=
UNIFI_NETWORK_PREFIX=/proxy/network/integration/v1
UNIFI_INSECURE_TLS=1
# Legacy fallback, only when official API cannot do the job
UNIFI_LEGACY_BASE_URL=https://192.168.1.1
UNIFI_USERNAME=
UNIFI_PASSWORD=
UNIFI_SITE=default
Use the interactive setup tool instead of editing by hand:
unifi-setup-openclaw  # OpenClaw
unifi-setup-hermes    # Hermes
It hides secret input, confirms keys/passwords, backs up the old env file, writes 0600, and can restart OpenClaw if installed as a user service. Hermes users should start a fresh Hermes session after setup so the skill catalog reloads.
| Command | Purpose | 
|---|---|
| unifi-status-openclaw /unifi-status-hermes | Show local access state and redacted env status for a specific target. | 
| unifi-setup-openclaw /unifi-setup-hermes | Interactive setup/update for API keys and controller URLs for a specific target. | 
| unifi-off-openclaw /unifi-off-hermes | Local break-glass disable. Renames the target unifi.env tounifi.env.disabled and restarts the gateway if present. | 
