package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"net/netip"
	"os"
)

var version = "0.0.0"
var help = fmt.Sprintf(Help, version)

func main() {
	var args = struct {
		Help    bool
		Version bool
	}{}

	flag.BoolVar(&args.Help, "h", false, "Help")
	flag.BoolVar(&args.Help, "help", false, "Alias for -h")
	flag.BoolVar(&args.Version, "version", false, "Version")
	flag.BoolVar(&args.Version, "version", false, "Alias for -v")
	flag.Parse()

	log.SetFlags(0)
	log.SetPrefix("Error: ")

	if args.Help {
		fmt.Print(help)
		os.Exit(0)
	}

	if args.Version {
		fmt.Print(version)
		os.Exit(0)
	}

	// Stdin
	var reader = csv.NewReader(os.Stdin)
	reader.Comma = ','
	reader.ReuseRecord = true
	reader.TrimLeadingSpace = true

	// Stdout
	var writer = csv.NewWriter(os.Stdout)
	writer.Comma = ','
	defer writer.Flush()

	is_header := true

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}

		// Skip the header line
		if is_header {
			if err = writer.Write(line); err != nil {
				log.Fatalln(err)
			}
			is_header = false
			continue
		}

		// Get IPs list
		ips_list, err := unwrapIP(line[0])
		if err != nil {
			log.Fatalln(err)
		}

		// Clone lines
		for _, ip := range ips_list {
			line[0] = ip
			if err = writer.Write(line); err != nil {
				log.Fatalln(err)
			}
		}

		writer.Flush()

		if err = writer.Error(); err != nil {
			log.Fatalln(err)
		}
	}
}

func unwrapIP(ip string) ([]string, error) {
	var err error
	var list []string

	p, err := netip.ParsePrefix(ip)
	if err != nil {
		return list, err
	}

	p = p.Masked()
	addr := p.Addr()
	for {
		if !p.Contains(addr) {
			break
		}
		list = append(list, addr.String())
		addr = addr.Next()
	}

	return list, err
}
