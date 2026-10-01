---
id: collect-261001-general-networking/general-networking/passkeys-fini-le-mot-de-passe-en-13-etapes-2026-1
title: "Installe une autorité de certification locale de confiance"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Apple", "Google", "Microsoft"]
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-general-networking/passkeys-fini-le-mot-de-passe-en-13-etapes-2026.md
source_anchor: ""
source_lines: [1, 34]
sha256: c50446e0941806f29002ae1ecb51374c6543668722cfd1860f25542782dbbfdc
---

# Installe une autorité de certification locale de confiance

Vingt-quatre milliards d’identifiants ont fuité en 2026, une bonne partie encore réutilisable pour se connecter à de vrais comptes, comme tech-insider.org l’a documenté en détail. Chaque nouvelle fuite pointe vers le même coupable : le mot de passe. Il se devine, se phishe, se réutilise et finit tôt ou tard dans une base de données volée. Les passkeys proposent une sortie technique à ce problème, pas une simple rustine. Ce tutoriel vous montre comment les activer sur vos comptes existants, puis comment construire de A à Z un système d’authentification par passkey dans une application Node.js, avec le code complet, les pièges à éviter et les réponses aux pannes les plus fréquentes.

Le sujet dépasse largement le confort de connexion. Pour une équipe technique française, les passkeys touchent à la fois la sécurité des comptes utilisateurs, le coût du support (réinitialisations de mot de passe, tickets liés au phishing) et, de plus en plus, la conformité réglementaire sous NIS2. Ce guide part du geste le plus simple, activer un passkey sur un compte que vous possédez déjà, jusqu’au plus technique, écrire le code serveur qui vérifie une signature WebAuthn. Vous n’avez pas besoin de suivre les deux volets dans l’ordre : un administrateur système peut s’arrêter après la partie configuration, un développeur peut sauter directement à la partie Node.js.

## Qu’est-ce qu’un passkey ? FIDO2 et WebAuthn expliqués simplement

Un passkey est un identifiant cryptographique qui remplace le couple identifiant-mot de passe. Techniquement, il repose sur deux standards complémentaires : **WebAuthn**, l’API du navigateur normalisée par le W3C, et **CTAP2**, le protocole défini par la FIDO Alliance pour parler à un authentificateur (empreinte digitale, Face ID, clé de sécurité). Ensemble, ils forment ce que l’industrie appelle FIDO2. La spécification WebAuthn Level 2 est une recommandation W3C publiée depuis avril 2021, et la version Level 3 poursuit son passage en recommandation candidate au W3C, ce qui montre que le standard continue d’évoluer plutôt que de stagner.

Contrairement à un mot de passe, un passkey génère une paire de clés asymétriques lors de l’inscription. La clé privée ne quitte jamais l’appareil ou le gestionnaire de clés de l’utilisateur. Seule la clé publique part vers le serveur. Un site piraté qui perd sa base de données ne récupère donc que des clés publiques, inutilisables pour se faire passer pour un utilisateur. C’est la différence fondamentale avec un hash de mot de passe, qui reste une cible même chiffré.

Deuxième propriété clé : la liaison à l’origine. Chaque passkey est cryptographiquement associé au domaine exact pour lequel il a été créé. Un faux site de phishing qui imite tech-insider.org sur un domaine différent ne pourra jamais déclencher une signature valide, même si la victime clique sur tout ce qu’on lui présente. Le phishing classique par formulaire devient structurellement impossible, pas seulement moins probable.

La FIDO Alliance, le consortium industriel derrière ces standards, regroupe des acteurs comme Google, Microsoft, Apple et Amazon autour d’un objectif commun : remplacer le mot de passe par de la cryptographie à clé publique pilotée depuis le navigateur ou le système d’exploitation. Le W3C, de son côté, standardise la brique WebAuthn qui permet à n’importe quel site web d’appeler ces authentificateurs depuis JavaScript, sans dépendre d’un plugin propriétaire.

Il faut distinguer deux catégories d’authentificateurs. Un **authentificateur de plateforme** vit directement dans l’appareil : la puce sécurisée d’un iPhone, le module TPM d’un PC Windows, le composant Titan M d’un Pixel. Un **authentificateur itinérant**, à l’inverse, se branche ou se connecte en Bluetooth, comme une clé YubiKey. Les deux parlent le même protocole CTAP2 et produisent le même type de credential côté serveur, ce qui simplifie grandement l’implémentation : votre backend n’a pas besoin de savoir lequel des deux l’utilisateur a en face de lui.

Autre notion utile avant de coder : le **credential découvrable** (discoverable credential, autrefois appelé resident key). Un passkey moderne stocke, en plus de la paire de clés, l’identifiant de l’utilisateur directement dans l’authentificateur. Résultat concret : l’utilisateur peut se connecter en cliquant simplement sur “se connecter avec un passkey” sans avoir d’abord tapé son nom d’utilisateur, puisque l’appareil sait déjà quels comptes il possède pour ce site. C’est ce mécanisme qui alimente l’autofill conditionnel que nous activerons à l’étape 13.

## Pourquoi passer aux passkeys en 2026

Le calcul est simple pour une équipe sécurité ou un particulier attentif. Un mot de passe seul ne résiste à aucune des attaques courantes citées dans les rapports d’incidents que couvre régulièrement ce site, du credential stuffing au phishing en passant par la réutilisation entre services. Ajouter un code SMS ou une application TOTP améliore les choses, mais ces méthodes restent vulnérables à des attaques de relais en temps réel où la victime saisit elle-même le code sur un faux site. Le tableau suivant résume les propriétés structurelles de chaque méthode, telles que définies par les spécifications WebAuthn et FIDO2.

| Méthode d’authentification | Résiste au phishing | Résiste au vol de base de données | Action requise à la connexion | 
|---|---|---|---|
| Mot de passe seul | Non | Non | Frappe manuelle | 
| Mot de passe + SMS OTP | Partiellement | Non (le mot de passe reste exposé) | Frappe + code reçu par SMS | 
| Mot de passe + TOTP (Authenticator) | Partiellement | Non (le mot de passe reste exposé) | Frappe + code à 6 chiffres | 
| Clé de sécurité physique (FIDO2) | Oui | Oui | Toucher la clé | 
| Passkey (plateforme ou synchronisable) | Oui | Oui | Biométrie ou code PIN de l’appareil | 

Cette résistance structurelle explique pourquoi les gestionnaires de mots de passe eux-mêmes intègrent désormais le stockage de passkeys plutôt que de s’y opposer. L’ANSSI, l’agence française chargée de la sécurité des systèmes d’information, pousse depuis plusieurs années vers une authentification forte pour les services sensibles, et CERT-FR documente régulièrement des campagnes de phishing qui contournent les mots de passe classiques et les seconds facteurs faibles. Les passkeys s’inscrivent directement dans cette logique de durcissement, sans demander à l’utilisateur final de retenir quoi que ce soit de nouveau.

Il y a aussi un argument économique, souvent sous-estimé côté produit. Chaque réinitialisation de mot de passe génère un ticket de support, un email de vérification, parfois un appel. Multiplié par une base d’utilisateurs de plusieurs dizaines de milliers de comptes, ce coût opérationnel pèse réellement sur une équipe support. Un passkey ne s’oublie pas de la même façon qu’un mot de passe : il vit dans l’appareil et se réutilise automatiquement, ce qui fait mécaniquement chuter le volume de demandes de réinitialisation une fois l’adoption suffisante. C’est un argument qui convainc souvent plus vite une direction produit que le discours purement sécurité.

