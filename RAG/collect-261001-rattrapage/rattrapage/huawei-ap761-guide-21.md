---
id: collect-261001-rattrapage/rattrapage/huawei-ap761-guide-21
title: "Guide ultra-complet — Huawei eKit AP761"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "incident"]
source: docs/RAG/collect-261001-rattrapage/huawei_ap761_guide.md
source_anchor: ""
source_lines: [2087, 2175]
sha256: 208fbddc899cd73de6ae4d7a38de0239b91f37cc09e49dc8530bc448779b9a78
---

# Guide ultra-complet — Huawei eKit AP761

**Remèdes :**
1. **Ne jamais enfermer** l'AP dans un coffret étanche non ventilé (chap. 5).
2. **Pare-soleil** : un simple auvent/champagneau au-dessus de l'AP (sans bloquer la ventilation ni le secteur radio) fait chuter la température de 10–15 °C.
3. **Orientation** : éviter le plein sud direct si possible ; un mur à l'est ou à l'ouest prend le soleil moins longtemps.
4. Vérifier que les **ouïes de ventilation** (s'il y en a — à vérifier sur la fiche du modèle exact) ne sont pas obstruées (nids, toiles, poussière).
5. Si le site dépasse durablement les specs : **ombrage obligatoire** ou déplacement — faire fonctionner un AP hors spec, c'est accepter des pannes et perdre la garantie.

**En hiver, l'inverse :** −40 °C supportés [constructeur], mais attention à la **condensation** lors des cycles gel/dégel — presse-étoupes vers le bas, pas de stagnation d'eau (chap. 5).

## 92. Cas 15 : rogue AP détecté — que faire concrètement

**Symptômes :** alerte WIDS « rogue AP détecté », un SSID suspect apparaît dans les scans.

**Procédure (ne pas improviser) :**
1. **Ne pas contre-attaquer immédiatement** (chap. 63–64) : classifier d'abord.
2. **Identifier :** noter BSSID (MAC), SSID, canal, RSSI vu par tes AP (triangulation approximative : l'AP qui le voit le plus fort est le plus proche).
3. **Est-il sur ton filaire ?** Chercher sa MAC dans les tables MAC des switches.
   - **Oui → interne :** remonter au port, aller voir physiquement (chap. 64). 90 % des cas = un équipement personnel.
   - **Non → externe :** usurpe-t-il ton SSID (evil twin) ou est-ce un voisin légitime ?
4. **Evil twin confirmé** (même SSID que toi, fort près de tes locaux) : activer la **contre-mesure ciblée** (désauthentification du rogue — fonction wIPS, à vérifier selon version), **alerter la direction**, conserver les logs (preuve), envisager un dépôt de plainte si malveillance avérée.
5. **Voisin légitime** (SSID différent, pas sur ton filaire) : **ne rien faire** d'agressif. Optimiser tes canaux pour coexister (chap. 37–38).

**Après incident :** noter dans le dossier de site (date, MAC, action). Si c'était un salarié : **sensibilisation** — et se demander pourquoi il a branché un AP pirate (chap. 64).

## 93. Cas 16 : après mise à jour firmware, comportements bizarres

**Symptômes :** depuis la mise à jour : clients qui ne s'associent plus, débit en berne, redémarrages, fonctions qui ont disparu des menus.

**Conduite à tenir :**
1. **Ne pas paniquer, mesurer** : comparer les indicateurs (chap. 67) avant/après. « Bizarre » n'est pas un diagnostic.
2. **Lire la release note** : le comportement a peut-être **changé volontairement** (nouveau défaut de canal, PMF activé par défaut, nom de commande modifié).
3. **Vérifier la config migrée** : `display current-configuration` — une migration peut avoir réinitialisé des paramètres (canaux repassés en auto, puissances par défaut).
4. **Si régression avérée → rollback** (chap. 76) : re-flasher l'ancienne version + restaurer la sauvegarde pré-mise à jour.
5. **Remonter au support** avec : versions avant/après, logs, description reproductible.

**Prévention (rappel chap. 74) :** site pilote d'abord, jamais de mise à jour massive un vendredi à 17h, sauvegarde avant, rollback testé. **Le firmware n'est pas un antivirus : on ne met à jour que pour une raison** (faille, bug qui te touche, fonction nécessaire).

## 94. Cas 17 : itinérance entre AP761 et AP indoor

**Symptômes :** en passant de la cour (AP761) au bâtiment (AP361 indoor), les clients coupent ou restent accrochés à l'AP extérieur.

**C'est le cas mixte classique :** deux modèles, deux environnements, un seul parcours utilisateur.

**Points de vigilance :**
1. **Même SSID, même sécurité, mêmes VLAN** des deux côtés — sinon le client change de réseau en passant la porte.
2. **Puissances équilibrées à la frontière** : l'AP761 en extérieur a tendance à « crier » plus fort que l'AP361 indoor. Résultat : le client reste accroché à l'AP761 depuis l'intérieur (à travers le mur, en 2.4 GHz) avec un débit minable. **Baisser la puissance 2.4 GHz de l'AP761** côté bâtiment, ou orienter son secteur **loin** de la façade.
3. **Zone de transition** : il faut un **recouvrement** à −67 dBm juste devant l'entrée — ni trou (coupure), ni recouvrement énorme (client sticky).
4. **Bande 2.4 GHz traîtresse** : elle traverse les murs, pas le 5 GHz. Un client qui entre dans le bâtiment reste en 2.4 GHz sur l'AP761 extérieur. **Débits de base relevés** (chap. 13) pour le forcer à lâcher.
5. **Tester le parcours réel** : sortir avec un appel en cours, entrer, noter où ça coupe (chap. 57).

**Astuce :** si la zone d'entrée est critique (accueil, contrôle d'accès), dédier un **petit AP indoor près de l'entrée** plutôt que de compter sur l'AP761 extérieur pour couvrir l'intérieur — ce n'est pas son travail (chap. 9 : l'arrière est sourd).

## 95. Cas 18 : je veux du Wi-Fi 7 — que faire avec mes AP761 ?

**Question légitime** quand la direction lit « Wi-Fi 7 » dans la presse. Réponse structurée en 3 temps :

**Temps 1 — Cadrer le besoin réel :**
- Quel problème le Wi-Fi 7 résoudrait-il **que le Wi-Fi 6 ne résout pas** ? (Débit par client > 500 Mbps ? Latence < 5 ms ? Densité > 100 clients/AP ?)
- Les **clients** sont-ils Wi-Fi 7 ? (En 2026 : une minorité. Un AP Wi-Fi 7 avec des clients Wi-Fi 6 = un AP Wi-Fi 6 cher.)
- Dans 80 % des cas PME, la réponse est : **le Wi-Fi 6 bien déployé suffit** — et l'argent est mieux investi dans le câblage, le PoE et la couverture.

**Temps 2 — Si le besoin est avéré :**
- **Ne pas jeter les AP761** : ils restent excellents pour la couverture générale. Ajouter des **AP Wi-Fi 7 (AP772E)** sur les **zones denses** (chap. 21, 99–100).
- **Refaire le budget** : uplink 2.5G/10G, PoE 802.3bt, câblage Cat6A — le Wi-Fi 7 ne vit pas sur une infra Wi-Fi 6 (chap. 25).
- **Cohabitation** : même SSID sur les deux générations, transition WPA2/WPA3 (chap. 47), canaux coordonnés.

**Temps 3 — Calendrier réaliste :**
- 2026–2027 : les clients Wi-Fi 7 deviennent majoritaires dans le neuf.
- Un AP761 bien installé aujourd'hui a encore **5–7 ans** de vie utile devant lui (chap. 105).
- **Ne pas acheter du Wi-Fi 7 « pour le futur »** en 2026 si le besoin n'existe pas : dans 3 ans, les produits seront meilleurs et moins chers.

## 96. Comparatif AP761 vs AP361 — tableau décisionnel

Les deux sont des **Wi-Fi 6 1.775 Gbps 2×2** [constructeur] — la différence n'est pas le Wi-Fi, c'est **l'enveloppe et l'usage**.

| Critère | **AP761** (outdoor) | **AP361** (indoor) |
|---|---|---|
| Usage prévu | Extérieur : cours, parkings, terrasses | Intérieur : bureaux, commerces, salles |
| Standard | Wi-Fi 6, 2×2, 1.775 Gbps | Wi-Fi 6, 2×2, 1.775 Gbps |
| Ports | 1× GE PoE-In + 1× SFP GE (combo) | 1× GE PoE |
| PoE | 802.3at/af, **17.7 W max** | 802.3at/af (à vérifier), **8.8 W max** [constructeur] |
| Antennes | **Directionnelles** 10/11 dBi, 65° | **Smart antennas** (omnidirectionnelles adaptatives) |
| Protection | **IP68**, −40 à +65 °C, 6 kV | Intérieur (pas d'IP68) |
| Montage | Mur / mât | Plafond / mur |
| Dimensions / poids | 200×200×69 mm, **1.91 kg** | 180×35 mm (disque), léger |
| Clients max | 1024 (512/radio) | 300 (recommandés, selon manuel eKit) |
| Prix indicatif | ~200–245 € HT (discomp.cz, 2026) | ~70–85 € HT (discomp.cz, 2026) |
| BLE | 5.2 | À vérifier sur la fiche du modèle exact |

