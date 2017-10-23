# OVCLI: A HPE Synergy Command Line Tool

OVCLI is a CLI tool to manage Hewlett Packard Enterprise Synergy resources. It provides IT admins with a cross-platform HPE Synergy CLI utility that can directly run on Linux, Windows, and MAC operating systems without worrying about installing/troubleshooting any programming language library dependencies.

OVCLI tool communicates with Synergy OneView Restful API interface. It complements HPE OneView GUI interface to provide quicker Synergy access in many use cases.

![image](https://user-images.githubusercontent.com/14317124/31856927-38e2c82a-b684-11e7-9924-cc12d617914a.png)


## [**Windows .EXE Download**](https://github.com/hjma29/ovcliexe/blob/master/ovcli.exe?raw=true)
For IT admins just want a simple small CLI tool, they can directly download the above .exe file. The image was compiled on windows2016 64-bit OS and tested also on windows2012R2 OS.


## Developer Source Code Access

OVCLI tool is written in Go bebind the scene. For developers, latest source code can be directly downloaded from the [github repository](https://github.com/hjma29/ovcli) by using <code>go get</code> command and further compilied into binary executable file for the platform.

## Release Notes
* Synergy 3.10 has been tested. Synergy 3.0 should work for most cases. 
* For questions, please send email to <hongjun.ma@hpe.com>

## Getting Started
* **Verify executable file** by running command <code>ovcli</code> either in the current directory or through system PATH environment setup. The succuss run should general the output like the sameple below.
<pre>
$ ovcli
Release Version: 0.1

ovcli is a Synergy OneView CLI tool, please use "--help" option
to explore what are next available options

Usage:
  ovcli [flags]
  ovcli [command]
  ...
  ...
</pre>

* **Create a Synergy configuration text file** including Synergy Mgmt IP address/hostname, login username and password in the current directory. It's a very simple text file with the following three lines(please update with your corresponding login credentials)

```
ipaddress: 192.168.1.1
username: Administrator
password: password
```

* **Populate OVCLI login configuration file** and test the connection with HPE Synergy using the above file by running "ovcli connect -f <filname>". A successful Synergy connection should come back with the Synergy API version. "500" below shows Synergy is running OneView version 3.10. (Behind the scene, the ovcli will copy the file to a new file called "appliance-credential.yml" and verify the Synergy connection using the yaml file contents.)


```
C:\Users\Administrator\Downloads>ovcli connect -f login.txt
Appliance Address      Username        Appliance Current Version
-----------------      --------        -------------------------
https://10.16.44.101   Administrator   500
```

* With login credential verified, you can **try different "ovcli show"** commands like the following examples. "--help" or "-h" flag should give you help contexts for the commands.

**Show Server Profiles**
<pre>
$ovcli show serverprofile
Name             Template       Hardware                              Hardware Type
DCA-ToR-Host17                  Top - Frame1 - CN7515010J, bay 7      SY 480 Gen9 1
DCA-Tor-Host28                  Middle - Frame 2 -CN75150484, bay 8   SY 480 Gen9 1
vsan node 1      vsan profile   Middle - Frame 2 -CN75150484, bay 1   SY 480 Gen9 3
vsan node 2      vsan profile   Middle - Frame 2 -CN75150484, bay 4   SY 480 Gen9 3
vsan node 3      vsan profile   Middle - Frame 2 -CN75150484, bay 6   SY 480 Gen9 3
</pre>

**Show Server Profile details for one specific profile**
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
**Show logical uplink details for all LIs**
```
→ ovcli show li --name all
------------------------------------------------------------------------------
SSA-DCA-3Frame-DCA-Solcenter-LIG-Copy
  UplinkSet: DCA-ToR
       Networks:
            Network Name             VlanID   Type
            ------------             ------   ----
            DCA-ACI-VMM-Pool__1041   1041     Tagged
            DCA-ACI-VMM-Pool__1042   1042     Tagged
            DCA-ACI-VMM-Pool__1043   1043     Tagged
            DCA-ACI-VMM-Pool__1044   1044     Tagged
            DCA-ACI-VMM-Pool__1045   1045     Tagged
            DCA-ACI-VMM-Pool__1046   1046     Tagged
            DCA-ACI-VMM-Pool__1047   1047     Tagged
            DCA-ACI-VMM-Pool__1048   1048     Tagged
            DCA-ACI-VMM-Pool__1049   1049     Tagged
            DCA-ACI-VMM-Pool__1050   1050     Tagged
            DCA-ACI-VMM-Pool__1051   1051     Tagged
            DCA-ACI-VMM-Pool__1052   1052     Tagged
            DCA-ACI-VMM-Pool__1053   1053     Tagged
            DCA-ACI-VMM-Pool__1054   1054     Tagged
            DCA-ACI-VMM-Pool__1055   1055     Tagged
            DCA-ACI-VMM-Pool__1056   1056     Tagged
            DCA-ACI-VMM-Pool__1057   1057     Tagged
            DCA-ACI-VMM-Pool__1058   1058     Tagged
            DCA-ACI-VMM-Pool__1059   1059     Tagged
            DCA-ACI-VMM-Pool__1060   1060     Tagged
            DCA-ACI-VMM-Pool__1061   1061     Tagged
            DCA-ACI-VMM-Pool__1062   1062     Tagged
            DCA-ACI-VMM-Pool__1063   1063     Tagged
            DCA-ACI-VMM-Pool__1064   1064     Tagged
            DCA-ACI-VMM-Pool__1065   1065     Tagged
            DCA-ACI-VMM-Pool__1066   1066     Tagged
            DCA-ACI-VMM-Pool__1067   1067     Tagged
            DCA-ACI-VMM-Pool__1068   1068     Tagged
            DCA-ACI-VMM-Pool__1069   1069     Tagged
            DCA-ACI-VMM-Pool__1070   1070     Tagged
            DCA-ACI-VMM-Pool__1071   1071     Tagged
            DCA-ACI-VMM-Pool__1072   1072     Tagged
            DCA-ACI-VMM-Pool__1073   1073     Tagged
            DCA-ACI-VMM-Pool__1074   1074     Tagged
            DCA-ACI-VMM-Pool__1075   1075     Tagged
            DCA-ACI-VMM-Pool__1076   1076     Tagged
            DCA-ACI-VMM-Pool__1077   1077     Tagged
            DCA-ACI-VMM-Pool__1078   1078     Tagged
            DCA-ACI-VMM-Pool__1079   1079     Tagged
            DCA-ACI-VMM-Pool__1080   1080     Tagged
            DCA-ACI-VMM-Pool__1081   1081     Tagged
            DCA-ACI-VMM-Pool__1082   1082     Tagged
            DCA-ACI-VMM-Pool__1083   1083     Tagged
            DCA-ACI-VMM-Pool__1084   1084     Tagged
            DCA-ACI-VMM-Pool__1085   1085     Tagged
            DCA-ACI-VMM-Pool__1086   1086     Tagged
            DCA-ACI-VMM-Pool__1087   1087     Tagged
            DCA-ACI-VMM-Pool__1088   1088     Tagged
            DCA-ACI-VMM-Pool__1089   1089     Tagged
            DCA-ACI-VMM-Pool__1090   1090     Tagged
            DCA-Host-mgmt2           150      Tagged
            DCA-VLAN-Pri             500      Tagged
            DCA-VLAN-Sec             510      Tagged
       UplinkPort:
            Enclosure                      IOBay   Port
            ---------                      -----   ----
            Middle - Frame 2 -CN75150484   6       Q3:1
            Middle - Frame 2 -CN75150484   6       Q4:1
            Top - Frame1 - CN7515010J      3       Q3:1
            Top - Frame1 - CN7515010J      3       Q4:1

  UplinkSet: Tech-Enablement-Testing
       Networks:
            Network Name     VlanID   Type
            ------------     ------   ----
            TE-Testing-300   300      Tagged
            TE-Testing-301   301      Tagged
            TE-Testing-302   302      Tagged
            TE-Testing-303   303      Tagged
            TE-Testing-304   304      Tagged
            vSAN_Network     2000     Tagged
       UplinkPort:
            Enclosure                      IOBay   Port
            ---------                      -----   ----
            Middle - Frame 2 -CN75150484   6       Q5
            Top - Frame1 - CN7515010J      3       Q5


Index   Enclosure                      IOBay   ModelName                                       PartNumber
1       Top - Frame1 - CN7515010J      3       Virtual Connect SE 40Gb F8 Module for Synergy   794502-B23
1       Top - Frame1 - CN7515010J      6       Synergy 20Gb Interconnect Link Module           779218-B21
2       Middle - Frame 2 -CN75150484   3       Synergy 20Gb Interconnect Link Module           779218-B21
2       Middle - Frame 2 -CN75150484   6       Virtual Connect SE 40Gb F8 Module for Synergy   794502-B23
3       CN7515048Q                     3       Synergy 20Gb Interconnect Link Module           779218-B21
3       CN7515048Q                     6       Synergy 20Gb Interconnect Link Module           779218-B21
```