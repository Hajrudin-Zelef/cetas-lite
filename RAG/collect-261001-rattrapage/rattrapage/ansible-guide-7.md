---
id: collect-261001-rattrapage/rattrapage/ansible-guide-7
title: "Guide Ansible — Automatisation système en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "attention"]
source: docs/RAG/collect-261001-rattrapage/ansible_guide.md
source_anchor: ""
source_lines: [1543, 1816]
sha256: 2996d6eb429c7894815df781e545bc28f361609967cc227bbf6302ca6a86f730
---

# batch par groupe avec max_fail_percentage
- hosts: all
  serial: "20%"
  max_fail_percentage: 10   # stop si plus de 10 % d'échecs dans la vague
  tasks: [...]
```

Repères de dimensionnement :

| Taille du parc | forks conseillé | serial conseillé |
|---|---|---|
| < 20 hôtes | 5–10 | inutile |
| 20–100 hôtes | 20–30 | `25%` pour les déploiements |
| 100–500 hôtes | 50–100 | `[10, "25%", "50%"]` |
| 500+ hôtes | 100+ + `host_pinned` | par vagues + `max_fail_percentage` |

Surveillez la charge du control node : chaque fork = un processus SSH + Python.

## 51. Gestion des erreurs : failed_when et changed_when

```yaml
- name: Vérifier la santé de l'application
  ansible.builtin.uri:
    url: http://localhost:8080/health
    status_code: 200
  register: sante
  # Échec seulement si le service ne répond pas après 3 essais
  retries: 3
  delay: 10
  until: sante.status == 200

- name: Commande dont le code 2 est normal
  ansible.builtin.command: /usr/bin/monoutil --verif
  register: verif
  failed_when: verif.rc not in [0, 2]
  changed_when: "'MODIFIE' in verif.stdout"

- name: Lecture seule, jamais marquée changed
  ansible.builtin.command: cat /etc/os-release
  changed_when: false
```

Par défaut, `failed_when` = code retour non nul. Surchargez-le dès qu'un outil
utilise des codes de sortie non standard.

## 52. block / rescue / always : le try/catch d'Ansible

```yaml
- name: Déploiement avec rollback
  block:
    - name: Arrêter l'application
      ansible.builtin.service:
        name: monapp
        state: stopped

    - name: Déployer la nouvelle version
      ansible.builtin.unarchive:
        src: /srv/releases/monapp-v2.5.tar.gz
        dest: /opt/monapp
        remote_src: true

  rescue:
    - name: Restaurer la version précédente
      ansible.builtin.unarchive:
        src: /srv/releases/monapp-v2.4.tar.gz
        dest: /opt/monapp
        remote_src: true

    - name: Alerter l'équipe
      ansible.builtin.debug:
        msg: "ROLLBACK effectué sur {{ inventory_hostname }}"

  always:
    - name: Redémarrer l'application dans tous les cas
      ansible.builtin.service:
        name: monapp
        state: started
```

`block` = tentative, `rescue` = exécuté **uniquement en cas d'échec**,
`always` = exécuté **dans tous les cas** (nettoyage, redémarrage).

## 53. Check mode (--check) et diff (--diff)

```bash
# Dry-run : montre ce qui CHANGERAIT sans rien modifier
ansible-playbook site.yml --check

# Dry-run + diff détaillé des fichiers
ansible-playbook site.yml --check --diff

# Forcer l'exécution réelle d'une tâche normalement ignorée en check
# (ex. : créer un répertoire temporaire nécessaire au diff)
- name: Préparer le répertoire de travail
  ansible.builtin.file:
    path: /tmp/work
    state: directory
  check_mode: false

# Forcer le dry-run d'une tâche dangereuse même sans --check
- name: Ne jamais exécuter sans --check
  ansible.builtin.command: /sbin/reboot
  check_mode: true   # tâche simulée même en exécution réelle : à manier avec soin
```

> **Limite du check mode** : les modules `command`/`shell` ne peuvent pas prédire
> leur effet ; Ansible les **saute** en `--check` sauf `check_mode: false`.
> Un `--check` vert ne garantit donc pas une exécution réelle verte.

## 54. Tags : exécuter une partie du playbook

```yaml
- name: Installer nginx
  ansible.builtin.apt:
    name: nginx
    state: present
  tags: [install, nginx]

- name: Configurer nginx
  ansible.builtin.template:
    src: nginx.conf.j2
    dest: /etc/nginx/nginx.conf
  tags: [config, nginx]
  notify: Recharger nginx

- name: Toujours exécuté (prérequis)
  ansible.builtin.apt:
    update_cache: true
    cache_valid_time: 3600
  tags: [always]
```

```bash
ansible-playbook site.yml --tags "config"        # que la configuration
ansible-playbook site.yml --tags "nginx"         # tout nginx
ansible-playbook site.yml --skip-tags "install"  # tout sauf l'installation
ansible-playbook site.yml --list-tags            # inventorier les tags
```

Convention d'équipe : `install`, `config`, `service`, `securite`, `always`, `never`
(pour les tâches destructrices type `reboot` à déclencher explicitement).

## 55. Délégation : delegate_to, run_once, local_action

```yaml
- name: Exécuter sur le control node (ex. : appel API central)
  ansible.builtin.uri:
    url: https://cmdb.example.lan/api/maint/{{ inventory_hostname }}
    method: POST
  delegate_to: localhost
  run_once: true        # une seule fois pour tout le play, pas par hôte

- name: Sauvegarder la base depuis le serveur de backup
  ansible.builtin.command: /usr/local/bin/dump.sh {{ inventory_hostname }}
  delegate_to: backup-01
  # la tâche tourne sur backup-01, mais UNE FOIS PAR HÔTE du play

- name: Ajouter temporairement au load balancer
  ansible.builtin.command: /usr/local/bin/lb_drain.sh {{ inventory_hostname }}
  delegate_to: lb-01
```

`delegate_to: localhost` + `run_once: true` : le motif standard pour les actions
centralisées (DNS, CMDB, tickets, supervision) dans un play multi-hôtes.

## 56. Boucles : loop et filtres associés

```yaml
- name: Installer plusieurs paquets
  ansible.builtin.apt:
    name: "{{ item }}"
    state: present
  loop:
    - nginx
    - certbot
    - htop

- name: Créer plusieurs utilisateurs
  ansible.builtin.user:
    name: "{{ item.nom }}"
    groups: "{{ item.groupes }}"
    shell: /bin/bash
  loop:
    - { nom: alice, groupes: sudo }
    - { nom: bob, groupes: docker }

- name: Déployer plusieurs templates
  ansible.builtin.template:
    src: "{{ item.src }}"
    dest: "{{ item.dest }}"
    mode: "0644"
  loop:
    - { src: app.conf.j2, dest: /etc/monapp/app.conf }
    - { src: log.conf.j2, dest: /etc/monapp/log.conf }

# Boucle sur un dictionnaire
- name: Configurer les vhosts
  ansible.builtin.template:
    src: vhost.conf.j2
    dest: "/etc/nginx/sites-enabled/{{ item.key }}.conf"
  loop: "{{ vhosts | dict2items }}"
```

Anciennes syntaxes `with_items`, `with_dict` : fonctionnelles mais dépréciées
dans la documentation — préférez `loop` + filtres Jinja2.

## 57. Conditions : when

```yaml
- name: Installer le paquet Debian uniquement sur Debian/Ubuntu
  ansible.builtin.apt:
    name: nftables
    state: present
  when: ansible_facts['os_family'] == "Debian"

- name: Redémarrer si le noyau a changé
  ansible.builtin.reboot:
  when:
    - kernel_update is changed
    - autoriser_reboot | default(false)

- name: Tâche réservée à la production
  ansible.builtin.debug:
    msg: "Je tourne en PROD"
  when: "'production' in group_names"

# when sur un block entier
- name: Bloc conditionnel
  when: installer_monitoring | default(true)
  block:
    - name: Installer node_exporter
      ...
```

`when` accepte une liste (ET logique). Pour OU logique : `when: a or b`.
Attention : `when: ma_variable` échoue si la variable n'existe pas —
utilisez `| default(...)` ou `is defined`.

## 58. become : l'élévation de privilèges

```yaml
# Au niveau du play (recommandé : homogène)
- hosts: all
  become: true
  become_user: root
  become_method: sudo
  tasks: [...]

# Ponctuellement sur une tâche
- name: Lire un fichier root en tant que simple utilisateur sudo
  ansible.builtin.command: cat /root/secret.txt
  become: true
  changed_when: false

# Devenir un autre utilisateur que root
- name: Exécuter en tant que postgres
  ansible.builtin.command: psql -c "SELECT 1"
  become: true
  become_user: postgres
  changed_when: false
```

Côté sudoers sur les cibles (`/etc/sudoers.d/ansible`) :

```text
deploy ALL=(ALL) NOPASSWD: ALL
```

> **Sécurité** : `NOPASSWD` est pratique mais large. En environnement sensible,
> restreignez les commandes autorisées et utilisez `become_ask_pass: true`
> avec `--ask-become-pass` pour les interventions manuelles.

## 59. SSH : clés, multiplexage, bastion, durcissement

