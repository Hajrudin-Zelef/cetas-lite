---
id: collect-261001-rattrapage/rattrapage/ansible-guide-3
title: "Guide Ansible — Automatisation système en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-rattrapage/ansible_guide.md
source_anchor: ""
source_lines: [464, 746]
sha256: 74a1f99080d008cf0b4c64fb88ea16735aa49737e53b212000c3fcb65be865a1
---

# 2. Version d'OS de tout le parc
ansible all -m ansible.builtin.command -a "cat /etc/os-release" | grep -E "SUCCESS|PRETTY"

# 3. Copier un fichier vers tous les hôtes
ansible all -m ansible.builtin.copy -a "src=/tmp/motd dest=/etc/motd mode=0644" -b

# 4. Créer un utilisateur d'astreinte partout
ansible all -b -m ansible.builtin.user -a "name=astreinte groups=sudo shell=/bin/bash"

# 5. Lancer une mise à jour sèche (check mode)
ansible all -b -m ansible.builtin.apt -a "upgrade=dist update_cache=yes" --check --diff

# 6. Vérifier qu'un port écoute
ansible webservers -m ansible.builtin.wait_for -a "port=80 timeout=5"

# 7. Collecter les facts d'un hôte et filtrer
ansible web-01 -m ansible.builtin.setup -a "filter=ansible_distribution*"

# 8. Exécuter une tâche en arrière-plan longue (async)
ansible all -b -m ansible.builtin.command -a "apt-get upgrade -y" -B 3600 -P 60

# 9. Tuer un processus bloqué par nom
ansible all -b -m ansible.builtin.command -a "pkill -f mon_processus_hs" --check

# 10. Tester la résolution DNS interne
ansible all -m ansible.builtin.command -a "getent hosts nas-01.example.lan"
```

## 18. Anatomie d'un playbook

```yaml
---
# playbooks/exemple_anatomie.yml
- name: Exemple d'anatomie d'un play           # 1. Nom du play (lisible dans les logs)
  hosts: webservers                            # 2. Cible : groupe d'inventaire
  become: true                                 # 3. Élévation de privilèges
  gather_facts: true                           # 4. Collecte des facts (défaut true)

  vars:                                        # 5. Variables du play
    http_port: 80

  pre_tasks:                                   # 6. Avant les rôles/tâches
    - name: Vérifier la connectivité
      ansible.builtin.ping:

  roles:                                       # 7. Rôles (optionnel)
    - role: mon_role_nginx

  tasks:                                       # 8. Tâches dans l'ordre
    - name: Installer nginx
      ansible.builtin.apt:
        name: nginx
        state: present
      notify: Redémarrer nginx                 # 9. Déclenche le handler si changed

  handlers:                                    # 10. Exécutés en fin de play, une fois
    - name: Redémarrer nginx
      ansible.builtin.service:
        name: nginx
        state: restarted

  post_tasks:                                  # 11. Après les handlers
    - name: Vérifier que le port écoute
      ansible.builtin.wait_for:
        port: "{{ http_port }}"
```

## 19. Playbook complet commenté : socle serveur Debian

```yaml
---
- name: Socle commun tous serveurs Debian/Ubuntu
  hosts: all
  become: true
  gather_facts: true

  vars:
    packages_socle:
      - vim
      - htop
      - curl
      - net-tools
      - unattended-upgrades

  tasks:
    - name: Mettre à jour le cache APT (Debian/Ubuntu uniquement)
      ansible.builtin.apt:
        update_cache: true
        cache_valid_time: 3600
      when: ansible_facts['os_family'] == "Debian"

    - name: Installer les paquets du socle
      ansible.builtin.apt:
        name: "{{ packages_socle }}"
        state: present

    - name: Déployer le MOTD d'entreprise
      ansible.builtin.copy:
        src: ../files/motd
        dest: /etc/motd
        owner: root
        group: root
        mode: "0644"

    - name: Activer les mises à jour automatiques de sécurité
      ansible.builtin.service:
        name: unattended-upgrades
        state: started
        enabled: true
```

Exécution :

```bash
ansible-playbook playbooks/socle.yml --check --diff   # dry-run d'abord !
ansible-playbook playbooks/socle.yml -l staging       # staging avant prod
ansible-playbook playbooks/socle.yml                  # production
```

## 20. Module apt : gérer les paquets Debian/Ubuntu

```yaml
- name: Mettre à jour le cache si vieux de plus d'une heure
  ansible.builtin.apt:
    update_cache: true
    cache_valid_time: 3600

- name: Installer une liste de paquets
  ansible.builtin.apt:
    name:
      - nginx
      - certbot
      - python3-certbot-nginx
    state: present

- name: Installer une version précise (épinglage)
  ansible.builtin.apt:
    name: postgresql=15.6-1.pgdg120+1
    state: present
    allow_downgrade: true

- name: Mise à jour de sécurité uniquement (équivalent unattended)
  ansible.builtin.apt:
    upgrade: dist
    update_cache: true

- name: Supprimer un paquet et purger sa configuration
  ansible.builtin.apt:
    name: apache2
    state: absent
    purge: true
    autoremove: true
```

> **Avertissement** : `upgrade: dist` sur un parc de production se fait en `serial: 1`
> ou par vagues (section 50), jamais en aveugle sur tous les nœuds à la fois.

## 21. Module copy : déployer des fichiers statiques

```yaml
- name: Déployer un fichier de configuration statique
  ansible.builtin.copy:
    src: files/ntp.conf            # relatif au playbook (ou au rôle : files/)
    dest: /etc/ntp.conf
    owner: root
    group: root
    mode: "0644"
    backup: true                   # garde une copie de l'ancien fichier
  notify: Redémarrer ntp

- name: Déployer un contenu inline (petits fichiers)
  ansible.builtin.copy:
    content: |
      # Géré par Ansible — ne pas modifier à la main
      127.0.0.1 localhost
    dest: /etc/hosts.d/ansible
    mode: "0644"

- name: Copier un répertoire entier
  ansible.builtin.copy:
    src: files/skel/
    dest: /etc/skel/
    mode: "0644"
```

`backup: true` crée `/etc/ntp.conf.12345~` : indispensable avant d'écraser un fichier
critique en production.

## 22. Module template : fichiers générés avec Jinja2

```yaml
- name: Générer la configuration nginx depuis un template
  ansible.builtin.template:
    src: nginx.conf.j2             # dans templates/ du rôle ou du playbook
    dest: /etc/nginx/nginx.conf
    owner: root
    group: root
    mode: "0644"
    backup: true
    validate: "nginx -t -c %s"     # valide AVANT d'écraser : sécurité maximale
  notify: Recharger nginx
```

Le paramètre `validate` est une **bonne pratique production** : `%s` est remplacé par le
fichier temporaire ; si la validation échoue, le fichier destination n'est pas touché.

## 23. Module file : fichiers, liens, répertoires, permissions

```yaml
- name: Créer un répertoire applicatif
  ansible.builtin.file:
    path: /opt/monapp
    state: directory
    owner: monapp
    group: monapp
    mode: "0750"

- name: Créer un lien symbolique
  ansible.builtin.file:
    src: /opt/monapp/releases/v2
    dest: /opt/monapp/current
    state: link

- name: Supprimer un fichier obsolète
  ansible.builtin.file:
    path: /etc/cron.d/vieux_job
    state: absent

- name: Toucher un fichier (créer s'il n'existe pas)
  ansible.builtin.file:
    path: /var/log/monapp/app.log
    state: touch
    mode: "0640"

- name: Appliquer récursivement des permissions
  ansible.builtin.file:
    path: /srv/data
    owner: www-data
    group: www-data
    mode: "0750"
    recurse: true
```

## 24. Modules service et systemd

```yaml
- name: S'assurer que nginx tourne et démarre au boot
  ansible.builtin.service:
    name: nginx
    state: started
    enabled: true

- name: Recharger (sans couper les connexions)
  ansible.builtin.systemd:
    name: nginx
    state: reloaded

- name: Redémarrer avec daemon-reload après changement d'unité
  ansible.builtin.systemd:
    name: monapp
    state: restarted
    daemon_reload: true

- name: Masquer un service indésirable
  ansible.builtin.systemd:
    name: telnet
    state: stopped
    enabled: false
    masked: true
```

`ansible.builtin.service` est agnostique (sysv/systemd) ; `ansible.builtin.systemd`
offre `daemon_reload`, `masked`, `scope=user`.

## 25. Modules user et group

```yaml
- name: Créer le groupe applicatif
  ansible.builtin.group:
    name: monapp
    gid: 2001
    state: present

