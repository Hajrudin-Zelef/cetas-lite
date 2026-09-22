---
id: briefing-general-tech-2026/05-networking-optics/08-ecoc-2026-malaga
title: "ECOC 2026 in Malaga"
domain: networking-optics
role: deep-dive
task: networking
actors: ["Credo", "FS.com", "Huawei", "MACOM", "Nokia", "OIF", "Telxius"]
dates: ["2026-09-09"]
keywords: ["asic", "dsp", "ethernet", "optics", "serdes", "training"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g06-8"
source_lines: [6090, 6166]
sha256: 008694a64bc1f55643e29ebc2fb93762831e0995606bb93af420038998a31502
---

# ECOC 2026 in Malaga

<a id="g06-8"></a>
### 6.8 ECOC 2026 in Malaga

ECOC — the European Conference on Optical Communications, the industry's
principal optical-networking exhibition — ran **September 21–23, 2026, in
Malaga, Spain**. As of this dossier's cutoff (September 22), the show was
**ongoing: no show reports had been published yet.** That is normal, not a
contradiction — exhibition reporting lands after the floor closes — but it
means every ECOC claim in this chapter is either a pre-show announcement or an
announced demo, never a reviewed result.

The one confirmed ECOC-2026 data point with independent standing comes from
the **Ethernet Alliance** (GlobeNewswire, September 9, 2026): a **live
multi-vendor interoperability demonstration at booth #2173**, covering 1.6T
OSFP optics, 224G SerDes with link training, and 800G AI-lossless networking
features (LLR/CBFC). The participant list reads as a who's who of the
ecosystem: Amphenol, Cisco, HPE, Huawei, Keysight and others. This is the
interoperability process working as designed — competing vendors proving
their 1.6T gear interworks on a show floor — and it is the substantive
counterpart to the single-vendor "verified" claims dissected in section 6.2.

On the standards front, **IEEE P802.3dj** — the project defining 1.6 Tb/s
Ethernet — continued progressing through 2026, tracking the commercial ramp
rather than leading it (a familiar pattern: 1.6T was shipping before the
standard was finished, as section 6.6 documents).

| ECOC 2026 data point | Booth | Status as of 22/09 |
|---|---|---|
| Ethernet Alliance multi-vendor interop demo (1.6T OSFP, 224G SerDes link training, 800G AI lossless LLR/CBFC; Amphenol, Cisco, HPE, Huawei, Keysight…) | #2173 | Announced 9/09, demo live during show |
| OIF 1600ZR multi-vendor demo | #2126 | Announced, demo live during show |
| MACOM live demos (incl. 3.2T / 448G-per-lane solutions) | #1055 | Announced 15/09 (see 6.10) |

#### What ECOC is, and why the timing matters

ECOC — the European Conference on Optical Communications — is the optical
networking industry's principal annual exhibition and conference, the
venue where the year's transceiver, DSP and coherent announcements are
staged for customers. The 2026 edition ran September 21–23 in Malaga. The
timing explains the September announcement cluster documented across this
chapter: vendors schedule launches for the week before ECOC so the press
releases are fresh when buyers walk the floor. FS.com's September 7
portfolio release, the OIF's September 9 agreement, the September 15 trio
(Credo, Telxius–Nokia, MACOM) and FS's September 21 Wi-Fi announcement are
not a coincidence of the news cycle — they are the industry's coordinated
drumbeat ahead of its main trade show.

#### The Ethernet Alliance demo, decoded

The booth #2173 demo's three items each address one layer of the 1.6T
interoperability problem. **1.6T OSFP optics** is the module layer: can
vendor A's transceiver talk to vendor B's switch port. **224G SerDes with
link training** is the electrical layer: the 224-gigabit-per-second
serializer/deserializer lanes inside the host ASIC must negotiate
equalization settings with the module at link-up — "link training" is that
negotiation, and demonstrating it multi-vendor proves the electrical
interface is truly standard, not just standard-adjacent. **800G AI-lossless
networking (LLR/CBFC)** is the fabric layer: link-level retry (LLR) and
credit-based flow control (CBFC) are the mechanisms that make an Ethernet
fabric behave losslessly for AI collective traffic — the feature set that
lets Ethernet credibly contest InfiniBand in the AI cluster. A participant
list spanning Amphenol, Cisco, HPE, Huawei and Keysight — component,
system and test vendors together — is what makes the demo more than a
press release: the test-equipment vendor in the booth is the tell that
real measurements are being taken.

#### P802.3dj: the standard chasing the shipments

IEEE P802.3dj is the project writing the 1.6 Tb/s Ethernet standard. Its
2026 status — progressing, not finished — while 1.6T is already shipping in
volume is the normal order of operations in Ethernet history, not a
dysfunction: 40G, 100G and 400G all shipped in volume before their IEEE
standards closed. The industry ships on multi-source agreements and
interoperability demos (like the one above) and standardizes afterwards.
The reader should therefore treat "the standard isn't finished" as
neutral information about process timing, never as doubt about the
technology's readiness — the shipments are the readiness proof.

