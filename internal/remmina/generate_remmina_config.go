package remmina

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var fileContentTempl = `[remmina]
ssh_tunnel_loopback=0
window_maximize=1
ssh_tunnel_passphrase=
name=%s
window_width=640
password=
ssh_forward_x11=0
ssh_proxycommand=
ssh_passphrase=
run_line=
precommand=
sshlogenabled=0
ssh_privatekey=/home/egor/.ssh/id_ed25519
ssh_tunnel_enabled=0
ssh_charset=
window_height=480
keyboard_grab=0
ssh_auth=1
ignore-tls-errors=1
postcommand=
server=%s
disablepasswordstoring=0
ssh_color_scheme=0
audiblebell=0
ssh_tunnel_username=
sshsavesession=0
ssh_hostkeytypes=
ssh_tunnel_password=
profile-lock=0
exec=
group=new-gonet
ssh_tunnel_server=
ssh_ciphers=
enable-autostart=0
ssh_kex_algorithms=
ssh_compression=0
ssh_tunnel_auth=0
ssh_tunnel_certfile=
notes_text=
sshlogfolder=
viewmode=1
sshlogname=
ssh_tunnel_privatekey=
protocol=SSH
ssh_stricthostkeycheck=0
username=egor_unikovskiy
`

type Config struct {
	Domain string
	Ip     string
	//Group  string
}

func GenerateRemminaFiles(servers []Config) {
	filenameTempl := "new-gonet_ssh_%s_%s.remmina"
	fmt.Printf("will be generate %d remmina config\n", len(servers))
	for _, s := range servers {
		clearDomain := replaceDot(s.Domain)
		clearIP := replaceDot(s.Ip)
		filename := fmt.Sprintf(filenameTempl, clearDomain, clearIP)
		fileContent := fmt.Sprintf(fileContentTempl, s.Domain, s.Ip)

		file, err := os.Create(filename)
		if err != nil {
			log.Fatalln("create failed:", err)
		}
		_, err = file.WriteString(fileContent)
		if err != nil {
			log.Fatalln("write failed:", err)
		}
	}
}

func replaceDot(line string) string {
	return strings.ReplaceAll(line, ".", "-")
}
