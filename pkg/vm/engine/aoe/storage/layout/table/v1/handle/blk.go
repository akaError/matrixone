package handle

import (
	"fmt"
	logutil2 "matrixone/pkg/logutil"
	"matrixone/pkg/vm/engine/aoe/storage/common"
	"matrixone/pkg/vm/engine/aoe/storage/layout/index"
	"matrixone/pkg/vm/engine/aoe/storage/layout/table/v1/col"
	"matrixone/pkg/vm/engine/aoe/storage/mock/type/chunk"
	"sync"
)

var (
	// allocTimes    = 0
	blkHandlePool = sync.Pool{
		New: func() interface{} {
			// allocTimes++
			// log.Infof("Alloc blk handle: %d", allocTimes)
			h := new(BlockHandle)
			h.Cols = make([]col.IColumnBlock, 0)
			// h.Cursors = make([]col.ScanCursor, 0)
			return h
		},
	}
)

type BlockHandle struct {
	ID          common.ID
	Cols        []col.IColumnBlock
	IndexHolder *index.BlockHolder
}

func (bh *BlockHandle) GetID() *common.ID {
	return &bh.ID
}

func (bh *BlockHandle) GetColumn(idx int) col.IColumnBlock {
	if idx < 0 || idx >= len(bh.Cols) {
		panic(fmt.Sprintf("Specified idx %d is out of scope", idx))
	}
	return bh.Cols[idx]
}

func (bh *BlockHandle) GetIndexHolder() *index.BlockHolder {
	return bh.IndexHolder
}

func (bh *BlockHandle) Close() error {
	if bhh := bh; bhh != nil {
		for _, col := range bhh.Cols {
			col.UnRef()
		}
		bhh.Cols = bhh.Cols[:0]
		// TODO
		// blkHandlePool.Put(bhh)
		bh = nil
	}
	return nil
}

func (bh *BlockHandle) InitScanCursor() []col.ScanCursor {
	cursors := make([]col.ScanCursor, len(bh.Cols))
	for idx, colBlk := range bh.Cols {
		colBlk.InitScanCursor(&cursors[idx])
		err := cursors[idx].Init()
		if err != nil {
			logutil2.Error(fmt.Sprintf("logic error: %s", err))
			panic(fmt.Sprintf("logic error: %s", err))
		}
	}
	return cursors
}

func (bh *BlockHandle) Fetch() *chunk.Chunk {
	// TODO
	return nil
}
