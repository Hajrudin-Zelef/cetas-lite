---
id: collect-261001-general-networking/general-networking/vaultwarden-auto-heberger-bitwarden-en-12-etapes-2026-1
title: "Mise à jour complète du système"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/vaultwarden-auto-heberger-bitwarden-en-12-etapes-2026.md
source_anchor: ""
source_lines: [1, 37]
sha256: 775ec07b7e7d6c65dc2fefadaccd3f3345478a003c6906e201df78a73e350521
---

# Mise à jour complète du système

Les identifiants volés ou réutilisés restent l’un des premiers vecteurs d’intrusion informatique, et la CNIL recommande explicitement l’usage d’un gestionnaire de mots de passe pour générer et stocker des secrets uniques. Le problème : confier l’intégralité de votre coffre-fort numérique à un service cloud tiers – souvent hébergé hors d’Europe – soulève des questions de souveraineté et de conformité au RGPD. **Vaultwarden** répond à ce dilemme. C’est une réimplémentation en Rust du serveur Bitwarden (anciennement *bitwarden_rs*), créée par Daniel García, qui pèse entre 10 et 50 Mo de RAM là où le serveur officiel en réclame plusieurs gigaoctets.

Avec plus de 62 900 étoiles sur GitHub et une licence AGPL-3.0, Vaultwarden implémente l’API Bitwarden : les applications officielles (extension navigateur, application de bureau, mobile, CLI) s’y connectent sans modification. Mieux, il débloque gratuitement des fonctions facturées par Bitwarden (TOTP intégré, pièces jointes, organisations, clés matérielles). Ce guide vous montre comment **auto-héberger** un serveur **Vaultwarden** de production en 12 étapes sur Ubuntu 24.04 LTS, avec Docker Compose, un reverse proxy Caddy gérant le HTTPS automatiquement, un jeton d’administration chiffré en Argon2, les inscriptions verrouillées, une protection fail2ban et des sauvegardes chiffrées. Comptez 40 à 60 minutes. Version cible : **Vaultwarden 1.37.2**, publiée le 22 août 2026 selon ReleaseLog.org – la 85<sup>e</sup> version référencée sur le dépôt GitHub du projet – qui exige désormais des clients en version 2026.8.0 ou supérieure, et déjà proposée en téléchargement sur SourceForge sous la forme d’une archive source d’environ 1 Mo, également disponible via Snap Store (canal stable mis à jour le 27 juillet 2026, canal candidate le 23 août 2026). Elle fait suite à la 1.37.1, datée du 3 août 2026 selon l’avis de sécurité Fedora (le 29 juillet 2026 selon ThereIsAnOpenSourceForThat.com), qui corrige la faille **CVE-2026-31812** et a été empaquetée sous la référence 1.37.1-1.fc44 pour Fedora 44 le même jour selon l’avis de sécurité Fedora, elle-même héritière de la 1.37.0 du 24 juillet 2026 (archive source de 739,1 Ko, exigeant des clients en version 2026.7.0 ou supérieure) qui corrige des problèmes d’invitations ainsi qu’un souci OpenSSL sous Alpine. *Mis à jour le 28 août 2026.*

## Pourquoi auto-héberger son gestionnaire de mots de passe en 2026 ?

Un gestionnaire de mots de passe centralise vos secrets les plus sensibles : accès bancaires, messageries, infrastructures professionnelles. Le confier à un service cloud commercial signifie accepter trois compromis. D’abord, la **dépendance** : si le fournisseur ferme, augmente ses tarifs ou subit une fuite, vous n’avez aucun levier. Ensuite, la **souveraineté** : un coffre hébergé sur une infrastructure soumise au Cloud Act américain peut, en théorie, faire l’objet d’une réquisition extraterritoriale, même si les données sont chiffrées de bout en bout. Enfin, le **coût récurrent** des formules familiales ou professionnelles, qui s’accumule année après année.

L’auto-hébergement renverse ces contraintes. Vous gardez la maîtrise complète du chiffrement, de l’emplacement physique des données et du calendrier de mise à jour. Pour un organisme français ou européen soumis au RGPD, héberger le coffre sur un VPS situé en France ou en Allemagne simplifie la cartographie des traitements et la réponse à une analyse d’impact. La CNIL place la sécurité des données par défaut au cœur de la conformité, et la maîtrise de l’hébergement en fait partie. Cet enjeu de souveraineté numérique structure d’ailleurs la stratégie nationale de cybersécurité française.

Vaultwarden ne fait aucun compromis sur le modèle de sécurité de Bitwarden. Le chiffrement reste *zero-knowledge* : votre mot de passe maître ne quitte jamais l’appareil client, et le serveur ne stocke que des données déjà chiffrées. Auto-héberger n’affaiblit donc pas la cryptographie ; cela déplace simplement le point de confiance du fournisseur vers votre propre serveur. La contrepartie est une responsabilité opérationnelle réelle : sauvegardes, mises à jour de sécurité et durcissement deviennent votre affaire. C’est précisément ce que couvrent les étapes suivantes.

## Vaultwarden vs Bitwarden officiel vs solutions cloud

Avant de déployer, il faut comprendre ce qui distingue Vaultwarden du serveur Bitwarden officiel auto-hébergé et des offres cloud. La différence majeure tient à l’empreinte : Vaultwarden est un binaire Rust unique, là où le déploiement officiel classique orchestre plusieurs conteneurs (application, base SQL, identité, icônes). Cette légèreté permet de faire tourner Vaultwarden sur le plus petit VPS européen, voire sur un Raspberry Pi.

| Critère | Vaultwarden 1.36 | Bitwarden auto-hébergé | Bitwarden Cloud | 
|---|---|---|---|
| Langage / moteur | Rust (binaire unique) | C# / .NET + SQL Server | Géré par l’éditeur | 
| Empreinte RAM au repos | ≈ 10–50 Mo | ≈ 2 Go recommandés | Sans objet | 
| Nombre de conteneurs | 1 (+ reverse proxy) | Plusieurs (stack complète) | Sans objet | 
| Base de données | SQLite, MySQL ou PostgreSQL | MS SQL / PostgreSQL | Géré | 
| TOTP, pièces jointes, organisations | Inclus gratuitement | Selon licence | Abonnement payant | 
| Clés matérielles (FIDO2 / YubiKey) | Inclus | Selon licence | Premium | 
| Statut officiel | Non officiel, communautaire | Officiel | Officiel | 
| Coût logiciel | Gratuit (AGPL-3.0) | Gratuit à payant | Abonnement | 

### Ce que Vaultwarden débloque gratuitement

Parce qu’il réimplémente l’API côté serveur, Vaultwarden ne connaît pas la notion d’abonnement « premium ». L’authentificateur TOTP intégré, les pièces jointes chiffrées, les rapports d’hygiène, les organisations avec collections partagées, l’accès d’urgence et la prise en charge des clés matérielles sont tous actifs par défaut ; la création d’organisations elle-même a d’ailleurs été corrigée dans la version 1.35.2 (web vault 2025.12.1) en janvier 2026, après un bug remonté par la communauté. C’est l’argument central de l’auto-hébergement face aux gestionnaires commerciaux que nous avons comparés ailleurs : vous obtenez l’équivalent fonctionnel d’une formule familiale ou entreprise sans coût d’abonnement. La contrepartie, répétons-le, est que la disponibilité et les sauvegardes reposent entièrement sur vous.

Un point de vigilance juridique : Vaultwarden n’est pas affilié à Bitwarden (8bit Solutions LLC). La marque « Bitwarden » reste la propriété de l’éditeur, et le projet le rappelle clairement. Vous utilisez les *clients* officiels Bitwarden contre un *serveur* communautaire – une combinaison parfaitement fonctionnelle, mais sans support de l’éditeur.

## Comment fonctionne le chiffrement zero-knowledge

Comprendre le modèle cryptographique éclaire toutes les décisions de durcissement qui suivent. Lorsque vous saisissez votre mot de passe maître dans un client Bitwarden, celui-ci ne part jamais en clair vers le serveur. Le client dérive localement une clé via la fonction KDF (PBKDF2-SHA256 à 600 000 itérations par défaut, ou Argon2id). De cette clé découlent deux éléments : une *clé de chiffrement* qui protège le contenu du coffre, et un *hachage d’authentification* – seul ce dernier, lui-même re-haché côté serveur, sert à prouver votre identité.

