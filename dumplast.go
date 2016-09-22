// 22 october 2015
package main

import (
	"fmt"
	"bytes"

	"github.com/andlabs/reallymine/command"
	"github.com/andlabs/reallymine/disk"
)

var zeroSector [disk.SectorSize]byte

func cDumpLast(d *disk.Disk) error {
	var sector []byte

	// TODO add -fakesize option of sorts
	pos := d.Size()
	iter, err := d.ReverseIter(pos)
	if err != nil {
		return err
	}

	found := false
	for iter.Next() {
		sector = iter.Sectors()
		pos = iter.Pos()
		if !bytes.Equal(sector, zeroSector[:]) {
			found = true
			break
		}
	}
	if err = iter.Err(); err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("non-empty sector not found")
	}

	fmt.Print(dumpSector(sector, pos))
	return nil
}

var dumplast = &command.Command{
	Name:		"dumplast",
	Args:		[]string{"disk"},
	Description:	"Hexdumps the last non-zero sector in disk.",
	Do:			cDumpLast,
}