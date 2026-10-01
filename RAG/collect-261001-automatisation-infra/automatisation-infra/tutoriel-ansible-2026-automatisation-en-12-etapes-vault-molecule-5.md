---
id: collect-261001-automatisation-infra/automatisation-infra/tutoriel-ansible-2026-automatisation-en-12-etapes-vault-molecule-5
title: "Créer un environnement virtuel Python dédié"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "distribution"]
source: docs/RAG/collect-261001-automatisation-infra/tutoriel-ansible-2026-automatisation-en-12-etapes-vault-molecule.md
source_anchor: ""
source_lines: [630, 796]
sha256: 5c9c4c1025c6ae0a7db028adfa216b32333954c3ba310a19872bb7e1274b201d
---

# Créer un environnement virtuel Python dédié

```
# playbooks/conditions-avancees.yml
---
- name: Démonstration des structures de contrôle avancées
  hosts: all
  become: true
  tasks:
    # Condition simple : exécuter selon la distribution
    - name: Installer les paquets (Debian/Ubuntu)
      ansible.builtin.apt:
        name: "{{ paquets_communs }}"
        state: present
      when: ansible_os_family == "Debian"
    - name: Installer les paquets (RedHat/CentOS)
      ansible.builtin.dnf:
        name: "{{ paquets_communs }}"
        state: present
      when: ansible_os_family == "RedHat"
    # Condition avec variable d'inventaire
    - name: Configurer Nginx comme serveur principal
      ansible.builtin.template:
        src: nginx-principal.conf.j2
        dest: /etc/nginx/conf.d/upstream.conf
      when: serveur_principal | default(false)
    # Boucle avec index et condition
    - name: Créer les utilisateurs avec des rôles différents
      ansible.builtin.user:
        name: "{{ item.nom }}"
        groups: "{{ item.groupes }}"
        shell: "{{ item.shell | default('/bin/bash') }}"
        state: present
      loop:
        - { nom: "dev01", groupes: "developers", shell: "/bin/zsh" }
        - { nom: "ops01", groupes: "sudo,docker" }
        - { nom: "staging", groupes: "www-data" }
      when: item.nom != "staging" or environnement == "staging"
    # Bloc avec gestion d'erreur (try/catch/finally)
    - name: Déployer l'application avec rollback
      block:
        - name: Sauvegarder la version actuelle
          ansible.builtin.copy:
            src: /var/www/app/
            dest: /var/www/app_backup/
            remote_src: true
        - name: Déployer la nouvelle version
          ansible.builtin.git:
            repo: "https://github.com/monprojet/app.git"
            dest: /var/www/app
            version: "{{ version_app | default('main') }}"
            force: true
        - name: Installer les dépendances
          ansible.builtin.command:
            cmd: composer install --no-dev --optimize-autoloader
            chdir: /var/www/app
      rescue:
        - name: Rollback – Restaurer la version précédente
          ansible.builtin.copy:
            src: /var/www/app_backup/
            dest: /var/www/app/
            remote_src: true
        - name: Notifier l'échec du déploiement
          ansible.builtin.debug:
            msg: "ERREUR : Déploiement échoué sur {{ inventory_hostname }}. Rollback effectué."
      always:
        - name: Nettoyer le répertoire de backup
          ansible.builtin.file:
            path: /var/www/app_backup
            state: absent
        - name: Vérifier que l'application répond
          ansible.builtin.uri:
            url: "http://localhost/health"
            status_code: 200
          register: sante_app
          ignore_errors: true
        - name: Afficher le statut de l'application
          ansible.builtin.debug:
            msg: "Application {{ 'opérationnelle' if sante_app.status == 200 else 'EN ERREUR' }}"
```
Le pattern `block/rescue/always` est l'équivalent Ansible du try/catch/finally. Il est indispensable pour les déploiements en production où un rollback automatique doit être déclenché en cas d'échec. Ce mécanisme est nettement plus fiable que la gestion manuelle des erreurs et constitue une bonne pratique que toute équipe DevOps devrait adopter.

## Étape 10 – Créer un Pipeline CI/CD avec Ansible

L'intégration d'Ansible dans un pipeline CI/CD automatise entièrement le cycle de déploiement, de la validation du code jusqu'à la mise en production. Cette étape transforme vos playbooks en workflows automatisés déclenchés par des événements Git. Voici comment intégrer Ansible avec GitHub Actions pour un déploiement continu.

```
# .github/workflows/deploy.yml
name: Déploiement Ansible
on:
  push:
    branches: [main]
  pull_request:
    branches: [main]
jobs:
  lint:
    name: Validation du code Ansible
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Installer les dépendances
        run: |
          python3 -m pip install ansible==13.5.0 ansible-lint yamllint
      - name: Valider la syntaxe YAML
        run: yamllint -d relaxed .
      - name: Linter Ansible
        run: ansible-lint playbooks/ roles/
      - name: Vérifier la syntaxe des playbooks
        run: |
          for playbook in playbooks/*.yml; do
            ansible-playbook "$playbook" --syntax-check
          done
  deploy-staging:
    name: Déployer en staging
    needs: lint
    runs-on: ubuntu-latest
    if: github.event_name == 'push'
    environment: staging
    steps:
      - uses: actions/checkout@v4
      - name: Configurer SSH
        run: |
          mkdir -p ~/.ssh
          echo "${{ secrets.SSH_PRIVATE_KEY }}" > ~/.ssh/id_ed25519
          chmod 600 ~/.ssh/id_ed25519
          ssh-keyscan -H ${{ secrets.STAGING_HOST }} >> ~/.ssh/known_hosts
      - name: Installer Ansible
        run: pip install ansible==13.5.0
      - name: Créer le fichier Vault
        run: echo "${{ secrets.VAULT_PASSWORD }}" > ~/.vault_pass
      - name: Déployer sur staging
        run: |
          ansible-playbook playbooks/deploiement-complet.yml \
            -i inventaire/staging.yml \
            --vault-password-file ~/.vault_pass \
            --extra-vars "environnement=staging version_app=${{ github.sha }}"
      - name: Test de fumée
        run: |
          curl -sf https://staging.exemple.fr/health || exit 1
  deploy-production:
    name: Déployer en production
    needs: deploy-staging
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/main'
    environment: production
    steps:
      - uses: actions/checkout@v4
      - name: Déployer en production (canary)
        run: |
          ansible-playbook playbooks/deploiement-complet.yml \
            -i inventaire/production.yml \
            --vault-password-file ~/.vault_pass \
            --limit web01 \
            --extra-vars "environnement=production"
      - name: Valider le canary
        run: curl -sf https://app.exemple.fr/health
      - name: Déployer sur tous les serveurs
        run: |
          ansible-playbook playbooks/deploiement-complet.yml \
            -i inventaire/production.yml \
            --vault-password-file ~/.vault_pass \
            --extra-vars "environnement=production"
```
Ce pipeline implémente un déploiement canary automatisé : la nouvelle version est d'abord déployée sur un seul serveur (`web01`), validée par un test de santé, puis propagée à l'ensemble de l'infrastructure. En cas d'échec du test canary, le déploiement s'arrête automatiquement, protégeant les utilisateurs d'une version défectueuse. C'est une stratégie éprouvée utilisée par les plus grandes entreprises technologiques.

## Étape 11 – Tester les Playbooks avec Molecule

Molecule est le framework de test officiel pour les rôles Ansible. Il permet de valider que vos rôles fonctionnent correctement dans un environnement isolé (conteneur Docker ou VM) avant de les déployer en production. Tester son infrastructure as code est aussi important que tester son code applicatif – une configuration défectueuse peut provoquer des pannes catastrophiques.

