<p align="center"><img src="https://avatars2.githubusercontent.com/u/24763891?s=400&u=c1150e7da5667f47159d433d8e49dad99a364f5f&v=4" height="300" width="300"></p>
<p align="center"><h2><em>[Multiverse]</em> Go Operating System Library</h2></p>

**URL** [multiverse-os.org](https://multiverse-os.org)

The Multiverse Go operating system (os) library is a drop-in replacement for the
Go standard `os` library that provides a consistent, secure and only supports
posix operating systems (e.g. Linux, Unix, and Android). 

The library aims to provide a feature complete and lighter weight `os` standard
library alternative; specifically focusing of with supporting very easy to use
ultra low level abstractions with secure defaults but maintaing support for maximum
customability.

#### Why build an alternative to the Go standard 'os' library?
The Multiverse OS operating system is in essence a complex abstraction layer over
one or more bare-metal machines running Debian OS. The primary purpose of this
is to provide ephemeral compartmentalization based security ontop of Debian OS
in a user-friendly and intuitive way. 

**Complete progrommatic access**
To the various aspects of both kernel and user space is needed to satisfy the
design requirements of Multiverse OS. In addition to being a spec build of 
many Multiverse OS protocols, it will provide a framework for consistent 
minimum level of security for Multiverse OS software. 

The library is designed from the ground up with a emphasis on security,
internationalization support, seamless clustering across both mobile and desktop
systems, patching and building linux kernel from source, support for virtual 
devices (e.g. usb, pci, network, and more), and network API that focuses on 
providing tools to work at the packet level by default. 


#### Full hardware virtualization compartmentalization
Multiverse OS uses full hardware virtualization to compartmentalize the various
aspects of the machine. Hypervisor functionality is built into the fabric of the
`os` library. Active development of Multiverse leverages existing hypversors
`qemu` and intel's `nemu` but will be replaced with Multiverse' `singularity`
hypervisor supporting full hardware virtualization.

Multiverse OS provides unparallelled security, and privacy by virtualizating the
entire user interface and disabling and locking down the host machine entirely.
The user never directly interacts with the host machine of any bare-metal
machine in the supercomputer-like cluster of any number of bare-metal machines
(one or more).

Multiverse OS was designed to provide a secure computing environment for general
use providing a consistent traditional desktop interface that condenses the
resources of one or more computers. 

The host machines run routers to transparently route all traffic over overlay
networks, VPNs, and the onion network, and the user interfaces with the cluster
through a dedicated controller virtual machine (VM); all Multiverse VMs are 
deterministic, use trustless image building locally by processing very simple
to review scripts provided by the `portal-gun` virtual machine manger. 

Each deterministically bult VM provides ephemeral computing environments resulting
in a general use desktop environment that provides a secure environment by
disabling all access to the host machine and operating in an entirely virtualized
environment using nested virtual machines and pci-passthrough. 

**Reaching the highest levels of computing environment security using trustless
images and fully virtualized virtual machine based compartmentalization but
still preformant enough to play modern FPS games.**

#### Security through open source: BIOS, Bootloaders , and Initramfs
Nesting virtual machines provides security against malware, direct attacks by
compartmentalizing activities and using ephemeral environments to automatically
revert back to deterministically determined stable states. These states are
guaranteed to be secure by providing tools in the Multiverse `os` library to
programmatically build Multiverse VM BIOS, initiramfs, and kernels that are
signed with the computer operaters key not the developers key providing
*trustless* security.

Security starts at boot, signed boots that carry through multiverse initramfs
implementation are a key component in Multiverse OS security. The init system 
ensures BIOS, initramfs and the eventual kernel are not tampered with, that
decyrption of hard-disks and creation of session keys happen in secure
environments. 


#### Modularity
The Multiverse `os` library is designed to be interdependent as possible, so 
that the core subpackages can be used without needing to include the entire
`os` library. 

#### Real-time emulation and support
Multiverse OS aims to be a real-time operating system, 


### Getting Started
Getting started with using the Multiverse `os` library is as easily as importing
the specific subpackage required by your script or application and initializing
the objects.

Documentation is stil under active development, so until a rough
draft can be published, the best resources for learning how to use the library
is reading over the source code.

```
import(
  // fs watch library; example described below
  watch "github.com/multiverse-os/os/fs/watch"
)
```


### Example Usage: Command-line Tools
A variety of command-line tool examples are developed alongside the `os`
library are all found in the `cmd/` folder.

The current command-line tools provides are:

  `fswatch`: A simple file and folder watcher, made to watch for changes and
  execute a bash/sh command in response to all or specific changes, to all or
  sppecific files in a folder. The `watch` library which it relies on, is built
  to exclusively use inotify directly through syscalls (it does not require
  standard Go `os` library or any external packages.)

  `api`: An API server that provides programmatic access to a Server operating
  system over a variety of interfaces including:
      * **Filesystem** based API; for example, how `sysfs` or `procfs` allows
        programmatic access to various aspects of the operating system by
        interacting with system files.

      * **REST API** served over HTTP, or Socket, using a variety of data
        encoding formats such as `JSON`, `CBOR`, or `XML`. Using SSE for push
        notifications.

      * **GRPC** served in a variety of data encoding formats including `JSON`,
        `CBOR`, and `XML`. **Implementation is still under discussion.**
      

### Contribute
Contributions through pull requests are
reviewed and accepted if minimum design requirements are satisfied. 

Multiverse OS is currently a collective of scientists and artists volunteering
towards a open community effort to build secure operating systems using cutting
edge virtualization technology aimed at general computing desktop usage.

**Email** [contact@multiverse-os.org](mailto:contact@multiverse-os.org)
**URL** [multiverse-os.org](https://multiverse-os.org)
