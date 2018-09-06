package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
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
	firmwareRevision(p, cln)
	enableIRTemperature(p, cln)
	isEnableIRTemperature(p, cln)
	irTemperature(p, cln)
	time.Sleep(time.Second)
	irTemperature(p, cln)
	time.Sleep(time.Second)
	irTemperature(p, cln)
	time.Sleep(time.Second)
	irTemperature(p, cln)

	// Disconnect the connection. (On OS X, this might take a while.)
	fmt.Printf("Disconnecting [ %s ]... (this might take up to few seconds on OS X)\n", cln.Address())
	cln.CancelConnection()

	<-done
}

func deviceName(profile *ble.Profile, client ble.Client) (string, error) {
	resp, err := readCharacteristics(profile, client, "1800", "2a00")
	s := string(resp)
	log.Print("DeviceName")
	log.Print("[read ] ", s)
	return s, err
}

func firmwareRevision(profile *ble.Profile, client ble.Client) (string, error) {
	resp, err := readCharacteristics(profile, client, "180a", "2a26")
	s := string(resp)
	log.Print("Firmware Revision")
	log.Print("[read ] ", s)
	return s, err
}

func irTemperature(profile *ble.Profile, client ble.Client) ([]byte, error) {
	resp, err := readCharacteristics(profile,
		client,
		"f000aa0004514000b000000000000000",
		"f000aa0104514000b000000000000000")
	log.Print("IR Temperature")
	log.Print("[read ] ", resp)
	return resp, err
}

func enableIRTemperature(profile *ble.Profile, client ble.Client) error {
	log.Print("enable IR Temperature")
	return writeIRTemperature(profile, client, []byte{1, 0})
}

func disableIRTemperature(profile *ble.Profile, client ble.Client) error {
	log.Print("disable IR Temperature")
	return writeIRTemperature(profile, client, []byte{0, 0})
}

func isEnableIRTemperature(profile *ble.Profile, client ble.Client) (bool, error) {
	value, err := readDescriptor(profile,
		client,
		"f000aa0004514000b000000000000000",
		"f000aa0104514000b000000000000000",
		"2902")
	log.Print("is enable IR Temperature?")
	log.Print("[read ] ", value)
	return true, err
}

func writeIRTemperature(profile *ble.Profile, client ble.Client, value []byte) error {
	err := writeDescriptor(profile,
		client,
		"f000aa0004514000b000000000000000",
		"f000aa0104514000b000000000000000",
		"2902",
		value)
	log.Print("[write] ", value)

	return err
}

func readCharacteristics(profile *ble.Profile,
	client ble.Client,
	service string,
	characteristics string) ([]byte, error) {
	sUUID, err := ble.Parse(service)
	if err != nil {
		log.Print("err: ", err)
		return nil, err
	}
	cUUID, err := ble.Parse(characteristics)
	if err != nil {
		log.Print("err: ", err)
		return nil, err
	}
	c, err := getCharacteristics(profile, sUUID, cUUID)
	if err != nil {
		log.Print("err: ", err)
		return nil, err
	}

	buf, err := client.ReadCharacteristic(c)
	if err != nil {
		log.Print("err: ", err)
		return nil, err
	}
	return buf, nil
}

func writeCharacteristics(profile *ble.Profile,
	client ble.Client,
	service string,
	characteristics string,
	value []byte) error {
	sUUID, err := ble.Parse(service)
	if err != nil {
		log.Print("err: ", err)
		return err
	}
	cUUID, err := ble.Parse(characteristics)
	if err != nil {
		log.Print("err: ", err)
		return err
	}
	c, err := getCharacteristics(profile, sUUID, cUUID)
	if err != nil {
		log.Print("err: ", err)
		return err
	}

	err = client.WriteCharacteristic(c, value, false)
	if err != nil {
		log.Print("err: ", err)
		return err
	}
	return nil
}

func writeDescriptor(profile *ble.Profile,
	client ble.Client,
	service string,
	characteristics string,
	descriptor string,
	value []byte) error {
	sUUID, err := ble.Parse(service)
	if err != nil {
		log.Print("err: ", err)
		return err
	}
	cUUID, err := ble.Parse(characteristics)
	if err != nil {
		log.Print("err: ", err)
		return err
	}
	dUUID, err := ble.Parse(descriptor)
	if err != nil {
		log.Print("err: ", err)
		return err
	}
	d, err := getDescriptor(profile, sUUID, cUUID, dUUID)
	if err != nil {
		log.Print("err: ", err)
		return err
	}

	err = client.WriteDescriptor(d, value)
	if err != nil {
		log.Print("err: ", err)
		return err
	}
	return nil
}

func readDescriptor(profile *ble.Profile,
	client ble.Client,
	service string,
	characteristics string,
	descriptor string) ([]byte, error) {
	sUUID, err := ble.Parse(service)
	if err != nil {
		log.Print("err: ", err)
		return nil, err
	}
	cUUID, err := ble.Parse(characteristics)
	if err != nil {
		log.Print("err: ", err)
		return nil, err
	}
	dUUID, err := ble.Parse(descriptor)
	if err != nil {
		log.Print("err: ", err)
		return nil, err
	}
	d, err := getDescriptor(profile, sUUID, cUUID, dUUID)
	if err != nil {
		log.Print("err: ", err)
		return nil, err
	}

	value, err := client.ReadDescriptor(d)
	if err != nil {
		log.Print("err: ", err)
		return nil, err
	}
	return value, nil
}

func getCharacteristics(profile *ble.Profile, service []byte, characteristics []byte) (*ble.Characteristic, error) {
	for _, s := range profile.Services {
		if !s.UUID.Equal(service) {
			continue
		}
		for _, c := range s.Characteristics {
			if !c.UUID.Equal(characteristics) {
				continue
			}
			return c, nil
		}
	}
	return nil, errors.New("not found Characteristic")
}

func getDescriptor(profile *ble.Profile, service []byte, characteristics []byte, descriptor []byte) (*ble.Descriptor, error) {
	c, err := getCharacteristics(profile, service, characteristics)
	if err != nil {
		return nil, err
	}

	for _, d := range c.Descriptors {
		if !d.UUID.Equal(descriptor) {
			continue
		}
		return d, nil
	}
	return nil, errors.New("not found Descriptor")
}
