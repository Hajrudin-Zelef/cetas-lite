---
id: collect-250926-servers-hardware/servers-hardware/how-to-serve-local-llms-anywhere-secure-remote-access-with-cloudflare-and-unsloth-unsloth--2
title: "how-to-serve-local-llms-anywhere-secure-remote-access-with-cloudflare-and-unsloth-unsloth-documentat"
domain: servers-hardware
role: reference
task: reference
actors: ["Unsloth"]
dates: []
keywords: ["mcp"]
source: docs/RAG/clean4/how-to-serve-local-llms-anywhere-secure-remote-access-with-cloudflare-and-unsloth-unsloth-documentat.md
source_anchor: ""
source_lines: [163, 271]
sha256: fccaac1d59f3af632129143410ac0ca4d7806aafae71ec56599c237ae17938c5
---

# how-to-serve-local-llms-anywhere-secure-remote-access-with-cloudflare-and-unsloth-unsloth-documentat

`run_server` suppresses the launch-path tunnel on Colab, since Colab already proxies the port. Ask for a shareable public link explicitly:

Colab refuses to open the link while the admin account still has its bootstrap password, that credential is visible to anyone who can load the page. Log in, change the password, then re-run. Any failure collapses to "no link" and the Colab proxy keeps working.

Publishing Unsloth means anyone holding the URL **and** a credential can use it.

- **Server-side tools run as your user.** Web search, Python and terminal execution are on by default, so anyone reaching the server with your API key can run code on that machine. Pass`--disable-tools` when exposing Unsloth, and keep the API key private. Every network-reachable launch says this in the banner.
- **The admin password gate** above is not optional on a launch-managed tunnel.
- **Rate limiting stays per-visitor.** The tunnel terminates at`127.0.0.1` , so every tunneled caller would otherwise share one bucket. When the socket peer is loopback, Unsloth honours Cloudflare's`CF-Connecting-IP` , which the edge sets and a tunneled client cannot forge. Behind your own reverse proxy, opt into`X-Forwarded-For` with`UNSLOTH_STUDIO_TRUST_FORWARDED=1` .
- **Local stdio MCP servers are revoked while a tunnel is live.** stdio MCP is auto-enabled only as a loopback convenience; a remote connector breaks that trust boundary, so the auto-default turns itself off. An explicit`UNSLOTH_STUDIO_ALLOW_STDIO_MCP=1` still wins.
- **The URL is random and disposable.** Quick Tunnels get a new hostname on every start. Treat the URL itself as a secret, and remember there is no way to pin it.

These are the exact statuses surfaced in the Remote access card. Anything else is collapsed to the generic `Cloudflare tunnel failed`.

`cloudflared is unavailable`

Not on `PATH`, not cached, and the download failed.

Check outbound access to `github.com`, or install `cloudflared` yourself so it is on `PATH`.

`cloudflared did not produce a URL`

The process exited before minting one.

Usually no outbound network. Retry; check a proxy or egress filter.

`cloudflared did not register a connection`

URL minted, but no edge connection, after the `http2` retry.

Your network blocks QUIC *and* HTTP/2 to Cloudflare's edge.

`Cloudflare URL was not reachable`

Registered, but the health probe never answered through the public URL.

Transient edge/DNS propagation; press Start again.

`cloudflared exited`

The connector died while online.

Start again; check for an OOM killer or a process supervisor reaping children.

`cloudflared could not be stopped`

Termination was never confirmed, so the slot is still held.

Press Stop again; if it persists, restart Unsloth.

Other symptoms:

**--secure**
 **exits immediately.** The tunnel failed; the message names`--no-secure` as the alternative. Fix connectivity first, since`--no-secure` publishes a raw port.
- **Start is greyed out in the UI.** Read the message under the card, it names the block reason from the table above.

`UNSLOTH_STUDIO_HOME`

Install root; the `cloudflared` cache lives in `<root>/bin`.

`UNSLOTH_STUDIO_BOOTSTRAP_TIMEOUT`

Seconds before an unsecured public launch shuts itself down. Default `3600`; `0` disables (and makes a headless public launch refuse).

`UNSLOTH_STUDIO_PASSWORD`

Initial admin password for headless launches.

`UNSLOTH_STUDIO_DISABLE_PUBLIC_CHECK`

`1` skips the third-party raw-port reachability probe on wildcard binds.

`UNSLOTH_STUDIO_TRUST_FORWARDED`

`1` honours `X-Forwarded-For` behind your own reverse proxy.

`UNSLOTH_STUDIO_ALLOW_STDIO_MCP`

`1` keeps stdio MCP servers enabled even with a tunnel live; `0` force-disables.

All four require a **UI session,** an API key is rejected with `403 Remote access requires a UI session.` A refused operation returns `409` with the block reason as the detail.

`GET`

`/api/settings/remote-access`

Current state, URL, owner, `can_start` / `can_stop`, `block_reason`.

`POST`

`/api/settings/remote-access/start`

Schedule a settings-owned start. Idempotent.

`POST`

`/api/settings/remote-access/stop`

Schedule a settings-owned stop. Does not change the auto-start preference.

`PUT`

`/api/settings/remote-access/auto-start`

`{"enabled": true|false}`. Rejected on Colab.

Last updated

Was this helpful?
