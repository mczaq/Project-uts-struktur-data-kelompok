package main

import (
	"fmt"
	"os"
)

type node struct {
	doctorName  string
	patientName string
	medicine    string
	next        *node
}

func main() {
	var head *node = nil
	var choice int
	for {
		fmt.Println("1. Masukkan data")
		fmt.Println("2. Lihat data")
		fmt.Println("3. Hapus data")
		fmt.Println("4. Keluar")
		fmt.Println("Masukkan pilihan Anda:")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			head = insertData(head)
		case 2:
			viewData(head)
		case 3:
			head = deleteData(head)
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Pilihan yang Anda masukkan salah")
		}
	}
}

func insertData(head *node) *node {
	var newNode *node = &node{}
	fmt.Println("Masukkan nama dokter:")
	fmt.Scan(&newNode.doctorName)
	fmt.Println("Masukkan nama pasien:")
	fmt.Scan(&newNode.patientName)
	fmt.Println("Masukkan obat:")
	fmt.Scan(&newNode.medicine)
	newNode.next = head
	head = newNode
	return head
}

func viewData(head *node) {
	if head == nil {
		fmt.Println("Data kosong")
		return
	}
	fmt.Println("Data kesehatan:")
	for current := head; current != nil; current = current.next {
		fmt.Println("Dokter:", current.doctorName)
		fmt.Println("Pasien:", current.patientName)
		fmt.Println("Obat:", current.medicine)
		fmt.Println("------------------------")
	}
}

func deleteData(head *node) *node {
	var doctorName string
	if head == nil {
		fmt.Println("Data kosong")
		return nil
	}
	fmt.Println("Masukkan nama dokter yang ingin dihapus:")
	fmt.Scan(&doctorName)
	if head.doctorName == doctorName {
		head = head.next
		return head
	}
	for current := head; current.next != nil; current = current.next {
		if current.next.doctorName == doctorName {
			current.next = current.next.next
			return head
		}
	}
	fmt.Println("Data tidak ditemukan")
	return head
}
