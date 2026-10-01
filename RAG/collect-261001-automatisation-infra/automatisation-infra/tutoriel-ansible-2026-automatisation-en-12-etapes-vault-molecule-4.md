---
id: collect-261001-automatisation-infra/automatisation-infra/tutoriel-ansible-2026-automatisation-en-12-etapes-vault-molecule-4
title: "Créer un environnement virtuel Python dédié"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: ["apache"]
source: docs/RAG/collect-261001-automatisation-infra/tutoriel-ansible-2026-automatisation-en-12-etapes-vault-molecule.md
source_anchor: ""
source_lines: [408, 629]
sha256: ecac6929bec8b67a84fa12dcd2dae2750c2857ebbe16efd9016d9985c62cf162
---

# Créer un environnement virtuel Python dédié

```
# Créer la structure d'un rôle avec ansible-galaxy
ansible-galaxy role init roles/serveur_web
# Structure générée :
# roles/serveur_web/
# ├── defaults/
# │   └── main.yml       # Variables par défaut (priorité la plus basse)
# ├── files/
# │   └── ...            # Fichiers statiques à copier
# ├── handlers/
# │   └── main.yml       # Handlers (redémarrages de services, etc.)
# ├── meta/
# │   └── main.yml       # Métadonnées et dépendances du rôle
# ├── tasks/
# │   └── main.yml       # Tâches principales
# ├── templates/
# │   └── ...            # Templates Jinja2
# ├── tests/
# │   └── ...            # Tests du rôle
# └── vars/
#     └── main.yml       # Variables du rôle (priorité haute)
# Installer un rôle depuis Galaxy
ansible-galaxy role install geerlingguy.docker
# Installer des collections depuis Galaxy
ansible-galaxy collection install community.general
ansible-galaxy collection install ansible.posix
```
Voici le contenu du rôle `serveur_web` qui installe et configure Nginx avec PHP :

```
# roles/serveur_web/tasks/main.yml
---
- name: Installer Nginx
  ansible.builtin.apt:
    name: nginx
    state: present
  notify: Démarrer Nginx
- name: Installer PHP et ses extensions
  ansible.builtin.apt:
    name:
      - "php{{ php_version }}-fpm"
      - "php{{ php_version }}-mysql"
      - "php{{ php_version }}-curl"
      - "php{{ php_version }}-mbstring"
      - "php{{ php_version }}-xml"
      - "php{{ php_version }}-zip"
    state: present
- name: Configurer les virtual hosts Nginx
  ansible.builtin.template:
    src: vhost.conf.j2
    dest: "/etc/nginx/sites-available/{{ item.nom }}.conf"
    mode: '0644'
  loop: "{{ vhosts }}"
  notify: Recharger Nginx
- name: Activer les virtual hosts
  ansible.builtin.file:
    src: "/etc/nginx/sites-available/{{ item.nom }}.conf"
    dest: "/etc/nginx/sites-enabled/{{ item.nom }}.conf"
    state: link
  loop: "{{ vhosts }}"
  notify: Recharger Nginx
- name: Supprimer le site par défaut
  ansible.builtin.file:
    path: /etc/nginx/sites-enabled/default
    state: absent
  notify: Recharger Nginx
- name: Créer les répertoires des applications
  ansible.builtin.file:
    path: "{{ item.racine }}"
    state: directory
    owner: www-data
    group: www-data
    mode: '0755'
  loop: "{{ vhosts }}"
- name: Configurer PHP-FPM
  ansible.builtin.template:
    src: php-fpm-pool.conf.j2
    dest: "/etc/php/{{ php_version }}/fpm/pool.d/www.conf"
    mode: '0644'
  notify: Redémarrer PHP-FPM
# roles/serveur_web/handlers/main.yml
---
- name: Démarrer Nginx
  ansible.builtin.service:
    name: nginx
    state: started
    enabled: true
- name: Recharger Nginx
  ansible.builtin.service:
    name: nginx
    state: reloaded
- name: Redémarrer PHP-FPM
  ansible.builtin.service:
    name: "php{{ php_version }}-fpm"
    state: restarted
# roles/serveur_web/templates/vhost.conf.j2
server {
    listen {{ http_port | default(80) }};
    server_name {{ item.nom }};
    root {{ item.racine }};
    index index.php index.html;
    location / {
        try_files $uri $uri/ /index.php?$query_string;
    }
    location ~ \.php$ {
        fastcgi_pass unix:/run/php/php{{ php_version }}-fpm.sock;
        fastcgi_param SCRIPT_FILENAME $realpath_root$fastcgi_script_name;
        include fastcgi_params;
    }
    location ~ /\.ht {
        deny all;
    }
    access_log /var/log/nginx/{{ item.nom }}_access.log;
    error_log /var/log/nginx/{{ item.nom }}_error.log;
}
```
Les templates Jinja2 (`.j2`) permettent de générer des fichiers de configuration dynamiques en utilisant les variables Ansible. Cette approche garantit que chaque serveur reçoit une configuration adaptée à son rôle et à ses spécificités, tout en maintenant un code source unique.

## Étape 8 – Déployer une Stack LAMP Complète

Nous allons maintenant assembler les rôles et les playbooks pour déployer une stack LAMP (Linux, Apache/Nginx, MySQL, PHP) complète sur notre infrastructure. Ce playbook principal orchestre l'ensemble du déploiement en appliquant les rôles appropriés à chaque groupe de serveurs. C'est l'étape qui transforme votre code Ansible en infrastructure réelle.

```
# playbooks/deploiement-complet.yml
---
- name: Phase 1 – Sécurité et configuration de base
  hosts: all
  become: true
  roles:
    - securite-base
  tags: [securite]
- name: Phase 2 – Déployer les serveurs web
  hosts: serveurs_web
  become: true
  roles:
    - serveur_web
  tags: [web]
  
  post_tasks:
    - name: Déployer la page de test
      ansible.builtin.copy:
        dest: "{{ vhosts[0].racine }}/index.php"
        content: |
          Serveur {{ inventory_hostname }} opérationnel";
          echo "
```
PHP " . phpversion() . "

";
          echo "
Serveur : " . gethostname() . "

";
          echo "
Date : " . date('Y-m-d H:i:s') . "

";
          phpinfo();
        owner: www-data
        group: www-data
        mode: '0644'
    - name: Vérifier que Nginx répond
      ansible.builtin.uri:
        url: "http://localhost"
        status_code: 200
      register: resultat_test
      retries: 3
      delay: 5
      until: resultat_test.status == 200
- name: Phase 3 – Déployer les serveurs de base de données
  hosts: serveurs_bdd
  become: true
  
  tasks:
    - name: Installer MySQL Server
      ansible.builtin.apt:
        name:
          - mysql-server
          - python3-pymysql
        state: present
    - name: Démarrer MySQL
      ansible.builtin.service:
        name: mysql
        state: started
        enabled: true
    - name: Configurer le mot de passe root MySQL
      community.mysql.mysql_user:
        name: root
        password: "{{ mysql_root_password }}"
        login_unix_socket: /var/run/mysqld/mysqld.sock
        state: present
    - name: Créer les bases de données
      community.mysql.mysql_db:
        name: "{{ item.nom }}"
        encoding: "{{ item.encoding }}"
        collation: "{{ item.collation }}"
        state: present
        login_user: root
        login_password: "{{ mysql_root_password }}"
      loop: "{{ mysql_bases }}"
    - name: Créer les utilisateurs MySQL
      community.mysql.mysql_user:
        name: "{{ item.nom }}"
        password: "{{ item.mot_de_passe }}"
        priv: "{{ item.privileges }}"
        host: "{{ item.hote }}"
        state: present
        login_user: root
        login_password: "{{ mysql_root_password }}"
      loop: "{{ mysql_utilisateurs }}"
# Exécuter le déploiement complet
# ansible-playbook playbooks/deploiement-complet.yml
# Exécuter uniquement la partie web
# ansible-playbook playbooks/deploiement-complet.yml --tags web
# Exécuter sur un seul serveur
# ansible-playbook playbooks/deploiement-complet.yml --limit web01
L'utilisation des **tags** permet d'exécuter sélectivement des parties du playbook, ce qui est particulièrement utile pendant le développement ou pour les mises à jour ciblées. L'option `--limit` restreint l'exécution à un serveur spécifique, idéal pour les déploiements progressifs (canary deployments) où vous testez d'abord sur un seul serveur avant de déployer sur l'ensemble de la flotte.

## Étape 9 – Gérer les Conditions, Boucles et Blocs

Ansible offre des structures de contrôle puissantes qui permettent d'adapter le comportement des playbooks en fonction du contexte d'exécution. Les conditions (`when`), les boucles (`loop`) et les blocs (`block/rescue/always`) rendent vos playbooks flexibles et robustes face aux situations variées que vous rencontrerez en production.

