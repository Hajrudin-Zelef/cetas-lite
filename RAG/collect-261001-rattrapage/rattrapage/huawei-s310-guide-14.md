---
id: collect-261001-rattrapage/rattrapage/huawei-s310-guide-14
title: "Guide ultra-complet — Huawei eKit S310"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-rattrapage/huawei_s310_guide.md
source_anchor: ""
source_lines: [2843, 3040]
sha256: 68d14e91f4388f147342c60680aef763e04f3ddf7ac30d89ae2c7df58b74ba1a
---

# Guide ultra-complet — Huawei eKit S310

**Diagnostic (dans l'ordre) :**
1. `display poe interface <port>` : PoE activé ? Budget restant ?
   (`display poe power-state`).
2. `display poe interface` : quelle **classe** détectée ? Classe 0 inattendue =
   pas de négociation.
3. Teste l'équipement sur un **autre port** PoE.
4. Teste avec une **jarretière courte** (câble trop long ou paires coupées).
5. `display logbuffer | include POE` : cause (overload, short, etc.).

**Solution :** selon le diagnostic — réactiver `poe enable`, libérer du budget
(reprioriser, section 61), changer le câble, ou remplacer l'équipement.
Si la caméra **démarre puis s'éteint** : limite `poe power` trop basse ou budget
tout juste (pics IR la nuit) → augmente la marge.

---

## 102. Cas n°4 — Port en err-disable / coupé par une protection

**Symptômes :** un port (ou plusieurs) passe down sans raison physique évidente ;
les logs parlent de BPDU, de storm, de port-security.

**Diagnostic :** `display logbuffer` → la **cause exacte** (BPDU protection,
storm-control action shutdown, port-security violation...).

**Solution :**
1. **Retire la cause** : switch sauvage débranché, équipement qui spamme
   déconnecté, MAC légitime ajoutée au sticky...
2. Réarme le port :
   ```
   interface GigabitEthernet0/0/5
    shutdown
    undo shutdown
   ```
3. Si ça recommence immédiatement : la cause est toujours là. Ne force pas en
   désactivant la protection « pour voir » — sauf en dernier recours documenté.

✅ **Ne désactive jamais une protection pour masquer une panne.** Elle a coupé
pour une raison ; trouve-la.

---

## 103. Cas n°5 — DHCP snooping bloque tout le monde

**Symptômes :** après activation du DHCP snooping, **personne** n'obtient d'IP
(ou un seul VLAN est touché).

**Diagnostic :**
1. `display dhcp snooping user-bind all` → vide ou incomplet ?
2. Le port vers le **serveur DHCP** est-il en `dhcp snooping trusted` ?
   **Sur chaque switch traversé** (cascade !).
3. Le snooping est-il activé **dans le bon VLAN** ?

**Solution :**
```
interface GigabitEthernet0/0/28   # vers le serveur DHCP
 dhcp snooping trusted
```
+ vérifie la chaîne complète. Teste l'obtention d'un bail **par VLAN** après
chaque correction.

⚠️ **En cascade :** le DHCP traverse SW-2 puis SW-1 : les **deux** uplinks
doivent être trusted. C'est l'oubli classique.

---

## 104. Cas n°6 — Débit de l'uplink saturé

**Symptômes :** lenteurs généralisées, surtout à certaines heures ; tout est
« up » mais tout rame.

**Diagnostic :**
1. `display interface <uplink>` → utilisation proche de 100 % en continu ?
2. `display port statistics` / compteurs : quel port génère le trafic ?
   (Miroir + Wireshark, section 85, pour identifier : sauvegarde cloud ? Caméra
   en haute résolution ? Boucle partielle ?)
3. Heure de pointe corrélée à une activité (sauvegardes à 14h ?).

**Solution :**
- Court terme : déplace les sauvegardes en heure creuse, baisse la résolution
  des caméras si inutilement haute.
- Structurel : passe l'uplink en **10G** (modèle X, section 5) ou en
  **Eth-Trunk** (section 55).
- Vérifie qu'il n'y a pas de **boucle partielle** (storm control, section 73).

---

## 105. Cas n°7 — Plus d'accès web au switch

**Symptômes :** `https://<ip>` ne répond plus (timeout), alors que le ping passe
(ou pas).

**Diagnostic :**
1. Le ping passe ? Non → problème IP/routage/VLAN (revois sections 20, 36).
   Oui → le serveur web est en cause.
2. En console/SSH : `display http server` (ou équivalent) → HTTP/HTTPS activés ?
3. Une **ACL** (section 78) bloque-t-elle ton poste ?
4. Certificat expiré ou navigateur qui bloque ? Teste un autre navigateur / en
   navigation privée.

**Solution :** réactive (`http secure-server enable`), corrige l'ACL, ou
repasse par la console pour diagnostiquer. **Toujours garder l'accès console
possible** : c'est la porte de secours quand le réseau est en vrac.

---

## 106. Cas n°8 — Mot de passe admin oublié

**Symptômes :** impossible de se connecter (console/web/SSH), personne n'a le
mot de passe.

**Diagnostic :** c'est bien un oubli, pas un piratage ? (Vérifie avec le
responsable avant toute manip.)

**Solution :**
- Via la console, procédure de **récupération** (mode BootROM / reset du mot de
  passe) : la procédure exacte est **à vérifier sur le guide du modèle exact**
  (elle varie selon les versions et peut nécessiter un redémarrage).
- En dernier recours : **reset usine** (section 93) + restauration de la
  sauvegarde (section 89) — d'où l'importance des sauvegardes !
- Après : mot de passe au **coffre** (section 117), procédure écrite.

✅ **Prévention :** 2 comptes admin nominatifs (pas de compte partagé), mots de
passe au coffre d'entreprise, test de récupération une fois par an.

---

## 107. Cas n°9 — Le stack ne se forme pas

**Symptômes :** le 2e switch reste indépendant, ou le stack se forme puis casse.

**Diagnostic :** `display stack` / `display stack topology` sur les deux.
1. **Versions logicielles identiques ?** Non → c'est (presque) toujours ça.
2. Câbles de stack bien branchés sur les **bons ports** (dédiés ou déclarés) ?
3. Priorités cohérentes (un seul master) ?
4. Le 2e switch a-t-il été **éteint puis rallumé après** câblage ?

**Solution :** aligne les versions **avant** (section 91), vérifie le câblage,
respecte l'ordre (master d'abord). Voir le tableau des pannes section 96.

---

## 108. Cas n°10 — Module SFP non reconnu / lien fibre down

**Symptômes :** pas de lien sur l'uplink fibre, LED SFP éteinte, ou lien qui
tombe par intermittence.

**Diagnostic :**
1. `display interface` : le module est-il **détecté** (transceiver info) ?
   Non → compatibilité (module non codé Huawei / tiers bas de gamme).
2. `display transceiver` : niveaux **Tx/Rx** (dBm). Rx trop faible (< -20 dBm
   typique) → fibre sale, cassée, ou trop longue pour le module.
3. Les **deux extrémités** ont-elles le même type (SX↔SX, LX↔LX) ? Un SX face à
   un LX ne marchera jamais.
4. Fibres **croisées** (Tx→Rx) ? Avec des jarretières duplex, vérifie le sens.

**Solution :** nettoie les connecteurs (stylo fibre), remplace la jarretière,
change le module pour un compatible, vérifie le budget optique. **N'achète que
des modules compatibles vérifiés** (section 16).

---

## 109. Cas n°11 — Débit anormal sur un port cuivre (100M au lieu de 1G)

**Symptômes :** un poste rame, `display interface` montre 100Mbit/s négociés.

**Diagnostic :**
1. **Câble** : 4 paires nécessaires pour le gigabit. Un câble abîmé (paires
   4-5/7-8 coupées) tombe en 100M. → VCT (section 86).
2. Négociation **forcée** d'un côté (100M full forcé côté PC, auto côté switch)
   → duplex mismatch possible. **Remets l'auto-négociation des deux côtés.**
3. Jarretière ou prise murale défectueuse → teste en direct au switch.

**Solution :** remplace le câble/jarretière, repasse en auto. **Ne force jamais
la vitesse/duplex** sauf équipement industriel qui l'exige (et alors des deux
côtés, documenté).

---

## 110. Cas n°12 — Téléphone IP : pas de voix / pas de VLAN voix

**Symptômes :** le téléphone s'allume (PoE OK) mais n'obtient pas d'IP voix, ou
reste dans le VLAN data.

**Diagnostic :**
1. `display voice-vlan status` : voice VLAN actif sur le port ?
2. `display mac-address vlan 20` : la MAC du téléphone y apparaît-elle ?
   Non → **OUI MAC** du constructeur non déclaré (section 38) : récupère les
   3 premiers octets de la MAC (`display mac-address interface ...`) et
   ajoute l'OUI.
3. Le téléphone est-il en mode VLAN **auto/LLDP-MED** ? (Menu du téléphone —
   à vérifier sur le modèle exact.)
4. Port en **hybrid** avec `tagged vlan 20` ? (Pas en access !)

**Solution :** ajoute l'OUI, vérifie le mode hybrid, teste avec un téléphone
dont l'OUI est déjà déclaré pour isoler (téléphone vs switch).

---

