---
id: collect-261001-general-networking/general-networking/proton-mail-vs-gmail-2026-100m-vs-1-8md-teste-5
title: "Configuration Thunderbird via Proton Bridge"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "Google", "Stripe", "United States"]
dates: []
keywords: ["arr", "gemini", "mai", "open source"]
source: docs/RAG/collect-261001-general-networking/proton-mail-vs-gmail-2026-100m-vs-1-8md-teste.md
source_anchor: ""
source_lines: [200, 289]
sha256: ef76bdc4af0089535599cd25bf6d7f60af4449df8a74757e7401868e9d80f217
---

# Configuration Thunderbird via Proton Bridge

Dans Gmail, configurez un **transfert automatique** vers votre nouvelle adresse Proton (Paramètres → Transfert et POP/IMAP). Ainsi, tout courriel reçu sur Gmail pendant la période de transition arrive aussi sur Proton. Prévoyez 30 à 60 jours de double routage avant désactivation, le temps que vos contacts mettent à jour vos coordonnées.

### Étape 5 : ajuster les services tiers

Mettez à jour votre adresse e-mail sur les services critiques : banque, impôts, sécurité sociale, factures, abonnements pro (LinkedIn, GitHub, Stripe). Activez la **vérification anti-phishing** Proton qui flagge automatiquement les tentatives d’usurpation. Désactivez votre compte Gmail seulement après 3 à 6 mois sans nouveau message d’importance reçu sur l’ancienne adresse.

## Avantages et inconvénients : récapitulatif

Synthèse des forces et faiblesses observées sur les deux plateformes après notre test de 60 jours en avril-mai 2026.

### Proton Mail : avantages

- **Chiffrement bout-en-bout zéro-accès** par défaut, audité par Cure53 et SEC Consult
- **Juridiction suisse** hors CLOUD Act et hors entraide judiciaire automatique avec les États-Unis
- Code source **open source** intégral pour les applications client
- Bundle **VPN + Drive + Pass + Wallet** à un prix imbattable
- Rapports de transparence trimestriels publiés
- Mode haute sécurité **Sentinel** disponible pour profils à risque
- Conformité **RGPD native** , FADP suisse renforcée depuis 2023
- Application mobile et web sobre, performante

### Proton Mail : inconvénients

- **1 Go gratuit** seulement, contre 15 Go chez Gmail
- IMAP/SMTP indirect via **Proton Bridge payant** (Mail Plus minimum)
- Filtre anti-spam inférieur à Gmail à la prise en main
- Recherche full-text sur les courriels chiffrés **plus lente**
- Pas d’édition collaborative documentaire intégrée
- Écosystème de tiers et extensions **plus restreint**
- Assistant IA Scribe moins riche fonctionnellement que Gemini

### Gmail / Google Workspace : avantages

- **15 Go gratuits** , écosystème grand public ultra-mature
- Filtre anti-spam de référence (99,9 % de blocage)
- Intégration **Gemini IA** profonde, gain de productivité réel
- Suite collaborative **Docs/Sheets/Slides/Meet** imbattable
- Recherche full-text instantanée sur des téraoctets
- Compatibilité **IMAP/SMTP/POP3 native**
- Marketplace de milliers d’extensions et intégrations
- Support et SLA Workspace Enterprise très complets

### Gmail / Google Workspace : inconvénients

- **Société américaine** , exposée au CLOUD Act, FISA, NSL
- Pas de **chiffrement bout-en-bout par défaut** (CSE réservé Enterprise)
- Code source **propriétaire** , pas d’audit externe possible
- Politique RGPD fragilisée par les arrêts **Schrems I et II**
- Ciblage publicitaire historique (suspendu depuis 2017 sur Gmail mais l’archi reste)
- Migration sortante difficile (verrou écosystème)
- VPN ou gestionnaire de mots de passe non inclus

## Verdict 2026 : qui choisir selon votre profil

Au terme de 60 jours d’usage croisé sur quatre boîtes professionnelles et deux personnelles, le verdict de notre rédaction est nuancé. **Proton Mail s’impose pour la souveraineté, la confidentialité et la sécurité par défaut** ; **Gmail conserve l’avantage en productivité collaborative et IA générative**. Aucun des deux ne ridiculise l’autre, mais chacun excelle dans des dimensions distinctes.

Si vous êtes une **profession réglementée** (avocat, médecin, notaire, journaliste), un **cadre dirigeant** manipulant des informations sensibles, un **activiste** ou simplement un Européen attaché à la maîtrise de ses données : **Proton Mail Unlimited à 12,99 €/mois** est notre recommandation principale. Le bundle VPN/Drive/Pass justifie largement le passage au plan payant et neutralise l’argument du « c’est plus cher ».

Si vous gérez une **scale-up B2B** avec 50+ collaborateurs déjà ancrée dans Google Cloud, BigQuery ou Looker, ou si vous êtes un **particulier** qui valorise avant tout l’écosystème Android/Chrome/YouTube/Photos : **Google Workspace Business Standard ou Google One Family** reste optimal. La friction de migration et le coût d’opportunité dépassent le bénéfice de souveraineté pour ce profil.

Pour un **établissement public français** ou européen, la directive politique post-Schrems II et les positions de la CNIL orientent fermement vers les alternatives souveraines : Proton Mail Business associé à La Suite numérique constitue un duo cohérent à 7,99 €/utilisateur/mois TTC, conforme à la doctrine *cloud au centre* du SGPSN.

## FAQ : Proton Mail vs Gmail en 2026

### Proton Mail est-il vraiment gratuit en 2026 ?

Oui, Proton Mail Free reste gratuit avec 1 Go d’espace, 150 messages par jour, un destinataire par adresse, l’application iOS et Android complètes et le chiffrement bout-en-bout. Aucun engagement, aucune carte bancaire requise pour s’inscrire. Toutefois, l’usage professionnel intensif (volumes, alias multiples, domaine personnalisé, Bridge IMAP/SMTP) requiert le plan Mail Plus à 4,99 €/mois TTC ou supérieur.

### Google peut-il lire mes courriels Gmail ?

Techniquement, oui. Google détient les clés de chiffrement en transit (TLS) et au repos (AES-128 GCM) et peut donc accéder au contenu pour appliquer ses filtres anti-spam, l’indexation, les fonctions Gemini et les fonctions Smart Compose. Le ciblage publicitaire basé sur le contenu Gmail a officiellement cessé en juin 2017, mais l’architecture technique reste la même. Pour une lecture impossible par le fournisseur, il faut activer S/MIME hébergé côté client (CSE) sur Workspace Enterprise, fonctionnalité complexe à déployer.

### Le CLOUD Act américain s’applique-t-il à Proton Mail ?

Non. Proton AG est une société de droit suisse, non soumise à la juridiction américaine. Le CLOUD Act ne s’applique qu’aux fournisseurs sous juridiction US (ou ayant une présence opérationnelle telle qu’un siège ou des actifs significatifs aux États-Unis). Proton est uniquement soumis aux requêtes émises par les autorités suisses dans le cadre de l’entraide judiciaire internationale, supervisées par le Ministère public de la Confédération.

### Puis-je migrer 10 ans de courriels Gmail vers Proton Mail ?

Oui, l’outil **Easy Switch** de Proton permet d’importer l’intégralité d’un compte Gmail (boîte de réception, libellés, contacts, calendriers) en une opération unique. Pour 50 Go de courriels, prévoyez 12 à 24 heures d’import en arrière-plan. Le plan Mail Plus (15 Go) ou Unlimited (500 Go) est nécessaire selon le volume. Easy Switch préserve les libellés et la date d’origine.

### Proton Mail offre-t-il IMAP et SMTP pour Outlook ou Thunderbird ?

Indirectement, via **Proton Bridge**. Bridge est une application légère qui s’installe sur macOS, Windows ou Linux, déchiffre les messages localement et expose un IMAP/SMTP local sur 127.0.0.1. Tout client compatible IMAP (Thunderbird, Outlook, Apple Mail, Mailspring, Spark) peut alors être configuré normalement. Bridge requiert un abonnement payant (Mail Plus minimum à 4,99 €/mois).

### Quel service est le plus performant contre le spam ?

Gmail conserve un avantage à la prise en main grâce à son volume de données d’entraînement et son moteur TensorFlow. Le taux de blocage atteint 99,9 % selon les communications officielles Google. Proton Mail démarre autour de 95 % et progresse à ~98 % après quelques semaines d’usage actif et de retours utilisateurs. La différence est faible pour un usage normal mais peut être notable pendant la période d’apprentissage.

### Proton Mail Business est-il adapté à une PME française ?

