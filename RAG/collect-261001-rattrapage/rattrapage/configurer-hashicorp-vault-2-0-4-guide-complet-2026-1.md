---
id: collect-261001-rattrapage/rattrapage/configurer-hashicorp-vault-2-0-4-guide-complet-2026-1
title: "Vault v2.0.4, built with go1.23"
domain: rattrapage
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["agent", "aws", "dpo", "incident", "license", "open source", "packaging"]
source: docs/RAG/collect-261001-rattrapage/configurer-hashicorp-vault-2-0-4-guide-complet-2026.md
source_anchor: ""
source_lines: [1, 44]
sha256: 807cf6033c8d8e5380e0aaaf74c4e85a4265067dc075549ef212a2a5d3e3f38b
---

# Vault v2.0.4, built with go1.23

Une clé API oubliée dans un dépôt Git public, un identifiant de base de données codé en dur dans une image Docker, un token AWS qui traîne dans une variable d’environnement CI/CD : c’est exactement ce type d’erreur qui a exposé des clés d’API AWS chez Braintrust début 2026, un incident que nous avions détaillé sur Tech Insider. Le problème n’est pas nouveau, mais son coût grimpe chaque année. HashiCorp Vault reste l’outil de référence pour centraliser, chiffrer et faire tourner ces secrets automatiquement, plutôt que de les laisser traîner dans des fichiers de configuration.

Ce tutoriel vous accompagne pas à pas pour installer, configurer et exploiter HashiCorp Vault en 2026, avec la version 2.1.0 (la dernière en date, passée en disponibilité générale le 2 septembre 2026 avec sa nouvelle interface Agent Registry UI, selon un billet du blog communautaire IBM et la page officielle des releases HashiCorp mise à jour le 15 septembre 2026), les nouveautés de licence depuis le rachat par IBM, le fork open source OpenBao, et les intégrations Kubernetes et Terraform qu’une équipe DevOps ou une RSSI française doit connaître avant de se lancer.

## Pourquoi centraliser ses secrets en 2026 : le contexte français et européen

Le stockage de secrets (mots de passe de base de données, clés API, certificats TLS, tokens cloud) dans des fichiers de configuration en clair reste l’une des causes les plus fréquentes de fuites de données. Un secret statique, copié dans un fichier YAML ou une image Docker, vit potentiellement pour toujours : il est recopié, archivé dans l’historique Git, dupliqué sur des postes de développeurs. Un secret dynamique, généré à la demande par un coffre-fort central et révoqué après usage, réduit drastiquement la fenêtre d’exposition.

Pour les entreprises soumises au RGPD, la gestion centralisée des secrets s’inscrit directement dans l’obligation de mettre en œuvre des « mesures techniques et organisationnelles appropriées » pour protéger les données personnelles. Un chiffrement au repos et en transit, une rotation automatisée des identifiants et un accès minimal par politique (principe du moindre privilège) sont des arguments que les délégués à la protection des données (DPO) et les RSSI documentent de plus en plus souvent dans leurs analyses d’impact (AIPD/DPIA). Le sujet est également suivi de près par les régulateurs : la CNIL rappelle régulièrement que la sécurisation des accès techniques fait partie intégrante des exigences du RGPD, même si le texte ne cite aucun outil en particulier.

Sur le plan strictement technique, l’écosystème a bougé en 2025-2026. HashiCorp a été racheté par IBM pour environ 6,4 milliards de dollars, une opération finalisée en février 2025. Depuis, Vault suit le cycle de support IBM : la branche 1.21.0, sortie en mars 2026 avec l’authentification SPIFFE native (rapportée par InfoQ le 28 mars 2026), a rapidement cédé la place à la version majeure 2.0, passée en disponibilité générale le 14 avril 2026 selon les notes de version officielles — la toute première nouvelle version majeure depuis le lancement de Vault 1.0 en 2018. IBM a par ailleurs publié en janvier 2026 une annexe de cycle de vie support précisant que les branches Vault Self-Managed 1.16, 1.19, 1.20 et 1.21 resteront maintenues entre le 30 avril 2026 et le 30 avril 2027 pour les organisations qui n’ont pas encore basculé vers la 2.x. Parallèlement, la licence Business Source License (BSL 1.1), adoptée par HashiCorp dès août 2023 pour l’édition Community, reste en vigueur : Vault n’est plus un logiciel librement redistribuable au sens strict, ce qui a motivé la naissance du fork OpenBao, hébergé par la Linux Foundation sous licence MPL 2.0. Nous détaillerons ce choix un peu plus loin, car il conditionne directement quelle édition installer.

## Prérequis : versions, matériel et comptes nécessaires

Avant de commencer, assurez-vous de disposer de l’environnement suivant. Ce tutoriel a été testé avec les versions disponibles au 20 août 2026.

| Composant | Version recommandée (août 2026) | Remarque | 
|---|---|---|
| HashiCorp Vault Community | 2.0.4 (publiée le 3-4 août 2026) | Licence BSL 1.1, gratuite pour un usage interne | 
| HashiCorp Vault Enterprise | 2.0.4 | Licence commerciale, namespaces et HSM inclus | 
| OpenBao (alternative open source) | Branche 2.x active | Fork Linux Foundation, licence MPL 2.0 | 
| Système d’exploitation | Ubuntu 24.04 LTS ou Debian 12 | Fonctionne aussi sur RHEL 9 et Windows Server 2022 | 
| Docker | 27.x ou supérieur | Pour le déploiement en conteneur | 
| Terraform | 1.9.x ou supérieur | Pour le provider Vault et l’automatisation IaC | 
| vault-k8s (intégration Kubernetes) | 1.7.6 (publiée le 6 août 2026) | Injecteur d’agent et sidecar pour secrets | 
| Kubernetes | 1.30 ou supérieur | Cluster EKS, AKS, GKE ou on-premise | 

Côté matériel, un serveur de test avec 2 vCPU et 4 Go de RAM suffit largement pour un cluster Vault en mode développement. Pour une mise en production réelle, HashiCorp recommande un minimum de 4 vCPU, 8 Go de RAM et un stockage SSD chiffré pour le backend de persistance (Raft intégré ou Consul). Vous aurez également besoin d’un compte cloud (AWS, Azure ou GCP) si vous comptez tester les moteurs de secrets dynamiques, ainsi que d’un accès administrateur à votre cluster Kubernetes si vous souhaitez suivre la partie intégration.

Prévoyez également un nom de domaine interne valide (par exemple `vault.entreprise.fr`) et un certificat TLS, auto-signé pour les tests ou émis par une autorité interne pour la production : Vault refuse par défaut toute connexion API en clair dès que le mode développement est désactivé, et c’est une bonne pratique à conserver. Si vous testez l’auto-unseal avec AWS KMS, créez à l’avance une clé symétrique dédiée dans la région européenne de votre choix (Paris ou Francfort typiquement), avec une politique IAM restreinte au seul rôle utilisé par Vault. Comptez enfin un accès SSH ou une console cloud pour au moins deux ou trois machines si vous voulez tester un cluster Raft en haute disponibilité plutôt qu’un nœud unique, ce qui correspond à la configuration réellement recommandée avant toute mise en production.

## Étape 1 : choisir son édition — Community, Enterprise, HCP Vault ou OpenBao

C’est la première décision, et elle a des conséquences sur toute la suite du projet. En 2026, le paysage s’est complexifié avec quatre options distinctes.

| Édition | Licence | Coût | Cas d’usage recommandé | 
|---|---|---|---|
| Vault Community | BSL 1.1 | Gratuit | Petites équipes, cluster auto-hébergé unique | 
| Vault Enterprise | Commerciale | Payant (devis) | Namespaces, réplication multi-site, intégration HSM, MFA plateforme | 
| HCP Vault Dedicated | SaaS géré par HashiCorp/IBM | À partir de ~450 $/mois (cluster dev extra-small) | Équipes qui ne veulent pas opérer l’infrastructure elles-mêmes | 
| OpenBao | MPL 2.0 (Linux Foundation) | Gratuit | Organisations exigeant une licence open source classique | 

OpenBao mérite une explication : le projet a été forké à partir de la dernière version de Vault encore sous licence MPL 2.0, la 1.14.0, avant le basculement vers la BSL. En 2026, OpenBao vit sa propre vie sur une branche 2.x, pilotée par un comité technique neutre sous l’égide de la Linux Foundation. Pour une administration publique française ou une entreprise dont la politique interne interdit les licences BSL, c’est aujourd’hui l’alternative la plus mature. La compatibilité fonctionnelle avec l’API Vault reste élevée, donc la majorité des commandes de ce tutoriel s’appliquent aussi à OpenBao, à quelques différences de packaging près.

