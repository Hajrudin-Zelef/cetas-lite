---
id: collect-261001-general-networking/general-networking/gophish-simuler-une-campagne-de-phishing-en-13-etapes-1
title: "1. Mise à jour système"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["dpo", "mai", "open source", "valuation"]
source: docs/RAG/collect-261001-general-networking/gophish-simuler-une-campagne-de-phishing-en-13-etapes.md
source_anchor: ""
source_lines: [1, 45]
sha256: a50edff3a06d46d73e64ed8a7aa81ec083dc8e384258fe373472d559e822b05e
---

# 1. Mise à jour système

Simuler une campagne de phishing en interne n’est plus réservé aux grandes banques ou aux ministères. Avec GoPhish, un framework open source écrit en Go, n’importe quelle PME ou DSI peut monter en quelques heures un exercice de sensibilisation réaliste, mesurer qui clique, qui saisit ses identifiants et qui signale l’email suspect. Le contexte y pousse : dans son panorama de la cybermenace 2025 publié le 13 mai 2026, l’ANSSI note que les attaquants s’appuient de plus en plus sur des services commerciaux et open source détournés pour cibler les entités françaises, souvent via l’ingénierie sociale. Ce tutoriel détaille, étape par étape, l’installation de GoPhish sur un serveur Linux, la configuration d’une campagne complète et les pièges à éviter, dans un cadre strictement légal et autorisé.

**Avertissement** : cet article s’adresse aux RSSI, administrateurs systèmes et équipes de sensibilisation qui testent leur propre organisation avec une autorisation écrite de la direction. Utiliser GoPhish contre des cibles qui n’ont pas donné leur accord constitue une infraction pénale en France (article 323-1 du Code pénal). Ne déployez jamais cet outil hors d’un périmètre autorisé.

## Pourquoi simuler des campagnes de phishing en 2026

Le phishing reste le premier vecteur d’intrusion en entreprise, et il change de nature. Les attaquants automatisent désormais la rédaction de leurs emails avec des générateurs de texte, ce qui réduit les fautes d’orthographe qui trahissaient autrefois les tentatives frauduleuses. Résultat : les employés distinguent de moins en moins bien un email légitime d’un email piégé. Une campagne de simulation GoPhish permet de mesurer ce risque concrètement, avec des chiffres exploitables plutôt qu’une simple formation théorique.

L’intérêt de GoPhish tient à trois choses. D’abord, c’est un outil entièrement gratuit et open source, disponible sur GitHub, sans licence à négocier. Ensuite, il tourne aussi bien sur une petite VM Linux que sur un serveur dédié, avec une empreinte système minimale puisqu’il est écrit en Go et distribué sous forme de binaire unique. Enfin, son modèle de données (profils d’envoi, modèles d’email, pages d’atterrissage, groupes de cibles, campagnes) reproduit fidèlement le déroulé d’une vraie attaque de phishing, ce qui rend les résultats pédagogiques pour les collaborateurs testés.

Ce tutoriel couvre l’installation depuis zéro, la sécurisation du service (TLS, pare-feu, utilisateur dédié, systemd), la création d’une campagne complète avec profil SMTP, modèle d’email et page d’atterrissage, l’analyse des résultats, ainsi que le cadre juridique français à respecter avant tout lancement.

Il ne s’agit pas d’un exercice ponctuel à cocher une fois par an pour satisfaire un audit. Les équipes qui obtiennent les meilleurs résultats traitent la simulation de phishing comme un cycle continu : mesurer, former, remesurer. Le rapport 2025 de l’ANSSI souligne d’ailleurs que l’ingénierie sociale reste le point d’entrée privilégié des attaquants, précisément parce qu’elle contourne les protections techniques classiques comme les pare-feux ou les antivirus. Un firewall bien configuré ne sert à rien si un comptable saisit ses identifiants sur une fausse page de connexion. C’est cette faille humaine, difficile à corriger par des correctifs logiciels, que GoPhish permet de quantifier puis de réduire progressivement.

## Comment fonctionne l’architecture de GoPhish

Avant de se lancer dans l’installation, il est utile de comprendre comment GoPhish organise ses composants, car cela évite ensuite bien des erreurs de configuration. Le binaire unique embarque en réalité deux serveurs HTTP distincts qui tournent dans le même processus mais écoutent sur des ports différents.

Le premier, le serveur d’administration, expose l’interface web depuis laquelle vous pilotez tout : création des profils SMTP, des modèles d’email, des pages d’atterrissage, des groupes de cibles et des campagnes. Cette interface communique avec une API REST interne, la même que vous pouvez interroger directement en HTTP si vous souhaitez automatiser la création de campagnes depuis un script externe plutôt que de tout faire à la souris.

Le second, le serveur de phishing, est celui qui reçoit les visites des cibles lorsqu’elles cliquent sur le lien contenu dans l’email. C’est lui qui génère l’identifiant unique par destinataire (visible dans l’URL sous forme de paramètre), qui sert la page d’atterrissage correspondante, et qui enregistre chaque étape du parcours de la cible dans la base de données. Les deux serveurs partagent la même base SQLite par défaut (`gophish.db`), ce qui explique pourquoi un problème de permissions sur ce fichier peut faire planter l’ensemble de l’application, pas seulement une partie.

Cette séparation en deux serveurs a une conséquence pratique directe pour le déploiement : vous pouvez exposer publiquement uniquement le serveur de phishing tout en gardant le serveur d’administration strictement privé, ce qui réduit considérablement la surface d’attaque de votre propre outil de test.

## Prérequis techniques et légaux

Avant de commencer, réunissez les éléments suivants. Ils conditionnent la réussite de l’installation et surtout la légalité de la campagne.

- Un serveur Linux (Debian 12, Ubuntu 22.04/24.04 ou équivalent) avec au moins 1 vCPU, 1 Go de RAM et 10 Go de disque
- Un accès root ou sudo sur ce serveur
- GoPhish en version 0.12.1, la dernière publiée sur le dépôt officiel GitHub
- Un relais SMTP autorisé à envoyer des emails pour votre domaine de test (serveur interne, ou service tiers configuré avec SPF/DKIM)
- Un nom de domaine ou sous-domaine dédié à la simulation (jamais le domaine de production de l’entreprise)
- Un certificat TLS valide pour l’interface d’administration (Let’s Encrypt convient)
- Une autorisation écrite signée par la direction générale et le DPO/RSSI, précisant le périmètre, la durée et les destinataires de la campagne
- Une liste de cibles au format CSV avec au minimum une colonne email

Sur le plan juridique, le guide RGPD de la CNIL sur la sécurité des données personnelles rappelle que toute collecte de données comportementales sur les employés (qui a cliqué, qui a saisi ses identifiants) doit respecter les principes de proportionnalité, de minimisation et de transparence. Concrètement, informez les salariés en amont qu’une campagne de sensibilisation incluant d’éventuelles simulations de phishing peut avoir lieu, consultez le CSE si nécessaire, et ne réutilisez jamais les résultats individuels à des fins disciplinaires. Consultez le guide pratique RGPD de la CNIL avant de lancer votre première campagne.

En pratique, la plupart des DSI structurent leur autorisation autour de quatre points fixes : le périmètre exact des personnes testées (par service, par site, ou l’ensemble de l’organisation), la durée de la campagne, le type de scénario simulé (facture frauduleuse, fausse alerte RH, fausse mise à jour de mot de passe), et la finalité déclarée de la collecte de données, qui doit rester la formation et jamais l’évaluation individuelle des salariés. Ce document signé sert aussi de preuve en cas de question du CSE ou d’un salarié qui découvrirait, après coup, qu’il a été ciblé par un test.

## Étape 1 : préparer le serveur Linux

Commencez par mettre à jour le système et installer les paquets de base nécessaires au téléchargement et à la décompression du binaire GoPhish.

