---
id: collect-261001-ia-llm/ia-llm/sukramj-go-unifi2mqtt-blob-head-claude-md-6ddcce2b-2
title: "sukramj-go-unifi2mqtt-blob-head-claude-md-6ddcce2b"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-ia-llm/sukramj-go-unifi2mqtt-blob-head-claude-md-6ddcce2b.md
source_anchor: ""
source_lines: [104, 200]
sha256: a9a26dcee8084c6ff80f2021ee024d7d4e610b7763635ad4159124a1b08c142f
---

# sukramj-go-unifi2mqtt-blob-head-claude-md-6ddcce2b

naive test, and silently lose VLAN mapping and the device topology.
- Publish default_entity_id . It is the only entity_id seed Home
Assistant still reads. Without a seed, HA derives the entity_id from
the localisedname — which is how German entity_ids appear in a
de-configured install.object_id is not published: the MQTT
discovery schemas areextra=REMOVE_EXTRA and accept it on 0 of 32
platforms (HA 2026.9), so it is dropped on arrival, silently.
- The add-on has four files that must agree about its options:
config.yaml options,config.yaml schema,translations/{en,de}.yaml andscript/run.sh . Each disagreement fails differently and quietly;internal/config/addon_test.go checks all four.
- Topic suffixes are an API. Every key in internal/coordinator/topics.go doubles as the Home Assistant entity key and the translation-table
lookup. Renaming one orphans the entity and its history in every
existing installation. The same goes forunique_id anddefault_entity_id ininternal/hass —TestIdentifiersAreStable pins the exact strings.
- The MQTT command handler must never block. It runs inline in the client's read loop, the same goroutine that decodes acknowledgements and feeds the keep-alive watchdog. It parses and enqueues, nothing else — anything slower makes the watchdog declare a healthy connection dead.
- Retained commands are replays, not requests. The broker
re-delivers the last retained message per filter on every reconnect. A
stale mosquitto_pub -r would otherwise power-cycle a real port on
every daemon start.
- Never publish optimistic state. After a command the affected object is re-polled and the state comes from the console, so a failed command snaps the Home Assistant entity back instead of lying.
- A classic-layer failure must degrade, never propagate. The facade
switches the affected capability off and the official path keeps
running. internal/unifi/facade_test.go pins this; breaking it turns
"site health is missing" into "the bridge is down".
- The classic API reports errors in the body, at HTTP 200. An
expired session arrives as meta.rc = "error" with a 200 status. A
client checking only status codes decodes an empty data array and
treats a logged-out session as "the site has no clients" — forever.
- Web assets must use relative references. The Home Assistant
Ingress proxy serves the page under a generated path prefix; an
absolute /app.css or/api/state escapes it and 404s.TestAssetReferencesAreRelative pins this.
- Never interpolate console values as HTML. A device named
<img onerror=…> is a legal UniFi name. The UI builds text nodes
only, and a test rejectsinnerHTML outright.
- Ownership of a retained config is recorded, not inferred. The
startup reconcile may retract a discovery config only if this
process published that exact topic since it started — a grow-only
claim list (hass.Claims ), which a retraction does not undo. The
payload check (unique_id prefix and the bridge availability
topic) survives as the shape half and is documented as never again
sufficient on its own: two UniFi consoles bridged to one broker emit
byte-identical config topics,unique_id s, availability topics and
state topics for the whole site plane, so no predicate over a payload
can separate them and "looks like mine" deletes the neighbour's
entities. The stated cost is that a config left by an earlier run
is no longer cleared either; it is logged once ascoordinator.reconcile_unclaimed with the remedy.
- Discovery is per-entity, and that is a decision with a condition
attached. One retained config per entity at
homeassistant/<platform>/<node_id>/<object_id>/config — not a
device bundle. ADR 0070 phase 9 measured, built and proved the bundle
path (RenderHamqtt /HamqttBundles render it and are byte-equal to
what ships) and then declined to publish it: a bundle is one retained
topic per device, so two consoles that cannot be told apart would
replace each other's whole entity set rather than overwrite it key
by key. It becomes available when this bridge has an identity that
distinguishes two consoles without re-registering entities Home
Assistant has already registered. Readnotes/adr0070-phase9-measurement.md 's closing section before
reopening it — the retraction form and the new identity pull in
opposite directions, and getting it wrong publishes a fleet with no
entities and no error.
- Never sweep a class whose source has not reported. An empty
announced set means "not polled yet", not "gone".
internal/coordinator/reconcile.go gates per class and treats a
device poll returning zero devices as not ready — an empty list is
usually a permission problem, and reading it as truth deletes every
device entity with its history.
- Never rebuild a topic in a second place. internal/hass receives
the topic layout through theTopics interface rather than
reconstructing it from config. Two copies drifting apart produce
entities that stay "unavailable" forever with nothing visibly wrong in
either package.
- Two surfaces, one facade. internal/unifi/integration is the
officially supported client and the default path;internal/unifi/classic is opt-in (CLASSIC_ENABLE ) and fills the
documented gaps. The coordinator only ever sees the combined facade
and theinternal/model types — never a raw API DTO.
- The classic API is undocumented and can break on any controller update. Every feature it powers must degrade gracefully: if the classic client is disabled or failing, the daemon keeps running on the official surface with those entities absent, never crashing.
- Rate limits are real. Honour 429 +Retry-After ; do not fan out
a per-device statistics call for every device on every tick without a
bounded worker pool.
- _id vs MAC. The Integration API keys objects by UUID, the
classic API by MongoDB_id , and both carry the MAC. The MAC is the
stable cross-API identity and the basis of MQTT topics and HAunique_id s — seeCONCEPT.md for the exact rule.
make release cross-compiles, bundles and extracts the notes; run it
locally first, since CI runs the identical command. Two checks that are
easy to skip and expensive to get wrong:
- docker build . — the Dockerfile is not exercised bymake check ,
so a broken one only surfaces when the release workflow runs.
- The add-on option set, schema: andscript/run.sh must agree. An
option missing fromrun.sh is silently ignored at runtime; one
missing fromschema: makes the Supervisor reject the config.
- CONCEPT.md is the design source of truth.
- README.md documents the MQTT topic layout, config
keys and quickstart paths (Docker, HA add-on, plain binary).
- changelog.md has the release history.
- config-template.yaml documents every
config field inline.
- addon/README.md andaddon/DOCS.md cover the Home Assistant add-on.
- The sibling project go-mtec2mqtt is the structural template for this repo — reuse its coordinator /
hass / state / web patterns, but not its LGPL headers.
