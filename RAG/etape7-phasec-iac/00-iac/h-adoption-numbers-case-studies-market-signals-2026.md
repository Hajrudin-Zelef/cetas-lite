---
id: etape7-phasec-iac/00-iac/h-adoption-numbers-case-studies-market-signals-2026
title: "H. Adoption numbers, case studies, market signals (2026)"
domain: step-7-phase-c-iac-platform-automation
role: deep-dive
task: reference
actors: []
dates: ["2025-02", "2026-03-31", "2026-06", "2026-08-19", "2026-08-25", "2026-09-15"]
keywords: ["acquisition", "agents", "license", "pricing"]
source: docs/RAG/etape7_phaseC_iac.md
source_anchor: ""
source_lines: [326, 375]
section: "Step 7 Phase C — IaC & Platform Automation"
sha256: bcce7669377291053d72764a5f68ff46800a3b3b6a25bf92a035e298cdeb7a42
---

# H. Adoption numbers, case studies, market signals (2026)

## H. Adoption numbers, case studies, market signals (2026)

- Terraform is described as "indistinguishable from infrastructure as code" with ~40k GitHub stars and 9k+ forks cited in 2023-era analyses; it remains the default among existing users in 2026 [secondary](https://www.develeap.com/magazine/About-Terraform-Licensing/).
- Ansible: 9.1/10 TrustRadius-style rating; G2 cites 366 reviews with enterprise skew (48.5%) [secondary](https://www.trustradius.com/compare-products/microsoft-system-center-vs-red-hat-ansible-automation-platform).
- Argo CD 24,160 vs Flux 8,407 GitHub stars (2026-09-15) — roughly a 3:1 adoption-mindshare ratio in the GitOps engine space [secondary](https://devtoollab.com/blog/best-gitops-tools).
- The IBM/HashiCorp $6.4B deal (closed Feb 2025) is the single largest IaC-vendor consolidation event of the cycle; its pricing aftershocks (RUM model, free-tier restructuring of 2026-03-31) are the main enterprise-budget storyline of 2026 [secondary](https://www.env0.com/blog/terraform-cloud).
- Community-verified workshop repos (e.g. platformrelay's OpenTofu workshop, claims re-verified 2026-08-25) show OpenTofu pinned in production lab baselines (1.10.x pins) even as 1.12.x is current — evidence of conservative enterprise upgrade cadence [secondary](https://github.com/platformrelay/opentofu-workshop/blob/HEAD/docs/claims-verification.md).
- **Gaps:** no verified 2026 market-share percentages (Ansible vs Puppet vs Chef; Terraform vs OpenTofu install base), no verified OpenTofu download counts, and no named enterprise case studies were captured in this pass. These are the highest-value follow-ups for a v2 of this file.

---

## I. Conflicts, ambiguities and gaps

**Conflicts / ambiguities:**
1. **HCP Terraform free tier (2026-03-31):** env0's early-2026 guide says the free tier was replaced by a $500 trial credit consumed at tier rates; later sources (HashiCorp blog via agentdeals, June 2026) say legacy free users were auto-migrated to an *enhanced* free tier (500 resources, unlimited users). Both may be true at different dates. Verify against the live pricing page before quoting. [secondary]
2. **OpenTofu "latest":** eosl.date (crawled Sept 2026) recommends 1.10.10 as latest; release trackers show 1.12.6 (2026-08-19). The eosl.date entry is stale. [secondary]
3. **AAP pricing:** aggregators ($5k/$10k/$14k yr "Tower" editions) vs Red Hat's official per-managed-node model — do not budget from aggregator figures. [secondary]
4. **Pulumi Individual tier:** "unlimited updates" (dev.to) vs "500 deployment minutes/month free" (Spacelift) — likely different meters (updates vs deployment minutes); not reconciled here. [secondary]
5. **HashiCorp acquisition close date:** "February 2025" (env0) vs "late 2024" (a knowledge-base audit note) — prefer February 2025 per the more specific source. [secondary]

**Gaps (not verified in this pass):**
- G1. Exact latest Terraform 1.x patch and 2026 minor-release contents.
- G2. Ansible Automation Platform official current version (2.5 vs 2.6 GA status) and official price book.
- G3. Event-Driven Ansible 2026 adoption figures.
- G4. Pulumi 3.x feature deltas vs 2.x.
- G5. Crossplane release line and CNCF maturity (v1 vs v2).
- G6. Kyverno version/graduation status; OPA/Gatekeeper 2026 releases.
- G7. Spacelift / env0 / Scalr current list pricing.
- G8. External Secrets Operator, SOPS, Sealed Secrets 2026 versions.
- G9. driftctl 2026 status.
- G10. Packer latest 2026 version.
- G11. Named enterprise IaC case studies with 2026 dates.

---

## J. Glossary

- **AAP** — Ansible Automation Platform (Red Hat commercial). **AWX** — its open-source upstream.
- **BSL 1.1** — Business Source License; source-available, bars competitive hosted/embedded use, converts to FOSS after the change date.
- **EDA** — Event-Driven Ansible (rulebook-driven automation).
- **EE** — Execution Environment (Ansible container image built with ansible-builder).
- **GitOps** — Git as the single source of truth for desired state; agents (Argo CD/Flux) converge the cluster.
- **HCP** — HashiCorp Cloud Platform (post-IBM-acquisition branding retained in 2026 sources).
- **RUM** — Resources Under Management; HCP Terraform's per-resource billing meter.
- **TACO** — Terraform Automation and Collaboration Software (Spacelift, env0, Scalr…).
- **OpenBao** — Linux-Foundation community fork of Vault (MPL 2.0).
- **ESO** — External Secrets Operator (Kubernetes).

---

