# sensortag-go

```
$ sudo ./explorer xx:xx:xx:xx:xx:xx
2018/09/06 16:37:52 dev: hci0 up
2018/09/06 16:37:52 dev: hci0 reset
2018/09/06 16:37:52 dev: hci0 down
2018/09/06 16:37:52 dev: hci0 opened
State: PoweredOn
Scanning...
2018/09/06 16:37:52 DATA: [ 08 00 20 03 ]

Peripheral ID:xx:xx:xx:xx:xx:xx, NAME:(CC2650 SensorTag)
  Local Name        = CC2650 SensorTag
  TX Power Level    = 0
  Manufacturer Data = [13 0 3 0 0]
  Service Data      = []

Connected
Service: 1800 (Generic Access)
  Characteristic  2a00 (Device Name)
    properties    read
    handle        2
    value         4343323635302053656e736f72546167 | "CC2650 SensorTag"
  Characteristic  2a01 (Appearance)
    properties    read
    handle        4
    value         0000 | "\x00\x00"
  Characteristic  2a04 (Peripheral Preferred Connection Parameters)
    properties    read
    handle        6
    value         5000a0000000e803 | "P\x00\xa0\x00\x00\x00\xe8\x03"

Service: 1801 (Generic Attribute)

Service: 180a (Device Information)
  Characteristic  2a23 (System ID)
    properties    read
    handle        10
    value         46d05200000e6c54 | "F\xd0R\x00\x00\x0elT"
  Characteristic  2a24 (Model Number String)
    properties    read
    handle        12
    value         4343323635302053656e736f72546167 | "CC2650 SensorTag"
  Characteristic  2a25 (Serial Number String)
    properties    read
    handle        14
    value         4e2e412e | "N.A."
  Characteristic  2a26 (Firmware Revision String)
    properties    read
    handle        16
    value         312e343220284d6179203130203230313729 | "1.42 (May 10 2017)"
  Characteristic  2a27 (Hardware Revision String)
    properties    read
    handle        18
    value         4e2e412e | "N.A."
  Characteristic  2a28 (Software Revision String)
    properties    read
    handle        20
    value         4e2e412e | "N.A."
  Characteristic  2a29 (Manufacturer Name String)
    properties    read
    handle        22
    value         546578617320496e737472756d656e7473 | "Texas Instruments"
  Characteristic  2a2a (IEEE 11073-20601 Regulatory Certification Data List)
    properties    read
    handle        24
    value         fe006578706572696d656e74616c | "\xfe\x00experimental"
  Characteristic  2a50 (PnP ID)
    properties    read
    handle        26
    value         010d0000001001 | "\x01\r\x00\x00\x00\x10\x01"

Service: 180f (Battery Service)
  Characteristic  2a19 (Battery Level)
    properties    read notify
    handle        29
    value         56 | "V"
  Descriptor      2902 (Client Characteristic Configuration)
    handle        31
    value         0000 | "\x00\x00"
  Descriptor      2908 (Report Reference)
    handle        32
    value         0401 | "\x04\x01"
  Descriptor      2904 (Characteristic Presentation Format)
    handle        33
    value         0400ad27011131 | "\x04\x00\xad'\x01\x111"

Service: f000aa0004514000b000000000000000
  Characteristic  f000aa0104514000b000000000000000
    properties    read notify
    handle        35
    value         00000000 | "\x00\x00\x00\x00"
  Descriptor      2902 (Client Characteristic Configuration)
    handle        37
    value         0000 | "\x00\x00"
  Characteristic  f000aa0204514000b000000000000000
    properties    read write
    handle        38
    value         00 | "\x00"
  Characteristic  f000aa0304514000b000000000000000
    properties    read write
    handle        40
    value         64 | "d"

Service: f000aa2004514000b000000000000000
  Characteristic  f000aa2104514000b000000000000000
    properties    read notify
    handle        43
    value         00000000 | "\x00\x00\x00\x00"
  Descriptor      2902 (Client Characteristic Configuration)
    handle        45
    value         0000 | "\x00\x00"
  Characteristic  f000aa2204514000b000000000000000
    properties    read write
    handle        46
    value         00 | "\x00"
  Characteristic  f000aa2304514000b000000000000000
    properties    read write
    handle        48
    value         64 | "d"

Service: f000aa4004514000b000000000000000
  Characteristic  f000aa4104514000b000000000000000
    properties    read notify
    handle        51
    value         000000000000 | "\x00\x00\x00\x00\x00\x00"
  Descriptor      2902 (Client Characteristic Configuration)
    handle        53
    value         0000 | "\x00\x00"
  Characteristic  f000aa4204514000b000000000000000
    properties    read write
    handle        54
    value         00 | "\x00"
  Characteristic  f000aa4404514000b000000000000000
    properties    read write
    handle        56
    value         64 | "d"

Service: f000aa8004514000b000000000000000
  Characteristic  f000aa8104514000b000000000000000
    properties    read notify
    handle        59
    value         000000000000000000000000000000000000 | "\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"
  Descriptor      2902 (Client Characteristic Configuration)
    handle        61
    value         0000 | "\x00\x00"
  Characteristic  f000aa8204514000b000000000000000
    properties    read write
    handle        62
    value         0002 | "\x00\x02"
  Characteristic  f000aa8304514000b000000000000000
    properties    read write
    handle        64
    value         64 | "d"

Service: f000aa7004514000b000000000000000
  Characteristic  f000aa7104514000b000000000000000
    properties    read notify
    handle        67
    value         0000 | "\x00\x00"
  Descriptor      2902 (Client Characteristic Configuration)
    handle        69
    value         0000 | "\x00\x00"
  Characteristic  f000aa7204514000b000000000000000
    properties    read write
    handle        70
    value         00 | "\x00"
  Characteristic  f000aa7304514000b000000000000000
    properties    read write
    handle        72
    value         50 | "P"

Service: ffe0
  Characteristic  ffe1
    properties    notify
    handle        75
  Descriptor      2902 (Client Characteristic Configuration)
    handle        77
    value         0000 | "\x00\x00"
  Descriptor      2901 (Characteristic User Description)
    handle        78
    value         4b6579205072657373205374617465 | "Key Press State"

Service: f000aa6404514000b000000000000000
  Characteristic  f000aa6504514000b000000000000000
    properties    read write
    handle        80
    value         7e | "~"
  Characteristic  f000aa6604514000b000000000000000
    properties    read write
    handle        82
    value         00 | "\x00"

Service: f000ac0004514000b000000000000000
  Characteristic  f000ac0104514000b000000000000000
    properties    read write
    handle        85
    value         00000000 | "\x00\x00\x00\x00"
  Characteristic  f000ac0204514000b000000000000000
    properties    read write
    handle        87
    value         0424160020 | "\x04$\x16\x00 "
  Characteristic  f000ac0304514000b000000000000000
    properties    read write
    handle        89
    value         0500 | "\x05\x00"

Service: f000ccc004514000b000000000000000
  Characteristic  f000ccc104514000b000000000000000
    properties    read notify
    handle        92
    value         060000000a00 | "\x06\x00\x00\x00\n\x00"
  Descriptor      2902 (Client Characteristic Configuration)
    handle        94
    value         0000 | "\x00\x00"
  Characteristic  f000ccc204514000b000000000000000
    properties    write
    handle        95
  Characteristic  f000ccc304514000b000000000000000
    properties    write
    handle        97

Service: f000ffc004514000b000000000000000
  Characteristic  f000ffc104514000b000000000000000
    properties    writeWithoutResponse write notify
    handle        100
  Descriptor      2902 (Client Characteristic Configuration)
    handle        102
    value         0000 | "\x00\x00"
  Descriptor      2901 (Characteristic User Description)
    handle        103
    value         496d67204964656e74696679 | "Img Identify"
  Characteristic  f000ffc204514000b000000000000000
    properties    writeWithoutResponse write notify
    handle        104
  Descriptor      2902 (Client Characteristic Configuration)
    handle        106
    value         0000 | "\x00\x00"
  Descriptor      2901 (Characteristic User Description)
    handle        107
    value         496d6720426c6f636b | "Img Block"
  Characteristic  f000ffc304514000b000000000000000
    properties    writeWithoutResponse write
    handle        108
  Descriptor      2901 (Characteristic User Description)
    handle        110
    value         496d6720436f756e74 | "Img Count"
  Characteristic  f000ffc404514000b000000000000000
    properties    read notify
    handle        111
    value         00 | "\x00"
  Descriptor      2902 (Client Characteristic Configuration)
    handle        113
    value         0000 | "\x00\x00"
  Descriptor      2901 (Characteristic User Description)
    handle        114
    value         496d6720537461747573 | "Img Status"

Waiting for 5 seconds to get some notifiations, if any.
```
