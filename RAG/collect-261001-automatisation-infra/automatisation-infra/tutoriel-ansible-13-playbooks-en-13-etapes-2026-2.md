---
id: collect-261001-automatisation-infra/automatisation-infra/tutoriel-ansible-13-playbooks-en-13-etapes-2026-2
title: "Mise à jour des paquets et installation de pipx"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-automatisation-infra/tutoriel-ansible-13-playbooks-en-13-etapes-2026.md
source_anchor: ""
source_lines: [33, 181]
sha256: 3f7dae519f06f04eca14e751a2918d86e7fc4682ce11c9b9b23148e102cc0053
---

# Mise à jour des paquets et installation de pipx

```
# Mise à jour des paquets et installation de pipx
sudo apt update && sudo apt install -y pipx python3-venv openssh-client
pipx ensurepath
# Installation du paquet Ansible community 13.x (inclut ansible-core 2.20)
pipx install --include-deps ansible
# Vérification de la version installée
ansible --version
ansible-playbook --version
ansible-galaxy collection list | head -20
```
La sortie attendue affiche `ansible [core 2.20.5]` avec le chemin vers la configuration et l’emplacement des collections. Si vous travaillez en environnement contrôlé sans accès Internet, téléchargez le wheel `ansible-13.6.0-py3-none-any.whl` sur PyPI depuis un poste connecté, puis installez-le hors-ligne avec `pipx install ansible-13.6.0-py3-none-any.whl`. Les collections requises peuvent être pré-empaquetées dans un tarball avec `ansible-galaxy collection download`.

### Vérifier la connectivité avec une cible

Avant tout playbook, validez que le contrôleur peut joindre une machine cible et y exécuter du Python distant. Le module `ping` d’Ansible n’est pas un ICMP : il ouvre une session SSH, copie un module Python sur la cible et exécute un test de présence. Si vous obtenez `SUCCESS => { "ping": "pong" }`, votre canal de transport est fonctionnel.

```
# Test rapide de connectivité avec une cible nommée web1
ansible -i web1.example.fr, all -m ping -u ansible --private-key=~/.ssh/ansible_ed25519
# Si la cible répond avec pong, vous pouvez exécuter un module ad-hoc
ansible -i web1.example.fr, all -m ansible.builtin.setup -u ansible
```
## Étape 2 : Structurer un projet Ansible selon les bonnes pratiques

Un projet Ansible mal structuré devient ingérable dès la dixième tâche. La documentation officielle recommande une arborescence claire qui sépare l’inventaire, les playbooks, les rôles, les variables et les fichiers chiffrés. Cette organisation est compatible avec l’**Ansible Automation Platform 2.6** et permet d’importer le projet dans un Execution Environment sans modification.

```
mon-projet-ansible/
├── ansible.cfg                 # Configuration locale du projet
├── requirements.yml            # Collections et rôles Galaxy
├── inventory/
│   ├── production/
│   │   ├── hosts.yml           # Inventaire YAML
│   │   ├── group_vars/
│   │   │   ├── all.yml
│   │   │   ├── webservers.yml
│   │   │   └── dbservers.yml
│   │   └── host_vars/
│   │       └── web1.yml
│   └── staging/
│       └── hosts.yml
├── playbooks/
│   ├── site.yml                # Orchestration globale
│   ├── webservers.yml
│   └── dbservers.yml
├── roles/
│   ├── common/
│   ├── nginx/
│   ├── mariadb/
│   └── php-fpm/
├── files/                      # Fichiers statiques à copier
├── templates/                  # Templates Jinja2
├── vault/
│   └── secrets.yml             # Secrets chiffrés
└── .ansible-lint               # Configuration du linter
```
Le fichier `ansible.cfg` placé à la racine du projet a priorité sur `/etc/ansible/ansible.cfg`. Il fixe l’inventaire par défaut, désactive la vérification de hostkey en première exécution (à réactiver après bootstrap) et active le forks parallèle à 50 pour exécuter les tâches sur 50 hôtes simultanément. Les paramètres ci-dessous représentent une configuration solide pour la production.

```
[defaults]
inventory = inventory/production/hosts.yml
roles_path = roles
collections_path = ./collections
host_key_checking = False
forks = 50
gathering = smart
fact_caching = jsonfile
fact_caching_connection = .ansible_cache
fact_caching_timeout = 7200
stdout_callback = yaml
callback_whitelist = profile_tasks, timer
log_path = .ansible.log
retry_files_enabled = False
interpreter_python = auto_silent
[ssh_connection]
pipelining = True
ssh_args = -o ControlMaster=auto -o ControlPersist=60s -o ServerAliveInterval=15
control_path = ~/.ansh-%%h-%%p-%%r
[privilege_escalation]
become = True
become_method = sudo
become_user = root
become_ask_pass = False
```
## Étape 3 : Construire un inventaire YAML statique et dynamique

L’inventaire est le cœur d’Ansible : il définit quels hôtes sont gérés, comment ils sont groupés et quelles variables s’appliquent à chacun. En 2026, le format YAML s’impose face au vieux format INI car il supporte naturellement la hiérarchie, les listes et les types complexes. Un inventaire bien conçu permet d’écrire un seul playbook qui s’adapte à plusieurs environnements (staging, production, DR) en changeant uniquement le fichier `hosts.yml`.

```
# inventory/production/hosts.yml
---
all:
  vars:
    ansible_user: ansible
    ansible_ssh_private_key_file: ~/.ssh/ansible_ed25519
    ansible_python_interpreter: /usr/bin/python3
  children:
    webservers:
      hosts:
        web1.example.fr:
          ansible_host: 10.10.1.11
          server_id: 1
        web2.example.fr:
          ansible_host: 10.10.1.12
          server_id: 2
      vars:
        nginx_worker_processes: 4
        php_version: "8.3"
    dbservers:
      hosts:
        db1.example.fr:
          ansible_host: 10.10.1.21
          mariadb_role: primary
      vars:
        mariadb_version: "11.4"
        mariadb_max_connections: 500
    paris_dc:
      children:
        webservers:
        dbservers:
```
Pour les environnements cloud, l’inventaire dynamique génère cette liste à la volée depuis l’API du fournisseur. La collection `amazon.aws` fournit le plugin `aws_ec2` qui interroge EC2 et groupe les instances par tag, région ou type. Un fichier `inventory/production/aws_ec2.yml` remplace le YAML statique et se rafraîchit à chaque exécution. La même logique s’applique avec `community.proxmox` pour Proxmox VE 9, `azure.azcollection` pour Azure ou `community.general.scaleway` pour Scaleway, prisé chez les clients soumis aux exigences de souveraineté française.

```
# inventory/production/aws_ec2.yml
---
plugin: amazon.aws.aws_ec2
regions:
  - eu-west-3
  - eu-west-1
filters:
  tag:Environment: production
  instance-state-name: running
keyed_groups:
  - prefix: tag
    key: tags.Role
  - prefix: az
    key: placement.availability_zone
hostnames:
  - tag:Name
  - private-ip-address
compose:
  ansible_host: private_ip_address
```
## Étape 4 : Écrire votre premier playbook idempotent

Un playbook est un fichier YAML qui décrit un état désiré. Ansible compare cet état à la réalité de chaque hôte et n’effectue que les actions nécessaires : c’est l’**idempotence**. Lancer le même playbook dix fois doit produire exactement le même résultat, sans déclencher de modifications inutiles. Ce comportement est garanti par les modules officiels de la collection `ansible.builtin`, qui gèrent automatiquement les états `present`, `absent` et `latest`.

