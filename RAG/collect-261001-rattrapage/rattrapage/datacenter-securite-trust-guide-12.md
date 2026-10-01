---
id: collect-261001-rattrapage/rattrapage/datacenter-securite-trust-guide-12
title: "Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés"
domain: rattrapage
role: reference
task: reference
actors: ["AWS", "EU"]
dates: []
keywords: ["attention", "aws", "distribution"]
source: docs/RAG/collect-261001-rattrapage/datacenter_securite_trust_guide.md
source_anchor: ""
source_lines: [1379, 1525]
sha256: 16ff3ac37fafb5d810da17eba5216646f6254a3ff420216b199300489f784564
---

# Sécurité datacenter — HSM, Root of Trust, FPGA, sécurité physique, gestion des clés

| Niveau | Méthode | Quand l'utiliser |
|---|---|---|
| **Clear** | Écrasement logique (1+ passes) | Réutilisation interne, disque sain |
| **Purge** | Commandes firmware (Secure Erase, crypto-erase), démagnétisation | Sortie du parc, disques défectueux mais lisibles |
| **Destroy** | Broyage, désintégration, incinération | Données top critiques, disques HS, fin de vie certifiée |

Procédure : **inventaire** (n° de série), **méthode tracée** par support, **PV de
destruction** signé, prestataire certifié pour le broyage (avec attestation). Les
**bandes de sauvegarde** suivent le même régime — une bande « jetée » lisible = fuite de
données.

## 101. Disques chiffrés (SED) et crypto-erase

Les disques **SED** (Self-Encrypting Drives, norme TCG Opal) chiffrent en permanence avec
une clé interne : la mise au rebut devient un **crypto-erase** (destruction de la clé =
données irrécupérables en millisecondes). Avantages : instantané, pas d'usure, fonctionne
même sur disque partiellement défectueux. Conditions : **activer le chiffrement dès le
déploiement** (un SED non activé = disque clair), gérer les mots de passe/clés (via le
contrôleur ou un KMIP/KMS — voir partie 5), et vérifier la conformité du firmware (des
failles ont existé sur des implémentations SED — patcher).

## 102. Déclassement de matériel : procédure

Checklist de sortie de parc pour tout serveur/équipement :

- [ ] Sauvegarde des données utiles (si réutilisation prévue ailleurs).
- [ ] **Effacement** selon NIST 800-88 (section 100) — niveau adapté à la sensibilité.
- [ ] Retrait et destruction séparée des **modules TPM/HSM** le cas échéant.
- [ ] Retrait des étiquettes d'inventaire sensibles, des configurations (BMC reset).
- [ ] PV de déclassement archivé (n° de série, méthode, date, opérateur).
- [ ] Si revente/don : uniquement après Purge minimum + contrôle.

## 103. Sécurité électrique et incendie : le lien énergie

La sécurité physique inclut l'énergie (métier de Zelef — voir partie 8 en détail) :

- **Incendie** : détection très précoce (aspiration), extinction **gaz** (inerte, sans
  eau) en salle serveurs, compartimentage coupe-feu, exercice d'évacuation annuel.
- **Dégât des eaux** : détection de fuite sous plancher technique, pas de tuyauterie
  d'eau au-dessus des baies.
- **Électrique** : double alimentation des baies critiques, onduleurs **secourus**
  (section 140), parafoudres, terre < 5 Ω.
- **Clim** : redondance N+1 minimum, sondes de température avec alertes, délestage
  automatique en cas de surchauffe.

## 104. Tableau : équipements physiques et budgets indicatifs

| Équipement | Ordre de grandeur | Remarque |
|---|---|---|
| Contrôle d'accès (lecteurs + logiciel, 5 portes) | 5–15 k€ | MIFARE DESFire, anti-passback |
| Biométrie (1 point, second facteur) | 1–3 k€ | En complément du badge |
| Vidéosurveillance (8 caméras + NVR 90 j) | 5–12 k€ | Rétention = coût de stockage |
| Alarme intrusion + télésurveillance | 2–5 k€ + abonnement | Tests trimestriels |
| Cage grillagée (10 m²) | 3–8 k€ | En salle mutualisée |
| Coffre-fort ignifuge (clés, HSM backup) | 1–4 k€ | Classe selon valeur protégée |
| Destruction (broyage, par passage) | 500–2 000 € | Prestataire certifié + PV |

Budgets **indicatifs France/EU 2026, à vérifier** sur devis — mais l'ordre de grandeur
compte : la sécurité physique d'une petite salle se chiffre en **dizaines de k€**, pas en
centaines.

# PARTIE 5 — GESTION DES CLÉS / KMS

## 105. Gestion des clés : le cycle de vie complet

Une clé n'est pas un fichier qu'on pose quelque part : c'est un **actif avec un cycle de
vie** :

```
 Génération (HSM/TRNG) → Distribution → Utilisation → Rotation → Archivage → Destruction
        │                     │              │            │            │            │
   cérémonie MofN        canal chiffré   journalisée   planifiée   chiffrée    certifiée
   (section 42)          (jamais clair)  (section 120) (section 108)(durée légale)(section 100)
```

Chaque transition est **documentée** : qui a fait quoi, quand, avec quelle autorisation.
Une clé sans cycle de vie géré, c'est une clé dont on ne connaît ni l'âge, ni les copies,
ni le moment où elle fuira.

## 106. Hiérarchie : KEK, DEK, clés maîtres

- **DEK** (Data Encryption Key) : chiffre les données. Nombreuses, à vie courte, stockées
  **chiffrées**.
- **KEK** (Key Encryption Key) : chiffre les DEK. Peu nombreuses, à vie longue, stockées
  dans le **HSM/KMS**, jamais en clair hors de lui.
- **Clé maîtresse / root key** : au sommet de la hiérarchie, générée en cérémonie,
  sauvegardée en MofN. Elle ne chiffre jamais de données directement.

Principe : **compromettre une DEK n'expose qu'un jeu de données** ; compromettre la KEK
expose tout ce qu'elle protège ; compromettre la maîtresse = tout reconstruire. D'où la
protection croissante à chaque niveau.

## 107. Envelope encryption : le pattern central

Le pattern utilisé par AWS KMS, GCP, Vault :

1. Le KMS génère une **DEK** (clé de données) via le HSM.
2. La DEK est **chiffrée par la KEK** → on obtient la DEK « enveloppée ».
3. Les données sont chiffrées avec la DEK **en clair en mémoire uniquement**.
4. On stocke : données chiffrées + DEK enveloppée. La DEK en clair est **oubliée**.

Pour déchiffrer : le KMS désenveloppe la DEK (opération HSM), déchiffre les données.
Avantages : la KEK ne voyage jamais, les DEK peuvent être nombreuses et rotées souvent,
le volume chiffré par une seule clé reste limité. **C'est le pattern à implémenter par
défaut** pour tout chiffrement applicatif maison.

## 108. Rotation : politiques et automatisation

La rotation limite l'exposition d'une clé dans le temps :

| Type de clé | Rotation recommandée | Notes |
|---|---|---|
| Clé TLS serveur | 90 jours (ou moins, ACME) | Automatisable à 100 % |
| DEK applicatives | 1 an ou par volume (ex. tous les 2^32 chiffrements) | Re-chiffrement ou double enveloppe |
| KEK | 2-3 ans | Opération planifiée, pas d'urgence |
| Clé d'AC intermédiaire | 3-5 ans | Planifiée avec la cérémonie |
| Clé d'AC racine | 10-20 ans (ou jamais, avec rollover) | Événement majeur, cérémonie complète |
| Clé de chiffrement de sauvegarde | 1-2 ans | Attention aux sauvegardes anciennes |

Règle : **automatiser** tout ce qui est automatisable (TLS, DEK), **planifier et
documenter** le reste. Une rotation manuelle « quand on y pense » n'arrive jamais.

## 109. Durées de vie recommandées par type de clé (détail)

Complément du tableau 108 — les justifications :

- **TLS 90 jours** : compromis entre charge opérationnelle et fenêtre d'exposition ;
  avec ACME/Let's Encrypt interne, le coût marginal est nul.
- **DEK 1 an** : limite le volume de données sous une même clé (principe de
  compartimentation) ; le re-chiffrement peut être progressif (lazy re-encryption).
- **KEK 2-3 ans** : la rotation d'une KEK = re-envelopper les DEK, opération lourde
  mais planifiable ; au-delà de 3 ans, le risque d'exposition silencieuse augmente.
- **Racine 10-20 ans** : sa rotation invalide toute la chaîne — on ne la fait que si
  compromise ou par obsolescence d'algorithme (ex. migration PQC, voir partie 9).

## 110. Séparation des rôles : qui peut quoi

Matrice type (à adapter) :

| Rôle | Générer | Utiliser | Sauvegarder | Détruire | Auditer |
|---|---|---|---|---|---|
| Crypto Officer | ✅ (cérémonie) | ✅ | ✅ | ❌ | ❌ |
| Opérateur applicatif | ❌ | ✅ (via API) | ❌ | ❌ | ❌ |
| Custodian (MofN) | ❌ | ❌ | ✅ (sa part) | ❌ | ❌ |
| Auditeur | ❌ | ❌ | ❌ | ❌ | ✅ |
| Admin système | ❌ | ❌ | ❌ | ❌ | ✅ (logs) |

