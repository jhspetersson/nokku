# nokku
Port-knocking client. Supports multiple hosts knocking.

### Usage

    nokku HOSTNAME[:PORT] [HOSTNAME2:]PORT2 [[HOSTNAME3:]PORT3 [[HOSTNAME4:]PORT4 [...]]]
    
Supply a hostname first, optionally with a port in the form of `host:port`. 
Then goes any number of ports. Supply another hostname option to switch the target address.
  
### Examples

Same host, several ports:

    nokku 192.168.251.34 31785 23077 46254
    nokku 192.168.251.34:31785 23077 46254

Multiple hosts, several ports on each:

    nokku 192.168.251.1 20100 20123 192.168.251.2 14500 14600
    nokku 192.168.251.1:20100 20123 192.168.251.2:14500 14600

Multiple hosts, one port on each:

    nokku 192.168.251.1 20100 192.168.251.2 14500 192.168.251.3 33500
    nokku 192.168.251.1:20100 192.168.251.2:14500 192.168.251.3:33500
    
Multiple hosts, same port:

    nokku 192.168.251.1 44555 192.168.251.2 192.168.251.3
    nokku 192.168.251.1:44555 192.168.251.2 192.168.251.3