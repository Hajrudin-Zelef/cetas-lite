---
id: collect-261001-rattrapage/rattrapage/huawei-usg6000-guide-15
title: "Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["arr", "revenue"]
source: docs/RAG/collect-261001-rattrapage/huawei_usg6000_guide.md
source_anchor: ""
source_lines: [2214, 2337]
sha256: 56a448d66ca25eb6477c526c2f2e2ac4ec9a1265db2202b0a8e7a9f7ab9f35c7
---

# Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)

## 134. Politique de mot de passe

- Longueur **≥ 12 caractères**, complexité (4 familles), **historique** (pas de réutilisation des 5 derniers), **durée de vie** 90 jours pour les comptes à privilèges.
- ⚠️ **Pas de mot de passe dans les scripts en clair**, pas de post-it, pas de « mot de passe d'équipe » partagé par 8 personnes.
- Comptes de **secours** : 1 compte « break-glass » local, mot de passe **sous enveloppe scellée** dans le coffre, **testé 1 fois par an**.

## 135. Protection Anti-DDoS / anti-attaque

L'USG6000 intègre des protections contre : **SYN flood, UDP flood, ICMP flood, scans, attaques par zone**.

```huawei
system-view
# Activer la détection d'attaque par zone (exemple untrust) :
[USG] firewall defend syn-flood enable zone untrust
[USG] firewall defend udp-flood enable zone untrust
# Seuils (à ajuster selon ton trafic normal - mesurer AVANT d'activer le blocage) :
[USG] firewall defend syn-flood threshold 10000 zone untrust   # paquets/s (exemple)
# Liste noire automatique :
[USG] firewall blacklist enable
save
```

⚠️ **Méthodologie** : 1) mesurer le trafic normal (pics), 2) activer en **détection seule**, 3) régler les seuils **au-dessus** des pics légitimes, 4) passer en **blocage**. Un seuil trop bas = **auto-DDoS** (tu bloques tes propres utilisateurs).

## 136. Filtrage géographique (si pertinent)

Si ton activité n'a aucune raison de recevoir du trafic de certains pays : blocage par **région/pays** dans les politiques (untrust→dmz, untrust→local).
⚠️ À manier avec prudence : CDN et services cloud ont des IP dans le monde entier — un blocage pays trop large **casse** des services légitimes (mises à jour, SaaS). Tester avant.

## 137. Désactiver les services inutiles

```huawei
# Inventaire des services actifs :
[USG] display tcp status
[USG] display udp status      # (selon version)
# Désactiver ce qui ne sert pas : telnet (133), ftp server, snmp v1/v2c si v3 en place...
[USG] undo ftp server enable
save
```

📋 Principe : **tout service actif est une surface d'attaque**. Si tu ne sais pas à quoi sert un port en écoute, cherche — puis coupe.

## 138. Journalisation des actions d'administration

- Activer les **logs de configuration** (qui a tapé quoi, quand) — `info-center` + envoi syslog.
- En HA/pilotage centralisé, les changements doivent passer par **demande → validation → application → vérification** (même à deux, même « vite fait »).
- 📋 Le « vite fait » non tracé, c'est la cause n°1 des « mais qui a touché à ça ?! ».

## 139. Sécurité des sauvegardes

- Les fichiers de config contiennent des **hashes de mots de passe** et des **PSK** : les stocker **chiffrés** (dossier chiffré, coffre), **jamais** sur un partage ouvert ni par mail.
- Droits d'accès restreints, **journal des accès** aux sauvegardes.

## 140. Mises à jour de sécurité : veille

- S'abonner aux **bulletins de sécurité Huawei** (alertes CVE).
- Criticité **critique** sur un firewall exposé = **patch sous 30 jours** (ou mesure compensatoire immédiate : règle de blocage, désactivation du service vulnérable).
- 📋 Noter chaque CVE applicable dans le journal (section 131) avec la décision (patché / compensé / non concerné + pourquoi).

## 141. Segmentation interne : ne pas s'arrêter au périmètre

Le firewall périmétrique ne protège pas contre : un poste infecté qui attaque le LAN, un prestataire branché en salle serveur, un IoT vérolé.
→ **Zones internes** (section 21 : IOT, GUEST, PROD), **politiques inter-VLAN** sur l'USG ou les switchs, **802.1X** sur les accès filaires sensibles. La défense en profondeur, pas en coquille d'œuf.

## 142. Audit de durcissement : grille annuelle

📋 Grille d'audit (1 fois/an) :
- [ ] Comptes : nominatifs ? anciens comptes supprimés ? break-glass testé ?
- [ ] Mots de passe : politique appliquée ? pas de défaut restant ?
- [ ] Accès : Telnet/HTTP désactivés ? sources restreintes ? pas d'admin depuis untrust ?
- [ ] Politiques : default deny ? règles 0-hit nettoyées ? DMZ→trust en deny ?
- [ ] UTM : signatures à jour ? profils actifs sur les bonnes règles ? exceptions revues ?
- [ ] VPN : PSK robustes ? certificats valides ? comptes VPN obsolètes supprimés ?
- [ ] HA : testée ? licences à jour des deux côtés ?
- [ ] Logs : syslog reçu ? rétention conforme ? rapports lus ?
- [ ] Firmware : version supportée ? CVE en cours traitées ?
- [ ] Physique : baie verrouillée ? console non branchée en permanence ?

---
---

# BLOC N — DÉPANNAGE TERRAIN (18 CAS)

> Méthode générale : **1)** qualifier (quoi, qui, quand, depuis quand), **2)** reproduire à la demande, **3)** observer (`display`, `debugging` ciblé), **4)** hypothèse unique, **5)** correction, **6)** vérification + traçabilité. Ne change jamais deux choses à la fois.

## 143. Méthode : l'ordre de traversée d'un paquet (à connaître par cœur)

```
Entrée interface → Zone source → [Session existante ? oui → forward]
                                   non → NAT destination (DNAT)
                                       → Routage (zone destination)
                                       → Politique de sécurité (1ère règle qui matche)
                                       → NAT source (SNAT)
                                       → UTM (profils de la règle)
                                       → Sortie interface
```

🔧 Quand « ça ne passe pas », remonte cette chaîne **dans l'ordre** : c'est presque toujours l'une de ces étapes. Les commandes reines : `display firewall session table`, `display security-policy`, `display nat-policy`, `display ip routing-table`, `debugging packet-filter`.

## 144. Cas 1 : « Rien ne passe » après une modification

- **Symptômes** : tout le trafic coupé (ou tout un sens) juste après un changement.
- **Diagnostic** : `display configuration` / comparer avec la sauvegarde ; `display security-policy rule all` (ordre), `display zone` (interface toujours dans sa zone ?).
- **Causes fréquentes** : `default action deny` activé sans règles de rattrapage ; interface **retirée de sa zone** par erreur ; `save` oublié puis reboot (l'ancienne config est revenue — ou l'inverse).
- **Solution** : restaurer la sauvegarde (section 116) ou corriger la modification. **Toujours** faire `display` AVANT/APRÈS chaque changement.

## 145. Cas 2 : une politique bloque un flux légitime

- **Symptômes** : une appli ne marche plus (ou n'a jamais marché) à travers l'USG.
- **Diagnostic** :
```huawei
[USG] display firewall session table source inside 192.168.10.25 destination outside 203.0.113.50
# Pas de session ? → bloqué avant création. Session en deny ?
[USG] display security-policy rule all   # quelle règle matche ? (ordre !)
# Test ciblé : autoriser temporairement en LOG pour voir :
# puis debugging :
[USG] debugging packet-filter ip source 192.168.10.25   # syntaxe à vérifier selon version
[USG] undo debugging all
```
- **Causes fréquentes** : règle trop restrictive (port oublié), **mauvais ordre** (un deny général avant le permit), objet d'adresse pas à jour, **plage horaire** active, critère `user` sans authentification.
- **Solution** : corriger la règle (ou son ordre avec `move rule`), re-tester, **retirer le debugging**.

## 146. Cas 3 : le retour ne passe pas (routage asymétrique / session)

