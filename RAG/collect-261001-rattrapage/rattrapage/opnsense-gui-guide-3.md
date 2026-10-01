---
id: collect-261001-rattrapage/rattrapage/opnsense-gui-guide-3
title: "OPNsense — Guide ultra-complet de l'interface web"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/opnsense_gui_guide.md
source_anchor: ""
source_lines: [411, 445]
sha256: db63862cdaf9b29046e7640d58c92df5563325f73d0b3c76c22c660bfd0d55aa
---

# OPNsense — Guide ultra-complet de l'interface web

1. **Dashboard** : gateways UP ? services verts ? mises à jour en
   attente ?
2. **Firewall > Log Files > Live View** : reproduisez le problème et
   regardez ce qui est **bloqué** (action Block, interface, règle).
   80 % des « pannes » se voient ici en 30 secondes.
3. **Interfaces > Diagnostics > Packet Capture** : le paquet arrive-t-il
   ? repart-il ? (sens aller/retour).
4. **Firewall > Diagnostics > States** : l'état existe-t-il ? Tuez-le
   après un changement de règle/NAT.
5. **VPN > … > Status** : tunnel UP ? handshake récent (WireGuard) ?
6. **System > Log Files > General** : erreurs de services.
7. **System > Configuration > History** : « qu'est-ce qui a changé
   juste avant la panne ? » → restaurez la révision précédente.

---

## 17. Raccourcis et productivité

- **Recherche globale** (loupe en haut) : le plus rapide pour trouver
  une page (« alias », « wireguard », « cron »…).
- **Favoris** : épinglez vos pages fréquentes.
- **API REST** : presque tout ce que fait la web UI est scriptable
  (clé API par utilisateur : System > Access > Users > API keys).
  Idéal pour : créer des alias en masse, automatiser les backups,
  superviser via scripts.
- **Thème sombre** disponible (System > Settings > General).
- **Widgets du dashboard** : ajoutez « Firewall Logs », « Gateways »,
  « Services », « Updates » pour un NOC miniature.

---

*Fin du guide. Avec la web UI vous couvrez 100 % des fonctions ;
avec le [guide CLI](opnsense_cli_guide.md) vous faites la même chose
au clavier. Les deux se complètent : la web UI pour découvrir, le CLI
pour automatiser et dépanner vite.*
