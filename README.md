# sensortag-go

```
$ sudo ./explorer xx:xx:xx:xx:xx:xx
2018/09/05 14:42:54 dev: hci0 up
2018/09/05 14:42:54 dev: hci0 reset
2018/09/05 14:42:54 dev: hci0 down
2018/09/05 14:42:54 dev: hci0 opened
State: PoweredOn
Scanning...
2018/09/05 14:42:54 DATA: [ 08 00 20 03 ]

Peripheral ID:xx:xx:xx:xx:xx:xx, NAME:(CC2650 SensorTag)
  Local Name        = CC2650 SensorTag
  TX Power Level    = 0
  Manufacturer Data = [13 0 3 0 0]
  Service Data      = []

Connected
Service: 1800 (Generic Access)
  Characteristic  2a00 (Device Name)
    properties    read
    value         4343323635302053656e736f72546167 | "CC2650 SensorTag"
  Characteristic  2a01 (Appearance)
    properties    read
    value         0000 | "\x00\x00"
  Characteristic  2a04 (Peripheral Preferred Connection Parameters)
    properties    read
    value         5000a0000000e803 | "P\x00\xa0\x00\x00\x00\xe8\x03"

Service: 1801 (Generic Attribute)

Service: 180a (Device Information)
  Characteristic  2a23 (System ID)
    properties    read
    value         46d05200000e6c54 | "F\xd0R\x00\x00\x0elT"
  Characteristic  2a24 (Model Number String)
    properties    read
    value         4343323635302053656e736f72546167 | "CC2650 SensorTag"
  Characteristic  2a25 (Serial Number String)
    properties    read
    value         4e2e412e | "N.A."
  Characteristic  2a26 (Firmware Revision String)
    properties    read
    value         312e343220284d6179203130203230313729 | "1.42 (May 10 2017)"
  Characteristic  2a27 (Hardware Revision String)
    properties    read
    value         4e2e412e | "N.A."
  Characteristic  2a28 (Software Revision String)
    properties    read
    value         4e2e412e | "N.A."
  Characteristic  2a29 (Manufacturer Name String)
    properties    read
    value         546578617320496e737472756d656e7473 | "Texas Instruments"
  Characteristic  2a2a (IEEE 11073-20601 Regulatory Certification Data List)
    properties    read
    value         fe006578706572696d656e74616c | "\xfe\x00experimental"
  Characteristic  2a50 (PnP ID)
    properties    read
    value         010d0000001001 | "\x01\r\x00\x00\x00\x10\x01"

Service: 180f (Battery Service)
  Characteristic  2a19 (Battery Level)
    properties    read notify
    value         56 | "V"
  Descriptor      2902 (Client Characteristic Configuration)
    value         0000 | "\x00\x00"
  Descriptor      2908 (Report Reference)
    value         0401 | "\x04\x01"
  Descriptor      2904 (Characteristic Presentation Format)
    value         0400ad27011131 | "\x04\x00\xad'\x01\x111"

Service: f000aa0004514000b000000000000000
  Characteristic  f000aa0104514000b000000000000000
    properties    read notify
    value         00000000 | "\x00\x00\x00\x00"
  Descriptor      2902 (Client Characteristic Configuration)
    value         0000 | "\x00\x00"
  Characteristic  f000aa0204514000b000000000000000
    properties    read write
    value         00 | "\x00"
  Characteristic  f000aa0304514000b000000000000000
    properties    read write
    value         64 | "d"

Service: f000aa2004514000b000000000000000
  Characteristic  f000aa2104514000b000000000000000
    properties    read notify
    value         00000000 | "\x00\x00\x00\x00"
  Descriptor      2902 (Client Characteristic Configuration)
    value         0000 | "\x00\x00"
  Characteristic  f000aa2204514000b000000000000000
    properties    read write
    value         00 | "\x00"
  Characteristic  f000aa2304514000b000000000000000
    properties    read write
    value         64 | "d"

Service: f000aa4004514000b000000000000000
  Characteristic  f000aa4104514000b000000000000000
    properties    read notify
    value         000000000000 | "\x00\x00\x00\x00\x00\x00"
  Descriptor      2902 (Client Characteristic Configuration)
    value         0000 | "\x00\x00"
  Characteristic  f000aa4204514000b000000000000000
    properties    read write
    value         00 | "\x00"
  Characteristic  f000aa4404514000b000000000000000
    properties    read write
    value         64 | "d"

Service: f000aa8004514000b000000000000000
  Characteristic  f000aa8104514000b000000000000000
    properties    read notify
    value         000000000000000000000000000000000000 | "\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"
  Descriptor      2902 (Client Characteristic Configuration)
    value         0000 | "\x00\x00"
  Characteristic  f000aa8204514000b000000000000000
    properties    read write
    value         0002 | "\x00\x02"
  Characteristic  f000aa8304514000b000000000000000
    properties    read write
    value         64 | "d"

Service: f000aa7004514000b000000000000000
  Characteristic  f000aa7104514000b000000000000000
    properties    read notify
    value         0000 | "\x00\x00"
  Descriptor      2902 (Client Characteristic Configuration)
    value         0000 | "\x00\x00"
  Characteristic  f000aa7204514000b000000000000000
    properties    read write
    value         00 | "\x00"
  Characteristic  f000aa7304514000b000000000000000
    properties    read write
    value         50 | "P"

Service: ffe0
  Characteristic  ffe1
    properties    notify
  Descriptor      2902 (Client Characteristic Configuration)
    value         0000 | "\x00\x00"
  Descriptor      2901 (Characteristic User Description)
    value         4b6579205072657373205374617465 | "Key Press State"

Service: f000aa6404514000b000000000000000
  Characteristic  f000aa6504514000b000000000000000
    properties    read write
    value         7e | "~"
  Characteristic  f000aa6604514000b000000000000000
    properties    read write
    value         00 | "\x00"

Service: f000ac0004514000b000000000000000
  Characteristic  f000ac0104514000b000000000000000
    properties    read write
    value         00000000 | "\x00\x00\x00\x00"
  Characteristic  f000ac0204514000b000000000000000
    properties    read write
    value         0424160020 | "\x04$\x16\x00 "
  Characteristic  f000ac0304514000b000000000000000
    properties    read write
    value         0500 | "\x05\x00"

Service: f000ccc004514000b000000000000000
  Characteristic  f000ccc104514000b000000000000000
    properties    read notify
    value         060000000a00 | "\x06\x00\x00\x00\n\x00"
  Descriptor      2902 (Client Characteristic Configuration)
    value         0000 | "\x00\x00"
  Characteristic  f000ccc204514000b000000000000000
    properties    write
  Characteristic  f000ccc304514000b000000000000000
    properties    write

Service: f000ffc004514000b000000000000000
  Characteristic  f000ffc104514000b000000000000000
    properties    writeWithoutResponse write notify
  Descriptor      2902 (Client Characteristic Configuration)
    value         0000 | "\x00\x00"
  Descriptor      2901 (Characteristic User Description)
    value         496d67204964656e74696679 | "Img Identify"
  Characteristic  f000ffc204514000b000000000000000
    properties    writeWithoutResponse write notify
  Descriptor      2902 (Client Characteristic Configuration)
    value         0000 | "\x00\x00"
  Descriptor      2901 (Characteristic User Description)
    value         496d6720426c6f636b | "Img Block"
  Characteristic  f000ffc304514000b000000000000000
    properties    writeWithoutResponse write
  Descriptor      2901 (Characteristic User Description)
    value         496d6720436f756e74 | "Img Count"
  Characteristic  f000ffc404514000b000000000000000
    properties    read notify
    value         00 | "\x00"
  Descriptor      2902 (Client Characteristic Configuration)
    value         0000 | "\x00\x00"
  Descriptor      2901 (Characteristic User Description)
    value         496d6720537461747573 | "Img Status"

Waiting for 5 seconds to get some notifiations, if any.
Disconnected
Done
```
