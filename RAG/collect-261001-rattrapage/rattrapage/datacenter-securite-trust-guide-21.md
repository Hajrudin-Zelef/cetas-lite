---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-21
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "arr", "gpu"]
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [2641, 2796]
sha256: b2bece9266c171fba097afa96f3ca0e533a087538297a183ba6a9e203e6c3813
---

# Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés

1. **Faut-il savoir coder en VHDL ?** Pour un POC : non (HLS + exemples). Pour de la
   production optimisée : oui, ou un prestataire.
2. **Un FPGA remplace-t-il un GPU pour l'IA ?** Non pour l'entraînement et les LLM ;
   oui pour de l'inférence **figée** à latence garantie (vision, scoring).
3. **Durée de vie d'une carte ?** 5-7 ans (support outils, disponibilité) ; le design
   est portable vers la génération suivante avec effort.
4. **Le FPGA consomme-t-il moins qu'un GPU ?** En **perf/watt sur workload fixe** :
   souvent oui. En absolu : 190 W reste 190 W à évacuer.
5. **Peut-on louer du FPGA ?** Oui (instances cloud F1/FPGA chez les hyperscalers) —
   bien pour le POC, à chiffrer vs achat au-delà de 6 mois.
6. **Bitstream et sécurité ?** Voir section 178 — à traiter comme un firmware critique.
7. **FPGA + virtualisation ?** PCI passthrough vers une VM dédiée ; le partage fin
   (vFPGA) reste un sujet de recherche/édition entreprise.
8. **Quel est le premier pas ?** Une carte d'occasion U250 (~1-2 k€) + Vivado + un
   tutoriel packet-processing : 1 mois pour savoir si votre équipe accroche.

# PARTIE 13 — APPROFONDISSEMENTS KMS / PHYSIQUE : LA MÉTHODE

## 184. Modèle de menaces : l'atelier en 2 heures

Avant d'acheter, **modéliser** : réunir 4-6 personnes (infra, sécu, métier, énergie)
et répondre par écrit :

1. **Actifs** : quelles clés ? quelles données ? (inventaire, section 119 n°9)
2. **Adversaires** : voleur opportuniste ? concurrent ? étatique ? insider ?
   (le niveau d'adversaire dimensionne tout — un HSM FIPS L3 contre un voleur de
   disques, c'est cohérent ; contre un étatique, c'est une couche parmi d'autres)
3. **Vecteurs** : vol physique, accès réseau, insider, fournisseur, panne ?
4. **Impacts** : chiffrer en euros l'heure d'arrêt et la fuite (section 146).
5. **Mesures existantes** : que couvre-t-on déjà ?
6. **Écarts** : la liste d'achats qui en découle — **priorisée par risque**, pas par
   catalogue vendeur.

Livrable : 2 pages. Refaire **annuellement** ou à chaque changement majeur. C'est
l'exercice qui évite d'acheter un HSM à 30 k€ quand le vrai risque est le mot de passe
BMC par défaut (section 67).

## 185. Vault : politiques et durcissement (exemples)

Principes (HashiCorp Vault, transposables) :

```
# Politique : une appli ne lit que ses secrets
path "secret/data/appli-paiement/*" {
  capabilities = ["read"]
}
# Pas de list sur la racine, pas d'accès aux secrets des autres applis
```

Durcissement :

- **Auto-unseal** sur HSM (section 113), jamais de Shamir manuel en production
  automatisée.
- Backend **Raft** chiffré, snapshots **chiffrés** et testés.
- **Audit device** activé (fichier + syslog), non désactivable par les opérateurs.
- Authentification : **AppRole** ou OIDC pour les applis, jamais de token root qui
  traîne (le token root sert à l'initialisation puis est **révoqué**).
- Rotation des **credentials dynamiques** (bases de données) : durée de vie courte
  (heures), pas de mots de passe statiques.

## 186. Rotation automatisée des DEK : pattern

```
1. Générer DEK_n+1 dans le KMS/HSM
2. Re-envelopper : chiffrer DEK_n+1 avec la KEK → stocker
3. Double lecture : l'appli accepte DEK_n (déchiffrement ancien) et chiffre en DEK_n+1
4. Migration : re-chiffrer progressivement (ou au fil de l'eau / lazy)
5. Retirer DEK_n de la lecture après la fenêtre de migration
6. Archiver DEK_n (chiffrée) pour les anciennes sauvegardes, puis détruire à échéance
```

Automatisable à 90 % (cron + API KMS). La fenêtre de double lecture (étape 3) est le
point critique : trop courte = données illisibles, trop longue = exposition prolongée.
La documenter dans la politique de rotation (section 108).

## 187. Sauvegarde 3-2-1 appliquée aux clés

La règle **3-2-1** (3 copies, 2 supports différents, 1 hors site) adaptée aux clés :

- **Copie 1** : HSM de production (opérationnelle).
- **Copie 2** : HSM backup sur site (réplication chiffrée, testée).
- **Copie 3** : sauvegarde chiffrée **hors site** (coffre distant ou HSM du site de
  secours) — la KEK de cette sauvegarde en **MofN** (section 43).
- **Jamais** : de copie en clair, même « temporaire », même « chez le prestataire de
  confiance ».

Tester la restauration de la copie 3 **une fois par an** (c'est elle qui sauve en cas
de sinistre majeur — section 138).

## 188. Plan de câblage de la zone crypto

- **Courant fort** : PDU A et B sur chemins séparés, étiquetage par couleur
  (ex. rouge = A, bleu = B), jamais de multiprise sauvage.
- **Courant faible** : VLAN crypto sur switch dédié ou ports dédiés, jarretières
  étiquetées, longueur adaptée (pas de boucles de 3 m pour 30 cm).
- **Séparation** : goulottes séparées fort/faible, croisements à 90°.
- **Documentation** : plan de brassage à jour, chaque câble numéroté aux deux bouts.
- **Sécurité** : aucun câble du VLAN crypto ne sort de la baie sans passer par le
  firewall documenté (section 35) ; les ports non utilisés du switch sont
  **désactivés**.

## 189. Journalisation : architecture SIEM pour la crypto

```
[HSM] ──syslog/TLS──┐
[Vault] ──audit──┐    │
[BMC] ──syslog──┤    ├─▶ [Collecteur] ─▶ [SIEM] ─▶ alertes + tableaux de bord
[Firewall] ─────┘    │         │
[Badgeuse] ──────────┘         └─▶ [Stockage WORM / immuable]
```

Règles :

- Transport **chiffré** (TLS) et **horodatage fiable** (NTP/NTS sur tous les
  équipements — sans temps fiable, la corrélation est impossible).
- Stockage **immuable** (WORM) pour les journaux d'audit : un attaquant ne doit pas
  pouvoir effacer ses traces.
- **Alertes** : échec d'auth HSM ×5, export de clé, accès physique hors heures,
  dérive d'horloge > 1 s.
- **Revue** : quotidienne automatisée (règles), hebdomadaire humaine (échantillon).

## 190. Exercice de crise : scénario « clé compromise »

Déroulé d'un exercice annuel (2 h, table ronde + technique) :

1. **Injection** : « le SIEM remonte un export de la KEK sauvegardes à 3 h du matin
   depuis le poste d'un admin en congés ».
2. **Décisions** : qui est convoqué ? (cellule de crise : RSSI, infra, direction,
   juriste) — dans quel ordre ?
3. **Actions** : confinement (section 137), périmètre, communication interne/externe.
4. **Technique** : révoquer, régénérer, re-chiffrer — **chronométrer** chaque étape.
5. **Débrief** : ce qui a coincé (mot de passe du coffre introuvable ? quorum
   injoignable un dimanche ?) → plan d'action.

C'est l'exercice le plus rentable de l'année : il révèle les failles
**organisationnelles** que la technique ne voit pas.

## 191. FAQ KMS (8 questions)

1. **Vault remplace-t-il un HSM ?** Non — il le **complète** (section 111). Vault sans
   HSM = clé maîtresse logicielle = point faible.
2. **Combien de Vault ?** 1 cluster (3 nœuds Raft) suffit pour une ETI ; séparer par
   environnement (prod/preprod) si les exigences divergent.
3. **Où stocker les secrets des conteneurs ?** Via l'agent Vault / CSI driver, jamais
   en variables d'environnement en clair ni dans les images.
4. **BYOK ou HYOK ?** (section 116) — BYOK par défaut, HYOK pour le top critique.
5. **Faut-il chiffrer les sauvegardes Vault ?** Oui, et tester leur restauration
   (section 138) — une sauvegarde Vault contient tous les secrets.
6. **Que faire des anciens secrets ?** Rotation puis **révocation** : un secret
   « ancien mais toujours valide » est une bombe à retardement.
7. **KMS cloud ou on-prem ?** Cloud si les workloads sont cloud (latence, intégration) ;
   on-prem si souveraineté ou workloads locaux — souvent les deux (hybride).
8. **Le post-quantique change-t-il le KMS ?** Le KMS doit devenir **crypto-agile**
   (section 153) : stocker l'identifiant d'algorithme avec chaque clé, prévoir la
   migration ML-KEM/ML-DSA.

