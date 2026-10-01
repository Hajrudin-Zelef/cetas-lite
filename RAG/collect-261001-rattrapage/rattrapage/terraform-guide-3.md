---
id: collect-261001-rattrapage/rattrapage/terraform-guide-3
title: "Terraform — Guide complet : Infrastructure as Code en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-rattrapage/terraform_guide.md
source_anchor: ""
source_lines: [425, 713]
sha256: 04c37475fc21d4ee4ad7dafd0defe53efb2f0fe2dc639c58eea542f13dff0a5b
---

# Terraform — Guide complet : Infrastructure as Code en production

| Type | Exemple |
|---|---|
| `string` | `"web-01"` |
| `number` | `4`, `3.14` |
| `bool` | `true`, `false` |
| `list(string)` | `["web", "prod"]` |
| `map(string)` | `{ env = "prod", role = "web" }` |
| `object({...})` | `{ name = string, cores = number }` |
| `tuple` | `["web", 4, true]` (types mixtes) |
| `set(string)` | `toset(["a", "b"])` (sans doublons, sans ordre) |
| `null` | valeur absente (laisse le défaut du provider) |

```hcl
# Interpolation et expressions
name = "${var.prefix}-web-01"          # interpolation (souvent implicite)
name = "${var.prefix}-web-01"          # équivalent moderne :
name = "${var.prefix}-web-01"

# Opérateurs
cores    = var.ha_enabled ? 8 : 4      # ternaire
disk_gb  = var.base_disk + 20          # arithmétique
is_prod  = var.env == "prod"           # comparaison
combined = concat(["a"], ["b"])        # fonctions
```

> 💡 Depuis Terraform 0.12, l'interpolation `"${...}"` autour d'une expression seule est inutile : `name = var.prefix` suffit. On garde `"${}"` uniquement pour assembler des chaînes.

---

## 14. HCL : les fonctions intégrées les plus utiles

```hcl
# Chaînes
upper("prod")                    # "PROD"
lower("Web-01")                  # "web-01"
format("web-%02d", 1)            # "web-01"
formatlist("%s-prod", ["a","b"]) # ["a-prod", "b-prod"]
join(",", ["a", "b"])            # "a,b"
split(",", "a,b")                # ["a", "b"]
replace("web-01", "-", "_")      # "web_01"
trimspace("  x  ")               # "x"

# Listes / maps
length(["a", "b"])               # 2
element(["a", "b"], 0)           # "a"
concat(["a"], ["b"])             # ["a", "b"]
distinct(["a", "a", "b"])        # ["a", "b"]
flatten([["a"], ["b", "c"]])     # ["a", "b", "c"]
toset(["a", "b"])                # set
keys({a = 1, b = 2})             # ["a", "b"]
values({a = 1, b = 2})           # [1, 2]
merge({a = 1}, {b = 2})          # {a = 1, b = 2}
lookup({a = 1}, "b", 0)          # 0 (défaut)

# Nombres / logique
max(1, 5, 3)                     # 5
min(1, 5, 3)                     # 1
ceil(4.2)                        # 5
floor(4.8)                       # 4
coalesce("", "defaut")           # "defaut" (1er non-vide/non-null)
coalescelist([], ["x"])          # ["x"]

# Fichiers / encodage
file("${path.module}/init.sh")   # contenu d'un fichier
templatefile("${path.module}/cloud-init.tftpl", { hostname = "web-01" })
jsonencode({a = 1})              # "{\"a\":1}"
jsondecode("{\"a\":1}")          # {a = 1}
base64encode("secret")           # "c2VjcmV0"
yamlencode({a = 1})              # "a: 1\n"

# Réseau / IP
cidrhost("10.0.0.0/24", 10)      # "10.0.0.10"
cidrsubnet("10.0.0.0/16", 8, 1)  # "10.0.1.0/24"

# Tester dans le REPL :
# terraform console <<< 'cidrhost("10.0.0.0/24", 10)'
```

---

## 15. Variables : déclaration et typage

```hcl
# variables.tf
variable "env" {
  description = "Environnement cible (dev, staging, prod)"
  type        = string
  default     = "dev"
}

variable "vm_count" {
  description = "Nombre de VM web"
  type        = number
  default     = 2
}

variable "tags" {
  description = "Tags Proxmox"
  type        = list(string)
  default     = ["terraform"]
}

variable "vm_spec" {
  description = "Gabarit d'une VM"
  type = object({
    cores  = number
    memory = number   # en Mo
    disk   = number   # en Go
  })
  default = {
    cores  = 2
    memory = 4096
    disk   = 40
  }
}

variable "node_affinities" {
  description = "Nœud préféré par rôle"
  type        = map(string)
  default = {
    web = "pve1"
    db  = "pve2"
  }
}
```

Règles :
- `description` **toujours** renseignée (c'est de la documentation).
- `type` **toujours** déclaré : Terraform valide les entrées avant tout appel provider.
- Pas de `default` = variable **obligatoire** (Terraform demandera la valeur en interactif — à bannir en CI : utilisez `-var` ou `.tfvars`).

---

## 16. Variables : validation des entrées

```hcl
variable "env" {
  description = "Environnement cible"
  type        = string
  default     = "dev"

  validation {
    condition     = contains(["dev", "staging", "prod"], var.env)
    error_message = "env doit valoir dev, staging ou prod."
  }
}

variable "vm_count" {
  type    = number
  default = 2

  validation {
    condition     = var.vm_count >= 1 && var.vm_count <= 50
    error_message = "vm_count doit être entre 1 et 50."
  }
}

variable "cidr" {
  type    = string
  default = "10.10.0.0/24"

  validation {
    # Vérifie que c'est un CIDR valide
    condition     = can(cidrhost(var.cidr, 0))
    error_message = "cidr doit être un bloc CIDR valide (ex. 10.10.0.0/24)."
  }
}

variable "hostname" {
  type    = string
  default = "web-01"

  validation {
    condition     = can(regex("^[a-z0-9-]+$", var.hostname))
    error_message = "hostname : minuscules, chiffres et tirets uniquement."
  }
}
```

> 💡 La validation s'exécute **avant** le plan : c'est le premier rempart contre les erreurs de saisie, bien avant tout appel à Proxmox.

---

## 17. Variables : les fichiers .tfvars et la précédence

```hcl
# envs/prod.tfvars
env      = "prod"
vm_count = 10
tags     = ["terraform", "prod"]
vm_spec = {
  cores  = 4
  memory = 8192
  disk   = 80
}
```

```bash
terraform plan -var-file=envs/prod.tfvars
terraform plan -var-file=envs/prod.tfvars -var="vm_count=12"  # surcharge ponctuelle
```

**Ordre de précédence** (du plus faible au plus fort) :

1. `default` dans la déclaration
2. `terraform.tfvars` / `*.auto.tfvars` (chargés automatiquement, ordre alphabétique)
3. `-var-file=...` (ordre de la ligne de commande)
4. `-var 'nom=valeur'`
5. Variable d'environnement `TF_VAR_nom`

```bash
export TF_VAR_env="staging"   # pratique en CI
```

> 🚨 **Sécurité** : les `.tfvars` contenant des secrets ne vont **jamais** dans Git (voir sections 54-55). Versionnez uniquement `terraform.tfvars.example`.

---

## 18. Outputs : exposer des valeurs

```hcl
# outputs.tf
output "vm_ipv4_addresses" {
  description = "Adresses IPv4 des VM web"
  value       = [for vm in proxmox_virtual_environment_vm.web : vm.ipv4_addresses]
}

output "vm_ids" {
  description = "IDs Proxmox des VM"
  value       = { for k, vm in proxmox_virtual_environment_vm.web : k => vm.vm_id }
}

output "db_password" {
  description = "Mot de passe généré pour la BDD"
  value       = random_password.db.result
  sensitive   = true   # masqué dans les logs et la sortie console
}
```

```bash
terraform output                  # affiche les outputs non sensibles
terraform output -json            # format JSON (pour scripts/Ansible)
terraform output -raw vm_ids      # valeur brute d'un output simple
```

> 💡 `sensitive = true` masque la valeur à l'écran, mais elle reste **en clair dans le state** (section 44). Ce n'est pas du chiffrement.

---

## 19. Locals : factoriser la logique

Les `locals` sont des valeurs calculées une fois et réutilisées : nommage, tags communs, listes dérivées.

```hcl
# locals.tf
locals {
  # Convention de nommage centralisée
  name_prefix = "${var.env}-${var.project}"

  # Tags communs à toutes les ressources
  common_tags = concat(var.tags, ["env:${var.env}", "managed-by:terraform"])

  # Liste des VM à créer : dérivée des variables
  web_nodes = [
    for i in range(var.vm_count) : {
      name = format("%s-web-%02d", local.name_prefix, i + 1)
      ip   = cidrhost(var.cidr, 10 + i)
    }
  ]

  # Map indexée par nom (pratique pour for_each)
  web_nodes_map = { for n in local.web_nodes : n.name => n }
}
```

Différences `locals` vs `variable` vs `output` :

| | Source | Modifiable par l'appelant | Visible hors module |
|---|---|---|---|
| `variable` | entrée externe | oui (`-var`, tfvars) | non (entrée) |
| `locals` | calcul interne | non | non |
| `output` | valeur exposée | non | oui |

---

## 20. Data sources : lire l'existant

Les data sources **lisent** des informations sans les gérer : un template existant, un nœud, un datastore.

