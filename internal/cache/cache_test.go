package cache

import (
	"fmt"
	"testing"
)

func TestCasche01(t *testing.T) {
	cacheTest := InitCache(5, 6, 3)

	if cacheTest.SN.casheSize != 5 {
		fmt.Printf("SN cache size should be 5, it is %d\n", cacheTest.SN.casheSize)
		t.Fail()
	}
	if cacheTest.HU.casheSize != 6 {
		fmt.Printf("HU cache size should be 6, it is %d\n", cacheTest.HU.casheSize)
		t.Fail()
	}
	if cacheTest.File.casheSize != 3 {
		fmt.Printf("File cache size should be 3, it is %d\n", cacheTest.SN.casheSize)
		t.Fail()
	}

	SerialEntries := []string{"First", "Second", "Third", "Fourth", "Fifth", "Sixth"}
	//order of loading into cache
	first := SerialEntries[0]
	second := SerialEntries[1]
	third := SerialEntries[2]
	fourth := SerialEntries[3]
	fifth := SerialEntries[4]

loop:
	for i, entry := range SerialEntries {
		cacheTest.SN.Contains(entry)
		firstNode, ok := cacheTest.SN.hashmap[entry]
		if !ok {
			fmt.Printf("Entry \"%s\"; i == %d; is not in the map\n", entry, i)
			fmt.Println("Further tests aborted, fatal error...")
			t.FailNow()
		}
		if i == 0 {
			if cacheTest.SN.first != firstNode && cacheTest.SN.last != firstNode {
				fNode := cacheTest.SN.first != firstNode
				lNode := cacheTest.SN.last != firstNode
				fmt.Printf("Node with key \"%s\" should be the first and the last entry\n", entry)
				fmt.Printf("First = %v\nLast = %v\n", fNode, lNode)
				fmt.Println("Further tests aborted, fatal error...")
				t.FailNow()
			}
			continue loop
		}
		if i == 1 {
			lastNode, ok := cacheTest.SN.hashmap[first]
			if !ok {
				fmt.Printf("Entry \"%s\"; i == %d; is not in the map", first, i)
				fmt.Println("Further tests aborted, fatal error...")
				t.FailNow()
			}
			if cacheTest.SN.first != firstNode && cacheTest.SN.last != lastNode {
				fNode := cacheTest.SN.first != firstNode
				lNode := cacheTest.SN.last != lastNode
				fmt.Printf("Node with key \"%s\" should be the first and Node with key \"%s\" the last entry\n", entry, first)
				fmt.Printf("First = %v\nLast = %v\n", fNode, lNode)
				fmt.Println("Further tests aborted, fatal error...")
				t.FailNow()
			}
			continue loop

		}
		if i == 4 {
			//one == fisrt node; five == last node
			one, ok := cacheTest.SN.hashmap[fifth]
			if !ok {
				fmt.Printf("Entry \"%s\"; i == %d; is not in the map\n", fifth, i)
				fmt.Println("Further tests aborted, fatal error...")
				t.FailNow()
			}
			two, ok := cacheTest.SN.hashmap[fourth]
			if !ok {
				fmt.Printf("Entry \"%s\"; i == %d; is not in the map\n", fourth, i)
				fmt.Println("Further tests aborted, fatal error...")
				t.FailNow()
			}
			three, ok := cacheTest.SN.hashmap[third]
			if !ok {
				fmt.Printf("Entry \"%s\"; i == %d; is not in the map\n", third, i)
				fmt.Println("Further tests aborted, fatal error...")
				t.FailNow()
			}
			four, ok := cacheTest.SN.hashmap[second]
			if !ok {
				fmt.Printf("Entry \"%s\"; i == %d; is not in the map\n", second, i)
				fmt.Println("Further tests aborted, fatal error...")
				t.FailNow()
			}
			five, ok := cacheTest.SN.hashmap[first]
			if !ok {
				fmt.Printf("Entry \"%s\"; i == %d; is not in the map\n", first, i)
				fmt.Println("Further tests aborted, fatal error...")
				t.FailNow()
			}
			//order check
			if cacheTest.SN.first != one && cacheTest.SN.last != five {
				fNode := cacheTest.SN.first != one
				lNode := cacheTest.SN.last != five
				fmt.Printf("Node with key \"%s\" should be the first and Node with key \"%s\" the last entry\n", fifth, first)
				fmt.Printf("First = %v\nLast = %v\n", fNode, lNode)
				fmt.Println("Further tests aborted, fatal error...")
				t.FailNow()
			}

			if cacheTest.SN.first.after != two {
				fmt.Println("Node two mismatch.")
				fmt.Println("Further tests aborted, fatal error...")
				t.FailNow()
			}

			if cacheTest.SN.last.before != four {
				fmt.Println("Node four cache node mismatch.")
				fmt.Println("Further tests aborted, fatal error...")
				t.FailNow()
			}
			if (cacheTest.SN.first.after.after != three) && (cacheTest.SN.last.before.before != three) {
				fmt.Println("Chainling broken")
				fmt.Println("Node three cache node mismatch.")
				fmt.Println("Further tests aborted, fatal error...")
				t.FailNow()
			}
			continue loop
		}

		if i == 5 {
			if _, ok := cacheTest.SN.hashmap[first]; ok {
				fmt.Printf("Node with key \"%s\" should not exist.\n", first)
				t.Fail()
			}
			if cacheTest.SN.entryCount != 5 {
				fmt.Printf("entryCount should be 5; is == %d\n", cacheTest.SN.entryCount)
				t.Fail()
			}
		}

	}

}
