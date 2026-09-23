---
id: etape6-phaseb-smb-networking/01-part-1-ubiquiti-unifi/6-adoption-figures-financials-2026
title: "6. Adoption figures / financials 2026"
domain: part-1-ubiquiti-unifi
role: deep-dive
task: reference
actors: ["United States"]
dates: ["2026-05", "2026-06-30", "2026-07"]
keywords: ["acquisition", "advisory", "agent", "disclosure", "export controls", "lawsuit", "revenue", "throughput"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [132, 187]
section: "Part 1 — Ubiquiti UniFi"
sha256: c7f34da805f8e150518c52219bf3d673ca1eca8756ade595c809bcd4d67cdefa
---

# 6. Adoption figures / financials 2026

## 6. Adoption figures / financials 2026

### 6.1 Fiscal 2026 results (year ended June 30, 2026) [official via press release; reported Aug 20, 2026]
- **Total revenue: $3,274.2M (+27.2% YoY)**; Q4 FY2026 revenue **$937.3M — record quarter (+23.5% YoY, +18.9% QoQ)**.
- Enterprise Technology: FY $2,972.3M (FY25: $2,254.3M, **+31.8%**); Q4 $868.3M (**92.6% of total**, +27.7% YoY).
- Service Provider Technology: FY $301.9M (−5.5%); Q4 $69.0M (−12.7%) — segment in structural decline.
- Q4 gross margin 45.8%; operating income $340.0M; GAAP EPS $4.70; non-GAAP EPS $4.73 (beat $4.48 consensus).
- Capital returns: quarterly dividend **$1.00/share**; **$500M share repurchase** authorization renewed through Sep 2027. Stock ~$606 after hours on results [secondary].
- Q3 FY2026 (Mar 31, 2026): revenue $788.2M; Q2 FY2026 (Dec 31, 2025): $814.9M — consistent growth trajectory through the year.
- Sources: https://www.morningstar.com/news/business-wire/20260820253715/ubiquiti-inc-reports-fourth-quarter-fiscal-2026-financial-results ; https://www.tradingkey.com/news/earnings/262124400-tradingkey

### 6.2 Positioning: enterprise push + homelab strength
- **Enterprise**: EFG/UXG-Enterprise ($1,999, NeXT AI SSL inspection, VRRP), UDM-Beast ($1,499, 25G IPS/IDS), Campus Core switches with MC-LAG/OSPF/BGP/VRRP — explicit targeting of Meraki/Fortinet/SonicWall price points (EFG vs FortiGate 400F/Meraki MX450/SonicWall 4700 comparisons in trade press [secondary]). Reviewers note inter-VLAN routing throughput as the weak spot vs price (2GT Media, Feb 2026) [secondary].
- **Homelab/prosumer**: STH review series on Pro XG switches (Sep 2026); LTT "100Gb/s Switches from Ubiquiti" (Apr 2026, enthusiastic); GitHub homelab roadmaps standardizing on USW-Pro-Max-24-PoE, UNAS Pro 4, UNVR-G2 [independent/secondary]. 9to5Mac (2025, still cited): Ubiquiti "should be today's default choice for many enterprise environments" on Wi-Fi 7 value.
- **Notable events**: UWC (UniFi World Conference) London 2026 — show floor with EFG Core (4×100G), Campus Core 32×100G, UniFi 5G, Express 7 3-pack, Pro AV line, G6 Edge cameras [unverified — YouTube]. ISC West 2026 (Apr 2026) — UDM Beast debut, UDR 5G Max, ENVR Core [secondary/YouTube].
- No hard unit-sales or market-share figures published by Ubiquiti; adoption evidence is indirect (revenue mix, sell-outs at retailers, review coverage) [uncertainty flagged].

---

## 7. Security incidents / CVEs 2026

### 7.1 Bulletin 062 — Mar 18, 2026 [secondary]
- **CVE-2026-22557**: path traversal in UniFi Network Application, **CVSS 10.0, unauthenticated**; fixed in 10.1.89 (official) / 10.2.97 (RC) / UX firmware 4.0.13. CyCognito flagged emergency-patch status for internet-exposed controllers.
- **CVE-2026-22558**: authenticated NoSQL injection → privilege escalation, CVSS 7.7.
- Source: https://www.cycognito.com/blog/emerging-threat-ubiquiti-unifi-network-application-path-traversal-cve-2026-22557/ ; https://cybersecuritynews.com/ubiquiti-unifi-vulnerabilities/

### 7.2 July 2026 advisory batch [secondary]
- **CVE-2026-50746** (UniFi Connect, CVSS 10.0, command injection; fixed 3.4.20); **CVE-2026-50747** (Talk, 9.9; fixed 5.2.2); **CVE-2026-50748** (Access, 9.9; fixed 4.2.29); CVE-2026-54400 (Access, 9.1); **CVE-2026-54402** (UniFi OS command injection, 9.9; fixed OS 5.1.19); CVE-2026-54403 (path traversal, chainable); CVE-2026-54404 (SQLi); **CVE-2026-55115** (Protect SSRF, 9.9; fixed 7.1.83); CVE-2026-55116 (UniFi OS improper access control, 9.0).
- Sources: https://thecyberexpress.com/cve-2026-50746-ubiquiti-unifi-os-vulnerability/ ; https://vulert.com/blog/ubiquiti-unifi-critical-flaws/

### 7.3 "21 critical flaws" disclosure — ~Aug 26, 2026 [secondary]
- 21 critical vulns (CVSS 8.2–10.0): auth bypass, command injection, privilege escalation across Network 10.5.67, OS Server 5.1.37, Protect 7.2.105, Access 4.3.5, Talk 5.3.2, Connect 3.24.22, UID Enterprise Agent 1.62.1, Connect Display Cast Pro 1.0.111, Protect AI Key 2.2.6. Report notes **active exploitation of UniFi OS earlier in 2026** [unverified detail — single secondary claim].
- Source: https://cybersecuritynews.com/21-critical-ubiquiti-unifi-flaws/

---

## 8. Lawsuits / acquisitions 2026

- **Ax Wireless v. Ubiquiti — ITC + N.D. Ill. (filed Feb 2, 2026)** [vendor-reported via UI regulatory filing]: Wi-Fi 6 patent infringement (US 10,079,707; 10,917,272; 11,646,927; 11,777,776; 11,812,134); complainant seeks **limited exclusion order + cease-and-desist** barring import/sale of wireless products in the US; parallel district-court damages action. UI: "plans to vigorously defend"; adverse ITC ruling "would have a material adverse effect." Loss not estimable. Source: https://www.tipranks.com/news/the-fly/ubiquiti-discloses-complaint-with-itc-filed-by-ax-wireless-thefly
- **Velocity patent complaint — D. Del. No. 1:26-cv-00942-UNA (filed Jul 30, 2026)** [secondary]: infringement of 802.11ax-related patents ('213, '832); accused products list confirms 2026 model names incl. UDR7, UX7, UDR-5G-Max, **U7 Pro XGS, U7 Pro XG Wall**, E7/E7 Campus/E7 Audience, U7 Pro Outdoor. Source: https://litnews.ai/new-suits/suit/1489/complaint.pdf
- **Kovalenko v. Ubiquiti, S.D.N.Y. No. 1:26-cv-06838 (2026)** [secondary]: suit by Ukrainian civilians alleging Ubiquiti airMAX equipment reached Russian forces (via intermediaries, despite US export controls since Mar 2022) and was used in drone attacks; alleges UI telemetry/firmware-update visibility and failure to geofence/disable. UI's Mar 2026 10-Q said the conflict had not materially affected operations. On **Sep 10, 2026**, Schall, Brown & Schwartz LLP announced a shareholder investigation into UI directors/officers over the allegations. Sources: https://hrtechedge.com/hr/ubiquiti-faces-lawsuit-over-alleged-russia-battlefield-network-role/ ; https://www.businesswire.com/news/home/20260910413500/en/Long-Term-UI-Shareholders-SBS-Law-Investigates-Ubiquiti-Inc.-Directors-and-Officers-Over-Alleged-Supplies-to-Russian-Forces
- **Acquisitions: none found in 2026** [uncertainty — no acquisition news surfaced].

---

## 9. Open questions / uncertainties
1. Exact retail SKUs + US street prices for the new Enterprise Campus Switch Core line (only Notebookcheck's $5,393/$4,314 figures found).
2. USW-Pro-XG-Aggregation and USW-Pro-XG-24/24-PoE/48-PoE US street prices — not confirmed.
3. UDM-Beast GA date: "early May 2026" inferred from review cluster; one Sep 2026 article re-reports launch.
4. EFG Core (4×100G) and 400G-class products: preview only, no launch confirmation.
5. No unit-sales/market-share figures from Ubiquiti; adoption is revenue-implied only.
6. Long-standing ui.com list prices (UDM-Pro $379 etc.) not re-verified against the live store.


---

