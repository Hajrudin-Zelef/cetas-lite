---
id: collect-250926-servers-hardware/servers-hardware/on-a-testa-le-minisforum-n5-max-ai-nas-le-nas-cinq-baies-qui-joue-aussi-les-mini-pc-surpui-3
title: "on-a-testa-le-minisforum-n5-max-ai-nas-le-nas-cinq-baies-qui-joue-aussi-les-mini-pc-surpuissants"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Mistral", "Nvidia"]
dates: []
keywords: ["amd", "benchmark", "diffusion", "gpu", "llama", "mistral", "nvidia"]
source: docs/RAG/clean4/on-a-testa-le-minisforum-n5-max-ai-nas-le-nas-cinq-baies-qui-joue-aussi-les-mini-pc-surpuissants.md
source_anchor: ""
source_lines: [95, 132]
sha256: 27905ef0ef48fb679191a5a001970a3df940b1e9ae18e24e6dbea05a00a06577
---

# on-a-testa-le-minisforum-n5-max-ai-nas-le-nas-cinq-baies-qui-joue-aussi-les-mini-pc-surpuissants

Afin dâenfoncer le clou, nous avons conduit des mesures sur Blender benchmark. LÃ , le Ryzen AI MAX+ 395 se comporte Ã nouveau trÃ¨s bien, mais notons surtout que les performances sont bien plus Ã©levÃ©es dÃ¨s lors que lâon met Ã contribution la partie GPU du processeur AMD. De maniÃ¨re gÃ©nÃ©rale, les puces Radeon font bien moins que les GeForce sur ce test, mais entre un GPU et un CPU, il nây a tout de mÃªme pas photo.

Nous terminons ces mesures Â« gÃ©nÃ©ralistes Â» avec PCMark, lâoutil synthÃ©tique par excellence. En simulant dâinnombrables scÃ©narios dâutilisation dâune machine (bureautique, visioconfÃ©rence, retouche photo, montage vidÃ©o, modÃ©lisation 3Dâ¦), il offre un large panorama du potentiel de la machine et, force est de constater que câest plutÃ´t pas mal. PrÃ¨s de 10 000 points au gÃ©nÃ©ral et, surtout, plus de 16 000 tant sur le test productivity que sur le test digital content creation, câest du jamais vu sur un mini-PC sans carte graphique dÃ©diÃ©e !

La puissance et donc lâimpact de la solution Radeon 8060S Ã©taient Ã vÃ©rifier de maniÃ¨re plus prÃ©cise. Pour ce faire, nous avons fixÃ© les choses avec 3DMark et nos trois scÃ¨nes de choix : Fire Strike, Time Spy Extreme et Steel Nomad. Ã chaque fois, les rÃ©sultats sont trÃ¨s intÃ©ressants. On reste loin dâune vraie carte graphique, mais les scores sont Ã©levÃ©s et laissent augurer de beaux rÃ©sultats en face de vrais jeux vidÃ©o, mÃªme des gourmands.

*Shadow of the Tomb Raider* dâabord. Il tourne plutÃ´t trÃ¨s bien, et ce, mÃªme sans activer la moindre Â« assistance Â» comme le FSR. En 1920 x 1080, avec les dÃ©tails au minimum, nous dÃ©passons les 200 images par seconde. Alors, forcÃ©ment, on passe les dÃ©tails au maximum et lÃ , câest mieux que nâimporte quel mini-PC jamais testÃ© Ã  Frandroid : 138 images par seconde et lâassurance de ne jamais baisser sous la barre des 100 ips.

On passe Ã  beaucoup plus costaud, le *Cyberpunk 2077* des Polonais de CD Projekt RED. Vitrine technologique largement utilisÃ©e par NVIDIA, le jeu tourne maintenant aussi trÃ¨s bien sur du matÃ©riel AMD et le N5 MAX AI NAS est lÃ  pour en tÃ©moigner. Nous sommes restÃ©s en 1920 x 1080 et nous avons activÃ© le FSR 3.1 avec gÃ©nÃ©ration dâimages, mais quelle claque. Avec les dÃ©tails sur Â« bas Â», on flirte avec les 260 ips et en Â« ultra Â» on est encore Ã  prÃ¨s de 170 ips alors pourquoi ne pas activer le ray tracing ? Ãa baisse bien sÃ»r, mais le FSR fait des merveilles : en ray tracing bas on est 150 ips et en ray tracing ultra, on frÃ´le les 90 ips. Impressionnant.

Rares sont les scÃ©narios pour lesquels le N5 MAX AI NAS pourrait ne pas Ãªtre prÃªt et, afin dâenfoncer le clou, nous avons conduit quelques tests en IA locale avec des benchs Â« gÃ©nÃ©ratifs Â» conÃ§us par UL, dÃ©jÃ Ã©diteur de PCMark et 3DMark. Le logiciel Procyon offre plusieurs scÃ©narios, nous en avons retenu deux. Pour la gÃ©nÃ©ration de textes â via Phi 3.5, Mistral 7B et Llama 3.1 â les rÃ©sultats sont intÃ©ressants, mais pas non plus de quoi sauter au plafond.

Il en va dâailleurs plus ou moins de mÃªme pour la gÃ©nÃ©ration dâimages. LÃ , Procyon se base sur Stable Diffusion 1.5 et profite dâun modÃ¨le en prÃ©cision FP16. Avec 8,389 images par seconde, nous sommes en face dâune configuration tout Ã fait capable, mais on sent quâelle peut mieux faire. La faute sans doute au fait que le NPU nâest pas mis Ã contribution : la puissance de traitement vient exclusivement du GPU Radeon 8060S.

## Consommation, chauffe et nuisances sonores

Le Ryzen AI MAX+ 395 est un processeur redoutablement puissant, cela ne fait maintenant plus guÃ¨re de doute. Il est dotÃ© dâune remarquable force de calcul CPU, mais nâa clairement pas Ã rougir de sa solution graphique intÃ©grÃ©e. ProblÃ¨me, une telle puce, qui plus est placÃ©e dans un espace aussi restreint que le N5 MAX AI NAS, Ã§a doit chauffer, non ?

*IntÃ©ressante, la solution de refroidissement dÃ©ployÃ©e par Minisforum distingue le CPU des cinq SSD.*

Eh bien, reconnaissons que nous avons Ã©tÃ© agrÃ©ablement surpris par le travail des ingÃ©nieurs de Minisforum. En effet, malgrÃ© les charges les plus lourdes, le processeur nâa jamais atteint le seuil des 85 Â°C. Si certains usagers trouveront cela chaud, Ã§a reste bien loin des limites fixÃ©es par AMD et aucun throttling nâa jamais Ã©tÃ© Ã dÃ©plorer. Belle performance rendue possible par lâutilisation dâun systÃ¨me avec 5 caloducs et soutenu par deux ventilateurs 80 mm de type blower.

Il en va dâailleurs de mÃªme pour tous les autres composants, et ce, que lâon parle des unitÃ©s de stockage 3,5 pouces â refroidies par deux ventilateurs de 92 mm â ou des unitÃ©s SSD au format M.2 qui bÃ©nÃ©ficient pour trois dâentre elles dâun dissipateur musclÃ© alors que les deux autres peuvent compter sur un petit ventilateur dÃ©diÃ© (60 mm).

HÃ©las, pour les nuisances sonores, le bilan est moins reluisant. Oh, Ã  Frandroid, nous avons dÃ©jÃ  vu plus bruyant, mais le Â« *silencieux mÃªme en pleine charge* Â» vantÃ© par Minisforum est, au mieux une jolie faÃ§on de prÃ©senter les choses, au pire un mensonge. En rÃ¨gle gÃ©nÃ©rale, il faut compter autour de 36 dB, mais dÃ¨s lors que lâon Â« charge la mule Â», cela devient compliquÃ© et avec un maximum de 43 dB, le N5 MAX AI NAS devient mÃªme trÃ¨s Â« sonore Â».

Enfin, il ne faut pas non plus nÃ©gliger lâaspect consommation Ã©lectrique dâune machine certes compacte, mais capable de se mesurer Ã presque toutes les situations. Minisforum peut compter sur la relative sobriÃ©tÃ© du Ryzen AI MAX+ 395, mais en pleine charge, nous avons relevÃ© jusquâÃ 90 W pour ce seul processeur et un maximum dâenviron 167 W pour lâensemble de la machine avec cinq disques durs. Cela dit, en dehors de la facture dâÃ©lectricitÃ©, ce nâest pas trop un problÃ¨me : Minisforum a prÃ©vu une alimentation de 250 W. On est tranquille.

## Prix et disponibilitÃ©

MalgrÃ© quelques critiques, notamment au niveau du logiciel, vous aurez compris que nous avons plutÃ´t apprÃ©ciÃ© ce vÃ©ritable couteau suisse que nous propose une marque pourtant bien chinoise. HÃ©las, Minisforum ne peut pas faire de miracle et Ã lâheure de la douloureuse, rares seront les usagers Ã se laisser tenter. Lâinflation que nous connaissons depuis maintenant un an ne montrant en plus aucun signe dâessoufflement.

Disponible depuis l’Ã©tÃ©, le N5 MAX AI NAS est commercialisÃ© en deux variantes. La premiÃ¨re â que nous avons testÃ©e â est conÃ§ue autour du Ryzen AI MAX+ 395, lequel nâest accompagnÃ© Â« que Â» de 64 Go de RAM et dâun SSD de 128 Go. Cette version est facturÃ©e **2 719 euros**. La seconde mouture est dotÃ©e du mÃªme processeur et du mÃªme SSD, mais vient avec 128 Go de RAM pour la bagatelle de **3 839 euros**. Sachant que du fait mÃªme de la conception du CPU, la RAM nâest pas Ã©volutive, il faudra faire le bon choix.

Il est utile de souligner quâaussi cher soit ce PC de Minisforum, il est finalement Ã lâimage de la plupart des autres PC basÃ©s sur le Ryzen AI MAX+ 395. Câest dâailleurs lÃ quâil tire son Ã©pingle du jeu : pour un prix somme toute trÃ¨s similaire, Minisforum parvient Ã lui adjoindre toutes ces fonctionnalitÃ©s NAS pour donner Ã son mini-PC un potentiel assez remarquableâ¦ mais qui nâintÃ©ressera sans doute pas tout le monde.

Ce contenu est bloquÃ© car vous n'avez pas acceptÃ© les cookies et autres traceurs. Ce contenu est fourni par Disqus.

