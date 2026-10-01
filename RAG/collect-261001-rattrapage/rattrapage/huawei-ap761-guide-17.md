---
id: collect-261001-rattrapage/rattrapage/huawei-ap761-guide-17
title: "Guide ultra-complet — Huawei eKit AP761"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ap761_guide.md
source_anchor: ""
source_lines: [1667, 1796]
sha256: b3dc84ee6c4013547420583d1ae07f9ec07a700ca68b4d3f3818cada1bbbeda5
---

# Guide ultra-complet — Huawei eKit AP761

**Bonnes pratiques :**
- Envoyer vers un **serveur syslog central** (celui de ton SI — voir ton guide Loki/Syslog si applicable), pas en local seul : si l'AP meurt, ses logs meurent avec lui.
- Niveau **informational** en routine, **debugging** seulement en dépannage ciblé (ça inonde).
- **Horodatage NTP obligatoire** (chap. 70) : des logs sans heure fiable sont inutilisables pour corréler avec d'autres équipements.
- **Rétention :** alignée sur tes obligations (au minimum quelques mois ; accès public = obligations légales, chap. 49).

**Ce qu'on cherche dans les logs :**
- `Deauth` en boucle sur un client → problème d'authentification ou client têtu.
- `Channel change` fréquent → DFS ou interférence (cas 8).
- `CAPWAP` down/up → problème réseau vers l'AC (mode Fit).
- `PoE` / `power` → alim instable (cas 1, cas 9).

## 70. NTP : pourquoi c'est critique

Un AP sans heure juste, c'est :
- Des **logs incohérents** (impossible de corréler avec le switch ou le firewall).
- Des **échecs TLS** vers le cloud (certificat « pas encore valide » ou « expiré ») → l'AP ne s'adopte pas (cas 2).
- Des **problèmes 802.1X** (Kerberos et les certificats détestent le décalage horaire).
- Des stats de supervision fausses.

**Configuration :**
```
[AP761] ntp-service unicast-server 192.168.10.1
# Idéalement 2 serveurs pour la redondance :
[AP761] ntp-service unicast-server 192.168.10.2
<AP761> display ntp-service status   # vérifier « synchronized »
```

**Architecture :** tes AP synchronisent sur **ton** serveur NTP interne (qui lui-même se synchronise sur des sources externes), pas directement sur Internet — moins de flux sortants, plus de contrôle. Si le site est isolé sans NTP interne, autoriser au moins un serveur public via le firewall.

**Vérification :** après chaque (re)démarrage, contrôler que l'AP est synchronisé avant de diagnostiquer autre chose. **Un AP qui vient de redémarrer avec une horloge à 1970 et qui refuse le cloud = NTP d'abord, diagnostic après.**

## 71. Alertes et seuils : que configurer

Ni trop ni trop peu : trop d'alertes = on ne les lit plus ; pas assez = on découvre les pannes par les utilisateurs. La config minimale viable :

| Alerte | Seuil | Canal | Priorité |
|---|---|---|---|
| AP hors ligne | Immédiat | SMS/app + mail | 🔴 Critique |
| Rogue AP non classifié | Immédiat | Mail | 🟠 Haute |
| Utilisation canal 5 GHz > 75 % | Pendant 30 min | Mail | 🟡 Moyenne |
| Client avec RSSI < −80 dBm nombreux | Tendance hebdo | Rapport | 🟡 Moyenne |
| Firmware obsolète | Mensuel | Mail | 🟢 Basse |
| Certificat RADIUS < 30 jours d'expiration | J-30, J-7 | Mail + SMS | 🔴 Critique (si 802.1X) |
| Espace flash AP < 20 % | Mensuel | Mail | 🟢 Basse |

**Règles d'or des alertes :**
1. **Chaque alerte doit déclencher une action connue.** « AP hors ligne → appeler l'astreinte, vérifier PoE » pas « AP hors ligne → 🤷 ».
2. **Tester les alertes** : débrancher volontairement un AP un jour calme et vérifier que l'alerte arrive. Une alerte jamais testée est une illusion.
3. **Astreinte :** qui reçoit quoi, quand. Le SMS critique à 3 h du matin pour un AP de terrasse de camping en hiver = à discuter (criticité par site).
4. **Éviter les flaps :** seuils avec hystérésis (alerter à 75 %, « rétabli » à 60 %) pour ne pas spammer.

## 72. Sauvegarde de configuration : méthodes

Trois méthodes complémentaires, pas concurrentes :

**1. Sauvegarde locale (flash de l'AP) — le minimum :**
```
<AP761> save
<AP761> display saved-configuration   # vérifier que c'est bien écrit
```
Protège contre le redémarrage, pas contre la foudre ni le vol.

**2. Export vers un serveur — la vraie sauvegarde :**
```
# Via SFTP depuis un poste d'administration :
sftp admin@192.168.10.11
sftp> get vrpcfg.zip /sauvegardes/AP761-Cour-Est_2026-09-27.zip
```
Automatiser : un script hebdomadaire qui aspire les configs de tous les AP (clef SSH dédiée, compte en lecture seule si possible).

**3. Version cloud — le filet de sécurité :**
En mode cloud, la config est dans la plateforme. **Exporter une copie** après chaque changement majeur + noter la version firmware. Le cloud n'est pas une sauvegarde si tu n'as plus accès au compte (départ du prestataire, litige) : **l'export local reste obligatoire**.

**Ce qu'on sauvegarde, au-delà de la config :**
- [ ] Fichier de config (`vrpcfg.zip` ou équivalent).
- [ ] Version firmware exacte (`display version`).
- [ ] Plan d'adressage, VLAN, SSID, clés (dans un coffre, pas dans le même fichier que la config en clair).
- [ ] Photos du montage, plan de canaux, relevés de survey.
- [ ] Dossier de site à jour (qui, quoi, où, quand).

**Test de restauration annuel** (chap. 73) : le seul moyen de savoir que ça marche.

## 73. Restauration de configuration

**Scénario :** l'AP761-Cour-Est a pris la foudre. Le remplaçant est sur l'établi. Objectif : 30 minutes chrono.

**Procédure :**
1. Noter la **MAC et le S/N** du nouvel AP (photo de l'étiquette).
2. Monter le nouvel AP (même fixation, même orientation — les photos du dossier aident).
3. L'alimenter en PoE, attendre le démarrage.
4. **Mode Fat :** injecter la config sauvegardée :
```
# Via SFTP : déposer la sauvegarde, puis :
<AP761> restore configuration from vrpcfg-sauvegarde.zip  # (syntaxe à valider selon version)
# ou copier-coller les blocs de la config en mode system-view, section par section
<AP761> reboot
```
5. **Mode Cloud :** adopter le nouvel AP dans le site (QR code), lui donner le **même nom** (`AP761-Cour-Est`), la config descend toute seule.
6. **Mode Fit :** déclarer la MAC du nouvel AP sur l'AC (ou laisser l'auto-découverte si configurée).
7. Vérifier : SSID diffusés, client test, supervision verte.
8. Mettre à jour le dossier de site (nouveau S/N, date de remplacement).

**Pièges :**
- La sauvegarde contient parfois l'**ancienne adresse MAC** ou des bindings : vérifier après restauration.
- **Version firmware différente** entre la sauvegarde et le nouvel AP : restaurer d'abord le même firmware (chap. 75), puis la config.
- En cloud, ne pas oublier de **retirer l'ancien AP** du site (sinon : AP fantôme « hors ligne » qui pollue les alertes).

## 74. Firmware : cycle de vie et bonnes pratiques

Le firmware d'un AP extérieur n'est pas un gadget : il corrige des failles, ajoute des fonctions radio, et parfois **change des comportements** (d'où le rollback, chap. 76).

**Politique de mise à jour recommandée :**
| Fréquence | Quoi | Comment |
|---|---|---|
| **À chaque faille critique** | Patch de sécurité | Sous 30 jours, après test |
| **2× par an** | Version stable (patch releases) | Fenêtre planifiée, un site pilote d'abord |
| **Jamais en automatique aveugle** | — | Toujours valider sur un AP test ou un site non critique |

**Règles :**
1. **Lire la release note** avant : qu'est-ce qui change ? Qu'est-ce qui est corrigé ? Y a-t-il des prérequis (bootloader, version intermédiaire) ?
2. **Un site pilote d'abord** : mettre à jour UN AP (ou un petit site), attendre 1–2 semaines, puis généraliser.
3. **Fenêtre de maintenance** : jamais en pleine journée d'activité. L'AP redémarre (coupure 2–5 min par AP).
4. **Sauvegarder avant** (chap. 72) : config + version actuelle notée.
5. **Ne pas sauter plusieurs versions majeures** d'un coup : certaines migrations exigent des paliers — à vérifier sur la fiche du modèle exact et la release note.

**Où trouver les firmwares :** portail support Huawei / plateforme cloud eKit (qui propose les versions validées). **Ne jamais** flasher un firmware trouvé sur un forum : risque de version trafiquée ou de variante régionale incompatible.

## 75. Mise à jour firmware — procédure pas à pas

