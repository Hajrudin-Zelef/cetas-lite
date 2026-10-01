---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/aeon-7-unifi-ai-network-management-blob-head-readme-md-63fc2239-2
title: "OpenClaw target"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "kill switch"]
source: docs/RAG/collect-261001-unifi-ubiquiti/aeon-7-unifi-ai-network-management-blob-head-readme-md-63fc2239.md
source_anchor: ""
source_lines: [131, 242]
sha256: da7911ac6795ab00430e59928fd0f33c3f6384a0d717cb2c2116dcc954a8ff4a
---

# OpenClaw target

| unifi-on-openclaw /unifi-on-hermes | Re-enable local UniFi access by restoring the target unifi.env . | 
| unifi-backup-openclaw /unifi-backup-hermes | Create a secure timestamped config snapshot before major changes. | 
| unifi-restore-openclaw /unifi-restore-hermes | Choose an available backup newest-first and run a dry-run restore plan. | 
Setup/status/break-glass commands are for the owner/operator. Backup and restore are included in the AI skill because they are part of safe change management; restore still requires explicit user selection and confirmation.
If a key may be compromised or you want to immediately stop the agent from using UniFi:
unifi-off-openclaw
This is a local kill switch. It does not revoke the real key in UniFi.
Then revoke the key in the UniFi UI:
- Open your UDM Pro / Cloud Gateway web console.
- Open Network.
- Go to Settings > Control Plane > Integrations.
- Find the API key/integration used by the agent.
- Delete/revoke/remove it.
- If legacy username/password was used, disable that service user, rotate the password, or remove its UniFi permissions.
After issuing a new key:
unifi-setup-openclaw
unifi-on-openclaw
unifi-status-openclaw
Use the local Network API key for device/client/site operations on your own console.
- 
Browse to your UDM Pro or Cloud Gateway: https://<UDM-Pro-IP-or-hostname>
- 
Sign in with an account allowed to manage integrations.
- 
Open the Network application.
- 
Open Settings.
- 
Open Control Plane.
- 
Open Integrations.
- 
Generate/create an API key.
- 
Copy it once and store it with the setup alias for your target: unifi-setup-openclaw # OpenClaw unifi-setup-hermes # Hermes
- 
Verify: unifi-status-openclaw # OpenClaw python3 ~/.openclaw/workspace/skills/unifi-api/scripts/unifi_api.py network-get /sites # Hermes python3 ~/.hermes/skills/unifi-api/scripts/unifi_api.py network-get /sites
Default local prefix:
UNIFI_NETWORK_PREFIX=/proxy/network/integration/v1
If every local official endpoint returns 404 while host/auth are correct, try the pluralized prefix some controller/docs revisions use:
UNIFI_NETWORK_PREFIX=/proxy/network/integrations/v1
Handing an AI API access to your network is powerful. Treat it like giving a junior network admin a constrained service account and a very detailed runbook.
Minimum safety baseline:
- Use a dedicated UniFi service/integration key.
- Prefer read-only or least-privilege access when UniFi exposes it.
- Keep write-capable keys out of chat history.
- Store keys only in the selected private env file, ~/.openclaw/unifi.env for OpenClaw or~/.hermes/unifi.env for Hermes, with mode0600 .
- Use unifi-backup-openclaw orunifi-backup-hermes before major changes, matching the agent target.
- Use official APIs before legacy private endpoints.
- Require exact object verification before writes: site, name, MAC/IP/model/id.
- Avoid changes that could sever the current management path unless you have a rollback path.
- Use unifi-off immediately if anything looks wrong.
- Revoke real keys in UniFi after any suspected compromise.
Recommended permission split:
| Key/account | Purpose | Access level | 
|---|---|---|
| Read-only API key | Daily monitoring, inventory, diagnostics | Read-only if available | 
| Write-capable API key | Explicit maintenance windows and targeted changes | Narrowest available write scope | 
| Legacy service account | Compatibility fallback only | Disabled unless needed | 
Never paste API keys, cookies, CSRF tokens, passwords, restore logs, or full raw config dumps into chat.
Create backup:
unifi-backup-openclaw  # OpenClaw
unifi-backup-hermes    # Hermes
Backups are stored under the selected agent home, for example:
~/.openclaw/unifi-backups/   # OpenClaw target
~/.hermes/unifi-backups/     # Hermes target
Example:
~/.openclaw/unifi-backups/backup-20260509-192255/
Backups use 0700 directories and 0600 files. They may contain sensitive network data including SSIDs, firewall rules, VLANs, port profiles, and possibly Wi-Fi credentials.
Dry-run restore:
unifi-restore-openclaw --backup ~/.openclaw/unifi-backups/backup-YYYYMMDD-HHMMSS
unifi-restore-hermes --backup ~/.hermes/unifi-backups/backup-YYYYMMDD-HHMMSS
Apply restore after explicit confirmation:
unifi-restore-openclaw --backup ~/.openclaw/unifi-backups/backup-YYYYMMDD-HHMMSS --apply
unifi-restore-hermes --backup ~/.hermes/unifi-backups/backup-YYYYMMDD-HHMMSS --apply
For non-interactive agent execution after explicit user confirmation, the same confirmation phrase can be supplied as a flag:
unifi-restore-openclaw --backup ~/.openclaw/unifi-backups/backup-YYYYMMDD-HHMMSS --apply --confirm 'RESTORE UNIFI CONFIG'
unifi-restore-hermes --backup ~/.hermes/unifi-backups/backup-YYYYMMDD-HHMMSS --apply --confirm 'RESTORE UNIFI CONFIG'
The restore tool requires this exact confirmation phrase before applying:
RESTORE UNIFI CONFIG
It will not delete current objects that are absent from the backup unless explicitly requested:
unifi-restore-openclaw --apply --delete-extra
unifi-restore-hermes --apply --delete-extra
Use --delete-extra only for intentional rollback to an older full state.
The helper exposes generic commands so the agent can call exact endpoints without guessing curl syntax.
Base:
https://api.ui.com/v1
| Command | Endpoint | What it does | AI use cases | 
|---|---|---|---|
| site-manager-get /hosts | GET /hosts | Lists UniFi OS hosts/consoles visible to the API key. | Fleet health, find the right console, identify offline hosts. | 
| site-manager-get /sites | GET /sites | Lists sites visible to the key. | Site inventory, mapping site names to controllers. | 
| site-manager-get /devices | GET /devices | Lists cloud-visible devices where available. | Global firmware/update/online summary. | 
| site-manager-post <path> | POST <path> | Generic cloud write/action route, if supported by the official docs. | Rare; use only with docs-confirmed path and explicit intent. | 
Base pattern:
https://<controller>/proxy/network/integration/v1
| Command | Endpoint pattern | What it does | AI use cases | 
|---|---|---|---|
| network-get /sites | GET /sites | Lists local Network sites. | Select exact site id before drilling into clients/devices. | 
| network-get /sites/<site_id>/devices | GET /sites/{site_id}/devices | Lists UniFi Network devices. | AP/switch/gateway health, firmware, adoption state, uplink checks. | 
| network-get /sites/<site_id>/clients | GET /sites/{site_id}/clients | Lists clients. | Client diagnostics, unknown device review, Wi-Fi quality checks. | 
| network-post <path> | POST <path> | Generic official local action route. | Guest authorization or supported actions from local docs. Requires explicit write acknowledgement. | 
Use only when the official APIs do not expose the needed function.
| Command | Endpoint pattern | What it does | AI use cases | 
|---|---|---|---|
| legacy-sites | GET /api/self/sites or/proxy/network/api/self/sites | Lists legacy sites. | Find legacy site names such as default . | 
| legacy-clients --site default | GET /api/s/<site>/stat/sta | Lists clients. | Deep client stats, RSSI, rates, AP association, legacy fields. | 
| legacy-devices --site default | GET /api/s/<site>/stat/device | Lists devices. | Detailed AP/switch/gateway stats and port/radio tables. | 
| legacy-get /api/s/<site>/rest/wlanconf | WLAN config read | Reads SSID/Wi-Fi config. | SSID/security/VLAN mapping audit. | 
| legacy-get /api/s/<site>/rest/networkconf | Network config read | Reads networks/VLANs/DHCP config. | VLAN/DHCP/DNS audit. | 
| legacy-get /api/s/<site>/rest/firewallrule | Firewall rules read | Reads firewall policy. | Security posture review and drift detection. | 
| legacy-get /api/s/<site>/rest/firewallgroup | Firewall groups read | Reads address/port groups. | Explain firewall rule targets and identify stale groups. | 
| legacy-get /api/s/<site>/rest/portconf | Port profiles read | Reads switch port profiles. | Validate VLAN/trunk/access profiles. | 
