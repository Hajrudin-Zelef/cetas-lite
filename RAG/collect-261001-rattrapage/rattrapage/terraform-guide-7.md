---
id: collect-261001-rattrapage/rattrapage/terraform-guide-7
title: "Terraform — Guide complet : Infrastructure as Code en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["aws", "memory"]
source: docs/RAG/collect-261001-rattrapage/terraform_guide.md
source_anchor: ""
source_lines: [1502, 1761]
sha256: 731891eedf3f8f47aac8af0cd7444bdf4c6b3c7b563bef97ad48bfeba0bc105b
---

# Terraform — Guide complet : Infrastructure as Code en production

## 45. Manipulation du state : mv, rm, pull, push

```bash
# Lister
terraform state list
terraform state show 'proxmox_virtual_environment_vm.nodes["web-01"]'

# Renommer une ressource dans le code SANS la recréer
# (ex. web_01 -> web dans le .tf) :
terraform state mv 'proxmox_virtual_environment_vm.web_01' \
                   'proxmox_virtual_environment_vm.nodes["web-01"]'

# Retirer du state SANS détruire l'objet réel (Terraform l'oublie)
terraform state rm 'proxmox_virtual_environment_vm.old'

# Réimporter ensuite si besoin
terraform import 'proxmox_virtual_environment_vm.old' pve1:qemu/109
```

> 💡 La méthode moderne pour les renommages : le bloc **`moved`** dans le code (traçable en Git) :
>
> ```hcl
> moved {
>   from = proxmox_virtual_environment_vm.web_01
>   to   = proxmox_virtual_environment_vm.nodes["web-01"]
> }
> ```

---

## 46. Workspaces : un state par environnement

```bash
terraform workspace new dev
terraform workspace new prod
terraform workspace list
terraform workspace select prod
terraform plan -var-file=envs/prod.tfvars
```

- Chaque workspace = un **state séparé** dans le même backend (`key` suffixé par le workspace).
- `${terraform.workspace}` utilisable dans le code :

```hcl
locals {
  name_prefix = "${terraform.workspace}-${var.project}"
}
```

⚠️ **Limites** : les workspaces partagent le même code et les mêmes variables par défaut — faciles à confondre (`apply` sur le mauvais workspace = catastrophe). Pour des environnements bien séparés, préférez des **dossiers distincts** (section 47). Workspaces = pratique pour dev/test éphémères, pas pour prod sérieuse.

---

## 47. Séparation des environnements : dev / staging / prod

Architecture recommandée :

```
infra/
├── modules/vm/                 # module partagé
├── envs/
│   ├── dev/
│   │   ├── main.tf             # appelle le module, petites VM
│   │   ├── backend.tf          # key = "proxmox/dev/terraform.tfstate"
│   │   └── terraform.tfvars
│   ├── staging/
│   │   ├── main.tf
│   │   ├── backend.tf          # key = "proxmox/staging/terraform.tfstate"
│   │   └── terraform.tfvars
│   └── prod/
│       ├── main.tf
│       ├── backend.tf          # key = "proxmox/prod/terraform.tfstate"
│       └── terraform.tfvars
```

| Principe | Mise en œuvre |
|---|---|
| 1 dossier = 1 state = 1 env | `envs/prod` ne touche jamais le state de dev |
| Code partagé | modules locaux ou registry |
| Différences | uniquement via `.tfvars` et variables |
| Promotion | `dev` → `staging` → `prod`, même code, plan relu à chaque étape |
| Accès | token Proxmox et droits S3 distincts par env |

```hcl
# envs/prod/main.tf : exemple minimal
module "web_farm" {
  source   = "../../modules/vm"
  env      = "prod"
  vm_count = 10
  # ...
}
```

---

## 48. Modules : créer son premier module

Un module = un dossier avec des `.tf`, une interface (`variables.tf` / `outputs.tf`).

```
modules/vm/
├── main.tf        # ressources
├── variables.tf   # interface d'entrée
├── outputs.tf     # interface de sortie
└── README.md      # documentation (but, exemple)
```

```hcl
# modules/vm/variables.tf
variable "name"      { description = "Nom de la VM"; type = string }
variable "node_name" { description = "Nœud Proxmox"; type = string }
variable "vm_id"     { description = "VMID"; type = number }
variable "cores"     { description = "vCPU"; type = number; default = 2 }
variable "memory"    { description = "RAM en Mo"; type = number; default = 4096 }
variable "ip_address" { description = "IP avec CIDR, ex. 10.10.0.11/24"; type = string }
variable "gateway"   { description = "Passerelle"; type = string; default = "10.10.0.1" }

# modules/vm/main.tf
resource "proxmox_virtual_environment_vm" "this" {
  name      = var.name
  node_name = var.node_name
  vm_id     = var.vm_id
  tags      = ["terraform", "module-vm"]

  cpu { cores = var.cores }
  memory { dedicated = var.memory }
  disk {
    datastore_id = "local-lvm"
    interface    = "scsi0"
    size         = 40
  }
  network_device { bridge = "vmbr0" }
  initialization {
    datastore_id = "local-lvm"
    ip_config {
      ipv4 { address = var.ip_address; gateway = var.gateway }
    }
  }
}

# modules/vm/outputs.tf
output "vm_id" { description = "VMID"; value = proxmox_virtual_environment_vm.this.vm_id }
output "name"  { description = "Nom";  value = proxmox_virtual_environment_vm.this.name }
```

## 49. Modules : utiliser un module (registry, git, local)

```hcl
# Module local
module "web_01" {
  source     = "../../modules/vm"
  name       = "web-01"
  node_name  = "pve1"
  vm_id      = 101
  cores      = 4
  memory     = 8192
  ip_address = "10.10.0.11/24"
}

# Module depuis le registry public
module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "~> 5.0"
  # ...
}

# Module depuis Git (tag figé !)
module "vm" {
  source = "git::https://git.entreprise.lan/infra/tf-modules.git//vm?ref=v1.2.0"
  # ...
}

# Accès aux outputs d'un module
output "web_01_id" {
  value = module.web_01.vm_id
}
```

```bash
terraform init -upgrade   # met à jour les modules vers la contrainte
terraform get             # télécharge les modules sans init complet
```

> 🚨 Toujours figer `version` (registry) ou `?ref=` (git). Un module flottant sur `main` = surprise garantie un lundi matin.

---

## 50. Bonnes pratiques de conception de modules

- [ ] **Une responsabilité** par module (ex. `vm`, pas `infra-complete`).
- [ ] Interface minimale : peu de variables obligatoires, des `default` sensés.
- [ ] `description` sur chaque variable et output.
- [ ] Pas de provider configuré dans un module partagé (le provider est configuré à la racine ; le module hérite).
- [ ] Pas de backend dans un module.
- [ ] Versionner (tag Git `v1.0.0`) + CHANGELOG.
- [ ] `README.md` avec exemple d'appel copiable.
- [ ] Tester avec `terraform plan` sur un exemple (`examples/`).

```
modules/vm/
├── main.tf
├── variables.tf
├── outputs.tf
├── README.md
├── CHANGELOG.md
└── examples/
    └── basic/
        └── main.tf     # exemple minimal, testé en CI
```

---

## 51. Import de ressources existantes (import block)

Une VM créée à la main dans Proxmox peut entrer sous gestion Terraform **sans être recréée**.

```hcl
# 1. Déclarer la ressource cible (config aussi proche que possible du réel)
resource "proxmox_virtual_environment_vm" "legacy_db" {
  name      = "db-legacy"
  node_name = "pve2"
  vm_id     = 200
  # ... (le reste sera ajusté après le premier plan)
}

# 2. Bloc d'import (Terraform ≥ 1.5) : traçable en Git, réapplicable
import {
  to = proxmox_virtual_environment_vm.legacy_db
  id = "pve2:qemu/200"   # format d'ID selon le provider (voir sa doc)
}
```

```bash
terraform plan -out=tfplan   # montre ce que Terraform voudrait CHANGER après import
# Ajuster le .tf jusqu'à obtenir "No changes" (ou des changements acceptés)
terraform apply tfplan
```

> 💡 Méthode : importez, lancez `plan`, ajustez la config pour coller au réel, recommencez jusqu'à zéro diff non voulu. Ne jamais importer en prod sans relecture du plan.

---

## 52. Provisioners : pourquoi les éviter

```hcl
# ❌ À ÉVITER : fragile, non idempotent, pas de gestion d'erreur fine
resource "proxmox_virtual_environment_vm" "web" {
  # ...
  provisioner "remote-exec" {
    inline = ["apt update", "apt install -y nginx"]
    connection {
      type        = "ssh"
      user        = "admin"
      private_key = file("~/.ssh/id_ed25519")
      host        = self.ipv4_addresses[0]
    }
  }
}
```

