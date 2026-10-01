---
id: collect-261001-fortinet/fortinet/itential-assets-blob-head-fortinet-fortigate-readme-md-1d0ba3ff-2
title: "itential-assets-blob-head-fortinet-fortigate-readme-md-1d0ba3ff"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/itential-assets-blob-head-fortinet-fortigate-readme-md-1d0ba3ff.md
source_anchor: ""
source_lines: [108, 149]
sha256: 56fc1f464eb043ec17ceaa6e8ac75e7139659d0b0e654fd13f794100fc6ab52b
---

# itential-assets-blob-head-fortinet-fortigate-readme-md-1d0ba3ff

Deploying this driver: copy device-drivers/fortigate-rest/ into your own automation repo — the one Itential Gateway already tracks, or a new one dedicated to your drivers. Don't add this repo (itential/assets) as an Itential Gateway repositories source just to pull one driver; that clones the entire community assets repo, which is unnecessarily large for a single driver.
After copying, update the repositories entry and working-directory paths in import.yaml to point at your own repo, then either:
iagctl db import import.yaml --force
or copy just the services/decorators blocks into your Itential Gateway's existing import.yaml.
Dependencies: requests>=2.28.0
Defines the fortios device type in Config Manager. This parser must be imported
before the FortiGate - Simple golden configuration tree can be created or used. It
tokenizes FortiOS CLI configuration (config / edit / set / next / end blocks)
using the cisco-ios lexer template, treating each line as a sequence of words
delimited by whitespace, with quoted strings preserved as single tokens. It adds a #
comment rule on top of the base cisco-ios rules to correctly skip the #-prefixed
metadata header lines (#config-version=..., #buildno=...) that FortiOS always
emits at the top of a configuration export.
Import via Config Manager → Configuration Parsers → Import.
One golden configuration tree is provided. It ships with no device bindings — bind it to your devices in Config Manager after importing.
Device type: fortios
Baseline configuration using literal matching. Covers admin session timeout, password
policy (status + minimum length), NTP sync to a named server, syslog forwarding, and a
handful of security-hardening checks: disallows disabling the HTTPS admin redirect,
disallows a wide-open trusthost1 0.0.0.0 0.0.0.0 management ACL, and disallows
telnet/HTTP management access on port1. Use this as a starting point when all devices
in a group are expected to share identical configuration values with no variation.
Before importing: The fortios parser must be registered in Config Manager first
(see Configuration Parsers above). Then bind the root node's devices to match your
Inventory Manager device name(s).
Dependencies: fortios parser
A netmiko-based automation showcase — every workflow here uses Config Manager and Gateway Manager's generic device operations (backUpDevice, getDeviceConfig, runComplianceForDevice, sendConfig), which need a device wired to Option 1 (netmiko) since sendConfig requires real set-config support. fortigate-rest alone (Option 2) can't run these workflows as-is — see Choosing API vs SSH — and a Hybrid Setup if you want to mix in REST calls elsewhere.
The project contains 4 workflows across 3 folders:
| Folder | Workflow | What it does | 
|---|---|---|
| Golden Configuration | FortiGate - Run Compliance | Runs a golden config compliance check against a device and returns the graded report | 
| Interface Configuration | FortiGate - Interface Configuration | Renders and pushes an interface config (IP, management access, description) with a manual approval gate, pre/post checks, and a diff | 
| Interface Configuration | FortiGate - Delete Interface Configuration | Resets an interface back to unconfigured, with the same approval gate | 
| Backup Configuration | FortiGate - Backup Configuration | Backs up a device's running configuration into Config Manager and returns the live config text | 
Each workflow starts with a JSON form collecting a device name (Itential::fortigate1 — replace with your own Inventory::NodeName) plus, for the interface workflows, the interface/IP/mask/access/description values. The interface workflows also carry two Jinja2 templates (FortiGate Interface Config, FortiGate Interface Delete) and a MOP command template (FortiGate Interface Check) used for both the pre- and post-push checks.
FortiGate - Run Compliance and the golden config tree above are already wired together — its runComplianceForDevice task targets the exact treeId that Golden Configurations/FortiGate - Simple.json ships with, and Config Manager's golden config import preserves that ID verbatim. Import the fortios parser and the FortiGate - Simple tree (see above) before this project, bind the tree to your device, and the compliance workflow resolves with no manual editing.
Dependencies:
| Dependency | Notes | 
|---|---|
| A netmiko-wired FortiGate node | See Option 1 above. Update the device-name default in each JSON form (or just type your own at run time) | 
| cluster_1 Itential Gateway cluster | The interface workflows' sendConfig tasks target a cluster namedcluster_1 — update this in the canvas if yours is named differently | 
| fortios parser +FortiGate - Simple golden config tree | Import both (see Configuration Parsers / Golden Configurations above) and bind the tree to your device before running FortiGate - Run Compliance |
