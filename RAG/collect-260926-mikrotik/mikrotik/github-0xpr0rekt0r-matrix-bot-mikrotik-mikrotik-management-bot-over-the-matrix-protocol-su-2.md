---
id: collect-260926-mikrotik/mikrotik/github-0xpr0rekt0r-matrix-bot-mikrotik-mikrotik-management-bot-over-the-matrix-protocol-su-2
title: "1. Create a working directory"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "parameters"]
source: docs/RAG/lot-mikrotik/forum/misc/github-0xpr0rekt0r-matrix-bot-mikrotik-mikrotik-management-bot-over-the-matrix-protocol-supports-all.md
source_anchor: ""
source_lines: [192, 371]
sha256: 401f32538fa50392547d49ebe9b85997634929d58a3bc2ede99e310b7001d9a8
---

# 1. Create a working directory

```
/ip service print
```
To enable TLS on port 8729, assign a certificate to the `api-ssl` service.
RouterOS can generate a self-signed one:

```
/certificate add name=api-ssl common-name=api-ssl key-size=2048 \
  key-usage=key-cert-sign,crl-sign,tls-server days-valid=3650
/certificate sign api-ssl
/ip service set api-ssl certificate=api-ssl
/ip service enable api-ssl
```
**RouterOS 3.x – 5.x (RouterOS API plaintext, port 8728 only):**

Port 8728 is enabled by default. No additional setup required.

The complete local setup, dependency management, formatting, linting, type
checking, testing, `direnv`, and Docker development workflow is documented in
`DEVELOPMENT.md`.

The bot uses `matrix-cli` for all Matrix communication. A session must be created once before the first start. The session is persisted in `bot_data/` and reused on every subsequent start.

Register a new Matrix account on your homeserver (e.g. via Element or the homeserver's admin panel). Use a dedicated account — do not use your personal account.

Example: `@mikrotik-bot:your.homeserver`

`docker compose up -d``docker compose exec -it matrix-bot-mikrotik /bin/bash`
Inside the container (full CLI reference):

`matrix-cli --mode auth --server "matrix.org" --user "@mikrotik-bot:matrix.org"`
Session credentials are written to `bot_data/` on the host and persist across container restarts.

After login, you need to establish E2EE trust. You can generate cross-signing keys or start a verification flow. Inside the container:

```
# Verify with another user session interactively:
matrix-cli --mode verify --user "@you:your.homeserver"
```
For full verification details see the upstream documentation.

`exit``docker compose down && docker compose up -d`
In your Matrix client:

```
/invite @mikrotik-bot:your.homeserver
```
```
# Start as a daemon
docker compose up -d
# Tail logs
docker compose logs -f
# Stop gracefully
docker compose down
```
Every message in the room is inspected by the bot. Messages that do not start with `!mtik` are silently ignored — normal conversation is unaffected.

```
!mtik <router_id> <path> [=key=value ...]
```
| Part | Description | 
|---|---|
| `!mtik` | Bot command prefix. Required. | 
| `router_id` | Key from `config.yaml` — identifies which router to target. | 
| `path` | RouterOS menu path with `/` separators (see below). | 
| `=key=value` | Write parameters in RouterOS format. Requires `ALLOW_WRITES=true` . | 

MikroTik organises all configuration as a menu tree. The path you use in a bot command is the same path you navigate in the RouterOS terminal — with `/` as a separator instead of a space. The `print` verb is implicit for read commands and is never typed.

| RouterOS terminal | Bot command | 
|---|---|
| `/ip address print` | `!mtik core-01 ip/address` | 
| `/interface print` | `!mtik core-01 interface` | 
| `/system resource print` | `!mtik core-01 system/resource` | 
| `/ip firewall filter print` | `!mtik core-01 ip/firewall/filter` | 
| `/ip address add address=X interface=Y` | `!mtik core-01 ip/address =address=X =interface=Y` | 

**Official API documentation:**


- REST API (RouterOS 7.1+): https://help.mikrotik.com/docs/display/ROS/REST+API
- RouterOS API (all versions): https://help.mikrotik.com/docs/display/ROS/API

Once you understand the path mapping, any resource visible in the RouterOS terminal can be queried directly — no need to memorise bot-specific syntax.

A quick way to confirm a router is reachable and see its load, uptime, and firmware version in one response.

```
!mtik core-01 system/resource
```
Equivalent terminal command:

```
/system resource print
```
Example response:

```
✅ `core-01` → `system/resource`
[
  {
    "uptime": "15d2h34m12s",
    "version": "7.14.3 (stable)",
    "cpu-load": "3",
    "free-memory": "198901760",
    "total-memory": "268435456",
    "architecture-name": "arm64"
  }
]
```
Requires `ALLOW_WRITES=true`. Write commands modify the running configuration immediately — there is no confirmation prompt.

```
!mtik core-01 ip/address =address=10.99.0.1/24 =interface=ether2
```
Equivalent terminal command:

```
/ip address add address=10.99.0.1/24 interface=ether2
```
Example response:

```
✅ `core-01` → `ip/address =address=10.99.0.1/24 =interface=ether2`
[{"ret": "*6"}]
```
`ret` is the internal ID assigned to the new entry by RouterOS. Verify the result with a follow-up read:

```
!mtik core-01 ip/address
```
| Situation | Response | 
|---|---|
| Unknown `router_id` | `❌ Unknown router \` xyz`. Known IDs: core-01, branch-02` | 
| Router unreachable | `❌ Router \` core-01`: Cannot reach 192.168.88.1:443: ...` | 
| Auth failure | `❌ Router \` core-01`: Authentication failed — check credentials in config.yaml` | 
| Write blocked | `❌ Router \` core-01`: Write operations are disabled. Set ALLOW_WRITES=true to enable.` | 
| Invalid path characters | `❌ Router \` core-01`: Invalid API path: '../../etc'` | 
| RouterOS API trap | `❌ Router \` legacy-04`: RouterOS API trap: no such command` | 

**Container exits immediately — no output**

`bot_data/` is likely empty (no Matrix session). Complete the Matrix Session Login steps first.

**Bot is running but does not respond to commands**

1. Confirm the bot account is a member of the Matrix room.
2. Verify `bot_data/` contains a valid session (non-empty directory).
3. Check for errors in the logs:

`docker compose logs -f | grep -iE "error|warn|critical"`
**Docker build fails: `frozen lockfile` error**

The lockfile is missing or out of sync with `pyproject.toml`. Regenerate it:

```
uv lock
docker compose -f docker-compose.dev.yaml build --no-cache
```
**RouterOS API returns 401**

- Check `username` and`password` in`config.yaml` .
- Confirm the user has the `api` policy:

```
/user print detail where name=api-bot
```
- If an IP restriction is set, confirm the container host's IP is included:

```
/user print detail where name=api-bot
```
**Legacy router (port 8728) connects but returns no data**

RouterOS 3.x/4.x may require enabling the API service explicitly:

```
/ip service enable api
/ip service set api port=8728
```
