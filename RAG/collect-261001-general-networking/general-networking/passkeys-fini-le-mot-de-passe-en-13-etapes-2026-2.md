---
id: collect-261001-general-networking/general-networking/passkeys-fini-le-mot-de-passe-en-13-etapes-2026-2
title: "Installe une autorité de certification locale de confiance"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Apple", "Google", "Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/passkeys-fini-le-mot-de-passe-en-13-etapes-2026.md
source_anchor: ""
source_lines: [35, 86]
sha256: 441e1981ea215f7a1794448f26ac0ee90bec339e229404fe2367938a923c5103
---

# Installe une autorité de certification locale de confiance

Dernier point, moins technique mais tout aussi réel : l’expérience utilisateur. Taper un mot de passe complexe sur un clavier de smartphone reste pénible. Un passkey remplace cette friction par un geste déjà familier, celui qui déverrouille l’appareil plusieurs dizaines de fois par jour. Les équipes produit qui ont testé les deux parcours constatent généralement un taux d’abandon plus faible sur l’écran de connexion lorsque le passkey est proposé en premier, pour la simple raison qu’il retire une étape de friction plutôt que d’en ajouter une.

## Prérequis : versions, outils et compatibilité

Avant de commencer, séparez bien les deux volets de ce tutoriel. Le premier volet, rapide, consiste à activer des passkeys sur des comptes que vous possédez déjà. Il ne demande qu’un smartphone récent et un navigateur à jour. Le second volet, plus technique, construit une authentification par passkey dans votre propre application. Voici ce qu’il vous faut pour le suivre dans de bonnes conditions.

Sur l’exigence HTTPS, une précision évite bien des heures perdues. WebAuthn considère qu’une origine est “sécurisée” soit parce qu’elle utilise HTTPS avec un certificat valide, soit parce qu’il s’agit strictement de `http://localhost`. Un serveur de développement exposé sur l’adresse IP de votre machine locale, du type `http://192.168.1.20:3000`, ne remplit aucune des deux conditions et fera échouer silencieusement l’appel à l’API, souvent sans message d’erreur explicite dans la console. C’est la source de confusion numéro un chez les développeurs qui découvrent WebAuthn, d’où l’intérêt de configurer mkcert dès le départ plutôt que d’y venir seulement après un premier échec.

- **Node.js 20 LTS** ou une version plus récente, avec npm installé
- Un éditeur de code (VS Code ou équivalent)
- La bibliothèque **@simplewebauthn/server** pour le backend, activement maintenue sur GitHub par son auteur Matthew Miller (MasterKale)
- La bibliothèque **@simplewebauthn/browser** pour le frontend
- Un certificat HTTPS local (via **mkcert** ) ou un tunnel type ngrok, car WebAuthn exige une origine sécurisée en dehors de localhost strict
- Un navigateur récent : Chrome, Edge, Firefox ou Safari supportent tous WebAuthn au moment de la rédaction
- Un appareil doté d’un capteur biométrique ou d’un code PIN système (Windows Hello, Touch ID, Face ID, ou verrouillage d’écran Android)

Le tableau ci-dessous récapitule la prise en charge des passkeys selon la plateforme, telle que documentée par les éditeurs eux-mêmes et par le référentiel MDN de l’API Web Authentication.

| Plateforme | Support des passkeys | Stockage / synchronisation | 
|---|---|---|
| Google Chrome (bureau et Android) | Oui | Gestionnaire de mots de passe Google | 
| Apple Safari (macOS et iOS) | Oui | Trousseau iCloud | 
| Microsoft Edge / Windows Hello | Oui | Compte Microsoft ou stockage local lié à l’appareil | 
| Mozilla Firefox | Oui (recours à un gestionnaire tiers ou à l’OS pour la synchronisation) | Dépend du gestionnaire configuré | 
| Gestionnaires tiers (Bitwarden, 1Password) | Oui | Coffre chiffré multiplateforme | 

## Qui a déjà déployé les passkeys en 2026

Les passkeys ne sont plus un concept de laboratoire. Google, Microsoft, Apple, GitHub, Amazon et PayPal proposent tous aujourd’hui la connexion par passkey sur leurs comptes grand public, chacun avec une approche légèrement différente selon son propre écosystème. Ce large soutien industriel compte autant que la technique elle-même : un standard n’a de valeur pratique que si les navigateurs, les systèmes d’exploitation et les grands services web l’implémentent tous de la même façon.

Chaque acteur illustre un cas d’usage différent. Google, avec l’échelle de son parc d’utilisateurs, a fait du passkey l’option mise en avant lors de la création ou de la sécurisation d’un compte, pour réduire mécaniquement le volume d’attaques par bourrage d’identifiants sur son infrastructure. GitHub cible un public de développeurs habitués à manier des clés SSH et des jetons d’accès, pour qui un passkey remplace naturellement un mot de passe de compte devenu la porte d’entrée vers du code source et des secrets de déploiement. PayPal, dans un contexte où chaque connexion touche potentiellement à de l’argent, met en avant la résistance au phishing comme argument de confiance directement dans son parcours de connexion. Microsoft, de son côté, capitalise sur Windows Hello déjà présent sur la majorité des postes professionnels pour convertir un geste de déverrouillage de session en méthode de connexion web à part entière.

Cette diversité de cas d’usage est une bonne nouvelle pour quiconque hésite encore à se lancer. Il n’existe pas un seul “bon” moment ou un seul “bon” secteur pour adopter les passkeys : le standard sert aussi bien un grand public non technique qu’une audience de développeurs, aussi bien un service financier qu’un outil de productivité interne.

## Étapes 1 à 4 : activer les passkeys sur vos comptes Google, Microsoft, Apple et GitHub

Avant de coder quoi que ce soit, prenez cinq minutes pour sécuriser vos comptes personnels. C’est la partie la plus rentable de ce tutoriel pour un lecteur non développeur.

### Étape 1 et 2 : Google et Microsoft

Sur un compte Google, ouvrez la gestion du compte, puis la section consacrée à la sécurité et à la connexion. Vous y trouverez une option pour créer un passkey, qui déclenche immédiatement une demande de vérification biométrique ou de code PIN via votre appareil. Une fois validée, le passkey apparaît dans la liste de vos méthodes de connexion, aux côtés du mot de passe que vous pouvez conserver comme méthode de secours ou supprimer plus tard.

Sur un compte Microsoft, la démarche est similaire depuis les paramètres de sécurité avancés. Windows Hello, s’il est déjà configuré sur votre PC, sert directement d’authentificateur de plateforme. L’intérêt pour une entreprise qui utilise déjà l’écosystème Microsoft est immédiat, puisque le passkey s’appuie sur un mécanisme que les utilisateurs connaissent déjà pour déverrouiller leur session.

Un détail pratique mérite d’être signalé pour ces deux comptes : une fois le passkey créé, ne supprimez pas immédiatement le mot de passe associé. Laissez-le comme filet de sécurité pendant quelques semaines, le temps de confirmer que la connexion par passkey fonctionne sans accroc sur tous vos appareils habituels, y compris après une mise à jour du système d’exploitation ou un changement de navigateur par défaut.

### Étape 3 et 4 : Apple et GitHub

Sur un iPhone ou un Mac, un passkey Apple se crée depuis les réglages du compte iCloud ou directement au moment de la connexion à un site compatible, via Touch ID ou Face ID. Le passkey se synchronise ensuite sur tous les appareils reliés au même identifiant Apple grâce au trousseau iCloud, chiffré de bout en bout.

GitHub, très utilisé par le lectorat technique de ce site, propose l’ajout d’un passkey depuis les paramètres de sécurité du compte, dans la même section que les clés SSH et les jetons d’accès personnels. Pour une équipe de développement, remplacer le mot de passe GitHub par un passkey ferme une porte d’entrée classique vers le code source et les secrets de déploiement.

