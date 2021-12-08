package partition

import "fmt"

func Add(args []string) {
	fmt.Println("add partition")
}

func add() error {
	return nil
}

func Drop(args []string) {
	fmt.Println("drop partition")
}

func drop() error {
	return nil
}

func Migrate(args []string) {
	fmt.Println("migrate partition")
}

func migrate() error {
	return nil
}

func Autopilot() {
	fmt.Println("autopilot partition")

	// TODO:
	// Check Partition Status

	// Add necessary partitions

	// Migrate partitions to Minio or Greenplum

	// Drop unnecessary partitions

}
