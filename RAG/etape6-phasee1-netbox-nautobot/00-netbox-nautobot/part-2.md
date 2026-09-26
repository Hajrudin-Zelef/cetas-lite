---
id: etape6-phasee1-netbox-nautobot/00-netbox-nautobot/part-2
title: "Phase E1 — NetBox & Nautobot: Network Source of Truth (Research File) (part 2)"
domain: phase-e1-netbox-nautobot-network-source-of-truth-research-fi
role: deep-dive
task: reference
actors: []
dates: ["2026-09-15"]
keywords: []
source: docs/RAG/etape6_phaseE1_netbox_nautobot.md
source_anchor: ""
source_lines: [58, 77]
section: "Phase E1 — NetBox & Nautobot: Network Source of Truth (Research File)"
sha256: c0b50ed2a2d2bd290b695096f2e6a9e37a6790f48a7005025806eea0d5e8654c
---

# Phase E1 — NetBox & Nautobot: Network Source of Truth (Research File) (part 2)

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

