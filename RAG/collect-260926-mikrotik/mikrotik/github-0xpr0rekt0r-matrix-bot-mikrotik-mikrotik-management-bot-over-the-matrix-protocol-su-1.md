---
id: collect-260926-mikrotik/mikrotik/github-0xpr0rekt0r-matrix-bot-mikrotik-mikrotik-management-bot-over-the-matrix-protocol-su-1
title: "1. Create a working directory"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/github-0xpr0rekt0r-matrix-bot-mikrotik-mikrotik-management-bot-over-the-matrix-protocol-supports-all.md
source_anchor: ""
source_lines: [1, 191]
sha256: 71bf99cd6081257382ac9a2aa6717e4a87fb2b52486e5b26f0323145d3e54c34
---

# 1. Create a working directory

A containerised MikroTik management bot that bridges Matrix rooms to MikroTik routers. Extends the `matrix-cli` base image and supports the full range of deployed MikroTik hardware.

```
┌──────────────────────────────────────────────────────────────┐
│  Container (UID 10001, read-only rootfs)                     │
│                                                              │
│  PID 1: mikrotik_entrypoint.sh                               │
│   └── python3 bot.py   →  matrix-cli listener                │
│                                │                             │
│                     ┌──────────┴──────────┐                  │
│                     ▼                     ▼                  │
│             REST API :443/80       RouterOS API :8728/8729   │
│             (RouterOS 7.1+)        (RouterOS 3.x – 6.x)      │
└──────────────────────────────────────────────────────────────┘
```
The entrypoint is PID 1. It starts the Python bot and restarts it on transient exit with exponential back-off.

The bot selects the router communication protocol automatically based on the `port` value in `config.yaml`. No extra flags are needed.

| Port | Protocol | RouterOS version | 
|---|---|---|
| `443` or`80` | RouterOS REST API (HTTPS/HTTP) | 7.1+ | 
| `8729` | RouterOS API over TLS ( `librouteros` ) | 6.x+ | 
| `8728` (or any other) | RouterOS API plaintext ( `librouteros` ) | 3.x+ | 

Port `8728` is unencrypted. Use it only on isolated management networks.
Prefer port `8729` for RouterOS 6.x routers whenever possible.


```
.
├── .github/
│   └── workflows/
│       ├── ci.yml              # lint and type-check on every push / PR
│       └── build.yml           # build and publish image on release / schedule
├── bot/
│   ├── bot.py                  # bot application logic
│   └── tests/                  # unit tests
├── pyproject.toml              # Python dependencies and tooling
├── uv.lock                     # locked dependency manifest — commit this
├── config/
│   └── routers_example.yaml    # copy to project root, rename to config.yaml, fill in, keep out of git
├── docker/
│   ├── Dockerfile
│   └── mikrotik_entrypoint.sh  # process supervisor
├── docker-compose.yaml         # production: pulls pre-built image from GHCR
├── docker-compose.dev.yaml     # development: builds image locally
├── .gitignore
└── README.md
```
| Tool | Minimum version | 
|---|---|
| Docker Engine | 24.x (BuildKit enabled) | 
| Docker Compose | v2.x | 
| Python | 3.14+ (development only) | 
| uv | latest stable (development only) | 
| GNU Make | development only | 
| direnv | optional (development only) | 

No tools beyond Docker are required to run the pre-built image. Python, uv, Make, and direnv are only needed for local development and checks.

For end users who want to run the bot without building anything.

```
# 1. Create a working directory
mkdir matrix-bot-mikrotik && cd matrix-bot-mikrotik
# 2. Download the production compose file
curl -O https://raw.githubusercontent.com/underhax/matrix-bot-mikrotik/main/docker-compose.yaml
# 3. Prepare directories and files — see Host Setup below
# 4. Complete Matrix Session Login — see Matrix Session Login below
# 5. Start
docker compose up -d
```
Run these commands once in your working directory before starting the container. The container runs as **UID/GID 10001** — all writable mount points must be owned by that UID.

```
mkdir -p bot_data
sudo chown -R 10001:10001 bot_data
```
```
# Download the example config and rename it
curl -o config.yaml https://raw.githubusercontent.com/underhax/matrix-bot-mikrotik/main/config/routers_example.yaml
sudo chown 10001:10001 config.yaml
chmod 0600 config.yaml
```
Fill in your router entries — see Configuration.

The configuration file contains security settings, a command whitelist, and router connection details. Download the annotated example to get started:

```
curl -o config.yaml https://raw.githubusercontent.com/underhax/matrix-bot-mikrotik/main/config/config_example.yaml
sudo chown 10001:10001 config.yaml && chmod 0600 config.yaml
```
Key sections:

| Section | Description | 
|---|---|
| `bot_user` | Full Matrix ID of the bot account — used to ignore its own messages | 
| `command_room` | Room ID where commands are accepted | 
| `admin_room` | Room ID for security alerts only — does not accept commands | 
| `allowed_users` | List of Matrix user IDs permitted to issue commands | 
| `allowed_commands` | Whitelist of RouterOS API paths the bot will execute | 
| `routers` | Router connection details — `router_id` maps to host, port, credentials | 

The `router_id` must match `[A-Za-z0-9_-]`, 1–64 characters.

```
# Security
bot_user: "@mikrotik-bot:your.homeserver"
command_room: "!yourRoomId:your.homeserver"
admin_room: "!adminRoomId:your.homeserver"
allowed_users:
  - "@admin:your.homeserver"
# Command whitelist
allowed_commands:
  - "system/resource"
  - "ip/address"
  # ... see config_example.yaml for the full default list
# Routers
routers:
  # RouterOS 7.1+ — REST API over HTTPS, valid CA-signed certificate
  core-01:
    host: "192.168.88.1"
    port: 443
    username: "api-bot"
    password: "strong-random-password"
    tls_verify: true
  # RouterOS 7.1+ — REST API over HTTPS, self-signed certificate
  branch-02:
    host: "10.0.1.1"
    port: 443
    username: "api-bot"
    password: "another-strong-password"
    tls_verify: false       # TLS is still in use; certificate is not verified
  # RouterOS 6.x — RouterOS API over TLS (librouteros)
  office-03:
    host: "10.10.0.1"
    port: 8729
    username: "api-bot"
    password: "yet-another-password"
    tls_verify: false       # RouterOS 6.x uses a self-signed cert on 8729
  # RouterOS 3.x / 4.x / 5.x / 6.x — RouterOS API plaintext (librouteros)
  legacy-04:
    host: "172.16.5.1"
    port: 8728
    username: "admin"
    password: "old-password"
    tls_verify: false
```
| Variable | Default | Description | 
|---|---|---|
| `ALLOW_WRITES` | `false` | Set to `true` to allow write (add/set) operations. Default is read-only. | 
| `BOT_CONFIG` | `/home/bot/config/config.yaml` | Config path inside the container. No need to change. | 
| `TARGETARCH` | `amd64` | Build target: `amd64` or`arm64` . For local dev builds only. | 

Set persistent overrides in a `.env` file in the working directory:

`ALLOW_WRITES=false`
Create a dedicated, least-privilege user on each router. Never use the `admin` account for automated access.

Connect to the router via Winbox, WebFig, or SSH and run:

```
# Read-only group (recommended default — matches ALLOW_WRITES=false)
/user group add name=api-readonly \
  policy=read,api,!local,!telnet,!ssh,!ftp,!reboot,!write,!policy,!test,!winbox,!password,!web,!sniff,!sensitive,!romon
# Read-write group (use only when ALLOW_WRITES=true)
/user group add name=api-readwrite \
  policy=read,write,api,!local,!telnet,!ssh,!ftp,!reboot,!policy,!test,!winbox,!password,!web,!sniff,!sensitive,!romon
```
```
/user add name=api-bot group=api-readonly password="strong-random-password"
```
Allow connections only from the container's host IP:

```
/user set [find name=api-bot] address=192.168.1.100/32
```
**RouterOS 7.1+ (REST API, ports 80/443):**

```
/ip service print
/ip service enable www-ssl
/ip service set www-ssl port=443
# Optional: disable plain HTTP if not needed
/ip service disable www
```
**RouterOS 6.x (RouterOS API, ports 8728/8729):**

Port 8728 (plaintext) and 8729 (TLS) are enabled by default. To verify:

