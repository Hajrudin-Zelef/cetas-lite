---
id: collect-261001-general-networking/general-networking/proton-mail-vs-gmail-2026-100m-vs-1-8md-teste-2
title: "Configuration Thunderbird via Proton Bridge"
domain: general-networking
role: reference
task: reference
actors: ["Google", "Microsoft", "Mistral"]
dates: []
keywords: ["dpo", "gemini", "mistral", "open source"]
source: docs/RAG/collect-261001-general-networking/proton-mail-vs-gmail-2026-100m-vs-1-8md-teste.md
source_anchor: ""
source_lines: [40, 81]
sha256: cf0204d269c6ed33c6776d13291449f05d6e9c8f9673e16ab7dd417dbeb7f1cb
---

# Configuration Thunderbird via Proton Bridge

Le chiffrement est le terrain où Proton Mail a construit son identité. Tous les courriels échangés entre deux comptes Proton bénéficient d’un **chiffrement bout-en-bout OpenPGP avec clés RSA 4096 bits ou Curve25519**, lui-même renforcé par AES-256 en stockage. Lorsque la destination est externe (par exemple Gmail), Proton propose deux modes : envoi en clair via TLS, ou envoi chiffré protégé par mot de passe transmis hors canal. La clé privée est stockée chiffrée côté client à l’aide du mot de passe utilisateur ; le serveur Proton ne peut donc *techniquement pas* déchiffrer le contenu, propriété démontrée publiquement par les audits du cabinet allemand *Cure53* et du laboratoire *SEC Consult*.

Gmail repose sur une architecture différente. Les courriels sont chiffrés **en transit (TLS 1.3) et au repos (AES-128 GCM)**, mais Google détient les clés et conserve un accès logique pour appliquer ses filtres anti-spam, son indexation et – sur les comptes Workspace activés – la fonctionnalité Gemini. Pour bénéficier d’un chiffrement bout-en-bout, l’administrateur doit activer **S/MIME hébergé côté client (CSE)**, fonctionnalité réservée aux plans Enterprise et qui requiert un fournisseur d’identité tiers ainsi que la gestion d’une PKI interne. Ce mode CSE existe depuis avril 2023 mais reste marginal en France et en Europe en raison de la complexité de déploiement et du coût.

### Modèle de menace et hypothèses

La différence pratique se concrétise dans le modèle de menace. Si l’on suppose un acteur malveillant ayant compromis les serveurs ou un acteur étatique exigeant un accès via assignation légale (FISA, National Security Letter aux États-Unis), Gmail sans CSE expose le contenu en clair tandis que Proton Mail expose uniquement les enveloppes (expéditeur, destinataire, horodatage) et les corps chiffrés. Andy Yen, fondateur et CEO de Proton AG, a rappelé en 2024 dans une intervention publiée sur le blog officiel sa conviction simple : on ne peut pas déchiffrer ce dont on ne possède pas les clés, et c’est par construction et non par accident.

L’ajout récent de **Proton Sentinel**, mode d’hyperdéfense activable en quelques clics depuis 2023, complète le dispositif : surveillance comportementale renforcée des connexions, blocage automatique des tentatives de phishing par typosquattage, alertes en temps réel et traçabilité avancée. Côté Google, les comptes les plus exposés (journalistes, dirigeants, activistes) peuvent activer **Advanced Protection Program**, qui impose des clés Titan FIDO2 et restreint les applications tierces. Les deux dispositifs convergent sur l’idée d’un mode haute sécurité optionnel, mais Sentinel reste le seul à couvrir un compte gratuit Proton, là où l’APP Google requiert au minimum un Workspace.

## Souveraineté des données : la fracture Suisse vs États-Unis

L’argument central de Proton dans le contexte européen est juridique. La Suisse n’est pas membre de l’Union européenne, mais bénéficie d’une décision d’adéquation de la Commission depuis 2000, renouvelée et confirmée. Le pays applique la **Loi fédérale sur la protection des données (LPD/FADP) révisée le 1er septembre 2023**, alignée sur le RGPD et publiée au *Recueil officiel* via la plateforme Fedlex. Surtout, la Suisse n’est pas soumise au **CLOUD Act américain** (Clarifying Lawful Overseas Use of Data Act, mars 2018) qui contraint tout fournisseur sous juridiction américaine à livrer les données stockées partout dans le monde sur réquisition.

Google, malgré sa *Data Region* Workspace permettant de localiser les données « primaires » dans l’UE depuis 2021, reste une société de droit américain. La **CNIL** a publié en novembre 2022 une mise en demeure très commentée concernant l’usage de Google Analytics, et plusieurs DPO de l’enseignement supérieur français ont engagé en 2023-2024 des migrations de Workspace vers des suites souveraines (BlueMind, Tchap, Mailo, Infomaniak). En décembre 2024, le ministère français de l’Éducation nationale a confirmé l’expérimentation de **La Suite numérique**, alternative open source à Google Workspace et Microsoft 365, déployée dans plusieurs académies pilotes. Ce mouvement souverain renforce l’argumentaire commercial de Proton Mail Business sur le marché public.

### Transparence des requêtes étatiques

Les rapports de transparence permettent de mesurer concrètement l’exposition. Proton publie **chaque trimestre** le nombre d’assignations reçues : 6 379 demandes contraignantes en 2023 selon le rapport officiel, dont la quasi-totalité émane d’autorités suisses respectant la procédure d’entraide judiciaire. La société a contesté plusieurs requêtes devant les tribunaux suisses et publié les ordonnances correspondantes. Google publie de son côté un *Transparency Report* mondial : plus de **196 000 demandes gouvernementales** sur le seul second semestre 2024 selon les données ouvertes, dont une part significative émanant des États-Unis sous procédure FISA et NSL – couverte par des ordres de bâillonnement empêchant toute notification à l’utilisateur.

## Tarifs Proton Mail vs Gmail en 2026

La comparaison tarifaire dépend du périmètre comparé. Si l’on rapporte le coût mensuel à un compte unique avec un domaine personnalisé et 15 Go d’espace, l’écart est minime entre Proton Mail Plus et Workspace Business Starter. Mais l’inclusion de bundles (VPN, Drive, Pass) bouscule l’équation à partir du tier Unlimited.

| Plan | Proton Mail (€ TTC) | Stockage | Gmail / Workspace (€ HT) | Stockage | 
|---|---|---|---|---|
| Gratuit | 0 €/mois | 1 Go | 0 €/mois | 15 Go | 
| Plus / Business Starter | 4,99 €/mois | 15 Go | 6,00 €/utilisateur/mois | 30 Go | 
| Unlimited / Standard | 12,99 €/mois | 500 Go | 12,00 €/utilisateur/mois | 2 To | 
| Family / Plus | 29,99 €/mois (6 utilisateurs) | 3 To | 18,00 €/utilisateur/mois | 5 To | 
| Business / Enterprise | 7,99 €/utilisateur/mois | 500 Go | Sur devis | 5 To+ négociable | 
| VPN inclus | Oui (Unlimited, 1 700+ serveurs) | – | Non (Google One VPN abandonné juin 2024) | – | 
| Pass (mots de passe) | Oui (Unlimited) | – | Non (Google Password Manager dans Chrome) | – | 

L’analyse est sans équivoque : pour un utilisateur souhaitant **e-mail + VPN + gestionnaire de mots de passe + cloud** dans une seule offre cohérente, **Proton Unlimited à 12,99 €/mois TTC** est plus avantageux que l’addition Gmail Plus (6 €) + un VPN tiers (8 à 12 €) + un gestionnaire de mots de passe (3 à 4 €) qui dépasse facilement **17 à 22 €/mois**. À l’inverse, pour une entreprise déjà ancrée dans l’écosystème Google (Sheets, Meet, AppSheet, BigQuery), le coût caché d’une migration et la perte d’intégration dépassent l’écart facial.

La promotion **Free Mobile / Mistral** annoncée en février 2025 a créé un précédent : opérateurs et fournisseurs européens commencent à grouper des abonnements à des outils souverains (Mistral, Proton, Infomaniak). Aucune offre groupée Free + Proton n’a été confirmée à ce jour, mais le marché évolue dans cette direction.

## Stockage et limites : 1 Go vs 15 Go en gratuit

Le talon d’Achille de Proton Mail Free saute aux yeux : **1 Go partagé entre Mail, Drive, Calendar et Pass**, contre **15 Go Gmail/Drive/Photos**. Pour un utilisateur grand public qui reçoit plusieurs gigaoctets de pièces jointes par an, le plan Proton gratuit sature en quelques mois. La limite d’envoi est également plus stricte : **150 messages par jour sur le compte Free Proton**, sans alias multiples, alors que Gmail autorise jusqu’à 500 envois quotidiens depuis l’interface web.

