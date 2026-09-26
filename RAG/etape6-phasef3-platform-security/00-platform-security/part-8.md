---
id: etape6-phasef3-platform-security/00-platform-security/part-8
title: "Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust (part 8)"
domain: step-6-phase-f3-platform-firmware-hardware-root-of-trust
role: deep-dive
task: hardware
actors: ["AMD", "AWS", "Google", "Intel", "Nvidia"]
dates: ["2026-09"]
keywords: ["amd", "aws", "gpus", "intel", "memory", "nvidia"]
source: docs/RAG/etape6_phaseF3_platform_security.md
source_anchor: ""
source_lines: [190, 200]
section: "Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust"
sha256: 696a371acaf1d543cd5d001750f6738d2d8799bcfe90597d34a6f262223575d1
---

# Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust (part 8)

- **DDRop** (disclosed September 2026) is an active DDR5 interposer attack against Intel TDX, AMD SEV-SNP, and Intel Scalable SGX [secondary](https://www.innovateksolutionsinc.com/news/ddrop-attack-undermines-intel-tdx-and-amd-sev-snp-memory-protection-researchers-say).
- Unlike passive interposer attacks (TEE.fail) or DDR4-only active attacks (Battering RAM), DDRop performs **write-dropping**: selectively dropping memory writes to break freshness [secondary].
- On Intel TDX, researchers turned write-dropping into full control of a protected VM: reading a victim's private memory, switching a victim machine into debug mode, and overwriting the launch measurement the VM uses to prove trusted state to remote customers [secondary].
- Two of the TDX results were shown only under TDX's default logical-integrity mode; the stronger cryptographic-integrity mode would block them, though attestation forgery would still work because the write happens inside the attacker's own VM [secondary].
- On AMD SEV-SNP the result is narrower: dropping writes during AMD's page-relocation feature let researchers copy one victim page's contents into another [secondary].
- All three technologies encrypt memory without a freshness check — the property DDRop exploits; Intel's older Client SGX is not affected because it uses a hardware integrity tree that catches stale data (though Intel has retired client SGX) [secondary].
- NVIDIA's confidential-computing GPUs are out of reach because their memory sits inside the chip package; Arm CCA was not tested but researchers say it may be affected too [secondary].
- No evidence of DDRop or comparable active interposers used outside a laboratory; no cloud service shown broken into [secondary].
- There is no simple patch: the weakness is in hardware design; closing it needs new memory-encryption hardware with both integrity and freshness; software mitigations (restricting abused memory-management features, verifying writes landed, detecting interposers at boot) raise the bar without removing the root cause [secondary].
- DDRop targets cloud servers specifically — TDX, Scalable SGX, and SEV-SNP are offered by AWS, Azure, and Google Cloud [secondary].

