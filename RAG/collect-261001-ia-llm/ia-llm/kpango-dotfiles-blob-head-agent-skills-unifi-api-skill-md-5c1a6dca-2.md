---
id: collect-261001-ia-llm/ia-llm/kpango-dotfiles-blob-head-agent-skills-unifi-api-skill-md-5c1a6dca-2
title: "Restart a device"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-ia-llm/kpango-dotfiles-blob-head-agent-skills-unifi-api-skill-md-5c1a6dca.md
source_anchor: ""
source_lines: [143, 238]
sha256: 36c48229de230b7eb7b69499b967c1c3eedf0ef0e8c2dfbdcc9c2002d679b167
---

# Restart a device

- Classic API cookie sessions expire; re-login on 401.
- Cloud Connector requires consoleId fromunifi.ui.com — find it in the console URL.
- Pagination default limit varies by endpoint; always check totalCount vscount .
The console (UDM/UDM Pro/etc.) is commonly reached over Tailscale SSH when set up that way; a
Tailscale IP is dynamic — resolve it fresh via tailscale status rather than hardcoding a stale
one across sessions. Managed APs/switches often have no SSH key provisioned (password auth only)
— if your ~/.ssh/config has host aliases for them, prefer those over raw IPs.
Credentials live in pass (exact entry names are per-deployment — check project memory for
which ones apply to your device):
- one entry typically holds the API key (see Authentication above)
- another typically holds the mgmt-account password, used for both SSH to the console's
root account (often key-based there already) and password-auth SSH to AP/switch aliases
Never interpolate pass show ... directly into a remote SSH command string (ssh host "$(pass show x)...") — this can trip the harness's command-safety classifier and leaks the secret into
shell history/process args. For password-auth SSH to APs/switches, use a local SSH_ASKPASS
script instead of sshpass (often absent) or expect — use mktemp (not a fixed path) and set
permissions explicitly before anything can read it, to avoid a symlink-preplant/TOCTOU race on a
predictable /tmp filename:
ASKPASS=$(mktemp)
trap 'rm -f "$ASKPASS"' EXIT INT TERM
chmod 700 "$ASKPASS"
cat > "$ASKPASS" << 'EOF'
#!/bin/bash
pass show <your-mgmt-password-entry>
EOF
SSH_ASKPASS="$ASKPASS" SSH_ASKPASS_REQUIRE=force \
  ssh -o PreferredAuthentications=password -o PubkeyAuthentication=no <ap-alias> 'uptime' < /dev/null
UniFi OS runs on a Debian (trixie) base with /etc/apt/sources.list.d/ubiquiti.list pointing at
Ubiquiti's own repo alongside the standard Debian repos. Ubiquiti ships custom builds of
several packages (iptables/libxtables12/libip4tc2/libip6tc2 with vendor iptables
extensions dpi128/dyn_random/geoip/ubnt_mark; frr for zebra/OSPF; dnsmasq, openvpn,
ppp, keepalived, lldpd, miniupnpd, wireguard-tools, xl2tpd,
xtables-addons-common, iproute2, and more) with higher version numbers stamped on the
Debian side. A plain apt upgrade/apt dist-upgrade — including one triggered incidentally
while installing an unrelated package — silently pulls the generic Debian build over the
Ubiquiti one wherever Debian's version compares higher, with two distinct failure modes observed
in production:
- ABI split, not a full downgrade: the iptables /iptables-restore binary andlibxtables.so.12 core library get replaced (new ABI) while individual extension.so files
under/usr/lib/aarch64-linux-gnu/xtables/ are left as the old vendor build (or vice versa,
depending on which packages actually had matching Debian candidates) — the resulting binary
vs. extension-module ABI mismatch causesiptables-restore to SIGSEGV on any ruleset that
uses the affected extensions (TCPMSS, GeoIP), not merely fail cleanly. This is the direct
mechanism behind a persistent "GatewayConfigurationError" (config apply crashes mid-write
instead of committing) and silently-truncated GeoIP country lists (the crash cuts the rule
list short partway through, which looks like a string-length limit but isn't one).
- Full package replacement: frr /libyang2 etc. get fully replaced, removing binaries
(/usr/sbin/zebra ) that udapi-server's config-apply pipeline depends on (OSPF), which trips
the same GatewayConfigurationError rollback for unrelated reasons.
Diagnose drift: dpkg -V <pkg> (empty/??5?????? output = files changed since install;
compare against /var/lib/dpkg/info/<pkg>.md5sums), grep -B2 -A2 <pkg> /var/log/apt/history.log
to date the replacement, apt-cache policy <pkg> to see whether the Ubiquiti-repo candidate is
still resolvable at all (it may not be, if Ubiquiti's own repo no longer serves an old version —
don't assume apt-get install --reinstall alone can restore the vendor build).
Check for an already-maintained fix on the device before building a new one — a prior
remediation may already exist as a small framework somewhere under /data/ or
/etc/apt/preferences.d/ that (a) pins the affected package set toward the vendor repo, (b)
preflight-checks dist-upgrade for would-be replacements before they happen, and (c) can restore
drifted files from the firmware's read-only lower layer (/mnt/.rofs/...), repackaging them for
dpkg. Check project memory for this device's specific path/tool name and run its own verify
step first before assuming a fresh repair is needed. Only recreate such a framework from scratch
if it's genuinely absent — and never duplicate it with an ad-hoc per-package Pin-Priority: -10
file, which starves apt of any candidate for that package (including the vendor one) rather
than steering it toward the vendor repo.
Never run a blanket iptables -F/ip6tables -F (any table) on this system to "reset" a
problem — NAT, fwmark-based policy routing, and Tailscale's own chains are all wired through the
same tables udapi-server manages, and flushing them independently of udapi-server's own state can
sever LAN routing/NAT until udapi-server rebuilds them (observed to cause a real outage). Before
touching a live ruleset, verify hypotheses non-destructively:
# Confirm a specific table's ruleset applies cleanly without touching the live kernel state
CHECK=$(mktemp)
iptables-save -t mangle > "$CHECK"
iptables-restore -w --counters --table mangle --noflush --test "$CHECK"
echo "exit: $?"   # 0 = would apply cleanly; 139 = SIGSEGV (see Package Safety above); 2 = parse error
rm -f "$CHECK"
Diagnosing a stuck inform/API TCP handshake (device shows stale last_seen in the
controller but ICMP/ARP to it work): check conntrack -L | grep dport=8080 for entries stuck in
SYN_RECV — the controller's reply direction shows several retransmitted SYN-ACK packets that
never get ACKed back, indicating an L2-adjacent (not routing) hangup rather than a controller
software crash. Cross-check by fetching the same device's networkconf/device documents directly
via mongo (below) to rule out a config-level cause (e.g. a WAN network object with the enabled
field entirely absent — see "Direct MongoDB queries" below) before assuming it's purely transport-layer.
Direct MongoDB queries (when the local controller API itself is unresponsive/crashing —
unifi-core's reverse proxy to the Java network app can wedge independently of udapi-server,
returning HTTP:000/502/503 on /proxy/network/* while still answering 401 for
unauthenticated probes; a systemctl restart unifi-core.service often clears this, but get
explicit confirmation first — it has been observed to cascade into UniFi's own self-healing full
system reboot when the device's underlying state was already unhealthy, not the restart alone):
mongo --quiet --port 27117 --eval 'db.device.find({},{name:1,last_seen:1,_id:0}).toArray()' ace
mongo --quiet --port 27117 --eval 'db.networkconf.find({purpose:"wan"},{name:1,enabled:1,wan_type:1,_id:0}).toArray()' ace
The ace database backs the classic controller; device/networkconf/power_supervisor are
the collections most relevant to connectivity troubleshooting. A WAN networkconf document
missing its enabled field entirely (not false — absent) has been the root cause of a
persistent internal mcad error flood (wan_man_get_primary(): UDAPI required field 'status' is missing, tens of thousands of log lines/hour) even while the interface itself forwards traffic
normally — fix via PUT /proxy/network/api/s/default/rest/networkconf/{id} (Classic REST API)
with the full existing document plus the single missing field, never a partial-document PUT.
On a DS-Lite WAN (wan_type: dslite), UniFi's own GeoIP enforcement only attaches to the
IPv4-in-IPv6 tunnel interface (ip6tnl1) — observed in production: native IPv6 traffic on the
