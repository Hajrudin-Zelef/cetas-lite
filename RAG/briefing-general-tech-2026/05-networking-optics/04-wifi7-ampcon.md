---
id: briefing-general-tech-2026/05-networking-optics/04-wifi7-ampcon
title: "Wi-Fi 7 campus and AmpCon"
domain: networking-optics
role: deep-dive
task: networking
actors: ["FS.com"]
dates: ["2026-09", "2026-09-21"]
keywords: ["dci"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g06-4"
source_lines: [5837, 5890]
sha256: 634697bd6614e2d80af8ec351905bcb6412fcf5bb61759b3b2143f6a146bf9db
---

# Wi-Fi 7 campus and AmpCon

<a id="g06-4"></a>
### 6.4 Wi-Fi 7 campus and AmpCon

On September 21, 2026 (via BusinessWire), FS.com confirmed a high-density
campus Wi-Fi 7 offering anchored by the **AP-N755** access point, rated at
**24.436 Gb/s** aggregate throughput. High-density campus — lecture halls,
stadiums, dense office floors — is the use case Wi-Fi 7 was designed for: the
aggregate figure comes from stacking multiple bands (2.4, 5 and 6 GHz) with
4096-QAM, 320 MHz channels and multi-link operation, serving hundreds of
clients per cell rather than pushing single-client peak rates.

The same announcement touched on FS's network-management software line:
**AmpCon-Campus** and **AmpCon-T** are confirmed products, while
**AmpCon-Campus 3.0** was described as an "upcoming release" — planned, not
shipped. The distinction matters for the reader: anything about 3.0 beyond the
announcement date is announced/planned, not delivered.

The Wi-Fi 7 move rounds out the picture of FS as a full-stack enterprise
vendor in 2026: from 1.6T AI fabrics down through DCI to the campus edge. It
also shows how the September 2026 news cycle concentrated vendor announcements
in the week around ECOC, even for products with no direct relation to the
optical-transport trade show.

#### The AP-N755 and the Wi-Fi 7 generation

The 24.436 Gb/s figure on the AP-N755 is an aggregate across radios, not a
single-client speed — the standard way Wi-Fi access points are specified,
and the number that matters for the high-density use case. Wi-Fi 7 (IEEE
802.11be) brings the access point three structural upgrades over Wi-Fi 6:
320 MHz channels (doubling the widest Wi-Fi 6 channel), 4096-QAM (denser
constellation, ~20% more bits per symbol), and multi-link operation (a
client using 2.4, 5 and 6 GHz simultaneously). In a lecture hall or stadium,
where hundreds of clients share one cell, the aggregate is the honest
metric: no single phone will see 24 Gb/s, but the cell can keep hundreds of
them fed. FS's announcement is one data point in the 2026 enterprise Wi-Fi
refresh cycle, in which every major campus vendor shipped a Wi-Fi 7
flagship — the differentiation, as always, is in radio resource management
software and price, not in the silicon, which comes from the same handful
of chip vendors.

#### AmpCon: the management layer, and the 3.0 that is not here yet

Network hardware without management software is unsellable to enterprise
IT, which is why the AmpCon line matters to the FS story. **AmpCon-Campus**
(the campus/wireless controller lineage) and **AmpCon-T** are confirmed,
shipping products. **AmpCon-Campus 3.0**, described in the September 21
release as an "upcoming release," is planned — and the dossier's rule for
planned software is strict: features, dates and capabilities attributed to
3.0 are announced, not delivered, until a release notice says otherwise.
The reader tracking FS's enterprise stack should therefore file 3.0 under
"watch for the release notes," alongside the usual caution that ".0"
management-platform releases in this industry historically arrive with a
long tail of point fixes.

