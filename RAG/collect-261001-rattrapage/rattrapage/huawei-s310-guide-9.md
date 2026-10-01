---
id: collect-261001-rattrapage/rattrapage/huawei-s310-guide-9
title: "Guide ultra-complet — Huawei eKit S310"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-rattrapage/huawei_s310_guide.md
source_anchor: ""
source_lines: [1705, 1929]
sha256: 44158ee488bc8700839b8fe456044bf0333bfe5713b13a03c9952215f45073bb
---

# Guide ultra-complet — Huawei eKit S310

⚠️ **Réfléchis avant de mettre `critical` partout :** si tout est critique,
rien ne l'est, et le switch coupera au hasard (en pratique : par numéro de port).

---

## 62. Planification PoE : la méthode (à faire AVANT d'acheter)

1. **Liste** chaque équipement PoE prévu : type, quantité, puissance max
   **côté switch** (fiche constructeur de l'équipement, pas « à peu près »).
2. **Additionne** : total = Σ puissances.
3. **Ajoute 20 % de marge** (pics, vieillissement, un équipement ajouté « vite fait »).
4. **Compare** au budget du modèle : total × 1,2 ≤ budget → OK.
5. **Vérifie par port** : aucun équipement > 30 W (sinon : injecteur 60/90 W
   externe ou autre gamme).
6. **Attribue les priorités** (section 61).

**Template de tableau de planification :**

| Port | Équipement | Puissance (W) | Priorité | Commentaire |
|---|---|---|---|---|
| GE0/0/1 | AP361 hall | 8,8 | high | |
| GE0/0/2 | Caméra PTZ parking | 25 | low | Classe 4 |
| ... | ... | ... | ... | |
| **Total** | | **... W** | | Marge : budget − total |

---

## 63. Calcul : combien d'AP Huawei eKit alimentables ?

**Données constructeur vérifiées (recherche web, septembre 2026) :**
- **AP361** (Wi-Fi 6 intérieur) : **8,8 W** max, PoE 802.3af.
- **AP761** (Wi-Fi 6 extérieur) : **17,7 W** max, PoE 802.3at recommandé
  (en 802.3af, fonctions limitées).

**Sur S310-24P4S (budget 400 W, 24 ports) :**

| Équipement | Calcul | Résultat |
|---|---|---|
| AP361 (8,8 W) | 400 ÷ 8,8 = 45,4 | **Limité par les ports : 24 AP** (211 W consommés, marge énorme) |
| AP761 (17,7 W) | 400 ÷ 17,7 = 22,6 | **22 AP max** (389 W) — le 23e ne montera pas |
| Mixte : 12× AP361 + 6× AP761 | 12×8,8 + 6×17,7 = 211,8 W | ✅ Largement OK |

**Sur S310-48P4S (budget 380 W, 48 ports) :**

| Équipement | Calcul | Résultat |
|---|---|---|
| AP361 (8,8 W) | 380 ÷ 8,8 = 43,1 | **43 AP max** (378 W) — pas 48 ! |
| AP761 (17,7 W) | 380 ÷ 17,7 = 21,4 | **21 AP max** (372 W) |
| Mixte : 24× AP361 + 10 caméras 15 W | 24×8,8 + 10×15 = 361 W | ✅ OK (19 W de marge — juste, prévoir priorités) |

⚠️ **Le 48P4S ne peut pas alimenter 48 AP761** : c'est le piège classique du
« 48 ports PoE ». Le nombre de ports n'est **pas** le nombre d'équipements
alimentables. Fais toujours le calcul.

🔧 En exploitation : `display poe power-state` régulièrement, et alerte si le
disponible passe sous 20 % du budget.

---

## 64. Dépannage PoE : arbre de décision

**Symptôme : l'équipement ne s'allume pas.**

1. `display poe interface GigabitEthernet0/0/X` → le port fournit-il ?
   - **Non / « off »** : vérifie `poe enable` sur le port, puis le **budget
     restant** (`display poe power-state`). Budget épuisé → débranche ou
     repriorise (section 61).
   - **Oui mais l'équipement reste éteint** : passe à 2.
2. **Câble** : teste avec une jarretière courte connue bonne. Les paires
   utilisées par le PoE (1-2/3-6 en mode A, 4-5/7-8 en mode B) peuvent être
   coupées alors que les données passent (cas tordu mais réel).
3. **Distance** : > 100 m = PoE non garanti. Mesure la longueur (VCT, section 96).
4. **Classe** : `display poe interface` montre la classe détectée. Classe 0
   inattendue = l'équipement ne négocie pas → teste avec un injecteur pour
   isoler (équipement HS vs switch).
5. **Port** : essaie le même équipement sur un autre port PoE. Si ça marche
   ailleurs → port suspect (surtension antérieure ?).
6. **Budget par port** : une limite `poe power` trop basse (section 60) empêche
   la montée en puissance (caméra PTZ qui démarre puis s'éteint = typique).

**Symptôme : l'équipement s'éteint par intermittence.**

- Budget **tout juste** suffisant : un pic (démarrage IR de la caméra la nuit)
  fait dépasser → le switch coupe le port le moins prioritaire. Augmente la
  marge ou change les priorités.
- `display logbuffer | include POE` : les coupures sont loggées avec la cause
  (overload, short-circuit...).

✅ **Bon réflexe :** devant tout équipement PoE « bizarre », commence par
`poe power-off` / `poe power-on` à distance. Ça règle 50 % des cas (équipement
planté) en 30 secondes.

---

## 65. QoS : rappels (le minimum pour ne pas tout casser)

- La QoS ne **crée** pas de bande passante : elle **choisit qui passe en premier**
  quand il y a congestion.
- Deux marquages : **802.1p** (priorité 0–7 dans le tag VLAN, niveau 2) et
  **DSCP** (6 bits dans l'en-tête IP, niveau 3).
- Le switch classe le trafic, le met dans des **files d'attente**, et les vide
  selon un **ordonnancement** (strict priority, WRR...).
- **Règle d'or :** marque **au plus près de la source** (le téléphone marque sa
  voix), et fais **confiance** (`trust`) à ce marquage sur le switch. Marquer
  partout « au cas où » = marquage incohérent = QoS inutile.

**Valeurs à connaître :**

| Trafic | DSCP | 802.1p | Nom |
|---|---|---|---|
| Voix (RTP) | EF (46) | 5 | Priorité max |
| Signalisation voix (SIP) | CS3 (24) | 3 | Haute |
| Visio interactive | AF41 (34) | 4 | Haute |
| Données critiques | AF31 (26) | 3 | Moyenne-haute |
| Best effort (défaut) | 0 | 0 | — |

---

## 66. Trust DSCP vs trust 802.1p : configuration

**Faire confiance au DSCP (recommandé quand les équipements marquent bien) :**

```
system-view
interface GigabitEthernet0/0/8
 trust dscp
quit
save
```

**Faire confiance au 802.1p (utile en L2 pur, sans IP) :**

```
system-view
interface GigabitEthernet0/0/8
 trust 8021p
quit
save
```

**Forcer une priorité (quand l'équipement ne marque pas, ex. vieux téléphone) :**

```
system-view
interface GigabitEthernet0/0/9
 port priority 5        # tout ce qui entre par ce port = priorité 5
quit
save
```

> La disponibilité exacte de `trust dscp` / `trust 8021p` en vue interface est
> **à vérifier sur la version logicielle du modèle exact**. Sur certaines
> versions eKit, la QoS se configure via des **traffic classifiers/behaviors/
> policies** (modèle MQC, voir section 67).

---

## 67. Files, ordonnancement et MQC (le modèle Huawei)

Sur VRP, la QoS avancée suit le modèle **MQC** (Modular QoS Command-line) en
3 étapes : **classifier** (qui ?) → **behavior** (quoi faire ?) → **policy**
(appliquée où ?).

```
system-view
# 1. Classifier : le trafic voix (DSCP EF)
traffic classifier VOICE operator or
 if-match dscp ef
quit
# 2. Behavior : priorité stricte
traffic behavior VOICE_PRIO
 queue af bandwidth pct 30        # exemple : file avec 30 % de bande garantie
quit
# 3. Policy : on lie les deux et on l'applique sur l'uplink
traffic policy QOS_LAN
 classifier VOICE behavior VOICE_PRIO
quit
interface GigabitEthernet0/0/28
 traffic-policy QOS_LAN outbound
quit
save
```

> Les paramètres exacts de files (`queue af`, `pq`, `wrr`, `shaping`) et leur
> disponibilité sont **à vérifier sur la version logicielle du modèle exact**.
> En PME, le **trust DSCP + priorité stricte pour la voix** (section 68) couvre
> 95 % des besoins sans MQC complet.

---

## 68. Exemple : prioriser la voix sur tout le switch

Objectif : la voix (VLAN 20, DSCP EF) passe toujours en premier sur l'uplink,
même quand un utilisateur sature le lien avec un gros transfert.

```
system-view
# Sur chaque port téléphone : confiance au marquage du poste
port-group pg-voice
 group-member GigabitEthernet0/0/1 to GigabitEthernet0/0/20
 trust dscp
quit
# Sur l'uplink : la file voix est servie en priorité stricte
# (via MQC selon version, ou trust dscp + ordonnancement par défaut)
interface GigabitEthernet0/0/28
 trust dscp
quit
save
```

🔧 Vérifications :

```
display qos interface GigabitEthernet0/0/28
display traffic-policy applied-record
```

**Test de validation :** lance un gros transfert (iperf) entre deux PC pendant
un appel : l'appel ne doit pas se dégrader. Si ça craque, la QoS n'est pas
appliquée au bon endroit (souvent : oubliée sur l'uplink).

---

## 69. Exemple : prioriser la visio (Teams/Zoom)

