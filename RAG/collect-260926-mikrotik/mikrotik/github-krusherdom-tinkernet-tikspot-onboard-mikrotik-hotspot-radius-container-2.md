---
id: collect-260926-mikrotik/mikrotik/github-krusherdom-tinkernet-tikspot-onboard-mikrotik-hotspot-radius-container-2
title: "github-krusherdom-tinkernet-tikspot-onboard-mikrotik-hotspot-radius-container"
domain: mikrotik
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/forum/dhcp-dns/github-krusherdom-tinkernet-tikspot-onboard-mikrotik-hotspot-radius-container.md
source_anchor: ""
source_lines: [116, 188]
sha256: d49f9432d250b7fc8283dedbc6cf967fa7177032c1330d8c3b19fe0db4b6aa8a
---

# github-krusherdom-tinkernet-tikspot-onboard-mikrotik-hotspot-radius-container

Guest-lookup **plugins** are portable JSON recipes that let guests log in with details
your own system already holds (room + surname, booking reference, mobile, email…). The
`plugins/` folder is the community catalog: the admin's **Guest lookup →
Browse catalog** reads its `index.json` and imports a recipe in one click (disabled and
without secrets until you review it). Point *Settings → Plugin catalog URL* at any
GitHub folder, `index.json` or single exported file to use your own catalog.

| **Plugin editor** — guided form: source, secrets, parameters, match rules, stay window, guest inputs, messages | **Built-in guest list** — no PMS? Paste a CSV and match against it | 
| **Test panel** — run a lookup with sample inputs and see the parsed records | **Guest lookup block** — drop it into any portal design; the fields come from the plugin | 

- How recipes work: `docs/guest-lookup-plugins.md` , field
reference`docs/plugins/README.md`
- Try it with the bundled demo guest API: `examples/guest-api/`
- Ready-made PMS/event recipes — supported-systems matrix in
`docs/plugins/README.md` :
  - **RMS Cloud** :`docs/plugins/rms-cloud.md` (mock:`examples/rms-mock/` )
  - **Mews** :`docs/plugins/mews.md` (mock:`examples/mews-mock/` )
  - **Apaleo** :`docs/plugins/apaleo.md` (mock:`examples/apaleo-mock/` )
  - **Cloudbeds** :`docs/plugins/cloudbeds.md` (mock:`examples/cloudbeds-mock/` )
  - **Eventbrite** attendees:`docs/plugins/eventbrite.md` (mock:`examples/eventbrite-mock/` )
  - **CSV / Google Sheet** , or the built-in guest list:`docs/plugins/csv-and-guest-list.md` (mock:`examples/csv-guest-list/` )
- Contribute one: `plugins/README.md`

Two paths, depending on RouterOS version:

1. **RouterOS 7.22+ — container "App"***(simplest; auto-provisions the network)* : add`deploy/tikspot.app.yml` with`/app add network=lan` and the
router pulls the public multi-arch image (`ghcr.io/krusherdom/tinkernet-tikspot` ,
published on each release tag) and creates the veth, bridge port, IP and NAT for you.
See`docs/deploy-app.md` .
2. **File-based***(any RouterOS 7 with the `container` package)* : build the arm64 tar,
upload it, create the veth, and`/container/add` . See`docs/deploy-rb5009.md` (RB5009 walk-through) and`docs/setup-mikrotik.md` (generic).

Then browse to `http://<container-ip>/admin`, complete the **setup wizard** (set an admin
password, point it at the router, **Auto-configure**), design the portal, and install the
hotspot shim files. `docs/backup-migrate.md` covers moving an
existing install to a new device.

All settings have container defaults and can be overridden via MikroTik
`/container/envs` (or `docker run -e`). The router connection, secrets and most options are
configured in the admin wizard and stored in `/data`.

| Env var | Default | Purpose | 
|---|---|---|
| `TIKSPOT_NAS_SECRET` | *(generated)* | RADIUS shared secret (router ⇄ container). Set it, or let setup generate one. | 
| `TIKSPOT_FREE_USERNAME` /`_PASSWORD` | `free` /`free` | Credential the free-login button submits | 
| `TIKSPOT_DB` | `/data/tikspot.db` | Shared SQLite path | 
| `TIKSPOT_DATA_DIR` /`TIKSPOT_ASSETS_DIR` | `/data` /`/data/assets` | Persistent data + branding uploads | 
| `TIKSPOT_HOST` /`TIKSPOT_PORT` | `0.0.0.0` /`80` | Listen address | 
| `TIKSPOT_RADIUS_CLIENT_NET` /`_NET6` | `0.0.0.0/0` /`::/0` | Source range FreeRADIUS trusts (the secret is the gate) | 
| `LOG_LEVEL` | `info` | Fastify log level | 

This is a LAN-management appliance. Within that scope: the admin password is **scrypt**-
hashed; sessions are **signed, httpOnly, SameSite** cookies; the admin login is
**rate-limited**; state-changing requests get an **Origin/CSRF** check; backups **redact
secrets** by default; and the RADIUS NAS secret **fails closed** (no insecure default).
Don't expose the admin port directly to the internet.

- Source is plain ES modules (Node 22, `"type": "module"` ); no build step at runtime.
- `cd app && npm run dev` runs the server with`--watch` (needs a local SQLite + the
RADIUS daemon for full function; most easily exercised inside the built container).
- Bookkeeping: `changelog.md` (per-version summary) and`changelog_detailed.md` (granular log).

Contributions are welcome — see `CONTRIBUTING.md` for dev setup, the PR
flow, and the < 250 MB image budget. Please open an issue (use the templates) before starting
non-trivial work. All PRs are gated on passing CI and maintainer review. By participating you
agree to the Code of Conduct.

- **Bugs / features:** open a GitHub issue
using the bug-report or feature-request template.
- **Security vulnerabilities:** please**don't** file a public issue — report privately via the
repo's**Security → Report a vulnerability** tab. See`SECURITY.md` .

MIT.
