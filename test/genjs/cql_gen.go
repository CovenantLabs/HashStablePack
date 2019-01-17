// Code generated by github.com/CovenantSQL/HashStablePack DO NOT EDIT.
import * as hspEncoder from 'hsp-encoder'

// PersonMarshalHash function
export function PersonMarshalHash(z) {
	let binary = []

	// map header, size 13
	binary.push(0x8d,)

	// encode z.F1 with type Float32
	binary.push(hspEncoder.appendFloat32(z.F1))

	// encode z.F2 with type Float64
	binary.push(hspEncoder.appendFloat64(z.F2))

	// encode z.Age2 with type Int
	binary.push(hspEncoder.appendInt(z.Age2))

	// encode z.Int16 with type Int16
	binary.push(hspEncoder.appendInt16(z.Int16))

	// encode z.Int32 with type Int32
	binary.push(hspEncoder.appendInt32(z.Int32))

	// encode z.Int64 with type Int64
	binary.push(hspEncoder.appendInt64(z.Int64))

	// encode z.Int8 with type Int8
	binary.push(hspEncoder.appendInt8(z.Int8))

	// encode z.Name with type String
	binary.push(hspEncoder.appendString(z.Name))

	// encode z.Age with type Uint
	binary.push(hspEncoder.appendUint(z.Age))

	// encode z.Uint16 with type Uint16
	binary.push(hspEncoder.appendUint16(z.Uint16))

	// encode z.Uint32 with type Uint32
	binary.push(hspEncoder.appendUint32(z.Uint32))

	// encode z.Uint64 with type Uint64
	binary.push(hspEncoder.appendUint64(z.Uint64))

	// encode z.Uint8 with type Uint8
	binary.push(hspEncoder.appendUint8(z.Uint8))


	const hash = new Uint8Array(b)
	return hash
}

// PersonMsgsize function
export function PersonMsgSize(z) {
	let size = 0
	size = 1 + 3 + hspEncoder.typeSizes.Float32Size + 3 + hspEncoder.typeSizes.Float64Size + 5 + hspEncoder.typeSizes.IntSize + 6 + hspEncoder.typeSizes.Int16Size + 6 + hspEncoder.typeSizes.Int32Size + 6 + hspEncoder.typeSizes.Int64Size + 5 + hspEncoder.typeSizes.Int8Size + 5 + hspEncoder.typeSizes.StringPrefixSize + z.Name.length + 4 + hspEncoder.typeSizes.UintSize + 7 + hspEncoder.typeSizes.Uint16Size + 7 + hspEncoder.typeSizes.Uint32Size + 7 + hspEncoder.typeSizes.Uint64Size + 6 + hspEncoder.typeSizes.Uint8Size
	return size
}