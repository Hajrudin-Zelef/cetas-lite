---
id: etape6-phaseb-smb-networking/07-supplementary-complementary-research-pass-round-3-september-/part-4
title: "Supplementary / Complementary Research Pass — Round 3 — September 22, 2026 (part 4)"
domain: supplementary-complementary-research-pass-round-3-september-
role: deep-dive
task: reference
actors: ["Intel"]
dates: ["2025-10", "2025-12", "2026-01-07", "2026-03", "2026-05", "2026-07-29"]
keywords: ["cost"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [1334, 1349]
section: "Supplementary / Complementary Research Pass — Round 3 — September 22, 2026"
sha256: d358ec6236244213e816ffa066751d0948013f301aa850a98f98c0b8692af6d3
---

# Supplementary / Complementary Research Pass — Round 3 — September 22, 2026 (part 4)

**"Agile Switches" details.** "Agile" = TP-Link's Easy Managed switch class (not a separate Pro line) [official]: https://www.omadanetworks.com/us/business-networking/omada-switch-agile/. The 2026 ISC West flyer (March 2026) lists the Agile series lineup [official]:
- ES205GP (5× GbE, 4× PoE+, 65 W budget), ES210GMP (10-port, 8× PoE+, 123 W), ES220GP (20-port, 16× PoE+, 150 W), ES228GP (28-port, 24× PoE+, 250 W), ES206XPP-M2 (6-port 2.5G with 4× PoE++, 120 W + 1× 10G SFP+). Feature pitch: cost-effective SDN compatibility, up-to-820 ft long-range PoE on gigabit models, IntelliRecover remote reboot. Source: https://static.tp-link.com/document/pdf/en/isc-west-2026/2026_ISC_WEST_Omada_Business_Switches.pdf
- The full Agile table adds: ES210XPP-M2 (8× 2.5G PoE++, 200 W), ES210X-M2, ES206X-M2 (2.5G non-PoE), ES228GMP (384 W), ES224G, ES220GMP, ES216G, ES210GP (63 W), ES208GP/ES208G, ES206GP. All desktop/wall-mount except 28-port models (rackmount); nearly all fanless. Source: https://www.omadanetworks.com/us/business-networking/omada-switch-agile/
- Feature limits (Easy Managed class): **no 802.1X control, no 802.1p priority, no LLDP-MED, no DHCP L2 relay, no trust mode** on port profiles [independent, from controller UI text]: https://github.com/daily-nerd/terraform-provider-omada/blob/HEAD/docs/SWITCH_CLASS_MATRIX.md. Agile switches require explicit per-site activation in the controller (community Q&A, Jan 2026) [secondary]: https://community.tp-link.com/en/business/threads/topic/853584.
- **Street prices for Agile models: not found** — flagged unknown.

#### B.2. D-Link Nuclias 2026

**2026 platform releases.** The next-gen Nuclias hardware controllers (DNH-1000 up to 500 devices, DNH-3000 up to 1,500 devices, software-based DNC-5000) were announced **October 2025** (PR Newswire) — pre-cutoff context, included only for continuity [vendor-reported]: https://beta.manilatimes.net/2025/10/16/tmt-newswire/pr-newswire/d-link-unveils-nuclias-network-controllers/2202000. **No 2026-versioned Nuclias Connect software release was found** — flagged unknown (latest version number not identified in 2026 sources).

**2026 Nuclias-compatible hardware.**
- **DAP-E9560 BE9500 Wi-Fi 7 ceiling-mount AP** and **DAP-X3060W AX3000 Wi-Fi 6 wall-plate AP** announced January 7, 2026 [vendor-reported]: DAP-E9560 = tri-band Wi-Fi 7, 320 MHz channels, 4096-QAM, 1× 10GbE PoE + 1× 2.5GbE LAN, WPA3 Enterprise, Nuclias-managed; DAP-X3060W = dual-band AX3000 in-room wall-plate, PoE, multiple GbE ports for IP phones/TVs. Source: https://facilityexecutive.com/d-link-unveils-next-generation-enterprise-wireless-lineup. Note: the PR wire copy carries a December 2025 wire date while trade press dates it Jan 7, 2026 — treat as late-Dec 2025/early-Jan 2026 announcement [secondary]: https://aapnews.aap.com.au/aapreleases/cision20251218AE49456.
- **Prices for DAP-E9560 / DAP-X3060W: not found** — flagged unknown.
- D-Link Japan's "Product Guide 2026-27" (May 2026) lists as planned releases (リリース予定): **DAP-E9560** (Wi-Fi 7 tri-band 2×2×2, 10G PoE 802.3bt), **DAP-E3620** (Wi-Fi 7 2×2), **DAP-E3620OU** (Wi-Fi 7 2×2 outdoor) — i.e., two additional Wi-Fi 7 APs still in "planned" status as of mid-2026 [vendor-reported]: https://files.actibookone.com/contents/4200/737165-20260519113745/original.pdf
- Nuclias-adjacent (surveillance, July 29, 2026): D-Link expanded the Nuclias surveillance portfolio for SMBs with BCB-P01/BCD-P01 Full HD IP cameras and BNR-004/BNR-008 NVRs [vendor-reported]: https://uktechnews.co.uk/2026/07/29/d-link-expands-nuclias-surveillance-portfolio-for-small-businesses/.

