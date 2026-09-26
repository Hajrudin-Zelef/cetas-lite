---
id: collect-260926-mikrotik/mikrotik/github-krusherdom-tinkernet-tikspot-onboard-mikrotik-hotspot-radius-container-1
title: "github-krusherdom-tinkernet-tikspot-onboard-mikrotik-hotspot-radius-container"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/lot-mikrotik/forum/dhcp-dns/github-krusherdom-tinkernet-tikspot-onboard-mikrotik-hotspot-radius-container.md
source_anchor: ""
source_lines: [1, 115]
sha256: 32c6fb7c587ad8640d11268f2b6cb88fd99b4a7c2abfe43535a8b53f48fbe939
---

# github-krusherdom-tinkernet-tikspot-onboard-mikrotik-hotspot-radius-container

A single, self-contained container that runs **on a MikroTik router** (RouterOS v7
`container` feature) and provides a complete Wi‑Fi hotspot stack — captive portal,
RADIUS authentication, and a web admin — with no external servers.

- **FreeRADIUS** authenticates hotspot users and applies speed / data / time limits via
MikroTik vendor attributes, sharing one SQLite database with the app.
- **Live captive portal** — the page your guests see. The router's hotspot redirects
clients to the container, which renders a customisable login page (free login, voucher
codes, or user accounts) and posts the final login back to the router.
- **Admin portal** — manage plans, vouchers and accounts; watch and kick active users;
design the portal page; audit activity; back up/restore; and run a guided first-time
setup that auto-configures the router for you.

The image stays **under 250 MB** so it fits hotspot-class MikroTik devices (hAP ax²/ax³,
RB5009, x86/CHR). Persistent data lives on a mounted `/data` volume.

**Version:** 0.10.0 — see `changelog.md`. LAN-management tool: the admin
runs over HTTP behind your router, not on the public internet.

**Tested on:** a MikroTik **RB5009 running RouterOS 7.22** (container + FreeRADIUS
confirmed end-to-end). **Targets RouterOS 7.22 and later** — the `/app`-based install
needs 7.22+; the file-based install also works on earlier 7.x with the `container` package.


**The captive portal** — the page guests see when they join the Wi‑Fi. Free login, voucher
codes and account login on one customisable page.

**The live page designer** — drag-and-drop blocks (logo, heading, text, login widgets) with a
desktop/mobile preview and a full colour picker. No re-uploading files to the router per change.

Designs are versioned: keep several, publish one, revert to any earlier version.

**The admin** — manage everything from one LAN-only web app.

| **Active users** — live sessions from RADIUS accounting; kick via CoA | **Plans** — MikroTik speed / data / time limits, incl. renew-at-midnight | 
| **Vouchers** — printable batches with optional date windows | **Router setup** — probe, auto-configure, or copy a manual script | 
| **Guest lookup** — plugins that check a guest against your PMS, a CSV, or a built-in list | **Plugin catalog** — browse community recipes on GitHub and import in one click | 
| **Settings** — every tunable in one place, with help text | **Help** — a checklist for the MikroTik-side setup traps | 

```
Guest device ──▶ MikroTik hotspot ──redirect──▶ Tikspot container (this project)
                       ▲                              │  live login page
                       └────── login POST ◀───────────┘  (FreeRADIUS auth + limits)
```
The router holds only a few tiny **redirect-shim** files (download them as a zip, or have
the app push them over the API). They hand the hotspot session to the container, which
hosts the real, editable page — so you customise the portal live, without re-uploading
files to the router for each change. FreeRADIUS and the Node app run inside the same
container and share `/data/tikspot.db`; the app projects plans/vouchers/accounts into the
RADIUS tables, and FreeRADIUS remains the single auth authority.

- **Four login types** per portal: one-tap**free** login,**voucher** codes, named**user accounts** , and**guest lookup** — verify guests against your own system (hotel
PMS, membership API…) via admin-authored recipes (JSON / XML / regex parsing, token
auth, stay-window checks). See`docs/guest-lookup-plugins.md` and the demo API in`examples/guest-api/` .
- **Plans** = MikroTik limits (rate`5M/5M` , data cap, session time). New:**"expire at
midnight"** plans that renew daily (sessions are CoA-disconnected at the router's local
midnight for a fresh quota) instead of a fixed time limit.
- **Voucher batches** with optional date-validity windows; printable voucher sheets.
- **MAC re-auth** ("remember device") so returning guests reconnect automatically.
- **Live page designer** — registry-driven blocks (logo, heading, text, columns, link
buttons, terms checkbox, sanitised HTML, all login widgets), per-block styling, themes
with background images,**drafts / publish / version history** , templates, multiple
designs, and themed connected / logged-out pages. The canvas is the real server render.
- **Announcements & settings** — banners on the portal and admin with severity and time
windows; a registry-driven Settings tab (portal texts, login method, log retention…).
- **Logs & reports** — bounded by a daily retention sweep, an app event log, CSV export,
and a reports view (logins per day, usage per plan, top users). Help tab with a setup
checklist and troubleshooting table.
- **Guided setup wizard** that probes the router and**auto-configures** the RADIUS client,
hotspot profile, DNS static and walled-garden — every object it creates is tagged with a
managed comment and can be**queried back / verified** from the admin (per-component
pass/fail with the raw RouterOS line).
- **Operations** : live active-users + kick (RADIUS CoA), RADIUS auth logs, an**admin
audit trail** , container/router**health** , and**backup/restore** (redacted of secrets
by default) to migrate between devices.

| Layer | Choice | 
|---|---|
| Runtime | Node.js 22 + Fastify | 
| Storage | SQLite via `better-sqlite3` (shared with FreeRADIUS`rlm_sql` ) | 
| Auth | FreeRADIUS (SQLite backend) + MikroTik vendor attributes; CoA for kick/expiry | 
| Supervision | s6-overlay | 
| Base | Alpine, multi-arch `arm64` +`amd64` , image < 250 MB | 

| Path | What it is | 
|---|---|
| `app/` | Node.js (Fastify) backend + static admin/portal assets | 
| `app/src/admin/` | Admin API: auth, setup wizard, plans/vouchers/accounts, backup, logs, validation, audit, rate-limiting | 
| `app/src/radius/` | RADIUS projection ( `sync` ), CoA (`coa` ), NAS secret (`nas` ),`clients.conf` rendering, midnight-expiry sweeper | 
| `app/src/portal/` | Captive-portal rendering + the served `/m/portal.js` client | 
| `app/src/mikrotik/` | RouterOS v7 REST client (auto-configure, verify, managed-object listing) | 
| `app/src/{db,design,mac,voucher,hotspot}/` | Schema/migrations, design model + block registry + templates, MAC re-auth, voucher sweeper, hotspot shim generator | 
| `app/src/plugins/` | Guest-lookup engine (recipe schema, templating, JSON/XML/regex parsers, matching, HTTP client, grants) | 
| `app/public/admin/` | Vanilla-JS admin SPA + the captive-portal page designer | 
| `app/test/` | `node --test` suites: pure unit tests, in-memory SQLite tests, route tests via`app.inject` , plugin engine + demo-API integration | 
| `examples/guest-api/` | Zero-dependency demo guest API + three ready recipes for trying guest lookup | 
| `docker/` | Multi-stage, multi-arch Dockerfile + s6 service tree ( `00-init` ,`db-init` ,`radiusd` ,`node` ) | 
| `docs/` | Setup & deployment guides (see below) | 
| `deploy/` | `tikspot.app.yml` — RouterOS 7.22+ container**App** manifest (self-provisions networking) | 
| `scripts/` | Build helpers (image size gate, RB5009 export) | 

Requires Docker with buildx (Docker Desktop on Windows is fine).

```
npm run build:local     # amd64 image (tikspot:dev) + the <250 MB size gate
npm run build:release   # multi-arch (arm64 + amd64)
npm run export:rb5009   # arm64 -> dist/tikspot-rb5009.tar (RouterOS docker-archive)
```
Run the unit tests:

`cd app && npm test      # node --test`
**Releases:** pushing a `v*` tag (e.g. `git tag v0.10.0 && git push origin v0.10.0`) runs the
release workflow, which builds the multi-arch image and publishes it to
`ghcr.io/krusherdom/tinkernet-tikspot` for the RouterOS App deploy below.

