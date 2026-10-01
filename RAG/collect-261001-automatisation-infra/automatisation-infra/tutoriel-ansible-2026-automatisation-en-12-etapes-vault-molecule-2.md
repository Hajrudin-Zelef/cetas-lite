---
id: collect-261001-automatisation-infra/automatisation-infra/tutoriel-ansible-2026-automatisation-en-12-etapes-vault-molecule-2
title: "Créer un environnement virtuel Python dédié"
domain: automatisation-infra
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-automatisation-infra/tutoriel-ansible-2026-automatisation-en-12-etapes-vault-molecule.md
source_anchor: ""
source_lines: [58, 247]
sha256: cacc7ccaeadd4e9bceeffccc3f64563a538ae680641bdb8331352db0acc5ae27
---

# Créer un environnement virtuel Python dédié

```
# Créer la structure du projet
mkdir -p ~/ansible-projet/{inventaire,playbooks,roles,group_vars,host_vars,fichiers,templates}
cd ~/ansible-projet
# Créer le fichier d'inventaire YAML
cat > inventaire/production.yml << 'EOF'
---
all:
  children:
    serveurs_web:
      hosts:
        web01:
          ansible_host: 192.168.1.10
          ansible_user: deploy
          http_port: 80
        web02:
          ansible_host: 192.168.1.11
          ansible_user: deploy
          http_port: 80
    serveurs_bdd:
      hosts:
        db01:
          ansible_host: 192.168.1.20
          ansible_user: deploy
          mysql_port: 3306
    serveurs_monitoring:
      hosts:
        monitor01:
          ansible_host: 192.168.1.30
          ansible_user: deploy
    production:
      children:
        serveurs_web:
        serveurs_bdd:
        serveurs_monitoring:
EOF
# Vérifier l'inventaire
ansible-inventory -i inventaire/production.yml --list
# Résultat attendu : liste JSON de tous les hôtes et groupes
# Tester la connectivité vers tous les serveurs
ansible all -i inventaire/production.yml -m ping
# web01 | SUCCESS => {"changed": false, "ping": "pong"}
# web02 | SUCCESS => {"changed": false, "ping": "pong"}
# db01 | SUCCESS => {"changed": false, "ping": "pong"}
# monitor01 | SUCCESS => {"changed": false, "ping": "pong"}
```
**Inventaire dynamique** : pour les environnements cloud (AWS, Azure, GCP), Ansible propose des plugins d’inventaire dynamique qui interrogent automatiquement l’API du fournisseur pour découvrir les instances. Par exemple, le plugin `amazon.aws.aws_ec2` génère l’inventaire à partir des tags EC2. Cette approche est indispensable dans les architectures auto-scalantes où les serveurs sont créés et détruits fréquemment.

Les variables d’hôte (`ansible_host`, `ansible_user`) et les variables personnalisées (`http_port`, `mysql_port`) permettent d’adapter le comportement des playbooks à chaque serveur sans modifier le code. C’est un principe fondamental d’Ansible : séparer les données de la logique d’automatisation.

## Étape 3 – Écrire le Fichier de Configuration ansible.cfg

Le fichier `ansible.cfg` centralise la configuration d’Ansible pour votre projet. Sans ce fichier, Ansible utilise les paramètres par défaut, ce qui peut entraîner des comportements inattendus, notamment la vérification des clés SSH qui bloque l’exécution sur de nouveaux serveurs. Un fichier de configuration bien pensé accélère considérablement le développement et le déploiement.

```
# Créer ansible.cfg à la racine du projet
cat > ~/ansible-projet/ansible.cfg << 'EOF'
[defaults]
# Chemin vers l'inventaire par défaut
inventory = inventaire/production.yml
# Utilisateur distant par défaut
remote_user = deploy
# Désactiver la vérification des clés SSH (dev uniquement)
host_key_checking = False
# Chemin vers les rôles
roles_path = roles
# Parallélisme : nombre de serveurs gérés simultanément
forks = 10
# Format de sortie coloré et lisible
stdout_callback = yaml
callbacks_enabled = timer, profile_tasks
# Répertoire pour les faits mis en cache
fact_caching = jsonfile
fact_caching_connection = /tmp/ansible_facts_cache
fact_caching_timeout = 3600
# Timeout de connexion SSH (secondes)
timeout = 30
# Réessayer les connexions échouées
retries = 3
[privilege_escalation]
# Utiliser sudo pour les tâches nécessitant des privilèges root
become = True
become_method = sudo
become_user = root
become_ask_pass = False
[ssh_connection]
# Optimisation SSH : réutiliser les connexions
pipelining = True
ssh_args = -o ControlMaster=auto -o ControlPersist=60s -o PreferredAuthentications=publickey
EOF
```
Le paramètre `pipelining = True` est une optimisation cruciale : il réduit le nombre d'opérations SSH nécessaires pour exécuter un module, ce qui peut accélérer l'exécution des playbooks de 30 à 50 %. Le paramètre `forks = 10` permet de gérer 10 serveurs simultanément au lieu de 5 par défaut, ce qui est particulièrement utile pour les déploiements à grande échelle.

**Sécurité** : en production, désactivez `host_key_checking = False` et utilisez plutôt un fichier `known_hosts` géré. Cette option est uniquement recommandée pour les environnements de développement et de test. De même, le paramètre `become_ask_pass = False` suppose que votre utilisateur `deploy` dispose de sudo sans mot de passe via le fichier sudoers – configurez cela sur chaque serveur cible avant de lancer vos playbooks.

## Étape 4 – Créer un Premier Playbook Ansible

Les playbooks sont le cœur d'Ansible : ils définissent les tâches à exécuter sur les serveurs cibles dans un format YAML déclaratif. Un playbook bien structuré est idempotent – il peut être exécuté plusieurs fois sans effet indésirable. Cette propriété est essentielle pour garantir la fiabilité des déploiements automatisés. Commençons par un playbook qui configure la sécurité de base d'un serveur.

```
# playbooks/securite-base.yml
---
- name: Configuration de sécurité de base des serveurs
  hosts: all
  become: true
  vars:
    utilisateur_deploy: deploy
    paquets_securite:
      - ufw
      - fail2ban
      - unattended-upgrades
      - logwatch
    ports_autorises:
      - "22"
      - "80"
      - "443"
  tasks:
    - name: Mettre à jour le cache APT et les paquets
      ansible.builtin.apt:
        update_cache: true
        upgrade: safe
        cache_valid_time: 3600
    - name: Installer les paquets de sécurité
      ansible.builtin.apt:
        name: "{{ paquets_securite }}"
        state: present
    - name: Créer l'utilisateur de déploiement
      ansible.builtin.user:
        name: "{{ utilisateur_deploy }}"
        groups: sudo
        shell: /bin/bash
        create_home: true
        state: present
    - name: Configurer l'authentification SSH par clé uniquement
      ansible.builtin.lineinfile:
        path: /etc/ssh/sshd_config
        regexp: "{{ item.regexp }}"
        line: "{{ item.line }}"
        state: present
        validate: sshd -t -f %s
      loop:
        - { regexp: '^#?PasswordAuthentication', line: 'PasswordAuthentication no' }
        - { regexp: '^#?PermitRootLogin', line: 'PermitRootLogin no' }
        - { regexp: '^#?PubkeyAuthentication', line: 'PubkeyAuthentication yes' }
      notify: Redémarrer SSH
    - name: Configurer le pare-feu UFW
      community.general.ufw:
        rule: allow
        port: "{{ item }}"
        proto: tcp
      loop: "{{ ports_autorises }}"
    - name: Activer UFW avec politique par défaut deny
      community.general.ufw:
        state: enabled
        policy: deny
        direction: incoming
    - name: Configurer Fail2Ban pour SSH
      ansible.builtin.copy:
        dest: /etc/fail2ban/jail.local
        content: |
          [sshd]
          enabled = true
          port = ssh
          filter = sshd
          logpath = /var/log/auth.log
          maxretry = 3
          bantime = 3600
          findtime = 600
        mode: '0644'
      notify: Redémarrer Fail2Ban
    - name: Activer les mises à jour automatiques de sécurité
      ansible.builtin.copy:
        dest: /etc/apt/apt.conf.d/20auto-upgrades
        content: |
          APT::Periodic::Update-Package-Lists "1";
          APT::Periodic::Unattended-Upgrade "1";
          APT::Periodic::AutocleanInterval "7";
        mode: '0644'
  handlers:
    - name: Redémarrer SSH
      ansible.builtin.service:
        name: sshd
        state: restarted
    - name: Redémarrer Fail2Ban
      ansible.builtin.service:
        name: fail2ban
        state: restarted
```
Exécutez ce playbook avec la commande suivante :

