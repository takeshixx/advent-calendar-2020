# Santa Operations Center (SOC) forensic support

Santas diligent helpers received non-xmas communication from the future(pcapng) in the Santa Operations Center (SOC). Apparently the grinch is going to send a file to Santas XMAS service causing a POLYGLOTAL problem on 24th of December. 
Santas investigators need your help! Can you FIND and EXTRACT the file grinch send to the XMAS SERVICE and inspect its contents...? 

Hint 1: Take a look at HTTP communications on PORT 24

Hint 2: Wireshark and Networkminer both can extract HTTP Objects from PCAP

Hint 3: Things are not as they seem...inspect the FILE you extracted thoroughly,think "polyglotal"


## Resources
- https://cloudshark.io/articles/making-custom-captures-examples/
- https://truepolyglot.hackade.org/
- https://www.rubyguides.com/2012/01/four-ways-to-extract-files-from-pcaps/

## Running
docker build -t day07_PCAPP .
sudo docker run -dit -p 7:80 dayxx_PCAPP

## Solution

- Download PCAP
- Find HTTP communication on Port 24
- Export HTTP Objects with Wireshark, Networkminer,binwalk...
- Open extracted file with zip(encrypted, needs password) --> open file with pdf reader --> see zip password
- Extract encrypted zip with password "p0ly6l0t5_4r3_phun" and find the flag.txt in the extracted data

```
encrypted ZIP Password: p0ly6l0t5_4r3_phun
```


