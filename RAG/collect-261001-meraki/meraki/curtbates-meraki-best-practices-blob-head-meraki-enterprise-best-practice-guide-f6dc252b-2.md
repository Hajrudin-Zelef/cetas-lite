---
id: collect-261001-meraki/meraki/curtbates-meraki-best-practices-blob-head-meraki-enterprise-best-practice-guide-f6dc252b-2
title: "curtbates-meraki-best-practices-blob-head-meraki-enterprise-best-practice-guide--f6dc252b"
domain: meraki
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["datacenter", "latency", "omni", "voice"]
source: docs/RAG/collect-261001-meraki/curtbates-meraki-best-practices-blob-head-meraki-enterprise-best-practice-guide--f6dc252b.md
source_anchor: ""
source_lines: [60, 115]
sha256: f9a0c676871014e0ef148e0f09dddb75c3457444f83f1e8759143e43296da85c
---

# curtbates-meraki-best-practices-blob-head-meraki-enterprise-best-practice-guide--f6dc252b

- Perform an active site survey + spectrum analysis (e.g., Ekahau/AirMagnet). Target ≥ 25 dB SNR across the coverage area, on both bands (don't survey 2.4 GHz only).
- Mounting heights:
  - Below ~8 ft (3 m): integrated/omni dipole antennas.
  - 8–25 ft (3–8 m): external downtilt omni antennas.
  - 25 ft+ or hard ceilings: wall-mount (10–15 ft, LED facing down) or use directional antennas.
- Use directional antennas above ~26 ft, outdoors, or where directional coverage is needed; point ceiling-mounted directional antennas straight down, tilt wall-mounted ones toward the floor.
- Keep SSIDs to ≤ 3 (a hard requirement in high density; ≤ 5 absolute max). Each extra SSID burns airtime on management frames — beyond 5 SSIDs you can lose 20%+ of capacity. Make one SSID per auth type (Splash / PSK / EAP) and consolidate same-auth SSIDs.
- Bridge mode is recommended for seamless L2 roaming (best for voice). All APs in the roaming domain (floor/RF profile) must share the same VLAN.
- Do not use NAT mode for VoIP/VPN/streaming — clients re-DHCP on every roam, breaking real-time sessions.
- For roaming across subnets at scale, the current Meraki solution is the Campus Gateway (VXLAN-tunneled SSID traffic; up to 5,000 APs / 50,000 clients per gateway). Distributed L3 roaming and "MX as concentrator" are no longer recommended for large-scale roaming.
| Setting | High-density recommendation | 
|---|---|
| Band selection | Dual-band with band steering if 2.4 GHz is needed; 5 GHz only if it isn't. | 
| Channel width | 20 MHz (VHT20) — preserves the number of non-overlapping channels and gives mixed-capability clients fair airtime. | 
| Minimum bitrate | Disable low legacy rates. Use ≥ 12 Mbps to exclude 802.11b and lift RF efficiency (Cisco SF uses 18 Mbps). Use 11 Mbps only if legacy 11b must be supported. | 
| DFS channels | Enable to reach up to 19 channels (US), improving channel reuse and lowering co-channel interference. | 
| Auto power | Auto RF tunes TX power every ~20 min, aiming for ≥3 neighbors heard. Constrain min/max TX power per RF profile in complex sites. | 
| RX-SOP | Tighten receive sensitivity to shrink cells in dense areas (e.g., 5 GHz −76/−78/−80 dBm high/med/low). Test first — too aggressive creates coverage holes. | 
| Client balancing | Enable in RF profiles for dense areas (off by default). | 
Fast roaming (mostly on by default): 802.11k (neighbor reports), 802.11i / PMK caching and OKC (faster 802.1X re-auth), and 802.11r / Fast BSS Transition — enable 802.11r explicitly under Access Control → Security for voice-heavy environments where clients support it.
- Add a shaping rule to give voice & video priority and set it to ignore the SSID per-client limit so real-time media isn't throttled; block/throttle P2P and recreational apps.
- MR APs automatically do multicast-to-unicast conversion (sends at negotiated rates — great for classroom video) and limit duplicate broadcasts to prevent storms and save client battery.
- Set each uplink's configured bandwidth limit to match the ISP's actual provisioned rate so the MX can shape correctly and avoid saturating the circuit.
- Configure multiple uplink-statistic test IPs (the default 8.8.8.8 , plus your ISP gateway and key remote VPN peers) for better monitoring and troubleshooting.
- Set security-list update frequency to hourly (AMP/IDS/URL lists) per uplink, including cellular.
- Load-balance across uplinks when you have similar-bandwidth redundant circuits.
- Use flow preferences to pin traffic to the right uplink — e.g., push guest traffic onto the cheaper secondary WAN to protect the primary.
- With multi-uplink Auto VPN, the MX builds tunnels on every reachable WAN interface and can dynamically move a flow to a healthier uplink when loss/latency/jitter degrade on the active path.
- Define performance/policy rules per application class so each flow takes the appropriate path; build these to match your traffic and circuit characteristics.
- Caveat: on MX 18.2 multi-WAN, SD-WAN and load-balancing policies do not apply to WAN3.
- In large client populations, set a global per-client bandwidth limit to prevent uplink saturation.
- Speedburst lets a client briefly exceed its cap (good for occasional large transfers) — enable it only where a few users need bursts, not where many need sustained high bandwidth simultaneously.
- Enable the default traffic-shaping rules — they already prioritize voice, software updates, and collaboration apps well for most deployments.
- Use Auto VPN for inter-site connectivity (Meraki-to-Meraki, same organization only).
- Use hub-and-spoke, not full mesh: a hub builds tunnels to all other hubs and to its configured spokes; a spoke builds tunnels only to its hubs. Making every MX a spoke (or everything a hub) degrades service at scale.
- Designate resource-rich sites (HQ, datacenter) as hubs. Auto VPN automatically builds redundant tunnels on all WAN interfaces that can reach the Meraki cloud, giving failover with no extra config.
- Deploy Client VPN with a Systems Manager (SM) policy pushed to endpoints — it removes manual user setup, improves the experience, and lets admins troubleshoot without walking users through config.
- Consider split tunneling to keep only corporate-bound traffic in the tunnel.
- Routed (NAT) mode when the MX terminates the ISP handoff directly — gives full L3 (NAT, routing, DHCP, multiple VLANs, the complete firewall/threat feature set).
- Passthrough / VPN-concentrator mode when an existing upstream L3 device handles routing — the MX behaves as an L2 device (typical for a datacenter Auto VPN concentrator). Add static routes on upstream L3 devices so VPN-destined traffic reaches the MX. Avoid running content filtering in this mode.
- Avoid stacking the MX behind another NAT device — it complicates cloud connectivity and can break Auto VPN.
- The MX is a stateful firewall: inbound is denied unless the flow originated inside or a forwarding rule exists.
- By default all VLANs can talk to each other — you must add L3 rules to stop it. Explicitly block guest VLANs from reaching business VLANs, and block specific destination IPs/URLs per VLAN as needed.
- L3 rules are processed top-down, first-match wins — order matters.
- L7 rules: be granular — broad category blocks (e.g., all "file sharing") can break legitimate apps like OneDrive. Block by country only when you're sure the traffic is malicious; over-blocking geographies breaks apps that pull resources from those regions.
- Scope inbound rules as narrowly as possible — only the required ports.
- Never use "Any" for allowed remote IPs except for a true public web server; even then, prefer an obscure public port rather than common web ports.
- Use 1:1 NAT to map a dedicated public IP to an internal host (keeps it off the MX WAN IP); use 1:Many NAT to multiplex services behind one public IP when addresses are scarce. For non-web services, restrict port and remote-IP ranges.
- AMP: enable — inspects HTTP downloads and blocks known-malicious files via the AMP cloud.
- IDS/IPS: enable — SNORT®-based detection/prevention on traffic through the MX; set to Prevention to actually block, not just alert.
- Keep security/signature lists updating hourly.
- Set IP Source Address Spoofing Protection to "Block" (not just "Log") so spoofed traffic is dropped at detection rather than merely recorded.
The most resilient topology combines redundancy at every layer:
| Layer | HA mechanism | 
|---|---|
| MX (edge) | Warm-spare HA pair (VRRP) + dual-ISP uplinks with automatic failover. Run HA management through a downstream L2 switch rather than a dedicated HA cable (the dedicated cable plus a shared L2 switch invites an STP loop). Allow all VLANs on the switch links so VRRP heartbeats reach across every VLAN. | 
