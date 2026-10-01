---
id: collect-261001-cisco/cisco/mcguiretechnology-enterprisedataplatform-blob-head-docs-source-systems-cisco-mer-77619d08-2
title: "mcguiretechnology-enterprisedataplatform-blob-head-docs-source-systems-cisco-mer-77619d08"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["governance", "incident", "license"]
source: docs/RAG/collect-261001-cisco/mcguiretechnology-enterprisedataplatform-blob-head-docs-source-systems-cisco-mer-77619d08.md
source_anchor: ""
source_lines: [149, 181]
sha256: 183f921f26a409c45bde425399f312bde959c1e5157a8615acf8bee364092c52
---

# mcguiretechnology-enterprisedataplatform-blob-head-docs-source-systems-cisco-mer-77619d08

- Switch port utilization mart
- License and subscription mart
- Configuration change mart
- Alert and incident correlation mart
- Asset and CMDB reconciliation mart
These marts should support network operations, asset governance, service management, license planning, security review, and audit readiness.
Recommended checks:
- Collection run completed successfully.
- Organization, network, device, client, license, and event counts are within expected ranges.
- Source object identifiers are present and unique within each object type.
- Device records include serial number, model, product type, and network assignment where available.
- Device status observations resolve to known devices.
- Client observations resolve to known networks and observation windows.
- License records resolve to known organizations or devices where applicable.
- Webhook events are validated and deduplicated where webhooks are enabled.
- API rate-limit responses are handled with backoff and retry behavior.
- Freshness meets the expected schedule.
Start with daily collection for organizations, networks, devices, licensing, configuration, admins, and templates.
Collect device status, uplink status, client observations, and event data more frequently when operational reporting requires it. Use webhooks for alert-driven workflows, but keep scheduled collection for reconciliation.
The Cisco Meraki Dashboard connector runbook should include:
- How to enable Dashboard API access
- How to create or validate API key access
- How to confirm organization and network permissions
- How to run a connection test
- How to run a manual collection
- How to review collection counts
- How to handle API rate limits and retry behavior
- How to configure and validate webhooks
- How to rotate API keys
- How to disable the connector safely
Meraki API coverage depends on product type, licensing, organization settings, and Dashboard permissions. Document collected product families and endpoint coverage for each implementation.
Client telemetry can include sensitive endpoint and user activity signals. Apply appropriate classification, retention, and access controls before making client-level detail broadly available.
Some operational telemetry is better collected through monitoring tools, syslog, SNMP, NetFlow, or SIEM pipelines. Use Dashboard API data as the cloud management source of authority and combine it with telemetry platforms where needed.
