---
id: collect-250926-servers-hardware/servers-hardware/on-a-testa-le-minisforum-n5-max-ai-nas-le-nas-cinq-baies-qui-joue-aussi-les-mini-pc-surpui-2
title: "on-a-testa-le-minisforum-n5-max-ai-nas-le-nas-cinq-baies-qui-joue-aussi-les-mini-pc-surpuissants"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Microsoft"]
dates: []
keywords: ["amd", "gpu"]
source: docs/RAG/clean4/on-a-testa-le-minisforum-n5-max-ai-nas-le-nas-cinq-baies-qui-joue-aussi-les-mini-pc-surpuissants.md
source_anchor: ""
source_lines: [51, 94]
sha256: d4947b010861a9d5dfb7f489b6978d53537fa529c3f93380243c2357f9e3863d
---

# on-a-testa-le-minisforum-n5-max-ai-nas-le-nas-cinq-baies-qui-joue-aussi-les-mini-pc-surpuissants

Avant cela, Ã©voquons aussi la partie RAM : Minisforum propose, au choix, 64 Go ou 128 Go de LPDDR5 soudÃ©s au CPU. Câest Ã prendre en compte : si le stockage et la connectique sont largement Ã©volutifs, il nâen est pas de mÃªme pour la mÃ©moire vive qui, rappelons le, est partagÃ©e entre le CPU et le GPU.

### LibertÃ© du systÃ¨me dâexploitation

Nous le disions, Minisforum livre son N5 MAX AI NAS avec un SSD M.2 dâune capacitÃ© de 128 Go sur lequel lâOS Miniscloud est prÃ©installÃ©. Fonctionnel, cet OS assure le service minimum pour un PC/NAS aussi polyvalent. Tout ce qui est partage de fichiers, gestion des utilisateurs et sauvegarde est au menu dâun OS qui ne peut toutefois pas concurrencer les logiciels de spÃ©cialistes du NAS depuis des annÃ©es. Synology bien sÃ»r, mais aussi ASUSTOR, QNAP ou Terramaster font beaucoup mieux avec un logiciel bien plus complet, plus agrÃ©able et plus fonctionnel.

*Encore trÃ¨s imparfaite, l’interface de l’OS Miniscloud contient encore de petites portions en chinois.*

Miniscloud manque parfois de rÃ©activitÃ© et mÃªme si Minisforum enchaÃ®ne les mises Ã jour, il souffre encore de pas mal de bugs handicapants. Plus gÃªnant encore, les services et les fonctionnalitÃ©s sont limitÃ©s, malgrÃ© la prÃ©sence, par exemple, dâun docker manager pour justement permettre de sâouvrir Ã beaucoup de choses pas prÃ©vues de base. Ce ne sont que quelques exemples, mais la partie surveillance matÃ©rielle du PC/NAS est assez limitÃ©e, certaines options apparaissent encore en chinois et lâinterface nâest pas toujours trÃ¨s Â« propre Â».

Reste que Minisforum est conscient des limites de Miniscloud et il autorise, voire il encourage la mise en place de votre propre OS. Pour ce faire, Ã la maniÃ¨re de certains fabricants comme Terramaster, un port USB-A a Ã©tÃ© placÃ© en interne pour simplifier le dÃ©ploiement de live-OS sur clÃ© USB.

Bien sÃ»r, il est aussi possible dâaccÃ©der au BIOS du N5 MAX AI NAS pour complÃ¨tement revoir le boot et dÃ©cider dâinstaller lâOS de son choix. Minisforum fait plutÃ´t bien les choses puisque tous les pilotes Windows sont disponibles au tÃ©lÃ©chargement depuis son site.

Mieux, puisque nous parlons dâune plateforme AMD (CPU, GPU, NPU et chipset), de nombreux pilotes existent, et ce, que lâon parle de Windows ou de Linux. On pourra dÃ¨s lors fort bien crÃ©er sa propre configuration ou se reposer sur des solutions tout-en-un comme OpenMediaVault, Rockstor, TrueNAS ou Unraid. Nous avons Ã©tÃ© impressionnÃ©s par la facilitÃ© avec laquelle nous avons Ã©tÃ© en mesure de dÃ©ployer ce genre de solutions et, surtout, par l’efficacitÃ© avec laquelle la chose a ensuite fonctionnÃ©.

Pour nos essais plus approfondis, nous avons toutefois privilÃ©giÃ© lâinstallation dâun Windows 11 afin de retomber sur quelque chose de plus classique en termes de mini-PC. Nous lâavons dit, lâaspect logiciel ne pose aucun problÃ¨me puisque tous les pilotes sont proposÃ©s par Minisforum, mais aussi parce que la plateforme AMD est une vieille connaissance. Comme sur nâimporte quel PC â ce que ce N5 MAX AI NAS est, finalement â il ne faut guÃ¨re que quelques dizaines de minutes pour disposer dâun Windows 11, parfaitement fonctionnelâ¦ mais sans licence.

## Et Ã lâusage, Ã§a tourne comment ?

Nous avons testÃ© le N5 MAX AI NAS dans ses deux configurations : en NAS et en mini-PC.

### En usage NAS

Bien que le N5 MAX AI NAS ait lâavantage de rÃ©unir les fonctions de NAS et de mini-PC au sein dâune seule et mÃªme machine, nous avons dÃ©cidÃ© de diviser nos essais en deux grandes catÃ©gories. Nous avons ensuite commencÃ© par le fonctionnement purement NAS de la machine.

Un fonctionnement qui passe par lâinsertion de plusieurs unitÃ©s dans les baies disponibles. Sans surprise, tout ce que nous avons dâun NAS de 2026 est ici prÃ©sent : des berceaux capables dâaccepter des unitÃ©s 2,5 pouces (vis fournies nÃ©cessaires) ou 3,5 pouces (systÃ¨me sans vis), des capacitÃ©s RAID jusquâaux RAID 5 et RAID 6 pour avoir une protection en cas de panne et le choix entre ZFS et ext4. Non, BTRFS nâest pas au menu.

Nous nây reviendrons pas ici, mais Ã lâheure actuelle, lâOS Miniscloud ne rend pas justice au matÃ©riel de Minisforum. Certes, il contient toutes les fonctionnalitÃ©s dâun NAS de 2026, mais dans la plupart des cas, Ã§a reste assez basique : les outils multimÃ©dias (Albums, Films, Musique) sont trÃ¨s pauvres, il nây a pas de module de surveillance vidÃ©o, le monitoring du systÃ¨me dispose dâune ergonomie horrible et le tout petit portail dâapplications souligne ces limites.

Pour autant, cela reste exploitable. Le partage des donnÃ©es ne pose aucun problÃ¨me et la puissance de la configuration permet de travailler Ã plusieurs sur le NAS sans difficultÃ©. Ci-dessous, nous avons mesurÃ© les performances du NAS via CrystalDiskMark et nous sommes un peu surpris des rÃ©sultats. Non que les dÃ©bits soient mauvais, mais nous sommes assez loin de saturer le contrÃ´leur 10 GbE.

Pour en avoir le cÅur net, nous avons rÃ©alisÃ© le mÃªme test CrystalDiskMark aprÃ¨s avoir changÃ© de systÃ¨me dâexploitation et troquÃ© Miniscloud pour Windows 11. Comme vous pouvez le voir, les rÃ©sultats nâont rien Ã voir, preuve que nos unitÃ©s de stockage (des SSD Kingston DC600M) ne sont pas en cause : le contrÃ´leur rÃ©seau se perd un petit peu en route, sans que ce soit scandaleux.

Avant de profiter plus complÃ¨tement du N5 MAX AI NAS sous Windows 11, nous avons continuÃ© nos tests NAS pour une petite vÃ©rification Â« de routine Â». Nous avons volontairement cassÃ© la pile RAID 6 que nous avions crÃ©Ã©e pour Ã©valuer le temps nÃ©cessaire Ã sa reconstruction. Eh bien sachez que le Ryzen AI MAX+ 395 fait des merveilles : câest simple, nous nâavions jamais vu un NAS reconstruire aussi vite notre pile constituÃ©e de 100 Go rÃ©partis en huit gros fichiers et de 10 Go rÃ©partis en plus de 4 000 petits fichiers. Un record, câest tout.

### En usage mini-PC

Lâautre versant du N5 MAX AI NAS, câest son utilisation en tant que mini-PC. LÃ , vous pourrez installer pas mal de choses, mais nous nous sommes contentÃ©s de Windows 11. Nous lâavons dit, la procÃ©dure dâinstallation nâa pas posÃ© de problÃ¨me et en quelques dizaines de minutes, nous nous retrouvons sur le bureau de lâOS de Microsoft. PrÃ©cision utile dâemblÃ©e : il est possible de gÃ©rer jusquâÃ quatre Ã©crans grÃ¢ce aux trois USB4 qui complÃ¨tent le port HDMI.

Il nây a en rÃ©alitÃ© pas grand-chose Ã dire de plus : le N5 MAX AI NAS est alors un PC Windows comme nâimporte quel PC Windows sachant que la configuration matÃ©rielle rend les choses extrÃªmement confortables avec ce minimum de 64 Go de RAM, ce processeur 16 cÅurs Zen 5 et ses 40 unitÃ©s de calcul RDNA 3.5. PlutÃ´t que de nous appesantir sur lâutilisation de la machine, nous vous proposons plutÃ´t dâen dÃ©couvrir les performances.

Sur Cinebench 2026, tout dâabord, câest la seule partie CPU qui est testÃ©e. Nous avons effectivement dÃ©laissÃ© le test GPU. Lâarchitecture Zen 5 du Ryzen AI Max+ 395 autorise dâhonnÃªtes scores single-thread/single-core, mais câest Ã©videmment en multi-threads que les choses deviennent plus intÃ©ressantes : avec plus de 7 100 points, le N5 MAX AI NAS est nettement devant le GMKtec EVO-T2S que nous testions il y a peu, mais aussi devant le MS-S1 Max signÃ© Minisforum. Les optimisations matÃ©rielles et, surtout, logicielles ne sont pas Ã©trangÃ¨res Ã cette belle performance.

