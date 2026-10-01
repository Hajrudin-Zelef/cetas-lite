---
id: collect-261001-rattrapage/rattrapage/terraform-guide-1
title: "Terraform — Guide complet : Infrastructure as Code en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["open source"]
source: docs/RAG/collect-261001-rattrapage/terraform_guide.md
source_anchor: ""
source_lines: [1, 183]
sha256: 58d2527f679280c8831f606acf072ed31209e97917cdf2773623346ad0696d6e
---

# Terraform — Guide complet : Infrastructure as Code en production

> **Public** : Zelef, chef de service systèmes & énergies, sysadmin orienté production.
> **Objectif** : maîtriser Terraform de l'installation au déploiement industriel, avec un focus sur Proxmox VE (provider bpg) et le cloud-init.
> **Versions couvertes** : Terraform ≥ 1.5 (vérifié avec 1.9/1.10), OpenTofu ≥ 1.8.
> **Règle d'or** : *on relit TOUJOURS le plan avant `apply`.*

---

## Sommaire

1. L'Infrastructure as Code : pourquoi Terraform
2. Déclaratif vs impératif
3. Idempotence : le pilier de la fiabilité
4. Installation de Terraform sur Debian/Ubuntu
5. Installation d'OpenTofu (alternative)
6. Configuration du shell et des outils
7. Vérification, versions et gestion multi-versions
8. Structure d'un projet Terraform
9. Le workflow de base : init → plan → apply
10. destroy : détruire proprement
11. La CLI Terraform : commandes essentielles
12. HCL : le langage — blocs, arguments, identifiants
13. HCL : types de valeurs et expressions
14. HCL : les fonctions intégrées les plus utiles
15. Variables : déclaration et typage
16. Variables : validation des entrées
17. Variables : les fichiers .tfvars et la précédence
18. Outputs : exposer des valeurs
19. Locals : factoriser la logique
20. Data sources : lire l'existant
21. count : répétition simple
22. for_each : répétition robuste par clé
23. Blocs dynamic : générer des sous-blocs
24. Conditions et boucles en expressions
25. Fonctions avancées : templatefile, fileset, jsonencode
26. terraform fmt et validate : l'hygiène du code
27. Providers : principe et installation
28. Version pinning : figer les versions
29. Le provider Proxmox bpg : installation et configuration
30. Les ressources : anatomie d'un bloc resource
31. Dépendances implicites et graphe
32. depends_on : dépendances explicites
33. Le bloc lifecycle : meta-arguments
34. prevent_destroy : le garde-fou anti-catastrophe
35. create_before_destroy : zéro interruption
36. ignore_changes : ignorer les dérives voulues
37. replace_triggered_by : forcer le remplacement
38. taint et -replace : marquer une ressource à recréer
39. Refresh et plan -refresh-only
40. Le state : pourquoi il existe
41. State local vs state distant
42. Backend S3 + verrouillage DynamoDB
43. Backend Terraform Cloud / HCP Terraform
44. Sauvegarde et protection du state
45. Manipulation du state : mv, rm, pull, push
46. Workspaces : un state par environnement
47. Séparation des environnements : dev / staging / prod
48. Modules : créer son premier module
49. Modules : utiliser un module (registry, git, local)
50. Bonnes pratiques de conception de modules
51. Import de ressources existantes (import block)
52. Provisioners : pourquoi les éviter
53. Alternatives aux provisioners (cloud-init, Ansible, packer)
54. Sécurité : jamais de secrets en clair
55. Chiffrer les tfvars : SOPS + age
56. Détection de dérive (drift detection)
57. CI/CD : pipeline GitLab de revue et d'apply
58. Revue de plan en équipe
59. Debug : TF_LOG et les traces
60. 10+ erreurs classiques et leurs solutions
61. Mise à jour des providers
62. Mise à jour de Terraform / OpenTofu
63. OpenTofu : différences avec Terraform
64. Migration Terraform → OpenTofu
65. Cas pratique 1 : parc de 10 VM Proxmox depuis du code
66. Cas pratique 2 : reprise après sinistre
67. Cas pratique 3 : réseau complet avec VLAN et cloud-init
68. Cas pratique 4 : cluster à 3 nœuds avec anti-affinité
69. Supervision et alerting autour de Terraform
70. Tester son code : terraform test et check blocks
71. Documenter : terraform-docs
72. Estimer les coûts : Infracost
73. Conventions de nommage et style d'équipe
74. Checklist de mise en production
75. Pense-bête de poche (cheat sheet)
76. Glossaire
77. Quiz : 10 questions (énoncés)
78. Quiz : réponses commentées
79. Pour aller plus loin
80. Annexe : fichiers de référence complets

> ⚠️ **Avertissement de version** : les exemples ont été écrits pour Terraform ≥ 1.5 / OpenTofu ≥ 1.8 et le provider `bpg/proxmox` ≥ 0.66. Vérifiez toujours la documentation du provider pour votre version exacte : les noms d'arguments changent parfois (ex. `proxmox_virtual_environment_vm`).

---

## 1. L'Infrastructure as Code : pourquoi Terraform

L'Infrastructure as Code (IaC) consiste à décrire son infrastructure (VM, réseaux, DNS, pare-feu…) dans des fichiers versionnés, au lieu de cliquer dans des interfaces ou de lancer des scripts ad hoc.

| Avant (manuel) | Après (Terraform) |
|---|---|
| Clics dans l'UI Proxmox, non reproductibles | Code versionné en Git, reproductible à l'identique |
| "Qui a créé cette VM ?" sans réponse | `git blame` répond en 10 secondes |
| Dérive de configuration invisible | `terraform plan` révèle toute dérive |
| Reprise après sinistre = réinstallation manuelle | `terraform apply` reconstruit le parc |

Terraform, créé par HashiCorp, est l'outil IaC déclaratif le plus répandu. Il gère le **cycle de vie complet** : création, mise à jour, suppression. Pour un chef de service systèmes, c'est l'assurance que le parc documenté = le parc réel.

---

## 2. Déclaratif vs impératif

- **Impératif** (Ansible ad hoc, scripts bash) : on décrit *comment* faire, étape par étape. Si une étape échoue à moitié, l'état est inconnu.
- **Déclaratif** (Terraform) : on décrit *l'état final souhaité*. Terraform calcule lui-même le chemin (créer, modifier, supprimer) pour y arriver.

```hcl
# Déclaratif : "je veux une VM nommée web-01 avec 4 vCPU"
resource "proxmox_virtual_environment_vm" "web_01" {
  name      = "web-01"
  node_name = "pve1"
  cpu {
    cores = 4
  }
}
# Terraform décide seul : créer la VM si absente, modifier si différente.
```

Avantage clé : **le code est la documentation vivante** de l'infrastructure.

---

## 3. Idempotence : le pilier de la fiabilité

Une opération idempotente peut être répétée sans effet de bord supplémentaire. `terraform apply` lancé 5 fois de suite ne crée qu'**une** VM.

```
1er apply  → 1 VM créée
2e apply   → "No changes. Your infrastructure matches the configuration."
3e apply   → idem
```

C'est ce qui rend Terraform sûr en production et en CI/CD : relancer un pipeline ne duplique rien. L'idempotence repose sur le **state** (section 40) : Terraform compare l'état désiré (code), l'état réel (fournisseur) et l'état mémorisé (state).

---

## 4. Installation de Terraform sur Debian/Ubuntu

Méthode officielle : dépôt APT HashiCorp (clé GPG vérifiée).

```bash
# 1. Prérequis
sudo apt update && sudo apt install -y gnupg software-properties-common curl

# 2. Clé GPG officielle HashiCorp
wget -O- https://apt.releases.hashicorp.com/gpg \
  | sudo gpg --dearmor -o /usr/share/keyrings/hashicorp-archive-keyring.gpg

# 3. Vérification de l'empreinte (doit afficher 798A EC65...)
gpg --no-default-keyring \
    --keyring /usr/share/keyrings/hashicorp-archive-keyring.gpg --fingerprint

# 4. Dépôt stable
echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] \
https://apt.releases.hashicorp.com $(lsb_release -cs) main" \
  | sudo tee /etc/apt/sources.list.d/hashicorp.list

# 5. Installation
sudo apt update && sudo apt install -y terraform

# 6. Vérification
terraform version
# Terraform v1.10.x on linux_amd64
```

> 💡 **Alternative sans dépôt** : télécharger le binaire sur `releases.hashicorp.com`, dézipper dans `/usr/local/bin`. Pratique sur serveur sans accès au dépôt.

---

## 5. Installation d'OpenTofu (alternative)

OpenTofu est le fork communautaire open source de Terraform (licence MPL-2.0), né après le changement de licence de Terraform (BSL) en 2023. Compatibilité quasi totale avec le langage HCL.

