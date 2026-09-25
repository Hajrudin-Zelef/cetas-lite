---
id: collect-250926-servers-hardware/servers-hardware/how-to-serve-local-ai-models-from-any-device-on-your-network-with-unsloth-lan-access-unslo
title: "how-to-serve-local-ai-models-from-any-device-on-your-network-with-unsloth-lan-access-unsloth-documen"
domain: servers-hardware
role: reference
task: reference
actors: ["Unsloth"]
dates: []
keywords: ["latency"]
source: docs/RAG/clean4/how-to-serve-local-ai-models-from-any-device-on-your-network-with-unsloth-lan-access-unsloth-documen.md
source_anchor: ""
source_lines: [1, 101]
sha256: ea504ec55b179c22adf921016bd967ab016b10a31970b1776457db0557d08a61
---

# how-to-serve-local-ai-models-from-any-device-on-your-network-with-unsloth-lan-access-unsloth-documen

You can run local AI models on any device in your home or office with LAN access using Unsloth. With this enabled, local models can be accessed from a phone, laptop, or another computer on the same Wi-Fi or wired network without sending anything to the cloud. No public URL or internet connection is needed - your AI stays on your machine and your network.

LAN access is the fastest way to use a local model from another device. Traffic never leaves your network, so latency is low.

Turn on LAN access and Unsloth answers at your machine's network address, like `http://192.168.1.42:8888`. Open it on your phone, your laptop, or any device on the same network.

It is free and it works offline. Nothing leaves your network and you never touch your router or firewall. For access from outside your network, use remote access instead.

Two ways to turn it on:

- **When you start Unsloth:** add`-H 0.0.0.0` to your`unsloth studio` command
- **While Unsloth is running:** go to Settings → API →**Remote & LAN** , find the**LAN access** card, and press**Start**

Firstly we will need to download the Unsloth Desktop app.

1. Launch the app
2. Then go to Settings → API → **Remote & LAN** .

Use this when Unsloth is already running and you want it on another device in the same building. Press **Start** on the **LAN access** card. `Online` means the address already answers, not just that it was requested.

The **Start automatically** toggle under the LAN access card puts Unsloth on the network each time it starts. Stopping LAN access now won't turn this off, the toggle is a separate preference from the running state.

After installing Unsloth manually, run in your terminal:

Unsloth binds to every interface, so the raw port is reachable from your network. No public Cloudflare URL is published, `-H 0.0.0.0` alone keeps you on the LAN. The first line printed is the address other devices use:

The address may be IPv6, in square brackets as above. Copy the whole URL including the brackets.

The raw port stays reachable on your network *and* a public Cloudflare URL is published. The banner warns you about exactly this, because it is the least private mode. The remote access page covers it.

1. Open **Settings → API → Remote & LAN** .
2. Find the **LAN access** card.
3. Click **Start** .
4. When the status reads **Online** , open the shown address from any device on the same network.

`unsloth studio`

this machine only

`unsloth studio -H 0.0.0.0`

your network

`unsloth studio -H 0.0.0.0 --cloudflare`

your network, plus a public Cloudflare URL

`unsloth studio --secure`

a public Cloudflare URL only, the LAN stays closed

`--secure` and LAN access are opposites: `--secure` forces a loopback bind, so nothing on your network can reach the raw port.

Putting Unsloth on the network means every device on that network can see it.

- **Server-side tools run as your user.** Web search, Python and terminal execution are on by default, so anyone reaching the server with your API key can run code on that machine. Pass`--disable-tools` when exposing Unsloth, and keep the API key private. Every network-reachable launch says this in the banner.
- **Unsloth checks whether the port leaks past your router.** On a wildcard bind, a reachability probe tests whether the raw port is accidentally reachable from the internet, and the startup banner says so if it is. That probe contacts`ifconfig.me` and`check-host.net` ; set`UNSLOTH_STUDIO_DISABLE_PUBLIC_CHECK=1` to skip both.

**Error: Unsloth Studio is already running on port 8888 (PID ...)**
**.** Another instance holds the port. Run`unsloth studio stop` first, or start this one on a different`--port` .
- **The address does not answer from another device.** Both devices must be on the same network. Guest Wi-Fi, client isolation ("AP isolation"), and VPNs are the usual culprits. Also check the machine's own firewall allows inbound connections on the port.
- **The address changed.** LAN IPs are assigned by your router and can change. Give the machine a static IP or a DHCP reservation in your router settings for a stable address.
- **Start is greyed out in the UI.** Read the message under the card, it names the block reason.

`UNSLOTH_STUDIO_DISABLE_PUBLIC_CHECK`

`1` skips the third-party raw-port reachability probe on wildcard binds.

`UNSLOTH_STUDIO_TRUST_FORWARDED`

`1` honours `X-Forwarded-For` behind your own reverse proxy.

All four require a **UI session,** an API key is rejected. A refused operation returns `409` with the block reason as the detail.

`GET`

`/api/settings/lan-access`

Current state, address, `can_start` / `can_stop`, `block_reason`.

`POST`

`/api/settings/lan-access/start`

Schedule a settings-owned start. Idempotent.

`POST`

`/api/settings/lan-access/stop`

Schedule a settings-owned stop. Does not change the auto-start preference.

`PUT`

`/api/settings/lan-access/auto-start`

`{"enabled": true|false}`.

Last updated

Was this helpful?
