---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-45
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "Meta", "Microsoft", "OpenAI", "United States"]
dates: ["2026-09-03", "2026-09-27"]
keywords: ["chatgpt", "gpu", "mai"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [3560, 3609]
sha256: 1bd143fadfda8be6fc48b8b77d45e17424d35efb5cefdbda53d378c6279765be
---

# IA — Le grand dossier

| Affaire / événement | Statut au 27/09/2026 | Source |
|---|---|---|
| **Anthropic vs auteurs (Bartz et al.)** | **Accord à 1,5 Md$ approuvé le 20 juillet 2026** par la juge Araceli Martínez-Olguín (Californie du Nord) — le plus gros accord de droit d'auteur connu aux US. ~500 000 œuvres concernées, ~3 000 $/titre, 91 % des ayants droit ont déposé une réclamation. Point juridique clé du juge William Alsup (juin 2025) : **l'entraînement sur des livres légalement acquis = fair use ; le stockage d'une bibliothèque centrale de 7+ millions de livres piratés (LibGen, PiLiMi) = pas du fair use**. La source des données compte autant que l'usage. | AP News, Reuters via eWeek, Authors Guild |
| **New York Times vs OpenAI/Microsoft** | En cours. Le **3 septembre 2026**, le département de la Justice US (administration Trump) a déposé une déclaration d'intérêt **en faveur d'OpenAI**, arguant que l'entraînement est « extrêmement transformatif » (fair use). Le NYT (porte-parole Graham James) dénonce un soutien aux « entreprises d'IA à mille milliards » au détriment des créateurs. | AFP via TechXplore, 03/09/2026 |
| **Seattle Times + Newsday vs OpenAI/Microsoft** | Nouvelle plainte déposée le **4 septembre 2026** (utilisation non autorisée + dilution de marque). | sourcetrail.com, sept. 2026 |
| **Hachette, Cengage, Elsevier et al. vs Meta** | Plainte déposée en **mai 2026** devant un tribunal de New York pour vol présumé de contenus protégés. | AFP, sept. 2026 |
| Accords de licence | AP, Axel Springer, Vox Media ont signé des licences avec OpenAI ; le NYT lui-même a un accord de licence avec Amazon pour ses contenus dans les outils IA. Double jeu assumé : contentieux d'un côté, licences de l'autre. | AFP, sept. 2026 |

#### Les positions (attribuées)

- **Les ayants droit** (Authors Guild, NYT, éditeurs) : l'entraînement massif sans licence détruit le marché des œuvres originales et, à terme, tarit la matière première dont l'IA a elle-même besoin (« collapse » du contenu humain de qualité).
- **Les laboratoires** (OpenAI, Anthropic, Meta) : l'entraînement est transformatif — le modèle n'apprend pas des œuvres, il apprend des *statistiques du langage* ; exiger une licence par œuvre rendrait l'entraînement impossible et concentrerait l'IA entre les mains des seuls géants capables de payer.
- **La doctrine qui émerge** (jugement Alsup, juin 2025) : distinction entre données légalement acquises (fair use plausible) et données piratées (pas de fair use). Les labs achètent désormais des corpus et signent des licences — le « Far West » du scraping 2020-2023 se referme.

> **Pour ton RAG.** Retiens la leçon opérationnelle : la *provenance* des données devient un actif
> juridique. Documente tes sources (licences, robots.txt respectés, dates de collecte) — c'est
> exactement ce que les juges regardent.

### 1.5. Coût environnemental et énergétique

#### Les chiffres (vérifiés, sources citées)

- **Consommation mondiale des data centers en 2026 : ~565 TWh** (+26 % sur un an), selon les données Gartner citées par Axis Intelligence (juin 2026). Les serveurs optimisés IA représentent **31 %** de ce total et devraient dépasser les serveurs classiques en 2027 (Gartner).
- **Projection AIE (avril/sept. 2026)** : ~950 TWh en 2030, soit ~3 % de la demande électrique mondiale ; la demande des seuls serveurs accélérés (IA) croît de ~30 %/an dans le scénario de base. Un rapport de l'ONU (juin 2026) avertit que les data centers pourraient consommer près du triple de la consommation électrique combinée du Pakistan, du Bangladesh et du Nigeria.
- **États-Unis** : les data centers représentaient ~4 % de l'électricité US en 2022 ; le Lawrence Berkeley National Laboratory (actualisé 2026) projette **9,5 à 15,3 % en 2030** (cas de référence : 11,8 %). Moody's (sept. 2026) chiffre à 110 Md$ les nouvelles centrales nécessaires d'ici 2030 rien que pour les US.
- **Irlande** : les data centers dépassent **20 %** de la consommation électrique nationale (AIE) ; l'objectif de 32 % en 2026 évoqué en 2024 est en débat — « à vérifier ».
- **Eau** : jusqu'à ~300 000 gallons/jour (~1 100 m³/jour) par data center hyperscale pour le refroidissement (chiffre cité par informedclearly.com, « à vérifier » sur source primaire).
- **Investissements réseau** : ~6 700 Md$ cumulés d'ici 2030 au niveau mondial seraient nécessaires pour les réseaux (chiffre cité, « à vérifier ») ; Goldman Sachs estime 720 Md$ d'infrastructures réseau rien que pour accompagner la demande des data centers.
- **Nucléaire** : les hyperscalers ont engagé plus de 100 Md$ dans 13 projets nucléaires totalisant 9,8 GW (2026, « à vérifier » sur source primaire) — le retour du nucléaire civil tiré par l'IA est un fait structurant.
- **Gartner (juin 2026)** : **40 % des data centers IA feront face à des contraintes opérationnelles liées aux pénuries d'électricité d'ici 2027**. Linglan Wang (Gartner) : « la disponibilité de l'énergie est la nouvelle contrainte du passage à l'échelle de l'IA » — plus les GPU.

#### Le débat (attribué, sans parti pris)

- **Les alarmistes énergétiques** : l'IA fait exploser une demande électrique que les réseaux ne peuvent pas suivre (files d'attente de raccordement de 5 ans, délais de transformateurs > 3 ans), fait grimper les factures des ménages (+25 à 30 Md$/an aux US selon Moody's) et menace les objectifs climatiques.
- **Les relativistes** : les data centers restent ~1,5 % de la consommation mondiale (2024) ; l'IA *optimise* aussi l'énergie (pilotage de réseaux, conception de matériaux, fusion) ; l'efficacité par calcul (FLOPS/watt) s'améliore ; et chaque révolution (électrification, climatisation, Internet) a suscité les mêmes alertes.
- **Le point aveugle** : presque personne ne publie la consommation *par modèle* de façon auditable. Les chiffres « une requête ChatGPT = X Wh » qui circulent sont des estimations non vérifiées — **à vérifier systématiquement** avant de les citer.

> **Pour un chef de service énergies.** C'est ton terrain : l'IA est à la fois une charge (tes
> futurs data centers/edge) et un outil (optimisation énergétique, maintenance prédictive).
> Le sujet « IA et énergie » mérite sa place dans ton RAG — avec les chiffres ci-dessus,
> sourcés et datés.

### 1.6. Éducation

- **Pour** : tuteur personnel infini, patient, adaptatif — la promesse du « 2 sigma » de Bloom (1984 : le tutorat individuel fait gagner 2 écarts-types) enfin scalable. Études 2024-2025 (Khan Academy/Khanmigo, « à vérifier » sur publications) rapportent des gains d'engagement.
- **Contre** : l'étude NBER (Cruces et al., fév. 2026) elle-même alerte : les gains de performance *avec* IA peuvent refléter une **délégation** (l'outil fait le travail) plutôt qu'un apprentissage ; des travaux cités (Bastani et al., 2025 ; Shen & Tamkin, 2026) montrent une **baisse de la performance ultérieure sans IA** quand l'utilisateur s'est contenté de déléguer. Traduction : l'IA rend meilleur *avec* elle, pas forcément *sans* elle.
- **Le consensus prudent** : l'IA est un excellent tuteur *si* l'apprenant reste engagé (effort actif) ; c'est un piège *s'il* délègue. Pour la formation de tes équipes : fais faire *d'abord* sans IA, *ensuite* avec IA, puis *de nouveau* sans IA pour vérifier l'acquis.

### 1.7. Synthèse du débat — tableau des positions

