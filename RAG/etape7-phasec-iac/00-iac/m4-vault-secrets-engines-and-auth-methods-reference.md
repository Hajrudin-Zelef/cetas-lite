---
id: etape7-phasec-iac/00-iac/m4-vault-secrets-engines-and-auth-methods-reference
title: "M4. Vault secrets engines and auth methods (reference)"
domain: step-7-phase-c-iac-platform-automation
role: deep-dive
task: reference
actors: ["AWS"]
dates: ["2026-09-15", "2026-09-22"]
keywords: ["agent", "apache", "aws", "distribution", "governance", "license", "pricing"]
source: docs/RAG/etape7_phaseC_iac.md
source_anchor: ""
source_lines: [582, 663]
section: "Step 7 Phase C — IaC & Platform Automation"
sha256: e4deeb536461e4f984b00766e042b9bba147b91d3d842ff749b764a9f4d957be
---

# M4. Vault secrets engines and auth methods (reference)

### M4. Vault secrets engines and auth methods (reference)

- **Secrets engines**: KV v2 (versioned), database (dynamic creds), PKI (internal CA), transit (encryption-as-a-service), SSH (signed keys/OTP), AWS/Azure/GCP (dynamic cloud creds), TOTP, KMIP [unverified — documented engine list].
- **Auth methods**: userpass, LDAP, OIDC/JWT, Kubernetes, AWS, Azure, GCP, AppRole, TLS certificates, GitHub [unverified — documented list].
- **HA**: integrated Raft storage; performance standby nodes and replication (DR + performance) on Enterprise/Standard tiers [secondary](https://infisical.com/blog/hashicorp-vault-pricing).
- **OpenBao parity goal**: drop-in API compatibility for the OSS Vault surface, MPL 2.0 governance [secondary](https://tech-insider.org/opentofu-vs-terraform-vs-pulumi-2026/).

### M5. cloud-init example (user-data pattern)

The canonical first-boot pattern for IaC-managed VMs (indented block, not executed):

    #cloud-config
    hostname: web-01
    manage_etc_hosts: true
    users:
      - name: deploy
        sudo: ALL=(ALL) NOPASSWD:ALL
        ssh_authorized_keys:
          - ssh-ed25519 AAAA...
    packages:
      - qemu-guest-agent
    runcmd:
      - [ systemctl, enable, --now, qemu-guest-agent ]
    final_message: "cloud-init done in $UPTIME seconds"

Pattern notes: keep user-data declarative; secrets go through Vault/ESO, never user-data; use `phone_home` or Ansible callbacks for boot reporting [official](https://docs.cloud-init.io/_/downloads/en/latest/pdf/).

### M6. EDA rulebook sketch (pattern)

    ---
    - name: Remediate full disk
      hosts: all
      sources:
        - ansible.eda.webhook:
            host: 0.0.0.0
            port: 5000
      rules:
        - name: disk usage over 90 percent
          condition: event.payload.usage > 90
          action:
            run_playbook:
              name: ops/cleanup_disk.yml

Pattern: event source → rulebook condition → playbook action; EDA server 1.2.12 is the 2.6-era supported component [secondary](https://github.com/fitbeard/automation-platform/blob/HEAD/README.md).

### M7. OpenTofu CLI quick reference (2026)

- `tofu init / plan / apply / destroy` — same verbs as Terraform; `-exclude` (1.9+) complements `-target` [secondary](https://tech-insider.org/opentofu-vs-terraform-vs-pulumi-2026/).
- `tofu test` — native test framework (`.tftest.hcl`); `tofu fmt`, `tofu validate` as in Terraform [unverified].
- Backend config supports early variable evaluation (1.8+) — variables/locals in `backend` blocks [secondary](https://tech-insider.org/opentofu-vs-terraform-vs-pulumi-2026/).
- Registry: OpenTofu registry mirrors Terraform provider namespace for compatible providers; 1.10+ adds OCI-registry distribution [secondary](https://tech-insider.org/opentofu-vs-terraform-vs-pulumi-2026/).

### M8. Argo CD vs Flux decision table (2026)

| Criterion | Argo CD (v3.5.x) | Flux (v2.9.x) |
|---|---|---|
| UI | Yes (resource tree) | No (CLI) |
| Architecture | Central server + Redis | In-cluster controllers |
| Image automation | Via Image Updater companion | Native controllers |
| Helm depth | Good | Stronger HelmRelease |
| Stars (2026-09-15) | 24,160 | 8,407 |
| License | Apache 2.0 | Apache 2.0 |
| Paid support | Akuity ($495/mo cited) | ControlPlane |
| Choose when | Devs need self-service UI | Small terminal-native platform team |

[secondary](https://devtoollab.com/blog/best-gitops-tools)

### M9. IaC security checklist (2026 baseline)

1. Encrypt state at rest: OpenTofu 1.7+ client-side encryption or backend SSE-KMS; never commit state to Git [secondary](https://tech-insider.org/opentofu-vs-terraform-vs-pulumi-2026/).
2. Ephemeral credentials: OpenTofu 1.11+ ephemeral values, Vault dynamic secrets, ESO short-lived sync [secondary](https://tech-insider.org/opentofu-vs-terraform-vs-pulumi-2026/).
3. Plan-time policy: Sentinel/OPA run tasks or conftest in CI [secondary](https://medium.com/@terraform-techie/best-infrastructure-as-code-governance-platforms-for-enterprises-in-2026-41fec153088e).
4. Admission policy: Kyverno/Gatekeeper on clusters fed by Argo CD/Flux [secondary](https://devtoollab.com/blog/best-gitops-tools).
5. Drift detection: scheduled plans + Argo CD sync status [secondary](https://www.theregister.com/tag/hashicorp).
6. Secrets hygiene: SOPS/ESO/Sealed Secrets; no plaintext in repos or user-data [unverified — established practice].
7. Lifecycle guards: OpenTofu 1.12 dynamic `prevent_destroy` for production data stores [secondary](https://github.com/christosgalano/christosgalano.github.io/blob/HEAD/_drafts/opentofu-vs-terraform-2026.md).
8. Supply chain: signed Argo CD images (cosign, SLSA 3); pinned collection/provider versions [official](https://github.com/zgfh/zgfh.github.io/blob/HEAD/content/docs/cncf/project/argo-cd/releasenote/argo-cd_v3.5.0_release_note.md).

*End of Step 7 Phase C file. 2026-09-22 cutoff. Single writer; no other workspace files modified.*

---

