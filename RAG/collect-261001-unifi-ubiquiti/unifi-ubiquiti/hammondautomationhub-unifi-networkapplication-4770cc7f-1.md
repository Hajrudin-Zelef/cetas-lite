---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hammondautomationhub-unifi-networkapplication-4770cc7f-1
title: "Fresh install"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hammondautomationhub-unifi-networkapplication-4770cc7f.md
source_anchor: ""
source_lines: [1, 118]
sha256: 414f7eeab4ac68c76d6c3c0c2a51e5f4a6eb8609e40f309c193ec750b6abebca
---

# Fresh install

This repo contains two different UniFi self-hosting paths. They are not interchangeable.
|  | UniFi OS Server (recommended by Ubiquiti) | Legacy Network Application (linuxserver.io) | 
|---|---|---|
| Script | install-unifi-os-server.sh | install-unifi-docker.sh | 
| Runtime | Official Ubiquiti installer + Podman | Docker Compose + external MongoDB | 
| Web UI | https://<host>:11443 | https://<host>:8443 | 
| Features | Full UniFi OS (Organizations, Site Manager, IdP, …) | Network controller only | 
| MongoDB | Bundled inside UOS (no separate container) | Separate mongo container (ARMv8.2 issues on Khadas) | 
| Ubiquiti support | Official self-hosting path | Community / legacy | 
If you want the updated UniFi OS Server, use install-unifi-os-server.sh — not install-unifi-docker.sh.
Production Bash installer for Ubiquiti's UniFi OS Server on Ubuntu/Debian (x86_64 and arm64, including Khadas VIM 4).
chmod +x install-unifi-os-server.sh
# Fresh install
sudo ./install-unifi-os-server.sh -y
# Migrate from native .deb and stop legacy Docker stack if present
sudo ./install-unifi-os-server.sh --migrate-from-deb --remove-legacy-docker -y
Open after install: https://<your-host-ip>:11443
Storage: UniFi OS Server’s official installer requires ≥15 GB free on /home (Ubiquiti recommends 25–40 GB total). Khadas VIM4 eMMC often does not meet this — free space or add NVMe/USB before installing.
Restore a .unf backup via Settings → System → Restore in the UOS wizard.
If you already ran install-unifi-docker.sh:
cd /root/unifi
sudo docker compose down
sudo ~/install-unifi-os-server.sh --migrate-from-deb --remove-legacy-docker -y
Then restore your .unf at https://<host>:11443.
Production-ready Bash installer and upgrader for the linuxserver.io UniFi Network Application Docker image with MongoDB, using Docker Compose.
Designed for Ubuntu/Debian hosts on x86_64 and arm64 (including boards like the Khadas VIM 4). The script auto-detects an existing install and upgrades in place while preserving data, or performs a clean fresh install when no legacy application is present.
Note: This is the legacy Network Application stack. Use it only if you explicitly want linuxserver + external MongoDB instead of UniFi OS Server.
- Auto-detect mode — upgrades an existing Docker Compose install, or fresh-installs when none is found
- Data preservation — tarball backup before upgrade, MongoDB document-count baseline, automatic rollback on verification failure
- Safety-first — aborts before destructive steps when data loss is possible; verbose [DETAIL] logging and a full session log under/tmp/
- Script lock — prevents concurrent runs against the same install
- Pre-flight checks — RAM, disk space, port conflicts, DNS/registry connectivity, Docker-in-Docker detection
- MongoDB guards — blocks unsupported major-version jumps and scans container logs for WiredTiger/upgrade errors
- Dynamic memory limits — adjusts UniFi MEM_LIMIT /MEM_STARTUP based on host RAM
- Idempotent Docker setup — installs docker.io anddocker-compose-v2 if missing
| Resource | Minimum | Recommended | 
|---|---|---|
| RAM | 1 GB (script aborts below) | 4 GB+ | 
| Disk | 5 GB free on install volume | 10 GB+ (more for large sites / backups) | 
| OS | Ubuntu or Debian (apt-based) | Ubuntu 22.04+ / Debian 12+ | 
| CPU | x86_64 with AVX (MongoDB >4.4) or arm64 with ARMv8.2-A for MongoDB 5+ | On arm64 without ARMv8.2-A (e.g. Khadas VIM 4), the script auto-selects MongoDB 4.4.18 | 
Must run on the host OS, not inside a Docker container. Docker (or permission to install it via sudo) is required.
The script checks for and uses: bash, curl, sudo, awk, tar, realpath, flock, openssl, and docker compose.
The file must be the raw Bash script. If you download GitHub’s HTML web page instead, Bash will fail with:
syntax error near unexpected token `newline'
<!DOCTYPE html>'
Verify before running:
head -1 install-unifi-docker.sh
# Expected: #!/usr/bin/env bash
file install-unifi-docker.sh
# Expected: Bourne-Again shell script (not HTML)git clone https://github.com/YOUR_USERNAME/unifi-networkapplication.git
cd unifi-networkapplication
chmod +x install-unifi-docker.sh
Use the raw URL, not the github.com/.../blob/... page:
curl -fsSL -o install-unifi-docker.sh \
  https://raw.githubusercontent.com/YOUR_USERNAME/unifi-networkapplication/main/install-unifi-docker.sh
chmod +x install-unifi-docker.sh
head -1 install-unifi-docker.sh   # must print: #!/usr/bin/env bash
From Windows/macOS/Linux where you already have the repo:
scp install-unifi-docker.sh root@<khadas-ip>:~/
ssh root@<khadas-ip> 'chmod +x ~/install-unifi-docker.sh && head -1 ~/install-unifi-docker.sh'git clone https://github.com/YOUR_USERNAME/unifi-networkapplication.git
cd unifi-networkapplication
chmod +x install-unifi-docker.sh
./install-unifi-docker.sh -y
Default install location: ~/unifi
Open the web UI after install:
https://<your-host-ip>:8443
MongoDB credentials and image tags are stored in ~/unifi/.env (mode 600 — back this up securely).
./install-unifi-docker.sh [options]
| Option | Description | 
|---|---|
| -d ,--dir PATH | Install directory (default: $HOME/unifi ) | 
| -t ,--tz TIMEZONE | Timezone (default: auto-detected, fallback UTC ) | 
| --unifi-tag TAG | linuxserver UniFi image tag (default: latest on upgrade) | 
| --mongo-tag TAG | Official MongoDB image tag (default: 7.0 on capable hosts; 4.4.18 auto-selected on arm64 without ARMv8.2-A; preserved on upgrade unless set) | 
| --network MODE | bridge (default) orhost | 
| --fresh | Force fresh install (refuses if legacy data exists at --dir ) | 
| --migrate-from-deb | Automated native .deb → Docker migration | 
| --backup-file PATH | .unf backup for migration (optional) | 
| --unifi-user USER | Native controller admin user (or UNIFI_CTRL_USER ) | 
| --unifi-pass PASS | Native controller password (or UNIFI_CTRL_PASS ) | 
| --keep-native-enabled | Leave native unifi.service enabled after verified migration | 
| -y ,--yes | Non-interactive; skip confirmation prompt | 
| -h ,--help | Show built-in help | 
First-time install (non-interactive):
./install-unifi-docker.sh -y
Custom install path and timezone:
./install-unifi-docker.sh -d /opt/unifi -t America/New_York -y
Upgrade existing install to latest UniFi (Mongo tag unchanged):
cd ~/unifi
/path/to/install-unifi-docker.sh -d ~/unifi -y
Pin UniFi version:
./install-unifi-docker.sh --unifi-tag 9.0.114 -y
Fresh install on empty directory only:
./install-unifi-docker.sh --dir /opt/unifi-new --fresh -y
The script chooses a mode automatically after scanning the system:
| Mode | When | What happens | 
|---|---|---|
| upgrade | docker-compose.yml +.env exist at--dir | Backup → baseline → pull new images → verify data → rollback on failure | 
| fresh | No recognized UniFi install | New stack with generated MongoDB credentials | 
| blocked | Native .deb , orphan data, or ambiguous state | Exits with instructions; no changes made | 
On upgrade, the script:
- Stops UniFi, captures a MongoDB document baseline
- Stops MongoDB gracefully
- Creates a full tarball under ../unifi-backups/unifi-pre-upgrade-*.tar.gz
- Pulls new images and restarts the stack
- Verifies collection counts and mongo-data/ size have not regressed
- Scans MongoDB logs for storage/upgrade errors
- Rolls back from the tarball if verification fails
The script never regenerates MongoDB passwords on upgrade. It aborts if:
- mongo-data/ exists without matching.env / compose files
- --fresh is used but an existing Docker install is detected
- --mongo-tag would downgrade MongoDB or skip a major version (e.g. 5.x → 7.x)
- RAM or disk space is insufficient for a safe backup
The script can backup, stop native UniFi, install Docker, restore your .unf, and set Inform Host via the API:
export UNIFI_CTRL_USER="admin"
export UNIFI_CTRL_PASS="your-controller-password"
curl -fsSL -o install-unifi-docker.sh \
