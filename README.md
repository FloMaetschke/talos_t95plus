## Attempt for T95Plus Support

### Notes:

```
root@node0:~# parted /dev/mmcblk2
GNU Parted 3.4
Using /dev/mmcblk2
Welcome to GNU Parted! Type 'help' to view a list of commands.
(parted) p
Model: MMC SEM64G (sd/mmc)
Disk /dev/mmcblk2: 62.5GB
Sector size (logical/physical): 512B/512B
Partition Table: gpt
Disk Flags:

Number  Start   End     Size    File system  Name      Flags
 1      4194kB  8389kB  4194kB               security
 2      8389kB  12.6MB  4194kB               uboot
 3      12.6MB  16.8MB  4194kB               trust
 4      16.8MB  62.5GB  62.5GB  ext4                   boot, esp

root@node0:~# fdisk -l /dev/mmcblk2
Disk /dev/mmcblk2: 58.24 GiB, 62537072640 bytes, 122142720 sectors
Units: sectors of 1 * 512 = 512 bytes
Sector size (logical/physical): 512 bytes / 512 bytes
I/O size (minimum/optimal): 512 bytes / 512 bytes
Disklabel type: gpt
Disk identifier: 67060000-0000-4B7E-8000-117E00004A3F

Device         Start       End   Sectors  Size Type
/dev/mmcblk2p1  8192     16383      8192    4M unknown
/dev/mmcblk2p2 16384     24575      8192    4M unknown
/dev/mmcblk2p3 24576     32767      8192    4M unknown
/dev/mmcblk2p4 32768 122142686 122109919 58.2G EFI System

```
UBoot Partition 2, Offset 16384

Talos Partition start at: 32768



<!-- markdownlint-disable MD041 -->

<p align="center">
  <h1 align="center">Talos Linux</h1>
  <p align="center">A modern OS for Kubernetes.</p>
  <p align="center">
    <a href="https://github.com/talos-systems/talos/releases/latest">
      <img alt="Release" src="https://img.shields.io/github/release/talos-systems/talos.svg?logo=github&logoColor=white&style=flat-square">
    </a>
    <a href="https://github.com/talos-systems/talos/releases/latest">
      <img alt="Pre-release" src="https://img.shields.io/github/release-pre/talos-systems/talos.svg?label=pre-release&logo=GitHub&logoColor=white&style=flat-square">
    </a>
  </p>
</p>

---

**Talos** is a modern OS for running Kubernetes: secure, immutable, and minimal.
Talos is fully open source, production-ready, and supported by the people at [Sidero Labs](https://www.SideroLabs.com/)
All system management is done via an API - there is no shell or interactive console.
Benefits include:

- **Security**: Talos reduces your attack surface: It's minimal, hardened, and immutable.
  All API access is secured with mutual TLS (mTLS) authentication.
- **Predictability**: Talos eliminates configuration drift, reduces unknown factors by employing immutable infrastructure ideology, and delivers atomic updates.
- **Evolvability**: Talos simplifies your architecture, increases your agility, and always delivers current stable Kubernetes and Linux versions.

## Documentation

For instructions on deploying and managing Talos, see the [Documentation](https://www.talos.dev/docs/latest/).

## Community

- Slack: Join our [slack channel](https://slack.dev.talos-systems.io)
- Support: Questions, bugs, feature requests [GitHub Discussions](https://github.com/talos-systems/talos/discussions)
- Forum: [community](https://groups.google.com/a/SideroLabs.com/forum/#!forum/community)
- Twitter: [@SideroLabs](https://twitter.com/SideroLabs)
- Email: [info@SideroLabs.com](mailto:info@SideroLabs.com)

If you're interested in this project and would like to help in engineering efforts or have general usage questions, we are happy to have you!
We hold a weekly meeting that all audiences are welcome to attend.

We would appreciate your feedback so that we can make Talos even better!
To do so, you can take our [survey](https://docs.google.com/forms/d/1TUna5YTYGCKot68Y9YN_CLobY6z9JzLVCq1G7DoyNjA/edit).

### Office Hours

- When: Mondays at 16:30 UTC.
- Where: [Google Meet](https://meet.google.com/day-pxhv-zky).

You can subscribe to this meeting by joining the community forum above.

> Note: You can convert the meeting hours to your [local time](https://everytimezone.com/s/6bb1045a).

## Contributing

Contributions are welcomed and appreciated!
See [Contributing](CONTRIBUTING.md) for our guidelines.

## License

<a href="https://github.com/talos-systems/talos/blob/master/LICENSE">
  <img alt="GitHub" src="https://img.shields.io/github/license/talos-systems/talos?style=flat-square">
</a>

Some software we distribute is under the General Public License family
of licenses or other licenses that require we provide you with the
source code.
If you would like a copy of the source code for this
software, please contact us via email: info at SideroLabs.com.
