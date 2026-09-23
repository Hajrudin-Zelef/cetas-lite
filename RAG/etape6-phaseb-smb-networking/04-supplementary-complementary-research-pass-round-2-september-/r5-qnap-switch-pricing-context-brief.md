---
id: etape6-phaseb-smb-networking/04-supplementary-complementary-research-pass-round-2-september-/r5-qnap-switch-pricing-context-brief
title: "R5. QNAP switch pricing context (brief)"
domain: supplementary-complementary-research-pass-round-2-september-
role: deep-dive
task: pricing
actors: ["EU", "United States"]
dates: ["2026-05"]
keywords: ["pricing", "research"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [639, 701]
section: "Supplementary / Complementary Research Pass — Round 2 — September 22, 2026"
sha256: 04cbd790cba3ce51af988f3e327e7d75e9a743999dc899a04348609ed259ea2b
---

# R5. QNAP switch pricing context (brief)

## R5. QNAP switch pricing context (brief)

- **QSW-1208-8C** (12-port unmanaged 10GbE: 4× SFP+ + 8× combo SFP+/10GBASE-T): list **$668.00** [secondary]; historical street ~$349.99 (2019 sale) [secondary]. No 2026 price update found — still sold as special-order [independent/secondary].
  Sources: https://www.unisolinternational.com/product/12-port-unmanaged-10gbe-switch/ ; https://forums.servethehome.com/index.php?threads/qnap-qsw-1208-8c-12-port-unmanaged-10gbe-switch-350.26885/
- **QSW-M1208-8C** (L2 managed version): refurbished **$959.37** [secondary].
  Source: https://techatlantix.com/qsw-m1208-8c.html

## R6. UniFi Talk 2026 — software & hardware updates (new detail)

- **UniFi Talk 5.0** (~May 2026, per "UniFi Talk Just Leveled Up" video, crawled ~May 2026): redesigned phone interface, improved audio, **Smart Attendant with Call Queues**, single-tap camera/door control from phones; "bigger breakthroughs coming this year" [secondary]. Some features in 5.2 RC [secondary].
  Source: https://www.youtube.com/watch?v=T0S2ndgoVnQ
- **UT-ATA (Talk ATA)** — new 2-line analog telephone adapter (first-look review ~Mar 2026): T.38 fax passthrough, USB-C power, desk/wall mounts included, UniFi Talk adoption [secondary].
  Source: https://www.hostifi.com/blog/first-look-at-the-new-unifi-talk-ata
- **Official UniFi Talk softphone: upcoming, not launched** — YouTube creator summary of a UniFi livestream notes Ubiquiti "plans to release the UniFi Talk softphone" [unverified]. An unofficial native Android SIP softphone for UniFi Talk exists on GitHub (josh2893/unifi-talk-phone-android) [independent].
  Sources: https://www.youtube.com/watch?v=T0S2ndgoVnQ (comment); https://github.com/josh2893/unifi-talk-phone-android
- **G3 Touch Pro (UTP-G3-Touch-Pro)** — current desktop VoIP phone: 5" HD touch, 5MP camera w/ privacy shutter, octa-core A53, dual GbE, Wi-Fi, Bluetooth (G3 handset), third-party SIP or UniFi plug-and-play, no extra licensing. Price: **£176.39** (linitx.com, in stock) [independent].
  Sources: https://linitx.com/product/ubiquiti-unifi-talk-phone-g3-touch-pro-utp-g3-touch-pro/18177 ; https://aceperipherals.com/collections/smart-home/products/ubiquiti-utp-g3-touch-pro-g3-touch-pro-unifi-talk-voip-phone
- Talk subscription: $9.99/mo US launch price incl. 3,000 min/mo (older reference); UK launch (£7.99/mo, Talk App 2.2.0) dates to Nov 2023 per ISPreview — **not a 2026 item; flag against misdating**.
  Source: https://www.ispreview.co.uk/index.php/2023/11/ubiquiti-finally-adds-support-for-the-uk-to-unifi-talk-app.html

## R7. UniFi Access 2026 — Enterprise Access Hub & Reader Pro (new detail)

- **Enterprise Access Hub** (new 2026): discussed in UniFi livestream ~Apr 2026 via creator summary [unverified — show summary only].
- **Access Reader Pro**: sent to reviewers but "not currently searchable on UniFi's website due to major screen tear issues on the device" — **pre-GA issue; flag** [unverified].
  Source: https://www.summarize.tech/summary?url=youtube.com%2Fwatch%3Fv%3Dg2EiKptU1Rs (summary of livestream 190)
- Current shipping readers: **UA-Reader-Lite** — NFC+BLE, 802.3af PoE, MIFARE DESFire, IP54, **£85.49 inc VAT** (netxl.com, 50 in stock) [independent]; **G3 Reader** — NFC card + Identity app + Touch Pass, handwave request-to-exit, IP55 [secondary].
  Sources: https://www.netxl.com/ip-access-control/ubiquiti-unifi-access-reader-ua-reader-lite/ ; https://dataworld.com.au/productpdf/download/file/id/8056/name/Ubiquiti_%257C_UniFi_Access_G3_Reader_-_WHITE.pdf/
- (New UniFi Jan 2026 items adjacent to Access: Fail-Secure/Fail-Safe Strike Lock, SuperLink antenna — from a Jan 2026 product roundup [secondary].)
  Source: https://www.youtube.com/watch?v=-BCVq3vrIpA

## R8. Homelab favorite switches 2026 — community picks (new detail)

- **Community consensus (ServeTheHome forums, LTT, guides, Sep 2026)** [independent/secondary]: TP-Link Omada series for value, Cisco SG200/SG300 (used, enterprise learning), NETGEAR prosumer (Plus/Plus-series), Aruba Instant On (enterprise features without controller), MikroTik for advanced users who accept the learning curve. UniFi popular for unified ecosystem but "more expensive… community feedback on reliability is mixed" [independent].
  Sources: https://vintagevinylnews.com/best-enterprise-switches-for-homelab/ ; https://forums.servethehome.com/index.php?threads/what-is-your-favourite-simple-unmanaged-switch.18380/
- **Budget 10G picks (LTT thread, 2026)** [independent/secondary]: MikroTik **CSS610-8G-2S+IN** (~$100, 8× 1G + 2× SFP+); **CSS326-24G-2S+RM** (~$150, 24× 1G + 2× SFP+); TRENDnet **TEG-30284** (~$210, 24× GbE + 4× 10G SFP+); TRENDnet **TEG-3284WS** (~$450, 24× 2.5G + 4× 10G SFP+); TRENDnet **TL2-F7080** (~$229, 8× 10G SFP+); used **QuantaMesh T1048-LY4R** (~$52, 48× 1G + 2× 10G SFP+) and **Dell N3048** (~$85, noisy 40mm fans).
  Source: https://linustechtips.com/topic/1623571-switch-recommendations/
- **Power-consumption note (STH forum, 2026)** [independent]: 24-port 10GBase-T is a "real power hog" — Netgear XS724EM idles ~60W; SFP+ is the way for low power; Arista 7050SX (48-port SFP+, used) idles ~70W with no licensing games; 8-port SFP+ MikroTik for low port density.
  Source: https://forums.servethehome.com/index.php?threads/recommendations-for-a-10gb-switch.46590/
- **Full-stack 2026 example (Medium, Jun 2026)** [independent]: UXG Fiber gateway + USW Pro 48 core + Flex 2.5G PoE + U7 Pro AP — VLAN-segmented, UniFi-only except the old Omada gear it replaced.
  Source: https://medium.com/@rashaad.s/homelab-chronicles-state-of-the-lab-2026-472c9344f040

## R9. Zyxel XGS2220 / GS2220 — 2026 status check (explicit finding)

- **No new Zyxel XGS2220-series launch in 2026 found.** The XGS2220 (10G SFP+ uplink Lite-L3) remains the current mid line; the **GS2220 series** (10/28/50-port L2 hybrid with tri-mode cloud/CLI/SNMP management and NebulaFlex Pro) remains the current gigabit line — both are older platforms still sold and supported [vendor-reported/secondary]. 2026 Zyxel switch news was the XS1935 (10GbE all-port, §3.2) and GS1915 ("Just Connect", §3.2) — **both SMB-focused, not XGS2220 refreshes**.
  Sources: https://www.zyxel.com/global/en/blogs/new-xgs2220-switches-are-amongst-our-most-versatile-yet ; https://technologyreseller.uk/zyxel-debuts-industrys-first-tri-mode-management-in-new-hybrid-switch/
- GS2220-10HP EU part (PN GS2220-10HP-EU0101F) still listed at EU distributors in 2026 [secondary].
  Source: https://distribucioninformatica.com/gen_pdf/gen_ficha.php?fichagen=3337

## R10. Round-2 verification log

1. EAP772 BE9300 vs PB Tech "BE11000" feature-block text — datasheet rate wins (BE9300) [flag noted above].
2. EAP772 listing "Date Created 23-09-2026" is one day after the research date — likely timezone/crawl artifact; treat as Sep 2026 snapshot [unverified].
3. UniFi Talk UK expansion (£7.99/mo) dates to Nov 2023 — exclude from 2026 claims.
4. Access Reader Pro screen-tear issue: single YouTube-summary source; pre-GA; treat cautiously [unverified].
5. UniFi Talk softphone: plans only, no launch date [unverified].
6. No new MikroTik switch launch in 2026 found — gap, not absence proof.
7. Zyxel USG FLEX H: still no 2026 refresh; S4 stands.
8. QNAP switch prices: no Sep-2026 snapshot — list prices from older sources only.
9. 25G/40G enterprise-grade SMB switches (FS.com, etc.) are Phase C scope — intentionally excluded here.


---

