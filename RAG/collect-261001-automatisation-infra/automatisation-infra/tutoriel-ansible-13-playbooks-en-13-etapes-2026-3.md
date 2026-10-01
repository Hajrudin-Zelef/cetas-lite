---
id: collect-261001-automatisation-infra/automatisation-infra/tutoriel-ansible-13-playbooks-en-13-etapes-2026-3
title: "Mise à jour des paquets et installation de pipx"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: ["benchmark"]
source: docs/RAG/collect-261001-automatisation-infra/tutoriel-ansible-13-playbooks-en-13-etapes-2026.md
source_anchor: ""
source_lines: [182, 332]
sha256: 88abe345a83b37e5f34fb9a7a5d5af67b1d605a04e8475989b1a2c0e40116adc
---

# Mise à jour des paquets et installation de pipx

```
# playbooks/bootstrap.yml
---
- name: Bootstrap initial des serveurs Linux
  hosts: all
  become: true
  gather_facts: true
  tasks:
    - name: Mettre à jour le cache APT
      ansible.builtin.apt:
        update_cache: true
        cache_valid_time: 3600
      when: ansible_os_family == "Debian"
    - name: Installer les paquets de base
      ansible.builtin.package:
        name:
          - curl
          - vim
          - htop
          - git
          - chrony
          - ufw
          - fail2ban
        state: present
    - name: Configurer le fuseau horaire Europe/Paris
      community.general.timezone:
        name: Europe/Paris
      notify: Redémarrer chrony
    - name: Activer le pare-feu UFW
      community.general.ufw:
        state: enabled
        policy: deny
        direction: incoming
    - name: Autoriser SSH sur le pare-feu
      community.general.ufw:
        rule: allow
        port: "22"
        proto: tcp
  handlers:
    - name: Redémarrer chrony
      ansible.builtin.service:
        name: chrony
        state: restarted
```
Exécutez ce playbook avec `ansible-playbook playbooks/bootstrap.yml`. Ansible affiche un résumé coloré : **ok** en vert pour les tâches sans changement, **changed** en jaune quand l’état a été modifié, **failed** en rouge en cas d’erreur. Sur 3 hôtes neufs, comptez environ 45 secondes pour le premier passage. Un guide OneUptime publié en février 2026 recommande d’ailleurs d’instrumenter ce résumé avec un petit script de benchmark qui journalise la durée en secondes (`duration_seconds`) ainsi que le décompte des tâches `ok`, `changed` et `failed` à chaque exécution, pour suivre la dérive de performance dans le temps. Le deuxième passage devrait afficher uniquement `ok=10 changed=0`, démontrant l’idempotence.

## Étape 5 : Maîtriser les variables, les facts et Jinja2

Les variables Ansible sont chargées depuis 22 emplacements différents avec une précédence stricte. Comprendre cet ordre est crucial pour éviter qu’une variable de `group_vars` n’écrase une valeur passée en ligne de commande. Les **facts**, eux, sont collectés automatiquement par le module `setup` au début de chaque play : ils exposent l’OS, l’IP, la mémoire, le hostname, les interfaces réseau et plus de 200 attributs par hôte.

| Source de variable | Précédence | Usage typique | 
|---|---|---|
| Extra vars (-e en CLI) | 1 (max) | Override ponctuel en pipeline CI/CD | 
| set_fact dans une tâche | 3 | Variables calculées dynamiquement | 
| Block vars | 5 | Scope local à un block de tâches | 
| Role vars (vars/main.yml) | 9 | Constantes internes au rôle | 
| Host vars (host_vars/host.yml) | 13 | Spécifique à un hôte | 
| Group vars (group_vars/group.yml) | 14 | Spécifique à un groupe | 
| Group vars (all.yml) | 15 | Valeurs par défaut globales | 
| Role defaults (defaults/main.yml) | 22 (min) | Valeurs override-ables par défaut | 

Les templates Jinja2 permettent de générer des fichiers de configuration dynamiques. La syntaxe `{{ variable }}` insère une valeur, `{% if %}` applique des conditions, `{% for %}` itère sur une liste. Combinées aux facts, elles produisent des configurations adaptées à chaque hôte sans duplication. L’exemple ci-dessous génère un `nginx.conf` dont le nombre de workers s’adapte au nombre de cœurs CPU détectés.

```
# templates/nginx.conf.j2
user www-data;
worker_processes {{ ansible_processor_vcpus | default(2) }};
pid /run/nginx.pid;
worker_rlimit_nofile {{ nginx_worker_rlimit | default(65535) }};
events {
    worker_connections {{ nginx_worker_connections | default(1024) }};
    use epoll;
    multi_accept on;
}
http {
    server_tokens off;
    sendfile on;
    tcp_nopush on;
    keepalive_timeout 65;
    gzip on;
    {% for upstream in nginx_upstreams %}
    upstream {{ upstream.name }} {
        {% for backend in upstream.servers %}
        server {{ backend }};
        {% endfor %}
    }
    {% endfor %}
    include /etc/nginx/sites-enabled/*;
}
```
## Étape 6 : Découper la logique en rôles réutilisables

Un rôle Ansible est un module logique autonome, versionné, testable, qui regroupe tâches, handlers, templates, fichiers, variables et métadonnées. La règle d’or : **un rôle, une responsabilité**. Un rôle `nginx` installe et configure Nginx. Un rôle `mariadb` gère la base. Un rôle `app-php` déploie le code applicatif. Cette séparation rend les rôles réutilisables entre projets et permet leur publication sur Ansible Galaxy.

```
# Création de la structure d'un rôle
ansible-galaxy role init roles/nginx
# Arborescence générée
roles/nginx/
├── defaults/main.yml      # Valeurs par défaut override-ables
├── files/                 # Fichiers à copier tel quel
├── handlers/main.yml      # Actions différées (restart, reload)
├── meta/main.yml          # Dépendances et métadonnées Galaxy
├── tasks/main.yml         # Tâches principales du rôle
├── templates/             # Templates Jinja2
├── tests/                 # Inventaire et playbook de test
└── vars/main.yml          # Variables internes
```
Le fichier `tasks/main.yml` contient la logique métier. Pour le rôle Nginx, il installe le paquet, dépose un template de configuration, active le service et vérifie qu’il répond. La directive `notify` déclenche un handler à la fin du play uniquement si la tâche a modifié l’état : Nginx ne sera redémarré que si la config a réellement changé.

```
# roles/nginx/tasks/main.yml
---
- name: Installer Nginx
  ansible.builtin.package:
    name: nginx
    state: present
- name: Déployer la configuration principale
  ansible.builtin.template:
    src: nginx.conf.j2
    dest: /etc/nginx/nginx.conf
    owner: root
    group: root
    mode: "0644"
    backup: true
    validate: nginx -t -c %s
  notify: Reload Nginx
- name: Déployer les vhosts depuis la liste
  ansible.builtin.template:
    src: vhost.conf.j2
    dest: "/etc/nginx/sites-available/{{ item.name }}"
    mode: "0644"
  loop: "{{ nginx_vhosts }}"
  notify: Reload Nginx
- name: Activer les vhosts
  ansible.builtin.file:
    src: "/etc/nginx/sites-available/{{ item.name }}"
    dest: "/etc/nginx/sites-enabled/{{ item.name }}"
    state: link
  loop: "{{ nginx_vhosts }}"
  notify: Reload Nginx
- name: S'assurer que Nginx est démarré
  ansible.builtin.service:
    name: nginx
    state: started
    enabled: true
```
## Étape 7 : Sécuriser les secrets avec Ansible Vault

Stocker des mots de passe, clés API ou tokens en clair dans un dépôt Git est une faute professionnelle en 2026. **Ansible Vault** chiffre des fichiers entiers ou des chaînes individuelles avec AES-256-CTR et une clé dérivée d’un mot de passe (PBKDF2 SHA-256, 10 000 itérations). Les fichiers vaultés sont déchiffrés en mémoire au runtime sans jamais toucher le disque, et leur contenu reste opaque pour quiconque ne possède pas le mot de passe.

