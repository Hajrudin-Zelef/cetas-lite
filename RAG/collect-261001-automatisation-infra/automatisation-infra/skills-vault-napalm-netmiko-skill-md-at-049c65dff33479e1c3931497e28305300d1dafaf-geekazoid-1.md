---
id: collect-261001-automatisation-infra/automatisation-infra/skills-vault-napalm-netmiko-skill-md-at-049c65dff33479e1c3931497e28305300d1dafaf-geekazoid-1
title: "skills-vault-napalm-netmiko-skill-md-at-049c65dff33479e1c3931497e28305300d1dafaf-geekazoid80-skills-"
domain: automatisation-infra
role: reference
task: reference
actors: ["AWS", "Anthropic", "Apple"]
dates: []
keywords: ["agents", "apache", "aws", "claude", "license"]
source: docs/RAG/collect-261001-automatisation-infra/skills-vault-napalm-netmiko-skill-md-at-049c65dff33479e1c3931497e28305300d1dafaf-geekazoid80-skills-.md
source_anchor: ""
source_lines: [1, 70]
sha256: d0ea7422ba05b385064278f53537befc74b3f194042434183511a8b4a35a8c77
---

# skills-vault-napalm-netmiko-skill-md-at-049c65dff33479e1c3931497e28305300d1dafaf-geekazoid80-skills-

| name | napalm-netmiko | 
|---|---|
| description | Use for any multi-vendor network device automation work that uses NAPALM, Netmiko, or both. Triggers include "napalm", "netmiko", "ConnectHandler", "send_command", "send_config_set", "device_type", "napalm get_facts", "napalm load_replace_candidate", "napalm load_merge_candidate", "compare_config", "commit_config", "discard_config", "ssh to a network device from python", "multi-vendor device library", "vendor-agnostic config push", "automate cisco / juniper / arista / panos with python (no ansible)", "transport layer for network automation", "network device CLI scraping". Covers the NAPALM driver model with idempotent replace and merge config pushes, the Netmiko transport with send_command and send_config_set, per-vendor quirks (Cisco IOS / IOS-XE / NX-OS / IOS-XR / ASA, JunOS, EOS, PANOS, FortiOS, F5, Linux), authentication patterns sourced from secrets-hygiene rather than YAML, single-threaded concurrency caveat (use nornir-automation for parallel fan-out), error handling (NetMikoTimeoutException, NetMikoAuthenticationException, NAPALMException family), and the diagnose-first, read-only-getter-before-state-change discipline. Self-authored from public NAPALM and Netmiko documentation; no upstream third-party Claude skill exists for either tool. Customised body, Apache-2.0. | 
| license | Apache-2.0 | 
| metadata |  | 

| version | 
|---|
| 1.0.0 | 

Multi-vendor network device automation in Python. Two libraries, one skill, because the decision is rarely "one or the other"; most fleets use Netmiko as the SSH transport and NAPALM as the abstraction over getters and config push, with NAPALM internally driving Netmiko on platforms where there is no native API.

**Skill marker**: When applying this skill, begin your reply with `[skill: napalm-netmiko]` on its own line so the transcript shows the skill fired. If multiple skills fire on the same reply, emit each marker on its own line at the top: transparency over neatness.


If a `CLAUDE.md` or `AGENTS.md` exists in the working directory, read it first to understand the automation estate (target vendors, driver choices, credential vaulting, run-mode conventions) before scripting. Only ask the user for information not already covered or specific to this task.

Before scripting, understand:

1. 
**Target estate**
  - Vendor(s) and OS version(s) (NAPALM driver supported, or Netmiko-only)?
  - Read-only getter, config diff / replace, or destructive push?
  - Single device, batched fleet, or rolling change?
2. 
**Library and dependencies**
  - NAPALM and Netmiko versions pinned?
  - Connection method (SSH, telnet legacy, NETCONF for JunOS / IOS-XR)?
  - Existing wrapper or orchestrator (Nornir, custom CLI)?
3. 
**Change posture**
  - Maintenance window and rollback plan?
  - Pre / post evidence the reviewer will expect?
  - Secrets path (env, Vault, runbook-specific)?

1. **Diagnose-first, read-only before state-change.** Every interaction starts with a NAPALM getter (`get_facts` ,`get_interfaces_counters` ,`get_config` ) or a Netmiko`send_command("show ...")` to capture baseline evidence BEFORE any`load_replace_candidate` or`send_config_set` . Never push config without a prior`compare_config` diff narrated to the user.
2. **Credentials from the secret store, never from YAML or inline.**`username` and`password` keys come from environment, vault, AWS Secrets Manager, HashiCorp Vault, or the platform's secret store. Testbed YAML and inventory files contain references, not values. See`secrets-hygiene` .
3. **Single connection at a time per device.** NAPALM and Netmiko are synchronous and single-threaded. For fleet fan-out (more than one device in flight) use`nornir-automation` or`pyats-network-automation` . Do not spawn ad-hoc threads around`ConnectHandler` ; the libraries are not thread-safe per-device.
4. **NAPALM `load_replace_candidate` is destructive by default.** It uploads a full candidate config and replaces running on commit. Always run`compare_config` first; surface the diff via AskUserQuestion for any production device before`commit_config` .`discard_config` if the diff is wrong.
5. **Netmiko `send_command_timing` over `send_command` for prompts that change mid-session** (banner accepts, paging confirmations,`wr mem` Y/N,`reload` at?).`send_command` waits for the device prompt regex to return; if the prompt has changed the call hangs to the read timeout.
6. **Always close the connection.** Use a`with` context manager for Netmiko (`ConnectHandler` ) or call`device.close()` for NAPALM in a`finally` block. Leaked sessions hold VTY lines and exhaust device session limits (Cisco IOS defaults to 5).
7. **No state-changing command without rollback.** NAPALM has`rollback()` (Junos and IOS-XR have native rollback; on IOS / NX-OS / EOS NAPALM emulates it via the previous config snapshot). Netmiko has no rollback; capture the pre-change config explicitly via`send_command("show running-config")` and store it before the change.

| Scenario | Pick | Why | 
|---|---|---|
| Read-only inventory across multi-vendor fleet | **NAPALM** | Normalised getter outputs (same dict shape across IOS / JunOS / EOS); useful for ingest into NetBox / CMDB. | 
| Config push with diff and rollback | **NAPALM** | `load_replace_candidate` +`compare_config` +`commit_config` +`rollback` is the canonical idempotent pattern. | 
| One-off `show` command with vendor-specific output | **Netmiko** | NAPALM normalises away vendor detail; raw Netmiko returns the unmodified CLI output. | 
| Bespoke interactive workflow ( `copy ftp:` with prompts, ROMmon recovery) | **Netmiko** | NAPALM is contract-driven; Netmiko exposes the raw session for prompt-by-prompt control via `send_command_timing` and`read_until_pattern` . | 
| Vendor not supported by NAPALM (FortiOS, F5, Aruba, MikroTik, etc.) | **Netmiko** | Netmiko's `device_type` registry covers ~80 platforms; NAPALM's first-party drivers cover ~6. | 
| Idempotent multi-vendor change automation | **NAPALM (driving Netmiko underneath)** | NAPALM IOS / NX-OS / EOS drivers use Netmiko as transport; you get the abstraction without losing the SSH path. | 

```
from napalm import get_network_driver
driver = get_network_driver("ios")           # or "iosxr", "nxos", "junos", "eos", "panos"
device = driver(
    hostname="r1.example.net",
    username=os.environ["NETOPS_USER"],
    password=os.environ["NETOPS_PASS"],
    optional_args={"port": 22, "transport": "ssh", "secret": os.environ["NETOPS_ENABLE"]},
)
device.open()
try:
    facts = device.get_facts()
    interfaces = device.get_interfaces()
finally:
    device.close()
```
First-party drivers (commit-supported): `ios`, `iosxr`, `nxos`, `nxos_ssh`, `junos`, `eos`, `panos`. Community drivers (varying maturity): `fortios`, `aruba`, `huawei`, `mikrotik`, `ros` (RouterOS), `linux`, `f5`, `srl` (Nokia SR Linux). Community drivers live in `napalm-<vendor>` packages.

