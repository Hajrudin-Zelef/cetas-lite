---
id: collect-261001-rattrapage/rattrapage/huawei-ekit-guide-21
title: "GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ekit_guide.md
source_anchor: ""
source_lines: [1780, 1852]
sha256: 5e473e523915bd54918776f37be685a2186e6ea7670cedcfdd2c000610f7a6be
---

# GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)

- **Puissance d'émission** : limitée par la réglementation locale (la fiche AR180 cite 23 dBm max, variable selon pays). Ne dépasse jamais les limites — amende + brouillage des voisins.
- **Bande 6 GHz** : pas autorisée partout — vérifier avant de déployer l'AP673H en 6 GHz (sinon : 2,4/5 GHz uniquement).
- **Données personnelles** : portail captif = collecte de données (e-mail, logs) → mentions légales, consentement, durée de conservation (selon droit local).
- **Vidéosurveillance** : déclaration/autorisation selon le pays + information des personnes filmées (panneaux).
- **Sécurité électrique** : la baie et l'onduleur suivent les normes locales (un disjoncteur dédié pour la baie, c'est la base).
En cas de doute : le distributeur local et un juriste valent mieux qu'une amende.

## 164. Les 10 questions que les clients te poseront (et tes réponses)

1. **« C'est chinois, c'est fiable ? »** → « C'est du matériel pro avec 3 ans de garantie via notre distributeur local, déployé dans 70+ pays. La fiabilité se joue surtout à l'installation et à la maintenance — c'est notre métier. »
2. **« Et si Internet tombe ? »** → Section 38 : le réseau local continue, seule la gestion cloud est aveugle.
3. **« Le cloud, c'est gratuit vraiment ? »** → Section 33 : gestion annoncée sans licence ; je te détaille ce qui est inclus et les options.
4. **« On pourra changer de prestataire ? »** → Section 103 : oui, réversibilité prévue au contrat dès le départ.
5. **« Ça capte jusqu'où ? »** → « On dimensionne par pièce/zone avec un AP par X m², et on valide par des mesures — pas des promesses. »
6. **« Pourquoi pas juste des répéteurs ? »** → « Un répéteur divise le débit par 2 à chaque saut et crée des coupures. Des AP câblés, c'est le seul vrai Wi-Fi pro. »
7. **« C'est sécurisé ? »** → Section 104 : VLAN, WPA2/3, isolation invités, pare-feu en option — et je te montre l'audit annuel.
8. **« Combien de temps pour installer ? »** → « Une boutique : 1-2 jours. Un hôtel 40 chambres : 1-2 semaines avec le câblage. Je te donne un planning écrit. »
9. **« Et dans 3 ans quand on grandit ? »** → Section 83-85 : l'architecture est prévue pour grandir, avec une trajectoire écrite.
10. **« Pourquoi vous et pas un autre ? »** → « Pour la méthode : visite technique, plan écrit, PV de recette, maintenance avec KPI. Le matériel, tout le monde le vend — le service, c'est nous. »

## 165. Index et fin

**Index rapide :** adressage (64) · app eKit (16-30) · audit (155) · budgets (148) · canaux Wi-Fi (146) · chiffrer (159) · cloud/SNC (31-45) · comparatif concurrents (149) · Datacom : bascule (83-90) · dépannage méthode (91) · dépannage cas 1-16 (92-100) · dépannage cas 17-25 (156) · équipe : formation (158) · fiches modèles (116-130) · garantie/RMA (88, 151) · glossaire (110) · hôtel (50-54) · interopérabilité (61) · iGuard (29, 121) · LED (161) · maintenance plan (101) · migration existant (62-63) · migration Datacom (84) · MiniFTTO (8, 52) · multi-sites (35) · onduleur/énergie (48, 72) · PoE (68-69) · portail captif (40, 154) · PRA (157) · PV de recette (160) · quiz (111-112) · réglementaire (163) · reset (150) · sauvegarde (89) · scénarios A-D (46-60) · cas pratiques 131-145 · sécurité (102-104) · sites : organisation (18) · SNC (31) · supervision (76-80) · USG (7, 130) · VLAN (66) · SSID (67).
**Ce guide est un document vivant** : à chaque déploiement, ajoute ton cas n°26, ta fiche modèle, ton retour distributeur. Dans un an, ce sera l'actif le plus précieux de ton service.
**Fin du guide — bon déploiement.**

---

## 166. DHCP : le régler comme un pro (et éviter 50 % des tickets)

- **Un scope par VLAN**, avec la passerelle du VLAN comme routeur par défaut et des DNS valides. L'erreur n°1 : un scope sans DNS ou avec un DNS injoignable → « Internet ne marche pas » alors que l'IP est bonne.
- **Baux** : staff/bureaux = 24 h ; invités/commerce = 4-8 h ; hôtel chambres = 12-24 h ; résidence longue = 7 jours (section 145). Un bail trop long sur un VLAN à fort turnover = scope épuisé (cas n°18).
- **Exclusions** : toutes les IP fixes d'infrastructure **hors** du scope (passerelle .1, switchs .2-.9, AP .10-.49, imprimantes, NVR). Le conflit d'IP (cas n°13) vient toujours d'ici.
- **Réservations** : pour les équipements qui doivent garder leur IP (imprimante, NVR, TPE) sans IP fixe en dur — le meilleur des deux mondes.
- **Option 43 / vendor-specific** : si les AP ont besoin d'options DHCP particulières pour trouver leur contrôleur/cloud — **à vérifier sur la documentation officielle** par modèle (en mode cloud pur, l'onboarding se fait via l'app, pas via DHCP).
- **Surveillance** : alerte « scope > 80 % » si le NMS le permet — sinon, contrôle mensuel (section 101).

## 167. DNS : le maillon invisible

- Si les clients ont une IP mais « rien ne s'affiche » : c'est **presque toujours le DNS** (cas n°5, diagnostic n°3).
- Sur l'AR : DNS de l'opérateur en primaire + un DNS public en secondaire (ex. fictifs : 1.1.1.1, 8.8.8.8 — à adapter selon politique locale).
- **Ne jamais** laisser un DNS interne injoignable comme DNS unique d'un VLAN invités.
- Portail captif : le DNS du VLAN invités doit résoudre **vite** — la détection captive en dépend (cas n°7).
- Test de validation : `nslookup` (ou équivalent) d'un nom externe depuis un client de chaque VLAN, le jour J et à chaque audit.

## 168. NAT/PAT sur la passerelle AR : l'essentiel

- En PME, la passerelle fait du **NAT/PAT** (masquerading) : tout le LAN sort avec l'IP publique du WAN. C'est le mode par défaut — vérifie juste qu'il est actif.
- **Redirection de ports** (port forwarding) : à éviter autant que possible (chaque port ouvert = une surface d'attaque). Si indispensable (caméras accessibles à distance, serveur local) : règle stricte, IP source limitée si possible, et **documentée** dans la fiche site.
- Alternative propre : **VPN** (IPsec/SSL sur AR/USG) pour l'accès distant — jamais d'admin exposée sur Internet.
- **UPnP** : désactivé par défaut en pro (un malware qui ouvre ses propres ports via UPnP, c'est le scénario classique).

## 169. Routage inter-VLAN : qui parle à qui

- Par défaut sur l'AR : les VLAN se voient (routage inter-VLAN actif). **Ce n'est pas ce que tu veux** pour les invités.
- Matrice type à appliquer :
  - GUESTS → Internet : oui ; → autres VLAN : **non**.
  - STAFF → Internet : oui ; → CAMERAS : non (sauf NVR) ; → VOIX : oui si besoin.
  - CAMERAS → NVR : oui ; → Internet : non (sauf accès distant via VPN).
  - MGMT → tout (admin) ; ← tout : **non** (personne n'administre depuis les invités).
- Implémentation : règles pare-feu/ACL sur l'AR ou l'USG — le niveau d'exposition dans l'app varie (**à vérifier sur la documentation officielle** ; sinon via l'interface locale).
- **Teste la matrice** le jour J avec des pings croisés — une règle oubliée = un VLAN invités qui voit les caméras.

## 170. QoS pratique : prioriser sans se noyer

- En PME, 3 classes suffisent : **voix/temps réel** (priorité max), **bureautique** (normal), **invités** (dégradé, avec plafond).
- Marque la voix en **DSCP EF** si les téléphones le font (la plupart oui) et fais confiance au marquage bout-en-bout sur le LAN.
- Le levier le plus efficace en PME n'est pas la QoS fine : c'est le **plafond de débit** sur le VLAN invités (10-20 Mbit/s partagés ou par client — valeurs fictives d'exemple à ajuster).
- Streaming/régie (sections 137) : VLAN dédié + priorité — et surtout pas de MAJ pendant l'événement.
- Ne promets jamais « la QoS réglera votre débit » si l'uplink est saturé : la QoS **arbitre** la pénurie, elle ne crée pas du débit.

## 171. 802.1X / NAC : ce que eKit peut et ne peut pas

