---
id: collect-261001-general-networking/general-networking/proton-mail-vs-gmail-2026-100m-vs-1-8md-teste-3
title: "Configuration Thunderbird via Proton Bridge"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "Google"]
dates: []
keywords: ["benchmarks", "gemini", "open source"]
source: docs/RAG/collect-261001-general-networking/proton-mail-vs-gmail-2026-100m-vs-1-8md-teste.md
source_anchor: ""
source_lines: [82, 130]
sha256: d42d2e64305a4ea8830e71f4f436cfae2473d40d18b374308f79e0d0409de213
---

# Configuration Thunderbird via Proton Bridge

Cette stratégie est revendiquée par Proton comme un *positionnement intentionnel* : le mode gratuit doit servir d’introduction et non de substitution durable. À titre de comparaison, Tutanota propose 1 Go gratuit (similaire), Mailbox.org démarre à 1 € HT/mois, Infomaniak kSuite démarre à 1,99 €/mois pour 15 Go, et Mailo propose 100 Mo gratuits. Le marché européen des messageries souveraines converge donc sur un même standard : la gratuité totale est la chasse gardée des géants américains amortis par la publicité ou la production de données, modèle économique que les acteurs européens refusent.

Pour un usage professionnel sérieux, la bascule vers **Mail Plus à 4,99 €/mois** rééquilibre le rapport : 15 Go d’espace, 10 alias, 1 domaine personnalisé, prise en charge IMAP/SMTP via Proton Bridge. La parité fonctionnelle avec Gmail gratuit est atteinte à ce niveau, et l’écart de prix avec Workspace Business Starter (6 €/utilisateur HT, soit ~7,20 € TTC) devient minime – voire favorable à Proton lorsqu’on inclut la TVA et les services additionnels.

## Fonctionnalités IA : Gemini intégré vs Proton Scribe chiffré

L’IA générative est le terrain où Gmail prend le plus gros avantage en 2026. **Gemini 1.5 Pro et Gemini 2.0** sont intégrés dans la barre latérale Gmail depuis mars 2025 sur les plans Workspace Business Standard et supérieurs (et désormais Business Starter avec une limite mensuelle de jetons). L’utilisateur peut résumer un fil de 50 messages, brouillonner une réponse, extraire automatiquement des engagements ou rechercher en langage naturel. Selon les benchmarks publiés par le blog Google, l’adoption de Smart Compose économise *jusqu’à 25 % du temps de rédaction* sur les courriels professionnels classiques.

Proton a riposté en juillet 2024 avec **Proton Scribe**, assistant de rédaction lancé d’abord pour les comptes Business et étendu aux particuliers Plus/Unlimited fin 2024. Spécificité majeure : **Scribe peut s’exécuter en local sur le poste utilisateur** (mode *on-device*) ou via une instance cloud chiffrée GDPR-compliant hébergée en Europe. Aucun courriel n’est envoyé en clair vers un serveur d’inférence externe. Cette architecture limite mécaniquement la richesse fonctionnelle (la fenêtre contextuelle ne traite qu’un message à la fois et non l’intégralité du fil), mais elle préserve la propriété zéro-accès qui fonde la promesse Proton.

### Comparatif fonctionnel détaillé

Sur les benchmarks usuels d’assistance e-mail, Gemini gagne nettement en richesse : génération de réponses personnalisées avec ton ajustable (formel, amical, court), extraction de tâches vers Google Tasks ou Calendar, traduction automatique en 100+ langues. Proton Scribe se concentre sur la **rédaction** (réécrire, formaliser, raccourcir) et la **traduction** (français, allemand, anglais, espagnol, italien) sans accès au reste du fil. Pour un utilisateur dont le quotidien est la collaboration multi-fil chez un client final, Gemini reste plus productif. Pour un utilisateur dont la priorité est la confidentialité – avocat, journaliste, médecin, fonctionnaire – Scribe offre la seule architecture où l’IA ne traverse jamais un serveur tiers en clair.

## Filtres anti-spam, anti-phishing et protection avancée

Le filtrage anti-spam est un domaine où la masse fait force. Google traite un volume de courriers entrants tel que ses modèles sont entraînés sur des données quasi exhaustives, expliquant les performances de référence du marché : selon les projections **ElectroIQ** pour 2025, Gmail devrait acheminer environ **376,4 milliards de courriels sur l’année**, soit près de **3,7 millions de messages échangés chaque jour** en moyenne. **Gmail bloque 99,9 % des spams, phishing et malwares** selon les statistiques publiées par Google sur son centre de presse Workspace, et la couche TensorFlow ajoute environ *100 millions de messages malveillants supplémentaires bloqués chaque jour* par rapport à l’ancien moteur. C’est un standard que peu de concurrents égalent.

Proton Mail s’appuie sur une architecture différente : **SpamAssassin renforcé par un modèle interne** entraîné sur les seuls signaux contextuels (sans lecture du contenu chiffré côté serveur, contrainte technique imposée par le zéro-accès). Les performances sont en retrait à la prise en main, mais le système apprend rapidement avec les retours utilisateurs. Notre test de 30 jours sur 4 boîtes (2 personnelles, 2 business) a montré un **taux de blocage initial de ~95 %** sur Proton, montant à **~98 %** au bout de trois semaines d’entraînement actif, contre **~99,7 %** sur Gmail dès le premier jour. L’écart est réel mais s’estompe.

Pour la protection anti-phishing avancée, les deux services proposent des dispositifs complémentaires. Proton intègre **PhishGuard**, qui analyse les en-têtes SPF/DKIM/DMARC et avertit visuellement sur les courriels suspects, plus la fonction **Sentinel** évoquée plus haut pour les profils à haut risque. Google complète Gmail par **Advanced Protection Program** et par les contrôles de sécurité embarqués dans Workspace (alertes Sandboxing, scan des pièces jointes via Chronicle Security). En entreprise, Workspace conserve un avantage net sur les fonctions SIEM, l’audit log centralisé et l’intégration avec les SOC tiers (Splunk, Sentinel, Elastic).

## Applications mobiles et clients de bureau

L’expérience client est un terrain de matchs très resserrés. **L’application Gmail iOS** conserve une note moyenne de 4,7/5 sur l’App Store français en avril 2026 et figure parmi les applications les plus téléchargées toutes catégories confondues. **L’application Proton Mail iOS** obtient une note de 4,7/5 également avec un volume d’avis plus modeste (autour de 60 000 contre plusieurs millions pour Gmail). Sur Android, le constat est similaire : Gmail est préinstallé sur la quasi-totalité des appareils, là où Proton Mail nécessite un téléchargement actif depuis Google Play ou F-Droid.

Le différentiel pour Proton se joue sur l’**open source**. Le code des applications Proton Mail web, iOS et Android est disponible sur github.com/ProtonMail. Cette transparence permet à des chercheurs en sécurité indépendants d’auditer le chiffrement et de soumettre des correctifs ; aucune équivalence n’existe pour Gmail dont le client web et les apps mobiles restent strictement propriétaires.

### Compatibilité IMAP/SMTP et clients tiers

Pour les utilisateurs accrochés à leurs clients de bureau (Thunderbird, Apple Mail, Outlook, Mailspring), **Gmail propose IMAP, SMTP et POP3 nativement** via OAuth 2.0 avec mots de passe d’application. Proton, par construction zéro-accès, ne peut pas exposer un protocole standard sans casser le chiffrement bout-en-bout. La société propose donc **Proton Bridge**, application de pont locale (macOS, Windows, Linux) qui déchiffre côté client et exposait un IMAP/SMTP local. Bridge est inclus à partir du plan Mail Plus à 4,99 €/mois et requiert un compte payant, ce qui constitue un coût d’entrée réel pour les utilisateurs très attachés à Outlook ou Thunderbird.

Voici un exemple de configuration Thunderbird minimaliste après installation de Proton Bridge :

```
# Configuration Thunderbird via Proton Bridge
Serveur IMAP entrant : 127.0.0.1
Port IMAP            : 1143
Sécurité             : STARTTLS
Authentification     : Mot de passe normal
Serveur SMTP sortant : 127.0.0.1
Port SMTP            : 1025
Sécurité             : STARTTLS
Authentification     : Mot de passe normal
# Bridge fournit un identifiant et un mot de passe spécifiques
# différents du mot de passe de votre compte Proton.
```
## Bundles et écosystème : Proton Suite vs Workspace

