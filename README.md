# ITLDIMS Program

## Setup
- [Install etcd](https://etcd.io/docs/v3.4/install/) and create a single node etcd (locally, if needed) by running the `etcd` command
- Clone the [itldims directory](https://github.com/yash-anand-fosteringlinux/itldims-cmd/tree/main/itldims)
- In the [etcd-inventory](https://github.com/yash-anand-fosteringlinux/CommandsAndCodes/tree/main/yes1/etcd-inventory) directory, `go run` [main.go](https://github.com/yash-anand-fosteringlinux/itldims-cmd/blob/main/itldims/main.go) for connecting with the API
- Run the `itldims` related commands from the cmd directory.
- If needed, Refer to the [Spreadsheet](https://docs.google.com/spreadsheets/d/1_oHivMUs1j4XZFSn3yZTKNkx50YXNYqU/edit#gid=1907009990) which consists of all the data that is uploaded to DB.

## Topics
- Setup
- Workings of the code
- Command Combinations
- - get
- - create 

# Workings of the code
- This code utilises a modifed version of  [main.go](https://github.com/yash-anand-fosteringlinux/Commands-and-Outputs/blob/main/Old-Keys-Input/main.go), where the etcd API url `localhost:8181/servers/` is connected with for displaying all the key-values.
- Data from the API Server is fetched and then the parsing of the data is done before the user inputs their argument/s to process the data from the API Server.
- `itldims` command is used to check connection with the API Server and `itldims get` subcommand is used to search user arguments from the API Server.
-  The method of placing user arguments into `localhost:8181/servers/<ServerType>/<ServerIP>/<Attribute>` is not used and grep like search is run through the data of `localhost:8181/servers`.
- If needed, user can search with a single key-component or value using `itldims get <input 1>`. The input entered by the user are then searched for and the key-values not needed are filtered out from the data in `localhost:8181/servers/`.

# GET Command Combinations
| S. No. | Command Combination               | Output Description                                      | Use-Case |
|-------|-----------------------------------|---------------------------------------------------------|------------|
| 1| `itldims get servers`              | Displays all the running Servers with their Server IPs | Helps user see all the running servers |
| 2| `itldims get types`       | Displays all the running Server Types | Helps User find all the server types |
| 3| `itldims get attributes`         | Displays all Attributes   | Helps User find all the attributes |
| 4| `itldims get <Server IP>`         | Displays all Attribute values of a specific Server IP  | Helps user find values of a specific server IP|
| 5| `itldims get <server Type>`         | Displays values of a specific server Type  | User can find values of a specific server Type |
| 6| `itldims get <Attribute>`   | Displays values of a specific Attribute   | User can find all the RAMs of all servers |
| 7| `itldims get <Server IP/Type> <attribute>` | Displays all value of attribute from specific server Type/IP | User can find if any attribute is 'None' on '10.249.221.22' |
| 8| `itldims get <Server IP/Type/Attribute> <Value>` | Displays all Server IPs containing a specific value  | User can find if any attribute is 'None' on '10.249.221.22' |
| 9| `itldims get <key>` | Displays value of a key, when the key is mentioned | User can find the value directly from the entire key |

## Outputs Of 'get' Commands
The possible combinations along with their outputs for the `itldims get` command have been provided below. For any output which is too lengthy, `. . . . .` has been used at the end to signify that the mentioned output gives complete data but is not being shown here completely.

**1. `itldims get servers`to get the following output:**
```
10.246.40.139
----------------------------
10.246.40.152
----------------------------
10.246.40.142
----------------------------
10.249.221.22
----------------------------
10.249.221.21
----------------------------
```

**2. `itldims get types` to get the following output:**
```
Physical
----------------------------
VM
----------------------------
```

**3. `itldims get attributes` to get the following output:**
```
LVM
----------------------------
NFS
----------------------------
Hostname
----------------------------
Gateway
----------------------------
PV
----------------------------
External_Disk
----------------------------
RAM
----------------------------
API
----------------------------
Internal_Partition
----------------------------
CPU
----------------------------
Environment
----------------------------
Netmask
----------------------------
External_Partition
----------------------------
data
----------------------------
Application
----------------------------
Internal_Disk
----------------------------
VG
----------------------------
OS
----------------------------
```

**4. `itldims get <Server IP>` or `itldims get 10.249.221.22`  to get the following output:**
```
Environment:
Production
----------------------------

PV:
PV Name=/dev/sda3
PV Size=101.00g
PV Name=/dev/sdb
PV Size=500.00g
----------------------------

Netmask:
255.255.255.128
----------------------------

RAM:
32GB
----------------------------

Gateway:
10.249.221.1
----------------------------

External_Partition:
u01:322GB
----------------------------

CPU:
8
----------------------------
. . . . .
```
    
**5. `itldims get <Server Type>` or `itldims get VM` to get the following output:**
```
Server IP: 10.249.221.21
Application:checkpost
----------------------------

Server IP: 10.249.221.22
OS:RHEL 8.7
----------------------------

Server IP: 10.249.221.22
External_Partition:u01:322GB
----------------------------
. . . . .
```

**6. `itldims get <Attribute>` or `itldims get RAM` to get the following output:**
```
Server IP: 10.249.221.22
RAM:32GB
----------------------------

Server IP: 10.246.40.142
RAM:32GB
----------------------------

Server IP: 10.249.221.21
RAM:32GB
----------------------------

Server IP: 10.246.40.139
RAM:32GB
----------------------------

Server IP: 10.246.40.152
RAM:32GB
----------------------------
```
    
**7. `itldims get <Server IP/Type> <attribute>`  or `itldims get 10.249.221.22 RAM`  to get the following output:**
```
RAM:
32GB
----------------------------
```

         
**8. `itldims get <Server IP/Type> <attribute>`  or `itldims get RAM 32GB`  to get the following output:**
```
Server IP: 10.246.40.152
RAM:32GB
----------------------------

Server IP: 10.249.221.21
RAM:32GB
----------------------------

Server IP: 10.249.221.22
RAM:32GB
----------------------------

Server IP: 10.246.40.142
RAM:32GB
----------------------------

Server IP: 10.246.40.139
RAM:32GB
---------------------------
```

**9. `itldims get <key>` or `itldims get /servers/VM/10.249.221.22/RAM` to get the following output:**
```
Key: /servers/VM/10.249.221.22/RAM
Value: 32GB
----------------------------
```

# CREATE Command Combinations
| S. No. | Command Combination               | Output Description                                      | Use-Case |
|-------|-----------------------------------|---------------------------------------------------------|------------|
| 1| `itldims create -k <key> -v <value>` | Posts the mentioned value to the specified key | Helps user post values of keys |

## Outputs Of 'create' Commands
The possible combinations along with their outputs for the `itldims create` command have been provided below.
**1. `itldims create -k <key> -v <value>` or `itldims create -k /servers/Physical/10.246.40.139/Hostname -v vahanapp18` to get the following output:**
``` 
Key: /servers/Physical/10.246.40.139/Hostname has been metered as vahanapp18 succesfully
```
