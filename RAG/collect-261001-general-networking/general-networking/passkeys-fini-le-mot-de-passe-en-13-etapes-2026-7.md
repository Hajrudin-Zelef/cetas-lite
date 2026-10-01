---
id: collect-261001-general-networking/general-networking/passkeys-fini-le-mot-de-passe-en-13-etapes-2026-7
title: "Installe une autorité de certification locale de confiance"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "Google"]
dates: []
keywords: ["attention", "exploit", "incident"]
source: docs/RAG/collect-261001-general-networking/passkeys-fini-le-mot-de-passe-en-13-etapes-2026.md
source_anchor: ""
source_lines: [436, 474]
sha256: f058cb69aee82dc26743651c5e6a44c0c90cae84184a5d42131eb9ff0c84731b
---

# Installe une autorité de certification locale de confiance

Sur le terrain réglementaire, la directive NIS2 pousse un nombre croissant d’entités françaises vers une authentification forte pour l’accès aux systèmes sensibles, un sujet que nous avons détaillé dans notre couverture de l’actualité cybersécurité du site. L’ANSSI recommande depuis longtemps l’authentification multifacteur pour les comptes à privilèges, et les passkeys remplissent cette exigence tout en supprimant le facteur de risque humain lié à la saisie d’un mot de passe sur un site frauduleux. En cas d’incident malgré tout, CERT-FR reste le point de contact national pour le signalement, joignable par courriel à [email protected] ou par téléphone au 3218, un service gratuit accessible depuis la France. Au niveau européen, l’ENISA documente également les bonnes pratiques d’authentification forte pour l’ensemble des États membres.

Pour une DSI qui prépare un dossier de conformité, gardez une trace écrite de trois éléments : la nature exacte des données collectées lors de l’inscription d’un passkey (avec attestationType à ‘none’, il s’agit essentiellement d’une clé publique et d’un identifiant de credential, rien de plus), la durée de conservation de ces données, et la procédure de suppression lorsqu’un utilisateur supprime son compte ou révoque un passkey. Cette documentation, courte à rédiger puisque le traitement est volontairement minimal, facilite grandement les échanges avec un délégué à la protection des données ou, le cas échéant, avec la CNIL.

## Foire aux questions

**Qu’est-ce qu’un passkey exactement ?**

Un passkey est une paire de clés cryptographiques qui remplace le couple identifiant-mot de passe. La clé privée reste sur votre appareil ou votre gestionnaire de mots de passe, la clé publique seule est envoyée au site auquel vous vous connectez.

**Les passkeys remplacent-elles complètement les mots de passe ?**

Pas du jour au lendemain. La plupart des services proposent le passkey en option à côté du mot de passe existant, puis encouragent progressivement la bascule complète une fois l’adoption suffisante.

**Que se passe-t-il si je perds mon téléphone contenant mes passkeys ?**

Si vos passkeys sont synchronisés via iCloud, Google ou un gestionnaire tiers, vous les récupérez en vous reconnectant à ce même compte cloud sur un nouvel appareil, après avoir vérifié votre identité par les moyens habituels de récupération de ce compte cloud. S’ils étaient liés uniquement à l’appareil perdu, vous devez utiliser votre méthode de récupération de secours définie au moment de l’inscription, ce qui souligne l’importance de ne jamais sauter cette étape lors de la configuration initiale.

**Les passkeys sont-elles vraiment plus sûres que les mots de passe ?**

Oui, structurellement. Elles suppriment la réutilisation de mot de passe, la vulnérabilité au phishing par formulaire et le risque lié au vol d’une base de données de hashs, trois causes majeures des compromissions de comptes documentées dans les rapports d’incidents. Aucune sécurité n’est absolue, mais ces trois vecteurs d’attaque comptent parmi les plus exploités au quotidien, ce qui rend leur suppression particulièrement significative.

**Puis-je utiliser un passkey créé sur iPhone pour me connecter depuis un PC Windows ?**

Oui, via un mécanisme appelé hybrid transport : le PC affiche un QR code, vous le scannez avec votre iPhone, qui confirme la connexion via Bluetooth de proximité sans jamais transmettre la clé privée elle-même.

**Les passkeys sont-elles compatibles avec le RGPD ?**

Oui. Elles traitent en réalité moins de données personnelles qu’un système de mot de passe classique, puisque aucun secret partagé n’est stocké côté serveur.

**Quelle est la différence entre un passkey et une clé de sécurité physique type YubiKey ?**

Les deux reposent sur le même standard FIDO2/WebAuthn. Une clé physique matérialise le secret sur un objet dédié que vous transportez, alors qu’un passkey plateforme vit dans votre téléphone ou votre ordinateur, éventuellement synchronisé dans le cloud.

**WebAuthn fonctionne-t-il sans connexion internet ?**

La cérémonie cryptographique locale (déverrouillage biométrique, signature) fonctionne hors ligne, mais l’échange avec le serveur pour vérifier cette signature nécessite bien sûr une connexion active au service concerné.

Vous disposez maintenant des deux moitiés du sujet : la partie utilisateur, pour verrouiller vos comptes personnels en quelques minutes, et la partie développeur, avec un projet Node.js fonctionnel que vous pouvez étendre ou brancher sur votre application existante. Les passkeys ne demandent ni budget matériel ni changement d’infrastructure lourd, seulement une bibliothèque bien choisie et une attention particulière aux quelques pièges décrits plus haut. Dans un contexte où les fuites d’identifiants continuent de s’accumuler, c’est l’un des rares changements techniques qui améliore simultanément la sécurité et le confort d’usage, plutôt que d’imposer l’un au détriment de l’autre.
