# NOTES — corpus `etape6-phasef3-platform-security`

## Génération

Mode auto (`RAG/_tools/build_rag.py`), cibles **45–90 lignes**, partition par titres H1 → H2 → H3 sans couper au milieu d'un paragraphe.

`first_is_content: true` (H1 unique = contenu) : dossier `00-platform-security` (`folder_name`), un fichier par bloc H2.

La source ne contient **aucune ancre** HTML : navigation via `INDEX.md` et `manifest.json`.

## Contenu

`Step 6 Phase F3 — Platform Firmware & Hardware Root of Trust` — 752 lignes, 14 chunks.

- 1. BIOS/UEFI fundamentals and the firmware supply chain
- 2. Secure Boot architecture and the 2026 certificate transition
- 3. BlackLotus and the Secure Boot bypass problem
- 4. Firmware CVEs of 2025–2026
- 5. Open firmware: Coreboot, OpenBMC, LinuxBoot
- 6. Root of Trust concepts and NIST SP 800-193
- 7. Silicon RoT implementations: Intel PFR, AMD PSP, Pluton
- 8. Caliptra and OpenTitan: open silicon Roots of Trust
- 9. TPM 2.0: chips, firmware TPMs, and the PQC transition
- 10. Hardware Security Modules (HSM)
- 11. Self-encrypting drives (SED): TCG Opal, FIPS drives, key management
- 12. Confidential computing: TDX, SEV-SNP, Arm CCA, GPU TEEs
- 13. Verification and audit tooling
- 14. Gaps, conflicts, and open items
- 15. Intel SGX lifecycle in 2026: enclaves vs confidential VMs
- 16. DDRop: September 2026 research undermining TEE memory protection
- 17. AMD SEV-SNP deep dive
- 18. UEFI internals: boot phases, NVRAM, SPI flash, SMM
- 19. Measured boot, DICE, SPDM, and attestation standards
- 20. Server OEM firmware management and supply-chain security
- 21. Post-quantum cryptography timeline for platform security
- 22. Standards and specification index
- 23. Glossary
- 24. BlackLotus anatomy and Secure Boot bypass mechanics
- 25. Firmware threat landscape 2024–2026
- 26. Open firmware ecosystem 2026
- 27. TPM 2.0 operational detail
- 28. HSM deployment patterns and selection guidance
- 29. SED standards detail
- 30. Procurement checklist: platform security
- 31. Firmware update and recovery playbook
- 32. Measured-boot and attestation flow (reference)
- 33. Threat-model matrix (platform security)
- 34. Source index (verbatim URLs used)
- 35. Linux Secure Boot: shim, MOK, distro signing
- 36. Windows Secured-Core and System Guard
- 37. Apple and Google platform RoTs
- 38. FIPS 140-3 and Common Criteria programs
- 39. Confidential-containers and enclave frameworks
- 40. Key management and code-signing infrastructure
- 41. CHIPSEC module catalog (audit reference)
- 42. fwupd/LVFS operations
- 43. TPM provisioning and TCG guidance
- 44. SED operations detail
- 45. PQC migration playbook (platform security)
- 46. Caliptra subsystem integration notes
- 47. BMC security hardening
- 48. Secure Boot certificate transition operations (2026)
- 49. Glossary extension
- 50. Intel TDX deep dive
- 51. Arm CCA deeper
- 52. NVIDIA confidential computing deeper
- 53. Firmware forensics and analysis tools
- 54. UEFI supply chain: IBV and OEM model
- 55. Key dates timeline 2024–2026
- 56. Cross-references to other Phase F tracks
- 57. Secure Boot: servers vs clients
- 58. Glossary additions

## Non audité

Aucun défaut de source recensé ; fidélité garantie par le vérificateur (concaténation ==
source), pas par une relecture sémantique.

## Régénérer

```bash
python3 RAG/_tools/build_rag.py
python3 RAG/_tools/verify_rag.py
```
