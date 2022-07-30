# EnvCheck
+ A tool for automatically detecting server environment information
+ It has only been verified on CentOS system for the time being

## Function list
+ Get HostName
  - Print host name
  - Check whether the host name meets the requirements (Only lowercase letters, numbers , < . > < - >)
+ Get CPU information
  - Architecture
  - Number of CPU cores
  - Whether the instruction set supports (avx2, AVX, bmi2) 
+ Get GPU information
  - Is there an NVIDIA graphics card
  - Number of graphics cards
  - Graphics card type
  - Is Nouveau installed
  - Whether the driver is installed
+ Get OS information
  - Operating system version
+ Get the host type
  - Virtual machine or physical machine
+ Get kernel information
  - Kernel version
+ Get memory information
  - Type
  - Size
+ Get network card information
  - Network card name
  - MAC address 
  - Rate
+ Get disk information
  - Size of system root path
  - Whether there is a mount point of the data volume </data>, the size of the mount point of </data>, and the type of Mount volume of </data>
  - Number of disks
  - Size of disk
  - Type of disk
+ Get the time zone
+ Turn off the firewall
+ Check whether docker is installed
+ Check whether k8s is installed
+ Get all ports currently in use

