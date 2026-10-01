---
id: collect-261001-rattrapage/rattrapage/terraform-guide-8
title: "Terraform — Guide complet : Infrastructure as Code en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/terraform_guide.md
source_anchor: ""
source_lines: [1762, 2009]
sha256: a43c7c0fbc86773fdc37f6161fe88bccc789d842fb815494f5eacc3459dc54ee
---

# Terraform — Guide complet : Infrastructure as Code en production

Problèmes :
- S'exécute **après** création, hors du modèle déclaratif (pas de drift detection dessus).
- Échec réseau/SSH = ressource "créée mais pas configurée" (état bâtard, `taint` manuel).
- Secrets de connexion dans le code/state.
- `local-exec` dépend de la machine qui lance Terraform (non portable en CI).

> 🚨 Règle d'équipe : **aucun provisioner dans le code partagé**. Ils restent un outil de dernier recours, documenté comme tel.

---

## 53. Alternatives aux provisioners (cloud-init, Ansible, Packer)

| Besoin | Solution propre |
|---|---|
| Config initiale (user, SSH, réseau, paquets) | **cloud-init** via `user_data_file_id` (natif Proxmox) |
| Configuration applicative | **Ansible** après `terraform apply` (inventaire généré depuis `terraform output -json`) |
| Image durcie/répétable | **Packer** pour builder le template, Terraform pour cloner |
| Script one-shot fiable | `null_resource` + `local-exec` avec `triggers` (explicite, versionné) |

```bash
# Chaînage Terraform -> Ansible (classique et robuste)
terraform apply -auto-approve tfplan
terraform output -json > /tmp/tf-outputs.json
ansible-playbook -i inventory_terraform.py site.yml
```

```hcl
# Exemple cloud-init complet (templates/cloud-init.tftpl)
#cloud-config
hostname: ${hostname}
manage_etc_hosts: true
package_update: true
packages: [qemu-guest-agent, curl, vim]
users:
  - name: ${admin_user}
    sudo: ALL=(ALL) NOPASSWD:ALL
    shell: /bin/bash
    ssh_authorized_keys: [${ssh_key}]
runcmd:
  - systemctl enable --now qemu-guest-agent
```

---

## 54. Sécurité : jamais de secrets en clair

Règles non négociables :

1. **Aucun secret dans les `.tf`** (ni token, ni mot de passe, ni clé privée).
2. **Aucun secret dans Git** — même en repo privé. L'historique Git n'oublie jamais.
3. Variables sensibles : `sensitive = true` + injection via environnement ou fichier chiffré.
4. Le state contient des secrets en clair → backend chiffré + accès restreint (section 44).

```hcl
variable "proxmox_api_token" {
  type      = string
  sensitive = true
  # PAS de default avec un vrai secret !
}
```

```bash
# Injection par environnement (CI ou poste admin)
export TF_VAR_proxmox_api_token="terraform@pve!tf-token=xxxxxxxx"
terraform plan
```

```bash
# Vérifier qu'aucun secret ne traîne (à mettre en pre-commit)
git secrets --scan
# ou : truffleHog / gitleaks
gitleaks detect --source . --verbose
```

> 💡 Si un secret a fuité dans Git : le **révoquer immédiatement** (rotation), puis nettoyer l'historique (`git filter-repo`). Le nettoyage seul ne suffit jamais.

---

## 55. Chiffrer les tfvars : SOPS + age

[SOPS](https://github.com/getsops/sops) chiffre les fichiers de variables avec `age` (simple, sans serveur).

```bash
# Installation
sudo apt install -y age
# sops : binaire depuis les releases GitHub
curl -sL https://github.com/getsops/sops/releases/latest/download/sops-v3.x.x.linux.amd64 \
  -o /tmp/sops && sudo install /tmp/sops /usr/local/bin/sops

# Générer une clé age (1 par admin, publique versionnée, privée gardée)
age-keygen -o ~/.config/sops/age/keys.txt
# Public key: age1xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

```yaml
# .sops.yaml (à la racine, versionné)
creation_rules:
  - path_regex: envs/prod/.*\.tfvars$
    age: age1xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

```bash
# Chiffrer le tfvars de prod
sops --encrypt --in-place envs/prod/secrets.tfvars

# Utilisation : déchiffrer à la volée (jamais de fichier clair sur disque)
sops --decrypt envs/prod/secrets.tfvars > /tmp/secrets.tfvars
terraform plan -var-file=envs/prod.tfvars -var-file=/tmp/secrets.tfvars
shred -u /tmp/secrets.tfvars
```

> 💡 Alternative : **Vault** (HashiCorp) via le provider `vault` + data sources, ou variables d'environnement injectées par le gestionnaire de secrets de la CI.

---

## 56. Détection de dérive (drift detection)

La dérive = la réalité s'écarte du code (modif manuelle dans l'UI Proxmox, script externe).

```bash
# Job planifié (cron ou CI nightly) : plan en lecture seule
terraform plan -refresh-only -detailed-exitcode -out=/tmp/drift.plan
# Codes de sortie : 0 = pas de changement, 1 = erreur, 2 = changements détectés !
```

Script de détection + alerte :

```bash
#!/bin/bash
# /usr/local/bin/terraform-drift-check.sh
set -euo pipefail
cd /opt/infra/envs/prod
terraform init -input=false >/dev/null
if terraform plan -refresh-only -detailed-exitcode -out=/tmp/drift.plan >/tmp/drift.log 2>&1; then
  echo "OK: aucune dérive"
else
  code=$?
  if [ $code -eq 2 ]; then
    echo "ALERTE: dérive détectée sur prod" | \
      mail -s "[Terraform] Drift prod" infra@entreprise.lan < /tmp/drift.log
    # ou webhook vers le canal d'astreinte
  else
    echo "ERREUR terraform (code $code)" >&2; exit 1
  fi
fi
```

```
# Cron : vérification quotidienne à 6h
0 6 * * * /usr/local/bin/terraform-drift-check.sh
```

> 💡 En CI (GitLab nightly, section 57), le même plan peut ouvrir une issue automatique. L'objectif : **zéro changement hors Terraform**, ou documenté via `ignore_changes`.

---

## 57. CI/CD : pipeline GitLab de revue et d'apply

```yaml
# .gitlab-ci.yml
stages: [validate, plan, apply]

variables:
  TF_IN_AUTOMATION: "true"

fmt-validate:
  stage: validate
  image: hashicorp/terraform:1.10
  script:
    - terraform fmt -check -recursive
    - terraform init -backend=false
    - terraform validate
  rules:
    - if: $CI_PIPELINE_SOURCE == "merge_request_event"

plan-prod:
  stage: plan
  image: hashicorp/terraform:1.10
  script:
    - cd envs/prod
    - terraform init -input=false
    - terraform plan -input=false -out=tfplan
    - terraform show -json tfplan > plan.json
  artifacts:
    paths: [envs/prod/tfplan, envs/prod/plan.json]
    expire_in: 1 week
  rules:
    - if: $CI_PIPELINE_SOURCE == "merge_request_event"

apply-prod:
  stage: apply
  image: hashicorp/terraform:1.10
  script:
    - cd envs/prod
    - terraform init -input=false
    - terraform apply -input=false -auto-approve tfplan
  rules:
    - if: $CI_COMMIT_BRANCH == "main"
  when: manual            # apply prod = validation humaine obligatoire
  environment: production
```

Points clés :
- `TF_IN_AUTOMATION=true` : sorties adaptées aux logs CI.
- Le **plan est un artefact** relu dans la MR, puis appliqué tel quel.
- `apply` en `manual` sur `main` uniquement : traçabilité totale.
- Secrets via variables CI masquées/protégées (`TF_VAR_proxmox_api_token`).

---

## 58. Revue de plan en équipe

Checklist du relecteur avant d'approuver un plan :

- [ ] Le plan correspond-il à l'intention de la MR (pas de changement surprise) ?
- [ ] Nombre de `to add / to change / to destroy` cohérent ?
- [ ] **Aucun `to destroy`** inattendu (surtout avec `-/+` = remplacement) ?
- [ ] Les `~` (modifications) touchent-ils des attributs sensibles (disque, réseau, `vm_id`) ?
- [ ] Les ressources avec `prevent_destroy` ne sont pas contournées ?
- [ ] Backend/state : le bon environnement (prod vs staging) ?
- [ ] Secrets : rien en clair dans le plan JSON (`terraform show -json tfplan | grep -i password`) ?
- [ ] Version des providers conforme au lock file ?

```bash
# Inspecter un plan en détail
terraform show tfplan                    # lisible humain
terraform show -json tfplan | jq '.resource_changes[] | {address, actions: .change.actions}'
```

> 💡 Astuce : commentez le résumé du plan dans la MR (bot ou copier-coller). Un plan non relu = un `apply` non relu.

---

## 59. Debug : TF_LOG et les traces

```bash
# Niveaux : TRACE, DEBUG, INFO, WARN, ERROR (défaut: désactivé)
export TF_LOG=DEBUG
export TF_LOG_PATH=/tmp/terraform-debug.log
terraform plan

# Tracer uniquement un provider (moins verbeux)
export TF_LOG_PROVIDER=DEBUG

# Désactiver
unset TF_LOG TF_LOG_PATH
```

