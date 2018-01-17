# Simple CLI for testing the go-hass library

A simple CLI for running commands with the [go-hass](https://github.com/pawal/go-hass) library.
Very limited support for devices, it is basically 'on' and 'off' that is supported. No support for groups.

## Example usage

List supported devices:
```bash
$ ./hasse 
light.kitchen_lights (Kitchen Lights): on
lock.kitchen_door (Kitchen Door): locked
light.ceiling_lights (Ceiling Lights): on
light.bed_light (Bed Light): off
lock.front_door (Front Door): locked
switch.decorative_lights (Decorative Lights): on
switch.ac (AC): off
```

Check status on a supported device:
```
$ ./hasse light.kitchen_lights
light.kitchen_lights state: on
```

Run on or off command on a supported device:
```
$ ./hasse light.kitchen_lights off
```
