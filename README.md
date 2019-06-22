# snmpbeat
Custom beat for snmp monitoring in firewall. Este codigo esta basado en https://github.com/isalgueiro/otilio

### Configurando snmpbeat.yml
```
snmpbeat:
  # Defines how often an event is sent to the output
  period: 1s

  # SNMP hosts to query
  hosts: ["192.168.1.1", "192.168.1.2"]

  # SMNP version: 1, 2c or 3
  version: 2c

  # SNMP community
  community: "public"

  # oids to query
  # (the starting dot is intended)
  oids:
    - {oid: ".1.3.6.1.2.1.1.1.0", name: sysDescr}
    - {oid: ".1.3.6.1.2.1.1.3.0", name: sysUpTime}
```
This will get oids `1.3.6.1.2.1.1.1.0` and `1.3.6.1.2.1.1.3.0` from SNMP servers at 192.168.1.1 and 192.168.1.2 and store them in `otilio-YYYY.MM.DD` index in Elasticsearch in fields `sysDescr` and `sysUpTime`.

SNMP V3 configuration example

```
snmpbeat:
  # Defines how often an event is sent to the output
  period: 1s

  # SNMP host to query
  hosts: ["127.0.0.1"]
  port: 10161

  # SMNP version
  version: 3

  # SNMP user security model parameters
  # currently only SHA auth and DES encryption supported ¯\_(ツ)_/¯
  user: "theuser"
  authpass: "theauthpassword"
  privpass: "theprivacyencryptionpassword"

  # oids to query
  # (the starting dot is intended)
  oids:
    - {oid: ".1.3.6.1.2.1.25.1", name: hrSystem}
```

### Runing

#### Modo Debug

```
./snmpbeat -c snmpbeat.yml -e -d "*"
```

## Instalacion de GO
Link : https://linuxize.com/post/how-to-install-go-on-centos-7/

Descargar goland usado "wget" o "curl"

> wget https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz

Verificar el tarball

> sha256sum go1.10.3.linux-amd64.tar.gz

El checksum deberia arrojar lo siguiente:

> fa1b0e45d3b647c252f51f5e1204aba049cde4af177ef9f2181f43004f901035  go1.10.3.linux-amd64.tar.gz

Extraer el archivo en "/usr/local/"

> fa1b0e45d3b647c252f51f5e1204aba049cde4af177ef9f2181f43004f901035  go1.10.3.linux-amd64.tar.gz

Ajustar la variable de entorno

> export GOROOT=/usr/local/go
> export PATH=$PATH:/usr/local/go/bin

Guardando el archivo y cargar el new PATH  a las variables de entorno

> source ~/.bash_profile

## Compilación de snmpbeat

## Árbol de carpetas
- $GOPATH: ruta al espacio de trabajo
   - snmpbeat.yml : Archivo de configuración requerido por snmpbeat
   - [src] : Carpetas que contiene el codigo fuente a ser compilado
       - [github.com]  :Carpeta adicional para agregar el codigo de github
          - [hoat23]
             - [snmpbeat] :Carpeta con el codigo fuente
          - [elastic]
             - [beats] :Carpeta con el codigo adicional para compilar el beats
                - [libbeat]      :Libreria muy importante para la compilación del código
                - [auditbeat]
                - [heartbeat]
                - [metricbeat]
                - [packetbeat]
                - [winlogbeat]
                - [deploy]
                - [dev-tools]
                - [docs]
                - [filebeat]
                - [generator]
                - [journalbeat]
                - [licenses]
                - [script]
                - [testing]
                - [vendor]
                - .editorconfigj
                - .gitattributes
                - .gitnore
                - .go-version
                - .travis.yml
                - .PYLITRC
                - Jenkisfile
                - LICENCE.txt
                - magefile.go    : Archivo make para la compilación con go
                - MAKE.BAT
                - Makefile
                - NOTICE.txt
                - reviewdog.yml
                - setup.yml
          


   
 
