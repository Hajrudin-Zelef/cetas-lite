---
id: collect-261001-rattrapage/rattrapage/terraform-guide-9
title: "Terraform — Guide complet : Infrastructure as Code en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "valuation"]
source: docs/RAG/collect-261001-rattrapage/terraform_guide.md
source_anchor: ""
source_lines: [2010, 2199]
sha256: 8e1441568dc39b53757e3b495ce59ddc6c600843bfaea40ff35d59f97962878a
---

# Terraform — Guide complet : Infrastructure as Code en production

```bash
# Autres leviers de debug
terraform plan -refresh=false     # plan sans interroger Proxmox (rapide, state seul)
terraform apply -parallelism=1    # sérialise (isole les problèmes de concurrence)
TF_CLI_ARGS_plan="-lock-timeout=60s" terraform plan   # attend le verrou au lieu d'échouer
```

> ⚠️ Les logs `TRACE`/`DEBUG` peuvent contenir des **secrets** (tokens, mots de passe). Ne les publiez jamais tels quels, purgez-les après usage.

---

## 60. 10+ erreurs classiques et leurs solutions

### Erreur 1 — `Error: Failed to install provider` / checksum mismatch
**Cause** : lock file corrompu ou miroir compromis.
**Solution** : `rm .terraform.lock.hcl && terraform init` (puis `git diff` du nouveau lock avant commit).

### Erreur 2 — `Error: acquiring the state lock`
**Cause** : un `apply` précédent a planté en laissant le verrou, ou deux apply concurrents.
**Solution** : vérifier qu'aucun apply ne tourne vraiment, puis `terraform force-unlock <ID>` (l'ID est dans le message d'erreur). Jamais à l'aveugle.

### Erreur 3 — `Error: Instance cannot be destroyed` (prevent_destroy)
**Cause** : le plan veut détruire une ressource protégée.
**Solution** : c'est le garde-fou qui fait son travail. Vérifier pourquoi la ressource est détruite (renommage ? `moved` oublié ?), corriger le code.

### Erreur 4 — `Provider produced inconsistent result after apply`
**Cause** : le provider retourne après création une valeur différente du plan (souvent un défaut côté API).
**Solution** : mettre à jour le provider ; contournement temporaire : `ignore_changes` sur l'attribut fautif.

### Erreur 5 — `Error: Duplicate resource` / VMID déjà utilisé
**Cause** : `vm_id` en dur déjà pris par une VM hors Terraform.
**Solution** : importer la VM existante (section 51) ou choisir un VMID libre (`pvesh get /cluster/nextid`).

### Erreur 6 — `Error: Reference to undeclared resource`
**Cause** : faute de frappe dans le nom local, ou ressource dans un module non préfixée par `module.`.
**Solution** : `terraform state list` pour voir les vrais noms ; corriger la référence.

### Erreur 7 — Cycle : `Error: Cycle: ...`
**Cause** : A dépend de B qui dépend de A (souvent via deux `depends_on` croisés).
**Solution** : `terraform graph` pour visualiser, supprimer la dépendance artificielle, passer par des data sources si besoin.

### Erreur 8 — `Error: Invalid for_each argument` (valeur non connue au plan)
**Cause** : `for_each` sur un attribut calculé (connu seulement après apply).
**Solution** : `for_each` exige des clés connues au plan. Restructurer (map statique en variable) ou découper en deux `apply`.

### Erreur 9 — Dérive permanente sur un attribut (ex. `tags` réordonnés)
**Cause** : le provider normalise différemment (ordre, casse).
**Solution** : `ignore_changes`, ou normaliser côté code (`sort()`), ou monter de version du provider.

### Erreur 10 — `terraform plan` interminable / timeouts API Proxmox
**Cause** : trop de ressources rafraîchies, API lente.
**Solution** : `-refresh=false` pour un plan rapide (à n'utiliser qu'en lecture), `-parallelism` réduit, vérifier la charge du nœud Proxmox.

### Erreur 11 — `Error: Failed to query available provider packages` (registry injoignable)
**Cause** : pas d'accès internet ou proxy non configuré.
**Solution** : `export HTTPS_PROXY=...`, ou miroir de providers local (`terraform init -backend=false` + `provider_installation` dans `.terraformrc`).

### Erreur 12 — Secrets visibles dans `terraform show` / logs CI
**Cause** : output sans `sensitive = true`, ou `TF_LOG=TRACE` en CI.
**Solution** : marquer `sensitive`, purger les logs, rotation du secret exposé.

## 61. Mise à jour des providers

```bash
# Voir les versions disponibles (registry)
terraform init -upgrade   # monte à la dernière version autorisée par la contrainte ~>

# Monter explicitement de branche (ex. 0.66 -> 0.70) :
# 1. Éditer versions.tf : version = "~> 0.70"
# 2. Lire le CHANGELOG du provider (breaking changes !)
# 3. terraform init -upgrade
# 4. terraform plan  → relire TOUT (un upgrade peut proposer des diffs)
# 5. Tester sur dev/staging avant prod
```

Procédure d'équipe :
1. Une MR par montée de provider (jamais mélangée avec du fonctionnel).
2. `terraform plan` sur **chaque environnement** (dev → staging → prod).
3. Ne jamais `-upgrade` en CI prod sans relecture humaine.

---

## 62. Mise à jour de Terraform / OpenTofu

```bash
# Vérifier la version requise par le projet
grep required_version versions.tf

# Mettre à jour le binaire (ex. via tenv)
tenv terraform install 1.10.5 && tenv terraform use 1.10.5

# Après changement de version MINEURE/MAJEURE :
terraform init
terraform plan   # le format du state peut migrer (sauvegarde auto en .backup)
```

| Saut de version | Risque | Conduite |
|---|---|---|
| Patch (1.10.1 → 1.10.2) | faible | direct, plan de contrôle |
| Mineur (1.9 → 1.10) | moyen | lire l'upgrade guide, tester dev/staging |
| Majeur (1.x → 2.x) | élevé | plan de migration, fenêtre de maintenance |

> ⚠️ `required_version` bloque l'équipe sur une version non prévue : mettez-la à jour **volontairement** dans la même MR que le test.

---

## 63. OpenTofu : différences avec Terraform

| Fonctionnalité | Terraform | OpenTofu |
|---|---|---|
| Licence | BSL | MPL-2.0 |
| `terraform {}` / CLI | `terraform` | `tofu` (alias `terraform` possible) |
| Chiffrement du state au repos | ❌ (délégué au backend) | ✅ natif (`key_provider`) |
| `removed` block | ❌ | ✅ (retirer proprement sans détruire) |
| Variables d'itération `for_each`/`count` dans `import` | partiel | ✅ (`for_each` sur blocs `import`) |
| Early evaluation / `static` | ❌ | ✅ (évaluation anticipée) |
| Registre par défaut | registry.terraform.io | idem (compatible) |

Exemple : chiffrement natif du state avec OpenTofu :

```hcl
# versions.tf (OpenTofu)
terraform {
  encryption {
    key_provider "pbkdf2" "main" {
      passphrase = var.state_passphrase   # via TF_VAR_state_passphrase
    }
    method "aes_gcm" "main" {
      keys = key_provider.pbkdf2.main
    }
    state {
      method = method.aes_gcm.main
    }
    plan {
      method = method.aes_gcm.main
    }
  }
}
```

---

## 64. Migration Terraform → OpenTofu

La migration est conçue pour être **sans douleur** (même langage, même state) :

```bash
# 1. Sauvegarde du state (obligatoire)
terraform state pull > /srv/backups/pre-tofu-$(date +%F).tfstate

# 2. Installer tofu (section 5), puis dans le projet :
tofu init          # lit le même backend, le même lock file

# 3. Vérifier : le plan doit être VIDE (aucun changement)
tofu plan
# "No changes. Your infrastructure matches the configuration."

# 4. Basculer la CI sur l'image opentofu/opentofu:latest
# 5. Optionnel : renommer terraform {} -> mêmes blocs (compatibles),
#    activer le chiffrement du state (section 63)
```

Points d'attention :
- `required_version` : le remplacer par la contrainte tofu si vous figez (`tofu` ignore `required_version` ? non — il l'honore ; utilisez la version tofu correspondante).
- Les providers du registry Terraform fonctionnent tels quels.
- Gardez le binaire `terraform` sous la main le temps de la transition (rollback = `terraform init` + `plan` vide).

---

## 65. Cas pratique 1 : parc de 10 VM Proxmox depuis du code

Objectif : 10 VM web identiques, réparties sur 2 nœuds, IP automatiques, cloud-init, tags.

```hcl
# versions.tf
terraform {
  required_version = ">= 1.9.0, < 1.11.0"
  required_providers {
    proxmox = { source = "bpg/proxmox", version = "~> 0.66" }
  }
  backend "s3" {
    bucket = "infra-terraform-state-prod"
    key    = "proxmox/web-farm/terraform.tfstate"
    region = "eu-west-3"
    encrypt = true
    dynamodb_table = "terraform-state-locks"
  }
}
```

