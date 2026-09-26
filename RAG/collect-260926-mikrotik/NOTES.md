# NOTES — collect-260926-mikrotik

- **Source** : `docs/RAG/lot-mikrotik/` (lot livré par le desktop, 691 `.md`).
- **Mode** : `files` (1 fiche = 1 chunk), avec garde-fou char (chunker v2,
  découpe au-delà de 8000 car. sur frontières de paragraphes).
- **Chunks** : 959 (691 fiches, 268 découpées).
- **Contenu** : RouterOS, forums (routing/BGP/OSPF, firewall/NAT, IPsec,
  WireGuard, VLAN/bridge, Wi-Fi/CAPsMAN, DHCP/DNS, QoS), CHR sur Proxmox,
  outils, troubleshooting.
- **Ancres** : aucune (`anchor_label` vide).
