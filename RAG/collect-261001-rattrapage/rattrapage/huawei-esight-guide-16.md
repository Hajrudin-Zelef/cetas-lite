---
id: collect-261001-rattrapage/rattrapage/huawei-esight-guide-16
title: "Huawei eSight — Guide ultra-complet d'exploitation terrain"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["agent", "incident"]
source: docs/RAG/collect-261001-rattrapage/huawei_esight_guide.md
source_anchor: ""
source_lines: [2374, 2550]
sha256: 102c81ae23f3f6bba452fe50042951466ecc9302a1c6804199d9a2370e9128f4
---

# Huawei eSight — Guide ultra-complet d'exploitation terrain

**À superviser :** AP en ligne/hors ligne par site, charge (utilisateurs
simultanés par AP), santé de l'AC (CPU, licences), interférences signalées.

**Seuils de départ :** > 10 % des AP d'un site hors ligne = major ;
AP avec > 30 utilisateurs = warning (densifier) ; AC CPU > 70 % = major.

**Réflexe dépannage :** « le Wi-Fi ne marche pas » → d'abord vérifier
dans eSight si les AP sont en ligne et si le switch PoE qui les alimente
va bien (souvent un problème d'énergie, pas radio — section 64).

## 125. Fiche réflexe : serveur et stockage

**Serveur (module server device management) — à superviser :** CPU,
mémoire, espace disque, état du matériel (disques RAID, alimentations,
ventilateurs), OS (services critiques).

**Stockage (OceanStor…) — à superviser :** capacité utilisée/disponible,
IOPS et latence, état des disques et contrôleurs, fonctions de protection
(anti-ransomware selon version).

**Seuils de départ :** disque > 80 % = warning, > 90 % = critical ;
disque RAID en défaut = major ; contrôleur down = critical.

---

# ANNEXE C — AUTOMATISATION AUTOUR D'ESIGHT

> eSight s'administre à la console, mais tout ce qui est **répétitif et
> vérifiable** gagne à être scripté à côté. Exemples en bash (à adapter).

## 126. Script : audit SNMP du parc avant découverte

Avant de lancer eSight sur un parc inconnu, auditer qui répond en SNMP.
Exemple fictif (nécessite `snmpwalk` ; community d'exemple à remplacer) :

```bash
#!/bin/bash
# audit_snmp.sh — teste la réponse SNMP d'une liste d'IP
# Usage : ./audit_snmp.sh liste_ip.txt
COMMUNAUTE="COMMUNAUTE_DE_TEST"   # exemple fictif — mettre la vraie
while read -r ip; do
  if snmpwalk -v2c -c "$COMMUNAUTE" -t 2 -r 1 "$ip" sysName.0 >/dev/null 2>&1; then
    echo "OK   $ip"
  else
    echo "KO   $ip"
  fi
done < "$1"
```

**Usage :** les KO alimentent la check-list de préparation (section 27-28)
avant la découverte eSight. On ne découvre que du OK.

## 127. Script : vérification des trap-targets

Vérifier que chaque équipement Huawei pointe bien vers eSight pour les
traps. Principe (via SSH en boucle, ou via le Smart Configuration Tool
d'eSight) : lire la configuration des `snmp-agent target-host` et comparer
à l'IP d'eSight. Tout écart = trap-target à corriger (voir cas n°3).

**Automatiser** : rejouer ce contrôle **tous les mois** — les
trap-targets « s'évaporent » lors des remplacements d'équipements ou des
restaurations de vieilles configs.

## 128. Script : contrôle des sauvegardes eSight

```bash
#!/bin/bash
# check_backup_esight.sh — alerte si pas de backup récent
# À brancher sur votre supervision système (appel quotidien)
REPERTOIRE_BACKUP="/mnt/sauvegardes/esight"
DELAI_MAX_HEURES=30
dernier=$(find "$REPERTOIRE_BACKUP" -type f -name '*.bak' -o -name '*.zip' \
  | xargs ls -t 2>/dev/null | head -1)
if [ -z "$dernier" ]; then
  echo "CRITICAL - aucun fichier de sauvegarde eSight trouve"
  exit 2
fi
age_heures=$(( ( $(date +%s) - $(stat -c %Y "$dernier") ) / 3600 ))
if [ "$age_heures" -gt "$DELAI_MAX_HEURES" ]; then
  echo "CRITICAL - derniere sauvegarde eSight il y a ${age_heures}h ($dernier)"
  exit 2
fi
echo "OK - derniere sauvegarde eSight il y a ${age_heures}h"
exit 0
```

**Règle** : ce script (ou équivalent) tourne **tous les jours** et remonte
vers votre supervision système — jamais de confiance aveugle dans la
planification eSight (voir cas n°7).

## 129. Inventaire : rapprocher eSight du parc réel

**Le rituel trimestriel :**
1. Exporter l'inventaire eSight (CSV).
2. Le rapprocher de la **source de vérité du parc** (CMDB, tableau
   d'immobilisations, inventaire physique).
3. Traiter les écarts : équipements dans eSight mais réformés (à
   supprimer), équipements réels absents d'eSight (à découvrir),
   modèles/versions divergents (à resynchroniser).
4. Consigner le rapprochement (date, écarts, actions).

Un inventaire eSight qui dérive du réel, c'est des licences payées pour
rien et des équipements non supervisés — les deux en même temps.

## 130. Superviser eSight depuis l'extérieur (sonde minimale)

Avec n'importe quel outil (même un cron + curl sur une autre machine) :

- **Ping** du serveur eSight toutes les minutes.
- **Test HTTPS** : la page de login répond-elle en < 5 s ?
- **Fraîcheur des données** : via l'API northbound ou un export,
  vérifier que des alarmes/perfs récentes existent (une supervision qui
  ne reçoit plus rien depuis 2 h est peut-être aveugle, pas calme).
- En cas d'échec : **alerte par un canal indépendant** d'eSight
  (SMS via une autre passerelle, appel d'astreinte).

---

# ANNEXE D — CAS PRATIQUES COMMENTÉS

## 131. Cas pratique n°1 : coupure du lien inter-sites, vendredi 18h05

**Contexte** : ETI, 3 sites. eSight Standard, 140 équipements. Lien
fibre opérateur entre le siège et le site B.

**Déroulé :**
- 18h05 : tempête — 1 alarme Critical « linkDown » + 34 alarmes
  « équipement injoignable » (tout le site B).
- 18h06 : l'astreinte reçoit le **SMS** (règle : Critical backbone → SMS).
  Connexion VPN, console eSight : la corrélation (configurée 2 mois plus
  tôt après un incident similaire) rattache les 34 alarmes à l'alarme racine.
- 18h10 : runbook « lien inter-sites down » appliqué : vérification
  topologie (lien vraiment down, pas un flap — historique stable),
  équipements d'extrémité injoignables des deux côtés → pas un problème
  d'équipement local.
- 18h15 : appel opérateur — incident connu sur la fibre, délai annoncé 22h.
  Alarme acquittée avec n° de ticket opérateur. **Masquage temporaire**
  des alarmes du site B (date de fin : lendemain 8h) pour garder la
  console lisible.
- 21h40 : retour opérateur, lien up, traps linkUp reçus, alarmes soldées
  en cascade. Levée du masquage. Vérification : tout le site B est revenu
  (synchronisation manuelle pour solder 2 alarmes fantômes).

**Ce qui a bien marché :** corrélation configurée, SMS reçu, runbook suivi,
pas de panique. **À améliorer :** le lien de secours 4G du site B n'a pas
basculé automatiquement — action planifiée au comité mensuel.

## 132. Cas pratique n°2 : migration SNMP v2c → v3 sur 120 équipements

**Contexte** : audit interne — community v2c identique partout depuis
5 ans. Décision : passage en SNMP v3 authPriv.

**Déroulé :**
1. **Préparation** : création du profil eSight `SNMPv3-authPriv`
   (utilisateur dédié, SHA/AES — exemples fictifs, mots de passe au coffre).
2. **Pilote** : 5 équipements (1 par famille). Déploiement de l'utilisateur
   v3 côté équipement via Smart Configuration Tool, bascule du profil
   eSight sur ces 5, vérification polling + traps pendant 1 semaine.
3. **Généralisation par vagues** : 20-30 équipements par soirée de
   maintenance, en commençant par les accès et en finissant par le cœur.
   Après chaque vague : contrôle des timeouts et des traps.
4. **Bascule** : une fois tout le parc en v3 et stable depuis 2 semaines,
   **suppression des community v2c** côté équipements + suppression du
   profil v2c dans eSight.
5. **Vérification finale** : audit snmpwalk (section 126) — 0 équipement
   répondant encore en v2c.

**Pièges évités :** ne pas supprimer v2c avant d'avoir validé v3 partout
(sinon supervision aveugle) ; traiter les équipements anciens ne
supportant pas v3 (maintien en v2c avec ACL strictes, plan de remplacement).

## 133. Cas pratique n°3 : déploiement eSight sur 3 sites

**Contexte** : groupe avec siège (80 équipements) + 2 sites distants
(30 chacun). Choix : **1 eSight au siège** (pas de hiérarchie — le WAN
est fiable et l'équipe est centralisée).

