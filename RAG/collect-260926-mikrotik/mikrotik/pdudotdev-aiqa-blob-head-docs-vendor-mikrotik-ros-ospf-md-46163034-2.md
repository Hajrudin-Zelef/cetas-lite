---
id: collect-260926-mikrotik/mikrotik/pdudotdev-aiqa-blob-head-docs-vendor-mikrotik-ros-ospf-md-46163034-2
title: "Reset interface-template timer properties"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/pdudotdev-aiqa-blob-head-docs-vendor-mikrotik-ros-ospf-md-46163034.md
source_anchor: ""
source_lines: [103, 144]
sha256: f6e4596fbb176885aba4cf23f75db5d38db1c0cf26b99c8e03305aff8ffe7361
---

# Reset interface-template timer properties

General rule: RouterOS uses an object-based model. Revert approaches depend on whether you're resetting a property on an object or removing an object entirely.
- Reset a property to default: set <id> <property>= (empty value) — e.g.,set 0 hello-interval= resets hello-interval to default on interface-template #0
- Remove an object: remove <id> — deletes the object entirely (use for area-range, static-neighbor, etc.)
- Changes take effect immediately (no commit step).
# Reset interface-template timer properties
/routing ospf interface-template
set <id> hello-interval=          # reset to 10s
set <id> dead-interval=           # reset to 40s
set <id> retransmit-interval=     # reset to 5s
set <id> transmit-delay=          # reset to 1s
set <id> cost=                    # reset to auto
set <id> priority=                # reset to 128 (RouterOS 7 default; note: NOT 1)
# Remove area type
/routing ospf area
set <id> type=default             # revert stub/nssa to normal
set <id> no-summaries=no          # remove totally-stubby flag
# Remove area range
/routing ospf area/range
remove <id>                       # delete the range object
# Remove authentication
/routing ospf interface-template
set <id> auth=none                # disables authentication
set <id> authentication-key=     # clears the key
Non-obvious exceptions:
| Scenario | Correct sequence | Gotcha | 
|---|---|---|
| Revert area type to normal | set <id> type=default | Area type field accepts default ,stub ,nssa — notnormal . | 
| Reset DR priority | set <id> priority= (empty) orset <id> priority=128 | Default is 128 in RouterOS 7, changed from 1 in RouterOS 6. | 
| Revert redistribution | set <instance-id> redistribute= (empty) | Redistribution is a comma-separated list on the instance object. Setting empty string removes all. To remove one source: set to the remaining list. | 
| Revert authentication | set <id> auth=none thenset <id> authentication-key= | Must unset both auth type and key. | 
- Forgetting without-paging causes SSH sessions to hang waiting for user input.
- The +ct suffix on username (e.g.,admin+ct ) disables colors and auto-completion for clean output parsing in automation scripts.
- RouterOS 7 OSPF configuration is completely restructured from RouterOS 6. Old ROS6 commands (/routing ospf network add ) do not work. The entire config model changed to instance/area/interface-template objects.
- DR priority default changed from 1 (ROS v6) to 128 (ROS v7). Migrated networks with strict DR election priorities may behave differently after upgrade.
- type=ptp nottype=point-to-point .ptmp notpoint-to-multipoint .ptp-unnumbered is a RouterOS-specific type for IP unnumbered links.
- terse flag onprint gives compact tabular output;detail gives verbose key-value output. Neither is the default — plainprint uses an intermediate format.
- LSA database uses /routing ospf lsa , notdatabase . The keyword islsa everywhere.
- Interface-template networks uses CIDR notation and matches by subnet, not by interface name. Useinterfaces to match by name or interface list.
- There is no show keyword — RouterOS uses/routing ospf neighbor print style. All OSPF menus are under/routing ospf/ .
- originate-default interacts without-filter-chain in subtle ways: when set toalways orif-installed , OSPF creates a synthetic default route and runs it through the filter chain, but the filter action (accept/reject) is ignored — the default is always originated. The filter can only set attributes.
- No per-VRF show command variants. All OSPF instances appear together; filter output by instance name to isolate VRF data.
- Area range with advertise=no suppresses the summary but also installs a blackhole route. This is the equivalent of IOSnot-advertise onarea range .
