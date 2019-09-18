# nokku
Port-knocking client. Supports multiple hosts knocking.

### Usage

    nokku HOSTNAME[:PORT] [PROTO] [pause] [HOSTNAME2:]PORT2 [pause] [[HOSTNAME3:]PORT3 [[HOSTNAME4:]PORT4 [...]]]
    
Supply a hostname first, optionally with a port in the form of `host:port`. 
Then goes any number of ports. Supply another hostname option to switch the target address.

Insert `tcp` (default) or `udp` command to switch to the corresponding protocol.

To add a pause to the sequence use `pause` (defaults to 1 second).
  
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
    
Switch to UDP and back to TCP:

    nokku 192.168.251.34 10100 udp 10200 tcp 10300
    
All UDP:

    nokku udp 192.168.251.1 20100 20123 192.168.251.2 14500 14600
    
Add pauses:

    nokku udp 192.168.251.1 20100 pause 20123 pause pause 192.168.251.2 14500 14600