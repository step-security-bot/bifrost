// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/aperturerobotics/bifrost/util/blockcompress/blockcompress.proto

package blockcompress

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// BlockCompress sets the type of compression to use.
type BlockCompress int32

const (
	// BlockCompress_NONE indicates no compression.
	BlockCompress_BlockCompress_NONE BlockCompress = 0
	// BlockCompress_SNAPPY indicates Snappy compression.
	BlockCompress_BlockCompress_SNAPPY BlockCompress = 1
	// BlockCompress_S2 indicates S2 compression.
	//
	// S2 is an extension of snappy. S2 is aimed for high throughput, which is why
	// it features concurrent compression for bigger payloads. Decoding is
	// compatible with Snappy compressed content, but content compressed with S2
	// cannot be decompressed by Snappy. This means that S2 can seamlessly replace
	// Snappy without converting compressed content. S2 is designed to have high
	// throughput on content that cannot be compressed. This is important so you
	// don't have to worry about spending CPU cycles on already compressed data.
	//
	// Reference: https://github.com/klauspost/compress/tree/master/s2
	BlockCompress_BlockCompress_S2 BlockCompress = 2
	// BlockCompress_LZ4 indicates LZ4 compression.
	BlockCompress_BlockCompress_LZ4 BlockCompress = 3
	// BlockCompress_ZSTD indicates z-standard compression.
	//
	// Zstandard is a real-time compression algorithm, providing high compression
	// ratios. It offers a very wide range of compression / speed trade-off, while
	// being backed by a very fast decoder. A high performance compression
	// algorithm is implemented.
	BlockCompress_BlockCompress_ZSTD BlockCompress = 4
)

var BlockCompress_name = map[int32]string{
	0: "BlockCompress_NONE",
	1: "BlockCompress_SNAPPY",
	2: "BlockCompress_S2",
	3: "BlockCompress_LZ4",
	4: "BlockCompress_ZSTD",
}

var BlockCompress_value = map[string]int32{
	"BlockCompress_NONE":   0,
	"BlockCompress_SNAPPY": 1,
	"BlockCompress_S2":     2,
	"BlockCompress_LZ4":    3,
	"BlockCompress_ZSTD":   4,
}

func (x BlockCompress) String() string {
	return proto.EnumName(BlockCompress_name, int32(x))
}

func (BlockCompress) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9cc39bccf4a1e784, []int{0}
}

func init() {
	proto.RegisterEnum("blockcompress.BlockCompress", BlockCompress_name, BlockCompress_value)
}

func init() {
	proto.RegisterFile("github.com/aperturerobotics/bifrost/util/blockcompress/blockcompress.proto", fileDescriptor_9cc39bccf4a1e784)
}

var fileDescriptor_9cc39bccf4a1e784 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0x4a, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x2c, 0x48, 0x2d, 0x2a, 0x29, 0x2d, 0x4a, 0x2d,
	0xca, 0x4f, 0xca, 0x2f, 0xc9, 0x4c, 0x2e, 0xd6, 0x4f, 0xca, 0x4c, 0x2b, 0xca, 0x2f, 0x2e, 0xd1,
	0x2f, 0x2d, 0xc9, 0xcc, 0xd1, 0x4f, 0xca, 0xc9, 0x4f, 0xce, 0x4e, 0xce, 0xcf, 0x2d, 0x28, 0x4a,
	0x2d, 0x2e, 0x46, 0xe5, 0xe9, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xf1, 0xa2, 0x08, 0x6a, 0xb5,
	0x31, 0x72, 0xf1, 0x3a, 0x81, 0x44, 0x9c, 0xa1, 0x22, 0x42, 0x62, 0x5c, 0x42, 0x28, 0x02, 0xf1,
	0x7e, 0xfe, 0x7e, 0xae, 0x02, 0x0c, 0x42, 0x12, 0x5c, 0x22, 0xa8, 0xe2, 0xc1, 0x7e, 0x8e, 0x01,
	0x01, 0x91, 0x02, 0x8c, 0x42, 0x22, 0x5c, 0x02, 0x68, 0x32, 0x46, 0x02, 0x4c, 0x42, 0xa2, 0x5c,
	0x82, 0xa8, 0xa2, 0x3e, 0x51, 0x26, 0x02, 0xcc, 0x98, 0xc6, 0x47, 0x05, 0x87, 0xb8, 0x08, 0xb0,
	0x24, 0xb1, 0x81, 0x9d, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x16, 0xcf, 0x08, 0xec,
	0x00, 0x00, 0x00,
}