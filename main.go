package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"golang.org/x/net/context"

	"github.com/currantlabs/ble"
	"github.com/currantlabs/ble/examples/lib/dev"
)

var (
	device = flag.String("device", "default", "implementation of ble")
	sd     = flag.Duration("sd", 5*time.Second, "scanning duration, 0 for indefinitely")
)

const name = "CC2650 SensorTag"

func main() {
	flag.Parse()

	d, err := dev.NewDevice(*device)
	if err != nil {
		log.Fatalf("can't new device : %s", err)
	}
	ble.SetDefaultDevice(d)

	// Default to search device with name of Gopher (or specified by user).
	filter := func(a ble.Advertisement) bool {
		return strings.ToUpper(a.LocalName()) == strings.ToUpper(name)
	}

	// Scan for specified durantion, or until interrupted by user.
	fmt.Printf("Scanning for %s...\n", *sd)
	ctx := ble.WithSigHandler(context.WithTimeout(context.Background(), *sd))
	cln, err := ble.Connect(ctx, filter)
	if err != nil {
		log.Fatalf("can't connect : %s", err)
	}

	mtu, err := cln.ExchangeMTU(512)
	if err != nil {
		log.Fatalf("can't exchange MTU : %s", err)
	}
	log.Printf("exchange MTU : %d", mtu)

	// Make sure we had the chance to print out the message.
	done := make(chan struct{})
	// Normally, the connection is disconnected by us after our exploration.
	// However, it can be asynchronously disconnected by the remote peripheral.
	// So we wait(detect) the disconnection in the go routine.
	go func() {
		<-cln.Disconnected()
		fmt.Printf("[ %s ] is disconnected \n", cln.Address())
		close(done)
	}()

	fmt.Printf("Discovering profile...\n")
	p, err := cln.DiscoverProfile(true)
	if err != nil {
		log.Fatalf("can't discover profile: %s", err)
	}

	deviceName(p, cln)

	// Disconnect the connection. (On OS X, this might take a while.)
	fmt.Printf("Disconnecting [ %s ]... (this might take up to few seconds on OS X)\n", cln.Address())
	cln.CancelConnection()

	<-done
}

func deviceName(profile *ble.Profile, client ble.Client) (string, error) {
	resp, err := readCharacteristics(profile, client, "1800", "2a00")
	log.Print("[read ] ", resp)
	return resp, err
}

func readCharacteristics(profile *ble.Profile,
	client ble.Client,
	service string,
	characteristics string) (string, error) {
	sUUID, err := ble.Parse(service)
	if err != nil {
		return "", err
	}
	cUUID, err := ble.Parse(characteristics)
	if err != nil {
		return "", err
	}
	c, err := getCharacteristics(profile, sUUID, cUUID)
	if err != nil {
		return "", err
	}

	buf, err := client.ReadCharacteristic(c)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func getCharacteristics(profile *ble.Profile, service []byte, characteristics []byte) (*ble.Characteristic, error) {
	for _, s := range profile.Services {
		if reflect.DeepEqual(service, s.UUID) {
			continue
		}
		for _, c := range s.Characteristics {
			if reflect.DeepEqual(characteristics, c.UUID) {
				continue
			}
			return c, nil
		}
	}
	return nil, errors.New("not found Characteristic")
}
