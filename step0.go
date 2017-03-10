package main

import (
	"encoding/hex"
	"fmt"
	"hash/fnv"
	"log"
	"reflect"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type info struct {
	id   []string
	vals []string
}

func (in *info) Stringify() string {
	return fmt.Sprintf("%v %v", hash(strings.Join(in.id, "_")), strings.Join(in.vals, " "))
}

func hash(s string) uint8 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return uint8(h.Sum32() % 256)
}

func (in *info) Add(value interface{}) {

	switch l := value.(type) {
	case (*layers.TCP):
		in.id = append(in.id, fmt.Sprintf("tcp_seq=%v;tcp=%v:%v", l.Seq, l.SrcPort, l.DstPort))
	case (*layers.UDP):
		in.id = append(in.id, fmt.Sprintf("udp=%v:%v", l.SrcPort, l.DstPort))
	case (*layers.ARP):
		in.id = append(in.id, fmt.Sprintf("arp=%v:%v", hex.EncodeToString(l.SourceHwAddress), hex.EncodeToString(l.DstHwAddress)))
	case (*layers.IPv4):
		in.vals = append(in.vals, fmt.Sprintf("%v", hash(l.SrcIP.String())))
		in.vals = append(in.vals, fmt.Sprintf("%v", hash(l.DstIP.String())))
	}

	reflected := reflect.ValueOf(value).Elem()
	if reflected.Type().Kind() != reflect.Struct {
		//		log.Println("Unsopported type", reflected.Type().Kind())
		return
	}

	for i := 0; i < reflected.NumField(); i++ {

		elem := reflected.Field(i)
		switch elem.Type().Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			in.vals = append(in.vals, fmt.Sprintf("%d", elem.Int()))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			in.vals = append(in.vals, fmt.Sprintf("%d", elem.Uint()))
		case reflect.Bool:
			var out uint8
			if elem.Bool() {
				out = 255
			}
			in.vals = append(in.vals, fmt.Sprintf("%d", out))
		default:
			//log.Println("Unknown field type -- %v", elem.Type().Kind())
		}
	}
}

func main() {
	handle, err := pcap.OpenOffline("/dev/stdin")
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		pkt := &info{}
		pkt.vals = []string{fmt.Sprintf("%v", packet.Metadata().Timestamp.Unix())}
		for _, l := range packet.Layers() {
			//fmt.Println("Addin layer", l.LayerType())
			pkt.Add(l)
		}
		fmt.Println(pkt.Stringify())
	}
}
