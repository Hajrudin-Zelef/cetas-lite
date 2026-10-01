---
id: collect-261001-rattrapage/rattrapage/huawei-ekit-guide-13
title: "GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ekit_guide.md
source_anchor: ""
source_lines: [1142, 1214]
sha256: 6c0e561eb5bde77276e854c56d639a8c8315c5225ede43220aea34761ad63e24
---

# GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)

**Symptômes :** MAJ firmware poussée hier soir, ce matin le site est en vrac.
**Diagnostic :**
1. Quels équipements ont été mis à jour ? Lesquels sont en panne ?
2. La **release note** mentionnait-elle un changement de comportement ? (ex. : nouveau mode de gestion par défaut)
3. La configuration a-t-elle **survécu** à la MAJ ? (comparer avec la sauvegarde d'avant-MAJ, section 89)
**Solution :** si un seul équipement : le **revenir en arrière** (downgrade vers la version précédente — procédure **à vérifier sur la documentation officielle** par modèle) ou restaurer sa config. Si tout le site : restaurer les configs une par une depuis la sauvegarde. **Leçon :** la prochaine fois, site pilote + 48 h d'observation (section 42). Et on ne fait pas de MAJ le vendredi.

## 100. Cas n°9 à n°16 — En rafale (symptômes → piste → solution)

**N°9 — Un switch ne distribue plus de PoE sur aucun port.** Piste : alimentation du switch (LED PWR ?), budget PoE désactivé après MAJ, surchauffe. Solution : vérifier l'alim, réactiver le PoE dans la config, laisser refroidir, RMA si le module PoE est mort.
**N°10 — Débit filaire plafonné à ~90 Mbit/s sur un PC.** Piste : lien négocié à 100 Mbit/s (câble ou prise défectueuse). Solution : tester le brin, refaire la prise, vérifier la négociation sur le port du switch.
**N°11 — Les téléphones IP se coupent / voix hachée.** Piste : pas de QoS, VLAN voix mélangé au data, ou switch saturé. Solution : VLAN 60 dédié, priorité QoS voix, vérifier que les téléphones sont bien sur le bon VLAN (LLDP-MED si supporté — à vérifier).
**N°12 — Une caméra ne remonte plus au NVR.** Piste : PoE (cycle PoE à distance d'abord !), IP changée (DHCP sans réservation), VLAN erroné. Solution : cycle PoE via le cloud, réservation DHCP ou IP fixe, vérifier le VLAN du port.
**N°13 — Conflit d'IP : « un autre appareil utilise cette adresse ».** Piste : IP fixe en doublon avec le scope DHCP, ou un équipement invité avec IP en dur. Solution : exclure les IP fixes du scope DHCP (toujours !), passer en réservation DHCP.
**N°14 — Le VPN IPsec ne monte pas (AR/USG).** Piste : paramètres IKE phase 1/2 incohérents entre les deux sites, NAT devant l'AR, PSK erroné, port UDP 500/4500 filtré. Solution : comparer les configs des deux côtés ligne par ligne, vérifier le NAT-T, tester sans NAT si possible.
**N°15 — Impossible d'accéder à l'interface locale d'un équipement.** Piste : gestion locale désactivée en mode cloud, mauvaise IP / mauvais VLAN depuis le poste, HTTP/HTTPS désactivé. Solution : se mettre sur le bon VLAN avec une IP du bon sous-réseau, vérifier que la gestion locale est autorisée (ou repasser temporairement en gestion locale selon le modèle).
**N°16 — L'app eKit ne se connecte plus à mon compte.** Piste : mot de passe expiré/modifié, compte verrouillé, problème réseau du téléphone, maintenance cloud Huawei. Solution : vérifier sur un 2e appareil, réinitialiser le mot de passe via l'e-mail du compte, vérifier le statut du cloud (distributeur / support Huawei), et en attendant : gestion locale des équipements si elle est activée.

---

## 101. Maintenance : le plan annuel type (par site)

| Fréquence | Action | Durée indicative |
|---|---|---|
| **Hebdomadaire** | Rituel cloud 15 min : dashboard global → sites orange/rouge → action ou ticket (section 43) | 15 min |
| **Mensuelle** | Vérifier les alertes récurrentes, les scopes DHCP (taux de remplissage), l'état des firmwares (N-1 max ?) | 30 min/site |
| **Trimestrielle** | Revue des accès (section 34), test d'un cycle PoE sur un AP témoin, dépoussiérage baie, vérification onduleur (voyants, test batterie) | 1 h/site |
| **Semestrielle** | Survey Wi-Fi rapide (tour avec smartphone : débits aux points clés), vérification étiquetage, mise à jour fiche site | 2 h/site |
| **Annuelle** | Mise à jour firmware planifiée (section 42), test de restauration de sauvegarde (section 89), révision du plan d'adressage, renégociation contrat de maintenance | 1/2 journée |

**Le contrat de maintenance, c'est ça que tu vends** : pas « on viendra si ça tombe en panne », mais ce plan, écrit, chiffré, avec les KPI de la section 80.

## 102. Sécurité des accès : le coffre et les règles

- **Coffre de mots de passe** (Bitwarden, KeePass — ton choix) : un coffre **par client**, partagé uniquement avec les techniciens habilités. Jamais de PSK ou d'admin dans un e-mail, un SMS ou un post-it sur la baie.
- **Mots de passe** : uniques par site et par équipement critique, 20+ caractères générés. Les exemples de ce guide (`BUREAU-Staff`, `192.168.10.1`) sont **fictifs** — ne les utilise jamais tels quels.
- **Comptes** : 1 compte propriétaire générique par tenant (section 17), comptes nominatifs pour les techniciens, retrait **immédiat** au départ d'un technicien ou à la fin d'une prestation.
- **Gestion locale** : change les identifiants par défaut **dès** la première configuration, désactive les protocoles non utilisés (Telnet si présent — préférer SSH/HTTPS).
- **SNMP** : v3 avec authentification si le modèle le supporte (les switchs eKit annoncent SNMPv1/v2c/v3 — utilise v3).
- **Journal** : note chaque accès admin sensible (qui, quand, quoi) dans le dossier du site.

## 103. Partage d'accès avec le client : le cadre contractuel

À écrire noir sur blanc dans le contrat :
- qui administre quoi (toi : tout ; client : lecture seule ou rien — section 28) ;
- ce qui se passe si le client modifie la config (intervention facturée, exclusion de garantie sur la partie modifiée) ;
- la **réversibilité** : en fin de contrat, transfert du tenant/compte au client ou à son nouveau prestataire (procédure + délai + format des exports). Un client « otage » de son prestataire, c'est illégal dans l'esprit et toxique commercialement — prévois la sortie dès l'entrée.
- les **données** : où sont hébergées les données de gestion (cloud Huawei), qui y a accès, durée de conservation.

## 104. Sécurité réseau : les 10 règles non négociables

1. VLAN invités **isolé**, toujours.
2. PSK robustes, uniques par site, **tournés** au moins 1×/an (et à chaque départ d'un employé sensible).
3. WPA3 quand les clients le permettent, WPA2-AES minimum — **jamais** de WPA/TKIP ni de réseau ouvert sans portail.
4. SSID staff **non diffusé** (obscurité ≠ sécurité, mais ça réduit le bruit).
5. Caméras et IoT sur VLAN dédiés, **sans accès Internet** sauf besoin (NVR accessible via VPN).
6. Gestion des équipements sur VLAN ou via cloud uniquement — pas d'admin depuis le Wi-Fi invités.
7. Mises à jour firmware suivies (une faille non patchée = la porte ouverte).
8. USG6000F-S (ou équivalent) dès que le client a des données sensibles : IPS/AV/filtrage URL activés, signatures à jour.
9. Sauvegardes chiffrées, stockées hors site.
10. **Personne ne branche** un équipement personnel sur le LAN staff (ou alors VLAN invités). Le PC portable du cousin « juste pour imprimer », c'est non.

## 105. Pièces de rechange : le stock minimal

| Pour | Stock conseillé | Pourquoi |
|---|---|---|
| Parc < 5 sites | 1 AP361 + 1 petit switch PoE (S220) | Couvre 80 % des pannes (l'AP est l'élément le plus exposé) |
| Parc 5-20 sites | + 1 AR180 + 1 S310 + modules SFP de secours + 2 alimentations | Autonomie RMA de 2 semaines |
| Client hôtel / critique | 2 AP + 1 switch **sur site client** (dans la réserve, étiquetés) | Intervention en minutes, pas en jours |
| Toujours | Câbles RJ45, jarretières fibre, étiquettes, colliers, 1 testeur | Le consommable qui manque un dimanche soir |

**Règle :** un équipement de secours **pré-configuré** (onboardé sur le site, config synchronisée) se remplace en 10 minutes. Un équipement « à configurer » se remplace en 2 heures.

