---
id: collect-261001-rattrapage/rattrapage/win11-guide-25
title: "Windows 11 en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: ["2023-10-10", "2024-10-08", "2025-10-14", "2025-11-11", "2026-09-27", "2026-10-13", "2026-11-10", "2027-10-12"]
keywords: ["agent", "copilot"]
source: docs/RAG/collect-261001-rattrapage/win11_guide.md
source_anchor: ""
source_lines: [4131, 4268]
sha256: 5c3210dc5526e208489269835b06656376e6cf4fe972af93a17b07b3ce3a5964
---

# Windows 11 en entreprise — Guide technique ultra-complet

### Q10 — Windows 10 est hors support depuis le 14/10/2025. Un poste incompatible avec Windows 11 doit rester en service 6 mois pour une application métier. Que proposez-vous ?
- [ ] A. Le laisser tel quel, « ça a toujours marché »
- [ ] B. ESU (Extended Security Updates) avec une date de sortie écrite à 6 mois + isolement réseau renforcé
- [ ] C. Le contournement de registre TPM
- [ ] D. Désactiver Windows Update pour « ne pas casser l'application »

**Réponse : B.** L'ESU est le seul cadre supporté pour gagner du temps, **avec** une date de sortie écrite et des mesures compensatoires (VLAN isolé, pas d'Internet direct, surveillance renforcée). A/C/D augmentent le risque sans le traiter.

---

## 79. Pour aller plus loin : documentation, outils, communautés

### Documentation officielle
- **Microsoft Learn — Windows 11** : documentation de référence (déploiement, sécurité, MDM)
- **Windows release health** : problèmes connus par build, contournements officiels
- **Security Compliance Toolkit** : baselines + Policy Analyzer
- **Cycle de vie Microsoft** : dates de fin de support par version (à vérifier avant chaque projet)

### Outils
- **Sysinternals Suite** (Process Explorer, Autoruns, TCPView, PsTools) — diagnostic
- **SetupDiag** — échecs de mise à niveau
- **MDT + ADK + WinPE** — déploiement Lite Touch
- **Intune** / **MECM** — gestion moderne / co-management
- **Microsoft Graph PowerShell SDK** — automatisation Intune/Entra
- **Wireshark** — analyse réseau
- **CrystalDiskInfo** — santé SMART
- **Agent GLPI** — inventaire de parc

### Communautés (francophones et internationales)
- **Microsoft Q&A** (forums officiels, étiquette `windows-11`)
- **Reddit r/sysadmin** — retours d'expérience terrain
- **Serveurs Discord / forums** d'admins systèmes francophones
- **Blogs MVP** : déploiement, Autopilot, sécurité Windows

### Certifications utiles
- **MD-102** (Endpoint Administrator) : Intune, Autopilot, conformité — la certification « poste de travail moderne »
- **AZ-800/AZ-801** (Windows Server Hybrid Administrator) : AD, GPO, hybride
- **SC-300** (Identity Administrator) : Entra ID, accès conditionnel

### Prochaines étapes pour votre parc (suggestions)
1. Finaliser l'inventaire de compatibilité (section 3.2) couplé à GLPI
2. Choisir la cible : MDT vs Autopilot (ou les deux : MDT pour le fixe, Autopilot pour les nomades)
3. Déployer BitLocker + LAPS à 100 % avec clés centralisées
4. Mettre en place les anneaux WUfB et la baseline de sécurité
5. Planifier la sortie des derniers Windows 10 / ESU avec dates écrites

---

## 80. Annexe A : tableau des builds et versions

| Version | Build | UBR (exemple) | Canal | Fin de support Pro | Fin de support Entreprise |
|---|---|---|---|---|---|
| 21H2 | 22000 | .2538 | GA | 10/10/2023 | 08/10/2024 |
| 22H2 | 22621 | .5258 | GA | 08/10/2024 | 14/10/2025 |
| 23H2 | 22631 | .5624 | GA | 11/11/2025 | 10/11/2026 |
| 24H2 | 26100 | .7171 | GA | 13/10/2026 | 12/10/2027 |

> 📌 Vérifiez les dates et UBR sur le site officiel du cycle de vie Microsoft : elles évoluent à chaque cumulative. Au 27/09/2026, la cible de déploiement est **24H2**.

```powershell
# Identifier précisément un poste
Get-ItemProperty "HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion" |
    Select-Object ProductName, DisplayVersion, CurrentBuild, UBR
# DisplayVersion = 24H2, CurrentBuild = 26100, UBR = révision cumulative
```

---

## 81. Annexe B : chemins de clés de registre les plus utiles

| Usage | Chemin |
|---|---|
| Version Windows | `HKLM\SOFTWARE\Microsoft\Windows NT\CurrentVersion` |
| Stratégies Windows Update | `HKLM\SOFTWARE\Policies\Microsoft\Windows\WindowsUpdate` (+ `\AU`) |
| Delivery Optimization | `HKLM\SOFTWARE\Policies\Microsoft\Windows\DeliveryOptimization` |
| BitLocker (stratégies) | `HKLM\SOFTWARE\Policies\Microsoft\FVE` |
| LAPS Windows | `HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\LAPS` (état) |
| Désactiver Copilot | `HKLM\SOFTWARE\Policies\Microsoft\Windows\WindowsCopilot` |
| Widgets / News | `HKLM\SOFTWARE\Policies\Microsoft\Dsh` |
| Store | `HKLM\SOFTWARE\Policies\Microsoft\WindowsStore` |
| Télémétrie | `HKLM\SOFTWARE\Policies\Microsoft\Windows\DataCollection` |
| Profils utilisateurs | `HKLM\SOFTWARE\Microsoft\Windows NT\CurrentVersion\ProfileList` |
| Exécution au démarrage (machine) | `HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Run` |
| Pare-feu (profils) | Configuré via `netsh` / `Set-NetFirewallProfile` (pas de clé unique) |
| WinRE | `reagentc /info` (pas de registre utile direct) |
| Certificats machine | `Cert:\LocalMachine\My` (fournisseur de certificats) |

---

## 82. Annexe C : modèle de plan de migration Windows 10 → 11

```text
PLAN DE MIGRATION WINDOWS 10 → WINDOWS 11
Entreprise : ______________________    Responsable : ______________________
Période : du ____/____/________ au ____/____/________

1. PÉRIMÈTRE
   - Postes totaux : ____    - Compatibles : ____    - À remplacer : ____
   - Postes sous ESU (avec date de sortie) : ____

2. MÉTHODE DE DÉPLOIEMENT
   [ ] MDT (Lite Touch)   [ ] Autopilot + Intune   [ ] Mixte   [ ] Manuelle
   Image / profil de référence : ________________________________________

3. VAGUES
   | Vague | Périmètre (service/site) | Nb postes | Date début | Date fin | Responsable |
   |-------|----------------------------|-----------|------------|----------|-------------|
   | Pilote| IT + volontaires          |           |            |          |             |
   | 1     |                            |           |            |          |             |
   | 2     |                            |           |            |          |             |
   | 3     |                            |           |            |          |             |

4. PRÉREQUIS PAR POSTE (checklist section 74)
   [ ] Compatibilité vérifiée   [ ] Sauvegarde OneDrive OK   [ ] Applications métier validées

5. COMMUNICATION
   - Annonce J-14 : __________   - Rappel J-3 : __________   - Support renforcé J+3 : __________

6. GESTION DES RISQUES
   | Risque | Probabilité | Impact | Mitigation |
   |--------|-------------|--------|------------|
   | KB fautive | Moyenne | Élevé | Anneaux WUfB + pause |
   | Application incompatible | Moyenne | Élevé | Validation vague pilote |
   | Clé BitLocker perdue | Faible | Critique | GPO sauvegarde AD obligatoire |

7. CRITÈRES DE SORTIE
   [ ] 100 % des postes compatibles migrés
   [ ] Chaque poste non migré a une fiche d'exception datée (remplacement / ESU)
   [ ] Clés BitLocker centralisées vérifiées à 100 %
   [ ] Documentation d'exploitation mise à jour

8. CLÔTURE
   Date : __________    Validation RSSI / Direction : __________
```

---

*Fin du guide — Windows 11 en entreprise (23H2/24H2). Bon déploiement !*
