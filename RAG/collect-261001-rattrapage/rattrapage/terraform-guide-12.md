---
id: collect-261001-rattrapage/rattrapage/terraform-guide-12
title: "Terraform — Guide complet : Infrastructure as Code en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "license", "open source"]
source: docs/RAG/collect-261001-rattrapage/terraform_guide.md
source_anchor: ""
source_lines: [2679, 2819]
sha256: 7ac897c3d2db2ded99804813a7204ee70bd9df77f6ff6a4d0b242bc572f6acd9
---

# Terraform — Guide complet : Infrastructure as Code en production

1. Quelle commande affiche les changements prévus **sans rien modifier** ?
2. Pourquoi faut-il relire le plan avant `terraform apply` en production ?
3. Quelle est la différence entre `count` et `for_each` pour créer 10 VM ?
4. Où sont stockés les secrets si on les met dans une variable sans précaution ? (2 endroits)
5. À quoi sert `terraform init` ? Citez 3 choses qu'il fait.
6. Que se passe-t-il si deux personnes lancent `terraform apply` en même temps avec un state local ?
7. Expliquez `prevent_destroy` et donnez un cas d'usage.
8. Comment importer une VM créée manuellement dans Proxmox sans la recréer ?
9. Pourquoi les provisioners `remote-exec` sont-ils déconseillés ? Citez 2 alternatives.
10. Quelle est la principale différence de licence entre Terraform et OpenTofu ?

---

## 78. Quiz : réponses commentées

1. **`terraform plan`** (idéalement `plan -out=tfplan`). C'est une opération de lecture seule : aucun changement n'est appliqué.
2. Parce que le plan est le **contrat d'exécution** : il liste créations, modifications et surtout **destructions**. Relire évite de détruire une BDD de prod par une coquille dans le code.
3. `count` indexe par position (0, 1, 2…) : retirer une VM du milieu décale les suivantes et peut provoquer des remplacements en cascade. `for_each` indexe par **clé** (`"web-01"`) : chaque VM a une identité stable, la suppression est chirurgicale.
4. Dans le **state** (`terraform.tfstate`, en clair dans les attributs) et potentiellement dans les **logs** (`TF_LOG=TRACE`, sortie `plan`/`apply`). D'où : backend chiffré + `sensitive = true` + jamais de secret en clair dans Git.
5. `terraform init` : (1) télécharge les **providers** requis, (2) télécharge les **modules**, (3) configure/initialise le **backend** (et migre le state si besoin).
6. **Corruption du state** : sans verrouillage (state locking), les deux apply écrivent en concurrence et le state devient incohérent. D'où l'obligation d'un backend distant avec verrou en équipe.
7. `lifecycle { prevent_destroy = true }` fait **échouer** tout plan qui voudrait détruire la ressource. Cas d'usage : BDD de production, NAS, toute ressource difficilement reconstructible.
8. Avec le bloc `import { to = ...; id = "..." }` (Terraform ≥ 1.5) ou `terraform import`, après avoir déclaré la `resource` correspondante dans le code, puis ajustement jusqu'à obtenir un plan sans diff non voulu.
9. Ils sont **hors modèle déclaratif** (pas de drift detection), fragiles (SSH/réseau), non idempotents, et stockent des secrets de connexion. Alternatives : **cloud-init** pour la config initiale, **Ansible** après l'apply pour la configuration applicative.
10. Terraform est sous **BSL** (Business Source License, usage restreint par HashiCorp) tandis qu'OpenTofu est sous **MPL-2.0**, une vraie licence open source gouvernée par la communauté (Linux Foundation).

---

## 79. Pour aller plus loin

| Ressource | Pourquoi |
|---|---|
| `registry.terraform.io/providers/bpg/proxmox` | Documentation officielle du provider Proxmox (arguments exacts par version) |
| `developer.hashicorp.com/terraform/docs` | Référence du langage et de la CLI |
| `opentofu.org/docs` | Documentation OpenTofu (chiffrement du state, `removed` blocks) |
| `cloudinit.readthedocs.io` | Référence cloud-init (modules, exemples) |
| `pve.proxmox.com/wiki` | API et concepts Proxmox VE (prérequis au provider) |
| `terraform-docs` | Génération de documentation des modules |
| `infracost` | Estimation de coûts sur plan |
| `checkov` / `tfsec` | Analyse statique de sécurité du code Terraform |
| `terragrunt` | Surcouche DRY pour gérer backend/variables sur N environnements |
| `packer` | Construction de templates Proxmox versionnés (avec le plugin Proxmox) |

**Feuille de route suggérée** :
1. Labo : déployer 3 VM avec le cas pratique §65 sur un Proxmox de test.
2. Mettre le state sur MinIO/S3 + verrou, brancher le drift check nightly.
3. Écrire le module `vm` d'équipe + pipeline GitLab (§57).
4. Exercice PRA : tout détruire et reconstruire depuis le code (§66).
5. Évaluer OpenTofu (§63-64) pour les nouveaux projets.

---

## 80. Annexe : fichiers de référence complets

### A. `backend.tf` — S3 + DynamoDB (prod)

```hcl
terraform {
  backend "s3" {
    bucket         = "infra-terraform-state-prod"
    key            = "proxmox/prod/terraform.tfstate"
    region         = "eu-west-3"
    encrypt        = true
    dynamodb_table = "terraform-state-locks"
  }
}
```

### B. `.gitignore` Terraform/OpenTofu

```gitignore
.terraform/
.tofu/
*.tfstate
*.tfstate.*
*.tfvars
!*.tfvars.example
.terraform.lock.hcl
override.tf
override.tf.json
*_override.tf
*_override.tf.json
crash.log
terraform.tfstate.backup
```

### C. `templates/cloud-init.tftpl` — template complet

```yaml
#cloud-config
hostname: ${hostname}
manage_etc_hosts: true
package_update: true
package_upgrade: false
packages:
  - qemu-guest-agent
  - curl
  - vim
  - htop
users:
  - name: ${admin_user}
    sudo: ALL=(ALL) NOPASSWD:ALL
    shell: /bin/bash
    ssh_authorized_keys:
      - ${ssh_key}
ntp:
  servers: [10.10.0.1, 1.1.1.1]
runcmd:
  - systemctl enable --now qemu-guest-agent
  - timedatectl set-timezone Europe/Paris
final_message: "Cloud-init termine pour ${hostname}"
```

### D. Script `drift-check.sh` complet (cron quotidien)

```bash
#!/bin/bash
# /usr/local/bin/terraform-drift-check.sh — à adapter par environnement
set -euo pipefail
ENV_DIR="/opt/infra/envs/prod"
ALERT_MAIL="infra@entreprise.lan"
cd "$ENV_DIR"
terraform init -input=false >/dev/null 2>&1
set +e
terraform plan -refresh-only -detailed-exitcode -out=/tmp/drift.plan >/tmp/drift.log 2>&1
code=$?
set -e
if [ "$code" -eq 0 ]; then
  echo "$(date -Iseconds) OK: aucune derive" | tee -a /var/log/terraform-drift.log
elif [ "$code" -eq 2 ]; then
  echo "$(date -Iseconds) ALERTE: derive detectee" | tee -a /var/log/terraform-drift.log
  mail -s "[Terraform] Derive detectee sur prod" "$ALERT_MAIL" < /tmp/drift.log
else
  echo "$(date -Iseconds) ERREUR terraform code=$code" | tee -a /var/log/terraform-drift.log
  exit 1
fi
```

---

*Fin du guide — bon courage, et rappelez-vous : **on relit toujours le plan avant `apply`**.* 🚀
