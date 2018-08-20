package covenant

// Code generated by github.com/CovenantSQL/HashStablePack DO NOT EDIT.

import (
	hsp "github.com/CovenantSQL/HashStablePack/marshalhash"
)

// MarshalHash marshals for hash
func (z Data) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	o = hsp.AppendBytes(o, []byte(z))
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Data) Msgsize() (s int) {
	s = hsp.BytesPrefixSize + len([]byte(z))
	return
}

// MarshalHash marshals for hash
func (z *HeaderTest) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 8
	o = append(o, 0x88, 0x88)
	o = hsp.AppendInt32(o, z.Version)
	o = append(o, 0x88)
	o = hsp.AppendString(o, z.TestName)
	o = append(o, 0x88)
	o = hsp.AppendBytes(o, z.TestArray)
	o = append(o, 0x88)
	if oTemp, err := z.Producer.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x88)
	o = hsp.AppendArrayHeader(o, uint32(len(z.GenesisHash)))
	for za0001 := range z.GenesisHash {
		if oTemp, err := z.GenesisHash[za0001].MarshalHash(); err != nil {
			return nil, err
		} else {
			o = hsp.AppendBytes(o, oTemp)
		}
	}
	o = append(o, 0x88)
	o = hsp.AppendArrayHeader(o, uint32(len(z.ParentHash)))
	for za0002 := range z.ParentHash {
		if z.ParentHash[za0002] == nil {
			o = hsp.AppendNil(o)
		} else {
			if oTemp, err := z.ParentHash[za0002].MarshalHash(); err != nil {
				return nil, err
			} else {
				o = hsp.AppendBytes(o, oTemp)
			}
		}
	}
	o = append(o, 0x88)
	if z.MerkleRoot == nil {
		o = hsp.AppendNil(o)
	} else {
		o = hsp.AppendArrayHeader(o, uint32(len(*z.MerkleRoot)))
		for za0003 := range *z.MerkleRoot {
			if (*z.MerkleRoot)[za0003] == nil {
				o = hsp.AppendNil(o)
			} else {
				if oTemp, err := (*z.MerkleRoot)[za0003].MarshalHash(); err != nil {
					return nil, err
				} else {
					o = hsp.AppendBytes(o, oTemp)
				}
			}
		}
	}
	o = append(o, 0x88)
	o = hsp.AppendTime(o, z.Timestamp)
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *HeaderTest) Msgsize() (s int) {
	s = 1 + 8 + hsp.Int32Size + 9 + hsp.StringPrefixSize + len(z.TestName) + 10 + hsp.BytesPrefixSize + len(z.TestArray) + 9 + z.Producer.Msgsize() + 12 + hsp.ArrayHeaderSize
	for za0001 := range z.GenesisHash {
		s += z.GenesisHash[za0001].Msgsize()
	}
	s += 11 + hsp.ArrayHeaderSize
	for za0002 := range z.ParentHash {
		if z.ParentHash[za0002] == nil {
			s += hsp.NilSize
		} else {
			s += z.ParentHash[za0002].Msgsize()
		}
	}
	s += 11
	if z.MerkleRoot == nil {
		s += hsp.NilSize
	} else {
		s += hsp.ArrayHeaderSize
		for za0003 := range *z.MerkleRoot {
			if (*z.MerkleRoot)[za0003] == nil {
				s += hsp.NilSize
			} else {
				s += (*z.MerkleRoot)[za0003].Msgsize()
			}
		}
	}
	s += 10 + hsp.TimeSize
	return
}

// MarshalHash marshals for hash
func (z *HeaderTest2) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 8
	o = append(o, 0x88, 0x88)
	o = hsp.AppendInt32(o, z.Version2)
	o = append(o, 0x88)
	o = hsp.AppendString(o, z.TestName2)
	o = append(o, 0x88)
	o = hsp.AppendBytes(o, z.TestArray2)
	o = append(o, 0x88)
	if oTemp, err := z.Producer2.MarshalHash(); err != nil {
		return nil, err
	} else {
		o = hsp.AppendBytes(o, oTemp)
	}
	o = append(o, 0x88)
	o = hsp.AppendArrayHeader(o, uint32(len(z.GenesisHash2)))
	for za0001 := range z.GenesisHash2 {
		if oTemp, err := z.GenesisHash2[za0001].MarshalHash(); err != nil {
			return nil, err
		} else {
			o = hsp.AppendBytes(o, oTemp)
		}
	}
	o = append(o, 0x88)
	o = hsp.AppendArrayHeader(o, uint32(len(z.ParentHash2)))
	for za0002 := range z.ParentHash2 {
		if z.ParentHash2[za0002] == nil {
			o = hsp.AppendNil(o)
		} else {
			if oTemp, err := z.ParentHash2[za0002].MarshalHash(); err != nil {
				return nil, err
			} else {
				o = hsp.AppendBytes(o, oTemp)
			}
		}
	}
	o = append(o, 0x88)
	if z.MerkleRoot2 == nil {
		o = hsp.AppendNil(o)
	} else {
		o = hsp.AppendArrayHeader(o, uint32(len(*z.MerkleRoot2)))
		for za0003 := range *z.MerkleRoot2 {
			if (*z.MerkleRoot2)[za0003] == nil {
				o = hsp.AppendNil(o)
			} else {
				if oTemp, err := (*z.MerkleRoot2)[za0003].MarshalHash(); err != nil {
					return nil, err
				} else {
					o = hsp.AppendBytes(o, oTemp)
				}
			}
		}
	}
	o = append(o, 0x88)
	o = hsp.AppendTime(o, z.Timestamp2)
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *HeaderTest2) Msgsize() (s int) {
	s = 1 + 9 + hsp.Int32Size + 10 + hsp.StringPrefixSize + len(z.TestName2) + 11 + hsp.BytesPrefixSize + len(z.TestArray2) + 10 + z.Producer2.Msgsize() + 13 + hsp.ArrayHeaderSize
	for za0001 := range z.GenesisHash2 {
		s += z.GenesisHash2[za0001].Msgsize()
	}
	s += 12 + hsp.ArrayHeaderSize
	for za0002 := range z.ParentHash2 {
		if z.ParentHash2[za0002] == nil {
			s += hsp.NilSize
		} else {
			s += z.ParentHash2[za0002].Msgsize()
		}
	}
	s += 12
	if z.MerkleRoot2 == nil {
		s += hsp.NilSize
	} else {
		s += hsp.ArrayHeaderSize
		for za0003 := range *z.MerkleRoot2 {
			if (*z.MerkleRoot2)[za0003] == nil {
				s += hsp.NilSize
			} else {
				s += (*z.MerkleRoot2)[za0003].Msgsize()
			}
		}
	}
	s += 11 + hsp.TimeSize
	return
}

// MarshalHash marshals for hash
func (z MyInt) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	o = hsp.AppendInt(o, int(z))
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z MyInt) Msgsize() (s int) {
	s = hsp.IntSize
	return
}

// MarshalHash marshals for hash
func (z Person1) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 3
	o = append(o, 0x83, 0x83)
	o = hsp.AppendString(o, z.Name1)
	o = append(o, 0x83)
	o = hsp.AppendInt(o, z.Age1)
	o = append(o, 0x83)
	o = hsp.AppendString(o, z.Address1)
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Person1) Msgsize() (s int) {
	s = 1 + 6 + hsp.StringPrefixSize + len(z.Name1) + 5 + hsp.IntSize + 9 + hsp.StringPrefixSize + len(z.Address1)
	return
}

// MarshalHash marshals for hash
func (z Person2) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 3
	o = append(o, 0x83, 0x83)
	o = hsp.AppendString(o, z.Name2)
	o = append(o, 0x83)
	o = hsp.AppendString(o, z.Address2)
	o = append(o, 0x83)
	o = hsp.AppendInt(o, z.Age2)
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Person2) Msgsize() (s int) {
	s = 1 + 6 + hsp.StringPrefixSize + len(z.Name2) + 9 + hsp.StringPrefixSize + len(z.Address2) + 5 + hsp.IntSize
	return
}

// MarshalHash marshals for hash
func (z *Struct) MarshalHash() (o []byte, err error) {
	var b []byte
	o = hsp.Require(b, z.Msgsize())
	// map header, size 3
	o = append(o, 0x83, 0x83)
	o = hsp.AppendMapHeader(o, uint32(len(z.Which)))
	for za0001, za0002 := range z.Which {
		o = hsp.AppendString(o, za0001)
		if za0002 == nil {
			o = hsp.AppendNil(o)
		} else {
			o = hsp.AppendInt(o, int(*za0002))
		}
	}
	o = append(o, 0x83)
	o = hsp.AppendBytes(o, []byte(z.Other))
	o = append(o, 0x83)
	o = hsp.AppendArrayHeader(o, uint32(Eight))
	for za0003 := range z.Nums {
		o = hsp.AppendFloat64(o, z.Nums[za0003])
	}
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Struct) Msgsize() (s int) {
	s = 1 + 6 + hsp.MapHeaderSize
	if z.Which != nil {
		for za0001, za0002 := range z.Which {
			_ = za0002
			s += hsp.StringPrefixSize + len(za0001)
			if za0002 == nil {
				s += hsp.NilSize
			} else {
				s += hsp.IntSize
			}
		}
	}
	s += 6 + hsp.BytesPrefixSize + len([]byte(z.Other)) + 5 + hsp.ArrayHeaderSize + (Eight * (hsp.Float64Size))
	return
}
