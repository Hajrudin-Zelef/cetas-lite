---
id: collect-250926-servers-hardware/servers-hardware/how-to-serve-local-llms-anywhere-secure-remote-access-with-cloudflare-and-unsloth-unsloth-
title: "how-to-serve-local-llms-anywhere-secure-remote-access-with-cloudflare-and-unsloth-unsloth-documentat"
domain: servers-hardware
role: reference
task: reference
actors: ["AWS", "Google", "Unsloth"]
dates: []
keywords: ["agent", "aws", "gpus", "mcp"]
source: docs/RAG/clean4/how-to-serve-local-llms-anywhere-secure-remote-access-with-cloudflare-and-unsloth-unsloth-documentat.md
source_anchor: ""
source_lines: [1, 271]
sha256: 41299353c733d3472417097aa0c4f5a2a433e888d4a9a2d3cd14cfa8a85ff6cc
---

# how-to-serve-local-llms-anywhere-secure-remote-access-with-cloudflare-and-unsloth-unsloth-documentat

Serve and deploy local AI models anywhere. You can securely access your local LLMs from another device over HTTPS using a Cloudflare tunnel via Unsloth. No Cloudflare account, domain setup, or port forwarding is required. Run private AI models on your own hardware, servers, or cloud GPUs and connect from anywhere like your phone.

Unsloth is an open-source project that allows you to train and run LLMs locally and with Cloudflare tunnel, you can access Unsloth from your mobile device, share access to a friend or coworker, host Unsloth on a server such as Google Colab, AWS, or even a personal server.

Unsloth can run 100% offline on your computer. Turn on remote access and you get a web link like `https://known-plates-desire-turkey.trycloudflare.com`. Open it on your phone, your laptop, or any browser, anywhere.

It is free. Nothing to sign up for or pay for and you never touch your router or firewall. Once Unsloth is installed, your computer reaches out to Cloudflare, and Cloudflare passes visitors back to it. The link is HTTPS, so the connection is encrypted.

Two ways to turn it on:

- **When you start Unsloth:** add`--secure` to your`unsloth studio` command
- **While Unsloth is running:** go to Settings → API →**Remote access** and press**Start**

Only one tunnel exists per Unsloth process, and only its owner can stop it.

Firstly we will need to download the Unsloth Desktop app.

1. Launch the app
2. Then go to Settings → Remote & LAN → Remote access.

Use this when Unsloth is already running and you want it on your phone or another machine. Press **Start**, then copy the **Remote URL** or scan the **QR** code. `Online` means the link already answers, not just that it was requested.

Three things worth knowing:

- Anyone with the URL and the password can sign in. Remote browsers log in as `unsloth` , set that under**Remote password** .
- The URL is new on every start and cannot be pinned.
- A raw `0.0.0.0` port stays open. Only`--secure` closes it.

If **Start** is greyed out, the card says why, usually the admin password still needs changing, or the tunnel belongs to the launch command.

`server_starting`

Unsloth is still starting.

`admin_password_change_required`

Set a remote password before exposing this server. *(desktop)* / Change the administrator password before exposing this server. In the desktop app, run `unsloth studio reset-password`. *(browser)*

`explicitly_disabled`

This launch used `--no-cloudflare`. Restart without it to enable remote access.

`launch_managed`

This tunnel is managed by the launch command.

`colab_managed`

This tunnel is managed by the Colab runtime.

`colab`

Remote access settings are managed by the Colab runtime.

Once a tunnel is online, two places switch to it automatically:

- **Settings → API → Usage examples** grows a**Secure HTTPS** toggle. On, every curl/Python/JavaScript snippet and coding-agent command is rewritten against the`trycloudflare.com` base instead of`localhost` . If you did not launch with`--secure` , an info tooltip reminds you:*"The 0.0.0.0 port is still reachable globally. For full security, launch Unsloth with* *--secure*
 *to expose only this HTTPS link."*
- The **API monitor** page's**Base URL** readout shows the tunnel origin, so a snippet copied there works from the remote device, and every request arriving through the tunnel is listed live.

After installing Unsloth manually, run in your terminal:

Unsloth stays bound to `127.0.0.1` and is published **only** through the tunnel. If the tunnel cannot come up, Unsloth **exits instead of falling back** to a raw port. The banner prints:

The raw port stays reachable on your network *and* a public Cloudflare URL is published. The banner warns you about exactly this, because it is the least private mode.

1. Open **Settings → API** .
2. Find the **Remote access** card.
3. Click **Start** .
4. When the status reads **Online** , copy the**Remote URL** or scan the**QR** code with your phone.

Both flags are accepted by `unsloth studio` (the plain-server path) and by `unsloth studio run`.

`--secure` / `--no-secure`

off

Publish **only** through Cloudflare. Forces a loopback bind, implies `--cloudflare`, and fails closed if the tunnel cannot start.

`--cloudflare` / `--no-cloudflare`

off

Also publish a public Cloudflare URL for a **non-****--api-only** **wildcard bind** (`0.0.0.0` or `::`). Has no effect on a loopback bind. `--no-cloudflare` forces it off but does **not** make a wildcard bind private.

Notes on flag handling:

- `--secure` ignores`-H` . If you pass another host it prints a note and binds`127.0.0.1` anyway.
- `--secure --no-cloudflare` is a contradiction and exits with code`2` .
- These flags belong to the plain-server path. Putting them *before* a subcommand (`unsloth studio --secure run ...` ) exits`2` with the corrected command, because Typer would otherwise silently drop them. Use`unsloth studio run --secure ...` .
- The choice is carried across Unsloth's internal re-exec as a tri-state ( `enabled` /`disabled` /`unset` ), so a stale`Docker ENV` or`systemd Environment=` can never re-enable a tunnel you opted out of on this invocation.

`unsloth studio`

this machine only

no

`unsloth studio --cloudflare`

this machine only

no, the flag is a no-op on loopback

`unsloth studio -H 0.0.0.0`

your network

no

`unsloth studio -H 0.0.0.0 --cloudflare`

your network

**yes**

`unsloth studio --secure`

this machine only

**yes**, and it is the only way in

`unsloth studio --api-only` (desktop backend)

as bound

no, unless `--secure`

Google Colab

Colab proxy

only via `start(cloudflare=True)`

The rule the backend applies: Colab never tunnels from the launch path; `--secure` always tunnels (even `--api-only`, for headless secure API serving); otherwise the tunnel starts only for a wildcard bind that is not `--api-only`.

If the tunnel does not come up, `--secure` refuses to keep running:

and exits `1`. This is deliberate: `--secure` means "no raw public port", so silently degrading to a loopback-only server that you believe is published would be worse than stopping.

On a wildcard bind Unsloth always states the tunnel's status, so a network-reachable launch is never silent:

- Cloudflare tunnel: ON. This is a PUBLIC internet URL: anyone who has it can reach this Unsloth.
- Cloudflare tunnel: ON. This Cloudflare URL is PUBLIC, and the raw port is also publicly reachable.
- Cloudflare tunnel: requested but failed to start.
- Cloudflare tunnel: OFF (default). / OFF (--no-cloudflare). / OFF for this mode.

Each variant adapts to whether the raw port was independently found reachable from the internet. That reachability probe contacts `ifconfig.me` and `check-host.net`; set `UNSLOTH_STUDIO_DISABLE_PUBLIC_CHECK=1` to skip both.

The first time a launch is about to publish Unsloth on a public URL while the admin account still holds its auto-generated bootstrap password, Unsloth stops and asks for a new one masked, confirmed, **before any server or tunnel exists**.

- **Terminal attached:** you are prompted. Aborting (Ctrl+C) refuses the launch.
- **No terminal:** Unsloth warns instead, never injects the bootstrap credential into the public page, and arms the bootstrap deadline, it shuts down after`UNSLOTH_STUDIO_BOOTSTRAP_TIMEOUT` (default 1 hour) unless the password is changed.
- **No terminal and no deadline** (`--api-only` , or`UNSLOTH_STUDIO_BOOTSTRAP_TIMEOUT=0` ): the launch is refused outright, because nothing would protect it.
**--secure**
 **with** **cloudflared**
 **provably unavailable:** refused with the seeded password preserved, so a failed tunnel cannot lock you out.

For headless setups, set the initial password non-interactively (only takes effect when none is set yet):

A literal `--password VALUE` is visible in `ps` and shell history, so prefer the env var or stdin. Rotate later with `unsloth studio reset-password`.

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
