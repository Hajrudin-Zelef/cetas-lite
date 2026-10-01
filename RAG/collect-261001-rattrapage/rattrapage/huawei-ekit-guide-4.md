---
id: collect-261001-rattrapage/rattrapage/huawei-ekit-guide-4
title: "GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-27"]
keywords: ["diffusion"]
source: docs/RAG/collect-261001-rattrapage/huawei_ekit_guide.md
source_anchor: ""
source_lines: [234, 319]
sha256: 4dd6ecaf176f0f86118e01201bbfa35357ecb7b8f3dfb3c679098d7f56bec827
---

# GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)

D'après la fiche technique (déploiement « Wi-Fi-based ») :
1. Mets l'AP sous tension (PoE ou adaptateur), attends qu'il diffuse son **Wi-Fi de gestion** (SSID de setup d'usine).
2. Sur le téléphone, ouvre l'app eKit, crée ou sélectionne le **site**, lance l'ajout d'équipement en mode Wi-Fi.
3. Connecte le téléphone au Wi-Fi de gestion de l'AP quand l'app le demande.
4. L'app pousse la configuration du projet (SSID, mot de passe, paramètres réseau) et **onboarde automatiquement** l'équipement au cloud.
5. Vérifie dans l'app que l'AP apparaît **en ligne** sur le site avant de passer au suivant.

**Conseil terrain :** fais l'onboarding **au bureau avant d'aller sur site** quand c'est possible (pré-staging) : tu scannes/configures tout au calme, sur site tu ne fais que brancher. Tu divises le temps d'intervention par deux.

## 20. Onboarding par scan de code-barres / QR : la méthode robuste

D'après la fiche technique (déploiement « barcode scanning-based ») :
1. Dans l'app, sur le site concerné, choisis l'ajout par **scan du SN** (le numéro de série sur l'étiquette du châssis ou de l'emballage).
2. Scanne chaque équipement (AP, switch, passerelle). L'app synchronise les SN vers le système eKit.
3. Quand les équipements sont branchés et ont accès à Internet, ils **remontent automatiquement** au cloud et rejoignent le site.
4. Cas d'usage typiques : ajout de capacité à un projet existant ; équipement dont le firmware ne supporte pas le déploiement par Wi-Fi (dans ce cas, **mets à jour d'abord**, ou passe par le scan).

**Réflexe :** photographie les étiquettes SN **avant** de fixer les AP au plafond. Une fois l'AP à 3 mètres de haut, tu ne veux pas ressortir l'échelle pour lire un SN.

## 21. Onboarding d'un switch ou d'une passerelle AR

Même logique que les AP, via l'app :
1. **Câblage d'abord** : la passerelle AR en tête (port WAN vers le modem/ONT de l'opérateur), le switch derrière (uplink vers l'AR), les AP sur les ports PoE du switch.
2. Allume dans l'ordre : **modem opérateur → AR → switch → AP** (l'ordre de mise en service détaillé en section 75).
3. Scanne ou déclare chaque équipement dans le site via l'app.
4. La passerelle AR récupère son accès Internet (DHCP ou PPPoE selon l'opérateur — paramètres à saisir dans l'app), puis les équipements derrière elle remontent au cloud **à travers elle**.
5. **Point de vigilance :** si l'AR n'a pas Internet, **rien ne remonte**. En cas d'échec global d'onboarding, commence toujours par vérifier l'accès Internet de la passerelle (section dépannage 91).

## 22. Premier paramétrage d'un site : l'ordre qui évite les allers-retours

1. Crée le site (nom explicite, adresse, fuseau horaire).
2. Déclare la **passerelle** (accès Internet : DHCP/PPPoE/IP statique + DNS).
3. Crée les **VLAN** (section 67) : gestion, utilisateurs, invités, IoT/caméras, voix si besoin.
4. Crée les **SSID** (section 68) : un SSID « staff » (WPA2/WPA3), un SSID « invités » (portail captif ou PSK simple, isolé), éventuellement un SSID IoT.
5. Déclare les **switchs** (activation PoE, VLAN natifs sur les ports AP).
6. Déclare les **AP** (association aux SSID, canaux/puissance — laisse l'auto par défaut au début, optimise après).
7. Teste : un client sur chaque SSID, un ping vers Internet, un test de débit.
8. **Sauvegarde** : exporte la configuration (section 89) et note-la dans le dossier du site.

## 23. Ce que tu configures dans l'app au quotidien

D'après les fiches techniques, l'app permet après déploiement : configuration, monitoring, inspection des équipements sur le cloud. Concrètement, attends-toi à pouvoir gérer :
- état en ligne/hors ligne des équipements par site ;
- SSID (création, PSK, diffusion ou non du SSID) ;
- paramètres radio de base (bande, canal, puissance — niveau de détail **à vérifier sur la documentation officielle** selon version de l'app) ;
- VLAN et adressage de base ;
- PoE (activation/désactivation par port — **à vérifier sur la documentation officielle**) ;
- redémarrage à distance d'un équipement ;
- mise à jour firmware (section 90).
Ce que tu ne feras **pas** dans l'app : du tuning radio fin, des ACL complexes, du routage avancé — pour ça, bascule l'équipement en gestion locale (CLI/web) si le modèle le permet (les switchs S310/S620 le permettent d'après leurs fiches).

## 24. L'app pour le technicien sur site : les bons réflexes

- **Mode hors-ligne partiel :** prépare tout ce qui est préparables avant de partir (site créé, équipements pré-déclarés par scan des cartons). Sur site, tu n'as besoin que d'Internet sur ton téléphone pour finaliser.
- **Batterie externe** dans la sacoche : l'app + le flash + les photos de chantier vident un téléphone en une demi-journée.
- **Photos systématiques** : baie, étiquettes, cheminement des câbles, position des AP. Tu les ranges dans le dossier du site — dans 6 mois, tu remercieras ton toi du passé.
- **Ne rends jamais un site « à l'aveugle »** : avant de partir, fais le tour avec le téléphone — tous les équipements verts dans l'app, un test client par SSID, et une capture d'écran de l'état du site.

## 25. Wi-Fi de gestion et sécurité pendant le déploiement

Pendant l'onboarding par Wi-Fi, l'équipement diffuse un SSID de configuration. Règles :
- Ne laisse **jamais** un équipement en mode setup sur un site en production : termine l'onboarding ou éteins-le.
- Change les mots de passe par défaut **dès** la première configuration (section 102).
- Si tu pré-stages au bureau, isole les équipements de ton réseau de production (un petit switch ou un VLAN dédié) pour éviter qu'un DHCP sauvage ne perturbe ton LAN.

## 26. Gestion des firmwares via l'app : le minimum vital

- L'app propose les **mises à jour** des équipements du site (via le mécanisme HOUP côté switch, et l'équivalent côté AP/AR — **à vérifier sur la documentation officielle** pour le détail par modèle).
- Règle d'or : **ne mets jamais à jour le jour de la mise en service**. Stabilise d'abord, mets à jour ensuite, en heures creuses, site par site.
- Avant toute mise à jour : **sauvegarde la config** (section 89), préviens le client, et ne mets à jour qu'**un site pilote** d'abord en multi-sites.
- Note la version installée dans le dossier du site (ex. fictif : `AR180 – V100R001C00SPC100 – 2026-09-27`). Si ça casse quelque chose, tu sais d'où tu reviens.

## 27. L'app eKit et les notifications : les régler pour ne pas devenir fou

- Active les alertes **équipement hors ligne** et **nouvel équipement détecté** pour tes sites : ce sont les deux alertes qui valent de l'or en maintenance.
- Désactive (ou regroupe) les alertes informatives en masse si l'app les propose — un technicien qui reçoit 50 notifications par jour finit par toutes les ignorer, y compris les importantes.
- En multi-sites, vérifie que les alertes sont bien **rattachées au site** : « AP hors ligne » ne sert à rien si tu ne sais pas dans quelle boutique.

## 28. Partager un site avec le client (ou pas)

- Question à trancher **contractuellement** : le client a-t-il accès à l'app sur son site ? Avantages : il redémarre un AP tout seul à 22 h. Risques : il change le PSK, casse un VLAN, puis t'appelle en disant « ça ne marche plus depuis votre passage ».
- Pratique recommandée : donne au client un accès **lecture seule** s'il existe (**à vérifier sur la documentation officielle** — sinon, pas d'accès du tout et un numéro d'astreinte), et garde l'administration chez toi dans le cadre du contrat de maintenance.
- Si le client exige l'admin : fais-le par écrit, avec une clause « toute modification hors contrat = intervention facturée ».

## 29. iGuard et la détection de caméras espion : mode d'emploi honnête

