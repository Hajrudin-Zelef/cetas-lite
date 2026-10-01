---
id: collect-261001-automatisation-infra/automatisation-infra/skills-vault-napalm-netmiko-skill-md-at-049c65dff33479e1c3931497e28305300d1dafaf-geekazoid-4
title: "skills-vault-napalm-netmiko-skill-md-at-049c65dff33479e1c3931497e28305300d1dafaf-geekazoid80-skills-"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-automatisation-infra/skills-vault-napalm-netmiko-skill-md-at-049c65dff33479e1c3931497e28305300d1dafaf-geekazoid80-skills-.md
source_anchor: ""
source_lines: [226, 237]
sha256: 12caf21c317053572be850c5f6d38dd7f05b0f726abca96a9481c680ad0dcf2c
---

# skills-vault-napalm-netmiko-skill-md-at-049c65dff33479e1c3931497e28305300d1dafaf-geekazoid80-skills-

1. About to call `commit_config()` on a production device without a narrated diff in the conversation.
2. About to pass a literal password string to `ConnectHandler` or NAPALM`password=` .
3. About to spawn a `ThreadPoolExecutor(max_workers=>20)` against unknown device population.
4. About to use `device_type="<vendor>_telnet"` for any device that is not a console-server reach to legacy hardware.
5. About to catch `Exception` and`pass` inside a config-push loop.
6. About to push `load_replace_candidate` with a candidate that omits sections present in the current running config (full replace; you may delete VTY ACLs, AAA, banner, etc.).
7. About to call `enable()` against a device whose`secret` is`None` because nobody set it; the call hangs to read timeout.
8. About to use NAPALM `rollback()` on IOS / NX-OS / EOS more than one commit after the change; NAPALM's emulated rollback window is one commit.
9. About to chain `send_config_set` calls without checking the returned output for`% Invalid input` /`error:` /`% Incomplete command` ; NAPALM raises, raw Netmiko returns the error in the output.
10. About to commit changes during a maintenance-window-adjacent boundary without coordinating against parallel automation runs (Nornir + ad-hoc + ServiceNow change windows can all hit the same device at once).

NAPALM gives you a normalised, idempotent multi-vendor abstraction with diff and rollback. Netmiko gives you the raw SSH transport when the abstraction does not fit. Use NAPALM by default for inventory and idempotent change; drop to Netmiko for vendors NAPALM does not cover or workflows that need prompt-by-prompt control. Both are synchronous per device; for fleet fan-out, hand the work to `nornir-automation`. Credentials always come from the secret store; configs always diff before commit; getters always re-run after commit; the response contract from `multi-vendor-network-ops` always applies on production change.
