## containers under the hood
- containers don’t emulate hardware like VMs do. Instead, they share the host OS kernel, but isolate processes using kernel features.
https://www.youtube.com/watch?v=wJdDWc6zO4U

## Linux building blocks
containers rely primarily on these following things to create an container
- namespaces
- control group (cgroups)
- union mount filesystems 
- container runtimes
### namespaces
- Namespaces provide **isolation** of resources so that a container thinks it has its own environment.
- there are different kinds of namespaces that provide different kinds of isolations
PID namespace: Isolates process IDs.
NET namespace: Provides a separate network stack.
MNT namespace: Isolates filesystem mount points. what files and directories it can see
UTS namespace: Isolates hostname/domain.
IPC namespace: Isolates interprocess communication.
USER namespace: Allows different UID/GID mappings (e.g., container root ≠ host root).
Each container is given its own set of namespaces, making it feel like it's running on its own system.
### control group (cgroups)
- they limit and monitor resource usage (CPU,memory,io etc)
- Example: A container can be restricted to use only 512MB of RAM or 20% of CPU.
- If it exceeds its limits, the kernel can throttle or kill processes.
### union mount filesystems (UnionFS)
https://www.youtube.com/watch?v=DfENwtNRlD4

- UnionFS is a concept of merging multiple layers of filesystems.

- It's used in containers so they can reuse common image layers (read-only) and just write differences (copy-on-write).
- OverlayFS is a real Linux implementation of this.
- It's the reason containers are:
  - Fast to start
  - Storage-efficient
  - Layered and composable

eg:
```bash
sudo mount -t overlay overlay -o lowerdir=lower,upperdir=upper,workdir=work merged 
```

``` bash
lower/
  └── etc/
        └── hostname (value: "from lower")

upper/
  └── etc/
        └── hostname (value: "from upper")
        └── resolv.conf
```
after mounting:
```bash
ls merged/etc
hostname       ← from upperdir (hides lowerdir version)
resolv.conf    ← only exists in upper
```
If a file is:
  - In both upper and lower → upper wins (copy-on-write behavior)
  - Only in lower → visible 
  - Only in upper → visible
If you modify or delete a file:
  - A copy is made into the upper layer.
  - You never change the original lowerdir(it is readonly).
OverlayFS needs a scratch area to manage internal metadata during operations like file copy, rename, etc. thats the work dir here. It must:
Be on the same filesystem as upperdir
Be empty
It’s not visible from the merged view — it's just for the kernel.
**Imagine:**
  - lowerdir = a glass plate with writing on it
  - upperdir = a transparent sheet on top where you can write
  - merged = what you see when looking through both
  - If you want to change something on the glass plate, you write on the transparent sheet instead. That’s **copy-on-write** in action.
 