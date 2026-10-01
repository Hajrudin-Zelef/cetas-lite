---
id: collect-261001-automatisation-infra/automatisation-infra/tutoriel-ansible-2026-automatisation-en-12-etapes-vault-molecule-1
title: "Créer un environnement virtuel Python dédié"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "distribution"]
source: docs/RAG/collect-261001-automatisation-infra/tutoriel-ansible-2026-automatisation-en-12-etapes-vault-molecule.md
source_anchor: ""
source_lines: [1, 57]
sha256: aba55e36cdfaa777384805a6a82fe69b2fae9e2fe32c24a49f9b0193909edd43
---

# Créer un environnement virtuel Python dédié

Ansible s’impose comme l’outil d’automatisation IT le plus recherché en 2026 : plus de 63 000 étoiles sur GitHub, une adoption massive dans les entreprises européennes, et une branche 2.20 qui était encore désignée comme Current-Latest sur la page de statut du Ansible Community Package en août 2026 (Ansible Documentation) — avant que l’Ansible Forum n’annonce, le 8 septembre 2026, la sortie d’ansible-core 2.21.4 et du community package 14.4.0, nouvelle ligne de publication de référence du projet. Ce tutoriel vous guide pas à pas dans la maîtrise d’Ansible, depuis l’installation jusqu’au déploiement d’une infrastructure multi-serveurs complète. Vous apprendrez à écrire des playbooks, organiser vos rôles, gérer les secrets avec Ansible Vault, et automatiser le provisionnement de serveurs en production. À la fin de ce guide, vous disposerez d’un projet fonctionnel déployant une stack LAMP sécurisée sur plusieurs serveurs.

*Dernière mise à jour : 8 septembre 2026 – ansible-core 2.21.4 et community package 14.4.0 (annoncés le 8 septembre 2026 par l’Ansible Forum), qui succèdent à la branche 2.20 restée Current-Latest jusqu’en août 2026 selon la documentation Ansible*

## Prérequis : Versions et Environnement Nécessaires

Avant de commencer ce tutoriel Ansible, assurez-vous de disposer de l’environnement suivant. Chaque composant a été testé avec les versions indiquées pour garantir la compatibilité de l’ensemble du projet.

| Composant | Version requise | Vérification | 
|---|---|---|
| Python | 3.10 ou supérieur | `python3 --version` | 
| pip | 23.0 ou supérieur | `pip3 --version` | 
| Ansible | 13.5.0 (installé à l’étape 1) | `ansible --version` | 
| Ansible Core | 2.20.4 | `ansible-core --version` | 
| SSH | OpenSSH 8.0+ | `ssh -V` | 
| Système hôte (nœud de contrôle) | Ubuntu 22.04+ / Debian 12+ / macOS | `lsb_release -a` | 
| Serveurs cibles | 2 VMs ou serveurs (Ubuntu/Debian) | Accès SSH root ou sudo | 

**Configuration matérielle minimale** : le nœud de contrôle Ansible nécessite seulement 1 Go de RAM et 1 vCPU. Les serveurs cibles doivent disposer d’au moins 512 Mo de RAM et de Python 3 installé. Pour ce tutoriel, nous utiliserons deux serveurs cibles : un serveur web et un serveur de base de données.

**Note importante** : Ansible fonctionne sans agent – il se connecte aux serveurs cibles via SSH. Aucune installation n’est requise sur les machines distantes, contrairement à Puppet ou Chef qui nécessitent un agent dédié. C’est l’un des avantages majeurs d’Ansible, particulièrement apprécié dans les environnements cloud européens soumis aux contraintes de souveraineté numérique.

## Étape 1 – Installer Ansible sur le Nœud de Contrôle

La première étape consiste à installer Ansible sur votre machine de contrôle. Ansible 13.5.0, publié le 25 mars 2026, reste la base stable utilisée dans ce tutoriel, mais le paysage des versions a continué d’évoluer depuis : ansible-core 2.20 était encore listé comme la version Current-Latest par la documentation officielle Ansible en août 2026, avant que l’Ansible Forum n’annonce, le 8 septembre 2026, la sortie d’ansible-core 2.21.4, qui prolonge désormais la ligne de publication communautaire au-delà de la branche 2.20. Red Hat continue par ailleurs de maintenir les branches précédentes pour les projets qui ne sont pas encore passés à la dernière version. L’installation via pip est la méthode recommandée car elle garantit l’accès à la version la plus récente, indépendamment des dépôts de votre distribution Linux.

```
# Créer un environnement virtuel Python dédié
python3 -m venv ~/ansible-env
source ~/ansible-env/bin/activate
# Installer Ansible 13 via pip
pip install ansible==13.5.0
# Vérifier l'installation
ansible --version
# ansible [core 2.20.4]
#   config file = None
#   configured module search path = ['~/.ansible/plugins/modules']
#   ansible python module location = ~/ansible-env/lib/python3.12/site-packages/ansible
#   executable location = ~/ansible-env/bin/ansible
#   python version = 3.12.x
# Vérifier la connectivité locale
ansible localhost -m ping
# localhost | SUCCESS => {
#     "changed": false,
#     "ping": "pong"
# }
```
**Alternative sur Ubuntu/Debian** : vous pouvez aussi utiliser le PPA officiel d’Ansible, mais la version disponible peut être en retard par rapport à pip. Pour les distributions Red Hat (RHEL, CentOS, Fedora), la plateforme Red Hat Ansible Automation Platform a continué d’évoluer tout au long de 2026 : Red Hat a expédié AAP 2.7 dès le 8 juin 2026, avant sa disponibilité générale officielle le 16 juin 2026 (Red Hat Developer), marquant la génération actuelle de la plateforme. La release de maintenance 2.7.20260824, publiée le 24 août 2026, a livré pas moins de 25 correctifs de sécurité, 16 améliorations de performance et d’exploitation, et 31 corrections de bugs selon la documentation Red Hat. Les branches précédentes restent activement maintenues : AAP 2.6 a reçu le 4 août 2026 un correctif regroupant 35 correctifs CVE ainsi que 10 améliorations de performance et de fonctionnalités (Red Hat Docs). Cette plateforme offre un support entreprise complet dans les environnements d’exécution, bien au-delà de ce que propose le simple PPA.

L’utilisation d’un environnement virtuel Python est fortement recommandée pour isoler Ansible et ses dépendances du système. Cela évite les conflits de paquets et permet de gérer facilement plusieurs versions d’Ansible sur la même machine – un point crucial dans les environnements de CI/CD où différents projets peuvent nécessiter des versions différentes.

## Étape 2 – Configurer l’Inventaire des Serveurs

L’inventaire Ansible définit les serveurs cibles et leur organisation en groupes. C’est le fichier de référence qui indique à Ansible quelles machines gérer. Vous pouvez utiliser le format INI (simple) ou YAML (structuré). Pour ce tutoriel, nous utiliserons le format YAML qui offre une meilleure lisibilité et prend en charge les structures imbriquées.

Créez un répertoire de projet et le fichier d’inventaire :

