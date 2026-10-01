---
id: collect-261001-general-networking/general-networking/proton-mail-vs-gmail-2026-100m-vs-1-8md-teste-1
title: "Configuration Thunderbird via Proton Bridge"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "Google", "Microsoft"]
dates: []
keywords: ["arr", "gemini", "mai", "open source"]
source: docs/RAG/collect-261001-general-networking/proton-mail-vs-gmail-2026-100m-vs-1-8md-teste.md
source_anchor: ""
source_lines: [1, 39]
sha256: b2305fadf203a3ad3c5b040ed59cded6a284b6ea44fd8d0461f383c5b4b66740
---

# Configuration Thunderbird via Proton Bridge

Le 28 avril 2026, le marché européen du courrier électronique entre dans une phase charnière. **Proton Mail**, basé à Genève, revendique désormais plus de **100 millions de comptes Proton** – un seuil confirmé en mai 2026 par l’analyse **SellCell**, qui classe le service parmi les principaux fournisseurs de messagerie axés sur la confidentialité – contre **2,5 milliards d’utilisateurs actifs Gmail** recensés en 2025 selon **XtendedView**. L’écart est spectaculaire – 25× en faveur de Google – mais l’arithmétique masque un transfert de fond : depuis l’arrêt Schrems II et l’application stricte du RGPD, des milliers d’administrations, de cabinets d’avocats et d’ETI françaises remettent à plat leur stack e-mail. Ce comparatif **Proton Mail vs Gmail 2026** tranche la question avec des données chiffrées, trois sources de tests indépendants et un guide de migration pas à pas.

Nous avons confronté les deux services sur dix-huit critères : tarifs en euros TTC, taille de la boîte de stockage, chiffrement bout-en-bout, conformité RGPD, exposition au CLOUD Act américain, intégration des assistants IA (Gemini, Proton Scribe), filtres anti-spam, applications mobiles, bundles (VPN, Drive, Pass) et qualité du service support. Verdict en fin d’article, mais l’enseignement principal est déjà clair : **Proton Mail** n’est plus une niche militante, c’est une alternative crédible aux suites américaines pour qui place la souveraineté des données et le chiffrement zéro-accès au-dessus de l’écosystème de productivité.

## État du marché mondial du courrier électronique en 2026

Le marché du courrier électronique reste l’un des plus concentrés du numérique mondial. Selon les données de Litmus et de StatCounter pour 2025, **Gmail capte 39,39 % des ouvertures globales** et conserve une avance considérable sur Apple Mail (58,3 % d’ouvertures sur mobile, mais combine iCloud et autres services), Outlook et Yahoo. **Proton Mail**, à l’inverse, se situe encore autour de **0,28 % de parts d’ouvertures dans le monde**, mais sa dynamique se lit dans le trafic direct : en mai 2025, ProtonMail.com a enregistré **550 300 visites**, en hausse de **23,95 %** sur un mois selon **ElectroIQ**. La société suisse, dirigée par Andy Yen, a franchi le cap des **100 millions de comptes** fin 2024 – chiffre repris en juin 2025 par une compilation **Wikipedia relayée par ElectroIQ** – et capitalise depuis sur le contexte réglementaire post-Schrems II.

L’environnement légal européen redéfinit profondément les arbitrages. La Cour de justice de l’Union européenne a invalidé en 2020 le Privacy Shield, puis le nouveau cadre Data Privacy Framework signé en 2023 reste contesté par Maximilian Schrems et noyb. La **CNIL française** a publié plusieurs mises en demeure visant Google Workspace dans les écoles et établissements publics entre 2022 et 2025, exigeant des garanties supplémentaires contre les transferts de données vers les États-Unis. Les autorités allemandes (BfDI) et italiennes (Garante) ont publié des positions similaires sur Microsoft 365 et Workspace, créant un climat propice aux alternatives européennes.

Côté Mountain View, **Google a intégré Gemini dans Gmail en mars 2025**, accélérant la course à l’IA générative dans la boîte de réception : Smart Compose, Smart Reply, résumé automatique de fil, recherche sémantique. La société a signé un partenariat à 1 milliard de dollars avec Apple en 2026 pour fournir le moteur de Siri, confirmant son rôle d’épine dorsale IA grand public. Mais cette même puissance algorithmique alimente le débat européen : un courriel chiffré côté serveur reste un courriel *lisible* par le fournisseur, donc indexable et utilisable pour entraîner des modèles. C’est la ligne de fracture sur laquelle Proton a bâti son discours commercial.

Sur le plan financier, **Alphabet** a publié 350 milliards de dollars de revenus consolidés sur l’exercice 2024 dont une part significative liée à Workspace, portée par une base Gmail désormais estimée à **2,5 milliards d’utilisateurs actifs** en 2025 selon **XtendedView** – un socle qui pèse directement sur les arbitrages des directions achats en entreprise. Proton AG, restée sous statut de société à mission depuis 2024 (le fondateur a transféré la majorité du capital à la *Proton Foundation*, structure non lucrative de droit suisse), ne publie pas son chiffre d’affaires précis, mais évoque dans ses derniers communiqués une **croissance soutenue à deux chiffres** et plus de 500 employés à Genève, Lausanne, Zurich, Skopje, Prague, Taipei et Vilnius. Le rapport de force économique reste donc colossal en faveur de Google, mais le rapport de force réglementaire en Europe penche dans l’autre sens.

## Tableau comparatif technique : 14 spécifications clés

Avant de creuser chaque critère, voici la photographie technique synthétisée à partir des pages de spécifications officielles de Proton Mail et de Google Workspace, enrichie par les tests de terrain réalisés en avril 2026.

| Spécification | Proton Mail | Gmail / Workspace | 
|---|---|---|
| Date de lancement | 16 mai 2014 (Genève, CERN) | 1er avril 2004 (Mountain View) | 
| Siège social | Genève, Suisse (canton de Genève) | Mountain View, Californie, USA | 
| Utilisateurs (2024-2026) | 100 M comptes Proton | 2,5 Md utilisateurs Gmail | 
| Stockage gratuit | 1 Go (limite 150 messages/jour) | 15 Go partagés Drive/Photos | 
| Stockage Plus / Standard | 15 Go (Mail Plus, 4,99 €/mois) | 30 Go par utilisateur (Business Starter) | 
| Stockage premium | 500 Go (Unlimited, 12,99 €/mois) | 2 To (Business Standard, 12 €/mois) | 
| Chiffrement par défaut | End-to-end zéro-accès (OpenPGP, AES-256) | TLS en transit, S/MIME en option | 
| Localisation des données | Suisse (Genève + Plan-les-Ouates) | USA + UE + monde (Workspace résidence UE optionnelle) | 
| Conformité RGPD | Native, droit suisse FADP renforcé | Conforme, DPF Privacy Framework, CLOUD Act applicable | 
| Code source | Open source (apps web/iOS/Android, Bridge) sur GitHub | Propriétaire | 
| Domaines personnalisés | 1 dans Plus, 3 dans Unlimited, 10 dans Business | 1 inclus dans Business Starter | 
| Assistant IA | Proton Scribe (chiffré côté client) | Gemini intégré (cloud Google) | 
| Bundles inclus | VPN, Drive, Pass, Calendar, Wallet | Drive, Calendar, Meet, Docs, Sheets, Slides | 
| Authentification | 2FA TOTP + clés FIDO2/WebAuthn + Sentinel | 2FA TOTP + clés Titan + Advanced Protection | 

La lecture du tableau impose une nuance : Gmail offre **15× plus de stockage gratuit** et un écosystème de productivité (Docs, Sheets, Meet) que Proton ne cherche pas à reproduire fonctionnalité pour fonctionnalité. Inversement, Proton délivre par défaut un chiffrement **zéro-accès** qui empêche techniquement la société elle-même de lire les courriels stockés – propriété qu’aucune offre Google, Microsoft ou Apple ne propose nativement. Le bon angle d’attaque dépend donc moins de la capacité brute que des priorités : confidentialité et souveraineté d’un côté, productivité collaborative et IA générative de l’autre.

## Chiffrement et sécurité : zéro-accès vs TLS

