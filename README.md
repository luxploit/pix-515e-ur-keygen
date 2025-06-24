# [luxploit](https://luxploit.net)/[pixlic](https://github.com/luxploit/pixlic)

A simple keygen for generating unrestricted licenses (UR), including activation for 3DES and AES Encryption for VPNs, for Cisco PIX 500 series firewall appliances aswell as "maxxed-out" licenses for Cisco ASA 5500 (non -X) series firewall appliances.

## Installation

Either clone this repo and run `./build.sh` in Git Bash for Windows or your local \*nix shell or download a pre-compiled version found in the releases tab [here](https://github.com/luxploit/pixlic/releases)

## Example Usage

**Run `./pixlic.lxb -list` to see all available PIX models**

Example for PIX 515e:

```
./pixlic.lxb PIX -serial "xxxxxxxxx" -model "PIX 515"
```

Example for ASA 5505:

```
./pixlic.lxb ASA -serial "xxxxxxxxx" -model "ASA 5505"
```

## Credits

- Algorithm Leak: [AntiCisco poster "Pinguin"](http://www.anticisco.ru/forum/viewtopic.php?f=2&t=920)
- Initial find that led to this project being possible: [Dominic's blog about the AntiCisco thread](https://blog.bjdch.org/2015/01/1976)
- PIX 525 UR Serial and Key for testing: [RuTracker's Massive PIX/ASA collection](https://rutracker.org/forum/viewtopic.php?t=831309)
- PIX 501 Future and more goodies: @\_stargo\_ over on the Serial Port discord
- Inspiration that led to this project: [ClabRetro's \"PIX Firewall Failover\" video](https://www.youtube.com/watch?v=719c5ed_ogc)
