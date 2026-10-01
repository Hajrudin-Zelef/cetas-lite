---
id: collect-261001-cisco/cisco/coree-du-nord-lazarus-vise-3-fabricants-de-drones-2026-1
title: "coree-du-nord-lazarus-vise-3-fabricants-de-drones-2026"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["open source", "research", "sol"]
source: docs/RAG/collect-261001-cisco/coree-du-nord-lazarus-vise-3-fabricants-de-drones-2026.md
source_anchor: ""
source_lines: [1, 57]
sha256: 4a8501181b0846bb199e3f38e2cc4fbdc35bad61a5c97565dd5b35f055f12abd
---

# coree-du-nord-lazarus-vise-3-fabricants-de-drones-2026

Trois fabricants européens de drones et d’équipements de défense. Une fausse offre d’emploi. Un cheval de Troie baptisé ScoringMathTea. En octobre 2025, les chercheurs de l’éditeur slovaque ESET ont mis au jour une campagne de cyberespionnage menée par le groupe Lazarus, aligné sur la Corée du Nord, contre l’industrie de défense du continent. Neuf mois plus tard, cette opération, surnommée **DreamJob**, reste l’un des cas les mieux documentés de vol technologique étatique visant le secteur des drones militaires en Europe.

Le groupe responsable n’est pas un inconnu. Washington le tient pour responsable de plusieurs milliards de dollars de vols de cryptomonnaies, de l’attaque contre Sony Pictures en 2014 et de la propagation du rançongiciel WannaCry en 2017. Cette fois, sa cible n’est plus une banque ni un studio de cinéma, mais le savoir-faire industriel qui permet de concevoir des drones aujourd’hui déployés en Ukraine. Voici les faits, les chiffres et les conséquences pour l’industrie de défense européenne.

## Que s’est-il passé : l’alerte d’ESET Research sur l’opération DreamJob

Le 23 octobre 2025, ESET Research publie un communiqué détaillant une nouvelle vague de l’opération DreamJob, une série de campagnes attribuées de longue date au groupe Lazarus. Cette fois, la cible est précise : le secteur européen des drones et véhicules aériens sans pilote (UAV). Selon l’éditeur, l’offensive a touché **successivement trois sociétés** situées en Europe centrale et du Sud-Est, actives dans la conception d’équipements militaires.

Peter Kálnai, le chercheur d’ESET à l’origine de la découverte, explique dans le communiqué de presse que l’objectif de la campagne semble viser le vol d’informations techniques sur la conception de drones. Un fichier malveillant analysé par l’équipe mentionnerait même explicitement un drone, ce qui renforce cette hypothèse. La découverte a ensuite été relayée par plusieurs médias spécialisés, dont CSO Online et Security Affairs, confirmant la portée internationale de l’alerte.

## Comment l’attaque a été menée : ingénierie sociale et faux recrutements

L’accès initial ne repose sur aucune faille logicielle. Il repose sur la crédulité humaine. C’est la marque de fabrique de l’opération DreamJob depuis sa première apparition, il y a plusieurs années.

### Le leurre : une offre d’emploi sur mesure

Les attaquants contactent leurs cibles avec une offre d’emploi fictive, présentée comme prestigieuse et financièrement alléchante. La victime reçoit un document leurre, accompagné d’un lecteur PDF trojanisé qui se fait passer pour un outil légitime. Dès l’ouverture, l’infection démarre en arrière-plan, sans que l’utilisateur s’en aperçoive.

### La chaîne d’infection et l’évasion

Pour contourner la détection, Lazarus insère ses charges malveillantes dans des chaînes d’exécution complexes : droppers, chargeurs et téléchargeurs, souvent dissimulés dans des projets open source hébergés sur GitHub. La dernière évolution du groupe, selon ESET, réside dans l’introduction de nouvelles bibliothèques de « DLL proxying » et la sélection de nouveaux projets open source à trojaniser, une méthode qui améliore l’évasion sans pour autant masquer l’identité du groupe aux yeux des chercheurs habitués à sa signature.

## ScoringMathTea : le cheval de Troie au cœur de la campagne

La charge utile principale porte un nom presque anodin : ScoringMathTea. C’est un cheval de Troie d’accès à distance (RAT) qui donne aux attaquants un contrôle quasi total du poste infecté. Une quarantaine de commandes sont disponibles : manipulation de fichiers et de processus, collecte d’informations système, ouverture de connexions réseau, et téléchargement de nouvelles charges depuis un serveur de commande et contrôle.

Ce RAT n’a rien de nouveau. La télémétrie d’ESET permet d’en retracer l’usage sur plusieurs années et plusieurs continents, toujours dans une logique d’espionnage industriel ciblé plutôt que de gain financier immédiat.

| Période | Cible ou leurre | Pays | Secteur | 
|---|---|---|---|
| Octobre 2022 | Documents usurpant l’identité d’Airbus | Portugal, Allemagne | Aérospatial | 
| Janvier 2023 | Entreprise technologique | Inde | Technologie | 
| Mars 2023 | Industriel de défense | Pologne | Défense | 
| Octobre 2023 | Société d’automatisation industrielle | Royaume-Uni | Industrie | 
| Septembre 2025 | Fabricant aérospatial | Italie | Aérospatial | 
| Mars à octobre 2025 | Trois fabricants de drones et de défense | Europe centrale et du Sud-Est | Défense, UAV | 

*Source : télémétrie ESET Research, communiqué du 23 octobre 2025 et analyse technique WeLiveSecurity.*

## Trois entreprises, un objectif : la technologie des drones

ESET n’a pas rendu publics les noms des trois sociétés visées, mais en a précisé le profil : elles produisent des équipements militaires actuellement utilisés en Ukraine dans le cadre de l’aide européenne. L’une d’entre elles fabriquerait au moins deux modèles de drones déployés sur le terrain, et participerait à la chaîne d’approvisionnement de drones monorotors avancés, un type d’appareil que la Corée du Nord développe activement de son côté. Ces trois cibles évoluent dans un marché européen des drones à usage commercial et militaire évalué à 7,58 milliards de dollars en 2025 et qui devrait grimper à 8,52 milliards en 2026, selon le cabinet Mordor Intelligence, une croissance qui explique en partie l’appétit de Pyongyang pour ce savoir-faire.

Ce choix n’a rien d’accidentel. Pyongyang investit massivement dans ses capacités de fabrication de drones militaires et s’appuie largement sur la rétro-ingénierie et le vol de propriété intellectuelle pour rattraper son retard technologique, dans un contexte où le marché européen des drones militaires est lui-même estimé à plus de 6,4 milliards de dollars à l’horizon 2026 par MarketsandMarkets. Les entreprises visées fabriquent justement des matériels que la Corée du Nord produit déjà sur son propre sol, et cherche à améliorer.

## Pourquoi maintenant : la guerre en Ukraine en toile de fond

Le calendrier n’est pas neutre. ESET relève que cette campagne coïncide avec le déploiement de soldats nord-coréens en Russie, dans la région de Koursk, pour soutenir Moscou. Les chercheurs avancent une hypothèse : Lazarus chercherait à collecter du renseignement sur les systèmes d’armement occidentaux employés dans le conflit russo-ukrainien, drones en tête.

Ce lien entre cyberespionnage et effort de guerre n’est pas isolé. D’autres services de renseignement, russes notamment, ont multiplié les opérations contre des cibles européennes liées au soutien militaire à Kiev, comme le montre la campagne attribuée au groupe Turla contre des intérêts français. La différence avec Lazarus tient à la finalité : moins la perturbation immédiate que l’accumulation discrète de savoir-faire industriel.

## Qui est le groupe Lazarus

Actif depuis au moins 2009 selon ESET, voire depuis 2007 selon d’autres estimations, Lazarus (également désigné HIDDEN COBRA) est un groupe de menace persistante avancée (APT) aligné sur les intérêts de la Corée du Nord. Il se distingue par la diversité de ses opérations : espionnage, sabotage et cybercriminalité financière à grande échelle.

La société de cybersécurité Group-IB le décrit ainsi : « **Lazarus Group est une organisation de menace persistante avancée parrainée par un État, attribuée à la République populaire démocratique de Corée (RPDC), active depuis au moins 2007** », selon une analyse publiée par Group-IB.

