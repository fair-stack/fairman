#!/bin/sh

ip=dockerOn serverip
password=aaaaaa
dir=/root/docker/cert

# create folder
if [ ! -d "$dir" ];then
echo ""
echo "$dir , not dir , will create"
echo ""
mkdir -p $dir
else
echo ""
echo "$dir , dir exist , will delete and create"
echo ""
rm -rf $dir
mkdir -p $dir
fi

cd $dir
# Create Root CertificateRSACreate Root Certificate
openssl genrsa -aes256 -passout pass:$password  -out ca-key.pem 4096
# establishCAestablish
openssl req -new -x509 -days 365 -key ca-key.pem -passin pass:$password -sha256 -out ca.pem -subj "/C=NL/ST=./L=./O=./CN=$ip"
# Create a server private key
openssl genrsa -out server-key.pem 4096
# Create a server signature request certificate file
openssl req -subj "/CN=$ip" -sha256 -new -key server-key.pem -out server.csr

# allocationIPallocation
echo subjectAltName = IP:$ip,IP:0.0.0.0 >> extfile.cnf
# takeDockertake
echo extendedKeyUsage = serverAuth >> extfile.cnf
# Create a server certificate file with valid signature
openssl x509 -req -days 365 -sha256 -in server.csr -CA ca.pem -CAkey ca-key.pem -passin "pass:$password" -CAcreateserial -out server-cert.pem -extfile extfile.cnf
# Create client private key
openssl genrsa -out key.pem 4096
# Create a client signature request certificate file
openssl req -subj '/CN=client' -new -key key.pem -out client.csr

# To make the key suitable for client authentication，To make the key suitable for client authentication
echo extendedKeyUsage = clientAuth >> extfile.cnf
echo extendedKeyUsage = clientAuth > extfile-client.cnf

# Create a client certificate file with valid signature
openssl x509 -req -days 365 -sha256 -in client.csr -CA ca.pem -CAkey ca-key.pem -passin "pass:$password" -CAcreateserial -out cert.pem -extfile extfile-client.cnf
# Delete excess files Delete excess files
rm -f -v client.csr server.csr extfile.cnf extfile-client.cnf
# Modify permissions，Modify permissions，Modify permissions
chmod -v 0400 ca-key.pem key.pem server-key.pem
# Certificate is externally readable，Certificate is externally readable
chmod -v 0444 ca.pem server-cert.pem cert.pem

# modifyDockermodify，modifyDockermodifyCAmodify
# vim /lib/systemd/system/docker.service
# ExecStart=/usr/bin/dockerd Add below
# --tlsverify --tlscacert=/etc/docker/ca.pem --tlscert=/etc/docker/server-cert.pem --tlskey=/etc/docker/server-key.pem -H tcp://0.0.0.0:2376 -H unix:///var/run/docker.sock
# ReloaddaemonReloaddocker
# systemctl daemon-reload
# systemctl restart docker
# check2376check
# netstat -lnpt
# If there is no monitoring，If there is no monitoringdocker.serviceIf there is no monitoring
# systemctl status docker
# Save relevant client certificate files locally
# ca.pem、cert.pem、key.pem
