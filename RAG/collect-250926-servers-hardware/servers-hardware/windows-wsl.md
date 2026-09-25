---
id: collect-250926-servers-hardware/servers-hardware/windows-wsl
title: "Windows (WSL)"
domain: servers-hardware
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/clean4/windows-wsl.md
source_anchor: ""
source_lines: [1, 40]
sha256: 2fcefec852b01ab4cae0e732693e7a2a2bde7d896296cb8a9056e229f41b3615
---

# Windows (WSL)

ExÃ©cutez OpenCode sur Windows avec WSL pour une expÃ©rience optimale.

MÃªme si OpenCode peut fonctionner directement sur Windows, nous recommandons dâutiliser Windows Subsystem for Linux (WSL) pour la meilleure expÃ©rience. WSL fournit un environnement Linux qui sâintÃ¨gre parfaitement aux fonctionnalitÃ©s dâOpenCode.

1. 
**Installez WSL**Si ce nâest pas encore fait, installez WSL Ã lâaide du guide officiel Microsoft.
2. 
**Installez OpenCode dans WSL**Une fois WSL configurÃ©, ouvrez votre terminal WSL et installez OpenCode avec lâune des mÃ©thodes dâinstallation.
3. 
**Utilisez OpenCode depuis WSL**Allez dans votre dossier de projet (accÃ©dez aux fichiers Windows via `/mnt/c/` ,`/mnt/d/` , etc.) et lancez OpenCode.

Si vous prÃ©fÃ©rez utiliser lâapplication Desktop OpenCode tout en exÃ©cutant le serveur dans WSL:

1. 
**DÃ©marrez le serveur dans WSL** avec`--hostname 0.0.0.0` pour autoriser les connexions externes:
2. 
**Connectez lâapplication Desktop** Ã`http://localhost:4096`

Pour la meilleure expÃ©rience web sous Windows:

1. 
**ExÃ©cutez `opencode web` dans le terminal WSL** plutÃ´t que dans PowerShell:
2. 
**AccÃ©dez-y depuis votre navigateur Windows** Ã`http://localhost:<port>` (OpenCode affiche lâURL)

Lancer `opencode web` depuis WSL garantit un accÃ¨s correct au systÃ¨me de fichiers et une bonne intÃ©gration terminal, tout en restant accessible depuis votre navigateur Windows.

WSL peut accÃ©der Ã  tous vos fichiers Windows via le rÃ©pertoire `/mnt/`:

- Lecteur `C:` â`/mnt/c/`
- Lecteur `D:` â`/mnt/d/`
- Et ainsi de suiteâ¦

Exemple:

- Gardez OpenCode dans WSL pour les projets stockÃ©s sur des lecteurs Windows: lâaccÃ¨s aux fichiers est fluide
- Utilisez lâextension WSL de VS Code avec OpenCode pour un flux de travail intÃ©grÃ©
- Votre configuration OpenCode et vos sessions sont stockÃ©es dans lâenvironnement WSL Ã  `~/.local/share/opencode/`
