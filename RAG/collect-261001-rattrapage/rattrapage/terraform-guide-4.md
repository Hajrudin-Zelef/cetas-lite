---
id: collect-261001-rattrapage/rattrapage/terraform-guide-4
title: "Terraform — Guide complet : Infrastructure as Code en production"
domain: rattrapage
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws", "memory"]
source: docs/RAG/collect-261001-rattrapage/terraform_guide.md
source_anchor: ""
source_lines: [714, 1007]
sha256: c42bb4ad105439525ee30c46a9dfc59a75ac8582035de02116671b578b28a4b3
---

# Terraform — Guide complet : Infrastructure as Code en production

```hcl
# providers.tf — configuration du provider bpg/proxmox
provider "proxmox" {
  endpoint  = var.proxmox_endpoint   # ex. https://pve1.exemple.lan:8006/
  api_token = var.proxmox_api_token # format "USER@REALM!TOKENID=SECRET" (voir §54)
  insecure  = false                 # true seulement avec CA autosignée interne
}
```

```hcl
# Lire le template Debian 12 existant (ne le crée pas, ne le modifie pas)
data "proxmox_virtual_environment_vms" "templates" {
  # selon la version du provider ; alternative : référencer un template connu
}

# Exemple générique avec un autre provider (concept identique) :
# data "proxmox_virtual_environment_datastores" ...
```

Exemple concret et portable avec le provider bpg :

```hcl
# On suppose un template cloud-init "debian-12-cloudinit" déjà présent (VMID 9000)
locals {
  template_vm_id = 9000
}

resource "proxmox_virtual_environment_vm" "web" {
  name      = "web-01"
  node_name = "pve1"
  vm_id     = 101

  clone {
    vm_id = local.template_vm_id   # clone depuis le template existant
    full  = true
  }
  # ...
}
```

> 💡 Règle : **data source = lecture**, **resource = gestion**. Si Terraform doit pouvoir le détruire, c'est une `resource`.

---

## 21. count : répétition simple

```hcl
resource "proxmox_virtual_environment_vm" "web" {
  count     = var.vm_count          # ex. 3
  name      = format("web-%02d", count.index + 1)  # web-01, web-02, web-03
  node_name = "pve1"
  vm_id     = 101 + count.index     # 101, 102, 103
  # ...
}
```

Références : `proxmox_virtual_environment_vm.web[0]`, `[1]`, `[2]`.

⚠️ **Limite de `count`** : les instances sont indexées par position. Supprimer l'élément du milieu décale tout :

```
web[0], web[1], web[2]  →  on passe count de 3 à 2
Terraform détruit web[2]... mais si on voulait retirer web[1],
c'est web[2] qui est détruit et web[1] est "modifié" en ex-web[2].
```

> 💡 Pour des VM nommées/identifiées durablement, préférez **`for_each`** (section 22).

---

## 22. for_each : répétition robuste par clé

```hcl
locals {
  vms = {
    "web-01" = { node = "pve1", cores = 4, memory = 8192, ip = "10.10.0.11" }
    "web-02" = { node = "pve1", cores = 4, memory = 8192, ip = "10.10.0.12" }
    "db-01"  = { node = "pve2", cores = 8, memory = 16384, ip = "10.10.0.21" }
  }
}

resource "proxmox_virtual_environment_vm" "nodes" {
  for_each  = local.vms
  name      = each.key               # "web-01", "web-02", "db-01"
  node_name = each.value.node
  # ...
  initialization {
    ip_config {
      ipv4 {
        address = "${each.value.ip}/24"
        gateway = "10.10.0.1"
      }
    }
  }
}
```

Références : `proxmox_virtual_environment_vm.nodes["web-01"]`.

| | `count` | `for_each` |
|---|---|---|
| Clé d'instance | index numérique | clé de map / valeur de set |
| Suppression d'un élément | décale les suivants (risque) | supprime uniquement cet élément |
| Idéal pour | N copies strictement identiques | objets nommés, hétérogènes |

> 🚨 On ne mélange pas `count` et `for_each` dans le même bloc (erreur de validation).

---

## 23. Blocs dynamic : générer des sous-blocs

Quand le nombre de sous-blocs varie (disques, interfaces réseau), `dynamic` les génère par boucle.

```hcl
variable "extra_disks" {
  type = list(object({
    size      = number
    datastore = string
  }))
  default = [
    { size = 50, datastore = "ceph-vm" },
    { size = 100, datastore = "ceph-vm" },
  ]
}

resource "proxmox_virtual_environment_vm" "db" {
  name      = "db-01"
  node_name = "pve2"

  disk {
    datastore_id = "local-lvm"
    size         = 40
    interface    = "scsi0"
  }

  # Un bloc disk par élément de var.extra_disks
  dynamic "disk" {
    for_each = var.extra_disks
    content {
      datastore_id = disk.value.datastore
      size         = disk.value.size
      interface    = "scsi${disk.key + 1}"   # scsi1, scsi2
    }
  }
}
```

- `disk.key` : index (liste) ou clé (map) de l'itération.
- `disk.value` : l'élément courant.
- Le nom du bloc `dynamic "disk"` doit correspondre exactement au nom du sous-bloc attendu par le provider.

---

## 24. Conditions et boucles en expressions

```hcl
locals {
  # Condition : choisir le datastore selon l'environnement
  vm_datastore = var.env == "prod" ? "ceph-vm" : "local-lvm"

  # Boucle for sur liste : construire des noms
  vm_names = [for i in range(3) : format("web-%02d", i + 1)]
  # => ["web-01", "web-02", "web-03"]

  # Boucle for sur map : filtrer
  prod_vms = {
    for name, spec in local.vms : name => spec
    if spec.env == "prod"
  }

  # Splat : extraire un attribut de chaque instance
  all_ips = proxmox_virtual_environment_vm.nodes[*].ipv4_addresses
  # Équivalent explicite :
  all_ips2 = [for vm in proxmox_virtual_environment_vm.nodes : vm.ipv4_addresses]
}
```

```hcl
# Ressource conditionnelle : count ternaire (0 = absente)
resource "proxmox_virtual_environment_vm" "bastion" {
  count     = var.env == "prod" ? 1 : 0
  name      = "bastion"
  node_name = "pve1"
  # ...
}
```

> 💡 Astuce : pour une ressource optionnelle avec `for_each`, utilisez `for_each = var.enabled ? toset(["x"]) : toset([])`.

## 25. Fonctions avancées : templatefile, fileset, jsonencode

```hcl
# cloud-init.tftpl (fichier template à côté du .tf)
#cloud-config
hostname: ${hostname}
manage_etc_hosts: true
users:
  - name: ${admin_user}
    sudo: ALL=(ALL) NOPASSWD:ALL
    ssh_authorized_keys:
      - ${ssh_key}
```

```hcl
# Utilisation
resource "proxmox_virtual_environment_vm" "web" {
  # ...
  initialization {
    user_data_file_id = proxmox_virtual_environment_file.cloud_config["web-01"].id
  }
}

# Générer un fichier cloud-init par VM depuis un template
resource "proxmox_virtual_environment_file" "cloud_config" {
  for_each     = local.vms
  content_type = "snippets"
  datastore_id = "local"
  node_name    = each.value.node

  source_raw {
    file_name = "${each.key}-cloud-config.yaml"
    data = templatefile("${path.module}/templates/cloud-init.tftpl", {
      hostname   = each.key
      admin_user = var.admin_user
      ssh_key    = var.ssh_public_key
    })
  }
}
```

```hcl
# fileset : lister des fichiers (ex. générer des snippets en masse)
locals {
  snippet_files = fileset("${path.module}/snippets", "*.yaml")
}

# jsonencode : passer une structure à une API
locals {
  metadata = jsonencode({
    env     = var.env
    owner   = "infra-team"
    created = timestamp()
  })
}
```

> ⚠️ `timestamp()` change à chaque plan → drift permanent. Ne l'utilisez jamais dans une ressource suivie, uniquement dans des `locals` éphémères ou des outputs informatifs.

---

## 26. terraform fmt et validate : l'hygiène du code

```bash
terraform fmt            # reformate tous les .tf du dossier (récursif avec -recursive)
terraform fmt -check     # vérifie sans modifier (idéal en CI : échoue si mal formaté)
terraform fmt -diff      # montre les différences

terraform validate       # vérifie la cohérence interne (types, références)
                         # SANS contacter les providers : rapide, à lancer avant plan
```

Pipeline minimal de qualité (à mettre en CI, section 57) :

```bash
terraform fmt -check -recursive
terraform init -backend=false
terraform validate
```

---

## 27. Providers : principe et installation

Un **provider** est un plugin (binaire) qui traduit le HCL en appels d'API : Proxmox, AWS, DNS, etc. Terraform Core orchestre ; les providers exécutent.

```hcl
terraform {
  required_providers {
    proxmox = {
      source  = "bpg/proxmox"   # <namespace>/<nom> sur le registry
      version = "~> 0.66"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 3.6"
    }
  }
}
```

```bash
terraform init   # télécharge les providers dans .terraform/
```

