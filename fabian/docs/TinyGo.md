
[List of supported boards](https://tinygo.org/docs/reference/microcontrollers)
[Linux install guide](https://tinygo.org/getting-started/install/linux/)
[Drivers](https://github.com/tinygo-org/drivers/)
[Devices](https://tinygo.org/docs/reference/devices/)

A compiled go binary is 1MB+, which includes the whole standard library.
These small boards have size limits and standard Go is obviously too big.

TinyGo is a new compiler and a new runtime implementation optimized for tiny results (and less general usage).

Besides small boards, it's also interesting for WASM (larger .wasm files increase spin up times. And we will be spinning up .wasm instead of single binary containers on Kubernetes in the future).