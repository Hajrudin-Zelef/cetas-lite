---
id: collect-261001-rattrapage/rattrapage/ansible-guide-4
title: "Guide Ansible — Automatisation système en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-rattrapage/ansible_guide.md
source_anchor: ""
source_lines: [747, 1006]
sha256: 25d0aade132dd9f6c7347c7fea043f6fe97d1eeaa98b42b774d79a1880f5f997
---

# Guide Ansible — Automatisation système en production

- name: Créer un utilisateur de service (sans shell, sans home)
  ansible.builtin.user:
    name: monapp
    group: monapp
    uid: 2001
    shell: /usr/sbin/nologin
    create_home: false
    system: true

- name: Créer un admin avec clé SSH
  ansible.builtin.user:
    name: alice
    groups: sudo
    append: true                  # ne pas écraser les groupes existants !
    shell: /bin/bash
    password: "{{ vault_alice_password_hash }}"  # hash, jamais en clair

- name: Déployer la clé publique SSH d'alice
  ansible.posix.authorized_key:
    user: alice
    key: "ssh-ed25519 AAAA... alice@poste"
    state: present
```

> **Sécurité** : `password` attend un **hash** (ex. `$y$j9T$...`), généré par
> `mkpasswd --method=yescrypt`. Ne jamais mettre un mot de passe en clair,
> même dans Vault : préférez toujours le hash.

## 26. Module cron

```yaml
- name: Sauvegarde quotidienne à 2h du matin
  ansible.builtin.cron:
    name: "sauvegarde quotidienne"
    user: root
    minute: "0"
    hour: "2"
    job: "/usr/local/bin/backup.sh >> /var/log/backup.log 2>&1"

- name: Job avec variables d'environnement
  ansible.builtin.cron:
    name: "nettoyage tmp"
    user: monapp
    minute: "30"
    hour: "3"
    weekday: "0"
    job: "/opt/monapp/bin/cleanup.sh"
    env:
      - name: PATH
        value: "/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"

- name: Supprimer un job obsolète
  ansible.builtin.cron:
    name: "ancien job"
    state: absent
```

Chaque entrée est identifiée par `name` : Ansible la retrouve et la met à jour de façon
idempotente au lieu de dupliquer la ligne.

## 27. Modules ufw et firewalld

```yaml
# Debian/Ubuntu : UFW
- name: Activer UFW avec politique restrictive
  community.general.ufw:
    state: enabled
    policy: deny
    direction: incoming

- name: Autoriser SSH (limité contre le brute-force)
  community.general.ufw:
    rule: limit
    port: "22"
    proto: tcp

- name: Autoriser HTTP/HTTPS
  community.general.ufw:
    rule: allow
    port: "80,443"
    proto: tcp

- name: Autoriser depuis le réseau d'administration uniquement
  community.general.ufw:
    rule: allow
    port: "5666"          # NRPE / supervision
    src: 192.168.20.0/24
```

```yaml
# RHEL-like : firewalld
- name: Ouvrir HTTPS en permanent
  ansible.posix.firewalld:
    service: https
    permanent: true
    immediate: true
    state: enabled
```

## 28. Modules lineinfile et blockinfile

Pour modifier **une ligne** dans un fichier existant sans template complet :

```yaml
- name: S'assurer que le forwarding IP est activé
  ansible.builtin.lineinfile:
    path: /etc/sysctl.conf
    regexp: '^#?net.ipv4.ip_forward\s*='
    line: 'net.ipv4.ip_forward=1'
    backup: true
  notify: Appliquer sysctl

- name: Forcer PermitRootLogin à no dans sshd_config
  ansible.builtin.lineinfile:
    path: /etc/ssh/sshd_config
    regexp: '^#?PermitRootLogin'
    line: 'PermitRootLogin no'
    validate: 'sshd -t -f %s'
  notify: Redémarrer sshd

- name: Insérer un bloc géré par Ansible
  ansible.builtin.blockinfile:
    path: /etc/hosts
    marker: "# {mark} BLOC ANSIBLE"
    block: |
      192.168.10.50 nas-01
      192.168.10.51 nas-02
```

> **Bonne pratique** : si vous gérez **tout** le fichier, préférez `template`.
> `lineinfile` est réservé aux ajustements chirurgicaux sur des fichiers
> partiellement gérés par le système.

## 29. Modules get_url et unarchive

```yaml
- name: Télécharger un binaire avec vérification de checksum
  ansible.builtin.get_url:
    url: https://github.com/prometheus/node_exporter/releases/download/v1.8.2/node_exporter-1.8.2.linux-amd64.tar.gz
    dest: /tmp/node_exporter.tar.gz
    checksum: sha256:81a8c4294a7a7a8b8e0c...   # checksum officiel
    mode: "0644"

- name: Extraire l'archive (idempotent grâce à creates)
  ansible.builtin.unarchive:
    src: /tmp/node_exporter.tar.gz
    dest: /opt/
    remote_src: true
    creates: /opt/node_exporter-1.8.2.linux-amd64/node_exporter
```

`télécharger + vérifier le checksum + extraire` : le trio de base pour déployer
un binaire tiers sans paquet officiel.

## 30. Module git : déployer du code

```yaml
- name: Cloner le dépôt applicatif sur une branche précise
  ansible.builtin.git:
    repo: "git@git.example.lan:equipe/monapp.git"
    dest: /opt/monapp
    version: main
    key_file: /root/.ssh/id_ed25519_deploy   # clé de déploiement lecture seule
    accept_hostkey: true
    force: false                              # ne pas écraser les modifs locales

- name: Mettre à jour vers un tag précis puis notifier
  ansible.builtin.git:
    repo: https://git.example.lan/equipe/monapp.git
    dest: /opt/monapp
    version: v2.4.1
  notify: Redémarrer monapp
  register: git_result
```

Le module retourne `changed: true` uniquement si le dépôt a réellement bougé :
le handler ne redémarre l'application qu'en cas de nouveau code.

## 31. Modules command, shell, raw : à utiliser en dernier recours

```yaml
- name: Commande simple sans shell (préféré quand on doit exécuter)
  ansible.builtin.command:
    cmd: /usr/local/bin/outil --check
  register: resultat
  changed_when: false          # lecture seule : jamais "changed"
  failed_when: resultat.rc not in [0, 2]

- name: Shell uniquement si pipes/redirections indispensables
  ansible.builtin.shell:
    cmd: "cat /var/log/app/*.log | grep -c ERROR || true"
  register: nb_erreurs
  changed_when: false

- name: Raw sur un équipement sans Python (switch, vieux système)
  ansible.builtin.raw: "show version"
```

**Règle d'or** : si un module existe, utilisez-le. `shell`/`command` ne sont pas
idempotents par nature : encadrez-les toujours de `changed_when` / `failed_when`
/ `creates` / `removes`.

## 32. Modules debug, set_fact, assert

```yaml
- name: Afficher une variable pendant le développement
  ansible.builtin.debug:
    msg: "L'hôte {{ inventory_hostname }} a l'IP {{ ansible_host }}"

- name: Afficher toute la structure d'une variable
  ansible.builtin.debug:
    var: ansible_facts['network']

- name: Mémoriser un calcul pour la suite du play
  ansible.builtin.set_fact:
    backup_dir: "/srv/backup/{{ inventory_hostname }}"

- name: Valider un prérequis avant de continuer
  ansible.builtin.assert:
    that:
      - ansible_facts['distribution'] in ['Debian', 'Ubuntu']
      - ansible_facts['distribution_major_version'] | int >= 11
    fail_msg: "Ce playbook exige Debian 11+ ou Ubuntu 22.04+"
    success_msg: "OS compatible"
```

`assert` en début de playbook évite d'appliquer une configuration à un OS non testé :
c'est un garde-fou production simple et efficace.

## 33. Variables : types, définitions, portée

```yaml
# group_vars/all.yml — types de variables
nom_app: monapp                                   # chaîne
http_port: 8080                                  # entier
tls_active: true                                 # booléen
serveurs_dns:                                    # liste
  - 192.168.10.1
  - 192.168.10.2
config_reseau:                                   # dictionnaire
  interface: eth0
  mtu: 1500
liste_mixte:                                     # liste de dictionnaires
  - { nom: web-01, role: frontal }
  - { nom: db-01, role: donnees }
```

Où définir les variables (de la plus générale à la plus spécifique) :

| Emplacement | Portée | Usage |
|---|---|---|
| `group_vars/all.yml` | Tout l'inventaire | Valeurs par défaut globales |
| `group_vars/<groupe>.yml` | Un groupe | Spécificités par rôle serveur |
| `host_vars/<hote>.yml` | Un hôte | Exceptions individuelles |
| `vars:` du play | Le play | Paramétrage du playbook |
| `vars:` de la tâche | La tâche | Cas ultra-local |
| `-e` ligne de commande | L'exécution | Surcharges ponctuelles (priorité max) |

## 34. Ordre de précédence des variables (21 niveaux)

