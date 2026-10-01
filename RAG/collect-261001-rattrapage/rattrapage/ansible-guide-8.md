---
id: collect-261001-rattrapage/rattrapage/ansible-guide-8
title: "Guide Ansible — Automatisation système en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "incident", "open source"]
source: docs/RAG/collect-261001-rattrapage/ansible_guide.md
source_anchor: ""
source_lines: [1817, 2066]
sha256: e836269074c61ae212c67295d6c4ae46ea7d13013076319d1a01a50dd3342311
---

# Guide Ansible — Automatisation système en production

```bash
# 1. Clé dédiée à Ansible (jamais la clé personnelle)
ssh-keygen -t ed25519 -f ~/.ssh/id_ed25519_ansible -C "ansible@control-node"

# 2. Déployer la clé sur le parc (une fois, avec mot de passe)
ssh-copy-id -i ~/.ssh/id_ed25519_ansible deploy@web-01

# 3. ssh_config du control node (~/.ssh/config)
Host *.example.lan
    User deploy
    IdentityFile ~/.ssh/id_ed25519_ansible
    ServerAliveInterval 30
```

Bastion dans l'inventaire (déjà vu section 11) ou global :

```ini
[all:vars]
ansible_ssh_common_args=-o ProxyJump=deploy@bastion.example.lan -o StrictHostKeyChecking=yes
```

Checklist SSH production :

- [ ] Clé ed25519 dédiée, `chmod 600`
- [ ] `ControlPersist` activé (section 7) pour la vitesse
- [ ] `host_key_checking = True` + `~/.ssh/known_hosts` rempli (via `ssh-keyscan` versionné)
- [ ] Port SSH non standard + `fail2ban` côté cibles (voir cas pratique 2)

## 60. Performance : pipelining, facts, stratégies

| Levier | Gain typique | Risque / condition |
|---|---|---|
| `pipelining = True` | ÷2 à ÷5 sur le temps | sudo sans `requiretty` |
| `ControlPersist` (section 7) | Connexions réutilisées | Aucun |
| `gather_facts: false` | −5 à 15 s par hôte | Uniquement si facts inutiles |
| `forks` adapté (section 50) | Linéaire jusqu'à saturation | Charge du control node |
| `strategy: free` | Hôtes rapides non bloqués | Logs entrelacés |
| `serial` bien dimensionné | Rolling updates sûrs | Plus lent par conception |
| Facts en cache (`fact_caching`) | Playbooks répétés | Redis ou JSON file |

Cache des facts pour les exécutions répétées :

```ini
[defaults]
fact_caching = jsonfile
fact_caching_connection = /tmp/ansible_facts_cache
fact_caching_timeout = 86400
```

## 61. Supervision : logs et callbacks

```ini
# ansible.cfg
[defaults]
log_path = ./logs/ansible.log
stdout_callback = yaml
# En cas d'échec : résumé dense des hôtes en échec
# (plugin community.general.log_plays ou callback splunk à configurer)
```

Plugins de callback utiles :

| Callback | Usage |
|---|---|
| `yaml` | Sortie lisible (défaut conseillé) |
| `dense` | Résumé compact en fin d'exécution |
| `timer` / `profile_tasks` | Temps par tâche (optimisation) |
| `logstash` / `splunk` | Envoi vers la supervision centralisée |
| `mail` | Notification email en fin de play |

```bash
# Activer le profiling ponctuellement
ANSIBLE_CALLBACKS_ENABLED=ansible.posix.profile_tasks ansible-playbook site.yml

# Journaliser en JSON pour ingestion (ELK/Loki)
ANSIBLE_STDOUT_CALLBACK=json ansible-playbook site.yml > run.json
```

Conservez les logs d'exécution 90 jours minimum : c'est la preuve d'audit
« qui a changé quoi, quand, sur quel serveur ».

## 62. ARA : enregistrer et visualiser les exécutions

[ARA](https://ara.recordsansible.org/) (Ansible Run Analysis) enregistre chaque
exécution dans une base SQLite/PostgreSQL avec interface web.

```bash
pipx inject ansible "ara[server]"

# ansible.cfg
[defaults]
callbacks_enabled = ara_default

# Lancer le serveur web ARA
ara-manage runserver 0.0.0.0:8000
```

Bénéfices : historique des `changed`, comparaison avant/après, recherche par hôte
ou par tâche, API REST pour la supervision. Alternative légère : le callback
`community.general.log_plays` qui écrit un fichier par hôte.

## 63. Sauvegarder les playbooks avec Git

```bash
cd ~/projets/ansible
git init
cat > .gitignore <<'EOF'
# Secrets et fichiers locaux
.vault_pass*
*vault_pass*
logs/
*.retry
collections/
roles_tiers/
# Inventaires sensibles éventuels
inventaires/*secret*
EOF
git add .
git commit -m "Socle Ansible initial"
```

Workflow d'équipe recommandé :

1. Une branche par changement (`git checkout -b feature/durcissement-ssh`).
2. `--check --diff` en local, puis revue de code (merge request).
3. `main` = ce qui tourne en production ; taggez les versions (`v2026.09.1`).
4. En cas d'incident : `git log --oneline -20` puis `git diff v2026.09.0..HEAD`
   pour identifier le changement fautif, rollback par `git revert`.

> **Règle absolue** : aucun secret en clair dans Git, même « temporairement ».
> L'historique Git n'oublie jamais (`git log -p` les retrouvera).

## 64. AWX / Ansible Controller : tour d'horizon

AWX (open source) et Ansible Controller (Red Hat, ex-Tower) ajoutent une
**interface web + API + ordonnanceur** au-dessus d'Ansible.

| Fonctionnalité | Apport vs CLI |
|---|---|
| Inventaires dynamiques | Synchronisés depuis cloud/CMDB |
| Credentials | Secrets chiffrés en base, jamais sur disque |
| Job templates | Playbooks paramétrables en un clic |
| Ordonnancement | Cron intégré avec historique |
| RBAC | Qui peut lancer quoi, sur quel inventaire |
| Workflows | Enchaîner plusieurs job templates |
| Notifications | Slack/email/webhook en fin de job |

Quand passer à AWX :

- [ ] Plusieurs opérateurs lancent des playbooks (besoin de traçabilité).
- [ ] Les exécutions doivent être planifiées (patching mensuel).
- [ ] Délégation à des non-experts (helpdesk relance un playbook via template).

Installation rapide (Docker Compose, labo) :

```bash
git clone https://github.com/ansible/awx.git
cd awx
# Suivre docs/awx-operator pour Kubernetes en production
```

> **Avertissement** : AWX en production se déploie via l'opérateur Kubernetes
> officiel, pas via un compose artisanal. Prévoyez PostgreSQL managé et des
> backups de la base AWX (c'est là que vivent les credentials).

## 65. Cas pratique 1 : déployer un serveur web complet (commenté)

Objectif : sur `web-01`, installer nginx + PHP-FPM, déployer un vhost, ouvrir le
pare-feu, vérifier le service. Idempotent et rejouable.

```yaml
---
- name: Déployer la stack web LEMP
  hosts: webservers
  become: true
  gather_facts: true

  vars:
    nom_domaine: intranet.example.lan
    php_version: "8.3"

  pre_tasks:
    # Garde-fou : on ne déploie que sur Debian/Ubuntu testés
    - name: Vérifier l'OS supporté
      ansible.builtin.assert:
        that:
          - ansible_facts['os_family'] == "Debian"
        fail_msg: "OS non supporté : {{ ansible_facts['distribution'] }}"

  tasks:
    # --- Installation ---
    - name: Installer nginx et PHP-FPM
      ansible.builtin.apt:
        name:
          - nginx
          - "php{{ php_version }}-fpm"
          - "php{{ php_version }}-mysql"
        state: present
        update_cache: true
        cache_valid_time: 3600
      tags: [install]

    # --- Configuration nginx ---
    - name: Déployer le vhost
      ansible.builtin.template:
        src: templates/vhost_lemp.conf.j2
        dest: /etc/nginx/sites-available/intranet.conf
        mode: "0644"
        validate: "nginx -t -c %s"   # jamais de nginx cassé en prod
      tags: [config]
      notify: Recharger nginx

    - name: Activer le site
      ansible.builtin.file:
        src: /etc/nginx/sites-available/intranet.conf
        dest: /etc/nginx/sites-enabled/intranet.conf
        state: link
      notify: Recharger nginx

    - name: Désactiver le site par défaut
      ansible.builtin.file:
        path: /etc/nginx/sites-enabled/default
        state: absent
      notify: Recharger nginx

    # --- Racine web ---
    - name: Créer la racine web
      ansible.builtin.file:
        path: /var/www/intranet
        state: directory
        owner: www-data
        group: www-data
        mode: "0755"

    - name: Déployer la page d'accueil
      ansible.builtin.template:
        src: templates/index.php.j2
        dest: /var/www/intranet/index.php
        owner: www-data
        group: www-data
        mode: "0644"

    # --- Pare-feu ---
    - name: Autoriser HTTP/HTTPS via UFW
      community.general.ufw:
        rule: allow
        port: "80,443"
        proto: tcp
      tags: [firewall]

