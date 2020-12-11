# Santa's Beamer

Since reindeer went on strike... Santa loves his Beamer (https://www.youtube.com/watch?v=IiXEenC3Ew4). 

But damnit... he haven´t turned it on for quite a while and forgot where his keys are. But due to new technologies he can now put through his phone app direct access to the ECUs to the Internet!!!

His car was explicitly build for him and he suspects the ignition sequence has something to do with xmas! But all this protocol weirdness is to much for him to handle. He first discovered that the ignition ECU accepts UDS messages:

- https://scapy.readthedocs.io/en/latest/api/scapy.contrib.automotive.uds.html
- https://automotive.softing.com/fileadmin/sof-files/pdf/de/ae/poster/UDS_Faltposter_softing2016.pdf

But since it is a beamer... this message is not wrapped by DoIP ... rather by High-Speed-Fahrzeug-Zugang (HSFZ):
https://scapy.readthedocs.io/en/latest/api/scapy.contrib.automotive.bmw.hsfz.html

Santa could also figure out the ECU device identifier is 0x37 and it can only be reached from the sender id of 0x13

Can you help turn the car on?

## Running

Insert flag in Dockerfile.

```bash
docker build -t dayxx_sbeamer .
sudo docker run -d -p xx:6801 dayxx_sbeamer
```

## Tipp

```bash
echo -e '\x00\x00\x00<length>\x00\x01\x13\x37<service bytes>'| nc <service> <port>
```

Todo scapy solution

## Solution

```bash
echo -e '\x00\x00\x00\x04\x00\x01\x13\x37\x24\x12'| nc 127.0.0.1 6801
```
