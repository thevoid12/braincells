# cryptography 
- cryptography is a vvast subject. here are few notes about it

- Cryptography is the science of securing information. It ensures confidentiality, integrity, and authenticity of data through various techniques. It is mainly categorized into:

    - Symmetric encryption: same key is used to encrypt and decrypt.

    - Asymmetric encryption: uses a key pair:public and private keys.

## Symmetric Encryption
- The same key is used for both encryption and decryption.
- It is fast and suitable for encrypting large volumes of data.
- Requires secure key sharing between sender and receiver.
Common examples: AES, ChaCha20

### AES:
- AES is a widely used symmetric encryption algorithm standardized by NIST.
- Operates on fixed-size blocks of 128 bits
- Key sizes: 128, 192, or 256 bits
- Highly secure and efficient in both hardware and software
- Used in applications like VPNs, disk encryption, HTTPS (TLS), and file encryption

AES does not define how to handle message integrity or authentication — for that, AEAD modes are often used. for aead [aead.md](./aead.md)

## Asymmetric Encryption
- Two keys are used public key (shared) and a private key (kept secret). 
- Data encrypted with the public key can only be decrypted with the private key.
- Slower than symmetric encryption but useful for key exchange, digital signatures, and authentication.
- Common examples: RSA, Ed25519, Elliptic Curve Cryptography (ECC)
  
### rsa:
- RSA is a widely-used asymmetric encryption algorithm based on the difficulty of factoring large prime numbers.
- Key sizes: usually 2048 or 4096 bits
- Supports both encryption and digital signatures
- Commonly used to encrypt small payloads (like symmetric keys), not large data

### ed25519
- Ed25519 is a digital signature algorithm based on elliptic curve cryptography.
- Provides strong security with small key sizes
- Extremely fast for signing and verifying
- Resistant to side-channel attacks
- Used in SSH (OpenSSH), TLS, code signing, and cryptocurrencies
- It is a modern and recommended alternative to RSA for digital signatures.

## SHA:
### SHA-256
- SHA-256 is a cryptographic hash function that:
- Outputs a fixed 256-bit (32-byte) hash
- Is deterministic, collision-resistant, and one-way
- Commonly used in digital signatures, password hashing, blockchain, and integrity checks

### SHA-512
- SHA-512 is similar to SHA-256 but produces a 512-bit (64-byte) output.
- Offers higher security margin due to longer output
- Slightly slower but more robust for high-security applications

---
1. How are SHA and RSA related?
They serve different purposes in cryptography, but are often used together in protocols like digital signatures.

RSA / Ed25519:

These are public-key cryptographic algorithms.

Used to generate a public-private key pair.

They are used for encryption, decryption, and digital signatures.

SHA256 / SHA512:

These are cryptographic hash functions.

Used to produce a fixed-size hash from input data (file, message, etc.).

They do not use keys and are not reversible.

How they work together (e.g., in digital signatures):
You hash the input (like a file or message) using SHA256 to produce a fixed-size digest.

Then, sign the hash with your private key using RSA or Ed25519.

On the other side, someone with your public key can verify the signature by:

Hashing the received file again

Checking if it matches the decrypted signature (i.e., the hash you originally signed)

So yes — RSA and Ed25519 work with key pairs; SHA does not.

2. Is SHA256 just hashing the input?
Yes.

SHA256 takes any input (text, file, bytes) and outputs a 256-bit (32-byte) hash.

It’s deterministic: same input → same hash.

It's not encryption. You cannot reverse it.

3. How does hashing an entire file with SHA256 work? What’s the input?
When you run sha256sum file.txt, the hash function takes the raw bytes of the file as input.

Even if one character in the file changes, the hash output will be completely different. This property is known as the avalanche effect.

Why hash a file?
To check data integrity

To compare two files (e.g., for duplicates)