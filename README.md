# OVCLI: A HPE Synergy Command Tool

OVCLI is a CLI tool to manage Hewlett Packard Enterprise Synergy. It leverages Synergy OneView Restful API interface to provide IT admins with quick CLI access to Synergy.

![image](https://user-images.githubusercontent.com/14317124/31856927-38e2c82a-b684-11e7-9924-cc12d617914a.png)

OVCLI is written in Go. For developers, latest source code and binary can be directly downloaded by ```go get``` command.

## Windows .EXE Download
## [ovcli.exe](https://github.com/hjma29/ovcliexe)
<a href="https://github.com/hjma29/ovcliexe" download="ovcli.exe">ovcli.exe</a>



```
$ovcli show serverprofile
Name             Template       Hardware                              Hardware Type
DCA-ToR-Host17                  Top - Frame1 - CN7515010J, bay 7      SY 480 Gen9 1
DCA-Tor-Host28                  Middle - Frame 2 -CN75150484, bay 8   SY 480 Gen9 1
vsan node 1      vsan profile   Middle - Frame 2 -CN75150484, bay 1   SY 480 Gen9 3
vsan node 2      vsan profile   Middle - Frame 2 -CN75150484, bay 4   SY 480 Gen9 3
vsan node 3      vsan profile   Middle - Frame 2 -CN75150484, bay 6   SY 480 Gen9 3
```

```
$show serverprofile --name "vsan node 1"
------------------------------------------------------------------------------
Name:                 vsan node 1
Description:
ProfileTemplate:      vsan profile
TemplateCompliance:   Unknown
ServerHardware:       Middle - Frame 2 -CN75150484, bay 1
ServerPower:          Off
ServerHardwareType:   SY 480 Gen9 3

Connections
ID   Name             Network          VLAN         MAC                 Port         Interconnect                                   Boot
1    vSAN Private 1   vSAN_Network     2000         F2:BA:A8:F0:01:B8   Mezz 3:1-d   Top - Frame1 - CN7515010J, interconnect 3      NotBootable
2    vSAN Private 2   vSAN_Network     2000         F2:BA:A8:F0:01:B9   Mezz 3:2-d   Middle - Frame 2 -CN75150484, interconnect 6   NotBootable
3    Vmotion Data 1   VLAN301-304      NetworkSet   F2:BA:A8:F0:01:BA   Mezz 3:1-c   Top - Frame1 - CN7515010J, interconnect 3      NotBootable
4    Vmotion dAta 2   VLAN301-304      NetworkSet   F2:BA:A8:F0:01:BB   Mezz 3:2-c   Middle - Frame 2 -CN75150484, interconnect 6   NotBootable
5    Mgmt 1           TE-Testing-300   300          F2:BA:A8:F0:01:BC   Mezz 3:1-b   Top - Frame1 - CN7515010J, interconnect 3      NotBootable
6    Mgmt 2           TE-Testing-300   300          F2:BA:A8:F0:01:BD   Mezz 3:2-b   Middle - Frame 2 -CN75150484, interconnect 6   NotBootable
```

##


* 2-day onsite 
* 3-day virtual

