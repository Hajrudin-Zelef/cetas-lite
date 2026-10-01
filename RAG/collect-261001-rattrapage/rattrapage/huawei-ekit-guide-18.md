---
id: collect-261001-rattrapage/rattrapage/huawei-ekit-guide-18
title: "GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ekit_guide.md
source_anchor: ""
source_lines: [1550, 1629]
sha256: 8cfd07b88d14ea53258c75e7245a9238e06b2429603ea20a21fb6e7314af4460
---

# GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)

**Contexte :** bureau de projet 15 personnes, connexion Internet instable (4G principale, VSAT en secours), coupures électriques fréquentes, budget serré.
**Matériel :** 1× AR180 (4G en USB si supporté — **à vérifier** — sinon routeur 4G devant l'AR), 1× S220-8P4S, 2× AP361, **onduleur 2000 VA + batteries étendues** (cf. guide onduleurs).
**VLAN :** 10 (staff), 20 (invités/partenaires).
**Commentaire :** ici l'**énergie** est le sujet n°1, pas le Wi-Fi : dimensionne l'onduleur pour 2-4 h (les coupures sont la norme), active le **redémarrage auto** partout, et mets le cloud eKit en mode « supervision quand ça passe » (les alertes « site hors ligne » seront fréquentes : règle les seuils pour ne pas spammer). Pré-stage tout au siège avant envoi sur le terrain — sur place, on ne fait que brancher.

## 141. Cas pratique — Cybercafé / salle de gaming (commenté)

**Contexte :** 30 PC gaming filaires, fibre 2 Gbit/s, tournois le week-end, 20 spectateurs en Wi-Fi, 4 caméras.
**Matériel :** 1× AR280 ou USG6000F-S150 (2 Gbit/s !), 1× S310-48P4X (**uplinks 10G**), 2× AP361 (spectateurs), onduleur 3000 VA.
**VLAN :** 10 (PC gaming — **latence minimale**), 20 (spectateurs Wi-Fi, bridé), 40 (caméras).
**Commentaire :** le gaming = **latence**, pas débit : pas de QoS qui ajoute du traitement inutile, pas de Wi-Fi pour les joueurs (filaire obligatoire), firmware stable (jamais de MAJ avant un tournoi). Le S310-48P4X se justifie ici (2 Gbit/s d'uplink). Mesure le ping, pas le débit, pour valider.

## 142. Cas pratique — Boulangerie / chaîne de 5 boutiques (commenté)

**Contexte :** 5 boutiques, chacune : 6 employés, 2 caisses, 3 caméras, Wi-Fi clients, pas d'IT sur place.
**Matériel (par boutique) :** 1× AR180, 1× S220-8P4S, 1-2× AP361. **Cloud :** 1 tenant, 5 sites, modèle « boutique standard ».
**VLAN :** 10 (caisses/staff), 20 (clients), 40 (caméras).
**Commentaire :** le multi-sites eKit **brille** ici : 1 modèle de config, 5 sites identiques, alertes par boutique avec le nom du responsable. Le technicien n'a besoin d'aller sur place que pour le physique. Procédure écrite remise à chaque responsable : « si Internet tombe : 1) vérifier la box, 2) appeler l'astreinte, 3) ne toucher à rien d'autre ». Et les caisses : jamais de MAJ entre 6 h et 20 h.

## 143. Cas pratique — Cabinet comptable / d'avocats (commenté)

**Contexte :** 12 personnes, données clients ultra-sensibles, secret professionnel, VPN vers 2 clients, 3 caméras.
**Matériel :** 1× USG6000F-S125 (IPS/AV/filtrage + VPN), 1× S220-8P4S ou S310-24P, 2× AP361, onduleur 1500 VA.
**VLAN :** 10 (collaborateurs), 20 (invités, isolé, **jamais** de PSK partagé avec le staff), 40 (caméras), 99 (gestion).
**Commentaire :** secret professionnel = **USG obligatoire**, WPA3, PSK tourné au départ de chaque collaborateur, registre des accès (106) tenu à jour, et **clause de confidentialité** avec toi le prestataire. Le Wi-Fi invités : PSK différent du staff, changé régulièrement. C'est aussi un client qui paiera un contrat de maintenance premium — propose-le.

## 144. Cas pratique — Station-service / boutique autoroute (commenté)

**Contexte :** boutique + 6 pompes (automates), TPE, 8 caméras, Wi-Fi clients, site isolé, 24h/24.
**Matériel :** 1× AR280 ou USG, 1× S310-24P4S, 2× AP361 (boutique) + 1× AP761 (piste, zone ATEX à respecter !), onduleur 2000 VA.
**VLAN :** 10 (boutique/TPE), 20 (clients), 40 (caméras), 50 (automates pompes — **isolé, critique**).
**Commentaire :** les **automates de pompes** = système critique, VLAN dédié, jamais touché sans l'accord de l'exploitant pétrolier (souvent un prestataire dédié — **coordination obligatoire**). Zone ATEX : aucun AP actif dans les zones classifiées sans certification — l'AP761 couvre **depuis** l'extérieur de la zone. Site 24/24 : fenêtres de maintenance à 3 h du matin, jamais de MAJ sans astreinte sur place.

## 145. Cas pratique — Résidence hôtelière / appart-hôtel (commenté)

**Contexte :** 30 appartements, séjours longs, box TV par appartement, 60+ clients simultanés le soir, laverie commune, parking.
**Matériel :** 1× AR280 ou USG6000F-S125, 2× S310-48P4S, 15× AP361 (1 pour 2 apparts) + 1× AP761 (parking/laverie), onduleur 3000 VA.
**VLAN :** 10 (staff/réception), 30 (résidents — portail captif, baux longs 7 jours), 40 (caméras), 70 (TV si IP).
**Commentaire :** séjours longs = **baux DHCP longs** (pas de coupure à minuit quand le bail expire en plein film), **isolation inter-clients** (les résidents restent des semaines : ils *essaieront* de voir le voisin), et un **débit équitable** par client (sinon 3 gamers monopolisent la fibre). La laverie et le parking : souvent oubliés dans le chiffrage — prévois-les dès le départ.

---

## 146. Tableau de référence — canaux Wi-Fi (plan de fréquences)

| Bande | Canaux utilisables (Europe/CEPT, ordre de grandeur) | Bonnes pratiques |
|---|---|---|
| 2,4 GHz | 1 à 13 (20 MHz) — canaux non chevauchants : **1, 6, 11** | N'utiliser que 1/6/11 en 20 MHz. Réserver le 2,4 GHz aux objets connectés et au secours. |
| 5 GHz | 36-64, 100-140 (20/40/80 MHz selon pays) | 40 MHz en bureau dense, 80 MHz si peu d'AP voisins. Éviter les canaux DFS (52-140) si des coupures radar sont constatées. |
| 6 GHz (Wi-Fi 7) | Selon réglementation locale | **Vérifier la disponibilité légale dans ton pays** avant de déployer l'AP673H en 6 GHz. |

**Règle :** en multi-AP, alterne les canaux (1/6/11 au rez, 6/11/1 à l'étage) et baisse la puissance plutôt que de monter les canaux en puissance. La réglementation locale (puissance max, canaux autorisés) prime toujours — **à vérifier sur la documentation officielle** et les textes locaux.

## 147. Tableau de référence — câblage et distances

| Support | Distance max | Usage eKit typique |
|---|---|---|
| Cuivre Cat6 (PoE) | **90 m** (lien permanent) + 10 m de jarretières | AP, caméras, téléphones |
| Cuivre Cat6A | 90 m (10G jusqu'à 100 m) | Uplinks 10G (S310-48P4X), serveurs |
| Fibre OM3/OM4 (multimode) | 300-400 m (10G) | Dorsales inter-bâtiments, colonnes |
| Fibre OS2 (monomode) | 10 km et + | Inter-sites, MiniFTTO (selon OLT) |
| SFP+ DAC (cuivre twinax) | 1-3 m (parfois 5-7 m) | Liaison switch-switch dans la même baie |

**Règles :** jamais de câble cuivre > 90 m pour du PoE (chute de tension + erreurs) ; en extérieur : **câble extérieur (UV) ou gaine**, jamais du câble intérieur au soleil ; teste chaque brin au testeur (continuité + cartographie des paires), pas « à l'œil ».

## 148. Tableau de référence — budgets types par scénario (ordres de grandeur)

| Scénario | Matériel eKit type | Échelle de budget matériel* |
|---|---|---|
| Boutique (C) | AR180 + S220-8P4S + 2 AP361 | € (le plus accessible) |
| Bureau 20 postes (A) | AR180 Pro + S310-24P + 3 AP361 | €€ |
| Restaurant (131) | AR180 + S220 + 2-3 AP | € |
| Entrepôt (D) | AR280 + S310-24P + 4-6 AP | €€€ |
| Hôtel 40 ch. (B) | AR280/USG + 2 S310-48P + 20 AP361 + 2 AP761 | €€€€ |
| École 12 classes (132) | AR280 + 2 S310-48P + 12-14 AP | €€€€ |
| Clinique (133) | USG-S125 + 2 S310-48P + 10 AP | €€€€ |

*Échelle relative uniquement — **les prix réels sont à vérifier auprès de ton distributeur** (ils varient par pays, stock et remise). Ajoute toujours : câblage (~30-40 % du matériel en rénovation), onduleur, main-d'œuvre, contrat de maintenance annuel (~15-20 % du matériel/an, ordre de grandeur).

## 149. Comparatif honnête — eKit vs alternatives PME

