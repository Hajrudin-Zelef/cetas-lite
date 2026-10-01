---
id: collect-261001-general-networking/general-networking/double-authentification-mfa-13-etapes-90-min-2026-5
title: "deploy-mfa-hardening.ps1"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Google", "Microsoft"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-general-networking/double-authentification-mfa-13-etapes-90-min-2026.md
source_anchor: ""
source_lines: [285, 331]
sha256: 1eb145dae6ec6c68760139ff3e1bcefdd9ce80cf2ce0d3e2f8de520471e05548
---

# deploy-mfa-hardening.ps1

Pour limiter les frictions, déployez toujours par vagues plutôt qu’en une seule bascule générale. Un premier groupe pilote, composé de l’équipe IT elle-même et de quelques utilisateurs volontaires, permet de repérer les applications legacy qui cassent avec l’accès conditionnel avant que le reste de l’entreprise ne s’en aperçoive. Ajoutez ensuite les comptes à privilèges, puis élargissez département par département en gardant environ une semaine entre chaque vague. Ce rythme laisse le temps au support de traiter les tickets sans être submergé, et donne une fenêtre pour ajuster les exclusions de la politique si un outil métier critique se révèle incompatible avec l’authentification héritée bloquée à l’étape 10.

## Au-delà de Microsoft 365 : Google Workspace, GitHub et AWS

La logique décrite dans ce tutoriel s’applique presque telle quelle à d’autres plateformes. Sur Google Workspace, la validation en deux étapes peut être rendue obligatoire au niveau de l’organisation depuis la console d’administration, avec un support natif des clés de sécurité FIDO2 pour les comptes à privilèges. Sur GitHub, l’authentification à deux facteurs est déjà exigée pour les contributeurs de code depuis plusieurs années, et la plateforme prend désormais en charge les passkeys pour une connexion sans mot de passe.

Du côté d’AWS, le compte racine (root) devrait systématiquement être protégé par un dispositif MFA matériel dédié, distinct de celui utilisé pour les comptes IAM quotidiens, afin d’isoler l’accès le plus critique du reste de l’organisation. Bonne nouvelle pour les équipes qui suivent ce guide : une même clé de sécurité FIDO2 physique peut généralement être enregistrée sur plusieurs de ces plateformes à la fois, puisque WebAuthn est un standard ouvert et non une technologie propriétaire à Microsoft.

La double authentification fonctionne aussi mieux quand elle s’appuie sur des mots de passe uniques et suffisamment longs pour chaque service, ce qui suppose en pratique d’utiliser un gestionnaire de mots de passe plutôt que de réutiliser la même combinaison partout. Notre comparatif des gestionnaires de mots de passe et notre tutoriel pour auto-héberger Vaultwarden détaillent les deux approches, commerciale et auto-hébergée, pour combler ce premier facteur avant d’ajouter le second.

## Foire aux questions

**Qu’est-ce que la double authentification (MFA) ?**

C’est une méthode de connexion qui exige au moins deux preuves d’identité différentes, par exemple un mot de passe et un code généré par une application, avant d’accorder l’accès à un compte. L’objectif est qu’un attaquant qui vole un seul facteur, le mot de passe le plus souvent, ne puisse pas se connecter seul.

**Le MFA par SMS protège-t-il encore contre le phishing en 2026 ?**

Partiellement. Il reste largement préférable à l’absence de MFA, mais il ne résiste pas aux kits de phishing AiTM ni au SIM swapping. Pour les comptes à privilèges, une méthode FIDO2 ou Windows Hello for Business est recommandée à la place.

**Quelle est la différence entre MFA et authentification sans mot de passe ?**

Le MFA classique conserve le mot de passe comme premier facteur et y ajoute une vérification supplémentaire. L’authentification sans mot de passe, via FIDO2 ou passkey, supprime entièrement le mot de passe au profit d’une preuve cryptographique liée à l’appareil.

**Combien de temps faut-il pour déployer le MFA dans une PME ?**

Le socle technique décrit dans ce tutoriel prend environ 90 minutes à mettre en place. La phase d’inscription des utilisateurs et de bascule progressive, elle, s’étale généralement sur deux à quatre semaines pour éviter un pic de tickets au support.

**Le MFA est-il obligatoire avec Microsoft 365 ?**

Microsoft impose déjà le MFA pour certaines actions administratives sensibles et pousse activement les nouveaux tenants vers l’activation des valeurs de sécurité par défaut. L’application complète à l’ensemble des comptes dépend toutefois encore de la configuration choisie par chaque organisation, d’où l’intérêt de suivre les étapes de ce tutoriel.

**Que faire en cas de perte d’une clé de sécurité FIDO2 ?**

Enregistrez toujours une seconde clé de secours au moment de l’inscription initiale. En cas de perte des deux, un Temporary Access Pass généré par un administrateur permet de réinscrire une nouvelle méthode sans jamais transmettre de mot de passe par un canal non sécurisé.

**Une clé FIDO2 fonctionne-t-elle avec plusieurs comptes à la fois ?**

Oui. Une même clé physique peut stocker des identifiants pour Microsoft, Google, GitHub et de nombreux autres services, puisque FIDO2 et WebAuthn sont des standards ouverts partagés par l’ensemble de l’industrie.

**Quelle est la meilleure méthode MFA contre le phishing en 2026 ?**

Les méthodes classées AAL3 par le NIST, à savoir les clés de sécurité FIDO2, les passkeys liés à l’appareil et Windows Hello for Business, offrent aujourd’hui le meilleur niveau de résistance documenté face aux techniques de contournement actuelles.

### Pour aller plus loin

La double authentification ne demande ni gros budget ni compétences exotiques : les 13 étapes ci-dessus tiennent en une matinée pour un tenant de taille moyenne. Ce qui prend réellement du temps, c’est la conduite du changement auprès des utilisateurs. Commencez par les comptes à privilèges, mesurez l’impact avec le rapport d’inscription avant chaque bascule en mode obligatoire, et gardez toujours un compte break-glass hors de portée de vos propres politiques.
