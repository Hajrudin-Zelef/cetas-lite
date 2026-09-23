# INDEX — Phase E2 — Ansible, Nornir, Terraform/OpenTofu & Python network automation libraries

Corpus `etape6-phasee2-ansible-nornir-terraform` · **16 fichiers** · 755 lignes source · ~9940 mots · partition exacte de `docs/RAG/etape6_phaseE2_ansible_nornir_terraform.md`.

## Mode d'emploi

1. Filtrer dans `manifest.json` (ou les tableaux ci-dessous) sur `domain`, `task`, `actors`, `dates` ou `keywords`.
2. Ouvrir 1 à 3 fichiers ciblés ; chaque fichier est une unité thématique auto-suffisante avec un en-tête YAML.
3. Pour un événement répété dans plusieurs sections, préférer le fichier marqué `canonical_for` (voir la table Événements canoniques).

## Domaines (dossiers → fichiers)

### `00-front-matter/`

| # | fichier | lignes source | rôle | tâche |
|---|---|---|---|---|
| 01 | [Phase E2 — Ansible, Nornir, Terraform/OpenTofu & Python network automation libraries](00-front-matter/overview.md) | 1–52 | reference | reference |
| 02 | [Wave 3 — AWX, Ansible Automation Platform & Event-Driven Ansible](00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md) | 53–107 | reference | reference |
| 03 | [Wave 6 — Python network libraries](00-front-matter/wave-6-python-network-libraries.md) | 108–155 | reference | reference |
| 04 | [Wave 9 — Ansible networking internals & validated patterns](00-front-matter/wave-9-ansible-networking-internals-validated-patterns.md) | 156–201 | reference | reference |
| 05 | [Wave 11 — Terraform/OpenTofu deep dive for networks](00-front-matter/wave-11-terraform-opentofu-deep-dive-for-networks.md) | 202–287 | reference | reference |

### `01-requirements-yml-pin-everything/`

| # | fichier | lignes source | rôle | tâche |
|---|---|---|---|---|
| 01 | [requirements.yml — pin everything](01-requirements-yml-pin-everything/overview.md) | 288–298 | deep-dive | reference |

### `02-playbooks-vlan-push-yml/`

| # | fichier | lignes source | rôle | tâche |
|---|---|---|---|---|
| 01 | [playbooks/vlan_push.yml](02-playbooks-vlan-push-yml/overview.md) | 299–339 | deep-dive | reference |

### `03-filter-to-one-site-for-canary-rollout/`

| # | fichier | lignes source | rôle | tâche |
|---|---|---|---|---|
| 01 | [Filter to one site for canary rollout](03-filter-to-one-site-for-canary-rollout/overview.md) | 340–359 | deep-dive | reference |

### `04-versions-pinned-binary-pinned-in-mise-toml-opentofu-1-12-6/`

| # | fichier | lignes source | rôle | tâche |
|---|---|---|---|---|
| 01 | [versions pinned; binary pinned in mise.toml (opentofu = "1.12.6")](04-versions-pinned-binary-pinned-in-mise-toml-opentofu-1-12-6/overview.md) | 360–378 | deep-dive | reference |

### `05-brownfield-adoption-import-existing-vlans-without-recreation/`

| # | fichier | lignes source | rôle | tâche |
|---|---|---|---|---|
| 01 | [Brownfield adoption: import existing VLANs without recreation](05-brownfield-adoption-import-existing-vlans-without-recreation/overview.md) | 379–389 | deep-dive | reference |

### `06-rulebooks-interface-flap-yml/`

| # | fichier | lignes source | rôle | tâche |
|---|---|---|---|---|
| 01 | [rulebooks/interface_flap.yml](06-rulebooks-interface-flap-yml/overview.md) | 390–415 | deep-dive | reference |

### `07-github-workflows-network-ci-yml-sketch/`

| # | fichier | lignes source | rôle | tâche |
|---|---|---|---|---|
| 01 | [.github/workflows/network-ci.yml (sketch)](07-github-workflows-network-ci-yml-sketch/overview.md) | 416–502 | deep-dive | reference |
| 02 | [Wave 17 — Vendor provider & collection catalog details](07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md) | 503–554 | deep-dive | reference |
| 03 | [Wave 19 — Glossary & concept map](07-github-workflows-network-ci-yml-sketch/wave-19-glossary-concept-map.md) | 555–607 | deep-dive | reference |
| 04 | [Wave 22 — Migration playbooks](07-github-workflows-network-ci-yml-sketch/wave-22-migration-playbooks.md) | 608–675 | deep-dive | reference |
| 05 | [Wave 25 — Key numbers at a glance (2026-09-22)](07-github-workflows-network-ci-yml-sketch/wave-25-key-numbers-at-a-glance-2026-09-22.md) | 676–755 | deep-dive | reference |

## Par tâche

- **reference** — [Phase E2 — Ansible, Nornir, Terraform/OpenTofu & Python network automation libraries](00-front-matter/overview.md), [Wave 3 — AWX, Ansible Automation Platform & Event-Driven Ansible](00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md), [Wave 6 — Python network libraries](00-front-matter/wave-6-python-network-libraries.md), [Wave 9 — Ansible networking internals & validated patterns](00-front-matter/wave-9-ansible-networking-internals-validated-patterns.md), [Wave 11 — Terraform/OpenTofu deep dive for networks](00-front-matter/wave-11-terraform-opentofu-deep-dive-for-networks.md), [requirements.yml — pin everything](01-requirements-yml-pin-everything/overview.md), [playbooks/vlan_push.yml](02-playbooks-vlan-push-yml/overview.md), [Filter to one site for canary rollout](03-filter-to-one-site-for-canary-rollout/overview.md), [versions pinned; binary pinned in mise.toml (opentofu = "1.12.6")](04-versions-pinned-binary-pinned-in-mise-toml-opentofu-1-12-6/overview.md), [Brownfield adoption: import existing VLANs without recreation](05-brownfield-adoption-import-existing-vlans-without-recreation/overview.md), [rulebooks/interface_flap.yml](06-rulebooks-interface-flap-yml/overview.md), [.github/workflows/network-ci.yml (sketch)](07-github-workflows-network-ci-yml-sketch/overview.md), [Wave 17 — Vendor provider & collection catalog details](07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md), [Wave 19 — Glossary & concept map](07-github-workflows-network-ci-yml-sketch/wave-19-glossary-concept-map.md), [Wave 22 — Migration playbooks](07-github-workflows-network-ci-yml-sketch/wave-22-migration-playbooks.md), [Wave 25 — Key numbers at a glance (2026-09-22)](07-github-workflows-network-ci-yml-sketch/wave-25-key-numbers-at-a-glance-2026-09-22.md)

## Par acteur

- **AWS** (3) — [00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md](00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md), [00-front-matter/wave-11-terraform-opentofu-deep-dive-for-networks.md](00-front-matter/wave-11-terraform-opentofu-deep-dive-for-networks.md), [07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md](07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md)
- **Meta** (1) — [00-front-matter/overview.md](00-front-matter/overview.md)

## Par date

- **2024-07** — [00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md](00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md)
- **2025-04** — [00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md](00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md)
- **2025-11-03** — [00-front-matter/overview.md](00-front-matter/overview.md)
- **2026-01-08** — [00-front-matter/overview.md](00-front-matter/overview.md), [07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md](07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md)
- **2026-03** — [00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md](00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md), [00-front-matter/wave-6-python-network-libraries.md](00-front-matter/wave-6-python-network-libraries.md)
- **2026-03-12** — [00-front-matter/overview.md](00-front-matter/overview.md), [00-front-matter/wave-6-python-network-libraries.md](00-front-matter/wave-6-python-network-libraries.md), [07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md](07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md)
- **2026-03-25** — [00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md](00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md), [07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md](07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md)
- **2026-04-16** — [00-front-matter/wave-6-python-network-libraries.md](00-front-matter/wave-6-python-network-libraries.md), [07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md](07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md)
- **2026-05** — [00-front-matter/overview.md](00-front-matter/overview.md), [00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md](00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md), [07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md](07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md)
- **2026-05-31** — [00-front-matter/overview.md](00-front-matter/overview.md), [07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md](07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md)
- **2026-06** — [00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md](00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md), [00-front-matter/wave-6-python-network-libraries.md](00-front-matter/wave-6-python-network-libraries.md)
- **2026-06-01** — [00-front-matter/overview.md](00-front-matter/overview.md), [07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md](07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md)
- **2026-06-12** — [00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md](00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md)
- **2026-06-13** — [00-front-matter/wave-6-python-network-libraries.md](00-front-matter/wave-6-python-network-libraries.md), [07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md](07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md)
- **2026-06-14** — [00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md](00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md), [07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md](07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md)
- **2026-06-22** — [00-front-matter/overview.md](00-front-matter/overview.md), [07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md](07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md)
- **2026-07** — [00-front-matter/overview.md](00-front-matter/overview.md), [00-front-matter/wave-9-ansible-networking-internals-validated-patterns.md](00-front-matter/wave-9-ansible-networking-internals-validated-patterns.md), [07-github-workflows-network-ci-yml-sketch/wave-25-key-numbers-at-a-glance-2026-09-22.md](07-github-workflows-network-ci-yml-sketch/wave-25-key-numbers-at-a-glance-2026-09-22.md)
- **2026-07-03** — [00-front-matter/wave-6-python-network-libraries.md](00-front-matter/wave-6-python-network-libraries.md), [07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md](07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md)
- **2026-07-23** — [00-front-matter/overview.md](00-front-matter/overview.md), [07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md](07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md)
- **2026-07-30** — [00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md](00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md), [07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md](07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md)
- **2026-08-19** — [00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md](00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md), [07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md](07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md), [07-github-workflows-network-ci-yml-sketch/wave-25-key-numbers-at-a-glance-2026-09-22.md](07-github-workflows-network-ci-yml-sketch/wave-25-key-numbers-at-a-glance-2026-09-22.md)
- **2026-08-21** — [00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md](00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md)
- **2026-08-30** — [00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md](00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md)
- **2026-09-08** — [00-front-matter/overview.md](00-front-matter/overview.md), [07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md](07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md)
- **2026-09-10** — [00-front-matter/overview.md](00-front-matter/overview.md), [07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md](07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md)
- **2026-09-22** — [00-front-matter/overview.md](00-front-matter/overview.md), [00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md](00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md), [00-front-matter/wave-6-python-network-libraries.md](00-front-matter/wave-6-python-network-libraries.md), [07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md](07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md), [07-github-workflows-network-ci-yml-sketch/wave-25-key-numbers-at-a-glance-2026-09-22.md](07-github-workflows-network-ci-yml-sketch/wave-25-key-numbers-at-a-glance-2026-09-22.md)
- **2026-11-30** — [00-front-matter/overview.md](00-front-matter/overview.md), [00-front-matter/wave-6-python-network-libraries.md](00-front-matter/wave-6-python-network-libraries.md), [07-github-workflows-network-ci-yml-sketch/wave-25-key-numbers-at-a-glance-2026-09-22.md](07-github-workflows-network-ci-yml-sketch/wave-25-key-numbers-at-a-glance-2026-09-22.md)
- **2027-05** — [00-front-matter/overview.md](00-front-matter/overview.md), [07-github-workflows-network-ci-yml-sketch/wave-25-key-numbers-at-a-glance-2026-09-22.md](07-github-workflows-network-ci-yml-sketch/wave-25-key-numbers-at-a-glance-2026-09-22.md)
- **2027-05-31** — [00-front-matter/overview.md](00-front-matter/overview.md)
- **2027-11-30** — [00-front-matter/overview.md](00-front-matter/overview.md)
- **2028-03** — [00-front-matter/overview.md](00-front-matter/overview.md), [00-front-matter/wave-9-ansible-networking-internals-validated-patterns.md](00-front-matter/wave-9-ansible-networking-internals-validated-patterns.md), [07-github-workflows-network-ci-yml-sketch/wave-22-migration-playbooks.md](07-github-workflows-network-ci-yml-sketch/wave-22-migration-playbooks.md)
- **2028-04-01** — [00-front-matter/overview.md](00-front-matter/overview.md), [07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md](07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md), [07-github-workflows-network-ci-yml-sketch/wave-22-migration-playbooks.md](07-github-workflows-network-ci-yml-sketch/wave-22-migration-playbooks.md), [07-github-workflows-network-ci-yml-sketch/wave-25-key-numbers-at-a-glance-2026-09-22.md](07-github-workflows-network-ci-yml-sketch/wave-25-key-numbers-at-a-glance-2026-09-22.md)
- **2028-05** — [00-front-matter/overview.md](00-front-matter/overview.md)

## Carte de couverture (lignes source)

| plage | fichier |
|---|---|
| 1–52 | etape6-phasee2-ansible-nornir-terraform/00-front-matter/overview.md |
| 53–107 | etape6-phasee2-ansible-nornir-terraform/00-front-matter/wave-3-awx-ansible-automation-platform-event-driven-ansible.md |
| 108–155 | etape6-phasee2-ansible-nornir-terraform/00-front-matter/wave-6-python-network-libraries.md |
| 156–201 | etape6-phasee2-ansible-nornir-terraform/00-front-matter/wave-9-ansible-networking-internals-validated-patterns.md |
| 202–287 | etape6-phasee2-ansible-nornir-terraform/00-front-matter/wave-11-terraform-opentofu-deep-dive-for-networks.md |
| 288–298 | etape6-phasee2-ansible-nornir-terraform/01-requirements-yml-pin-everything/overview.md |
| 299–339 | etape6-phasee2-ansible-nornir-terraform/02-playbooks-vlan-push-yml/overview.md |
| 340–359 | etape6-phasee2-ansible-nornir-terraform/03-filter-to-one-site-for-canary-rollout/overview.md |
| 360–378 | etape6-phasee2-ansible-nornir-terraform/04-versions-pinned-binary-pinned-in-mise-toml-opentofu-1-12-6/overview.md |
| 379–389 | etape6-phasee2-ansible-nornir-terraform/05-brownfield-adoption-import-existing-vlans-without-recreation/overview.md |
| 390–415 | etape6-phasee2-ansible-nornir-terraform/06-rulebooks-interface-flap-yml/overview.md |
| 416–502 | etape6-phasee2-ansible-nornir-terraform/07-github-workflows-network-ci-yml-sketch/overview.md |
| 503–554 | etape6-phasee2-ansible-nornir-terraform/07-github-workflows-network-ci-yml-sketch/wave-17-vendor-provider-collection-catalog-details.md |
| 555–607 | etape6-phasee2-ansible-nornir-terraform/07-github-workflows-network-ci-yml-sketch/wave-19-glossary-concept-map.md |
| 608–675 | etape6-phasee2-ansible-nornir-terraform/07-github-workflows-network-ci-yml-sketch/wave-22-migration-playbooks.md |
| 676–755 | etape6-phasee2-ansible-nornir-terraform/07-github-workflows-network-ci-yml-sketch/wave-25-key-numbers-at-a-glance-2026-09-22.md |

