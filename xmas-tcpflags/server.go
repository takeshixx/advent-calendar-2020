// Copyright 2012 Google, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file in the root of the source
// tree.

// arpscan implements ARP scanning of all interfaces' local networks using
// gopacket and its subpackages.  This example shows, among other things:
//   * Generating and sending packet data
//   * Reading in packet data and interpreting it
//   * Use of the 'pcap' subpackage for reading/writing
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	ifacePtr := flag.String("iface", "lo", "interface")
	flag.Parse()

	// Get a list of all interfaces.
	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	for _, iface := range ifaces {
		if iface.Name != *ifacePtr {
			continue
		}
		wg.Add(1)
		// Start up a scan on each interface.
		go func(iface net.Interface) {
			defer wg.Done()
			if err := scan(&iface); err != nil {
				log.Printf("interface %v: %v", iface.Name, err)
			}
		}(iface)
	}
	// Wait for all interfaces' scans to complete.  They'll try to run
	// forever, but will stop on an error, so if we get past this Wait
	// it means all attempts to write have failed.
	wg.Wait()
}

// scan scans an individual interface's local network for machines using ARP requests/replies.
//
// scan loops forever, sending packets out regularly.  It returns an error if
// it's ever unable to write a packet.
func scan(iface *net.Interface) error {
	// We just look for IPv4 addresses, so try to find if the interface has one.
	//var addr *net.IPNet

	// Open up a pcap handle for packet reads/writes.
	handle, err := pcap.OpenLive(iface.Name, 65536, true, pcap.BlockForever)
	if err != nil {
		return err
	}
	defer handle.Close()

	// Start up a goroutine to read in packet data.
	stop := make(chan struct{})
	go readPackets(handle, iface, stop)
	defer close(stop)
	for {
		/* 		*/
		// We don't know exactly how long it'll take for packets to be
		// sent back to us, but 10 seconds should be more than enough
		// time ;)
		fmt.Println("Sleeping 10")
		time.Sleep(10 * time.Second)
	}
}

func readPackets(handle *pcap.Handle, iface *net.Interface, stop chan struct{}) {
	src := gopacket.NewPacketSource(handle, layers.LayerTypeEthernet)
	in := src.Packets()
	for {
		var packet gopacket.Packet
		select {
		case <-stop:
			return
		case packet = <-in:
			ipLayer := packet.Layer(layers.LayerTypeIPv4)
			if ipLayer == nil {
				fmt.Printf("maybe IPv6 packet? Skipping")
				continue
			}
			ip := ipLayer.(*layers.IPv4)
			tcpLayer := packet.Layer(layers.LayerTypeTCP)
			if tcpLayer == nil {
				fmt.Printf("Invalid TCP layer found, skipping\n")
				continue
			}
			tcp := tcpLayer.(*layers.TCP)

			if ip != nil && tcp != nil {
				fmt.Printf("IP and TCP layer there\n")
				if tcp.DstPort != 24 {
					continue
				}
				if tcp.FIN && tcp.PSH && tcp.URG {
					fmt.Printf("Found a proper XMAS packet!\n")
					// Write our scan packets out to the handle.
					if err := sendICMP(handle, iface, ip.SrcIP.String()); err != nil {
						log.Printf("error writing packets on %v: %v", iface.Name, err)
					}
				} else {
					fmt.Printf("IP and TCP found, but not a proper XMAS packet!\n")
				}
			} else {
				fmt.Printf("No IP or TCP layer found\n")
			}
		}
	}
}

func sendICMP(handle *pcap.Handle, iface *net.Interface, addr string) error {
	fmt.Printf("Source MAC: %v\n", iface.HardwareAddr)
	if iface.Name == "lo" {
		na, _ := net.ParseMAC("00:00:00:00:00:00")
		iface.HardwareAddr = na
	}
	eth := layers.Ethernet{
		SrcMAC:       iface.HardwareAddr,
		DstMAC:       net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		EthernetType: layers.EthernetTypeIPv4,
	}
	addrs, err := iface.Addrs()
	if err != nil {
		return err
	}
	sip, _, _ := net.ParseCIDR(addrs[0].String())
	ip := layers.IPv4{
		SrcIP:    sip,
		DstIP:    net.ParseIP(addr),
		Version:  4,
		TTL:      64,
		Protocol: layers.IPProtocolICMPv4,
	}
	icmp := layers.ICMPv4{
		TypeCode: layers.CreateICMPv4TypeCode(layers.ICMPv4TypeDestinationUnreachable, layers.ICMPv4CodeCommAdminProhibited),
	}
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
	payload := []byte("XMAS TIME! Take your token: " + os.Getenv("XMAS_SECRET"))
	if err := gopacket.SerializeLayers(buf, opts, &eth, &ip, &icmp, gopacket.Payload(payload)); err != nil {
		return err
	}
	if err := handle.WritePacketData(buf.Bytes()); err != nil {
		return err
	}
	return nil
}
