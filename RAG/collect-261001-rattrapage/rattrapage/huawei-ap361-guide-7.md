---
id: collect-261001-rattrapage/rattrapage/huawei-ap361-guide-7
title: "Huawei eKit AP361 — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "diffusion"]
source: docs/RAG/collect-261001-rattrapage/huawei_ap361_guide.md
source_anchor: ""
source_lines: [920, 1107]
sha256: c715b9a36aa2723ff583cefcaa88ff97b3f720c6eb5be4e0e576a71ee741ac4b
---

# Huawei eKit AP361 — Guide ultra-complet

```
<AP361> system-view
[AP361] wlan
[AP361-wlan-view] ap-id 0
[AP361-wlan-ap-0] radio 0                    # radio 0 = 2,4 GHz (convention)
[AP361-wlan-radio-0/0] channel 20mhz 6
[AP361-wlan-radio-0/0] eirp 14              # puissance en dBm (indicatif)
[AP361-wlan-radio-0/0] radio enable
[AP361-wlan-radio-0/0] quit
[AP361-wlan-ap-0] radio 1                    # radio 1 = 5 GHz (convention)
[AP361-wlan-radio-0/1] channel 40mhz 44
[AP361-wlan-radio-0/1] eirp 14
[AP361-wlan-radio-0/1] radio enable
[AP361-wlan-radio-0/1] quit
[AP361-wlan-ap-0] quit
[AP361-wlan-view] quit
[AP361] save
```

Vérification :

```
<AP361> display wlan radio all
<AP361> display wlan ap all
```

## 57. Tableau récapitulatif : profil radio « PME standard »

| Paramètre | 2,4 GHz | 5 GHz |
|---|---|---|
| Canal | 1 / 6 / 11 (alternés) | 36 / 40 / 44 / 48 (ou 40 MHz : 36, 44…) |
| Largeur | 20 MHz | 40 MHz (80 si justifié) |
| Puissance initiale | 14 dBm | 14 dBm |
| Band steering | — (bande d'accueil IoT) | Actif vers 5 GHz |
| Airtime fairness | Actif | Actif |
| 802.11b (débits 1-11 Mbit/s) | **Désactivés** si possible | N/A |

> 🧹 **Désactiver les bas débits 802.11b** : ça force les vieux clients à parler
> plus vite et libère de l'airtime. Attention : certains très vieux équipements
> (ou badges) en ont besoin → testez avant en condition réelle.

## 58. Checklist radio

- [ ] Plan de canaux documenté (tableau AP × canal 2,4 × canal 5)
- [ ] Largeurs : 20 MHz en 2,4 GHz, 40 MHz en 5 GHz (sauf justification)
- [ ] Puissances réglées manuellement, notées
- [ ] Band steering actif sur SSID bureautique
- [ ] Mesure de couverture faite (-60 dBm cible, -67 dBm mini voix)
- [ ] Zones de recouvrement vérifiées entre AP adjacents
- [ ] Canaux DFS évités près des zones radar (ou assumés et documentés)

---
---

# G. SSID ET SÉCURITÉ D'ACCÈS

## 59. Créer un SSID via l'app eKit

1. `eKit > Site > Wi-Fi > + SSID`.
2. Nom (SSID) : explicite, sans accent ni espace si possible
   (`Bureau-Etage2`, pas `Bureau Étage 2 café ☕`).
3. Bande : les deux (sauf SSID IoT → 2,4 GHz only, voir §60).
4. Sécurité : voir §60-61.
5. VLAN : voir §73.
6. Options : diffusion du SSID (oui sauf cas particulier, §63), isolation client
   (§93), limite de débit (§101).
7. Appliquer → vérifier la propagation sur tous les AP (1-2 min).

**Convention de nommage conseillée** (multi-sites) :

| SSID | Usage | Exemple |
|---|---|---|
| `STE-Bureau` | Collaborateurs | Préfixe société, pas de site dedans (le même SSID partout = roaming inter-sites naturel) |
| `STE-Invite` | Visiteurs | Portail captif |
| `STE-IoT` | Objets connectés | 2,4 GHz only, WPA2-PSK, VLAN isolé |
| `STE-Voix` | Téléphones Wi-Fi (optionnel) | QoS prioritaire, 5 GHz de préférence |

## 60. WPA2-PSK vs WPA3-SAE : que choisir en 2026

| Critère | WPA2-PSK (AES) | WPA3-SAE | Mode transition WPA2/WPA3 |
|---|---|---|---|
| Sécurité | Correcte si clé robuste (attaque par dictionnaire possible) | **Supérieure** (SAE résiste aux attaques hors ligne) | Les deux proposés |
| Compatibilité | Universelle | Problèmes avec vieux clients (Windows 7, vieilles imprimantes, IoT bas de gamme) | Bonne (négociation) |
| Recommandation | SSID IoT / legacy | SSID bureautique **si le parc client est récent** | ✅ **Choix par défaut** en PME |

**Recommandation terrain** :

- **SSID bureautique** : **WPA3-SAE**, ou **transition WPA2/WPA3** si le parc est
  hétérogène. Testez avec les PC les plus vieux du parc avant de généraliser.
- **SSID invité** : WPA2-PSK suffit (rotation de clé régulière) + portail captif.
- **SSID IoT** : WPA2-PSK (beaucoup d'objets ne connaissent pas WPA3).

> 🔑 **Clé robuste** : 20+ caractères, générée aléatoirement, stockée dans le
> coffre de l'équipe. Exemple **fictif** : `Trombone-Cactus-4821-Bureau!Vert`.
> Jamais de clé = nom de société + année.

## 61. SAE (WPA3-Personal) : points d'attention

- Le handshake SAE remplace le 4-way handshake PSK : plus de capture de handshake
  exploitable hors ligne.
- **PMF (Protected Management Frames)** : requis/recommandé avec WPA3 (voir §94).
  Certains vieux clients refusent de s'associer si PMF = required → utilisez
  « optional » en transition.
- Si un client **ne s'associe plus** après passage en WPA3 pur : repassez en
  transition et identifiez le client fautif (voir §123).

## 62. 802.1X / EAP : l'authentification entreprise

Pour un vrai contrôle d'accès (un identifiant par utilisateur, révocable), il faut
passer au **WPA2/WPA3-Enterprise (802.1X)** :

```
Client ---(EAP)---> AP361 ---(RADIUS)---> Serveur RADIUS (ex: NPS Windows, FreeRADIUS)
```

| Méthode EAP | Certificat serveur | Certificat client | Usage |
|---|---|---|---|
| **PEAP-MSCHAPv2** | Oui (requis) | Non | Le plus courant en PME (login/mot de passe AD) |
| **EAP-TLS** | Oui | **Oui** | Le plus sûr (PKI), plus lourd à déployer |
| EAP-TTLS | Oui | Non | Alternative à PEAP |

**Prérequis** :

- Serveur RADIUS joignable depuis les AP (ports UDP 1812/1813).
- Certificat serveur **valide** (sinon les clients affichent des alertes et les
  utilisateurs cliquent « accepter » → faille homme-du-milieu).
- 🔎 Vérifiez que votre version eKit expose la configuration RADIUS pour
  l'AP361 (sinon : SSID PSK + portail, ou gamme supérieure).

**En PME sans AD** : un **FreeRADIUS** sur un petit Linux fait le job pour
20-50 utilisateurs. Mais ça ajoute une brique à maintenir : pesez le pour et le
contre face à un WPA3-SAE avec clé robuste + rotation trimestrielle.

## 63. Portail captif (captive portal) : pour les invités

Le portail captif redirige le visiteur vers une page d'authentification
(CGU, code, identifiant/mot de passe temporaire).

**Bonnes pratiques** :

- SSID **dédié** (`STE-Invite`), VLAN **isolé** (voir §76), débit **limité**
  (ex. 10/10 Mbit/s par client).
- **Isolation client** activée : les invités ne se voient pas entre eux.
- Page de CGU courte, en français, avec durée de validité du ticket.
- **Ne jamais** mettre le portail captif comme seule sécurité du SSID
  bureautique : le trafic avant authentification n'est pas chiffré.

🔎 Les capacités de portail (interne à l'AP vs externe, personnalisation) dépendent
de la version eKit : vérifiez dans l'app.

## 64. SSID invité : le kit complet

- [ ] SSID dédié, WPA2-PSK (clé simple, affichée à l'accueil, **changée
      mensuellement**)
- [ ] Portail captif avec CGU
- [ ] VLAN invité dédié, **sans accès** au LAN (ACL ou firewall, voir §76)
- [ ] Isolation client activée
- [ ] Limite de débit par client (ex. 10 Mbit/s)
- [ ] Bande : les deux (ou 5 GHz only si zone dense)

## 65. SSID caché (non diffusé) : pour ou contre ?

**Contre** (position recommandée) :

- Ça n'apporte **aucune sécurité réelle** (le SSID se découvre en 30 secondes
  avec un sniffer).
- Ça complique la vie des utilisateurs (saisie manuelle) et **casse** le band
  steering et certains mécanismes de roaming.
- Les clients configurés en « SSID caché » **sondent en permanence** à la
  recherche du réseau → pollution radio + batterie.

**Pour** (seul cas acceptable) : SSID technique temporaire, ou contrainte
contractuelle explicite. Sinon : **diffusez vos SSID**.

## 66. Combien de SSID ? La règle des 3-4

Chaque SSID consomme de l'**airtime** (beacons ~10 fois/seconde par SSID et par
bande). Au-delà de 4-5 SSID par AP, la surcharge devient mesurable.

| Nb de SSID | Verdict |
|---|---|
| 1-2 | Idéal |
| 3-4 | ✅ Standard PME (bureau, invité, IoT, +1) |
| 5-6 | Limite haute, à justifier |
| 7+ | ⛔ Non : fusionnez ou utilisez le VLAN par d'autres moyens |

## 67. Exemples CLI — SSID et sécurité (indicatif VRP)

