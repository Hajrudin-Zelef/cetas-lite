---
id: collect-261001-automatisation-infra/automatisation-infra/skills-vault-napalm-netmiko-skill-md-at-049c65dff33479e1c3931497e28305300d1dafaf-geekazoid-2
title: "skills-vault-napalm-netmiko-skill-md-at-049c65dff33479e1c3931497e28305300d1dafaf-geekazoid80-skills-"
domain: automatisation-infra
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-automatisation-infra/skills-vault-napalm-netmiko-skill-md-at-049c65dff33479e1c3931497e28305300d1dafaf-geekazoid80-skills-.md
source_anchor: ""
source_lines: [71, 152]
sha256: 5ebdfd0f3b97efae641756b203655a82c1fe68b7f538d25d58db218e15f5f8f2
---

# skills-vault-napalm-netmiko-skill-md-at-049c65dff33479e1c3931497e28305300d1dafaf-geekazoid80-skills-

| Getter | Returns | Use case | 
|---|---|---|
| `get_facts()` | hostname, model, serial, OS version, uptime, interface list | Inventory baseline, version audit | 
| `get_interfaces()` | per-interface state, MAC, MTU, speed, last_flapped | Interface-down sweep | 
| `get_interfaces_counters()` | rx / tx / errors / drops counters | Error-rate baseline | 
| `get_lldp_neighbors()` | per-port LLDP neighbour name + port | Topology discovery | 
| `get_arp_table()` | IP -> MAC -> interface mapping | L2 / L3 reconciliation | 
| `get_mac_address_table()` | MAC -> VLAN -> port | L2 troubleshooting | 
| `get_route_to(destination, protocol=None)` | per-protocol route entries with next-hop, metric, age | Routing decision tree (pair with `bgp-analysis` and`igp-routing-analysis` ) | 
| `get_bgp_neighbors()` | per-peer state, prefixes received / sent, uptime | BGP audit baseline | 
| `get_config(retrieve="all", sanitized=True)` | running, startup, candidate config strings | Pre-change snapshot; `sanitized=True` strips passwords for safe storage | 
| `get_environment()` | CPU, memory, fans, temperature, PSU | Health monitoring | 

The whole getter list is at `napalm.readthedocs.io/en/latest/support/`; coverage varies by driver, check the support matrix before assuming a getter exists.

```
device.open()
try:
    # 1. Capture pre-change baseline (mandatory; even with NAPALM rollback).
    pre = device.get_config(retrieve="running", sanitized=False)
    # 2. Load candidate config.
    device.load_replace_candidate(filename="r1.candidate.cfg")
    # OR additive: device.load_merge_candidate(config="interface Gi0/1\n description LINK_TO_R2\n")
    # 3. Diff. NEVER skip this step for production.
    diff = device.compare_config()
    if not diff:
        device.discard_config()
        return  # No-op; nothing to do.
    print(diff)  # Surface to user via AskUserQuestion before commit on production devices.
    # 4. Commit.
    device.commit_config()
    # 5. Verify post-change. Re-read getters; compare against expectations.
    post_state = device.get_interfaces()
    # ... assertions ...
except Exception:
    device.discard_config()  # Drop the candidate if anything went wrong before commit.
    raise
finally:
    device.close()
```
Rollback after commit (NAPALM-native on Junos / IOS-XR; emulated by NAPALM on IOS / NX-OS / EOS by re-pushing the snapshot): `device.rollback()`. The rollback window on emulated platforms is one commit; subsequent commits overwrite the snapshot.

```
from netmiko import ConnectHandler
device = {
    "device_type": "cisco_ios",   # see netmiko.ssh_dispatcher for the full list
    "host": "r1.example.net",
    "username": os.environ["NETOPS_USER"],
    "password": os.environ["NETOPS_PASS"],
    "secret": os.environ["NETOPS_ENABLE"],
    "port": 22,
    "fast_cli": False,            # leave False unless you have profiled the impact
    "session_log": "r1.session.log",
}
with ConnectHandler(**device) as conn:
    conn.enable()                  # IOS / NX-OS / ASA: enter privileged mode
    output = conn.send_command("show ip interface brief")
    conn.send_config_set(["interface Gi0/1", " description LINK_TO_R2"])
    conn.save_config()             # vendor-specific: write memory / commit / etc.
```
Common `device_type` values: `cisco_ios`, `cisco_xe`, `cisco_xr`, `cisco_nxos`, `cisco_asa`, `cisco_ftd`, `juniper_junos`, `arista_eos`, `paloalto_panos`, `fortinet`, `f5_tmsh`, `f5_linux`, `hp_procurve`, `hp_comware`, `huawei`, `mikrotik_routeros`, `linux`. Append `_telnet` for Telnet (almost always wrong in 2026; only justified for console-server reach to legacy gear).

| Method | Returns / does | Notes | 
|---|---|---|
| `send_command(cmd, expect_string=None, read_timeout=10)` | Output until prompt regex matches | Default; use for stable `show` commands. | 
| `send_command_timing(cmd, delay_factor=1, last_read=2)` | Output by read timeout, no prompt expectation | Use when prompt may change (paging, Y/N confirmations, banner). | 
| `send_config_set(config_commands, exit_config_mode=True)` | Pushes a list of commands inside config mode | Returns the combined output; check for vendor error markers ( `% Invalid input` ,`error: ...` ). | 
| `send_config_from_file(filename, **kwargs)` | Same, sourced from file | Convenient for reviewable change packs. | 
| `enable()` | Enters privileged mode | Requires `secret` in the device dict for IOS / NX-OS / ASA. | 
| `save_config()` | Vendor-specific persist | Maps to `write memory` (IOS),`commit` (Junos),`copy run start` (NX-OS),`write` (FortiOS). | 
| `disconnect()` | Closes the SSH session | Use the `with` context manager instead; it calls this on exit. | 
| `read_until_pattern(pattern, read_timeout=10)` | Reads until the regex pattern appears | Build interactive workflows (ROMmon, image-copy progress). | 

- **SSH key preferred** over password where the platform supports it. Pass`use_keys=True, key_file="/path/to/key"` to NAPALM`optional_args` or to Netmiko device dict. Combine with`passphrase` if the key is encrypted;`passphrase` itself comes from the secret store.
- **Per-platform enable / privilege quirks:**  - Cisco IOS / NX-OS / ASA: separate enable secret; pass `secret=` ; call`enable()` for Netmiko or set in`optional_args` for NAPALM.
  - JunOS: no enable mode; class-based privilege via `class super-user` etc.
  - PAN-OS: API key preferred over username/password (`api_key=` in NAPALM`optional_args` ); Netmiko works against the CLI but the official panxapi route is faster.
  - FortiOS: use Netmiko `fortinet` ;`vdom` switch via`conn.send_command("config vdom")` if multi-VDOM.
- Cisco IOS / NX-OS / ASA: separate enable secret; pass 
- **MFA / TACACS / RADIUS:** Netmiko handles standard prompts. For Cisco TACACS where the second prompt is "Password:" expecting the same string, no extra config needed. For RSA token + PIN concatenation, set`password = pin + token_code` at call time (token captured outside Python; never persisted).
- **Inventory pattern:** YAML or TOML inventory holds device hostnames +`device_type` +`secrets_path` . Loader looks up`secrets_path` in the secret store at run time. NEVER plaintext credentials in the inventory file.

