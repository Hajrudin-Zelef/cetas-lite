---
id: collect-250926-servers-hardware/servers-hardware/parts-de-marche-gpu-2026-nvidia-90-amd-a-8-3
title: "Windows (PowerShell)"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: []
keywords: ["amd", "antitrust", "dram", "gpu", "hbm", "intel", "nvidia", "research"]
source: docs/RAG/clean4/parts-de-marche-gpu-2026-nvidia-90-amd-a-8.md
source_anchor: ""
source_lines: [85, 149]
sha256: 1032ae14ba438992f264f2c51e8d11911d2bde57354c895a9b4f9735c7f82557
---

# Windows (PowerShell)

Pour l’acheteur français, la conséquence la plus tangible de ce marché déséquilibré se lit sur l’étiquette. Entre la TVA, la faiblesse récurrente de l’euro face au dollar, la pénurie de mémoire et une offre volontairement limitée, les cartes graphiques se négocient couramment bien au-dessus de leur prix conseillé de lancement. Le mouvement s’est encore accentué cet été : entre juin et août 2026, les prix constatés chez le revendeur américain Newegg ont grimpé de **39 %** pour la RTX 5060 Ti 16 Go et de **36 %** pour la RTX 5070, selon une analyse de Tech-Insider – signe que la flambée gagne désormais le milieu de gamme et pas seulement les cartes premium comme la RTX 5090, lancée à 1 999 dollars début 2026. Le tableau ci-dessous rappelle les tarifs de référence des principales cartes actuelles ; les prix de détail réels en Europe leur sont, en 2026, systématiquement supérieurs.

| Carte graphique | Fabricant | Prix conseillé de lancement (USD) | 
|---|---|---|
| GeForce RTX 5090 | Nvidia | 1 999 $ | 
| GeForce RTX 5080 | Nvidia | 999 $ | 
| GeForce RTX 5070 Ti | Nvidia | 749 $ | 
| Radeon RX 9070 XT | AMD | 599 $ | 
| Radeon RX 9070 | AMD | 549 $ | 
| Arc B580 | Intel | 249 $ | 

Cette inflation matérielle alimente un débat plus politique sur la souveraineté technologique européenne. Comme nous l’analysions à propos du Chips Act 2.0 et de ses 120 milliards d’euros, l’Union européenne cherche à réduire sa dépendance vis-à-vis des fondeurs asiatiques et américains. Mais aucune de ces initiatives ne produira de carte graphique « made in Europe » avant plusieurs années : à court terme, le consommateur reste captif d’un marché dominé par une seule entreprise.

## RTX 50 contre RX 9000 : le duel qui n’a plus lieu

Sur le papier, la génération actuelle offrait pourtant les ingrédients d’une belle bataille. Face aux GeForce RTX 50 de Nvidia, AMD alignait ses Radeon RX 9070 et RX 9070 XT, saluées par la critique pour leur rapport performance-prix. Notre comparatif RX 9070 XT contre RTX 5070 Ti montrait d’ailleurs une AMD très compétitive en 1440p, à un tarif inférieur.

Mais les parts de marché GPU racontent une vérité cruelle : la qualité d’un produit ne suffit pas. La force de l’écosystème Nvidia – pilotes matures, DLSS, adoption par les développeurs de jeux, notoriété de la marque – pèse au moins autant que les performances brutes. AMD peut proposer la meilleure carte à budget égal ; si le grand public continue d’associer « carte graphique » et « GeForce », le déséquilibre perdure.

Le prochain rendez-vous se jouera sur le terrain de l’architecture. AMD prépare une nouvelle génération, tandis que le calendrier des processeurs et des GPU s’accélère, comme nous le décrivions dans notre dossier Zen 6 et la feuille de route AMD. Reste à savoir si un simple saut technologique suffira à briser des habitudes d’achat aussi ancrées.

## Ce que disent les données de Steam

Les livraisons de cartes ne disent pas tout : encore faut-il savoir ce que les joueurs utilisent réellement. Là encore, l’enquête matérielle de Steam, qui recense la configuration de dizaines de millions de joueurs à travers le monde, confirme la mainmise de Nvidia : au deuxième trimestre 2025, ses GPU équipaient précisément **74,88 %** des machines répertoriées, contre **17,32 %** pour AMD et **7,44 %** pour les solutions graphiques intégrées d’Intel, selon les chiffres de Valve relayés par HardwareTimes.

Ce décalage entre le parc installé (où AMD conserve une présence historique) et les ventes récentes (où la marque s’effondre) est révélateur. Il signifie que le fossé se creuse trimestre après trimestre : chaque nouvelle vague de mises à niveau renforce mécaniquement la position de Nvidia, car les acheteurs récents choisissent massivement GeForce. Sans inflexion, la part de marché GPU d’AMD dans le parc actif continuera de s’éroder dans les années à venir.

### Vérifier sa propre carte graphique

Curieux de savoir dans quel camp vous vous situez ? Un simple relevé suffit à identifier votre carte, sous Windows comme sous Linux :

```
# Windows (PowerShell)
Get-CimInstance Win32_VideoController | Select-Object Name, AdapterRAM
# Linux
lspci | grep -Ei 'vga|3d|display'
```
## Contexte historique : comment AMD a perdu la bataille du GPU

Le duopole n’a pas toujours été aussi déséquilibré. Au milieu des années 2010, ATI puis AMD tenaient tête à Nvidia, dépassant régulièrement 30 à 40 % du marché des cartes graphiques dédiées. Le tournant s’est amorcé avec l’arrivée du *ray tracing* matériel et de la génération d’images par IA : Nvidia a imposé le DLSS dès 2018, prenant plusieurs années d’avance sur une technologie devenue déterminante.

À cela s’est ajoutée une série de choix stratégiques d’AMD : priorité donnée aux processeurs Ryzen et aux puces des consoles PlayStation et Xbox, ressources d’ingénierie limitées côté Radeon, et une communication moins agressive. Le rachat d’ATI en 2006, censé faire d’AMD un géant du calcul graphique, n’a jamais tenu toutes ses promesses sur le marché grand public.

La crise de la mémoire de 2025-2026 n’a fait qu’accélérer un déclin déjà engagé. En délaissant le très haut de gamme et en se retrouvant en concurrence directe avec ses propres processeurs à graphique intégré sur l’entrée de gamme, AMD s’est enfermé dans un segment intermédiaire de plus en plus étroit – pendant que Nvidia raflait la mise en haut et qu’Intel s’installait en bas.

## Cinq prédictions pour le marché GPU d’ici 2029

- **AMD reste sous la barre des 10 % à court terme.** Sans une génération capable de rivaliser sur le haut de gamme et de renverser l’image de marque, Radeon devrait osciller entre 5 et 12 % du marché des cartes dédiées pendant plusieurs trimestres encore.
- **Intel consolide sa troisième place.** Porté par le segment budget et par le professionnel/IA, Arc pourrait durablement se stabiliser autour de 1 à 3 %, un socle modeste mais inédit pour le fondeur.
- **Les prix resteront élevés tant que dure la crise mémoire.** Tant que la mémoire HBM et la DRAM seront happées par les datacenters, les cartes graphiques grand public conserveront des tarifs gonflés, au moins jusqu’en 2027.
- **Le marché des cartes dédiées poursuit sa contraction.** Avec une croissance annuelle moyenne prévue à −3,3 % jusqu’en 2029, les GPU intégrés et les APU grignoteront le bas de gamme, réduisant le volume total de cartes vendues.
- **Nvidia conserve plus de 85 % du marché.** Sauf choc réglementaire majeur – enquête antitrust européenne, régulation de l’IA – la domination de Nvidia sur les parts de marché GPU devrait perdurer, portée par l’inertie de son écosystème.

## Foire aux questions

### Quelle est la part de marché de Nvidia en 2026 ?

Selon Jon Peddie Research, Nvidia détenait environ 90 % du marché des cartes graphiques dédiées de bureau au premier trimestre 2026, après avoir culminé à 94 % fin 2025. La firme vend donc environ neuf cartes graphiques sur dix.

### Pourquoi AMD perd-il des parts de marché sur les cartes graphiques ?

AMD cumule plusieurs handicaps : renoncement au très haut de gamme avec RDNA 4, retard historique sur les fonctions d’IA comme le DLSS, écosystème logiciel moins étoffé et notoriété de marque inférieure. La firme est tombée à 5 % fin 2025, son plus bas niveau historique, avant de remonter à environ 8 % début 2026.

### Intel Arc peut-il concurrencer Nvidia et AMD ?

