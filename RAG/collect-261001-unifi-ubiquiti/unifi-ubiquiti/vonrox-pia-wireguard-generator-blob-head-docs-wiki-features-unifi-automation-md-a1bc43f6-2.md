---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/vonrox-pia-wireguard-generator-blob-head-docs-wiki-features-unifi-automation-md-a1bc43f6-2
title: "vonrox-pia-wireguard-generator-blob-head-docs-wiki-features-unifi-automation-md-a1bc43f6"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/vonrox-pia-wireguard-generator-blob-head-docs-wiki-features-unifi-automation-md-a1bc43f6.md
source_anchor: ""
source_lines: [96, 176]
sha256: 9930c00a7fd8b7f37ca0b069b07e8a7c4a44f47c1ece76a0e5e35dc13c3e05fd
---

# vonrox-pia-wireguard-generator-blob-head-docs-wiki-features-unifi-automation-md-a1bc43f6

  "unifi": { "url": "https://192.168.1.1", "certificate": "unifi-console.pem" },
  "dns": "10.0.0.243",
  "tunnels": [
    { "network": "WireGuard PIA CZ", "region": "czech" },
    { "network": "WireGuard US East", "region": "us_east" }
  ]
}network is the VPN Client name exactly as the console shows it;region is a PIA region id
(--list-regions prints them).dns takes the same values as the app, and a tunnel may override
it. The file holds no secrets. Add"site" if the site was renamed, and"selfHosted": true for a
Network application that is not on a UniFi OS console.
- 
Put the credentials in the environment, or in files named by PIA_USERNAME_FILE and
friends — the form Docker secrets and systemdLoadCredential= use.
- 
Rehearse: node scripts/pia-unifi-sync.mjs --config pia-unifi-sync.json --list-networks node scripts/pia-unifi-sync.mjs --config pia-unifi-sync.json --dry-run A dry run signs in everywhere, registers a real key with PIA, and prints the fields it would change — private keys named but never shown — without writing to the console.
- 
Schedule it. examples/systemd/ has a service and a timer for twice a day; the equivalent
cron line is0 4,16 * * *  cd /opt/PIA-WireGuard-Generator && node scripts/pia-unifi-sync.mjs --config /etc/pia-unifi-sync.json --quiet
Twice daily is a comfortable margin against both the token's day and the registration's idle timeout. Every run rotates the key pair, which is what the app does too.
Anywhere with Node 20+ and curl that can reach both the internet and the console: a NAS, a Raspberry Pi, the machine that already runs Home Assistant. It cannot run on the gateway, which has no Node, and a script placed on UniFi OS does not survive a firmware update in any case.
With no always-on host, a timer is beside the point — you want to fix the tunnels when you notice
they are down. scripts\pia-unifi-sync.cmd does that:
- Put pia-unifi-sync.json andpia-unifi-sync.env in the repository root (copy them fromexamples\ ). Both are already in.gitignore .
- Double-click scripts\pia-unifi-sync.cmd . With no arguments it performs a dry run — it
registers fresh keys with PIA and prints what it would change, without writing to the console.
- When the output looks right, run it again with --apply .
Credentials travel in the environment, never on the command line, so they do not appear in the window title, the scroll buffer, or another user's process list. The launcher refuses to start with a clear message if Node is missing or either file is absent.
scripts\pia-unifi-sync.ps1 does the same job with the credentials encrypted by DPAPI for your
Windows account instead of sitting in pia-unifi-sync.env:
powershell -File scripts\pia-unifi-sync.ps1 -SetCredentials
powershell -File scripts\pia-unifi-sync.ps1 --list-networks
powershell -File scripts\pia-unifi-sync.ps1 --apply
-SetCredentials prompts for the PIA username, PIA password and UniFi API key (nothing typed is
echoed) and stores them in %LOCALAPPDATA%\pia-unifi-sync\credentials.xml, in a folder only your
account can open. Arguments behave exactly as they do for the .cmd: none is a dry run, --apply
writes, anything else is passed through. -ForgetCredentials deletes the store.
What it buys over the .env: nothing in the repository to commit by accident or lose with a
checkout; a copied file, a backup, another account or a pulled disk cannot decrypt it; and a
password containing " works. What it does not buy: a program already running as you can decrypt
it, just as it could read a file you own. The API key is a site-admin credential either way —
revoke it in the console when you are done with it.
--diagnose reports the three facts that decide whether this can work, and how:
node scripts/pia-unifi-sync.mjs --config pia-unifi-sync.json --diagnose
It reads the console's certificate — whether it is self-signed, whether it is a CA or a leaf, and what names it carries — then performs the same read the sync performs and reports what came back that a curl-based client would not be able to see: response header names, whether a session cookie was set, and whether a CSRF token was picked up. Header names only; no value is ever printed.
It then describes each VPN Client the configuration names, field by field — which is the only
way to learn what your particular console puts in a row, since no schema is published. Field names
are always listed; values only for the few short flags that decide how a row must be written
(wireguard_client_mode, whether a preshared key is enabled, DNS pulling, default route). Every
secret — the WireGuard keys, the configuration file, and any x_ field — is reported as
present, absent, blank, looks redacted or not a key, never shown. A tunnel name that
matches a network which is not a WireGuard VPN Client is reported, not described.
Add --probe-write to also write each configured row back unchanged. That is the only honest
way to learn whether your credential authorises a write without changing anything: the body is
byte-identical to what the console just sent, so the gateway re-provisions the tunnel briefly but
its configuration does not change.
The first write comes from a client that has sent nothing — its PUT is its first request, so it carries no cookie or CSRF token. A write made after the read would replay whatever the read was handed, and so could not tell a cookie the console sets from one it requires. Only if that bare write is refused on authorisation, and the read did leave a session behind, is the write repeated with it. The report lists both, and the verdict follows the difference: a cookie that is set but not needed does not block the desktop app; a write accepted only with the cookie does. With a password sign-in there is no bare client to try, and the report says the question stays open.
Unless the row is not safe to round-trip — then "byte-identical" would be a lie. A private key
that is not a real key, a preshared key that is masked (or missing while in use), any x_ field
that came back masked or blank, or a row missing the fields the sync needs: writing it back could
store a mask or a blank over the real value. The probe checks first and refuses, reporting which
field. That refusal is itself the finding: it means no tool can safely round-trip that row.
scripts/pia-unifi-sync.mjs      argument parsing, credentials, exit codes
scripts/unifi-sync/sync.mjs     reads credentials from the environment; re-exports the rest
resources/js/core/unifi-sync.js config validation, row patching, the per-tunnel loop
scripts/unifi-sync/unifi.mjs    the console client: sign-in, CSRF, networkconf GET/PUT, certificate pinning
scripts/unifi-sync/exec.mjs     runs `curl -q --config -` through execFile — no shell at all
scripts/unifi-sync/crypto.mjs   X25519 from node:crypto, the reference the test suite already trusts
The sync itself lives in the application's tree rather than beside the script,
because none of it is specific to a command line: it takes a parsed
configuration and two injected clients and returns a report. It is built from
three steps the desktop app can use separately — inspectTunnels matches each
configured tunnel to its row and costs nothing, prepareTunnel registers one
key with PIA, and applyTunnel writes one row back. syncTunnels is those
three composed, and it registers and writes one tunnel at a time rather than
batching every registration ahead of every write: a key PIA has issued does
nothing until the gateway is using it, so the gap between the two is kept
small. test/unifi-sync.test.js asserts that ordering.
The PIA side is the app's own HttpClient and PiaClient, so PIA requests are pinned to the
bundled CA and addKey replies are validated before anything is written, exactly as in the app.
The UniFi side uses node:https directly, because it may need to read response headers — a
console that sets a session cookie on an API-key request has to be followed, and the curl config
