---
id: collect-261001-rattrapage/rattrapage/terraform-guide-6
title: "Terraform — Guide complet : Infrastructure as Code en production"
domain: rattrapage
role: reference
task: reference
actors: ["AWS"]
dates: ["2026-09-26"]
keywords: ["attention", "aws"]
source: docs/RAG/collect-261001-rattrapage/terraform_guide.md
source_anchor: ""
source_lines: [1270, 1501]
sha256: 565b44401a8661319b89199c9a4f57b63bfd365efa96d7ef8a892845d34b12c8
---

# Terraform — Guide complet : Infrastructure as Code en production

```hcl
# Exemple : nom unique par génération pour permettre la coexistence
resource "proxmox_virtual_environment_vm" "web" {
  name = "web-${random_id.gen.hex}"
  lifecycle { create_before_destroy = true }
}
resource "random_id" "gen" {
  byte_length = 2
  keepers = { config_version = var.config_version }  # change => nouvelle VM
}
```

---

## 36. ignore_changes : ignorer les dérives voulues

Quand un attribut est modifié **hors Terraform** volontairement (ex. tags posés par une autre équipe, nombre de vCPU ajusté à chaud), on évite que chaque plan propose de le réécrire.

```hcl
resource "proxmox_virtual_environment_vm" "web" {
  # ...
  lifecycle {
    ignore_changes = [
      tags,            # les tags gérés hors Terraform sont ignorés
      cpu[0].cores,    # ajustement CPU à chaud ignoré (adressage d'un sous-bloc)
    ]
  }
}
```

> ⚠️ À utiliser avec parcimonie : chaque `ignore_changes` est un angle mort du drift detection (section 56). Documentez **pourquoi** dans un commentaire.

## 37. replace_triggered_by : forcer le remplacement

Quand une ressource dépend d'un fichier externe (cloud-init, script) que Terraform ne "voit" pas comme un changement de la ressource elle-même :

```hcl
resource "proxmox_virtual_environment_file" "cloud_config" {
  for_each     = local.vms
  content_type = "snippets"
  datastore_id = "local"
  node_name    = each.value.node
  source_raw {
    file_name = "${each.key}-cloud-config.yaml"
    data      = templatefile("${path.module}/templates/cloud-init.tftpl", {
      hostname = each.key
      # ...
    })
  }
}

resource "proxmox_virtual_environment_vm" "nodes" {
  for_each = local.vms
  # ...
  lifecycle {
    # Si le cloud-init change, recréer la VM (le provider bpg ne le fait pas seul)
    replace_triggered_by = [
      proxmox_virtual_environment_file.cloud_config[each.key].id
    ]
  }
}
```

> 💡 Alternative moins brutale : versionner le nom du fichier (`cloud-config-v2.yaml`) pour tracer les générations.

---

## 38. taint et -replace : marquer une ressource à recréer

Une VM est corrompue (disque endommagé, OS vérolé) mais sa config Terraform est bonne : on la marque à recréer.

```bash
# Ancienne méthode (deprecated mais encore vue) :
terraform taint 'proxmox_virtual_environment_vm.nodes["web-01"]'

# Méthode moderne (planifie explicitement le remplacement) :
terraform plan -replace='proxmox_virtual_environment_vm.nodes["web-01"]' -out=tfplan
terraform apply tfplan

# Annuler un taint :
terraform untaint 'proxmox_virtual_environment_vm.nodes["web-01"]'
```

Le plan affiche alors `-/+` (remplacement) pour cette seule ressource.

---

## 39. Refresh et plan -refresh-only

```bash
# Synchronise le state avec la réalité (sans modifier l'infra)
terraform refresh

# Aperçu des dérives SANS proposer de corrections (lecture seule)
terraform plan -refresh-only -out=refresh.plan
```

Cas d'usage : quelqu'un a modifié une VM dans l'UI Proxmox → `plan -refresh-only` montre ce qui a dérivé, sans proposer de tout écraser. C'est la base du **drift detection** (section 56).

> ⚠️ `terraform refresh` seul modifie le state : préférez `plan -refresh-only` + `apply` du refresh plan pour garder une trace.

---

## 40. Le state : pourquoi il existe

Le **state** (`terraform.tfstate`, JSON) est la mémoire de Terraform : il mappe chaque `resource` du code vers l'objet réel (VMID Proxmox, adresse IP…).

```json
{
  "version": 4,
  "resources": [
    {
      "type": "proxmox_virtual_environment_vm",
      "name": "web_01",
      "instances": [{ "attributes": { "vm_id": 101, "name": "web-01" } }]
    }
  ]
}
```

Sans state, Terraform ne saurait pas que `web_01` = VMID 101 : il voudrait tout recréer. Le state contient aussi les **outputs** et les **dépendances**.

> 🚨 Le state peut contenir des **secrets** (mots de passe en clair dans les attributs). Il doit être **chiffré au repos, accès restreint, jamais dans Git** (sauf repo privé avec git-crypt — déconseillé).

---

## 41. State local vs state distant

| | Local (`terraform.tfstate`) | Distant (S3, Terraform Cloud…) |
|---|---|---|
| Stockage | disque du poste | service partagé |
| Travail en équipe | ❌ (conflits, écrasements) | ✅ |
| Verrouillage | ❌ (risque de corruption si 2 apply) | ✅ (state locking) |
| Secrets | en clair sur le poste | chiffrables au repos |
| Sauvegarde | manuelle | versioning natif |

**Règle** : state local = POC solo sur son poste. Dès qu'on est 2 ou qu'on automatise → **backend distant obligatoire**.

---

## 42. Backend S3 + verrouillage DynamoDB

Le classique : state dans un bucket S3 (versionné, chiffré), verrou dans DynamoDB.

```hcl
# backend.tf
terraform {
  backend "s3" {
    bucket         = "infra-terraform-state-prod"
    key            = "proxmox/prod/terraform.tfstate"
    region         = "eu-west-3"
    encrypt        = true
    dynamodb_table = "terraform-state-locks"
    # use_lockfile = true   # alternative moderne (S3 natif, sans DynamoDB)
  }
}
```

```bash
# Création du bucket + table (une fois, à la main ou via une stack bootstrap)
aws s3api create-bucket --bucket infra-terraform-state-prod --region eu-west-3 \
  --create-bucket-configuration LocationConstraint=eu-west-3
aws s3api put-bucket-versioning --bucket infra-terraform-state-prod \
  --versioning-configuration Status=Enabled
aws s3api put-bucket-encryption --bucket infra-terraform-state-prod \
  --server-side-encryption-configuration \
  '{"Rules":[{"ApplyServerSideEncryptionByDefault":{"SSEAlgorithm":"AES256"}}]}'
aws dynamodb create-table --table-name terraform-state-locks \
  --attribute-definitions AttributeName=LockID,AttributeType=S \
  --key-schema AttributeName=LockID,KeyType=HASH \
  --billing-mode PAY_PER_REQUEST --region eu-west-3
```

**Alternatives on-prem** (sans AWS) :
- **Backend HTTP** maison (simple serveur avec verrou) ;
- **Terraform Cloud / HCP Terraform** (section 43) ;
- **OpenTofu** : chiffrement natif du state (`tofu` + `key_provider`, section 63) ;
- MinIO (compatible S3) auto-hébergé : même config, `endpoints = { s3 = "https://minio.lan" }`.

> ⚠️ Changer de backend = `terraform init -migrate-state` (Terraform propose de copier l'ancien state). Toujours sauvegarder avant.

---

## 43. Backend Terraform Cloud / HCP Terraform

```hcl
terraform {
  cloud {
    organization = "mon-entreprise"
    workspaces {
      name = "proxmox-prod"
    }
  }
}
```

Avantages : state hébergé + verrouillé + versionné, historique des runs, revue de plan dans l'UI, variables sensibles chiffrées, intégration VCS (plan auto sur PR).
Inconvénients : dépendance à un SaaS externe, offre gratuite limitée.

```bash
terraform login          # authentification (token stocké en clair dans ~/.terraform.d/credentials.tfrc.json !)
terraform init           # migre le state local vers le cloud
```

> 💡 Pour un usage 100 % on-prem, préférez S3/MinIO + DynamoDB ou le backend HTTP.

---

## 44. Sauvegarde et protection du state

Checklist de protection du state distant :

- [ ] **Versioning** activé sur le bucket (restauration d'une version antérieure en 1 clic)
- [ ] **Chiffrement au repos** (SSE-S3 ou SSE-KMS)
- [ ] **Accès restreint** : IAM limité à l'équipe infra + rôle CI (pas de `s3:*` large)
- [ ] **MFA Delete** sur le bucket (optionnel mais fort)
- [ ] **Sauvegarde externe** : `terraform state pull > backup-$(date +%F).tfstate` chiffré (gpg) vers un stockage séparé
- [ ] **Jamais dans Git** en clair

```bash
# Sauvegarde manuelle chiffrée
terraform state pull | gpg --encrypt --recipient infra@entreprise.lan \
  > /srv/backups/terraform/prod-$(date +%F).tfstate.gpg

# Restauration d'urgence (ATTENTION : écrase le state distant)
gpg --decrypt prod-2026-09-26.tfstate.gpg | terraform state push -
```

> 🚨 `state push` écrase sans verrou collaboratif : à n'utiliser qu'en dernier recours, équipe prévenue.

---

