# NOTES — corpus `etape6-phaseb-smb-networking`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

La source contient 19 titres H1 : un dossier par document H1 (un entête éventuel devient `00-front-matter`).

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 6 Extension — Phase B: SMB / Prosumer Networking` — 2612 lignes, 42 chunks.

- 1. Switch launches 2026 (USW Pro / Aggregation / Enterprise)
- 2. Gateways 2026
- 3. UniFi OS / Network Application software 2026
- 4. Wi-Fi 7 APs (U7) — lineup & pricing, Sep 2026
- 5. Street prices, dated Sep 2026 (key SKUs)
- 6. Adoption figures / financials 2026
- 7. Security incidents / CVEs 2026
- 8. Lawsuits / acquisitions 2026
- 9. Open questions / uncertainties
- 1. TP-Link / Omada
- 2. D-Link
- 3. Zyxel
- 4. EnGenius (brief)
- 5. Grandstream (brief)
- 6. Cross-vendor street-price comparison (2.5G / 10G managed switches, observed Sep 2026)
- S1. TP-Link Omada Pro — full lineup and official MSRP (was: price unconfirmed)
- S2. TP-Link "Agile" easy-managed series — fully detailed (was: thin sourcing)
- S3. TP-Link Wi-Fi 7 AP — EAP770 BE11000 MSRP (Omada 2026 price list)
- S4. Zyxel USG FLEX H — 2026 street pricing (was: absent)
- S5. Zyxel Wi-Fi 7 APs — 2026 current lineup and prices (was: brief)
- S6. NETGEAR — AV/IT M4350 switching line (new vendor)
- S7. HPE Aruba Instant On — SMB line (new vendor)
- S8. D-Link additions: DQS-5000 25G/100G DC switch + Nuclias controllers (was: Nuclias roadmap thin)
- S9. UniFi Enterprise ECS line — official confirmations + memory surcharge (was: SKU/price unconfirmed)
- S10. UISP Fiber XGS — UniFi/UISP PON line (new detail; complements §1 fiber)
- S11. UniFi ecosystem 2026 — G6 Entry intercom + G6 180 camera (new detail)
- S12. UniFi 400G — status re-check (Sep 22, 2026)
- S13. Remaining gaps (not researched in this pass)
- Supplementary / Complementary Research Pass — September 22, 2026 (third wave; S14–S18)
- S14. UISP Fiber — XG (non-symmetric) ONU variant + UK price snapshots (was: XGS ONU only, ZA/KE prices only)
- S15. D-Link Nuclias Unity — cloud platform launched April 2, 2026 (was: Oct 2025 hardware controllers only)
- S16. TP-Link Omada Wi-Fi 7 EAP portfolio — SKUs and September 2026 street prices (was: EAPs absent)
- S17. Flagship / adjacent Wi-Fi 7 APs — Zyxel WBE660S and Grandstream GWN7670LR (was: S5 covered budget NWA + WBE510D/630S; Grandstream had no Wi-Fi 7)
- S18. New SKUs / products not previously captured: Instant On 1960 multi-gig SKU, UniFi Travel Router Long-Range, homelab 25G options
- R1. TP-Link Omada Wi-Fi 7 APs — EAP772 / EAP772-Outdoor / EAP783 (complements S3 which had EAP770 BE11000)
- R2. NETGEAR Pro Wi-Fi 7 APs — WBE710 (BE9400) / WBE700 / WBE750 (complements S6)
- R3. Aruba Instant On Wi-Fi 7 — status re-check (complements S7)
- R4. MikroTik 25G/100G switch pricing — Sep 2026 snapshots (new detail)
- R5. QNAP switch pricing context (brief)
- R6. UniFi Talk 2026 — software & hardware updates (new detail)
- R7. UniFi Access 2026 — Enterprise Access Hub & Reader Pro (new detail)
- R8. Homelab favorite switches 2026 — community picks (new detail)
- R9. Zyxel XGS2220 / GS2220 — 2026 status check (explicit finding)
- R10. Round-2 verification log
- R1. TP-Link Omada Cloud Standard licensing — pricing RESOLVED (was: gap #1)
- R2. D-Link DXS-3130 — street pricing (was: launch announced, no prices)
- R3. Grandstream Wi-Fi 7 — lineup CONFIRMED (corrects main §5 flag)
- R4. NETGEAR business Wi-Fi 7 — WBE710 / WBE750 detail (new)
- R5. UniFi Pro XG — official US list prices + memory surcharge (Sep 2026)
- R6. UniFi Access — G6 Pro Entry + UK line pricing (new)
- R7. Zyxel XS1935 — retail pricing surfaces (was: unannounced)
- R8. EnGenius 2026 — Cloud-Lite + L3 core/aggregation launches (new)
- R9. UniFi UNAS storage lineup — full 2026 detail (was: name-dropped in §2.3)
- R10. Homelab / prosumer favorites 2026 (new)
- R11. Market shock — US FCC foreign-made consumer-router import ban, Mar 2026 (new; verify against FCC record)
- R12. Re-checks that held (no change)
- R13. Round-2 verification log / sources used
- T1. MikroTik — homelab/SMB angle (new vendor to this file)
- T2. Grandstream Wi-Fi 7 — NEW launch evidence (corrects S13 gap #7)
- T3. Cross-vendor Wi-Fi 7 AP comparison (SMB-relevant, Sep 2026)
- T4. Zyxel XGS/XMG switch street prices 2026 (was: absent or MSRP-only)
- T5. Aruba Instant On — 2026 status refresh
- T6. 2026 cost shock — memory prices, tariffs, and their SMB impact (new dimension)
- T7. Remaining gaps after Pass #2
- Supplementary / Complementary Research Pass — September 22, 2026 (fourth wave; S19–S23)
- S19. TP-Link Omada Pro AP portfolio MSRP + controller pricing status (was: Omada Pro pricing partly open)
- S20. Zyxel XS1935 street prices captured (was: pricing unannounced)
- S21. NETGEAR — two NEW M4350 models launched at ISE 2026 (was: no confirmed 2026 launches)
- S22. Grandstream Wi-Fi 7 — 2026 GA status resolved (was: GA unconfirmed)
- S23. D-Link DQS-5000 deep-dive + Instant On availability snapshot
- Supplementary pass #4 — verification notes / open gaps
- Supplementary / Complementary Research Pass — September 22, 2026 (fourth wave; S19–S24)
- S19. Zyxel XGS2220 series — 2026 street prices + GS1915 price data points (was: old-series context only)
- S20. NETGEAR AV/IT lines — M4250 full price table + M4500 context + extra M4350 SKUs (was: M4350 only)
- S21. MikroTik 100G/25G street pricing — CRS520 / CRS504 / CRS304 (was: community mentions only)
- S22. EnGenius 2026 detail — ECW526 street price, ECW520, ECS5512F (was: ECW515/ECW536S only)
- S23. UniFi UDR7 — GA confirmed, September 2026 regional prices/availability (was: name only)
- S24. Grandstream GWN7700M unmanaged + 2026 supply/availability notes (was: GWN7700MP only; no supply section)
- Supplementary / Complementary Research Pass — 2026-09-22 (independent research passes A/B/C)
- T1. D-Link DAP-E3620 Series — NEW Wi-Fi 7 BE3600 APs (announced Aug 19, 2026) (was: absent)
- T2. D-Link DAP-E9560 — BE9500 Wi-Fi 7 AP street prices (was: named in Japan guide only)
- T3. Zyxel USG FLEX 50H / 50HP — smaller SMB models (was: only 100H–700H in S4)
- T4. Grandstream GWN7810/7816 — retail pricing Sep 2026 (was: specs only, no prices)
- T5. NETGEAR M4500 — 25G/100G Pro AV line pricing Sep 2026 (complements S6 M4350)
- T6. HPE Aruba CX 6000 — NEW CX-series switch for SMB/retail (NRF 2026) (was: absent)
- T7. Aruba Instant On — divestiture status (complements S7; was: convergence-only context)
- T8. TP-Link Omada hardware controller pricing — Sep 2026 snapshots (closes the §1.2/S1/S13 gap)
- T9. EnGenius AirGuard + ECW515 detail (complements §4)
- T10. Round-3 verification log
- U1. TP-Link regulatory exposure — supplemental detail (developing story, [unverified] beyond cited secondary sources)
- U2. TP-Link Omada Cloud Standard license pricing (new, closes a gap)
- U3. TP-Link EAP783 Wi-Fi 7 (new product data point)
- U4. Zyxel XS1935 — port-map corrections and pricing status
- U5. Zyxel Nebula Pro Pack — actual pricing found (closes earlier gap)
- U6. Grandstream GWN7816 / GWN7816P (new — above the GWN7800 line)
- U7. Aruba Instant On 1930 — street-price snapshots (JL683B)
- U8. NETGEAR WBE750 Wi-Fi 7 (ecosystem data point)
- U9. EnGenius 2026 enterprise launches (new — above the SMB ECS line)
- U10. 2026 homelab consensus — 10G/2.5G managed picks (independent community data)
- U11. D-Link — extra 2026 price points
- U12. UniFi Access — 2026 device detail (ecosystem context)
- U13. Open gaps after Pass #3
- B1. TRENDnet 2026 — full portfolio snapshot (new vendor; was: only R8 budget-pick mentions)
- B2. Zyxel GS1915 — street prices (closes the repeated §3.2 gap)
- B3. EdgeRouter / EdgeMAX — 2026 status (new; not previously researched in this file)
- B4. Ax Wireless v. Ubiquiti ITC — status re-check (Sep 22, 2026; new)
- B5. Addendum to R1 — Omada Cloud Standard multi-year license data points (new)
- B6. Addendum to R5 — USW-Pro-XG-Aggregation checkout totals and street snapshots (new)
- B7. Verification log
- T1. NETGEAR — NEW M4350 models at ISE 2026 (Jan 2026) [vendor-reported via trade press]
- T2. Aruba Instant On 1960 — NEW 12-port multi-gigabit model (S0F35A) [independent]
- T3. HPE Aruba CX 6000 — NEW L2 branch/retail switch series, NRF 2026 [secondary]
- T4. UniFi Network Application release train — correction/addition [official via community trackers]
- T5. TP-Link Omada Pro — Wi-Fi 6 AP MSRPs + controller street prices [official/independent]
- T6. TP-Link EAP772 — Sep 2026 street prices (listings created Sep 22, 2026) [independent]
- T7. Zyxel XS1935-12F — retail sightings (was: price unannounced) [independent]
- T8. D-Link DXS-3130 — street prices + full specs (Sep 2026) [independent]
- T9. UISP Fiber XGS — 2026 retail price snapshots [independent]
- T10. Grandstream GWN7816/GWN7816P — pricing [independent/official]
- T11. Supply/availability context — the 2026 memory-price crisis ("memflation") [independent/secondary]
- T12. Round-3 verification log / remaining gaps
- V1. Ubiquiti E7 family — commercial-product confirmation (corrects the "unconfirmed" header-table row)
- V2. NETGEAR M4350 — MAP prices Sep 2026 (closes the "street prices not captured" gap)
- V3. EnGenius ECS8830F / ECS8854F — datasheet detail (complements the Jul 16, 2026 launch)
- V4. TP-Link EAP772 — UK price data point (complements R1)
- V5. Open gaps after Round 5
- U14. UniFi Flex 2.5G-8-PoE (USW-Flex-2.5G-8-PoE) — pricing and availability snapshot (new)
- U15. Aruba Instant On 1930 — per-SKU street snapshots (new; was: series overview only)
- U16. UniFi gateway / cloud-gateway price re-check (new; complements §9)
- U17. UNVR Gen 2 — UniFi NVR pricing (ecosystem context, new)
- U18. TP-Link ER8411 + OC200 — street price add-ons (complements existing lines)
- U19. MikroTik CRS520 — EU/ZA price add-ons (complements S21)
- U20. Round-4 verification log
- U1. NETGEAR — NEW 2026 hardware: M4350-16M4V + M4350-16C at ISE 2026 (corrects "no confirmed NEW NETGEAR switch model in 2026" flag)
- U2. Zyxel GS1915 — street prices surface (resolves T7 gap #3 for GS1915); XS1935-12F shipping
- U3. Grandstream GWN7670 — GA confirmed + EU street price (corrects T7 gap #1)
- U4. EnGenius ECW510 — $129 MSRP detail (complements §4)
- U5. TP-Link Omada Pro controller — explicit finding: no Omada-Pro-branded controller SKU
- U6. Aruba Instant On — 2026 status re-check (explicit finding: no new hardware; CX 6000 is adjacent enterprise)
- U7. Fourth-wave verification log + remaining gaps
- V1. Zyxel NWA210BE — specs and pricing RESOLVED (was: U13 gap #3 — model name only)
- V2. Zyxel XS1935-12F — street pricing found (was: U13 gap #2 — regional pricing/availability unannounced)
- V3. TP-Link EAP783 — BE19000 vs BE22000 naming conflict RESOLVED (was: U13 gap #5)
- V4. Grandstream GWN7816 — stacking support RESOLVED + price points (was: U13 gap #6)
- V5. Aruba Instant On 1930R successor — re-check (was: U13 gap #7)
- V6. D-Link DQS-5000 — second price attempt + spec correction (was: U13 gap #4)
- V7. MikroTik — 2026 launches: CRS804 DDQ 400GbE + CRS304-4XG-IN (new detail)
- V8. UniFi U7 Pro Max — September 2026 price snapshots (flag: 2024 launch, not 2026)
- V9. Wave-5 verification log / remaining gaps
- V1. UniFi Network Application standalone — 2026 EOL transition to UniFi OS Server (was: not researched)
- V2. UniFi Protect 7.1 (May 2026) — feature detail (was: only security-fix versions mentioned)
- V3. Zyxel ZyWALL ATP — 2026 status (was: absent)
- V4. Round-3b verification log
- S25. UniFi U7 Pro XGS / U7 Pro XG / E7 family — launch CONFIRMED (resolves retained gap)
- S26. UniFi Cloud Gateway Max (UCG-Max) — 2026 street pricing and availability
- S27. NETGEAR Insight cloud subscription — 2026 pricing state
- S28. MikroTik 2026 hardware wave — MWC 2026 announcements and RouterOS switches
- S29. TP-Link Omada 2026 access-point launches — EAP775-Wall/Outdoor, EAP783, EAP787, EAP725
- S30. Grandstream GWN7672 (BE11000 Wi-Fi 7) — street pricing confirmed (resolves retained gap)
- S31. Omada SDN Controller software state — no v6.4 as of 22 Sept 2026
- S32. UniFi Network Application — 10.7 status and release-train notes
- S33. This wave's gap closure summary

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
