---
id: collect-261001-ia-llm/ia-llm/novotnyllc-agent-utilities-blob-head-plugins-agent-utilities-skills-unifi-networ-bbd82e93
title: "novotnyllc-agent-utilities-blob-head-plugins-agent-utilities-skills-unifi-networ-bbd82e93"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["guardrails"]
source: docs/RAG/collect-261001-ia-llm/novotnyllc-agent-utilities-blob-head-plugins-agent-utilities-skills-unifi-networ-bbd82e93.md
source_anchor: ""
source_lines: [1, 32]
sha256: 121b00728012b6f7c306300ce7ceb64c1621e031af88729cc6f94c5fa10f71d0
---

# novotnyllc-agent-utilities-blob-head-plugins-agent-utilities-skills-unifi-networ-bbd82e93

| name | unifi-network-api | 
|---|---|
| description | Use the official UniFi Network Integration API for managed UniFi console automation, especially firewall policies, policy ordering, zones, networks, clients, devices, DNS policies, traffic matching lists, and other Network app configuration. Use when the user asks Codex to inspect or change UniFi Network settings from UDM Pro. | 
Use UniFi Network's managed Integration API before browser UI automation for repeatable Network app configuration changes.
Set SKILL_DIR to the absolute directory containing this loaded SKILL.md; the shell working directory is not necessarily the skill directory. Before running the examples below, configure UNIFI_SSH_TARGET with the user's SSH target and UNIFI_CONSOLE_HOST with the user's console hostname. Guard these variables before any refresh or API request:
: "${SKILL_DIR:?Resolve the absolute directory containing this loaded SKILL.md and set SKILL_DIR}"
: "${UNIFI_SSH_TARGET:?Set UNIFI_SSH_TARGET to the configured SSH target}"
: "${UNIFI_CONSOLE_HOST:?Set UNIFI_CONSOLE_HOST to the configured console host}"
- Live console OpenAPI spec: /usr/lib/unifi/webapps/ROOT/api-docs/integration.json on the user-configuredUNIFI_SSH_TARGET .
- Saved snapshot: references/current-openapi.json from the operator's UDM Pro, currently Network10.5.67 .
- Notes and examples: references/api-usage.md .
- Multicast, mDNS, and IGMP debugging: references/multicast-troubleshooting.md .
Before relying on a schema, prefer refreshing or checking the live console spec:
: "${SKILL_DIR:?Resolve the absolute directory containing this loaded SKILL.md and set SKILL_DIR}"
: "${UNIFI_SSH_TARGET:?Set UNIFI_SSH_TARGET to the configured SSH target}"
scp "$UNIFI_SSH_TARGET:/usr/lib/unifi/webapps/ROOT/api-docs/integration.json" "$SKILL_DIR/references/current-openapi.json"
jq -r '[.info.title,.info.version,.openapi,.servers[0].url] | @tsv' "$SKILL_DIR"/references/current-openapi.json
Query the spec with jq; do not paste the whole OpenAPI file into context.
- Read references/api-usage.md for auth, base URLs, and guardrails.
- Inspect the exact endpoint/schema in references/current-openapi.json .
- Use an API key from the user or environment; do not mine browser cookies, local storage, or password/session stores.
- For mutations, fetch current state first, prepare the minimal change, and preserve unrelated fields/order.
- Prefer Integration API writes over browser UI control.
- Use browser control only when the API lacks coverage or to confirm a confusing UI-only behavior.
- Avoid direct Mongo writes for persistent configuration. Use Mongo only for read-only forensics or emergency rollback investigation.
- Never add persistent shell hooks or custom on-boot scripts for UniFi config.
After firewall or policy changes:
- Verify API/controller state with GET or, if needed, read-only Mongo.
- Verify user-defined ordering if rule order matters.
- Verify compiled gateway behavior over SSH with iptables-save ,ipset , or equivalent live firewall inspection.
- Remove any temporary runtime-only emergency rules once managed policy rules are compiled and confirmed.
For host-scoped region blocking, use explicit destination-scoped firewall policies rather than global CyberSecure region blocking when the user wants only one host affected.
