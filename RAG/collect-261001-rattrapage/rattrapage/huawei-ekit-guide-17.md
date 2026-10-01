---
id: collect-261001-rattrapage/rattrapage/huawei-ekit-guide-17
title: "GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)"
domain: rattrapage
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["diffusion", "incident"]
source: docs/RAG/collect-261001-rattrapage/huawei_ekit_guide.md
source_anchor: ""
source_lines: [1477, 1549]
sha256: 6379a88bd9440365866d07e6797c679daf507ae80408da6d11d8c95b607e3a57
---

# GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)

- **Modèles :** S125 (~600 utilisateurs, 1,7 Gbit/s NGFW) / S150 (~1000, 2,1 Gbit/s) / S200 (~2500, 3 Gbit/s). Débits « enterprise mix » : 1 / 1,2 / 1,8 Gbit/s avec tout activé.
- **Fonctions :** pare-feu, IPS, antivirus, filtrage URL (500 M d'URL), IPsec (2000-4000 tunnels), SSL VPN (100 inclus, jusqu'à 1000-2000), portail captif local, SSO RADIUS/AD/LDAP, SD-WAN sécurisé.
- **Dimensionnement :** on dimensionne sur le **débit tout activé**, pas le débit brut (section 7).
- **Cas d'usage :** PME avec données sensibles, hôtel (portail + sécurité), multi-sites avec VPN.
- **À vérifier :** licences des fonctions avancées et abonnement signatures auprès du distributeur avant chiffrage.

---

## 131. Cas pratique — Restaurant 80 couverts (commenté)

**Contexte :** restaurant avec terrasse, 12 employés, caisse + 2 TPE, Wi-Fi clients gratuit, 4 caméras, pas d'IT.
**Matériel :** 1× AR180, 1× S220-8P4S, 2× AP361 (salle + terrasse couverte), 1× AP761 si terrasse très exposée (à la place du 2e AP361).
**VLAN :** 10 (caisse/TPE/staff), 20 (clients, portail captif), 40 (caméras).
**Commentaire :** le point critique, ce sont les **TPE** : jamais de MAJ en service, VLAN 10 prioritaire, et un **câble 4G de secours** sur la caisse si le restaurateur l'accepte (un samedi soir sans TPE = catastrophe). Le portail captif clients = page avec la carte des desserts + avis Google (marketing). Puissance radio modérée : la terrasse ne doit pas arroser toute la rue (sinon 30 passants squattent le Wi-Fi).

## 132. Cas pratique — École primaire 12 classes (commenté)

**Contexte :** 12 classes, bureau direction, 300 élèves (pics : 1 classe = 30 tablettes), fibre 1 Gbit/s.
**Matériel :** 1× AR280, 2× S310-48P4S (un par aile), 12-14× AP361 (1 par classe + couloirs), onduleur 3000 VA.
**VLAN :** 10 (admin/profs), 20 (élèves, filtré), 40 (caméras), 70 (TBI/vidéoprojecteurs si IP).
**Commentaire :** la densité est le sujet : 30 tablettes × 12 classes qui se connectent à 8 h = **storm DHCP et pics radio**. Baux DHCP courts, band steering vers 5 GHz, et surtout **filtrage web** (USG6000F-S125 ou filtrage DNS) — une école sans filtrage, c'est un incident avec les parents assuré. Prévoir 1 AP par classe, pas 1 pour 2 : les murs de classes sont souvent en dur.

## 133. Cas pratique — Clinique 20 lits (commenté)

**Contexte :** clinique privée, 20 lits, 30 soignants, dossiers patients informatisés, Wi-Fi patients, 10 caméras, exigence de confidentialité.
**Matériel :** 1× USG6000F-S125 (sécurité : données de santé), 1× AR280 ou l'USG en tête, 2× S310-48P4S, 10× AP361, onduleur 3000 VA + groupe si possible.
**VLAN :** 10 (soignants/dossiers — chiffré, audité), 20 (patients/invités, isolé), 40 (caméras), 60 (téléphonie/DECT).
**Commentaire :** ici la **sécurité n'est pas une option** : données de santé = USG obligatoire, journaux conservés, accès tracés (registre section 106), WPA3 si les terminaux suivent. Le Wi-Fi soignants (chariots de soins, tablettes) = SSID dédié prioritaire, jamais mélangé au Wi-Fi patients. Prévoir une **procédure écrite** remise à la direction (qui accède à quoi).

## 134. Cas pratique — Agence bancaire / assurance (commenté)

**Contexte :** agence 10 personnes, VPN vers le siège, GAB (distributeur de billets) à connecter, exigence de sécurité maximale.
**Matériel :** 1× USG6000F-S125 (VPN IPsec vers siège), 1× S220-8P4S ou S310-24P, 2× AP361.
**VLAN :** 10 (postes agence), 20 (invités, isolé), 50 (GAB — **totalement isolé**, accès siège uniquement via VPN), 99 (gestion).
**Commentaire :** le GAB ne touche **jamais** le LAN de l'agence : VLAN dédié + règles pare-feu strictes + VPN vers la banque. C'est typiquement un besoin où le **cahier des charges de la banque** prime sur tes habitudes : lis-le avant de chiffrer. eKit convient pour l'agence ; le GAB suit les règles du groupe bancaire.

## 135. Cas pratique — Espace de coworking 40 postes (commenté)

**Contexte :** 40 postes en open space + 3 bureaux fermés + salle de réunion, clients qui changent chaque mois, imprimante partagée, domiciliation.
**Matériel :** 1× AR280 (ou USG si besoin), 1× S310-48P4S, 4× AP361 (densité !), onduleur 2000 VA.
**VLAN :** 10 (staff coworking), 20 (coworkers — portail captif avec code mensuel), 21 (bureaux fermés — 1 VLAN par bureau si exigé), 40 (caméras/contrôle d'accès).
**Commentaire :** le **turnover** est le sujet : PSK coworkers **changé chaque mois** (procédure section 152), portail captif avec CGU (responsabilité du gérant en cas d'abus — mentions légales !), isolation inter-clients **impérative** (un coworker ne doit pas voir le PC du voisin). La salle de réunion = AP dédié proche + priorité visio.

## 136. Cas pratique — Salle de sport / fitness (commenté)

**Contexte :** 500 m², 200 adhérents/jour, musique en streaming, écrans, contrôle d'accès par badge, 6 caméras, Wi-Fi adhérents.
**Matériel :** 1× AR180 Pro, 1× S310-24P4S, 3× AP361 (salle muscu, cardio, accueil), onduleur 1500 VA.
**VLAN :** 10 (staff/musique/caisse), 20 (adhérents, portail captif), 40 (caméras + contrôle d'accès).
**Commentaire :** la **musique** passe souvent par le réseau (streaming) : SSID ou VLAN prioritaire pour ne pas hacher quand 50 adhérents se connectent à 18 h. Humidité/chlore si piscine : AP standard **hors** zone humide, ou modèle adapté. Les vestiaires : pas d'AP dedans (confidentialité), couverture depuis le couloir.

## 137. Cas pratique — Lieu de culte / grande salle (commenté)

**Contexte :** salle de 500 places, streaming des cérémonies (caméras + diffusion), Wi-Fi public modéré, 4 caméras.
**Matériel :** 1× AR280, 1× S310-24P4S, 3-4× AP361/AP572 (densité le jour d'affluence), onduleur 2000 VA.
**VLAN :** 10 (régie/streaming — **prioritaire absolu**), 20 (public, limité), 40 (caméras).
**Commentaire :** le **streaming** est critique et gourmand (montant) : VLAN dédié, QoS, et surtout **ne jamais faire de MAJ le jour d'une cérémonie**. Le Wi-Fi public : débit bridé (5 Mbit/s), sinon 500 fidèles en direct tuent le streaming. Tester en conditions réelles (répétition) avant le jour J.

## 138. Cas pratique — Pharmacie (commenté)

**Contexte :** pharmacie 100 m², 6 employés, robot de stockage (parfois en réseau), TPE, ordonnances électroniques, 4 caméras, Wi-Fi clients discret.
**Matériel :** 1× AR180, 1× S220-8P4S, 2× AP361, onduleur 1000 VA.
**VLAN :** 10 (officine : postes, TPE, robot), 20 (clients, discret voire non diffusé en continu), 40 (caméras).
**Commentaire :** le **robot** et les **TPE** = critiques : IP fixes, VLAN 10, jamais de MAJ en journée. Les données de santé (ordonnances) imposent la même rigueur qu'une clinique (section 133) à petite échelle. Le Wi-Fi clients : utile mais secondaire — ne jamais le laisser impacter l'officine.

## 139. Cas pratique — Garage / concession automobile (commenté)

**Contexte :** showroom + atelier (poussière, métal), 15 employés, diagnostic véhicules en Wi-Fi (valises), 8 caméras, Wi-Fi clients showroom.
**Matériel :** 1× AR180 Pro, 1× S310-24P4S, 3× AP361 (showroom, atelier, bureau), onduleur 1500 VA.
**VLAN :** 10 (bureaux/valises diag), 20 (clients showroom), 40 (caméras).
**Commentaire :** l'**atelier** est un enfer radio (métal, ponts élévateurs) : survey obligatoire, AP en hauteur dégagée, et les **valises de diagnostic** = SSID dédié si elles sont capricieuses en roaming. Poussière : baie fermée, nettoyage annuel des AP. Les caméras extérieures (parc véhicules) : AP761 + caméras sur mâts avec parafoudre.

## 140. Cas pratique — ONG / bureau de terrain (commenté)

