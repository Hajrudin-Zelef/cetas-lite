---
id: collect-261001-rattrapage/rattrapage/terraform-guide-2
title: "Terraform — Guide complet : Infrastructure as Code en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["open source"]
source: docs/RAG/collect-261001-rattrapage/terraform_guide.md
source_anchor: ""
source_lines: [184, 424]
sha256: cc24681602de528047e7bb0925898f0636dca7a3d1a0bda02868ac54317599b9
---

# Terraform — Guide complet : Infrastructure as Code en production

```bash
# Dépôt officiel OpenTofu (Debian/Ubuntu)
curl -fsSL https://get.opentofu.org/opentofu.gpg \
  | sudo tee /etc/apt/keyrings/opentofu.gpg >/dev/null
curl -fsSL https://packages.opentofu.org/opentofu/tofu/any-any/stable/tofu.list \
  | sudo tee /etc/apt/sources.list.d/opentofu.list >/dev/null
sudo chmod a+r /etc/apt/keyrings/opentofu.gpg /etc/apt/sources.list.d/opentofu.list
sudo apt update && sudo apt install -y tofu

# Vérification
tofu version
# OpenTofu v1.9.x
```

| Critère | Terraform | OpenTofu |
|---|---|---|
| Licence | BSL (source disponible, usage restreint) | MPL-2.0 (vraiment open source) |
| Registre | registry.terraform.io | Compatible registry Terraform + le sien |
| Chiffrement du state | Non natif | Oui (natif, section 42) |
| Gouvernance | HashiCorp | Linux Foundation / communauté |

> 💡 En 2026, les deux coexistent. OpenTofu est le choix prudent pour éviter tout risque de licence. La migration est couverte section 64.

---

## 6. Configuration du shell et des outils

```bash
# Complétion bash (à ajouter dans ~/.bashrc)
complete -C /usr/bin/terraform terraform
# Pour OpenTofu :
complete -C /usr/bin/tofu tofu

# Alias pratiques
alias tf='terraform'
alias tfp='terraform plan'
alias tfa='terraform apply'

# Éditeur : VS Code + extension "HashiCorp Terraform"
# (coloration HCL, fmt à la sauvegarde, validation)
```

Fichiers à ignorer dans Git (`.gitignore`) :

```gitignore
# Terraform
.terraform/
*.tfstate
*.tfstate.*
*.tfvars
!example.tfvars
.terraform.lock.hcl
override.tf
override.tf.json
*_override.tf
*_override.tf.json
crash.log
crash.*.log

# OpenTofu
.tofu/
```

---

## 7. Vérification, versions et gestion multi-versions

```bash
terraform version          # version installée
terraform version -json   # format JSON (utile en CI)

# Figer la version requise dans le code (OBLIGATOIRE en prod) :
```

```hcl
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

**Gestion multi-versions** : `tfenv` (ou `tenv`, qui gère aussi tofu) :

```bash
# tenv : gère terraform + tofu + terragrunt
curl -sL https://github.com/tofuutils/tenv/releases/latest/download/tenv_Linux_x86_64.tar.gz \
  | tar xz -C /tmp && sudo mv /tmp/tenv /usr/local/bin/
tenv terraform install 1.10.5
tenv terraform use 1.10.5
```

> ⚠️ En équipe, **tout le monde utilise la même version** (via `.terraform-version` ou la CI). Un state écrit par une version majeure différente peut poser problème.

---

## 8. Structure d'un projet Terraform

```
infra-proxmox/
├── README.md                  # but du projet, prérequis
├── versions.tf                # required_version + required_providers
├── providers.tf               # configuration des providers
├── backend.tf                 # configuration du backend (state distant)
├── variables.tf               # déclarations de variables
├── terraform.tfvars.example   # exemple de valeurs (JAMAIS de secrets)
├── locals.tf                  # valeurs locales calculées
├── main.tf                    # ressources principales
├── outputs.tf                 # sorties
├── modules/
│   └── vm/                    # module local "vm"
│       ├── main.tf
│       ├── variables.tf
│       └── outputs.tf
└── envs/
    ├── dev.tfvars
    ├── staging.tfvars
    └── prod.tfvars
```

Règles :
- **Un dossier = une stack** (un state). Ne jamais mélanger prod et dev dans le même dossier.
- Noms de fichiers en minuscules, suffixe `.tf`.
- `main.tf` ne doit pas faire 2000 lignes : découper par thème (`network.tf`, `vms.tf`, `storage.tf`).

---

## 9. Le workflow de base : init → plan → apply

```bash
cd infra-proxmox/

# 1. Initialisation : télécharge providers + modules, configure le backend
terraform init

# 2. Plan : calcule les actions SANS rien modifier (lecture seule)
terraform plan -out=tfplan

# 3. RELIRE LE PLAN (toujours !) puis appliquer exactement ce plan
terraform apply tfplan
```

Le plan affiche :

```
Terraform will perform the following actions:

  # proxmox_virtual_environment_vm.web_01 will be created
  + resource "proxmox_virtual_environment_vm" "web_01" {
      + name      = "web-01"
      + node_name = "pve1"
      ...
    }

Plan: 1 to add, 0 to change, 0 to destroy.
```

Symboles : `+` créer, `~` modifier, `-` détruire, `-/+` remplacer (détruire puis recréer), `<=` lire (data source).

> 🚨 **Règle d'or** : `terraform apply` sans `-out`/`tfplan` recalcule un plan à la volée — acceptable en dev, **interdit en prod**. En prod : `plan -out=tfplan` → relecture humaine → `apply tfplan`.

---

## 10. destroy : détruire proprement

```bash
# Détruire TOUTE la stack (demande confirmation)
terraform destroy

# Détruire avec plan préalable (recommandé en prod)
terraform plan -destroy -out=tfplan-destroy
terraform apply tfplan-destroy

# Détruire une seule ressource (ciblée)
terraform destroy -target=proxmox_virtual_environment_vm.web_01
```

> ⚠️ `-target` est un outil de dépannage, pas un mode de fonctionnement normal : il crée des dépendances partielles et peut laisser le state incohérent. À utiliser avec parcimonie.

**Protection** : combinez `prevent_destroy` (section 34) sur les ressources critiques pour rendre `destroy` impossible par accident.

---

## 11. La CLI Terraform : commandes essentielles

| Commande | Usage |
|---|---|
| `terraform init` | Initialise le dossier (providers, backend, modules) |
| `terraform plan` | Affiche les changements prévus |
| `terraform apply` | Applique les changements |
| `terraform destroy` | Détruit l'infrastructure gérée |
| `terraform fmt` | Reformate le code HCL |
| `terraform validate` | Vérifie la syntaxe/cohérence (sans provider) |
| `terraform show` | Affiche le state ou un plan |
| `terraform output` | Affiche les outputs |
| `terraform state list` | Liste les ressources du state |
| `terraform import` | Importe une ressource existante |
| `terraform taint` | Marque une ressource à recréer |
| `terraform workspace` | Gère les workspaces |
| `terraform providers` | Liste les providers requis |
| `terraform graph` | Génère le graphe de dépendances (DOT) |
| `terraform console` | REPL pour tester des expressions |
| `terraform force-unlock` | Libère un verrou bloqué (avec l'ID !) |

```bash
# Exemples utiles au quotidien
terraform output -json | jq .                    # outputs en JSON
terraform state list | grep proxmox_virtual_environment_vm
terraform console <<< 'format("%s-%02d", "web", 1)'   # tester une expression
```

---

## 12. HCL : le langage — blocs, arguments, identifiants

HCL (HashiCorp Configuration Language) : syntaxe `bloc "type" "nom" { argument = valeur }`.

```hcl
# Bloc "resource" : type = proxmox_virtual_environment_vm, nom local = web_01
resource "proxmox_virtual_environment_vm" "web_01" {
  name      = "web-01"   # argument : chaîne
  node_name = "pve1"     # argument : chaîne
  cpu {                  # sous-bloc imbriqué
    cores = 4            # argument : nombre
  }
  tags = ["web", "prod"] # argument : liste
}
```

- Le **nom local** (`web_01`) n'existe que dans le code : c'est l'identifiant pour les références (`proxmox_virtual_environment_vm.web_01.name`).
- Le **nom réel** (`web-01`) est l'argument `name` vu par Proxmox.
- Commentaires : `#` ou `//` (ligne), `/* ... */` (bloc).
- Indentation : 2 espaces (`terraform fmt` l'impose).

## 13. HCL : types de valeurs et expressions

