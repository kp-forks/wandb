//go:build !noasm && gc && arm64 && !amd64
// +build !noasm,gc,arm64,!amd64

package sha1cd

import (
	"math"
	"unsafe"

	shared "github.com/pjbgf/sha1cd/internal"
)

// blockARM64 hashes the message p into the current state in dig.
// Both m1 and cs are used to store intermediate results which are used by the collision detection logic.
//
//go:noescape
func blockARM64(dig *digest, p sliceHeader, m1 []uint32, cs [][5]uint32)

func block(dig *digest, p []byte) {
	m1 := [shared.Rounds]uint32{}
	cs := [shared.PreStepState][shared.WordBuffers]uint32{}

	for len(p) >= shared.Chunk {
		// Only send a block to be processed, as the collission detection
		// works on a block by block basis.
		ips := sliceHeader{
			base: uintptr(unsafe.Pointer(&p[0])),
			len:  int(math.Min(float64(len(p)), float64(shared.Chunk))),
			cap:  shared.Chunk,
		}

		blockARM64(dig, ips, m1[:], cs[:])

		col := checkCollision(m1, cs, dig.h)
		if col {
			dig.col = true

			blockARM64(dig, ips, m1[:], cs[:])
			blockARM64(dig, ips, m1[:], cs[:])
		}

		p = p[shared.Chunk:]
	}
}
