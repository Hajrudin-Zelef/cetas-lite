---
id: collect-261001-rattrapage/rattrapage/huawei-ar720-guide-12
title: "Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/huawei_ar720_guide.md
source_anchor: ""
source_lines: [2080, 2250]
sha256: bdd39187e94dc52fabb205825b82c4a3c46b4388b18ad98bde7a289d36125838
---

# Guide Huawei NetEngine AR720 — Routeurs d'entreprise PME/Agences

Cocher, dater, archiver. Une checklist non archivée = une checklist non faite.

## 108. Dépoussiérage et environnement

- Le flux d'air va de gauche à droite : dépoussiérer les grilles d'entrée (gauche) au
  moins une fois par trimestre dans un local poussiéreux.
- Ne jamais poser d'objets sur le routeur (il dissipe par le châssis aussi).
- Contrôler que le local ne dépasse pas 45 °C en pointe (été, clim en panne).
- Vérifier le serrage des cordons d'alimentation et l'état des câbles RJ45 (languettes
  cassées = faux contacts intermittents = tickets fantômes).

## 109. Gestion des pièces de rechange

Pour un parc d'agences : **1 routeur de spare pour ~10 sites** (minimum 1 au total).
Le spare est pré-configuré avec un template, stocké avec sa console-câble et une clé USB
contenant les firmwares validés. Procédure de remplacement : noter le numéro de série,
restaurer la config du site (adapter hostname/IP), tester WAN/LAN/VPN, `save`.

## 110. Suivi des firmwares et des failles

- S'abonner aux bulletins de sécurité Huawei pour la gamme AR.
- Ne pas appliquer chaque version le jour de sa sortie : attendre les retours, tester en
  labo/spare, puis déployer par vagues (1 site pilote → 30 % → 100 %).
- Tenir un registre : site / version actuelle / version cible / date de passage.

---
---

# Partie M — Durcissement

## 111. Principe : réduire la surface d'attaque

Un routeur de branche exposé sur Internet est scanné en permanence. Le durcissement =
désactiver tout ce qui ne sert pas, restreindre tout ce qui sert, journaliser le reste.

## 112. Désactiver les services inutiles

```
[AGENCE-DAKAR-AR720]undo telnet server enable
[AGENCE-DAKAR-AR720]undo ftp server enable
[AGENCE-DAKAR-AR720]undo http server enable
```

Ne garder que : SSH (admin), SNMP (supervision, v3), NTP (client). Le serveur web
d'administration : à n'activer que si l'équipe l'utilise vraiment, et jamais exposé sur
le WAN (politique untrust→local : refuser les ports 80/443).

## 113. SSH durci : clés plutôt que mots de passe (quand possible)

```
[AGENCE-DAKAR-AR720]ssh user admin authentication-type all
[AGENCE-DAKAR-AR720]ssh server-source -i Vlanif 10
```

- Préférer l'authentification par clé (selon support de la version) ; à défaut, mot de
  passe long + `authentication-mode aaa`.
- `ssh server-source -i Vlanif 10` : le serveur SSH n'écoute que sur l'interface LAN
  (pas sur le WAN). Vérifier la syntaxe exacte sur la version (`?` dans la CLI).
- Changer le port SSH par défaut ? Ça réduit le bruit des scans mais n'est pas une
  sécurité en soi. Le vrai gain = pas de SSH joignable depuis Internet du tout
  (administration via le tunnel VPN).

## 114. ACL d'administration : qui a le droit de se connecter

```
[AGENCE-DAKAR-AR720]acl number 2005
[AGENCE-DAKAR-AR720-acl-basic-2005]rule 5 permit source 192.168.10.0 0.0.0.255
[AGENCE-DAKAR-AR720-acl-basic-2005]rule 10 permit source 10.0.0.0 0.0.255.255
[AGENCE-DAKAR-AR720-acl-basic-2005]quit
[AGENCE-DAKAR-AR720]user-interface vty 0 4
[AGENCE-DAKAR-AR720-ui-vty0-4]acl 2005 inbound
[AGENCE-DAKAR-AR720-ui-vty0-4]quit
```

Seuls le LAN de l'agence et le réseau du siège peuvent ouvrir une session VTY.

## 115. SNMPv3 uniquement, jamais en clair sur le WAN

Rappel section 77 : `snmp-agent sys-info version v3` seul (retirer v2c/v1), communauté
supprimée, traps vers le superviseur via le tunnel. Vérifier qu'aucune politique
untrust→local n'autorise UDP 161.

## 116. Protection du plan de contrôle (CPU)

Limiter ce que le CPU accepte pour éviter qu'une attaque ou une tempête ne le mette à
genoux :

- Activer les protections anti-attaque intégrées (selon version : `anti-attack` /
  `cpu-defend policy` — **syntaxe à vérifier sur la fiche du modèle exact**).
- Principe : définir des seuils (pps) par protocole vers la zone local (ARP, ICMP, SSH…).
- Tester après activation : une politique trop stricte coupe l'administration légitime.

## 117. Journaliser les accès et les refus

```
[AGENCE-DAKAR-AR720]security-policy
[AGENCE-DAKAR-AR720-policy-security]rule name Log-Refus-WAN
[AGENCE-DAKAR-AR720-policy-security-rule-Log-Refus-WAN]source-zone untrust
[AGENCE-DAKAR-AR720-policy-security-rule-Log-Refus-WAN]destination-zone local
[AGENCE-DAKAR-AR720-policy-security-rule-Log-Refus-WAN]action deny
[AGENCE-DAKAR-AR720-policy-security-rule-Log-Refus-WAN]quit
[AGENCE-DAKAR-AR720-policy-security]quit
```

+ `info-center` vers le syslog (section 79). Relire périodiquement : des refus SSH en
boucle depuis la même IP = scan → bannir à la main ou via le superviseur.

## 118. Mots de passe : politique minimale

- Longueur ≥ 12, uniques par site et par usage (console ≠ admin ≠ PSK VPN ≠ SNMP).
- `cipher` partout (jamais en clair dans la config).
- Rotation : à chaque départ d'un technicien qui les connaissait, et au moins une fois
  par an pour les comptes partagés.
- Stockage : coffre de l'équipe, jamais dans un tableur partagé ni dans un e-mail.

## 119. Checklist de durcissement avant mise en service

```
[ ] Telnet/FTP/HTTP désactivés (ou justifiés)
[ ] SSH seul, restreint par ACL, pas joignable du WAN
[ ] SNMP v3 uniquement, communauté v2c supprimée
[ ] Politique untrust->local : uniquement IKE (500/4500) si VPN, rien d'autre
[ ] Politique untrust->trust : uniquement les NAT server documentés
[ ] Mots de passe en cipher, uniques par site
[ ] info-center vers syslog actif
[ ] CPU-defend / anti-attaque activé et testé
[ ] save + sauvegarde externe
```

---
---

# Partie N — Cas pratiques complets commentés

## 120. Cas pratique n°1 — Agence standard : config complète commentée

Scénario : agence de 30 personnes. WAN principal PPPoE (fibre FAI, login
`agence.dakar@fai.exemple`), WAN secours box 4G (192.168.8.1). LAN 192.168.10.0/24,
invités 192.168.20.0/24, équipements 192.168.30.0/24. Tunnel IPSec vers le siège
(192.168.0.0/24). Toutes les valeurs d'exemple sont fictives.

```
# --- Identité et base ---
[AR720]sysname AGENCE-DAKAR-AR720
# Nom explicite : site + modèle. On sait où on est dans chaque terminal.

[AGENCE-DAKAR-AR720]clock timezone Dakar add 00:00
# Fuseau pour des logs exploitables. NTP configuré plus bas.

# --- Accès console ---
[AGENCE-DAKAR-AR720]user-interface console 0
[AGENCE-DAKAR-AR720-ui-console0]authentication-mode password
[AGENCE-DAKAR-AR720-ui-console0]set authentication password cipher MotDePasseConsoleFictif123
[AGENCE-DAKAR-AR720-ui-console0]idle-timeout 10 0
[AGENCE-DAKAR-AR720-ui-console0]quit

# --- Admin SSH ---
[AGENCE-DAKAR-AR720]aaa
[AGENCE-DAKAR-AR720-aaa]local-user admin password cipher MotDePasseAdminFictif456
[AGENCE-DAKAR-AR720-aaa]local-user admin privilege level 15
[AGENCE-DAKAR-AR720-aaa]local-user admin service-type ssh terminal
[AGENCE-DAKAR-AR720-aaa]quit
[AGENCE-DAKAR-AR720]stelnet server enable
[AGENCE-DAKAR-AR720]ssh user admin authentication-type password
[AGENCE-DAKAR-AR720]user-interface vty 0 4
[AGENCE-DAKAR-AR720-ui-vty0-4]authentication-mode aaa
[AGENCE-DAKAR-AR720-ui-vty0-4]protocol inbound ssh
[AGENCE-DAKAR-AR720-ui-vty0-4]acl 2005 inbound
[AGENCE-DAKAR-AR720-ui-vty0-4]quit
# SSH uniquement, filtré par ACL (définie plus bas).

