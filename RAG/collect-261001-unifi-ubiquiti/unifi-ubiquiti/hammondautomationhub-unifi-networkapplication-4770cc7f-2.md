---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/hammondautomationhub-unifi-networkapplication-4770cc7f-2
title: "Fresh install"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "memory"]
source: docs/RAG/collect-261001-unifi-ubiquiti/hammondautomationhub-unifi-networkapplication-4770cc7f.md
source_anchor: ""
source_lines: [119, 210]
sha256: 61faa7af11bb5affeef798bae1af394aec0be70a0978900317381920841eb432
---

# Fresh install

  https://raw.githubusercontent.com/HammondAutomationHub/unifi-networkapplication/main/install-unifi-docker.sh
chmod +x install-unifi-docker.sh
./install-unifi-docker.sh --migrate-from-deb -y
Backup sources (first match wins):
- --backup-file /path/to/backup.unf if provided
- Live API backup using --unifi-user /--unifi-pass (native service must be running)
- Latest .unf from native autobackup (/usr/lib/unifi/data/backup/autobackup/ )
Prefer credentials for a fresh backup. Autobackup files older than 24 hours trigger a warning.
What is automated:
| Step | Automated | 
|---|---|
| Create/find .unf backup | Yes | 
| Stop native unifi.service | Yes (stays enabled for rollback if migration fails) | 
| Install Docker Compose stack | Yes | 
| Upload .unf to new controller | Yes (API) | 
| Set Inform Host + Override | Yes (requires credentials) | 
| Disable native service | Yes, automatic after restore + verification succeed | 
If API restore fails, the script leaves Docker running and prints the backup path for manual restore in the web UI.
- Settings → System → Backup → download .unf
- sudo systemctl stop unifi
- ./install-unifi-docker.sh --fresh -y
- Restore in the web wizard; set Inform Host manually
After the initial install, routine updates can use Compose directly:
cd ~/unifi
sudo docker compose pull
sudo docker compose up -d
For community-safe upgrades with backup and verification, re-run the installer:
./install-unifi-docker.sh -d ~/unifi -y
View logs:
cd ~/unifi
sudo docker compose logs -f
sudo docker compose logs -f unifi-db
Install script log (each run):
/tmp/unifi-install-YYYYMMDD-HHMMSS.log
Default bridge mode publishes:
| Port | Protocol | Purpose | 
|---|---|---|
| 8443 | TCP | Web UI / controller API | 
| 8080 | TCP | Device inform (HTTP) | 
| 8843 | TCP | Guest portal HTTPS | 
| 8880 | TCP | Guest portal HTTP | 
| 3478 | UDP | STUN | 
| 10001 | UDP | Device discovery | 
| 1900 | UDP | UPnP / SSDP | 
Use --network host if you need the controller to bind directly on the host network stack.
Per linuxserver.io documentation:
| UniFi Network Application | MongoDB | 
|---|---|
| 8.1+ | 3.6 – 7.0 | 
| 9.0+ | 3.6 – 8.0 | 
- Fresh install defaults to MongoDB 7.0 .
- Upgrade preserves the installed Mongo tag unless you pass --mongo-tag .
- Do not use latest for MongoDB in production; pin a version and upgrade one major at a time.
syntax error / <!DOCTYPE html> on line 7
The file on disk is an HTML page, not the installer. Remove it and re-download using Getting the script above. Do not save the GitHub “blob” browser URL in a browser “Save as” dialog.
rm -f ./install-unifi-docker.sh
# Then use git clone, raw curl, or scp — and verify: head -1 install-unifi-docker.sh
Script says another instance is running
Wait for the other install to finish, or remove a stale lock only if no installer is running:
# Only if you are certain no install is in progress
sudo rm -f /run/lock/unifi-docker-install.lock
Port 8443 already in use
Stop the conflicting service or choose a different host. Native UniFi (.deb) must be migrated manually (see above).
MongoDB crash-loop on old x86_64 CPU (no AVX)
./install-unifi-docker.sh --mongo-tag 4.4.18 -y
MongoDB Illegal instruction on ARM64 (Khadas VIM 4, Raspberry Pi 4, etc.)
MongoDB 5.0+ and 4.4.19+ require ARMv8.2-A. The script (v1.4.6+) detects this and defaults to 4.4.18 automatically. If a previous run wrote Mongo 7.0 data or .env, wipe the failed data and re-run:
cd /root/unifi   # or your --dir path
sudo docker compose down
sudo rm -rf mongo-data/*
curl -fsSL -o install-unifi-docker.sh \
  https://raw.githubusercontent.com/HammondAutomationHub/unifi-networkapplication/main/install-unifi-docker.sh
chmod +x install-unifi-docker.sh
sudo ./install-unifi-docker.sh --migrate-from-deb -y
Low memory on ARM boards (e.g. Khadas VIM 4)
The script sets conservative MEM_LIMIT values automatically. If the host has less than 2 GB RAM, add swap and ensure adequate free disk before upgrading.
Upgrade rolled back
Restore is attempted automatically. Manual recovery:
ls -lt "$(dirname ~/unifi)/unifi-backups/"
# Extract the latest pre-upgrade tarball if neededunifi-networkapplication/
├── install-unifi-os-server.sh   # Official UniFi OS Server (Podman) — recommended
├── install-unifi-docker.sh      # Legacy linuxserver Network Application + MongoDB
└── README.md
After install:
~/unifi/                      # or your --dir path
├── docker-compose.yml
├── .env                        # secrets — chmod 600
├── init-mongo.sh
├── config/                     # UniFi application data
└── mongo-data/                 # MongoDB data files
MIT — see LICENSE if present, or distribute under MIT terms as stated in the script header.
This project is community-maintained and not affiliated with Ubiquiti, linuxserver.io, or MongoDB Inc. Test upgrades in a lab when possible, keep .env and backup archives safe, and review linuxserver.io release notes before pinning tags in production.
