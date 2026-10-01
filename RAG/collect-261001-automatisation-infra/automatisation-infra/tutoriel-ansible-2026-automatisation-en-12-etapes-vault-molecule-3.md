---
id: collect-261001-automatisation-infra/automatisation-infra/tutoriel-ansible-2026-automatisation-en-12-etapes-vault-molecule-3
title: "Créer un environnement virtuel Python dédié"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "memory"]
source: docs/RAG/collect-261001-automatisation-infra/tutoriel-ansible-2026-automatisation-en-12-etapes-vault-molecule.md
source_anchor: ""
source_lines: [248, 407]
sha256: c8a6b21f4541a127a3eeaf4095649082bb7404952591ee094bf77f9074e7dcc7
---

# Créer un environnement virtuel Python dédié

```
# Lancer le playbook en mode vérification (dry-run)
ansible-playbook playbooks/securite-base.yml --check --diff
# Exécuter réellement le playbook
ansible-playbook playbooks/securite-base.yml
# Résultat attendu :
# PLAY [Configuration de sécurité de base des serveurs] ************************
# 
# TASK [Gathering Facts] ********************************************************
# ok: [web01]
# ok: [web02]
# ok: [db01]
# ok: [monitor01]
# 
# TASK [Mettre à jour le cache APT et les paquets] ******************************
# changed: [web01]
# changed: [web02]
# changed: [db01]
# changed: [monitor01]
# 
# TASK [Installer les paquets de sécurité] **************************************
# changed: [web01]
# changed: [web02]
# changed: [db01]
# changed: [monitor01]
# 
# PLAY RECAP ********************************************************************
# web01    : ok=8  changed=6  unreachable=0  failed=0  skipped=0
# web02    : ok=8  changed=6  unreachable=0  failed=0  skipped=0
# db01     : ok=8  changed=6  unreachable=0  failed=0  skipped=0
# monitor01: ok=8  changed=6  unreachable=0  failed=0  skipped=0
```
Les **handlers** sont des tâches spéciales qui ne s'exécutent que lorsqu'elles sont notifiées par une tâche qui a effectué un changement. Dans notre exemple, SSH ne sera redémarré que si la configuration a réellement été modifiée. Cette mécanique évite les redémarrages inutiles et garantit la stabilité du système pendant les déploiements.

## Étape 5 – Organiser les Variables avec group_vars et host_vars

La gestion des variables est un aspect fondamental de tout projet Ansible professionnel. Au lieu de coder les valeurs en dur dans les playbooks, Ansible propose une hiérarchie de variables qui permet d'adapter le comportement à chaque groupe de serveurs ou à chaque hôte individuel. Les répertoires `group_vars` et `host_vars` suivent une convention de nommage stricte qui correspond aux groupes et aux hôtes définis dans l'inventaire.

```
# Variables partagées par tous les serveurs
cat > group_vars/all.yml << 'EOF'
---
# Configuration commune à tous les serveurs
ntp_serveur: "fr.pool.ntp.org"
timezone: "Europe/Paris"
locale: "fr_FR.UTF-8"
# Utilisateur de déploiement
utilisateur_deploy: deploy
cle_ssh_deploy: "ssh-ed25519 AAAA... deploy@control"
# Paquets communs à installer sur tous les serveurs
paquets_communs:
  - curl
  - wget
  - vim
  - htop
  - git
  - net-tools
  - python3-pip
# Configuration des logs
log_retention_jours: 30
rsyslog_serveur: "monitor01"
EOF
# Variables spécifiques aux serveurs web
cat > group_vars/serveurs_web.yml << 'EOF'
---
# Configuration Apache/Nginx
serveur_web: nginx
nginx_worker_processes: auto
nginx_worker_connections: 1024
nginx_keepalive_timeout: 65
# Configuration PHP
php_version: "8.3"
php_memory_limit: "256M"
php_upload_max_filesize: "64M"
php_max_execution_time: 30
# Domaines virtuels
vhosts:
  - nom: "app.exemple.fr"
    racine: "/var/www/app"
    ssl: true
  - nom: "api.exemple.fr"
    racine: "/var/www/api"
    ssl: true
EOF
# Variables spécifiques aux serveurs de base de données
cat > group_vars/serveurs_bdd.yml << 'EOF'
---
# Configuration MySQL/MariaDB
mysql_root_password: "{{ vault_mysql_root_password }}"
mysql_port: 3306
mysql_bind_address: "0.0.0.0"
mysql_max_connections: 150
mysql_innodb_buffer_pool_size: "1G"
# Bases de données à créer
mysql_bases:
  - nom: app_production
    encoding: utf8mb4
    collation: utf8mb4_unicode_ci
  - nom: app_staging
    encoding: utf8mb4
    collation: utf8mb4_unicode_ci
# Utilisateurs MySQL
mysql_utilisateurs:
  - nom: app_user
    mot_de_passe: "{{ vault_app_db_password }}"
    privileges: "app_production.*:ALL"
    hote: "192.168.1.%"
EOF
# Variables spécifiques à un hôte
cat > host_vars/web01.yml << 'EOF'
---
# web01 est le serveur principal
serveur_principal: true
nginx_worker_connections: 2048
EOF
```
La hiérarchie de priorité des variables Ansible suit un ordre précis : les variables en ligne de commande (`-e`) ont la priorité la plus haute, suivies des variables de tâche, puis des variables de bloc, de rôle, de playbook, de `host_vars`, de `group_vars`, et enfin de l'inventaire. Comprendre cette hiérarchie est essentiel pour éviter les comportements inattendus dans les projets complexes.

Notez l'utilisation de `{{ vault_mysql_root_password }}` dans les variables de base de données. Ces valeurs sensibles sont stockées de manière chiffrée avec Ansible Vault, que nous configurerons à l'étape suivante. Ne stockez jamais de mots de passe en clair dans vos fichiers de variables – c'est une règle de sécurité fondamentale que trop de projets négligent.

## Étape 6 – Sécuriser les Secrets avec Ansible Vault

Ansible Vault permet de chiffrer les données sensibles (mots de passe, clés API, certificats) directement dans votre dépôt Git. Les fichiers chiffrés sont déchiffrés automatiquement lors de l'exécution des playbooks, sans exposer les secrets en clair. Cette fonctionnalité est indispensable pour les équipes qui versionnent leur infrastructure as code tout en respectant les bonnes pratiques de sécurité.

```
# Créer un fichier de secrets chiffré
ansible-vault create group_vars/serveurs_bdd/vault.yml
# Entrez le mot de passe de chiffrement quand demandé
# Contenu du fichier vault.yml (édité dans votre éditeur par défaut)
---
vault_mysql_root_password: "MonMotDePasseSecurise2026!"
vault_app_db_password: "AppUserP@ssw0rd2026"
vault_backup_encryption_key: "cle-de-chiffrement-backup-32chars"
# Chiffrer un fichier existant
ansible-vault encrypt fichiers/certificat-ssl.pem
# Éditer un fichier chiffré
ansible-vault edit group_vars/serveurs_bdd/vault.yml
# Voir le contenu déchiffré sans modifier
ansible-vault view group_vars/serveurs_bdd/vault.yml
# Exécuter un playbook avec Vault
ansible-playbook playbooks/deploiement.yml --ask-vault-pass
# Ou utiliser un fichier de mot de passe (pour CI/CD)
echo "MonMotDePasseVault" > ~/.vault_pass
chmod 600 ~/.vault_pass
ansible-playbook playbooks/deploiement.yml --vault-password-file ~/.vault_pass
# Chiffrer une seule variable (inline)
ansible-vault encrypt_string 'MotDePasseSecret' --name 'ma_variable_secrete'
# Résultat :
# ma_variable_secrete: !vault |
#   $ANSIBLE_VAULT;1.1;AES256
#   62313365396662343062653031623164...
```
**Bonne pratique** : utilisez des identifiants Vault multiples pour séparer les secrets par environnement. Par exemple, un identifiant `dev` pour les secrets de développement et `prod` pour la production. Ajoutez la ligne `vault_password_file = ~/.vault_pass` dans votre `ansible.cfg` pour éviter de taper le mot de passe à chaque exécution. N'oubliez pas d'ajouter `.vault_pass` à votre `.gitignore` pour ne jamais committer le fichier de mot de passe en clair.

Dans les pipelines CI/CD comme GitHub Actions, stockez le mot de passe Vault en tant que secret de l'environnement et injectez-le au moment de l'exécution. Cette approche garantit que les secrets ne sont jamais exposés dans les logs de build ni dans l'historique Git.

## Étape 7 – Structurer le Code avec les Rôles Ansible

Les rôles Ansible permettent d'organiser le code en composants réutilisables et modulaires. Un rôle encapsule les tâches, les variables, les templates, les fichiers et les handlers liés à une fonctionnalité spécifique. Cette structure favorise la réutilisation du code entre projets et simplifie la maintenance à long terme. Ansible Galaxy, le dépôt communautaire, propose des milliers de rôles prêts à l'emploi – à noter que la feuille de route du projet a fixé le feature freeze de la collection 14.0 à mars 2026 (Ansible Community Documentation), un jalon à surveiller si vos rôles dépendent de modules récents.

