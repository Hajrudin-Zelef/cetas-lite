---
id: etape6-phased3-bgp-ha/00-bgp-ha/wave-2-route-policy-in-the-dc-fabric-prefix-lists-route-maps
title: "Wave 2 — Route policy in the DC fabric: prefix lists, route maps, communities, RPKI"
domain: phase-d3-bgp-underlay-and-high-availability-in-the-data-cent
role: deep-dive
task: regulation
actors: []
dates: ["2026-09-22"]
keywords: ["dci"]
source: docs/RAG/etape6_phaseD3_bgp_ha.md
source_anchor: ""
source_lines: [97, 144]
section: "Phase D3 — BGP underlay and high availability in the data center"
sha256: b0b445578fd718e1a4693dac03c1cbd58a7f885f14a479463eebd20bccc59223
---

# Wave 2 — Route policy in the DC fabric: prefix lists, route maps, communities, RPKI

## Wave 2 — Route policy in the DC fabric: prefix lists, route maps, communities, RPKI

### 2.1 Where policy is applied in a BGP DC fabric

- Policy enforcement points in a typical eBGP-underlay + iBGP-EVPN design [independent][official]:
  - **Underlay eBGP sessions:** inbound prefix-lists limiting accepted routes to expected loopback/peer-link ranges; outbound filters on `redistribute connected` via route-maps matching only loopback and P2P prefixes (Arista lab pattern: `RM-CONN-2-BGP` matching `PL-LOOPBACKS` and `PL-P2P-UNDERLAY`) [independent — Arista ATD EVPN lab guide].
  - **DCI / border leaves:** full route-maps with prefix-lists, AS-path filters, and community tagging toward WAN/transit [independent].
  - **EVPN overlay:** route-maps matching on EVPN route type or embedded IP prefix (Cisco IOS-XE documents `match ip address prefix-list` against the EVPN prefix field for Type-2/3/5 routes) [official — Cisco BGP EVPN VXLAN config guides, IOS-XE 17.12–17.15].
- Arista EOS offers **`peer-filter` with `match as-range`** for dynamic BGP peering (`bgp listen range`), letting operators accept only peers whose ASN falls in an expected range — useful for zero-touch leaf onboarding [independent — Arista ATD lab guide].
- `maximum-routes` (prefix limits) on fabric peer-groups (e.g., 12000 in the Arista lab example) guard against route leaks blowing up FIB/RIB [independent].

### 2.2 Prefix lists and route maps: DC-specific patterns

- Underlay prefix-list pattern: permit loopback /32s (e.g., `1.1.1.0/24 eq 32`) and P2P /31s (e.g., `10.0.0.0/8 le 31`); deny everything else — keeps the underlay RIB minimal and deterministic [independent].
- Route-map evaluation order matters: first-match wins; sequences evaluated in order; an implicit deny-all terminates non-matching routes. Cisco documents this explicitly for EVPN route-maps [official].
- Common set actions in DC fabrics: `set local-preference` (iBGP overlay path preference), `set community` (tagging for downstream policy), `set ip next-hop` (rare in underlay; used at DCI), MED for DCI exit selection [independent].

### 2.3 Communities: standard, extended, and large

- **Standard communities** (RFC 1997, `ASN:value`) are used for basic tagging (e.g., marking DCI-learned routes); **extended communities** (RFC 4360) carry EVPN route targets and encapsulation types — mandatory in the EVPN overlay [official].
- **Large communities** (RFC 8092, 12 bytes: `ASN:function:value`) are increasingly used in DC fabrics for structured policy: site/role/function encoding, and for automated remediation (e.g., blackhole communities) [secondary][independent].
- Operational notes [independent]:
  - `send-community` (and `send-community extended` / `large`) must be enabled per neighbor — a classic misconfiguration is tagging routes that are never sent because the send knob is off.
  - Community-based filtering at DCI borders is the standard defense against route leaks from tenants.
- Well-known blackhole community **65535:666** (RFC 7999) is an ISP convention; inside DC fabrics, operators more often use locally-defined large communities for discard policies [secondary][unverified for specific DC deployments].

### 2.4 RPKI and origin validation in DC contexts

- RPKI (RFC 6480 series; ROV per RFC 6483) is a **WAN/Internet** control; inside the private-ASN DC fabric, prefix ownership is established by design (private address space, closed ASN plan), so RPKI provides little value on fabric links themselves [independent].
- Where RPKI does touch the DC: **border leaves / DCI routers** peering with transit or cloud interconnects validate received Internet routes against ROAs — standard ISP practice extended to the DC edge [secondary].
- **Status as of 2026-09-22:** no evidence of RPKI/ROV being deployed *inside* DC underlays in production; vendor NOSes that support RPKI (IOS-XR, Junos, EOS) document it for Internet-edge use cases, not fabric underlays [official][secondary]. **Gap:** operator survey data on RPKI at DC borders is thin.
- Alternative in-fabric protections actually used: strict prefix-lists, max-prefix limits, AS-path filters, and GTSM/TTL-security on fabric eBGP sessions (RFC 5082) — the latter guards against off-link spoofed BGP packets [independent].

### 2.5 Wave-2 verification notes and open items

- **Verified:** Arista lab route-map/prefix-list/peer-filter patterns; Cisco EVPN route-map match semantics; community RFC numbers; RPKI's edge-only role.
- **Open:** production RPKI adoption at DC borders (survey); large-community conventions across operators.
- *End of Wave 2.*

---

## Wave 3 — Multi-chassis LAG: per-vendor architectures, control planes, and failure behavior

### 3.0 Why MC-LAG still matters in a BGP/EVPN world

- MC-LAG (dual-homing a server or downstream switch to two physical switches that present one logical LAG) remains the standard for **L2-attached hosts, appliances, and firewalls** that cannot do L3 ECMP; in EVPN fabrics it coexists with the anycast-VTEP/all-active multihoming model [independent].
- All implementations share the same building blocks: an **inter-switch link** (peer link/ISL/VLTi/ICL) carrying control traffic plus a backup data path, a **keepalive/heartbeat** on a separate path for split-brain detection, and **role election** (primary/secondary or active/standby) deciding who keeps member ports up when the inter-switch link fails [independent].

