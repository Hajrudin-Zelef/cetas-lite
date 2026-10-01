---
id: collect-261001-rattrapage/rattrapage/maintenance-windows-guide-3
title: "Maintenance et exploitation Windows en entreprise"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["exploit"]
source: docs/RAG/collect-261001-rattrapage/maintenance_windows_guide.md
source_anchor: ""
source_lines: [349, 520]
sha256: 92f49b3f785b4e30512b861d693a81324a1bc9e293d27e2dfcd29866df404ca3
---

# Maintenance et exploitation Windows en entreprise

- [ ] Inventaire complet : rapprochement CMDB ↔ AD ↔ supervision.
- [ ] Test de restauration PRA **partiel** (un serveur, un scénario).
- [ ] Revue des droits : groupes sensibles AD, administrateurs locaux (LAPS).
- [ ] Mise à jour des images de référence / séquences de déploiement.
- [ ] Nettoyage WSUS (§15), purge des journaux archivés.
- [ ] Revue des contrats : garanties, supports, licences.

**Annuelle :**

- [ ] Test PRA **complet** (§127).
- [ ] Audit de sécurité AD (mots de passe, Kerberoasting, délégation).
- [ ] Plan de renouvellement matériel (serveurs > 5 ans).
- [ ] Bilan d'exploitation : MTTR, disponibilité, incidents majeurs — présenté
      à la direction.
- [ ] Mise à jour du DTI (§128) et du plan de maintenance (§4).

---

## 10. Patch Tuesday : principe et calendrier

Microsoft publie les mises à jour de sécurité le **2ᵉ mardi de chaque mois**
(« Patch Tuesday », vers 19h heure de Paris). Le cycle d'exploitation mensuel type :

| Jour | Action |
|---|---|
| J (mardi) | Publication. Lecture des notes, repérage des correctifs critiques / 0-day. |
| J+1 → J+7 | Déploiement anneau **test** (§11), validation applicative. |
| J+7 → J+14 | Déploiement anneau **pilote** (échantillon représentatif). |
| J+14 → J+21 | Déploiement **production**, par vagues, dans les fenêtres (§19). |
| J+21 → J+30 | Rattrapage des non-conformes, rapport de conformité (§8). |

**Hors cycle** : les correctifs **hors bande** (0-day exploité, ex. PrintNightmare)
suivent une procédure accélérée : test réduit → pilote restreint → production sous
48-72 h, avec validation du changement exceptionnel (§130).

Types de mises à jour à connaître :

| Type | Contenu | Redémarrage |
|---|---|---|
| Correctif cumulatif mensuel (B) | Sécurité + qualité cumulées | Oui (souvent) |
| Correctif hors bande (OOB) | Urgence sécurité | Oui |
| Correctif .NET | Sécurité .NET Framework | Oui |
| Mise à jour de la pile de maintenance (SSU) | Fiabilise Windows Update lui-même | Parfois |
| Pilotes | Via WU (optionnel) | Variable |

> Règle : ne jamais « sauter » un cumulatif pour gagner du temps : ils sont cumulatifs,
> le mois suivant inclut le précédent. Un serveur à jour en N-1 se met à jour en N
> sans étape intermédiaire.

---

## 11. Anneaux de déploiement : test, pilote, production

| Anneau | Contenu | Délai après Patch Tuesday | Objectif |
|---|---|---|---|
| **Test** | 2-5 machines par OS/rôle (VM de labo) | J+1 | Détecter les régressions grossières |
| **Pilote** | 5-10 % du parc, utilisateurs volontaires, tous métiers | J+7 | Valider en conditions réelles |
| **Production** | Reste du parc, par vagues (serveurs non critiques → critiques) | J+14 | Conformité ≥ 98 % |

**Serveurs** : ne jamais patcher tous les membres d'un cluster / tous les DC la même
nuit. Règle : **N+1 minimum** — toujours un membre sain pendant l'opération.
Exemple : 2 DC → DC1 semaine 1, DC2 semaine 2 ; cluster Hyper-V → nœud par nœud
avec migration des VM.

**Critères de passage d'un anneau au suivant :**

- 0 BSOD / 0 non-démarrage sur l'anneau précédent ;
- applications métier validées (ouverture, impression, batch) ;
- taux d'échec d'installation < 2 %.

**Blocage** : si un correctif casse une application, on le **décline/bloque pour
l'anneau production** et on documente (KB, symptômes, contournement) en attendant
le correctif suivant. On ne laisse pas un parc diverger sans trace écrite.

---

## 12. Stratégie de ciblage par criticité

Ne traitez pas un serveur de fichiers et un poste de la même façon :

| Classe | Exemples | Patch | Redémarrage |
|---|---|---|---|
| P1 critique | DC, DNS/DHCP, hyperviseurs, ERP | Fenêtre dédiée, un par un | Planifié, annoncé |
| P2 important | Serveurs applicatifs, fichiers, impression | Fenêtre mensuelle | Planifié |
| P3 standard | Postes utilisateurs | Heures d'activité + report (§19) | Différé 7 jours max |
| Labo | Machines de test | Immédiat | Libre |

**Groupes de ciblage** : en WSUS via groupes d'ordinateurs ; en WUfB/Intune via
groupes Entra ID ou anneaux de mise à jour ; en GPO via OU ou filtrage de sécurité.
Nommez les groupes de façon explicite : `WU-Anneau-Test`, `WU-Anneau-Pilote`,
`WU-Prod-Serveurs-P1`, `WU-Prod-Postes`.

---

## 13. WSUS : architecture et dimensionnement

> ⚠️ WSUS est **déprécié** par Microsoft depuis septembre 2024 (maintenu, mais sans
> évolution). Ne lancez plus un nouveau WSUS en 2026 sauf contrainte (réseau isolé,
> parc 100 % local sans Intune). La cible moderne est Windows Update for Business (§16).

Architecture classique : un serveur **WSUS amont** synchronisé depuis Microsoft
Update, des **réplicas aval** par site, les clients configurés par GPO.

| Paramètre | Recommandation |
|---|---|
| OS | Windows Server 2022 (rôle WSUS) |
| Stockage contenu | 150-250 Go minimum (selon produits/langues sélectionnés) |
| Base | WID (Windows Internal Database) jusqu'à ~10 000 clients ; SQL au-delà |
| Ports | 8530 (HTTP) / 8531 (HTTPS) par défaut depuis Server 2012 |
| Produits | Uniquement ceux du parc (chaque produit/langue = Go téléchargés) |
| Classifications | Correctifs critiques, sécurité, cumulatifs, SSU |

**Pièges classiques :**

- Sélectionner « tous les produits » → des centaines de Go et une console inutilisable.
- Oublier les **langues** : limitez aux langues du parc.
- Laisser les mises à jour **remplacées** (superseded) approuvées → base obèse.

---

## 14. WSUS : installation et configuration

Installation du rôle (PowerShell) :

```powershell
# Sur le serveur WSUS (Server 2022)
Install-WindowsFeature -Name UpdateServices -IncludeManagementTools
# Configuration initiale : contenu sur D:, base WID
& 'C:\Program Files\Update Services\Tools\wsusutil.exe' postinstall CONTENT_DIR=D:\WSUS
```

Configuration de base (à faire dans la console, une fois) :

1. **Assistant de configuration** : choisir l'amont (Microsoft Update), le proxy si
   besoin, les **langues** (français + anglais), les **produits** (Windows 11,
   Windows Server 2019/2022/2025, Office si géré…), les **classifications**.
2. **Planification de synchronisation** : 2-4 fois/jour (ex. 06:00, 12:00, 18:00).
3. **Groupes d'ordinateurs** : créer `Test`, `Pilote`, `Prod-Serveurs`, `Prod-Postes`.
4. **Approbations** : manuelles par groupe (recommandé) ou règle d'approbation
   automatique limitée (ex. correctifs critiques → groupe Test uniquement).

GPO client type (appliquée par anneau) :

```text
Configuration ordinateur > Stratégies > Modèles d'administration >
Composants Windows > Windows Update > Gérer les mises à jour proposées par WSUS
  - Spécifier l'emplacement du service : http://wsus.contoso.local:8530
  - Activer le ciblage côté client : WU-Prod-Postes
```

Vérification côté client :

```powershell
# Le client voit-il le WSUS ?
(Get-ItemProperty 'HKLM:\SOFTWARE\Policies\Microsoft\Windows\WindowsUpdate').WUServer
# Forcer une détection
USOClient StartScan   # Windows 10/11 (remplace wuauclt /detectnow)
# Rapport d'état
Get-WindowsUpdateLog  # génère WindowsUpdate.log lisible sur le Bureau
```

---

## 15. WSUS : maintenance du serveur (Server Cleanup Wizard)

Un WSUS non entretenu meurt étouffé (base > 10 Go, console qui timeout). Maintenance
**mensuelle** :

1. **Décliner les mises à jour remplacées** : dans la console, vue « Toutes les mises
   à jour », filtrer Statut d'approbation = Non approuvée + État = Remplacée →
   clic droit **Refuser**. Ou en PowerShell :

