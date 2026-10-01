---
id: collect-261001-automatisation-infra/automatisation-infra/skills-vault-napalm-netmiko-skill-md-at-049c65dff33479e1c3931497e28305300d1dafaf-geekazoid-3
title: "skills-vault-napalm-netmiko-skill-md-at-049c65dff33479e1c3931497e28305300d1dafaf-geekazoid80-skills-"
domain: automatisation-infra
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-automatisation-infra/skills-vault-napalm-netmiko-skill-md-at-049c65dff33479e1c3931497e28305300d1dafaf-geekazoid80-skills-.md
source_anchor: ""
source_lines: [153, 225]
sha256: d73700ed808b6f805fe778e6a67dfb0b0cae45b7dc39c6657bcb0cb04dbc2b44
---

# skills-vault-napalm-netmiko-skill-md-at-049c65dff33479e1c3931497e28305300d1dafaf-geekazoid80-skills-

| Vendor / OS | Quirk | Mitigation | 
|---|---|---|
| Cisco IOS-XE | `terminal length 0` paging fix is set automatically by Netmiko, but config-mode commands still page on`show running-config interface` | Use NAPALM `get_config` instead, or wrap in`terminal length 0` . | 
| Cisco NX-OS | `show running-config` truncates with`! NX-OS image file is: bootflash:///nxos.X.X.X.bin` header that breaks naive diff | NAPALM `get_config(sanitized=True)` strips this; for Netmiko, post-process. | 
| Cisco IOS-XR | Two-stage commit: `commit` may trail`Uncommitted changes found, commit them?` if config is dirty | NAPALM handles via `commit_config()` ; raw Netmiko needs`send_command_timing("commit\n", expect_string="]:")` . | 
| Cisco ASA | Privilege levels matter; `enable` may be required even for some`show` commands depending on TACACS profile | Always call `conn.enable()` early. | 
| JunOS | Config push requires `configure private` (concurrent-edit safe) or`configure exclusive` (locks) | NAPALM uses `configure private` ; if you need`exclusive` (long change windows), use Netmiko + explicit`send_config_set(["configure exclusive", ...])` . | 
| JunOS | `set system services ssh root-login` may be`deny` ; root SSH disabled by default | Use a non-root account; never re-enable root SSH. | 
| Arista EOS | `eAPI` (HTTPS+JSON) is faster than SSH for read-heavy work | NAPALM `eos` driver uses eAPI by default; pass`optional_args={"transport": "ssh"}` to force SSH if eAPI is disabled. | 
| PAN-OS | API key has no inactivity expiry but rotates on admin password change | Capture key once via `keygen` ; refresh on auth failure. | 
| FortiOS | Per-VDOM context; `config vdom / edit root` switches scope mid-session | Track current VDOM in your wrapper; assume nothing. | 
| F5 BIG-IP | iControl REST is the modern path; tmsh CLI is the Netmiko `f5_tmsh` route | Prefer iControl REST for new code; Netmiko only for break-glass. | 
| Linux network device (Cumulus, SONiC, FRR-on-debian) | Netmiko `linux` works but you lose the network-device prompt assumptions | Use raw paramiko or fabric for these; Netmiko adds little. | 

NAPALM and Netmiko are SYNCHRONOUS PER DEVICE. They are NOT thread-safe on the same `device` / `ConnectHandler` object across threads. For parallel fan-out across many devices:

- Preferred: `nornir-automation` (purpose-built parallel runner with NAPALM and Netmiko plugins).
- Acceptable: `concurrent.futures.ThreadPoolExecutor` with one connection per worker thread (each worker creates and disposes its own`ConnectHandler` ); cap the worker count at something the device infrastructure can handle (10 to 20 typical for production fleets, after surveying device session limits).
- Never: shared `ConnectHandler` across threads,`multiprocessing` against the same SSH session, asyncio against Netmiko (it is not async; wrapping it in`loop.run_in_executor` works but at that point use Nornir).
- For Cisco-dominant fleets with native parallel needs, `pyats-network-automation` 's`pcall` idiom is also a fit.

```
from netmiko.exceptions import NetMikoTimeoutException, NetMikoAuthenticationException
from napalm.base.exceptions import (
    ConnectionException,
    MergeConfigException,
    ReplaceConfigException,
    CommitError,
    SessionLockedException,
)
try:
    device.open()
except ConnectionException as e:
    # Network unreachable, port closed, SSH version mismatch, host key mismatch
    log.error("connect failed", host=host, error=str(e))
except NetMikoTimeoutException:
    # SSH banner timeout; usually firewall in path or device CPU pinned
    ...
except NetMikoAuthenticationException:
    # Credentials wrong; back off; do NOT retry in a loop, you will lock the account
    ...
```
NAPALM `MergeConfigException` and `ReplaceConfigException` indicate the candidate failed validation (syntax errors, unsupported commands). `CommitError` is post-validation rollback by the device itself. `SessionLockedException` (JunOS) means another user holds `configure exclusive`; back off and surface the lock holder.

- Capture pre-change `show running-config` (or NAPALM`get_config` ) explicitly; do not rely on NAPALM rollback being native.
- Run `compare_config` and surface the diff via AskUserQuestion before`commit_config` . The 9-element response contract from`multi-vendor-network-ops` applies (Summary / Goal / Devices / Diff / Risk / Pre-checks / Procedure / Post-checks / Rollback).
- Have the rollback command ready in a separate file (not just in NAPALM's emulated rollback).
- Verify post-commit via getters: `get_interfaces()` for link state,`get_bgp_neighbors()` for peering,`get_route_to()` for reachability. Do not claim "done" without fresh evidence; see`completion-gate` .
- For state-changing chunks, plan-mode entry should fire `engineering:deploy-checklist` per`plan-time-tooling` .

- `multi-vendor-network-ops` : umbrella; this skill is the Python transport / abstraction specialist underneath the diagnose-first methodology.
- `nornir-automation` : fleet orchestration; Nornir's NAPALM and Netmiko plugins drive this skill's libraries in parallel.
- `ansible-network-modules` : declarative alternative; YAML-first; better fit for change-managed environments where Python code review is heavier than playbook review.
- `pyats-network-automation` : Cisco-dominant alternative for fleet automation with structured Genie parsing.
- `bgp-analysis` /`igp-routing-analysis` : protocol-depth specialists; NAPALM getters feed their decision trees.
- `acl-rule-analysis` : when the change is an ACL push, the audit discipline lives there.
- `secrets-hygiene` : credential sourcing; never plaintext in YAML or scripts.
- `systematic-debugging` : Phase 1 boundary evidence often comes from a NAPALM getter or a Netmiko`show` .
- `completion-gate` : post-deploy verification gate; getter re-read is the evidence.
- `plan-time-tooling` : state-changing chunks fire`engineering:deploy-checklist` ; new automation framework choice fires`engineering:architecture` .
- `bash-defensive` : wrapper scripts around python entrypoints follow defensive-bash discipline.

1. Plaintext credentials in `inventory.yaml` or`device.yaml` . Always reference the secret store.
2. Sharing a `ConnectHandler` across threads. Not thread-safe; use Nornir or per-worker connections.
3. `send_command` against a vendor banner-accept prompt. Use`send_command_timing` instead.
4. NAPALM `load_replace_candidate` without`compare_config` first. The diff is the audit trail.
5. Forgetting to call `device.close()` /`conn.disconnect()` ; leaks VTY lines.
6. Assuming NAPALM has the getter you need without checking the support matrix.
7. Mixing NAPALM commit with manual config-mode interactions in the same session; NAPALM expects to own the session.
8. Telnet `device_type` in production. Plain text on the wire. Console-server reach to legacy gear is the only acceptable case; document why.
9. Setting `fast_cli=True` without profiling. The default is`False` for a reason; flipping it can return partial output on slow devices.
10. Catching all `Exception` and continuing; you will mask`CommitError` and silently push half-changes. Catch the specific exception classes.

