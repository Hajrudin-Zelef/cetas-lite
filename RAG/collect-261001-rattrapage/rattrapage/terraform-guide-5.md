---
id: collect-261001-rattrapage/rattrapage/terraform-guide-5
title: "Terraform — Guide complet : Infrastructure as Code en production"
domain: rattrapage
role: reference
task: reference
actors: ["Meta"]
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-rattrapage/terraform_guide.md
source_anchor: ""
source_lines: [1008, 1269]
sha256: f225a707ef294e80d17a71de5596d4c6b1bce2f04eee74449c148bad3980a2ad
---

# Terraform — Guide complet : Infrastructure as Code en production

- Registre officiel : `registry.terraform.io` (Terraform) — OpenTofu sait aussi l'utiliser.
- Les binaires sont vérifiés par signature (`.terraform.lock.hcl` fige les hashes).
- **Ne jamais** commiter `.terraform/` (binaires lourds) — voir `.gitignore` section 6.

---

## 28. Version pinning : figer les versions

Contraintes de version (opérateurs) :

| Contrainte | Signification |
|---|---|
| `= 0.66.0` | exactement cette version |
| `~> 0.66` | `>= 0.66.0, < 0.67.0` (recommandé : correctifs + mineurs compatibles) |
| `>= 0.66, < 0.70` | intervalle explicite |
| `>= 1.9.0, < 2.0.0` | pour `required_version` de Terraform lui-même |

```hcl
terraform {
  required_version = ">= 1.9.0, < 1.11.0"
  required_providers {
    proxmox = {
      source  = "bpg/proxmox"
      version = "~> 0.66"     # fige la branche 0.66.x
    }
  }
}
```

`.terraform.lock.hcl` (versionné en Git !) fige le **hash exact** du binaire :

```bash
terraform init -upgrade   # autorise la montée de version dans la contrainte
terraform providers lock  # régénère le lock file
```

> 🚨 En prod : contrainte `~>` + lock file versionné + `init` sans `-upgrade` en CI. Une montée de provider ne se fait que volontairement, avec relecture du changelog.

---

## 29. Le provider Proxmox bpg : installation et configuration

Le provider communautaire `bpg/proxmox` est la référence pour Proxmox VE (API native, pas de dépendance SSH).

```hcl
# versions.tf
terraform {
  required_version = ">= 1.9.0, < 1.11.0"
  required_providers {
    proxmox = {
      source  = "bpg/proxmox"
      version = "~> 0.66"
    }
  }
}
```

```hcl
# providers.tf
variable "proxmox_endpoint" {
  description = "URL API Proxmox, ex. https://pve1.lan:8006/"
  type        = string
}
variable "proxmox_api_token" {
  description = "Token API 'USER@REALM!TOKENID=SECRET'"
  type        = string
  sensitive   = true
}

provider "proxmox" {
  endpoint  = var.proxmox_endpoint
  api_token = var.proxmox_api_token
  insecure  = var.proxmox_insecure_tls  # true si CA interne/autosignée
  # ssh { ... }  # optionnel : uniquement pour certaines opérations (ex. upload ISO)
  tmp_dir   = "/var/tmp"
}
```

**Créer le token API côté Proxmox** (privilège minimal) :

```bash
# Sur le nœud Proxmox, en root :
pveum user add terraform@pve --comment "Terraform automation"
pveum usertoken add terraform@pve tf-token --privsep 0
# Noter : terraform@pve!tf-token=<SECRET>

# Rôle minimal (ajuster selon besoins : VM.Allocate, VM.Clone, Datastore.AllocateSpace...)
pveum role add TerraformRole -privs "VM.Allocate VM.Clone VM.Config.CDROM VM.Config.CPU VM.Config.Disk VM.Config.HWType VM.Config.Memory VM.Config.Network VM.Config.Options VM.Audit VM.PowerMgmt Datastore.AllocateSpace Datastore.Audit Sys.Audit"
pveum acl modify / -user terraform@pve -role TerraformRole
```

> 💡 Le token s'injecte via `TF_VAR_proxmox_api_token` ou un `.tfvars` **hors Git** (sections 54-55). Jamais en clair dans le code.

---

## 30. Les ressources : anatomie d'un bloc resource

```hcl
resource "proxmox_virtual_environment_vm" "web_01" {
  # ── Identité ──
  name      = "web-01"
  node_name = "pve1"        # nœud d'hébergement
  vm_id     = 101           # VMID Proxmox (unique par cluster)
  tags      = ["terraform", "prod", "web"]

  # ── Matériel ──
  cpu {
    cores = 4
    type  = "host"          # ou x86-64-v2-AES pour la portabilité
  }
  memory {
    dedicated = 8192        # Mo
  }
  disk {
    datastore_id = "ceph-vm"
    file_format  = "raw"
    interface    = "scsi0"
    size         = 40        # Go
  }
  network_device {
    bridge = "vmbr0"
    vlan_id = 10
  }

  # ── Démarrage ──
  started = true
  on_boot = true

  # ── Cloud-init ──
  initialization {
    datastore_id = "local-lvm"
    ip_config {
      ipv4 {
        address = "10.10.0.11/24"
        gateway = "10.10.0.1"
      }
    }
  }
}
```

Attributs **calculés** (connus après création) : `ipv4_addresses`, `vm_id` (si auto), etc. Ils sont utilisables dans d'autres ressources et outputs — c'est le chaînage.

---

## 31. Dépendances implicites et graphe

Terraform construit un **graphe de dépendances** à partir des références entre ressources : si B utilise un attribut de A, B est créé **après** A, et détruit **avant** A.

```hcl
# Le snippet cloud-init doit exister AVANT la VM qui le référence :
# dépendance implicite via proxmox_virtual_environment_file.cloud_config[...].id
resource "proxmox_virtual_environment_vm" "web" {
  for_each = local.vms
  # ...
  initialization {
    user_data_file_id = proxmox_virtual_environment_file.cloud_config[each.key].id
  }
}
```

Visualiser le graphe :

```bash
terraform graph | dot -Tpng > graph.png   # nécessite graphviz
terraform graph -type=plan | dot -Tsvg > plan.svg
```

> 💡 95 % des dépendances doivent être **implicites** (par référence). C'est plus lisible et moins fragile que `depends_on`.

---

## 32. depends_on : dépendances explicites

À n'utiliser que quand il n'y a **aucune référence d'attribut** mais un ordre requis (ex. attendre qu'un script côté Proxmox ait préparé un stockage).

```hcl
resource "proxmox_virtual_environment_vm" "db" {
  name      = "db-01"
  node_name = "pve2"
  # ...

  # N'a de sens QUE sans référence directe entre les deux ressources
  depends_on = [
    proxmox_virtual_environment_vm.nas,  # le NAS doit exister avant la BDD
  ]
}
```

> ⚠️ `depends_on` crée un ordre strict mais **n'exporte aucune donnée**. Si vous avez besoin d'un attribut de l'autre ressource, utilisez une référence directe (dépendance implicite) plutôt que `depends_on`.

---

## 33. Le bloc lifecycle : meta-arguments

```hcl
resource "proxmox_virtual_environment_vm" "db" {
  # ...
  lifecycle {
    prevent_destroy       = true   # §34
    create_before_destroy = true   # §35
    ignore_changes        = [tags]  # §36
    replace_triggered_by  = [proxmox_virtual_environment_file.cloud_config["db-01"].id]  # §37
  }
}
```

| Meta-argument | Effet |
|---|---|
| `prevent_destroy = true` | `apply`/`destroy` échoue si la ressource doit être détruite |
| `create_before_destroy = true` | crée le remplaçant avant de détruire l'ancien |
| `ignore_changes = [...]` | ignore certaines dérives au plan |
| `replace_triggered_by = [...]` | force le remplacement quand la référence change |

---

## 34. prevent_destroy : le garde-fou anti-catastrophe

```hcl
resource "proxmox_virtual_environment_vm" "db_prod" {
  name      = "db-prod-01"
  node_name = "pve2"
  # ...

  lifecycle {
    prevent_destroy = true
  }
}
```

Si un plan veut détruire cette VM :

```
Error: Instance cannot be destroyed
  on main.tf line 42:
  42: resource "proxmox_virtual_environment_vm" "db_prod" {
Resource ... has lifecycle.prevent_destroy set, but the plan calls for this
resource to be destroyed.
```

**À mettre sur** : BDD de prod, NAS, tout ce qui est difficilement reconstructible.
**Retrait volontaire** : enlever le bloc, `apply`, puis détruire — jamais en urgence, toujours relu.

---

## 35. create_before_destroy : zéro interruption

Par défaut, un remplacement = détruire puis recréer (**coupure**). Avec `create_before_destroy`, Terraform crée d'abord le remplaçant.

```hcl
resource "proxmox_virtual_environment_vm" "web" {
  # ...
  lifecycle {
    create_before_destroy = true
  }
}
```

⚠️ **Contraintes** :
- Le nom/ID doit pouvoir coexister temporairement : avec Proxmox, deux VM ne peuvent pas partager le même `vm_id` ni le même nom → prévoyez un nom avec suffixe ou un `vm_id` dynamique.
- Doublonne temporairement les ressources (capacité à prévoir).

