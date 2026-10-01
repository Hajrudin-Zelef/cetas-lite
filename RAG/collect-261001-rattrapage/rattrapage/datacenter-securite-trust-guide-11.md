---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-11
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "dpo"]
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [1261, 1378]
sha256: f88a4ee0a70192e2d2f25e8537fd9e2c02f418cbc448b76dc7451d4513b17219
---

# Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés

```
  ┌──────────────────────────────────────────────────────────────┐
  │ CERCLE 1 — Périmètre : clôture, portail, éclairage, rondes    │
  │ ┌──────────────────────────────────────────────────────────┐ │
  │ │ CERCLE 2 — Bâtiment : sas d'entrée, accueil, badges       │ │
  │ │ ┌──────────────────────────────────────────────────────┐ │ │
  │ │ │ CERCLE 3 — Salle : contrôle d'accès biométrique,      │ │ │
  │ │ │ vidéo, anti-tailgating                                │ │ │
  │ │ │ ┌──────────────────────────────────────────────────┐ │ │ │
  │ │ │ │ CERCLE 4 — Rack : serrure, cage client, scellés   │ │ │ │
  │ │ │ │ HSM + coffre des sauvegardes                     │ │ │ │
  │ │ │ └──────────────────────────────────────────────────┘ │ │ │
  │ │ └──────────────────────────────────────────────────────┘ │ │
  │ └──────────────────────────────────────────────────────────┘ │
  └──────────────────────────────────────────────────────────────┘
```

Chaque cercle **retarde** et **détecte**. Aucun n'est infaillible seul : la sécurité physique
est un **empilement de délais** (le temps de franchir > le temps d'intervention).

## 92. Périmètre : bâtiment, clôture, sas

- **Clôture** : 2 m minimum, anti-escalade, avec détection (câble périmétrique ou vidéo).
- **Éclairage** : périmètre éclairé la nuit (dissuasion + efficacité des caméras).
- **Portail** : contrôle d'accès véhicules, barrières, registre des immatriculations.
- **Sas d'entrée** : double porte à interverrouillage (la 2e ne s'ouvre que si la 1re est
  fermée) — anti-tailgating dès l'entrée du bâtiment.
- Pour une PME sans datacenter propre : ces exigences se **délèguent à l'hébergeur**
  (demander ses certifications et son rapport d'audit — voir section 126).

## 93. Contrôle d'accès : badges, biométrie, anti-passback

- **Badges** : technologie **MIFARE DESFire** ou équivalent (ne pas utiliser les vieux
  MIFARE Classic, clonables en minutes). Un badge = une personne, jamais de badge
  « générique ».
- **Biométrie** : empreinte ou veine du doigt en **second facteur** pour les salles
  critiques (jamais seule : on ne « révoque » pas un doigt).
- **Anti-passback** : le système refuse une 2e entrée sans sortie enregistrée — tue le
  prêt de badge.
- **Zonage** : chaque badge n'ouvre que les zones nécessaires à sa fonction (principe du
  moindre privilège physique).
- **Révocation** : procédure de désactivation **immédiate** en cas de départ/perte
  (SLA < 1 h).

## 94. Cages, racks, serrures : le cloisonnement

- **Cages grillagées** : dans une salle mutualisée, chaque client a sa cage fermée à clé —
  vos HSM et serveurs critiques n'y sont accessibles qu'à vos habilités.
- **Racks verrouillés** : serrures à clé ou à code **par rack**, pas une clé unique pour
  toute la salle.
- **Scellés** : sur les serveurs sensibles (HSM, serveurs PKI), scellés numérotés dont
  l'état est contrôlé à chaque intervention.
- **Baies aveugles** : panneaux obturateurs sur les U vides (sécurité + airflow —
  double bénéfice, voir section 142).

## 95. Vidéosurveillance : caméras, enregistrement, durées

- **Couverture** : entrées/sorties, couloirs, allées de baies, coffre-fort, zone de
  réception — **pas de zones aveugles** sur les chemins critiques.
- **Qualité** : résolution permettant l'identification (pas juste « une silhouette »),
  infrarouge ou éclairage pour la nuit.
- **Enregistrement** : **90 jours minimum** de rétention pour les zones critiques
  (standard d'audit courant), stockage **protégé** (local sécurisé + réplication).
- **Supervision** : affichage en temps réel à l'accueil/sécurité + détection de
  mouvement avec alertes hors heures ouvrées.
- **Conformité** : affichage réglementaire (information des personnes filmées), registre
  des accès aux enregistrements — pas de conseil juridique ici, mais à cadrer avec le
  DPO/juriste (voir section 128).

## 96. Détection intrusion : alarmes, capteurs

- **Alarme volumétrique** : détecteurs de mouvement dans les salles hors présence
  permanente.
- **Contacts d'ouverture** : portes, baies sensibles, coffre.
- **Report** : vers un centre de télésurveillance ou l'astreinte, avec **levée de doute**
  vidéo.
- **Tests** : déclenchement testé **trimestriellement** (une alarme jamais testée est une
  décoration).
- **Couplage** : intrusion + coupure électrique simultanées = scénario d'attaque à
  intégrer au plan de réponse (voir section 136).

## 97. Gestion des visiteurs : procédure complète

1. **Pré-enregistrement** : visite annoncée, motif, zones autorisées, accompagnateur désigné.
2. **Accueil** : pièce d'identité vérifiée, badge **visiteur** distinct (couleur
   différente), registre signé (papier ou électronique, conservé 1 an minimum).
3. **Escorte permanente** : aucun visiteur seul en salle technique — **sans exception**,
   y compris les prestataires connus.
4. **Restitution** : badge rendu, sortie enregistrée, vérification qu'il ne reste rien
   (outils, documents).
5. **Prestataires récurrents** (maintenance clim, électriciens) : même procédure, avec en
   plus une **charte de confidentialité** signée et une sensibilisation aux zones
   interdites (photos interdites en salle).

## 98. Politique anti-tailgating et escortes

Le **tailgating** (suivre quelqu'un qui badge) est le vecteur n°1 d'intrusion physique.
Contre-mesures : sas à interverrouillage, tourniquets pleine hauteur aux entrées
critiques, **culture d'entreprise** (« je ne tiens pas la porte, c'est la procédure » —
à faire valider par la direction pour que personne ne se sente malpoli), et
**rappels réguliers**. Un test d'intrusion physique annuel (auditeur qui tente d'entrer)
est le meilleur indicateur.

## 99. Personnel : habilitations, séparation des tâches

- **Habilitation** : l'accès aux salles critiques est nominatif, revu **semestriellement**
  (les droits s'accumulent avec les changements de poste — revue obligatoire).
- **Séparation** : celui qui administre les HSM n'est pas celui qui contrôle les accès
  physiques ; l'auditeur n'est pas l'exploitant (même logique qu'en section 44).
- **Départs** : checklist de sortie incluant la révocation des accès physiques **le jour
  même**.
- **Sensibilisation** : phishing, tailgating, photos en salle — 1 session/an minimum,
  traçée.

## 100. Destruction de médias : NIST 800-88

La norme **NIST SP 800-88** définit 3 niveaux de sanitization :

