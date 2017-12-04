package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/stellar/go/keypair"
)

var (
	prefix   = flag.String("prefix", "", "prefix")
	seed     = flag.String("seed", "", "seed")
	nWorkers = flag.Int("n", 1, "number of workers")
)

func usage() {
	fmt.Fprintln(os.Stderr, "usage: stellar-keypairs [ -n nWorkers ] [ -p prefix | -s seed ]")
	os.Exit(1)
}

func main() {
	flag.Parse()

	if flag.NArg() != 0 {
		usage()
	}

	if *prefix != "" && *seed != "" {
		usage()
	}

	if *prefix == "" && *seed == "" {
		err := generateKeyPairRandom()
		if err != nil {
			log.Fatal(err)
		}
	}

	if *seed != "" {
		err := generateKeyPairSeed(*seed)
		if err != nil {
			log.Fatal(err)
		}
	}

	if *prefix != "" {
		if !isValidPrefix(*prefix) {
			log.Fatalf("prefix %s is not base32", *prefix)
		}
		if *nWorkers == 1 {
			err := generateKeyPairPattern(*prefix)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err := generateKeyPairPatternParallel(*prefix)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func generateKeyPairRandom() error {
	kp, err := keypair.Random()
	if err != nil {
		return err
	}

	err = printKeys(kp)
	if err != nil {
		return err
	}

	return nil
}

func generateKeyPairSeed(seed string) error {
	kp, err := keypair.Parse(seed)
	if err != nil {
		return err
	}

	fmt.Println("Seed (secret key)", seed)
	fmt.Println("Public key", kp.Address())

	return nil
}

func generateKeyPairPattern(prefix string) error {
	var kp *keypair.Full
	var err error

	for {
		kp, err = keypair.Random()
		if err != nil {
			return err
		}
		if keyHasPrefix(kp, prefix) {
			break
		}
	}

	err = printKeys(kp)
	if err != nil {
		return err
	}

	return nil
}

func generateKeyPairPatternParallel(prefix string) error {
	keypairs := make(chan *keypair.Full, 0)

	for w := 1; w <= *nWorkers; w++ {
		go worker(w, keypairs, prefix)
	}

	for kp := range keypairs {
		err := printKeys(kp)
		if err != nil {
			return err
		}
	}

	return nil
}

func worker(id int, keypairs chan<- *keypair.Full, prefix string) {
	for {
		kp, err := keypair.Random()
		if err != nil {
			log.Println(err)
			continue
		}
		if keyHasPrefix(kp, prefix) {
			keypairs <- kp
		}
	}
}

func keyHasPrefix(kp keypair.KP, prefix string) bool {
	return strings.HasPrefix(kp.Address(), prefix)
}

func printKeys(kp *keypair.Full) error {
	fmt.Println("Seed (secret key)", kp.Seed())
	fmt.Println("Public key", kp.Address())
	return nil
}

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"

func isValidPrefix(prefix string) bool {
	if len(prefix) < 1 {
		return false
	}
	if prefix[0] != 'G' {
		return false
	}
	for _, r := range prefix {
		if !strings.ContainsRune(alphabet, r) {
			return false
		}
	}
	return true
}
