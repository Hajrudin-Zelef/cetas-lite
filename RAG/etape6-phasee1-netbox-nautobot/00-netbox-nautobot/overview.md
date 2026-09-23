---
id: etape6-phasee1-netbox-nautobot/00-netbox-nautobot/overview
title: "Phase E1 — NetBox & Nautobot: Network Source of Truth (Research File)"
domain: phase-e1-netbox-nautobot-network-source-of-truth-research-fi
role: deep-dive
task: reference
actors: []
dates: ["2024-05", "2026-01-06", "2026-06-30", "2026-08-25", "2026-09", "2026-09-02", "2026-09-15", "2026-09-22"]
keywords: ["research", "apache", "benchmarks", "governance", "license"]
source: docs/RAG/etape6_phaseE1_netbox_nautobot.md
source_anchor: ""
source_lines: [1, 77]
section: "Phase E1 — NetBox & Nautobot: Network Source of Truth (Research File)"
sha256: 75174b0103e2693b79d81644ee69d4231a08fbbd8570b9f1d49860e90ec37cd0
---

# Phase E1 — NetBox & Nautobot: Network Source of Truth (Research File)

> **Scope:** NetBox and Nautobot as network source-of-truth platforms (IPAM/DCIM/automation), their ecosystems, commercial offerings, head-to-head comparison, adjacent IPAM/DCIM tools, and real-world deployments.
> **Date coverage:** research current through **2026-09-22**.
> **Method:** read-only web research (browser_search / browser_open). No live-browser visits, no forms, nothing sent externally. No identifiers guessed. All new facts carry provenance tags; open items, gaps, conflicts and unverified claims are registered in the final log.
> **Provenance legend:** `[official]` = vendor/project's own docs, release notes, blog, GitHub releases; `[vendor-reported]` = vendor marketing, press releases, analyst quotes paid by vendor; `[independent]` = independent benchmarks, third-party analysis, community surveys; `[secondary]` = tech press, aggregators, re-reported; `[unverified]` = single-source, forum/Reddit/HN claims, or cannot be corroborated.

---

## Wave 1 — NetBox core: version line and platform features (2024–2026)

### 1.1 Version timeline (latest first, as of 2026-09-22)

- **NetBox v4.7.0 — released 2026-09-02** `[official]` — latest stable release as of 2026-09-22. Source: NetBox release notes, https://github.com/netbox-community/netbox/blob/HEAD/docs/release-notes/version-4.7.md
  - Requires **PostgreSQL 15 or later** (drops PostgreSQL 14 support).
  - Requires the PostgreSQL **ltree extension**; the release replaces `django-mptt` with ltree for hierarchical tables. Migration holds `ACCESS EXCLUSIVE` locks on hierarchical tables and backfills rows, blocking reads/writes for their duration (can last several minutes on large deployments); migrations are not reversible in practice. `rebuild_config_context_cache` runs one `UPDATE` per device and VM — maintenance window recommended `[official]`.
  - Requires **Redis 6.0 or later** (drops Redis 5.x).
- **NetBox v4.6.9 — 2026-08-25** `[official]` — enhancements: InfiniBand 4X interface types; 100GBase-X-SFP112 interface type; HPE Synergy interconnect link interface type. Performance: prefetch cable terminations to avoid N+1 queries on GraphQL cable fetches. Source: https://github.com/netbox-community/netbox/blob/HEAD/docs/release-notes/version-4.6.md
- **NetBox v4.6.4 — 2026-06-30** `[official]` — enhancements: 1C8P:8C1P breakout cable profile; JSON-schema enum dropdowns in module/device profile attribute forms; `dns_name` of primary & out-of-band IPs in event payloads. Performance: S3 image-attachment view optimization, chunked bulk updates to custom field data.
- **NetBox v4.5.0 — 2026-01-06** `[official]` — breaking changes: Python 3.12/3.13/3.14 required (drops 3.10/3.11); GraphQL queries filtering by object IDs/enums must use filter lookups (e.g. `id: {exact: 123}`); rendering device/VM configuration now requires the `render_config` permission; API token plaintext retrieval removed (`ALLOW_TOKEN_RETRIEVAL` removed); tokens can no longer be reassigned between users; `/api/dcim/cable-terminations/` endpoint is read-only (set terminations via `/api/dcim/cables/`); UI view for swapping A/Z circuit terminations removed; experimental HTMX navigation removed; webhooks payload uses `object_type` instead of `model`; `/api/extras/object-types/` removed (use `/api/core/object-types/`). Source: https://github.com/netbox-community/netbox/blob/HEAD/docs/release-notes/version-4.5.md
- **NetBox v4.1** (late 2024) `[official]` — notable features: Circuit Groups (circuits assignable to multiple groups with primary/secondary/tertiary priority); VLAN Group ID Ranges (multiple VID ranges per group, e.g. `1-100,1000-1999`); Nested Device Modules (module bays on modules for hierarchical submodules); Rack Types (like device types, auto-populate rack attributes); **Plugins Catalog Integration** — NetBox UI integrates with the canonical plugins catalog hosted by NetBox Labs, letting users browse plugins and check updates natively; User Notifications. Source: https://github.com/netbox-community/netbox/blob/HEAD/docs/release-notes/version-4.1.md
- Historical cadence note: NetBox v4.0 was released May 2024 with the UI refresh; the v4.x line is the current major line `[secondary]`. v4.7 exists on GitHub release-notes docs; no v4.8 observed as of 2026-09-22 `[unverified]` (absence of evidence).

### 1.2 Core data-model features (stable across v4.x)

- **DCIM:** sites, locations (nested via ltree in v4.7), racks/rack roles/types, devices, device types/modules, manufacturers, platforms, interfaces (including InfiniBand 4X and SFP112 types added 2026), cables with A/Z terminations and breakout profiles (1C8P:8C1P added 2026), front/rear ports, power feeds/panels, device bays, inventory items `[official]`.
- **IPAM:** VRFs, route targets, RIRs, aggregates, prefixes (with available-prefix allocation), IP ranges, IP addresses (primary/OOB, DNS name), VLANs/VLAN groups with multi-range VID support, FHRP groups, service templates `[official]`.
- **Circuits:** providers, provider networks, circuit types, circuits with A/Z terminations, circuit groups (added v4.1) `[official]`.
- **Virtualization:** clusters, cluster types/groups, virtual machines, VM interfaces `[official]`.
- **Wireless:** wireless LANs, wireless links (added in v3.x, carried in v4.x) `[secondary]`.
- **Tenancy:** tenants, tenant groups; contacts & assignments `[official]`.
- **Config contexts:** hierarchical JSON data merged by weight, applied to devices/VMs; platform-level config contexts also apply to platform children (behavior change in v4.5) `[official]`.
- **Customization:** custom fields, custom links, custom validators, export templates, webhooks, event rules, notification groups, saved filters, bookmarks, config templates (Jinja2) `[official]`.
- **Permissions:** object-level permissions with constraints; `render_config` permission split out in v4.5 `[official]`.

### 1.3 API surface

- **REST API:** full CRUD across all models, OpenAPI schema, token authentication; breaking note: cable-terminations endpoint read-only since v4.5, terminations managed via cables endpoint `[official]`.
- **GraphQL API:** single-endpoint queries; since v4.5, ID/enum filters require explicit lookups (`id: {exact: 123}`); performance work in v4.6.9 on cable-termination prefetching `[official]`.
- Webhooks/event payloads now reference `object_type` (with parent app label) instead of `model` since v4.5 `[official]`.

### 1.4 Deployment

- Official deployment: PostgreSQL 15+, Redis 6+, Python 3.12–3.14, WSGI (gunicorn/uvicorn) behind a reverse proxy; Docker Compose reference deployment maintained by the project; Kubernetes via community Helm charts `[official][secondary]`.
- Upgrade tooling: `upgrade.sh` script, `post_upgrade` management commands; v4.7 upgrade explicitly warns of extended duration and irreversibility of the mptt→ltree migration `[official]`.

### 1.5 Origin and governance

- NetBox was created by Jeremy Stretch at DigitalOcean and open-sourced in 2016; it is now stewarded by **NetBox Labs** (commercial entity founded 2023, backers include the original creator) `[secondary]`. Community repo: `netbox-community/netbox` on GitHub, Apache 2.0 license `[official]`.
- NetBox Labs hosts the canonical plugins catalog at https://netboxlabs.com/netbox-plugins/ and integrates it into the product UI (since v4.1) `[official]`.

---

### 1.6 NetBox v4.7.0/4.7.1 deep dive (September 2026) `[official]`

Source for all items below: https://github.com/netbox-community/netbox/blob/HEAD/docs/release-notes/version-4.7.md

- **Cooling infrastructure modeling** (new in v4.7): deliberately mirrors the power model.
  - New models: `CoolingSource` (facility cooling plant — chiller, cooling tower, dry cooler, facility water system — scoped to site/location); `CoolingFeed` (coolant loop delivered from a source to a rack); device components `CoolingIntake` (coolant intake on a device, e.g. server cold-plate connection or CDU facility-water inlet) and `CoolingOutflow` (coolant outlet on a CDU or manifold supplying downstream equipment); intakes reference their upstream outflow; device-type templates exist for both components; CDUs and manifolds are ordinary devices carrying these components.
  - Lightweight attributes: `cooling_method` (air, liquid, hybrid, immersion) on Device/DeviceType/ModuleType; `cooling_capability` (air-only, hybrid, liquid-only) and `cooling_capacity` on Rack and RackType.
  - REST endpoints: `/api/dcim/cooling-sources/`, `/api/dcim/cooling-feeds/`, `/api/dcim/cooling-intakes/`, `/api/dcim/cooling-intake-templates/`, `/api/dcim/cooling-outflows/`, `/api/dcim/cooling-outflow-templates/`.
- **Channelized subinterfaces (breakout modeling):** new `channels` field on Interface (number of physical channels an interface is divided into); each channel is a subinterface of the new generic `channel` type bound via `channel_id`; a single cable terminates to the channelized parent while NetBox traces a distinct cable path per channel subinterface. Fields also on interface templates (with a `parent` FK added to InterfaceTemplate).
- **Multi-protocol application services:** `protocol`+`ports` replaced by unified `port_mappings` (API: flat list of `protocol/port` strings, e.g. `["tcp/80", "udp/53"]`); legacy fields retained as deprecated in REST/GraphQL but read-only/derived at ORM level; GraphQL `ServiceProtocolEnum` members renamed (e.g. `ROLE_TCP` → `TCP`); new `port_mappings`/`protocol`/`port` filters.
- **Module bay types:** new `ModuleBayType` model (e.g. SFP28 cage, PCIe x16 slot); assignable to bays, bay templates, and module types; NetBox validates bay/module type-set overlap before permitting installation; templates propagate bay types to instantiated bays; `is_bay_compatible`/`is_module_compatible` read-only API fields.
- **Relocating installed modules:** an installed module can be moved to another module bay (including on a different device) via PATCH on `module_bay`; the module's entire subtree (components, own bays, child modules) moves; names/labels/positions derived from templates are re-resolved for the destination bay.
- **Background processing for REST bulk writes:** append `?background=true` to a bulk write → enqueues a job, returns `HTTP 202 Accepted` with job ID/URL; validation is deferred to the worker (202 = accepted, not succeeded).
- **Per-object errors for bulk operations:** failed bulk create/update returns `{"detail": ..., "errors": [{"index": N, "errors": {...}}]}` correlating errors to object indices; bulk remains all-or-none.
- **Pre-rendered config context data:** device/VM merged config context is cached on the object; invalidated on upstream changes and repopulated by a non-blocking background job; reads fall back to on-demand rendering during the window; always included in REST representations (`DeviceWithConfigContextSerializer` merged into base; `?exclude=config_context` ignored).
- **Snapshot-aware event rule conditions:** conditions can inspect pre/post-change snapshots (`changed`/`unchanged` operators, `snapshots.prechange.<attr>` / `snapshots.postchange.<attr>` dot paths); new `regex` operator; unresolvable attributes fail closed with a logged error.
- **Plugin API additions:** `GenericObjectChoiceField`/`GenericObjectFormMixin`; plugins can register custom Jinja filters and inject config-template context; plugins can extend core GraphQL types with fields/filters; plugins can register custom Event Rule action types by subclassing `EventRuleAction`.
- **Other notable v4.7 enhancements:** `end_of_life` date field on device types and module types (lifecycle planning); bulk import of cables with multiple terminations per side; interface `mac_address` field now writable (creates/updates primary MAC); read-only `is_primary` on MAC address API; record/display of background job execution time; `WEBHOOK_DEFAULT_TIMEOUT` (mandatory lower if `RQ_DEFAULT_TIMEOUT` ≤ 60s); `BULK_UPDATE_CHUNK_SIZE` config parameter; NetBox releases published to production PyPI; denormalized fields maintained via PostgreSQL triggers instead of Python signals (performance).
- **Deprecations (to be removed in v5.0):** core **custom scripts deprecated in favor of a dedicated plugin** (will be removed in v5.0); `JINJA2_FILTERS` renamed to `JINJA_FILTERS`; Rack `form_factor`/`width`/`outer_*` fields deprecated — rack type assignment will become mandatory; `protocol`/`ports` on services deprecated in favor of `port_mappings`; MPTT `NestedGroupModel` deprecated in favor of `NestedLtreeGroupModel`.
- **v4.7.1 (2026-09-15)** `[official]`: fixes pg_dump-restore of ltree cascade triggers (v4.7.0 dumps restored without them; upgrading reinstalls triggers but does not repair already-stale values — repair procedure documented); adds default module type profile for transceivers; adds InfiniBand 2X interface types for HDR/NDR/XDR (HDR100, NDR200, XDR400).
- **Stack upgrade:** Django 6.1 in v4.7; django-tables2 3.0; social-auth-app-django 6.0 / social-auth-core 5.1 (test SSO before upgrading); migration from django-pglocks to django-pgware; mail settings now via `MAILERS` (individual `EMAIL_*` settings no longer defined); `EMAIL['SERVER']` mandatory to send mail; custom-link templates receive a sanitized request (only id/path/path_info/method/GET/user); URL custom fields validated against `ALLOWED_URL_SCHEMES` (scheme-less values stored as absolute https URLs).

---

