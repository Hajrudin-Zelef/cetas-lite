---
id: collect-261001-automatisation-infra/automatisation-infra/tutoriel-ansible-2026-automatisation-en-12-etapes-vault-molecule-7
title: "Créer un environnement virtuel Python dédié"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: ["open source"]
source: docs/RAG/collect-261001-automatisation-infra/tutoriel-ansible-2026-automatisation-en-12-etapes-vault-molecule.md
source_anchor: ""
source_lines: [966, 1050]
sha256: 783d3d363156425d56206853e95eb971d0ff375114b05d5d8ff7dd87a3a9c08f
---

# Créer un environnement virtuel Python dédié

| Problème | Cause | Solution | 
|---|---|---|
| `UNREACHABLE!` – Host unreachable | Clé SSH non configurée ou mauvais utilisateur | Vérifier `ssh -i ~/.ssh/id_ed25519 user@host` . Ajouter`ansible_ssh_private_key_file` dans l'inventaire. | 
| `Permission denied (publickey)` | L'utilisateur n'a pas de clé autorisée sur le serveur cible | Copier la clé : `ssh-copy-id user@host` . Vérifier les permissions de`~/.ssh/authorized_keys` (600). | 
| `Missing sudo password` | `become: true` sans accès sudo sans mot de passe | Ajouter `deploy ALL=(ALL) NOPASSWD:ALL` dans`/etc/sudoers.d/deploy` ou utiliser`--ask-become-pass` . | 
| `Module not found` | Collection manquante dans l'environnement | Installer la collection : `ansible-galaxy collection install community.general` . | 
| `Vault password not provided` | Fichier Vault chiffré sans mot de passe | Ajouter `--ask-vault-pass` ou configurer`vault_password_file` dans`ansible.cfg` . | 
| `changed` à chaque exécution | Module non-idempotent (shell/command) | Utiliser les modules dédiés ou ajouter `creates:` /`removes:` pour vérifier l'état. | 
| Playbook très lent | Gathering de facts complet sur chaque exécution | Activer `fact_caching` , utiliser`gather_subset` , ou`gather_facts: false` . | 
| `Jinja2 template error` | Variable non définie dans le contexte du template | Utiliser le filtre `{{ variable \| default('valeur') }}` pour définir des valeurs par défaut. | 
| `YAML syntax error` | Indentation incorrecte ou caractères spéciaux | Utiliser `yamllint` et`ansible-lint` . Les chaînes avec`:` doivent être entre guillemets. | 
| `could not find or access` un fichier template | Chemin relatif incorrect vers le template | Les templates doivent être dans `roles/NOM/templates/` ou spécifier le chemin absolu. | 

**Commandes de débogage essentielles** : utilisez `-vvv` (triple verbose) pour voir les connexions SSH détaillées, `-vvvv` pour le débogage complet incluant les scripts envoyés aux serveurs, et `--step` pour exécuter les tâches une par une avec confirmation. La commande `ansible-playbook --list-tasks` affiche toutes les tâches sans les exécuter, utile pour vérifier la logique avant un déploiement.

## Astuces Avancées pour Ansible en Production

Ces techniques avancées vous permettront d'exploiter pleinement la puissance d'Ansible dans des environnements de production exigeants. Elles sont utilisées par les équipes DevOps les plus performantes pour gérer des centaines, voire des milliers de serveurs.

**Callback plugins pour l'observabilité.** Les plugins de callback envoient les résultats d'exécution vers des systèmes externes. Le plugin `community.general.slack` notifie un canal Slack à chaque déploiement, tandis que `ara` (ARA Records Ansible) stocke l'historique complet dans une base de données consultable via une interface web. Configurez-les dans `ansible.cfg` avec `callbacks_enabled = ara_default, slack`.

**Dynamic includes vs static imports.** La différence entre `include_tasks` (dynamique) et `import_tasks` (statique) est subtile mais importante. Les imports sont résolus au parsing du playbook, ce qui permet l'utilisation de tags et de `--list-tasks`. Les includes sont résolus à l'exécution, ce qui permet l'utilisation de boucles et de conditions dynamiques. Utilisez `import` par défaut et `include` uniquement quand vous avez besoin de comportement dynamique.

**Custom modules en Python.** Quand aucun module existant ne correspond à votre besoin, vous pouvez créer un module personnalisé en Python. Placez-le dans `library/` à la racine du projet ou dans `roles/NOM/library/`. Le module reçoit les arguments en JSON et retourne un résultat structuré. Cette approche est particulièrement utile pour interagir avec des API internes ou des systèmes propriétaires.

**Ansible Navigator.** Ansible Navigator (`ansible-navigator`) est le successeur de `ansible-playbook` en ligne de commande. Il offre une interface interactive avec des vues détaillées pour chaque tâche, le support natif des execution environments (conteneurs), et une meilleure gestion des collections. Pour les utilisateurs de la plateforme Red Hat Ansible Automation Platform, il remplace progressivement les outils CLI traditionnels – d'autant plus que le patch 2.6.20260325, livré le 25 mars 2026, fait évoluer l'écosystème avec Automation Controller 4.7.19, Automation Hub 4.11.7 et Event-Driven Ansible 1.2.7.

## Structure Complète du Projet

Voici la structure finale du projet Ansible complet que nous avons construit au fil de ce tutoriel. Cette organisation suit les bonnes pratiques recommandées par la documentation officielle d'Ansible et peut être versionnée dans un dépôt Git.

```
ansible-projet/
├── ansible.cfg                          # Configuration Ansible du projet
├── inventaire/
│   ├── production.yml                   # Inventaire de production
│   └── staging.yml                      # Inventaire de staging
├── group_vars/
│   ├── all.yml                          # Variables communes
│   ├── serveurs_web.yml                 # Variables des serveurs web
│   └── serveurs_bdd/
│       ├── vars.yml                     # Variables non-sensibles
│       └── vault.yml                    # Secrets chiffrés (Vault)
├── host_vars/
│   └── web01.yml                        # Variables spécifiques à web01
├── playbooks/
│   ├── securite-base.yml                # Hardening de sécurité
│   ├── deploiement-complet.yml          # Déploiement full-stack
│   └── conditions-avancees.yml          # Structures de contrôle
├── roles/
│   ├── securite-base/
│   │   ├── tasks/main.yml
│   │   ├── handlers/main.yml
│   │   └── templates/
│   └── serveur_web/
│       ├── defaults/main.yml
│       ├── tasks/main.yml
│       ├── handlers/main.yml
│       ├── templates/
│       │   ├── vhost.conf.j2
│       │   └── php-fpm-pool.conf.j2
│       └── molecule/
│           └── default/
│               ├── molecule.yml
│               ├── converge.yml
│               └── verify.yml
├── fichiers/                            # Fichiers statiques
├── templates/                           # Templates globaux
├── .github/
│   └── workflows/
│       └── deploy.yml                   # Pipeline CI/CD
├── requirements.yml                     # Collections Galaxy requises
├── .gitignore                           # Exclure .vault_pass, *.retry
└── README.md
```
Créez un fichier `requirements.yml` pour documenter les collections et rôles Galaxy nécessaires à votre projet. Cette pratique garantit que tout développeur peut reproduire l'environnement avec un simple `ansible-galaxy install -r requirements.yml`. Incluez les versions exactes pour assurer la reproductibilité des déploiements.

### Couverture associée

## FAQ – Questions Fréquentes sur Ansible

### Ansible est-il gratuit ?

Oui, Ansible (le projet open source) est entièrement gratuit et distribué sous licence GPL v3. La version communautaire disponible via pip ou les gestionnaires de paquets Linux inclut toutes les fonctionnalités essentielles : playbooks, rôles, Vault, inventaires dynamiques et modules. Red Hat propose la plateforme Ansible Automation Platform (AAP) en version payante, qui ajoute une interface web, le contrôle d'accès par rôle, la gestion centralisée des inventaires et le support entreprise – la lignée 2.5 a par exemple reçu le correctif 2.5-22 dès le 1er mars 2025, avant qu'AAP 2.6 ne soit diffusé le 8 octobre 2025 avec des bundles x86_64 de 3,74 Go et 2,96 Go. Pour la majorité des cas d'usage, la version communautaire est amplement suffisante.

