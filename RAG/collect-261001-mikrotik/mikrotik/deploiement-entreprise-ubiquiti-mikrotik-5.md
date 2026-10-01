---
id: collect-261001-mikrotik/mikrotik/deploiement-entreprise-ubiquiti-mikrotik-5
title: "Déploiement entreprise — Ubiquiti & MikroTik (référence complète)"
domain: mikrotik
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-mikrotik/deploiement_entreprise_ubiquiti_mikrotik.md
source_anchor: ""
source_lines: [681, 695]
sha256: 8fa22e927bdcfbb6bffe737fec12ffd670e258fb50dd3bab3ed1befd0c39bcfb
---

# Déploiement entreprise — Ubiquiti & MikroTik (référence complète)

- **MikroTik** : rester sur la branche **stable**, lire le changelog
  (les upgrades RouterOS sont rapides mais parfois cassants sur les
  configs exotiques) ; Netinstall en plan B ultime.
- **EdgeRouter** : upgrades moins fréquents, sauvegarder `config.boot`
  avant (`cp /config/config.boot /config/config.boot.bak`).
- **UniFi** : contrôleur backupé avant chaque upgrade ; ne pas upgrader
  contrôleur + firmwares AP le même jour en prod.
- **Trimestriel** : comptes inactifs, règles firewall obsolètes,
  certificats, capacité (ports, CPU), test de restore backup.

---

*Fin de la référence. Compagnons : [Déploiement Cisco & Huawei](deploiement_entreprise_cisco_huawei.md)
(architecture 3-tier détaillée), [Guide CLI OPNsense](opnsense_cli_guide.md),
[Guide web OPNsense](opnsense_gui_guide.md), [Scénarios OPNsense](opnsense_deploiements.md).*
