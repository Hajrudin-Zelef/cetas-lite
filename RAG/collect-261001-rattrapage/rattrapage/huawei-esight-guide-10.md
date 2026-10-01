---
id: collect-261001-rattrapage/rattrapage/huawei-esight-guide-10
title: "Huawei eSight — Guide ultra-complet d'exploitation terrain"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["agent", "distribution", "incident"]
source: docs/RAG/collect-261001-rattrapage/huawei_esight_guide.md
source_anchor: ""
source_lines: [1454, 1628]
sha256: e9793e6955eace96286b2c047215e00d304d21f140a24d196b2fb102087fa892
---

# Huawei eSight — Guide ultra-complet d'exploitation terrain

- Des **interfaces northbound** existent pour l'intégration et le
  développement secondaire (détails : REST ? SOAP ? modèles de données ? →
  **à vérifier sur la documentation officielle** de votre version, les
  API évoluent).
- Cas d'usage : alimenter un portail interne, créer des tickets
  automatiquement dans l'ITSM, exporter des indicateurs vers une
  plateforme BI.
- **Conseil** : avant de développer, vérifiez ce que l'export standard
  (rapports programmés, traps northbound) couvre déjà — 80 % des besoins
  ne nécessitent aucun code.

## 75. Intégration avec l'ITSM (tickets)

Le couple gagnant : **alarme eSight → ticket automatique**.

- Via l'API northbound ou via e-mail parser côté ITSM (GLPI, etc.) :
  à la création d'une alarme critical/major, création d'un ticket avec
  le contexte (équipement, alarme, heure).
- À la résolution (clear), clôture ou commentaire automatique du ticket.
- **Ne pas tout ticketiser** : seules les sévérités critical/major
  (sinon le centre de tickets devient une poubelle).
- Bénéfice : traçabilité complète incident → alarme → action, et des
  KPI d'exploitation fiables (section 15).

## 76. Notifications SMS et passerelles

L'interface **SMS server** permet à eSight d'envoyer des SMS d'alarmes
via une passerelle (modem GSM, passerelle IP de l'opérateur…).

- Réserver le SMS aux **critical** et à l'astreinte (section 41).
- Tester mensuellement (changement de carte SIM, d'APN, de passerelle :
  le SMS est le canal qui tombe en panne en silence).
- Prévoir un **canal de secours** (appel vocal automatisé ou messagerie
  d'astreinte) : si eSight est down, il ne peut pas prévenir qu'il est down
  — d'où la supervision d'eSight lui-même (section 16).

## 77. eSight et eKit : que faut-il savoir ?

Huawei **eKit** est la marque/offre Huawei orientée distribution et PME
(équipements et outils simplifiés pour les petits réseaux). **Aucune
documentation publique consultée ne décrit une intégration native
eSight ↔ eKit** → **à vérifier sur la documentation officielle** et
auprès de votre partenaire.

En pratique :
- Si votre parc eKit est petit (< 20 équipements), l'édition **Compact**
  ou les outils cloud/Simplifiés associés à eKit peuvent suffire —
  un eSight complet serait surdimensionné.
- Si le parc grandit ou se mélange avec du Huawei entreprise classique,
  eSight redevient pertinent comme superviseur unifié.
- Ne présumez d'aucune compatibilité : **testez la supervision d'un
  équipement eKit depuis eSight en maquette avant d'acheter**.

## 78. Coexister avec d'autres outils (Zabbix, PRTG…)

eSight n'a pas besoin d'être seul :
- **Zabbix/Prometheus** pour la supervision fine des serveurs et
  applicatifs, **eSight** pour le réseau Huawei : complémentarité saine.
- Éviter les **doubles pollings agressifs** sur les mêmes équipements
  (deux NMS qui interrogent en SNMP toutes les 30 s = charge CPU
  inutile sur les petits switches).
- Répartir les rôles par écrit : quel outil est **référent** pour
  quelles alarmes (sinon deux équipes se renvoient la balle).

---

# 14. DÉPANNAGE : 17 CAS TERRAIN

> Format : **Symptômes → Diagnostic → Solution**. Les chemins de menus
> exacts dépendent de la version (**à vérifier sur la documentation
> officielle**).

## 79. Cas n°1 — Équipement non découvert

**Symptômes** : l'équipement n'apparaît pas dans l'inventaire après
la tâche de découverte, alors qu'il est bien dans la plage IP.

**Diagnostic :**
1. L'équipement répond-il au **ping** depuis le serveur eSight ?
   Non → problème réseau/routage/firewall avant même SNMP.
2. L'agent **SNMP est-il activé** sur l'équipement ? (`display snmp-agent`
   sur Huawei — la commande exacte dépend du modèle.)
3. La **community / l'utilisateur v3** correspond-elle au profil eSight ?
   Tester avec un outil externe (snmpwalk) depuis le serveur eSight :
   si snmpwalk échoue, le problème est côté équipement ou réseau,
   pas côté eSight.
4. Une **ACL** côté équipement bloque-t-elle l'IP d'eSight ?
5. Le **timeout/retry** du profil est-il adapté (WAN lent) ?

**Solution :** corriger le point bloquant (activer SNMP, aligner la
community, ouvrir l'ACL, augmenter le timeout), puis **relancer une
découverte ciblée** sur l'IP unique — pas tout le /16.

## 80. Cas n°2 — SNMP ne répond pas / timeouts en masse

**Symptômes** : des dizaines d'équipements passent « SNMP timeout »
simultanément.

**Diagnostic :**
1. **Tout le monde ou un segment ?** Tout le monde = problème côté
   eSight ou réseau central (lien, firewall). Un segment = problème
   local (lien montant, ACL).
2. Vérifier que l'**IP source** d'eSight n'a pas changé (nouvelle VM,
   nouvelle carte réseau — le cas classique après migration).
3. Vérifier les **ACL SNMP** des équipements : autorisent-elles toujours
   l'IP d'eSight ?
4. Charge du serveur eSight : un polling trop agressif + serveur
   sous-dimensionné = timeouts (voir cas n°13).
5. Un changement de **mot de passe SNMP v3** ou de community déployé
   en masse sans mettre à jour le profil eSight ?

**Solution :** selon la cause — corriger le routage/firewall, mettre à
jour l'IP dans les ACL (ou utiliser une IP virtuelle stable), aligner
les credentials, espacer le polling.

## 81. Cas n°3 — Les traps n'arrivent pas (alarmes en retard ou absentes)

**Symptômes** : eSight voit l'équipement en polling (up), mais les
pannes ne génèrent pas d'alarmes en temps réel — elles n'apparaissent
qu'au prochain cycle de polling, voire jamais.

**Diagnostic :**
1. Côté équipement : le **trap-target** pointe-t-il vers la bonne IP
   d'eSight (port UDP 162) ?
2. Un **firewall** filtre-t-il l'UDP 162 entre l'équipement et eSight ?
   (UDP = pas de connexion à tester au telnet ; utilisez un compteur
   ou une capture.)
3. Côté eSight : le service de réception des traps est-il démarré ?
4. L'équipement **génère-t-il** vraiment le trap ? Provoquer un événement
   de test (down/up d'un port de test) et observer.

**Solution :** corriger le trap-target / le firewall, redémarrer le
service de traps si nécessaire, puis valider avec un événement de test.
**Règle** : après chaque ajout d'équipement, le test trap fait partie
de la recette (section 31).

## 82. Cas n°4 — Tempête d'alarmes (storm)

**Symptômes** : des centaines d'alarmes en quelques minutes, console
inutilisable.

**Diagnostic :** identifier l'événement déclencheur (panne du cœur,
boucle réseau, flap d'un lien d'agrégation, redémarrage en cascade
après coupure électrique).

**Solution immédiate :**
1. Ne pas tout acquitter : **trier par heure** et identifier la
   première alarme (souvent la cause racine).
2. Appliquer la **corrélation** si configurée ; sinon, traiter manuellement
   la cause racine.
3. Si la tempête vient d'un équipement qui flap : **masquage temporaire**
   avec date de fin, le temps de réparer.

**Solution durable :** ajouter les règles de corrélation et d'agrégation
qui manquaient (sections 42, 47). Chaque tempête = une leçon documentée.

## 83. Cas n°5 — La topologie ne se met pas à jour

**Symptômes** : liens affichés up alors qu'ils sont down, équipements
fantômes, nouveaux liens absents.

**Diagnostic :**
1. La **découverte de topologie** (LLDPvoisinage) est-elle activée et
   planifiée ?
2. Le **LLDP est-il activé** sur les équipements (des deux côtés du lien) ?
3. Les traps linkUp/linkDown arrivent-ils (voir cas n°3) ?
4. Rafraîchissement manuel : forcer une resynchronisation de la vue.

**Solution :** activer LLDP partout (c'est aussi utile pour le
dépannage manuel), planifier la resynchronisation topologique
(quotidienne la nuit, par exemple), nettoyer les équipements
décommissionnés.

## 84. Cas n°6 — Fausse alerte « équipement down » (flapping de supervision)

