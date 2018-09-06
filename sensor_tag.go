package main

import (
	"errors"
	"log"

	"github.com/currantlabs/ble"
)

type SensorTag struct {
	profile *ble.Profile
	client  ble.Client
}

func NewSensorTag(profile *ble.Profile, client ble.Client) *SensorTag {
	return &SensorTag{profile: profile, client: client}
}

func (s *SensorTag) DeviceName() (string, error) {
	resp, err := readCharacteristics(s.profile, s.client, "1800", "2a00")
	value := string(resp)
	log.Print("DeviceName")
	log.Print("[read ] ", value)
	return string(value), err
}

func (s *SensorTag) FirmwareRevision() (string, error) {
	resp, err := readCharacteristics(s.profile, s.client, "180a", "2a26")
	value := string(resp)
	log.Print("Firmware Revision")
	log.Print("[read ] ", value)
	return string(value), err
}

func (s *SensorTag) IRTemperature() ([]byte, error) {
	resp, err := readCharacteristics(s.profile,
		s.client,
		"f000aa0004514000b000000000000000",
		"f000aa0104514000b000000000000000")
	log.Print("IR Temperature")
	log.Print("[read ] ", resp)
	return resp, err
}

func (s *SensorTag) EnableIRTemperature() error {
	log.Print("enable IR Temperature")
	return s.writeIRTemperatureConfig([]byte{1, 0})
}

func (s *SensorTag) DisableIRTemperature() error {
	log.Print("disable IR Temperature")
	return s.writeIRTemperatureConfig([]byte{0, 0})
}

func (s *SensorTag) IsEnableIRTemperature() (bool, error) {
	value, err := readDescriptor(s.profile,
		s.client,
		"f000aa0004514000b000000000000000",
		"f000aa0104514000b000000000000000",
		"2902")
	log.Print("is enable IR Temperature?")
	log.Print("[read ] ", value)
	return true, err
}

func (s *SensorTag) writeIRTemperatureConfig(value []byte) error {
	err := writeDescriptor(s.profile,
		s.client,
		"f000aa0004514000b000000000000000",
		"f000aa0104514000b000000000000000",
		"2902",
		value)
	log.Print("[write] ", value)

	return err
}

func (s *SensorTag) KeyStateCharacteristic() (*ble.Characteristic, error) {
	sUUID, err := ble.Parse("ffe0")
	if err != nil {
		log.Print("err: ", err)
		return nil, err
	}
	cUUID, err := ble.Parse("ffe1")
	if err != nil {
		log.Print("err: ", err)
		return nil, err
	}
	return getCharacteristics(s.profile, sUUID, cUUID)
}

func (s *SensorTag) EnableKeyState() error {
	log.Print("enable KeyState")
	return s.writeKeyStateConfig([]byte{1, 0})
}

func (s *SensorTag) DisableKeyState() error {
	log.Print("disable KeyState")
	return s.writeKeyStateConfig([]byte{0, 0})
}

func (s *SensorTag) IsEnableKeyState() (bool, error) {
	value, err := readDescriptor(s.profile,
		s.client,
		"ffe0",
		"ffe1",
		"2902")
	log.Print("is enable Key State?")
	log.Print("[read ] ", value)
	return true, err
}

func (s *SensorTag) writeKeyStateConfig(value []byte) error {
	err := writeDescriptor(s.profile,
		s.client,
		"ffe0",
		"ffe1",
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