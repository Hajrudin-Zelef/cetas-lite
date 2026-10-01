---
id: collect-261001-rattrapage/rattrapage/huawei-ap361-guide-16
title: "Huawei eKit AP361 — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: ["2026-09-27"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ap361_guide.md
source_anchor: ""
source_lines: [2520, 2583]
sha256: cc4f9b9c9fd46c8a5b68ffcc81f6d56579d72cdadfa3133eb09487585b6eba57
---

# Huawei eKit AP361 — Guide ultra-complet

**R8.** (1) SSID non appliqué à cet AP (oubli de template) ; (2) radio
désactivée ou en panne ; (3) canal DFS en attente de détection radar ;
(4) puissance trop faible ; (5) clients qui s'associent puis repartent faute de
DHCP sur le VLAN. (4 réponses suffisent.)

**R9.** **Ne jamais couper l'alimentation pendant l'écriture flash.**
Procédure : lire les release notes → sauvegarder la config → fenêtre hors
heures ouvrées → **1 AP pilote**, validation 24 h → généralisation **par
vagues** (jamais tout d'un coup) → vérifier l'homogénéité des versions →
conserver l'image N-1 deux semaines.

**R10.** 12 AP × 10 W = 120 W ; 4 caméras × 12 W = 48 W ; total = 168 W ;
× 1,2 = **201,6 W > 185 W** → **non**, ça ne passe pas avec la marge.
Solutions : switch à budget supérieur (250-370 W), ou répartir sur deux
switchs, ou accepter de fonctionner sans marge (déconseillé).

---
---

# T. POUR ALLER PLUS LOIN

## 160. Pour aller plus loin

**Documentation officielle (à récupérer)** :

- Datasheet officielle AP361 — ekit.huawei.com (vérifiez la révision hardware
  de votre lot)
- *WLAN Hardware Installation and Maintenance Guide* — chapitre AP361 (LED,
  reset, specs détaillées, couples de serrage)
- *Configuration Guide* eKit de votre version logicielle + release notes
- Guide de démarrage rapide fourni dans le carton (QR code)

**Dans votre bibliothèque (déjà livrés)** :

- `onduleurs_ups_guide.md` — dimensionner l'onduleur de la baie PoE (§23)
- `proxmox_guide.md` — héberger le serveur syslog / RADIUS / Zabbix en VM
- `zabbix_guide.md` — template de supervision §109
- `debian_ubuntu_guide.md` — serveur RADIUS (FreeRADIUS), syslog, scripts
- `ssh_guide.md` — durcir l'accès SSH §145
- `loki_guide.md` / `prometheus_guide.md` / `grafana_guide.md` — supervision avancée

**Sujets d'approfondissement** :

1. **Étude de couverture (site survey)** : avec un logiciel de survey
   (Ekahau, NetSpot…), passez du « au pif » au dimensionné — indispensable
   au-delà de 20 AP ou en environnement contraint (hôpital, entrepôt).
2. **802.1X + PKI interne** : EAP-TLS avec votre CA (voir guide Debian/Ubuntu)
   pour un contrôle d'accès sans mot de passe partagé.
3. **NAC** : coupler le RADIUS à une politique d'accès dynamique (VLAN
   assigné par utilisateur).
4. **Wi-Fi 6E / Wi-Fi 7** : quand la densité ou les usages (AR/VR, vidéo 4K
   massive) l'exigeront — l'AP361 est Wi-Fi 6 2,4/5 GHz, pas 6 GHz.
5. **Analyse spectrale** : un analyseur de spectre (même USB à ~300 €) pour
   traquer les interférences non-Wi-Fi que l'analyseur logiciel ne voit pas.

**Dernier mot terrain** : un réseau Wi-Fi, c'est 20 % d'achat et 80 % de
méthode — plan radio écrit, étiquetage, mesures, documentation, maintenance.
L'AP361 fait correctement son travail si **vous** faites correctement le vôtre.
Bon chantier. 🔧

---

*Fin du guide — Huawei eKit AP361 — rédigé le 2026-09-27 pour Zelef.*
*Vérifiez les points marqués 🔎 sur la fiche du modèle exact avant mise en production.*
