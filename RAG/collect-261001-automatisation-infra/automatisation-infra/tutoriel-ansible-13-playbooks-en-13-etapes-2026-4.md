---
id: collect-261001-automatisation-infra/automatisation-infra/tutoriel-ansible-13-playbooks-en-13-etapes-2026-4
title: "Mise à jour des paquets et installation de pipx"
domain: automatisation-infra
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-automatisation-infra/tutoriel-ansible-13-playbooks-en-13-etapes-2026.md
source_anchor: ""
source_lines: [333, 487]
sha256: 6d3cf73700e0ad26d2eaaeac02cd42e18f67e3aa9261da4961053c76df8187b1
---

# Mise à jour des paquets et installation de pipx

```
# Créer un fichier chiffré
ansible-vault create vault/secrets.yml
# Le contenu (saisi dans l'éditeur)
mariadb_root_password: "M0tDeP@sse-Tr3sFort-2026!"
api_key_stripe: "sk_live_51HZk..."
ssl_private_key: |
  -----BEGIN PRIVATE KEY-----
  MIIEvQIBADANBgkqhkiG9w0BAQEFAASC...
  -----END PRIVATE KEY-----
# Éditer un fichier déjà chiffré
ansible-vault edit vault/secrets.yml
# Voir le contenu en clair sans modifier
ansible-vault view vault/secrets.yml
# Re-chiffrer avec un nouveau mot de passe
ansible-vault rekey vault/secrets.yml
# Chiffrer une chaîne unique pour l'inclure dans un fichier YAML
ansible-vault encrypt_string 'B@seD3Donn33s' --name 'db_password'
```
En production, ne tapez jamais le mot de passe Vault au clavier. Stockez-le dans un fichier hors-Git pointé par `vault_password_file = ~/.ansible/vault_pass` dans `ansible.cfg`, ou intégrez Ansible Vault avec HashiCorp Vault ou AWS Secrets Manager via un script `vault_password_executable`. Les pipelines CI/CD GitHub Actions ou GitLab CI doivent injecter le mot de passe via une variable secrète chiffrée, jamais en clair dans `.gitlab-ci.yml`.

## Étape 8 : Orchestrer un déploiement complet avec site.yml

Le playbook `site.yml` est le point d’entrée qui orchestre l’ensemble du déploiement. Il invoque les rôles dans l’ordre logique (base de données avant applicatif, applicatif avant load balancer) et applique des stratégies différentes selon le groupe : déploiement en série pour les bases, déploiement parallèle pour les serveurs web sans état, déploiement progressif (canary) pour les applicatifs critiques.

```
# playbooks/site.yml
---
- name: Bootstrap commun à tous les hôtes
  hosts: all
  become: true
  roles:
    - common
- name: Configurer les serveurs de base de données
  hosts: dbservers
  become: true
  serial: 1                  # Un nœud à la fois
  roles:
    - mariadb
- name: Configurer les serveurs web
  hosts: webservers
  become: true
  strategy: free             # Chaque hôte avance à son rythme
  roles:
    - php-fpm
    - nginx
    - app
- name: Configurer le load balancer
  hosts: loadbalancers
  become: true
  roles:
    - haproxy
- name: Vérifications post-déploiement
  hosts: webservers
  tasks:
    - name: Vérifier que l'application répond en HTTP 200
      ansible.builtin.uri:
        url: "https://{{ inventory_hostname }}/health"
        status_code: 200
        validate_certs: true
        return_content: true
      register: health_check
      failed_when: "'OK' not in health_check.content"
```
Lancez le déploiement complet en passant le mot de passe Vault et un tag pour cibler une partie spécifique. La commande `--check` active le **mode dry-run** qui simule les changements sans les appliquer, et `--diff` affiche les différences ligne à ligne sur les fichiers de configuration. C’est l’arme absolue pour valider une modification avant production.

```
# Déploiement complet en production avec Vault
ansible-playbook playbooks/site.yml \
  --inventory inventory/production/hosts.yml \
  --vault-password-file ~/.ansible/vault_pass \
  --extra-vars "deploy_version=1.42.0"
# Simulation (dry-run) avec affichage des diffs
ansible-playbook playbooks/site.yml --check --diff
# Limiter à un sous-ensemble d'hôtes
ansible-playbook playbooks/site.yml --limit webservers
# Exécuter uniquement les tâches taguées "config"
ansible-playbook playbooks/site.yml --tags config
# Reprendre depuis la tâche qui a échoué la dernière fois
ansible-playbook playbooks/site.yml --start-at-task "Déployer les vhosts"
```
## Étape 9 : Installer des collections depuis Ansible Galaxy

Ansible Galaxy est le hub officiel de partage de collections et de rôles communautaires. Les collections regroupent modules, plugins et rôles autour d’un éditeur ou d’une technologie : `community.general` apporte plus de 400 modules génériques, `amazon.aws` couvre l’intégralité d’AWS, `community.docker` gère Docker et Docker Compose, `kubernetes.core` manipule Kubernetes. Les éditeurs tiers suivent le mouvement : Splunk a par exemple mis à jour en mars 2026 les playbooks Ansible fournis avec la version 26.3.0-71 de sa plateforme, en y ajoutant des bindings natifs pour accélérer le traitement des données ; Confluent a de son côté publié le 23 juin 2026 la collection Ansible 8.3.0, qui déploie Confluent Platform 8.3.0 sans rupture de compatibilité ; et Cisco documente désormais son propre modèle de playbook Ansible pour BroadWorks dans le guide Release 26, daté du 7 avril 2025. Côté outillage complémentaire, le générateur de diagrammes **ansible-playbook-grapher** a lui aussi progressé, avec sa version 2.9.0 publiée le 1er février 2025 et packagée pour Python 3.9 et 3.11 sur piwheels, bien pratique pour visualiser l’enchaînement des rôles d’un projet volumineux. Le fichier `requirements.yml` liste les dépendances du projet, garantissant la reproductibilité.

```
# requirements.yml
---
collections:
  - name: ansible.posix
    version: ">=2.0.0"
  - name: community.general
    version: ">=10.0.0"
  - name: community.mysql
    version: ">=3.13.0"
  - name: amazon.aws
    version: ">=9.0.0"
  - name: kubernetes.core
    version: ">=5.2.0"
roles:
  - name: geerlingguy.docker
    version: 7.5.0
  - src: https://github.com/MyOrg/role-monitoring.git
    scm: git
    version: v2.1.0
    name: monitoring
```
```
# Installation des collections et rôles dans le projet
ansible-galaxy install -r requirements.yml
# Installation isolée par projet (recommandée)
ansible-galaxy collection install -r requirements.yml -p ./collections
# Mise à jour forcée
ansible-galaxy collection install -r requirements.yml --force
# Lister les collections installées
ansible-galaxy collection list
```
## Étape 10 : Tester avec ansible-lint et molecule

Un rôle non testé est un rôle cassé qui s’ignore. **ansible-lint** analyse statiquement vos playbooks et rôles pour détecter les anti-patterns, les modules dépréciés et les violations des bonnes pratiques. **Molecule**, lui, exécute réellement vos rôles dans des conteneurs Docker ou des VM Vagrant éphémères, vérifie l’idempotence sur deux passages successifs et lance des tests d’intégration avec Testinfra ou Verifier Ansible.

```
# Installation des outils de test
pipx install ansible-lint
pipx install molecule
pipx inject molecule molecule-plugins[docker]
# Linter le projet entier
ansible-lint playbooks/site.yml
# Configuration .ansible-lint
---
profile: production           # Profil le plus strict
exclude_paths:
  - .cache/
  - collections/
skip_list:
  - yaml[line-length]         # Tolère les lignes longues dans les templates
warn_list:
  - experimental
# Initialiser Molecule pour un rôle
cd roles/nginx
molecule init scenario --driver-name docker
# Exécuter le test complet : create, converge, idempotence, verify, destroy
molecule test
# Itérer pendant le développement
molecule converge          # Applique le rôle
molecule verify            # Lance les vérifications
molecule login             # SSH dans le conteneur
molecule destroy           # Nettoie
```
Le test d’idempotence est particulièrement précieux : Molecule exécute le rôle deux fois et échoue si le second passage signale des changements. C’est la garantie que votre code respecte le contrat fondamental d’Ansible. Couplez ces tests à GitHub Actions ou GitLab CI pour les exécuter à chaque pull request, et vous obtenez un pipeline DevOps de niveau industriel.

## Étape 11 : Intégrer Ansible dans GitHub Actions

Automatiser l’exécution d’Ansible dans un pipeline CI/CD transforme l’outil en moteur GitOps : un push sur la branche `main` déclenche un déploiement, un commit sur `staging` rebuild l’environnement de pré-production, une pull request lance ansible-lint et molecule. La clé SSH du compte de service et le mot de passe Vault sont stockés dans GitHub Secrets, jamais dans le repo.

