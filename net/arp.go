package net

import "path/filepath"
import "fmt"
import "io"
import "os"

// Configure the ARP cache parameters for the given interface.  This
// makes containers react more quickly to a change in the MAC address
// associated with an IP address.
func ConfigureARPCache(root, name string) error {
	if err := sysctl(root, fmt.Sprintf("net/ipv4/neigh/%s/base_reachable_time", name), "5"); err != nil {
		return err
	}
	if err := sysctl(root, fmt.Sprintf("net/ipv4/neigh/%s/delay_first_probe_time", name), "2"); err != nil {
		return err
	}
	if err := sysctl(root, fmt.Sprintf("net/ipv4/neigh/%s/ucast_solicit", name), "1"); err != nil {
		return err
	}
	return nil
}

func sysctl(root, variable, value string) error {
	f, err := os.OpenFile(filepath.Join(root, "/proc/sys/", variable), os.O_WRONLY, 0)
	if err != nil {
		return err
	}

	n, err := io.WriteString(f, value)
	if err != nil {
		return err
	}
	if n < len(value) {
		return io.ErrShortWrite
	}

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}
