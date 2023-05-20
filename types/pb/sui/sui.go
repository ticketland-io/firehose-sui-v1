package sui_checkpoint_v1

import (
	"encoding/binary"
	"encoding/hex"
	"time"

	"github.com/streamingfast/bstream"
)

func (b *CheckpointData) AsRef() bstream.BlockRef {
	return bstream.NewBlockRef(b.ID(), b.Checkpoint.SequenceNumber)
}

func (b *CheckpointData) ID() string {
	return uint64ToID(b.Checkpoint.SequenceNumber)
}

func (b *CheckpointData) Number() uint64 {
	return b.Checkpoint.SequenceNumber
}

func (b *CheckpointData) PreviousID() string {
	if b.Checkpoint.SequenceNumber == 0 {
		return ""
	}

	return uint64ToID(b.PreviousNum())
}

func (b *CheckpointData) PreviousNum() uint64 {
	if b.Checkpoint.SequenceNumber == 0 {
		return 0
	}

	return b.Checkpoint.SequenceNumber - 1
}

func (b *CheckpointData) Time() time.Time {
	return time.UnixMilli(int64(b.Checkpoint.TimestampMs))
}

func (b *CheckpointData) LIBNum() uint64 {
	number := b.Number()
	if number <= bstream.GetProtocolFirstStreamableBlock {
		return number
	}

	// Since there is no forks blocks on Aptos, I'm pretty sure that last irreversible block number
	// is the block's number itself. However I'm not sure overall how the Firehose stack would react
	// to LIBNum == Num so to play safe for now, previous block of current is irreversible.
	return b.Number() - 1
}

func uint64ToHash(height uint64) []byte {
	id := make([]byte, 8)
	binary.BigEndian.PutUint64(id, height)

	return id
}

func uint64ToID(height uint64) string {
	return hex.EncodeToString(uint64ToHash(height))
}
