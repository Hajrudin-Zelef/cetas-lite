---
id: collect-260926-mikrotik/mikrotik/danikf-tik4mcp-blob-head-plugin-skills-mikrotik-firewall-skill-md-0ebefe2a
title: "danikf-tik4mcp-blob-head-plugin-skills-mikrotik-firewall-skill-md-0ebefe2a"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "parameters"]
source: docs/RAG/lot-mikrotik/forum/firewall-nat/danikf-tik4mcp-blob-head-plugin-skills-mikrotik-firewall-skill-md-0ebefe2a.md
source_anchor: ""
source_lines: [1, 66]
sha256: ba05a7d4d3b5bd770dfb73b0b364d7d720fe58babd006eeeb3aaac2c4e5bab36
---

# danikf-tik4mcp-blob-head-plugin-skills-mikrotik-firewall-skill-md-0ebefe2a

| name | mikrotik-firewall | 
|---|---|
| description | Understand, audit, and safely change the MikroTik RouterOS firewall via tik4mcp. Use when the user wants to read/explain firewall rules, review or harden the firewall, add/modify/reorder filter or NAT or mangle or RAW rules, set up masquerade or port-forwards, manage address-lists, or wants firewall best-practice recommendations. | 
Drive the firewall through the tik4mcp tools (mikrotik_command for reads and all writes; the curated
read tools for quick looks). This skill carries the logic and best practices; for exact field
semantics and the current authoritative ruleset, consult the MikroTik docs linked below — prefer them
over memory, since RouterOS details change between versions.
- Tables: filter (allow/deny),nat (address translation),mangle (marking),raw (pre-connection-tracking drops). Plus/ip/firewall/address-list and connection tracking.
- Chains: input = traffic to the router itself;forward = traffic through the router
(LAN↔WAN);output = traffic from the router. NAT usessrcnat /dstnat .
- Evaluation: rules run top-down, first match wins, per chain. Order is everything — a rule's effect depends entirely on what precedes it.
- Connection state: each packet is new /established /related /invalid /untracked .
Acceptingestablished,related,untracked early and droppinginvalid is the backbone of a stateful
firewall.
To read & explain a firewall: mikrotik_command with /ip/firewall/filter/print,
/ip/firewall/nat/print, /ip/firewall/mangle/print, /ip/firewall/raw/print,
/ip/firewall/address-list/print, /ip/firewall/connection/print. Walk the rules in order and
describe what each does and what reaches the bottom.
Mirror MikroTik's recommended ruleset (see Firewall & QoS case studies, incl. "Building Advanced Firewall"):
- Use interface lists WAN andLAN (/interface/list +/interface/list/member ) and match onin-interface-list /out-interface-list instead of raw interface names — rules survive topology
changes.
- input chain: acceptestablished,related,untracked → dropinvalid → accept ICMP (don't
over-filter it) → accept fromLAN (and/or a management address-list) → drop everything else.
- forward chain: FastTrackestablished,related (IPv4, big performance win) → acceptestablished,related,untracked → dropinvalid → accept LAN→WAN → drop WAN traffic that isn't
dst-NATed → drop bogon sources → drop the rest.
- raw /prerouting: drop bogons/reserved ranges and obviously bad packets before conntrack to
save CPU under load.
- NAT: use action=masquerade for a dynamic WAN address; useaction=src-nat only for a
static public IP. Port-forwards arechain=dstnat action=dst-nat .
- Don'ts: never omit the final "drop the rest" (protects newly-added interfaces); don't FastTrack
IPsec traffic without a policy-bypass rule; don't expose router services to WAN .
A wrong rule can lock you out instantly. Always:
- Confirm with the user before any filter/NAT change; restate the intended effect.
- Protect your own access first. If managing remotely, ensure an accept rule for your
management source precedes any newdrop . For a multi-rule change, run it throughmikrotik_safe_batch instead of separatemikrotik_command calls: it holds RouterOS Safe
Mode across the whole batch on one session, so if a rule locks you out (or the connection drops)
every change is rolled back automatically. Usecommit=false first as a dry-run to confirm the
router accepts the rules. (A singlemikrotik_command cannot do this — Safe Mode only spans one
session; see the safe-mode note inmikrotik-admin .)
- Order matters — be explicit. New filter rules append to the end by default (often after a
drop, so they never match). Use =place-before=<id> on add, or/ip/firewall/filter/move with=numbers=<id> =destination=<pos> . Always re-print and verify order after changing.
- Prefer reversible steps: add new rules with =disabled=yes , verify, then/ip/firewall/filter/enable . Disable rather than remove while testing.
Example — add a stateful baseline accept at the top of input:
/ip/firewall/filter/add · =chain=input · =action=accept ·
=connection-state=established,related,untracked · =place-before=0 · =comment=baseline est/rel
Example — masquerade for a dynamic WAN:
/ip/firewall/nat/add · =chain=srcnat · =action=masquerade · =out-interface-list=WAN
Example — apply a several-rule baseline atomically with auto-rollback via mikrotik_safe_batch
(commit=false first to dry-run, then commit=true):
steps = [
  { command: "/ip/firewall/filter/add", parameters: ["=chain=input","=action=accept","=connection-state=established,related,untracked","=place-before=0","=comment=baseline est/rel"] },
  { command: "/ip/firewall/filter/add", parameters: ["=chain=input","=action=drop","=connection-state=invalid","=comment=drop invalid"] },
  { command: "/ip/firewall/filter/add", parameters: ["=chain=input","=action=accept","=in-interface-list=LAN","=comment=allow LAN to router"] },
  { command: "/ip/firewall/filter/add", parameters: ["=chain=input","=action=drop","=comment=drop all else"] }
]
If any step is rejected — or you lose the connection partway — RouterOS reverts the whole set.
When asked to review or harden, report findings against these:
- input ends indrop (no implicit accept exposure)?invalid dropped? router services not reachable fromWAN ?
- Stateful accept present and near the top of both chains? FastTrack enabled on forward ?
- Masquerade vs src-nat correct for the WAN type? Any overly broad dst-nat /port-forwards?
- Management restricted to an address-list? Useful rules disabled or shadowed by an earlier match?
- Address-lists current; logging on key drops for visibility.
- Firewall & QoS (overview) · Packet Flow in RouterOS
- Filter · NAT · Mangle · Address-lists
- Common matchers & actions · Connection tracking · Case studies (incl. Building Advanced Firewall)
See the mikrotik-admin skill for transports, the router inventory, and the global safety rules.
