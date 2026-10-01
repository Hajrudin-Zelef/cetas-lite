---
id: collect-261001-automatisation-infra/automatisation-infra/tutoriel-ansible-2026-automatisation-en-12-etapes-vault-molecule-6
title: "Créer un environnement virtuel Python dédié"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-automatisation-infra/tutoriel-ansible-2026-automatisation-en-12-etapes-vault-molecule.md
source_anchor: ""
source_lines: [797, 965]
sha256: a5622f2a999aa76a64f412b82dcc3a485d46535158f1428f688160448cbdb1f2
---

# Créer un environnement virtuel Python dédié

```
# Installer Molecule avec le driver Docker
pip install molecule molecule-docker
# Initialiser Molecule dans un rôle existant
cd roles/serveur_web
molecule init scenario --driver-name docker
# molecule/default/molecule.yml
---
dependency:
  name: galaxy
driver:
  name: docker
platforms:
  - name: ubuntu-test
    image: ubuntu:22.04
    pre_build_image: true
    command: /sbin/init
    privileged: true
    volumes:
      - /sys/fs/cgroup:/sys/fs/cgroup:ro
provisioner:
  name: ansible
  inventory:
    host_vars:
      ubuntu-test:
        php_version: "8.3"
        http_port: 80
        vhosts:
          - nom: "test.local"
            racine: "/var/www/test"
            ssl: false
verifier:
  name: ansible
# molecule/default/converge.yml
---
- name: Converge
  hosts: all
  become: true
  roles:
    - role: serveur_web
# molecule/default/verify.yml
---
- name: Verify
  hosts: all
  become: true
  tasks:
    - name: Vérifier que Nginx est installé
      ansible.builtin.package:
        name: nginx
        state: present
      check_mode: true
      register: nginx_check
      failed_when: nginx_check.changed
    - name: Vérifier que Nginx écoute sur le port 80
      ansible.builtin.wait_for:
        port: 80
        timeout: 10
    - name: Vérifier que le répertoire web existe
      ansible.builtin.stat:
        path: /var/www/test
      register: web_dir
      failed_when: not web_dir.stat.exists
# Exécuter les tests
molecule test
# Résultat attendu :
# --> Test matrix
# └── default
#     ├── dependency
#     ├── lint
#     ├── cleanup
#     ├── destroy
#     ├── syntax
#     ├── create
#     ├── prepare
#     ├── converge
#     ├── idempotence
#     ├── verify
#     └── destroy
# 
# PLAY RECAP ********
# ubuntu-test : ok=12  changed=0  unreachable=0  failed=0
```
Le test d'idempotence est particulièrement important : Molecule exécute le playbook une deuxième fois et vérifie qu'aucune tâche ne produit de changement. Si une tâche indique `changed` lors de la seconde exécution, cela signifie que le playbook n'est pas idempotent – un défaut qui peut entraîner des comportements imprévisibles lors de réexécutions en production.

## Étape 12 – Optimiser les Performances Ansible

À mesure que votre infrastructure grandit, les performances d'Ansible deviennent un enjeu critique. Un playbook qui prend 5 minutes sur 10 serveurs peut nécessiter 50 minutes sur 100 serveurs si les optimisations ne sont pas en place. Voici les techniques essentielles pour accélérer vos déploiements.

| Optimisation | Impact | Configuration | 
|---|---|---|
| Pipelining SSH | Réduction de 30-50 % du temps SSH | `pipelining = True` | 
| Fact caching (JSON) | Élimination du gathering sur les exécutions suivantes | `fact_caching = jsonfile` | 
| Forks (parallélisme) | Gestion de N serveurs simultanés | `forks = 20` | 
| Stratégie free | Chaque serveur avance indépendamment | `strategy: free` | 
| Mitogen (plugin) | Accélération de 1,25x à 7x | Plugin tiers, installation séparée | 
| Async tasks | Exécution en arrière-plan pour les tâches longues | `async: 300, poll: 5` | 
| gather_facts: false | Suppression de la collecte de facts inutile | Directive dans le play | 

```
# Playbook optimisé pour les déploiements à grande échelle
---
- name: Déploiement optimisé
  hosts: serveurs_web
  become: true
  strategy: free          # Chaque serveur avance à son rythme
  gather_facts: false     # Pas de gathering si les facts ne sont pas nécessaires
  serial: "30%"           # Déployer par lots de 30% des serveurs
  pre_tasks:
    - name: Collecter uniquement les facts réseau (plus rapide)
      ansible.builtin.setup:
        gather_subset:
          - network
          - hardware
  tasks:
    - name: Télécharger l'artefact (tâche longue en async)
      ansible.builtin.get_url:
        url: "https://releases.exemple.fr/app-{{ version_app }}.tar.gz"
        dest: /tmp/app.tar.gz
      async: 300
      poll: 0
      register: download_async
    - name: Pendant ce temps, préparer le répertoire
      ansible.builtin.file:
        path: /var/www/app-{{ version_app }}
        state: directory
        owner: www-data
        mode: '0755'
    - name: Attendre la fin du téléchargement
      ansible.builtin.async_status:
        jid: "{{ download_async.ansible_job_id }}"
      register: job_result
      until: job_result.finished
      retries: 30
      delay: 10
    - name: Extraire l'application
      ansible.builtin.unarchive:
        src: /tmp/app.tar.gz
        dest: /var/www/app-{{ version_app }}
        remote_src: true
    - name: Basculer le lien symbolique
      ansible.builtin.file:
        src: /var/www/app-{{ version_app }}
        dest: /var/www/app
        state: link
        force: true
      notify: Recharger Nginx
```
La directive `serial: "30%"` est cruciale pour les déploiements sans interruption : elle limite le déploiement simultané à 30 % de la flotte, garantissant que 70 % des serveurs restent disponibles à tout moment. Combinée avec la stratégie `free`, chaque serveur progresse à son rythme sans attendre les plus lents du lot.

Le pattern `async/poll` permet de lancer des tâches longues (téléchargements, compilations) en arrière-plan et de continuer avec d'autres tâches en parallèle. Cela peut diviser le temps total d'exécution par deux ou plus pour les playbooks qui incluent des opérations I/O intensives.

## 5 Pièges Courants à Éviter avec Ansible

Même les administrateurs systèmes expérimentés tombent dans ces pièges classiques d'Ansible. Connaître ces erreurs vous évitera des heures de débogage et des incidents en production.

**Piège n°1 : Oublier l'idempotence.** Utiliser le module `command` ou `shell` au lieu des modules Ansible dédiés rend vos playbooks non-idempotents. Par exemple, `command: apt install nginx` s'exécute à chaque fois, tandis que `ansible.builtin.apt: name=nginx state=present` ne fait rien si Nginx est déjà installé. Préférez toujours les modules spécialisés.

**Piège n°2 : Stocker des secrets en clair dans Git.** C'est la faille de sécurité la plus courante dans les projets Ansible. Utilisez systématiquement Ansible Vault pour chiffrer les mots de passe, clés API et certificats. Un audit de sécurité révèle souvent des fichiers `group_vars` contenant des mots de passe en clair – une pratique qui peut coûter cher en cas de fuite du dépôt.

**Piège n°3 : Négliger la gestion des erreurs.** Un playbook sans blocs `rescue` peut laisser un serveur dans un état incohérent en cas d'échec à mi-parcours. Utilisez toujours `block/rescue/always` pour les opérations critiques, et `--check` pour valider les changements avant de les appliquer.

**Piège n°4 : Utiliser `become: true` globalement.** Donner les privilèges root à toutes les tâches est une violation du principe du moindre privilège. Appliquez `become: true` uniquement aux tâches qui le nécessitent (installation de paquets, modification de fichiers système). Les tâches de copie de fichiers ou de configuration d'applications ne nécessitent généralement pas root.

**Piège n°5 : Ignorer les tags et les limites.** Exécuter un playbook entier pour une petite modification est risqué et lent. Utilisez les tags (`--tags`) pour cibler des sections spécifiques et `--limit` pour restreindre l'exécution à certains serveurs. En production, commencez toujours par un déploiement limité avant d'étendre à l'ensemble de l'infrastructure.

## Dépannage : 8+ Problèmes Fréquents et Solutions

Voici les problèmes les plus fréquemment rencontrés lors de l'utilisation d'Ansible, avec leurs solutions testées et validées.

