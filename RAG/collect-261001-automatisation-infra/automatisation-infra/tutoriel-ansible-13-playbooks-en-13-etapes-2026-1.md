---
id: collect-261001-automatisation-infra/automatisation-infra/tutoriel-ansible-13-playbooks-en-13-etapes-2026-1
title: "Mise à jour des paquets et installation de pipx"
domain: automatisation-infra
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["agent", "arr", "aws", "distribution", "mai"]
source: docs/RAG/collect-261001-automatisation-infra/tutoriel-ansible-13-playbooks-en-13-etapes-2026.md
source_anchor: ""
source_lines: [1, 32]
sha256: 9ee4f50512d502bd0a0db0ddecc73a1d64e343ed7cf69264e927e9f793b030c5
---

# Mise à jour des paquets et installation de pipx

Publié le 28 avril 2026 — Ansible reste l’outil de référence pour automatiser le provisionnement de serveurs, le déploiement d’applications et la gestion de configuration à grande échelle. Avec la sortie d’**ansible-core 2.20.5 le 21 avril 2026** et du paquet communautaire **Ansible 13.6.0** le même jour, l’écosystème entre dans une phase de maturité où les playbooks YAML pilotent aussi bien des fermes de VM Proxmox que des clusters Kubernetes managés sur AWS. La documentation officielle d’Ansible, mise à jour en août 2026, confirme d’ailleurs un rythme de sortie d’environ deux versions majeures par an, signe d’un projet qui continue d’accélérer plutôt que de ralentir. Ce tutoriel guide les administrateurs systèmes et ingénieurs DevOps français à travers **13 étapes pratiques**, depuis l’installation jusqu’au déploiement d’un projet complet de stack web LAMP avec rôles, vault et inventaire dynamique.

Ce guide cible les ingénieurs systèmes, SRE et architectes cloud qui souhaitent industrialiser leurs déploiements sans agent. Vous apprendrez à structurer un projet selon les bonnes pratiques officielles de la documentation Ansible, à écrire des playbooks idempotents, à chiffrer vos secrets avec Ansible Vault et à exécuter vos rôles sur 100 hôtes Linux ou Windows en parallèle. À la fin, vous disposerez d’un projet réutilisable pour automatiser un cluster web complet en moins de 4 minutes.

## Pourquoi Ansible domine l’automatisation IT en 2026

Ansible s’est imposé comme la lingua franca de l’automatisation infrastructure depuis le rachat par Red Hat le 16 octobre 2015. En avril 2026, le dépôt officiel `ansible/ansible` sur GitHub franchit la barre des 68 000 étoiles et reste l’un des projets Python les plus actifs au monde — une popularité qui se traduit très concrètement en entreprise puisque, selon les données 6sense pour l’année 2025, plus de 2 734 entreprises dans le monde s’appuyaient déjà sur Red Hat Ansible Automation Platform pour industrialiser leurs chaînes de build. Contrairement à Puppet ou Chef, qui exigent un agent installé sur chaque nœud cible, Ansible fonctionne en **mode push agentless** via SSH (ou WinRM/PSRP pour Windows), ce qui réduit drastiquement la surface d’attaque et la maintenance — un argument de poids quand on sait qu’une enquête Red Hat menée en juillet 2026 auprès de 900 décideurs IT révèle que plus de 50 % d’entre eux évitent encore les plateformes d’automatisation jugées trop complexes.

Cette simplicité explique l’adoption massive en France et en Europe. Les administrateurs de la SNCF, d’Orange, du Crédit Agricole et de l’Éducation nationale s’appuient sur des playbooks Ansible pour déployer des milliers de VM RHEL et Ubuntu chaque semaine. Selon les baromètres salariaux 2026 publiés par Free-Work et LeHibou, un ingénieur Ansible confirmé en Île-de-France facture entre **55 000 et 78 000 euros bruts annuels**, avec des pointes à 95 000 euros pour les profils certifiés Red Hat Certified Engineer (RHCE) maîtrisant l’**Ansible Automation Platform 2.6**.

L’arrivée d’**Event-Driven Ansible** dans AAP 2.6 — sortie officiellement en octobre 2025 et embarquant ansible-core 2.16 sous le capot, selon l’API endoflife.date —, qui déclenche des playbooks en réaction à des événements observés (alertes Prometheus, webhooks GitHub, syslog), pousse l’outil au-delà du simple provisionnement — une bascule que le terrain confirmait déjà : lors du sommet Ansible 2025, une enquête Steampunk.si publiée en mai 2025 révélait que 63 % des répondants utilisaient Ansible Automation Platform comme distribution Ansible de référence, loin devant le simple ansible-core communautaire. Une analyse Red Hat publiée en janvier 2026 souligne que cette version a aussi apporté une interface utilisateur repensée et des fonctionnalités d’inventaire assistées par IA. Cette boucle de remédiation automatique transforme Ansible en couche opérationnelle réactive, capable de redémarrer un service, de retirer un nœud d’un load balancer ou de scaler une stack en quelques secondes après détection d’une anomalie. Et la suite est déjà annoncée : selon CRN, Red Hat prévoit la disponibilité générale d’**Ansible Automation Platform 2.7** dans les prochaines semaines après mai 2026, avec un orchestrateur d’automatisation piloté par IA proposé en tech preview plus tard dans l’année.

## Prérequis : versions, OS et machines cibles

Avant d’attaquer le premier playbook, vérifiez que votre poste de contrôle et vos hôtes cibles respectent la matrice officielle de compatibilité d’ansible-core 2.20. Cette **nouvelle ligne 2.20**, passée en disponibilité générale le 3 novembre 2025 et désormais marquée « Current – Latest » par la documentation Ansible — au même titre que le paquet communautaire Ansible 2.20, lui aussi étiqueté « Current – Latest » depuis août 2026 —, prend en charge Python 3.12 à 3.14 sur le contrôleur et bénéficie d’un support de sécurité garanti jusqu’en novembre 2026, contre Python 3.11 à 3.13 pour la branche 2.19 précédente, sortie en GA le 21 juillet 2025. C’est ce qui en fait le choix recommandé pour tout nouveau projet en production. La documentation officielle, mise à jour en août 2026, confirme que la branche 2.20.x reste la version stable de référence et recommande spécifiquement ansible-core 2.20.4 comme base par défaut pour tout playbook démarré aujourd’hui. Quant à la 2.19, elle reste largement déployée et n’est pas encore obsolète : passée de la révision 2.19.3 en octobre 2025 à son dernier correctif connu, la 2.19.13, le 8 septembre 2026 selon endoflife.date, son support GA s’est arrêté le 19 mai 2026 pour une fin de vie complète programmée au 30 novembre 2026 — seuls les utilisateurs encore sur la 2.18 doivent migrer sans attendre, cette branche ayant atteint sa fin de vie dès décembre 2025.

| Composant | Version minimale | Version recommandée 2026 | Notes | 
|---|---|---|---|
| ansible-core | 2.18 | 2.20.5 | Sortie le 21 avril 2026 | 
| Paquet Ansible community | 11 | 13.6.0 | Inclut 100+ collections | 
| Python (contrôleur) | 3.12 | 3.13 | 3.14 supporté | 
| Python (cibles Linux) | 3.9 | 3.12 | 3.14 supporté | 
| OpenSSH | 7.6 | 9.x | Authentification par clé obligatoire | 
| PowerShell (cibles Windows) | 5.1 | 7.4 | Module win_* requis | 
| Sudo / become | 1.8 | 1.9.13 | Mode non interactif | 

Pour ce tutoriel, vous aurez besoin d’un poste de contrôle sous Ubuntu 24.04 LTS, Debian 12 ou Fedora 40, ainsi que de trois VM cibles (deux serveurs web et un serveur de base de données). Vous pouvez les provisionner en local avec Vagrant, sur Proxmox VE 9.1 ou directement sur AWS/Scaleway. Comptez environ 4 Go de RAM sur le contrôleur et 1 Go par VM cible. Une connectivité SSH par clé entre le contrôleur et chaque cible est obligatoire, l’authentification par mot de passe étant fortement déconseillée en 2026.

## Étape 1 : Installer Ansible 13 sur le poste de contrôle

L’installation officielle recommandée passe par `pipx`, qui isole Ansible dans un environnement Python dédié et évite les conflits avec les paquets système. Sur Ubuntu 24.04 ou Debian 12, exécutez les commandes suivantes pour disposer du paquet community 13.6.0 complet. Cette méthode fonctionne identiquement sur Fedora, RHEL 9, openSUSE et même WSL 2 sous Windows 11.

