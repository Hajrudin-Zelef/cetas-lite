---
id: collect-261001-rattrapage/rattrapage/huawei-ap761-guide-23
title: "Guide ultra-complet — Huawei eKit AP761"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ap761_guide.md
source_anchor: ""
source_lines: [2288, 2405]
sha256: e0f73db074ad1a7b04ad1efc5318a894713321f33b1d0f1c2aab843a1fa60416
---

# Guide ultra-complet — Huawei eKit AP761

**Checklist annuelle électrique (par AP ou par départ) :**
- [ ] **Tension PoE** au pied de l'AP si accessible (ou négociation affichée côté switch : af/at ?).
- [ ] **Budget PoE** du switch : toujours suffisant ? (On a ajouté des AP depuis ? — chap. 25.)
- [ ] **Serrage** des connexions dans les coffrets (dilatation thermique annuelle).
- [ ] **Continuité de terre** : mât, boîtier, parafoudres (ohmmètre — valeur selon norme locale).
- [ ] **Parafoudres :** état, date de mise en service, remplacement si > 5 ans ou après coup encaissé (chap. 26).
- [ ] **Onduleur du local** : autonomie réelle vs dimensionnement (les AP + le switch PoE sont une charge non négligeable — voir ton guide onduleurs).
- [ ] **Thermographie** (si caméra dispo) : un connecteur ou un injecteur qui chauffe anormalement = un problème en gestation.

**Tracer les mesures** : un tableau de suivi annuel (tension, terre, observations) permet de voir les dérives — c'est ça, la maintenance préventive, par opposition au « on répare quand ça casse ».

## 104. Contrôle radio annuel (audit)

Une fois par an, refaire un **mini-survey** (chap. 60, version allégée) : 30 minutes par site.

**Protocole allégé :**
1. Relever les canaux/largeurs/puissances **réels** (ont-ils dérivé ? un voisin a-t-il changé de canal ?).
2. Mesurer le RSSI en 3 points de référence (les mêmes chaque année — les noter dans le dossier).
3. Comparer aux relevés de l'année précédente : **une dérive de 10 dB = quelque chose a changé** (végétation, nouvel obstacle, AP voisin, AP qui fatigue).
4. Vérifier la répartition 2.4/5 GHz des clients (chap. 42).
5. Vérifier les rogues (chap. 64) : un an, c'est long — il a pu s'en installer.

**Ce qu'on cherche :** pas la perfection, mais les **dérives**. Un réseau Wi-Fi extérieur vit : la végétation pousse, les voisins changent de box, les usages évoluent. L'audit annuel, c'est le moment où on le constate **avant** que les utilisateurs s'en plaignent.

## 105. Gestion du cycle de vie et fin de vie

**Durée de vie réaliste d'un AP761 extérieur : 5–7 ans.** Après :
- Les condensateurs et les joints vieillissent (surtout en plein soleil).
- Le Wi-Fi des clients a évolué (en 2030, le Wi-Fi 7 sera banal).
- Le support firmware se tarit (fin de support constructeur — à vérifier sur la fiche du modèle exact).

**Signes de fin de vie :**
- Redémarrages inexpliqués de plus en plus fréquents.
- Débit qui se dégrade sans cause externe (composants RF qui dérivent).
- Plus de firmware depuis 2+ ans (failles non corrigées = risque sécurité).
- Pièces de rechange introuvables.

**Planification du renouvellement :**
1. **Inventaire** : âge de chaque AP, date d'achat, S/N (le dossier de site, chap. 101).
2. **Provisionner** le renouvellement à 5 ans : budgéter le remplacement par tiers (un tiers du parc par an sur 3 ans = pas de big bang).
3. **Ne pas attendre la panne** : un AP extérieur qui meurt un vendredi soir en plein événement = une urgence évitable.
4. **Recyclage** : DEEE — les AP contiennent des métaux et des composants à traiter en filière adaptée. Effacer la config avant mise au rebut (`reset` + vérification).

**Et le Wi-Fi 7 dans tout ça ?** Le renouvellement naturel (5–7 ans) est **le bon moment** pour migrer vers le Wi-Fi 7 : les produits seront matures, les clients aussi, et tu ne jettes pas du matériel amorti (chap. 95).

## 106. Durcissement : checklist sécurité

Le durcissement, c'est la différence entre « ça marche » et « ça marche et c'est sûr ». Checklist à appliquer à chaque AP761, **dès l'installation** :

**Comptes et accès :**
- [ ] Compte admin par défaut **changé** (ou désactivé si l'auth se fait via le cloud).
- [ ] Pas de compte « invité »/« test » qui traîne.
- [ ] Mots de passe ≥ 16 caractères, uniques par site (coffre de mots de passe d'équipe).
- [ ] Sessions web en **HTTPS uniquement** (HTTP désactivé ou redirigé).

**Réseau de management :**
- [ ] VLAN management dédié (chap. 52), pas de management depuis les VLAN invités/IoT.
- [ ] ACL : seuls les postes d'administration et le superviseur accèdent au management.
- [ ] Telnet **désactivé**, SSHv2 uniquement (chap. 107).

**Wi-Fi :**
- [ ] Pas de WEP/TKIP, pas de SSID ouvert sans portail (chap. 44).
- [ ] PMF au minimum optionnel sur les SSID sensibles (chap. 51).
- [ ] WIDS activé, alertes configurées (chap. 63).
- [ ] SSID invité : isolation client + VLAN + ACL Internet seul (chap. 49, 53).

**Suivi :**
- [ ] NTP configuré (chap. 70), syslog centralisé (chap. 69).
- [ ] Firmware à jour (chap. 74), alertes configurées (chap. 71).
- [ ] Sauvegarde faite et testée (chap. 72–73).

**Revue :** repasser cette checklist **une fois par an** et après chaque changement majeur. Un durcissement qui n'est pas revu se dégrade (nouveau SSID « temporaire » oublié, compte prestataire jamais supprimé…).

## 107. Durcissement : management (SSH, SNMP, web)

**SSH (STelnet) :**
```
[AP761] ssh server enable
[AP761] undo telnet server enable        # couper Telnet : mots de passe en clair, inacceptable
[AP761] ssh server timeout 60
[AP761] ssh authentication-retries 3
# Restreindre par ACL aux postes d'admin :
[AP761] acl number 2000
[AP761-acl-basic-2000] rule permit source 192.168.10.0 0.0.0.255
[AP761-acl-basic-2000] quit
[AP761] ssh server acl 2000
```
- Clés SSH plutôt que mots de passe si la version le supporte (à vérifier).
- Désactiver l'accès SSH depuis les VLAN utilisateurs : le management ne répond que sur le VLAN 10.

**SNMP :** v3 uniquement, auth+priv, ACL vers le superviseur (chap. 68). **Désactiver v1/v2c** explicitement — un community `public` qui traîne, c'est une lecture du réseau offerte.

**Web :** HTTPS uniquement. Si l'interface web n'est pas utilisée (tout en CLI/cloud), la **désactiver** : moins de surface d'attaque.

**Comptes :** principe du moindre privilège — un compte « lecture seule » pour la supervision, un compte « admin » pour les changements, **jamais de compte partagé** (« admin » connu de 10 personnes = 0 traçabilité). En cas de départ d'un technicien : révoquer **le jour même**.

## 108. Durcissement : Wi-Fi (chiffrement, PMF, WIDS)

**Chiffrement :**
- Bannir WEP, WPA-TKIP, SSID ouverts sans portail (chap. 44).
- Objectif : WPA3-SAE ou transition sur les SSID salariés (chap. 46–47), WPA2-PSK fort minimum partout ailleurs.
- Clés PSK ≥ 20 caractères, uniques par SSID, rotation trimestrielle pour l'invité.

**PMF (802.11w) :** optionnel minimum sur les SSID salariés, obligatoire si le parc le supporte (chap. 51). C'est la protection contre la désauthentification forgée — l'attaque la plus simple et la plus courante.

**WIDS :** activé avec alertes (chap. 63). Revoir **mensuellement** les détections : un WIDS dont personne ne lit les alertes ne sert à rien.

**SSID :**
- Pas d'information sensible dans les noms (chap. 43).
- SSID invité isolé (VLAN + ACL + isolation client).
- Désactiver les SSID inutilisés : chaque SSID actif est de l'airtime consommé et une surface d'attaque.

**Physique :** un AP761 accessible au public (terrasse, camping) = risque de **reset physique** (bouton) ou de vol. Câble antivol (chap. 8), fixation hors de portée (> 3 m), et **en cas de vol : révoquer** — changer les clés PSK des SSID (l'AP volé contient la config… sauf si elle était chiffrée — à vérifier sur la fiche du modèle exact) et retirer l'AP du cloud.

## 109. Pense-bête de poche — page 1 : specs

À imprimer, plastifier, garder dans la caisse à outils.

