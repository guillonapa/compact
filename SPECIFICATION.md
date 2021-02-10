# Technical Specification

## Languages and Tools

- Optimizer: Go.
- Compressor: Go.
- Client: Go.

## Client API

- `Optimize(image)`
- `OptimizeAs(image, optimizer.type)`
- `Compress(image, path)`
- `CompressAs(image, compressor.type, path)`
- `Decompress(file, path)`
- `DecompressAs(file, compressor.type, path)`
- `OptimizeAndCompress(image, path)`
- `OptimizeAndCompressAs(image, optimizer.type, compressor.type, path)`

## Optimizer API

- Optimizing types: `Web`, `Print`, `Mobile`
- `Optimize(ImageIr, type)` -> `ImageIr`
- `shrink(ImageIr, scale)` -> `ImageIr`
- `scalingFactor(ImageIr, type)` -> `float`

## Compressor API

- Compression types: `Simple`, `Experimental`, etc.
- `Compress(ImageIr, type`) -> `CompactImage`
- `Decompress(CompactImage)` -> `ImageIr`

## Internal

- `ImageIr`: `struct` for image internal representation
- `Read(image)` -> `ImageIr`
- `Draw(ImageIr)` -> `image`
