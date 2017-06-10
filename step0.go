package main

import (
	"encoding/hex"
	"fmt"
	"hash/fnv"
	"log"
	//	"math"
	"reflect"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type info struct {
	id   []string
	vals []string
}

func (in *info) Stringify() string {
	return fmt.Sprintf("%v %v", hash32(strings.Join(in.id, "_")), strings.Join(in.vals, " "))
}

func hash8(s string) uint8 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return uint8(h.Sum32() % 256)
}

func hash32(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

var (
	total  = map[string]float64{}
	totalN = map[string]float64{}
)

func addVal(to []string, name string, val int64) []string {

	/*
		v := float64(val)
		total[name] += v
		totalN[name]++

		var dist float64

		if v != 0 {
			dist = math.Sqrt(math.Abs((total[name] / totalN[name]) / v))
		}
	*/

	return append(to, fmt.Sprintf("%f", float32(val)))
}

func (in *info) Add(value interface{}) {

	var valToAnalyze interface{}

	switch l := value.(type) {
	case (*layers.ICMPv4):
		in.id = append(in.id, fmt.Sprintf("icmp_id=%v;seq=%v", l.Id, l.Seq))
		valToAnalyze = l
	case (*layers.IPv4):
		in.id = append(in.id, fmt.Sprintf("ip=%v:%v", l.SrcIP.String(), l.DstIP.String()))
		in.vals = append(in.vals, fmt.Sprintf("%v", hash8(l.SrcIP.String())))
		in.vals = append(in.vals, fmt.Sprintf("%v", hash8(l.DstIP.String())))
		valToAnalyze = l
	case (*layers.TCP):
		in.id = append(in.id, fmt.Sprintf("seq=%v;tcp=%v:%v", l.Seq, l.SrcPort, l.DstPort))
		valToAnalyze = l
	case (*layers.UDP):
		in.id = append(in.id, fmt.Sprintf("udp=%v:%v", l.SrcPort, l.DstPort))
		valToAnalyze = l
	case (*layers.ARP):
		in.id = append(in.id, fmt.Sprintf("arp=%v:%v", hex.EncodeToString(l.SourceHwAddress), hex.EncodeToString(l.DstHwAddress)))
		valToAnalyze = l
	default:
		// skip any other network layer for now to save spaace
		return
	}

	if valToAnalyze == nil {
		return
	}

	reflected := reflect.ValueOf(valToAnalyze).Elem()
	if reflected.Type().Kind() != reflect.Struct {
		//		log.Println("Unsopported type", reflected.Type().Kind())
		return
	}

	tp := reflected.Type()

	for i := 0; i < reflected.NumField(); i++ {
		elem := reflected.Field(i)
		elemName := tp.Field(i).Name

		switch elem.Type().Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			in.vals = addVal(in.vals, elemName, elem.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			in.vals = addVal(in.vals, elemName, int64(elem.Uint()))
		case reflect.Bool:
			if true { // tooooo many vals; tooo slow
				var out uint8
				if elem.Bool() {
					out = 1
				}
				in.vals = append(in.vals, fmt.Sprintf("%d", out))
			}
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

	var first *time.Time

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {

		if first == nil {
			first = &packet.Metadata().Timestamp
		}

		pkt := &info{}
		pkt.vals = []string{fmt.Sprintf("%v", -first.Sub(packet.Metadata().Timestamp).Seconds())}
		for _, l := range packet.Layers() {
			//fmt.Println("Addin layer", l.LayerType())
			pkt.Add(l)
		}
		fmt.Println(pkt.Stringify())
	}
}
