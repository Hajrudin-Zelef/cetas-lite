---
id: collect-261001-automatisation-infra/automatisation-infra/tutoriel-ansible-13-playbooks-en-13-etapes-2026-5
title: "Mise à jour des paquets et installation de pipx"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-automatisation-infra/tutoriel-ansible-13-playbooks-en-13-etapes-2026.md
source_anchor: ""
source_lines: [488, 638]
sha256: 1f8c6c6084f8abe0e90f9012bfc4ef65f7f299c04f5e51e317d60e93631f1d08
---

# Mise à jour des paquets et installation de pipx

```
# .github/workflows/ansible-deploy.yml
name: Ansible Deploy
on:
  push:
    branches: [main]
  pull_request:
    branches: [main]
jobs:
  lint:
    runs-on: ubuntu-24.04
    steps:
      - uses: actions/checkout@v5
      - uses: actions/setup-python@v6
        with:
          python-version: "3.13"
      - name: Installer ansible-lint
        run: pip install ansible==13.6.0 ansible-lint
      - name: Lancer le linter
        run: ansible-lint playbooks/
  deploy:
    needs: lint
    if: github.ref == 'refs/heads/main'
    runs-on: ubuntu-24.04
    steps:
      - uses: actions/checkout@v5
      - uses: actions/setup-python@v6
        with:
          python-version: "3.13"
      - name: Installer Ansible
        run: pip install ansible==13.6.0
      - name: Configurer la clé SSH
        env:
          SSH_KEY: ${{ secrets.ANSIBLE_SSH_KEY }}
        run: |
          mkdir -p ~/.ssh
          echo "$SSH_KEY" > ~/.ssh/id_ed25519
          chmod 600 ~/.ssh/id_ed25519
          ssh-keyscan -H ${{ secrets.PROD_HOSTS }} >> ~/.ssh/known_hosts
      - name: Écrire le mot de passe Vault
        run: echo "${{ secrets.VAULT_PASSWORD }}" > .vault_pass
      - name: Installer les collections
        run: ansible-galaxy install -r requirements.yml
      - name: Déployer
        run: |
          ansible-playbook playbooks/site.yml \
            -i inventory/production/hosts.yml \
            --vault-password-file .vault_pass
```
## Étape 12 : Optimiser les performances sur 100+ hôtes

À petite échelle, Ansible est rapide. À grande échelle, sa sérialisation par défaut peut faire exploser le temps d’exécution. Cinq leviers permettent d’accélérer drastiquement vos playbooks. La **parallélisation** via `forks` détermine combien d’hôtes sont traités en simultané : la valeur par défaut de 5 doit passer à 50 ou 100. Le **pipelining SSH** évite l’écriture de fichiers temporaires sur la cible. Le **fact caching** stocke les facts collectés pendant 2 heures pour ne pas les regénérer. Le **strategy free** autorise chaque hôte à avancer indépendamment des autres. Enfin, l’**async** exécute les tâches longues en arrière-plan.

| Optimisation | Gain typique | Mise en œuvre | Risque | 
|---|---|---|---|
| forks = 50 | ~10x sur 100 hôtes | ansible.cfg [defaults] | Charge réseau accrue | 
| pipelining = True | 30-50% | ansible.cfg [ssh_connection] | Incompatible avec sudo requiretty | 
| SSH ControlPersist | 20-40% | ssh_args | Aucun | 
| fact_caching jsonfile | 15-25% | ansible.cfg [defaults] | Facts potentiellement obsolètes | 
| strategy: free | 20-60% | Au niveau du play | Ordre d’exécution non garanti | 
| async + poll: 0 | Variable | Au niveau de la tâche | Gestion d’erreur complexe | 
| gather_subset: min | 10-15% | Au niveau du play | Facts limités | 

```
# Exemple de tâche async pour un long apt upgrade
- name: Mise à jour majeure du système (peut durer 15 min)
  ansible.builtin.apt:
    upgrade: dist
  async: 1800        # Tue la tâche après 30 min
  poll: 0            # Lance et passe à la suivante immédiatement
  register: upgrade_job
- name: Attendre la fin de la mise à jour
  ansible.builtin.async_status:
    jid: "{{ upgrade_job.ansible_job_id }}"
  register: upgrade_result
  until: upgrade_result.finished
  retries: 60
  delay: 30
```
## Étape 13 : Projet complet — déployer une stack LAMP en 4 minutes

Pour conclure, voici le projet final : un déploiement complet de stack LAMP (Linux + Nginx + MariaDB + PHP-FPM) sur 3 hôtes Ubuntu 24.04, avec firewall, SSL Let’s Encrypt, monitoring par Node Exporter et secrets vaultés. Le tout s’exécute en 3 minutes 47 secondes sur des VM 2vCPU/2Go en local et permet de servir une application PHP en production. Le code complet est disponible ci-dessous, prêt à être adapté à vos besoins.

```
# playbooks/lamp-stack.yml
---
- name: Stack LAMP complète sur Ubuntu 24.04
  hosts: lamp_cluster
  become: true
  vars_files:
    - ../vault/secrets.yml
  vars:
    app_name: monapp
    app_domain: monapp.example.fr
    php_version: "8.3"
    mariadb_version: "11.4"
  pre_tasks:
    - name: Vérifier la distribution
      ansible.builtin.assert:
        that:
          - ansible_distribution == "Ubuntu"
          - ansible_distribution_version is version("24.04", ">=")
        fail_msg: "Ubuntu 24.04 LTS minimum requis"
  roles:
    - role: common
      tags: [common, base]
    - role: mariadb
      when: "'dbservers' in group_names"
      tags: [db, database]
    - role: php-fpm
      when: "'webservers' in group_names"
      tags: [php, app]
    - role: nginx
      when: "'webservers' in group_names"
      tags: [web, nginx]
    - role: certbot
      when: "'webservers' in group_names"
      tags: [ssl, security]
    - role: node_exporter
      tags: [monitoring]
  post_tasks:
    - name: Test final HTTP 200
      ansible.builtin.uri:
        url: "https://{{ app_domain }}"
        status_code: 200
      delegate_to: localhost
      run_once: true
```
Lancez le déploiement complet avec la commande ci-dessous. Sur 3 VM neuves, l’exécution typique observée en environnement de test produit la sortie suivante : 87 tâches exécutées sur 3 hôtes, dont 64 ayant déclenché un changement, en 3 minutes 47 secondes. Aux passages suivants, l’idempotence ramène ce chiffre à 0 changement et le runtime à environ 12 secondes (essentiellement collecte des facts et vérifications).

```
# Exécution
ansible-playbook playbooks/lamp-stack.yml \
  -i inventory/production/hosts.yml \
  --vault-password-file ~/.ansible/vault_pass \
  --diff
# Sortie attendue (premier passage)
PLAY RECAP ********************************************************
db1.example.fr   : ok=24 changed=18 unreachable=0 failed=0 skipped=4
web1.example.fr  : ok=31 changed=23 unreachable=0 failed=0 skipped=2
web2.example.fr  : ok=31 changed=23 unreachable=0 failed=0 skipped=2
Playbook run took 0 days, 0 hours, 3 minutes, 47 seconds
```
## Pièges courants à éviter avec Ansible

Cinq pièges reviennent systématiquement chez les équipes qui adoptent Ansible. Le premier consiste à utiliser `command` ou `shell` à la place des modules dédiés : ces commandes brutes ne sont pas idempotentes et marquent toujours `changed=true`. Préférez `ansible.builtin.apt`, `ansible.builtin.service` ou `ansible.builtin.lineinfile`. Le second est l’oubli de `become: true` sur les tâches root, qui produit silencieusement des erreurs de permission.

Le troisième piège est le hardcoding de mots de passe dans `group_vars/all.yml` : tout fichier sensible doit être chiffré par Vault, sans exception. Le quatrième concerne les **handlers** qui ne sont pas notifiés : un handler n’est exécuté que si la tâche qui le notifie a renvoyé `changed`. Vérifiez vos noms de handlers à l’identique. Enfin, ignorer la sortie `--check --diff` avant production est la première cause d’incidents : 40 % des erreurs de déploiement sont détectables en mode dry-run.

## Dépannage et erreurs courantes

