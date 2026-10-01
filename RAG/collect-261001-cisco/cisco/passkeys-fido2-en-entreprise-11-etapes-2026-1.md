---
id: collect-261001-cisco/cisco/passkeys-fido2-en-entreprise-11-etapes-2026-1
title: "Activer l'action requise WebAuthn Passwordless"
domain: cisco
role: reference
task: reference
actors: ["Apple", "Google", "Microsoft"]
dates: []
keywords: ["mai"]
source: docs/RAG/collect-261001-cisco/passkeys-fido2-en-entreprise-11-etapes-2026.md
source_anchor: ""
source_lines: [1, 49]
sha256: 3d2a7ef859da06b3b04de565c3ffbca0701506c1d3f01d44635f333419a3e0b5
---

# Activer l'action requise WebAuthn Passwordless

Les mots de passe restent la première porte d’entrée des attaquants en Europe, et 2026 marque un tournant : les entreprises françaises basculent en masse vers les passkeys (clés d’accès FIDO2/WebAuthn). Entre le rapport ENISA qui pousse les organisations de l’UE à structurer leur passage au sans mot de passe, la sortie de la YubiKey 5.8 en juillet 2026 et le rebranding par Okta de son authentificateur FIDO2 en « Passkey » début août, le sujet n’a jamais été aussi concret. Ce tutoriel détaille, étape par étape, comment déployer les passkeys dans votre entreprise sans casser la productivité ni ouvrir de nouvelles failles.

Contrairement aux guides théoriques, cet article part du terrain : inventaire des services à risque, choix entre passkeys synchronisées et liées à l’appareil, intégration avec un fournisseur d’identité (IdP), configuration de clés matérielles YubiKey, tests d’interopérabilité multi-navigateurs, et plan de reprise en cas de perte d’appareil. Vous repartirez avec un projet complet et fonctionnel, prêt à être adapté à votre organisation.

## Pourquoi migrer vers les passkeys en 2026 ?

Les passkeys reposent sur la norme FIDO2/WebAuthn : une paire de clés cryptographiques remplace le mot de passe, la clé privée ne quittant jamais l’appareil de l’utilisateur ou son gestionnaire de clés sécurisé. Le résultat, c’est une authentification résistante au phishing, puisqu’il n’existe plus de secret à voler ou à intercepter sur un faux site.

Selon le rapport *State of Passkeys 2026* de la FIDO Alliance, publié en mai 2026, plus de 5 milliards de passkeys sont désormais en circulation dans le monde ; la notoriété du grand public atteint 90 % (contre 75 % un an plus tôt), et 75 % des internautes ont activé au moins une passkey, même si seuls 49 % l’utilisent réellement au quotidien. Côté entreprises, 68 % des organisations ont déjà déployé, sont en train de déployer, ou lancent un déploiement de passkeys pour l’authentification de leurs employés en 2026. Un an plus tôt, l’Alliance notait déjà que 87 % des entreprises interrogées aux États-Unis et au Royaume-Uni avaient entamé ce virage. Le mouvement est donc bien engagé, et l’Europe n’est pas en reste : l’ENISA (Agence de l’Union européenne pour la cybersécurité) recommande explicitement ce type de déploiement dans son panorama des menaces 2025, publié à l’automne et régulièrement mis à jour.

Sur le plan réglementaire, la norme **NIST SP 800-63-4**, publiée en juillet 2025, est la première à qualifier formellement les passkeys synchronisées au niveau d’assurance AAL2 (Authentication Assurance Level 2). C’est un signal fort : les auditeurs et les régulateurs acceptent désormais cette méthode comme une authentification forte conforme, y compris dans des contextes réglementés comme la finance ou la santé.

Autre évolution notable : selon un rapport publié le 10 août 2026, des chercheurs ont démontré que les passkeys synchronisées stockées dans le cloud du navigateur ou de l’OS pouvaient, dans certains scénarios, être ciblées via les flux de récupération de compte ou la mémoire du navigateur (The Hacker News). Cette découverte ne remet pas en cause l’intérêt des passkeys, mais elle justifie pleinement la démarche « risque par service » détaillée plus bas : les comptes à privilèges méritent des clés matérielles liées à l’appareil, pas uniquement des passkeys synchronisées dans le cloud.

## Prérequis techniques et organisationnels

Avant de lancer le moindre pilote, réunissez les éléments suivants. Un déploiement de passkeys touche l’identité, les postes de travail et les navigateurs : mieux vaut cadrer large dès le départ.

- Un fournisseur d’identité (IdP) compatible WebAuthn : Okta Identity Engine (API 2026, authentificateur renommé « Passkey » depuis le 1er août 2026), Microsoft Entra ID, ou Keycloak en version récente pour les déploiements auto-hébergés.
- Des navigateurs à jour : Chrome, Edge, Firefox et Safari supportent tous WebAuthn Level 2 nativement depuis plusieurs versions ; vérifiez simplement que votre politique de gestion de navigateur d’entreprise (MDM/EMM) n’en bloque pas les API.
- Des clés de sécurité matérielles pour les comptes à haut risque : YubiKey 5.8 (lancée le 24 juillet 2026, avec autorisation matérielle des signatures) ou équivalent compatible FIDO2/CTAP2.
- Un annuaire d’entreprise (Active Directory, Azure AD / Entra ID, ou LDAP) déjà relié à votre IdP.
- Un outil MDM pour gérer les appareils BYOD et professionnels (Intune, Jamf, ou équivalent).
- Du temps : comptez 6 à 12 mois pour un déploiement complet, de l’intégration IdP jusqu’à l’inscription généralisée des utilisateurs.
- Un budget clés matérielles : entre 25 € et 60 € par clé selon le modèle, pour les comptes Tier 1 et Tier 2.

Sur le plan organisationnel, désignez un référent identité (souvent le RSSI ou un architecte IAM) et un référent support utilisateur : les deux premières semaines de pilote génèrent presque toujours des questions sur la récupération de compte.

## Étape 1 : Cartographier les services par niveau de risque

La première erreur d’un déploiement de passkeys, c’est de vouloir tout migrer en même temps. La bonne pratique documentée en 2026 consiste à classer les services en niveaux de risque (Tier 1 à Tier 3) et à cibler d’abord les comptes les plus sensibles.

| Niveau | Exemples de services | Type de passkey recommandé | Priorité de déploiement | 
|---|---|---|---|
| Tier 1 (critique) | Consoles admin cloud, VPN d’entreprise, systèmes financiers | Passkey liée à l’appareil (clé matérielle YubiKey) | Semaines 1 à 6 | 
| Tier 2 (sensible) | Messagerie professionnelle, CRM, outils RH | Passkey synchronisée + clé matérielle de secours | Mois 2 à 4 | 
| Tier 3 (standard) | Intranet, outils collaboratifs internes | Passkey synchronisée (Apple/Google/Microsoft) | Mois 4 à 9 | 
| Tier 4 (BYOD/externe) | Portails prestataires, applications SaaS tierces | Passkey synchronisée avec MFA de secours | Mois 6 à 12 | 

Cette matrice sert aussi de base à votre dossier de conformité RGPD et NIS2 : elle prouve que le choix des mécanismes d’authentification est proportionné au risque, ce que les auditeurs demandent de plus en plus explicitement.

## Étape 2 : Choisir entre passkeys synchronisées et liées à l’appareil

C’est la décision technique la plus structurante du projet. Les passkeys **synchronisées** (synced) sont répliquées via le trousseau cloud d’Apple, le gestionnaire de mots de passe Google, ou Windows Hello via Microsoft Entra : elles offrent un confort d’usage élevé (l’utilisateur retrouve sa passkey sur tous ses appareils) mais dépendent de la sécurité du compte cloud associé. Les passkeys **liées à l’appareil** (device-bound), comme celles stockées sur une YubiKey, ne quittent jamais le composant matériel : elles sont plus contraignantes à utiliser mais offrent la meilleure garantie contre l’exfiltration.

D’après la FIDO Alliance, 47 % des organisations ayant déployé des passkeys en entreprise utilisent aujourd’hui un mix des deux approches, ce qui confirme qu’il n’existe pas de réponse unique. Pour un cabinet d’avocats ou une PME de service, des passkeys synchronisées suffisent largement sur les comptes Tier 3. Pour une banque ou un opérateur d’importance vitale, les comptes à privilèges devraient rester sur des clés liées à l’appareil.

## Étape 3 : Configurer votre fournisseur d’identité (exemple Keycloak)

