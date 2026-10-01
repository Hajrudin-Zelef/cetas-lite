---
id: collect-261001-rattrapage/rattrapage/ansible-guide-1
title: "Guide Ansible — Automatisation système en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/ansible_guide.md
source_anchor: ""
source_lines: [1, 194]
sha256: 5004db0a2f858e28ad29801d0c9ddb6b2be1d8170131f8fbb167bfe39909a9ce
---

# Guide Ansible — Automatisation système en production

> **Public** : administrateurs systèmes, chefs de service systèmes & énergies, équipes d'exploitation.
> **Angle** : pratique, production, sans agent, idempotent.
> **Versions couvertes** : Ansible Core 2.15 → 2.19 (Debian 12/13, Ubuntu 22.04/24.04).
> **Avertissement de version** : la syntaxe `ansible.builtin.*` (FQCN) est recommandée depuis Ansible 2.10+ ;
> `gather_facts` remplace l'ancien `gather_subset: all` implicite ; le plugin d'inventaire Proxmox
> s'appelle `community.general.proxmox` et nécessite la collection `community.general`.
> Vérifiez toujours la version cible avec `ansible --version` avant d'appliquer un playbook.

---

## 1. Pourquoi Ansible et où il se positionne

Ansible est un moteur d'automatisation **sans agent** (*agentless*) : le nœud de contrôle pilote les
machines gérées en SSH (Linux/Unix) ou WinRM (Windows). Aucun démon à installer, aucun port
supplémentaire à ouvrir sur les cibles.

| Outil | Agent | Langage | Idempotent | Usage typique |
|---|---|---|---|---|
| Ansible | Non (SSH) | YAML + Jinja2 | Oui (modules) | Config management, déploiements, orchestration |
| Puppet | Oui | DSL Puppet | Oui | Grands parcs, modèle pull |
| Chef | Oui | Ruby | Oui | Idem, culture dev |
| SaltStack | Oui (minion) | YAML + Python | Oui | Temps réel, event-driven |
| Terraform | Non | HCL | Oui | Provisionnement infra (complémentaire, pas concurrent) |

Ansible excelle quand on veut **déclarer l'état désiré** : « ce paquet doit être installé »,
« ce fichier doit contenir ceci », « ce service doit tourner ». On le relance sans crainte :
c'est l'**idempotence**.

## 2. Concepts fondamentaux

- **Control node** : la machine (ou le conteneur) où Ansible est installé et d'où partent les exécutions.
  Exigences : Python 3.10+, SSH client, Ansible Core.
- **Managed nodes** : les machines gérées. Exigences minimales : un serveur SSH et Python 3
  (`/usr/bin/python3`). Pas d'agent Ansible.
- **Inventory** : la liste des nœuds gérés, leurs groupes et leurs variables.
- **Module** : l'unité d'action (ex. `ansible.builtin.apt`, `ansible.builtin.copy`).
  Chaque module est **idempotent par conception** : il ne change l'état que si nécessaire.
- **Task** : un appel à un module avec des paramètres, exécuté sur un ensemble d'hôtes.
- **Play** : un ensemble de tâches appliquées à un groupe d'hôtes.
- **Playbook** : un fichier YAML contenant un ou plusieurs plays. C'est le livrable versionné en Git.
- **Idempotence** : propriété selon laquelle exécuter N fois une tâche produit le même état final
  qu'une seule exécution. C'est ce qui rend Ansible sûr en production.

```text
Control node ──SSH──> managed node 1 (web-01)
             ──SSH──> managed node 2 (web-02)
             ──SSH──> managed node 3 (db-01)
```

## 3. Architecture et fonctionnement interne

1. Ansible lit l'inventaire et construit la liste des hôtes cibles.
2. Pour chaque tâche, il génère un petit script Python (le *module*) et le transfère via SFTP/SCP.
3. Le script s'exécute sur le nœud géré avec les privilèges demandés (`become`).
4. Le résultat revient en JSON ; Ansible décide `changed` / `ok` / `failed`.
5. Les **handlers** notifiés s'exécutent en fin de play, une seule fois.

Points clés pour un sysadmin :

- Le transfert utilise SFTP par défaut (`scp_if_ssh` dans les vieilles versions).
- Le **pipelining** SSH (`pipelining = True` dans `ansible.cfg`) évite le transfert de fichier
  et accélère fortement les exécutions, mais exige `requiretty` désactivé côté sudo.
- Les **facts** sont collectés une fois par hôte (sauf `gather_facts: false`).

## 4. Installation sur Debian / Ubuntu via APT

```bash
# Debian 12/13 et Ubuntu 22.04/24.04 : paquet officiel (version parfois en retard)
sudo apt update
sudo apt install -y ansible

# Vérification
ansible --version
ansible-community --version 2>/dev/null || true
```

Le paquet Debian/Ubuntu fournit `ansible-core` + une sélection de collections.
Pour la dernière version stable, préférez **pipx** (section suivante).

Pré-requis Python sur les nœuds gérés (Debian/Ubuntu récents l'ont déjà) :

```bash
# À exécuter sur chaque nœud géré si Python manque
sudo apt update && sudo apt install -y python3
```

## 5. Installation via pipx (méthode recommandée)

`pipx` isole chaque outil Python dans son propre environnement virtuel : pas de conflit
avec le Python système, mises à jour propres.

```bash
# 1. Installer pipx
sudo apt update
sudo apt install -y pipx python3-venv
pipx ensurepath
# Rouvrir le shell ou : source ~/.bashrc

# 2. Installer Ansible
pipx install --include-deps ansible

# 3. Vérifier
ansible --version
which ansible   # ~/.local/bin/ansible

# 4. Mettre à jour plus tard
pipx upgrade ansible
```

Avantages en production : version figée par environnement, `pipx list` pour l'audit,
rollback simple (`pipx install ansible==9.5.1`).

## 6. Vérification de l'installation et versions

```bash
ansible --version
# ansible [core 2.17.3]
#   config file = /etc/ansible/ansible.cfg
#   python version = 3.12.x ...

# Lister les collections installées
ansible-galaxy collection list

# Lister les modules d'une collection
ansible-doc -l | grep -i apt

# Aide détaillée d'un module (à lire avant usage en prod)
ansible-doc ansible.builtin.apt

# Tester la connectivité vers l'inventaire
ansible all -i inventaire.ini -m ansible.builtin.ping
```

> **Avertissement de version** : `ansible` (paquet communautaire) ≠ `ansible-core`.
> Le premier embarque des dizaines de collections ; le second est le moteur seul.
> En production, partez d'`ansible-core` et n'ajoutez que les collections nécessaires.

## 7. Configuration : ansible.cfg complet et commenté

Ansible lit la configuration dans cet ordre : `ANSIBLE_CONFIG` (variable d'env) >
`./ansible.cfg` (répertoire courant) > `~/.ansible.cfg` > `/etc/ansible/ansible.cfg`.

```ini
# ~/projets/ansible/ansible.cfg — à poser à la racine du projet
[defaults]
# Inventaire par défaut du projet
inventory = ./inventaires/production.ini
# Rôles et collections versionnés avec le projet
roles_path = ./roles:./roles_tiers
collections_path = ./collections
# Utilisateur SSH par défaut (surchargé par l'inventaire si besoin)
remote_user = deploy
# Clé privée par défaut
private_key_file = ~/.ssh/id_ed25519_ansible
# Parallélisme : nombre de forks simultanés (défaut 5, trop faible)
forks = 30
# Pipelining : accélère SSH, exige sudo sans requiretty
pipelining = True
# Ne pas vérifier les clés d'hôtes inconnus en labo UNIQUEMENT
# En production : host_key_checking = True + ssh_known_hosts
host_key_checking = True
# Timeout SSH en secondes
timeout = 30
# Callback lisible (yaml) au lieu du pavé JSON une-ligne
stdout_callback = yaml
# Journaliser chaque exécution (supervision)
log_path = ./logs/ansible.log
# Désactiver le warning sur les collections (bruit en CI)
deprecation_warnings = False
# Interdire les plugins non signés ? (durcissement)
# allow_unsafe_lookups = False

[privilege_escalation]
# become: sudo par défaut, mot de passe demandé si become_ask_pass
become = True
become_method = sudo
become_user = root
become_ask_pass = False

[ssh_connection]
# Multiplexage SSH : réutilise les connexions (gros gain de vitesse)
ssh_args = -o ControlMaster=auto -o ControlPersist=60s -o ServerAliveInterval=30
# Nombre de tentatives avant échec
retries = 3

[colors]
# Couleurs lisibles même sur fond clair
changed = yellow
```

Checklist de revue d'un `ansible.cfg` en production :

