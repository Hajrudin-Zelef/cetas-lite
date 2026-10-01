---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/vonrox-pia-wireguard-generator-blob-head-docs-wiki-features-unifi-automation-md-a1bc43f6-1
title: "vonrox-pia-wireguard-generator-blob-head-docs-wiki-features-unifi-automation-md-a1bc43f6"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: ["2026-09-02", "2026-09-13"]
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/vonrox-pia-wireguard-generator-blob-head-docs-wiki-features-unifi-automation-md-a1bc43f6.md
source_anchor: ""
source_lines: [1, 95]
sha256: 731f3ba4b5ce243055d6e93ca6eb38e84510c177a0a48ec17e8bca58dc9c8fa9
---

# vonrox-pia-wireguard-generator-blob-head-docs-wiki-features-unifi-automation-md-a1bc43f6

| title | UniFi Automation | 
|---|---|
| aliases |  | 
| tags |  | 
| created | 2026-09-02 | 
| updated | 2026-09-13 | 
| sources |  | 
| status | active | 
| confidence | medium | 
| UniFi Sync | Automation | Headless | 
| features | automation | unifi | 
scripts/pia-unifi-sync.mjs does the desktop app's round trip without the desktop: sign in to
Private Internet Access, generate a key pair, register it with a server in a chosen region, and
write the result straight into a WireGuard VPN Client on a UniFi gateway (UCG Ultra, UDM, UDR,
or a self-hosted Network application). Run it from a timer and the tunnels stop needing you.
A PIA WireGuard registration is not a credential you hold; it is a row in a server's peer table
that exists only while the peer keeps talking. PIA documents none of this. Their own
manual-connections scripts document the auth
token's expiry and the port-forward signature's two months, and say nothing at all about how long
a key registration lives.
What is known comes from community operators, and is worth quoting accurately rather than in the form it usually gets repeated:
- Idle expiry. thrnz/docker-wireguard-pia ,
the most widely used PIA WireGuard container, reports that keys "seem to expire at PIA's end
after several hours of inactivity" and that settingPersistentKeepalive "may be enough to
prevent this from happening". Note the hedging — this is observed behaviour, not a specification.
- A second key displacing the first. This one is usually repeated as a flat rule, and it is not. The source scopes it to a three-way coincidence: a second key registered to the same endpoint address using the same auth token, which "in practice is probably only an issue if multiple containers share the same auth token, and if they also happen to pick the same vpn endpoint". Two tunnels in different regions do not meet that condition.
So a gateway reboot, a firmware update, or a WAN outage that outlasts the keepalive leaves the console showing Not Established until someone generates a new configuration and pastes it in. The account token itself lasts about a day, which is why the generator has to sign in again rather than reuse one.
Since the documented cause is inactivity, the cheapest fix is to stop the tunnel going idle.
If your VPN Client was created by uploading a .conf, check that it carries
PersistentKeepalive = 25 under [Peer]. That costs nothing, needs no script, and every refresh
it avoids is a refresh that cannot go wrong.
It may not be available: the networkconf row a console returns has no keepalive field of its
own, so on a tunnel configured by hand in the UI there may be nothing to set. --diagnose lists
the fields your console actually returns, which will tell you either way.
Worth stating plainly before anything else. The script drives
/proxy/network/api/s/<site>/rest/networkconf — the console's private API, the one its own
web UI calls. Ubiquiti publishes no schema for it, does not version it, and promises nothing
about it. The maintainers of the longest-running client for it
(Art-of-WiFi/UniFi-API-client) put a
disclaimer at the top of their README saying exactly that.
This is not a shortcut around a supported path. There isn't one:
- The official UniFi Network Integration API (/proxy/network/integration/v1 ) cannot do it.
From Network 10.0 it gained real CRUD for networks, firewall policies, DNS and ACLs — but its
"network" object is a VLAN: name,vlanId , subnet, DHCP. Nopurpose , novpn_type , no
WireGuard field. Its only VPN surface is read-only, and lists VPN servers, not clients. On
Network 9.x it cannot touch network configuration at all.
- The Site Manager API (api.ui.com ) is read-only — inventory, sites, ISP metrics.
- There is no official CLI, config file, Ansible collection or Terraform provider.
- Every community tool that can configure WireGuard on a UniFi gateway — the Terraform providers, the PHP clients — uses this same private endpoint. The providers built on the official API are structurally incapable of it.
What this means in practice. The path itself has been stable for years. What breaks is
payload validation: a firmware update changes a field's name or type, and writes start failing
with 400 api.err.*. That is a loud failure, not a silent one, and this tool is built to stop
rather than guess — it refuses to write a row missing the fields it knows, and prints what it
found instead. If a Network update ever breaks it, expect a clear refusal and an unchanged
gateway, not a mangled tunnel.
A VPN Client is a networkconf row with purpose: "vpn-client" and vpn_type: "wireguard-client". The script reads the row by the name shown in the console, and replaces
exactly the fields a registration changes:
| Field | New value | 
|---|---|
| x_wireguard_private_key | the freshly generated private key | 
| wireguard_public_key (when present) | its public half | 
| wireguard_client_peer_public_key | the server key addKey returned | 
| wireguard_client_peer_ip ,wireguard_client_peer_port | the endpoint | 
| ip_subnet | the peer address, /32 | 
| wireguard_client_configuration_file (rows created by upload) | a rendered .conf matching the above | 
Routing, DNS pulling, which networks use the tunnel, and every other setting are sent back
unchanged. The row is then PUT to rest/networkconf/<id>, which is what the console's own UI
does on Apply, and the gateway re-provisions the tunnel — expect a few seconds of interruption
per tunnel, so schedule the run for a quiet hour.
If a row lacks the fields above the script refuses to write it and prints the fields it did find; Ubiquiti publishes no schema for this API, and guessing at one is how a tunnel ends up silently misconfigured.
It also refuses when the console returns a masked secret — a run of asterisks where a stored
key belongs. Some builds do that so reading a row does not disclose it, which is good practice and
a trap for anything that reads a row and writes it back: send the mask and the console stores the
mask. For a preshared key that means the tunnel stops handshaking and nothing in the reply says
so. The same bug was reported against a Terraform provider as
ubiquiti-community/terraform-provider-unifi#490.
The private key is exempt, because a refresh replaces it outright.
- 
Create the VPN Clients once in the console (Settings → VPN → VPN Client → WireGuard), by uploading a configuration from the app. Give each a name you will refer to.
- 
Choose how the script signs in. Either: 
  - an API key — Settings → Control Plane → Integrations on UniFi Network 9 and later, set as
UNIFI_API_KEY ; or
  - a dedicated local admin without multi-factor authentication, set as UNIFI_USERNAME andUNIFI_PASSWORD . A UniFi account with MFA cannot be used unattended.
 The API key is preferred: it can be revoked on its own, and a leak does not expose a password. Whether a key is accepted on the networkconf route depends on the Network version.--dry-run does not answer that — it registers keys with PIA and then stops, without ever issuing the
write (syncTunnels skipsupdateNetwork entirely whendryRun is set). Use--diagnose --probe-write , which writes one row back unchanged and reports whether the console
accepted it.
- an API key — Settings → Control Plane → Integrations on UniFi Network 9 and later, set as
- 
Trust the console's certificate. A console ships with a self-signed certificate, and the script never disables verification. Export the certificate from your browser (the padlock → certificate → export as Base64/PEM) and point unifi.certificate at the file. The script then
accepts that certificate and no other, regardless of the name it carries, which is stricter
than what a browser does after you click through the warning. If your console has a real
certificate for the name you reach it by, omit the setting.
- 
Write the configuration — copy examples/pia-unifi-sync.example.json :{
