## what is virtualization?
goated video: https://www.youtube.com/watch?v=zLJbP6vBk2M
- Virtualization is the process of creating a virtual (rather than physical) version of something, such as a server, storage device, network resource, or operating system. 
- Virtualization lets you share the hardware resources of the underlying physical server to be consumed by virtual servers running on it.
- we sepearate the os out of the physical hardware so that they can  move around

## Evolution of Virtualization:
### bare metal:
![bare metail](img/1.png)
- simple and easy way. only way before virtualization
- bare metal is installing everything directly on our physical server
- With a bare metal system, the operating system, binaries/libraries, and applications are installed and run directly onto the physical hardware.
- this can lead to some problems like:
  - dependency hell: assume 2 application runs in the same machine and both needs golang but different version. this will create a problem in setting up and lot of configuration needs to be changed to make both the application work.
  - Low utilization efficiency
  - Large blast radius: if something failed all the application will crash
  - Slow start up & shut down speed (minutes): beacause we need to setup all the application
  - Very slow provisioning & decommissioning: because of all these complexity it takes time to provision and deprovision
![alt text](img/2.png)

### virtualization
![alt text](img/3.png)
- if we wanna move to a new physical server, we need to set up the same old hypervisor and that all we can copy the os in quickly in minutes unlike bare metal where everything needs to be setup from scratch based on the new physical hardware
- here we can run mulitple instances of different os on top of the hypervisor. 
- example if we wanna run a windows program in mac, i will install a hypervisor virutual box and install windows as one instance and run (type 2)

### types of hypervisor
- type 1 
- type 2

### type1 hypervisor:
![alt text](img/4.png)

- A Type 1 hypervisor, also known as a bare-metal hypervisor, is a virtualization layer that runs directly on the physical hardware of the host machine without an underlying operating system.
- Runs directly on hardware: No OS in between.
- Examples
  - VMware ESXi
  - Microsoft Hyper-V (in server mode)
  - Xen (Citrix Hypervisor)
  - KVM (Linux-based, sometimes debated but generally considered Type 1 due to kernel integration)
- when you log into a type1 hv unlike type2 since there is no os you have very few details in them like just the ip address of hypervisor. so things are manages with the help of another machine with management software through which you connect to the hypervisor with the ip and manage the server
![alt text](img/5.png)
- for example we pay for VMware ESXi software which installs hv in physical servers and they earn money in giving additional features while management console like moving os instances between hypervisors, move to free rack when an instance have heavy load etc etc
#### overprovisioning hypervisors
- note that we can overprovisioned hypervisors which is recommended
- An overprovisioned hypervisor refers to a situation where a hypervisor is configured to allocate more virtual resources to its VMs than the physical hardware actually has usually in terms of CPU, RAM, or storage.
- this is under the assumption that not all VMs will fully utilize their assigned resources at the same time.
eg:
You have a physical server with:
8 physical CPU cores
32 GB RAM
VM Name	vCPUs Assigned	RAM Assigned
WebServer1	4 vCPUs	8 GB	
WebServer2	4 vCPUs	8 GB	
DBServer	6 vCPUs	12 GB
AppServer	4 vCPUs	8 GB
otal Virtual Resources Assigned
vCPUs: 4 + 4 + 6 + 4 = 18 vCPUs
RAM: 8 + 8 + 12 + 8 = 36 GB
eventhough we over provisioned it is immposible to allocate more than what it has.because of this over provisioning, suppose db server load increases, it can use upto 12GB thus it can pull off ram from other places to maintain. 
note that this will create a problem if everything tries to exceed 
### type2 hypervisor:
- A Type 2 hypervisor, also known as a hosted hypervisor, is virtualization software that runs on top of a host operating system (OS) like Windows, macOS, or Linux. It relies on the host OS to access hardware resources such as CPU, memory, and disk.
- Think of it as "an app that creates and manages virtual machines inside your regular OS."
- install a os, install a hypervizor in top of os and multiple instance of os 
- eg: if we wanna run a windows program in mac, i will install a hypervisor virutual box and install windows as one instance and run 
![alt text](img/6.png)
**cons:**
- Slower than Type 1 due to extra OS layer.
- Host OS + hypervisor + guest OS = more RAM/CPU usage.
- Vulnerable if host OS is compromised.
