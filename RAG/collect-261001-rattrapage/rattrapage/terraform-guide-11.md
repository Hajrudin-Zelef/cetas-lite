---
id: collect-261001-rattrapage/rattrapage/terraform-guide-11
title: "Terraform — Guide complet : Infrastructure as Code en production"
domain: rattrapage
role: reference
task: reference
actors: ["AWS", "Meta"]
dates: []
keywords: ["aws", "cost", "open source"]
source: docs/RAG/collect-261001-rattrapage/terraform_guide.md
source_anchor: ""
source_lines: [2450, 2678]
sha256: 2ae320cb2f72cc721d13011f3c04ecc99bdf5d1c6ef6fd279b6dfba5d9ce9e39
---

# Terraform — Guide complet : Infrastructure as Code en production

```bash
# Exemple : alerte si un verrou de state dure trop longtemps (DynamoDB)
aws dynamodb scan --table-name terraform-state-locks --region eu-west-3 \
  --query 'Items[?attribute_exists(LockID)]' | jq length
```

> 💡 Ajoutez un **dashboard "IaC"** : dernier apply par env, résultat du drift check, version des providers. C'est ce que l'astreinte regarde en premier.

---

## 70. Tester son code : terraform test et check blocks

```hcl
# tests/web.tftest.hcl — tests natifs (Terraform ≥ 1.6)
mock_provider "proxmox" {}   # simule le provider, aucun appel réel

run "naming_convention" {
  command = plan
  variables {
    vm_count = 3
  }
  # Vérifie que les noms suivent web-XX
  assert {
    condition     = alltrue([for k in keys(proxmox_virtual_environment_vm.web) : can(regex("^web-[0-9]{2}$", k))])
    error_message = "Les VM doivent se nommer web-XX."
  }
}
```

```hcl
# check blocks : assertions évaluées à chaque plan/apply (Terraform ≥ 1.5)
check "vm_count_coherent" {
  assert {
    condition     = var.vm_count == length(local.vms)
    error_message = "Incohérence entre vm_count et la map des VM."
  }
}
```

```bash
terraform test          # exécute les .tftest.hcl
```

---

## 71. Documenter : terraform-docs

```bash
# Installation
curl -sL https://github.com/terraform-docs/terraform-docs/releases/latest/download/terraform-docs-v0.x.x-linux-amd64.tar.gz \
  | tar xz -C /tmp && sudo mv /tmp/terraform-docs /usr/local/bin/

# Génère la doc du module (variables, outputs) en Markdown
terraform-docs markdown table modules/vm/ > modules/vm/README.md
```

Extrait généré :

```markdown
## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| name | Nom de la VM | `string` | n/a | yes |
| cores | vCPU | `number` | `2` | no |
```

> 💡 En CI : `terraform-docs --check` échoue si le README n'est pas à jour → documentation toujours synchronisée.

---

## 72. Estimer les coûts : Infracost

Infracost estime le coût mensuel d'un plan (utile surtout pour le cloud public ; pour Proxmox on-prem, adaptez avec des coûts internes au kWh/To).

```bash
# Installation
curl -fsSL https://raw.githubusercontent.com/infracost/infracost/master/scripts/install.sh | sh

# Estimation sur un plan
terraform plan -out=tfplan
terraform show -json tfplan > plan.json
infracost breakdown --path plan.json
```

Pour l'on-prem, une approche simple : un `locals` de coût interne exposé en output :

```hcl
locals {
  # Coûts internes mensuels (exemples, à ajuster)
  cost_vcpu_gb_ram = 2.5   # €/mois par vCPU
  cost_gb_disk     = 0.08  # €/mois par Go Ceph
}
output "cout_mensuel_estime" {
  value = sum([
    for k, vm in proxmox_virtual_environment_vm.web :
    4 * local.cost_vcpu_gb_ram + 40 * local.cost_gb_disk
  ])
}
```

---

## 73. Conventions de nommage et style d'équipe

```hcl
# Fichiers : minuscules, par thème
# versions.tf, providers.tf, backend.tf, variables.tf, locals.tf,
# main.tf (ou vms.tf, network.tf), outputs.tf, checks.tf

# Ressources : snake_case, nom local explicite
resource "proxmox_virtual_environment_vm" "web_prod_01" { ... }  # ✅
resource "proxmox_virtual_environment_vm" "vm1" { ... }          # ❌

# Noms réels : <env>-<role>-<nn>
name = "prod-web-01"   # ✅  (tri, filtrage, inventaire)

# Tags : toujours env + managed-by + rôle
tags = ["terraform", "prod", "web"]

# Variables : noms explicites, description obligatoire
variable "proxmox_endpoint" { ... }   # ✅
variable "endpoint" { ... }           # ❌ trop vague
```

Règles d'équipe (à mettre dans `CONTRIBUTING.md`) :
- `terraform fmt` avant chaque commit (pre-commit hook).
- Pas de valeur en dur : tout ce qui varie par env → variable.
- Un `README.md` par stack expliquant : but, prérequis, `init/plan/apply`, environnements.

## 74. Checklist de mise en production

Avant le premier `apply` sur la prod :

**Code**
- [ ] `terraform fmt -check` OK, `terraform validate` OK
- [ ] `required_version` + `required_providers` figés, lock file versionné
- [ ] Aucun secret en clair (`gitleaks detect` OK)
- [ ] `prevent_destroy` sur les ressources critiques (BDD, NAS)
- [ ] Variables validées (`validation` blocks), descriptions renseignées

**State & backend**
- [ ] Backend distant configuré (S3 versionné + chiffrement + DynamoDB, ou équivalent)
- [ ] Sauvegarde du state testée (pull → restore sur un backend de test)
- [ ] Accès IAM restreints (équipe infra + rôle CI uniquement)

**Processus**
- [ ] Pipeline CI : fmt → validate → plan sur MR, apply manuel sur main
- [ ] Drift check nightly en place + alerte
- [ ] Procédure de reprise après sinistre écrite et testée
- [ ] `CONTRIBUTING.md` / conventions d'équipe diffusées
- [ ] Plan relu par un pair avant apply (jamais seul sur prod)

**Exploitation**
- [ ] Monitoring (Prometheus/pve-exporter) + alerting sur le parc
- [ ] Sauvegardes PBS configurées et répliquées hors site
- [ ] Documentation (`terraform-docs`) à jour

---

## 75. Pense-bête de poche (cheat sheet)

```bash
# ── Quotidien ──
terraform init                          # 1ère fois / nouveau module/provider
terraform plan -out=tfplan              # plan relu, sauvegardé
terraform apply tfplan                  # apply EXACTEMENT ce plan
terraform output -json | jq .           # voir les outputs

# ── Dépannage ──
terraform validate                      # cohérence du code
terraform plan -refresh-only            # voir la dérive sans toucher
terraform state list                    # que contient le state ?
terraform force-unlock <ID>             # libérer un verrou (ID du message)
TF_LOG=DEBUG terraform plan             # debug (purger les logs après)

# ── Recréer / importer ──
terraform plan -replace='<addr>' -out=tfplan   # recréer une ressource
terraform import '<addr>' '<id-provider>'      # importer l'existant

# ── Workspaces ──
terraform workspace new dev && terraform workspace select dev

# ── OpenTofu ──
tofu plan / tofu apply                  # mêmes commandes, binaire tofu
```

| Symbole du plan | Signification |
|---|---|
| `+` | création |
| `~` | modification en place |
| `-` | destruction |
| `-/+` | remplacement (destroy puis create) |
| `<=` | lecture (data source) |

---

## 76. Glossaire

| Terme | Définition |
|---|---|
| **IaC** | Infrastructure as Code : décrire l'infra dans des fichiers versionnés |
| **HCL** | HashiCorp Configuration Language, le langage de Terraform |
| **Provider** | Plugin traduisant le HCL en appels API (Proxmox, AWS…) |
| **Resource** | Objet géré par Terraform (VM, disque, DNS…) |
| **Data source** | Lecture d'un objet existant, sans le gérer |
| **State** | Fichier JSON mémorisant le mapping code ↔ objets réels |
| **Backend** | Où est stocké le state (local, S3, Terraform Cloud…) |
| **Plan** | Calcul des actions nécessaires, sans rien modifier |
| **Apply** | Exécution des actions du plan |
| **Drift** | Écart entre le réel et le code (modif hors Terraform) |
| **Taint** | Marquer une ressource comme à recréer |
| **Module** | Paquet réutilisable de configuration Terraform |
| **Workspace** | Instance isolée du state dans un même backend |
| **Lock** | Verrou du state distant empêchant deux apply concurrents |
| **Idempotence** | Propriété : réappliquer ne change rien |
| **for_each / count** | Meta-arguments de répétition de ressources |
| **Dynamic block** | Génération de sous-blocs par boucle |
| **Lifecycle** | Meta-bloc contrôlant création/destruction (prevent_destroy…) |
| **Provisioner** | Script exécuté à la création (à éviter) |
| **cloud-init** | Standard de configuration initiale des VM cloud |
| **OpenTofu** | Fork open source (MPL-2.0) de Terraform |
| **SOPS** | Outil de chiffrement de fichiers (avec age) |
| **TF_LOG** | Variable d'activation des logs de debug |

---

## 77. Quiz : 10 questions (énoncés)

