---
id: collect-261001-general-networking/general-networking/pasqal-170-m-ipo-nasdaq-quantique-francais-2026-3
title: "Exemple minimaliste avec Pulser, le SDK officiel Pasqal"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Google", "Microsoft", "Nvidia"]
dates: []
keywords: ["aws", "gpu", "ipo", "mai", "nvidia", "open source"]
source: docs/RAG/collect-261001-general-networking/pasqal-170-m-ipo-nasdaq-quantique-francais-2026.md
source_anchor: ""
source_lines: [73, 127]
sha256: 17742aece2b5f3c8deb7690f7f51bc2f7073efcdab42ae0afb6ef11543e08717
---

# Exemple minimaliste avec Pulser, le SDK officiel Pasqal

Le paysage compétitif du quantique mondial s’est densifié en 2025-2026. **IBM** a annoncé son processeur Condor à 1 121 qubits supraconducteurs et vise 100 000 qubits à l’horizon 2033. **Google** a lancé en décembre 2024 son processeur Willow à 105 qubits, avec une démonstration majeure de correction d’erreurs. **Quantinuum** a fusionné ses activités H1 et H2 et lève régulièrement à des valorisations dépassant les 10 milliards de dollars. **PsiQuantum**, le pari photonique américano-australien, a sécurisé près d’**1 milliard de dollars** sur 2024-2025 et construit deux usines au Queensland et à Chicago.

Sur la verticale spécifique des atomes neutres, Pasqal n’est plus seul. Le concurrent américain **QuEra**, issu de Harvard/MIT, a livré sa machine Aquila (256 qubits) chez AWS Braket et lève également à des valorisations licornes. **Atom Computing**, autre américain, a atteint 1 180 qubits en 2024 – chiffre dépassant nominalement la roadmap Pasqal, mais sans capacité commerciale équivalente. La différenciation de Pasqal repose donc moins sur la course aux qubits que sur l’intégration HPC, la qualité opérationnelle et la base de clients industriels payants.

## Tableau comparatif : Pasqal vs concurrents quantiques mondiaux 2026

| Acteur | Pays | Technologie | Qubits 2026 | Valorisation | Cotation | 
|---|---|---|---|---|---|
| Pasqal | France | Atomes neutres | 140 → 250 | ~1,2-1,5 Md$ (privé) | Nasdaq prévu 2026 | 
| IBM Quantum | États-Unis | Supraconducteurs | 1 121 (Condor) | Branche IBM | NYSE (IBM) | 
| Google Quantum AI | États-Unis | Supraconducteurs | 105 (Willow) | Branche Alphabet | NASDAQ (GOOGL) | 
| Quantinuum | USA/UK | Ions piégés | ~56 (H2) | ~10 Md$ | IPO en préparation | 
| IonQ | États-Unis | Ions piégés | ~64 | ~10 Md$ (capi.) | NYSE: IONQ | 
| PsiQuantum | USA/Australie | Photonique | roadmap 1M | ~6 Md$ (privé) | Non coté | 
| QuEra | États-Unis | Atomes neutres | 256 (Aquila) | n.c. (privé) | Non coté | 
| Quandela | France | Photonique | 12-50 | n.c. (privé) | Non coté | 
| C12 Quantum | France | Nanotubes de carbone | roadmap 2030 | ~150 M€ (privé) | Non coté | 
| Alice & Bob | France | Qubits de chat | ~10 | ~100 M€ (privé) | Non coté | 

## L’écosystème quantique français : Pasqal, Quandela, Alice & Bob, C12

La France compte aujourd’hui quatre pépites quantiques structurantes : **Pasqal** (atomes neutres), **Quandela** (photonique), **Alice & Bob** (qubits de chat tolérants aux fautes) et **C12 Quantum Electronics** (nanotubes de carbone). Cette diversité de modalités est unique au monde – aucun autre pays européen ne dispose d’un tel portefeuille de paris technologiques différenciés. Pasqal est cependant la plus avancée commercialement, avec ses machines déjà déployées chez des clients industriels et son passage prévu en Bourse.

Cette dynamique s’appuie sur le **Plan Quantique national** français, doté de 1,8 milliard d’euros sur 2021-2026 selon les annonces gouvernementales relayées par France 2030 – un dispositif qui continue d’irriguer des projets ciblés, à l’image des **4 millions d’euros** accordés en février 2026 au projet InterQo, mené par Pasqal et Welinq dans le cadre du programme France 2030 i-Demo Régionalisé. Le plan a notamment financé l’installation de machines quantiques au **GENCI** (Grand Équipement National de Calcul Intensif), au **TGCC** de Bruyères-le-Châtel et au Jülich Supercomputing Centre en Allemagne via le projet HPCQS. Pasqal a remporté en 2022 le contrat HPCQS pour fournir deux ordinateurs quantiques destinés à être couplés à des supercalculateurs européens.

« Le quantique français a passé l’étape de la recherche : il est désormais dans la phase d’industrialisation. Pasqal en est le meilleur exemple », analyse Maud Vinet, fondatrice de Quobly et figure de l’écosystème quantique français, dans une intervention à la French Quantum Week 2025.

## Intégration HPC : GENCI, HPCQS, Microsoft Azure Quantum

La stratégie commerciale de Pasqal repose largement sur le couplage avec le **calcul haute performance**. Plutôt que de vendre des QPU isolés, l’entreprise propose des architectures hybrides où le quantique accélère des sous-routines spécifiques (optimisation combinatoire, simulation moléculaire, échantillonnage) au sein de workflows HPC classiques. Cette approche convient parfaitement aux besoins industriels de 2026, où l’avantage quantique pur reste à démontrer.

Côté cloud, Orion Alpha est accessible via **Microsoft Azure Quantum**, la Google Cloud Marketplace – canal ouvert en mai 2025 –, le portail propriétaire Pasqal Cloud et, depuis décembre 2025, la plateforme QaaS de **Scaleway**, une intégration dévoilée lors de l’événement ai-PULSE 2025 à Paris. Cette disponibilité multi-cloud est un argument commercial puissant : un développeur peut prototyper un algorithme en quelques minutes sans investir dans une machine on-premise à plusieurs millions d’euros. Le SDK officiel de Pasqal, `pulser`, est open source et permet de programmer des séquences d’impulsions optiques avec une abstraction haut niveau. Sur le plan logiciel, Pasqal a annoncé le 16 mars 2026 l’intégration de **NVIDIA CUDA-Q** à son runtime **QRMI**, renforçant l’interopérabilité entre ses QPU à atomes neutres et les infrastructures GPU dans les workflows hybrides HPC-quantique.

```
# Exemple minimaliste avec Pulser, le SDK officiel Pasqal
import numpy as np
from pulser import Pulse, Sequence, Register
from pulser.devices import AnalogDevice
# Creation d'un registre de 5 atomes en ligne
register = Register.line(5, spacing=5)
# Creation de la sequence
sequence = Sequence(register, AnalogDevice)
sequence.declare_channel("rydberg_global", "rydberg_global")
# Impulsion de Rabi (5 microsecondes)
pulse = Pulse.ConstantPulse(5000, 2*np.pi, 0, 0)
sequence.add(pulse, "rydberg_global")
# Simulation locale ou execution distante sur Orion Alpha via Pasqal Cloud
print(sequence)
```
Le déploiement en France via **GENCI** est emblématique de la stratégie hybride. La machine « Ruby » (Orion Beta), installée en 2024, est utilisée par les laboratoires académiques français pour des cas d’usage en simulation de matériaux, en théorie des champs et en optimisation logistique. Cette approche a été renforcée en juillet 2026 par un nouveau déploiement adossé à un supercalculateur, co-financé à 50 % par l’**EuroHPC Joint Undertaking** pour un budget total de **12 millions d’euros**. Le même mois, le système Orion de Pasqal, doté de **140 qubits**, a mis en service le tout premier ordinateur quantique à atomes neutres d’Italie, une étape rapportée par Quantum Computing Report en juillet 2026. C’est un pont entre la recherche fondamentale et les applications industrielles que peu de concurrents internationaux ont su construire.

## Plan Quantique européen : la souveraineté technologique en jeu

L’opération Pasqal s’inscrit dans un contexte politique européen en pleine recomposition. La Commission européenne a engagé des travaux préparatoires pour un **Quantum Act**, projet législatif visant à structurer l’écosystème quantique du continent et à réduire la dépendance aux acteurs américains et chinois. Preuve concrète de cette ambition, la ligne pilote européenne **Q-PLANET**, pilotée par Pasqal, a été lancée en juillet 2026 avec un budget de **50 millions d’euros** et réunit **37 partenaires** répartis dans **12** États membres. Selon le portail Digital Strategy de l’UE, le programme phare Quantum Flagship a déjà mobilisé près d’1 milliard d’euros sur dix ans (2018-2028), et de nouveaux financements seront déployés au titre du programme Horizon Europe.

