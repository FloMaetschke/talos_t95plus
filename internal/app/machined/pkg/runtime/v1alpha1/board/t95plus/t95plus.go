// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package t95plus

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/talos-systems/go-procfs/procfs"
	"golang.org/x/sys/unix"

	"github.com/talos-systems/talos/internal/app/machined/pkg/runtime"
	"github.com/talos-systems/talos/pkg/copy"
	"github.com/talos-systems/talos/pkg/machinery/constants"
)

var (
	bin       = fmt.Sprintf("/usr/install/arm64/u-boot/%s/u-boot-rockchip.bin", constants.BoardT95Plus)
	off int64 = 512 * 64
  dtb       = "/dtb/rockchip/rk3399-nanopi-r4s.dtb"      /// TODO: ADD DTB?
)

// T95Plus represents the T95Plus Android TV Box board.
//
// Reference: https://t95plus.com
type T95Plus struct{}

// Name implements the runtime.Board.
func (n *T95Plus) Name() string {
	return constants.BoardT95Plus
}

// Install implements the runtime.Board.
func (n *T95Plus) Install(disk string) (err error) {
	file, err := os.OpenFile(disk, os.O_RDWR|unix.O_CLOEXEC, 0o666)
	if err != nil {
		return err
	}

	defer file.Close() //nolint:errcheck

	uboot, err := os.ReadFile(bin)
	if err != nil {
		return err
	}

	log.Printf("writing %s at offset %d", bin, off)

	amount, err := file.WriteAt(uboot, off)
	if err != nil {
		return err
	}

	log.Printf("wrote %d bytes", amount)

	// NB: In the case that the block device is a loopback device, we sync here
	// to esure that the file is written before the loopback device is
	// unmounted.
	if err := file.Sync(); err != nil {
		return err
	}

	src := "/usr/install/arm64" + dtb
	dst := "/boot/EFI" + dtb

	if err := os.MkdirAll(filepath.Dir(dst), 0o600); err != nil {
		return err
	}

	return copy.File(src, dst)
}

// KernelArgs implements the runtime.Board.
func (n *T95Plus) KernelArgs() procfs.Parameters {
	return []*procfs.Parameter{
		procfs.NewParameter("console").Append("tty0").Append("ttyS2,1500000n8"),
		procfs.NewParameter("sysctl.kernel.kexec_load_disabled").Append("1"),
	}
}

// PartitionOptions implements the runtime.Board.
func (n *T95Plus) PartitionOptions() *runtime.PartitionOptions {
	return &runtime.PartitionOptions{PartitionsOffset: 2048 * 10}
}
