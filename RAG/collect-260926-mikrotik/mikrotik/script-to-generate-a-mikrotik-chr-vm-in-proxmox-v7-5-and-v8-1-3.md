---
id: collect-260926-mikrotik/mikrotik/script-to-generate-a-mikrotik-chr-vm-in-proxmox-v7-5-and-v8-1-3
title: "script-to-generate-a-mikrotik-chr-vm-in-proxmox-v7-5-and-v8-1"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/lot-mikrotik/chr-proxmox/script-to-generate-a-mikrotik-chr-vm-in-proxmox-v7-5-and-v8-1.md
source_anchor: ""
source_lines: [218, 290]
sha256: 892b8f5c24b487f223b877670be090c69fbadac1759b6f4b1d9a791a7ac5ca73
---

# script-to-generate-a-mikrotik-chr-vm-in-proxmox-v7-5-and-v8-1

|  | TAG=$(echo $line \| awk '{print $1}') | 
|  | TYPE=$(echo $line \| awk '{printf "%-10s", $2}') | 
|  | FREE=$(echo $line \| numfmt --field 4-6 --from-unit=K --to=iec --format %.2f \| awk '{printf( "%9sB", $6)}') | 
|  | ITEM=" Type: $TYPE Free: $FREE " | 
|  | OFFSET=2 | 
|  | if [[ $((${#ITEM} + $OFFSET)) -gt ${MSG_MAX_LENGTH:-} ]]; then | 
|  | MSG_MAX_LENGTH=$((${#ITEM} + $OFFSET)) | 
|  | fi | 
|  | STORAGE_MENU+=("$TAG" "$ITEM" "OFF") | 
|  | done < <(pvesm status -content images \| awk 'NR>1') | 
|  | VALID=$(pvesm status -content images \| awk 'NR>1') | 
|  | if [ -z "$VALID" ]; then | 
|  | echo -e "\n${RD}⚠ Unable to detect a valid storage location.${CL}" | 
|  | echo -e "Exiting..." | 
|  | exit | 
|  | elif [ $((${#STORAGE_MENU[@]} / 3)) -eq 1 ]; then | 
|  | STORAGE=${STORAGE_MENU[0]} | 
|  | else | 
|  | while [ -z "${STORAGE:+x}" ]; do | 
|  | STORAGE=$(whiptail --backtitle "Proxmox VE Helper Scripts" --title "Storage Pools" --radiolist \ | 
|  | "Which storage pool you would like to use for the Mikrotik RouterOS CHR VM?\n\n" \ | 
|  | 16 $(($MSG_MAX_LENGTH + 23)) 6 \ | 
|  | "${STORAGE_MENU[@]}" 3>&1 1>&2 2>&3) \|\| exit | 
|  | done | 
|  | fi | 
|  | msg_ok "Using ${CL}${BL}$STORAGE${CL} ${GN}for Storage Location." | 
|  | msg_ok "Virtual Machine ID is ${CL}${BL}$VMID${CL}." | 
|  | msg_info "Getting URL for Mikrotik RouterOS CHR Disk Image" | 
|  | URL=https://download.mikrotik.com/routeros/7.12.1/chr-7.12.1.img.zip | 
|  | sleep 2 | 
|  | msg_ok "${CL}${BL}${URL}${CL}" | 
|  | wget -q --show-progress $URL | 
|  | echo -en "\e[1A\e[0K" | 
|  | FILE=$(basename $URL) | 
|  | msg_ok "Downloaded ${CL}${BL}$FILE${CL}" | 
|  | msg_info "Extracting Mikrotik RouterOS CHR Disk Image" | 
|  | gunzip -f -S .zip $FILE | 
|  | STORAGE_TYPE=$(pvesm status -storage $STORAGE \| awk 'NR>1 {print $2}') | 
|  | case $STORAGE_TYPE in | 
|  | nfs \| dir) | 
|  | DISK_EXT=".qcow2" | 
|  | DISK_REF="$VMID/" | 
|  | DISK_IMPORT="-format qcow2" | 
|  | ;; | 
|  | btrfs \| zfspool) | 
|  | DISK_EXT="" | 
|  | DISK_REF="$VMID/" | 
|  | DISK_FORMAT="subvol" | 
|  | DISK_IMPORT="-format raw" | 
|  | ;; | 
|  | esac | 
|  | DISK_VAR="vm-${VMID}-disk-0${DISK_EXT:-}" | 
|  | DISK_REF="${STORAGE}:${DISK_VAR:-}" | 
|  | msg_ok "Extracted Mikrotik RouterOS CHR Disk Image" | 
|  | msg_info "Creating Mikrotik RouterOS CHR VM" | 
|  | qm create $VMID -tablet 0 -localtime 1 -cores $CORE_COUNT -memory $RAM_SIZE -name $HN \ | 
|  | -tags proxmox-helper-scripts -net0 virtio,bridge=$BRG,macaddr=$MAC$VLAN$MTU \ | 
|  | -onboot 1 -ostype l26 -scsihw virtio-scsi-pci | 
|  | qm importdisk $VMID ${FILE%.*} $STORAGE ${DISK_IMPORT:-} 1>&/dev/null | 
|  | qm set $VMID \ | 
|  | -scsi0 "$DISK_REF" \ | 
|  | -boot order=scsi0 \ | 
|  | -description "<div align='center'><a href='https://mymangamedia.com/'><img src='https://avatars.githubusercontent.com/u/20906949'/></a> | 
|  | # Mikrotik RouterOS CHR | 
|  | <a href='https://mymangamedia.com/'><img src='https://avatars.githubusercontent.com/u/20906949' /></a> | 
|  | </div>" >/dev/null | 
|  | msg_ok "Mikrotik RouterOS CHR VM ${CL}${BL}(${HN})" | 
|  | if [ "$START_VM" == "yes" ]; then | 
|  | msg_info "Starting Mikrotik RouterOS CHR VM" | 
|  | qm start $VMID | 
|  | msg_ok "Started Mikrotik RouterOS CHR VM" | 
|  | fi | 
|  | msg_ok "Completed Successfully!\n" |
