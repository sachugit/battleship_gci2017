package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func read_input(filePath string) []string {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return lines
}

func parse_data(space_list string, grid_size int) [][]string {

	p1 := make([][]string, grid_size)
	space_list_data := strings.Split(space_list, ",")

	for i := 0; i < grid_size; i++ {
		p1[i] = make([]string, grid_size)
		for j := 0; j < grid_size; j++ {
			p1[i][j] = "_"
		}
	}
	for i := range space_list_data {
		space_coord := strings.Split(space_list_data[i], ":")
		x, _ := strconv.Atoi(space_coord[0])
		y, _ := strconv.Atoi(space_coord[1])

		p1[x][y] = "B"
	}
	return p1
}

func calculate_missile_hit(missiles []string, spaces [][]string) int {

	hit_count := 0
	for i := range missiles {
		vals := strings.Split(missiles[i], ":")

		x, _ := strconv.Atoi(vals[0])
		y, _ := strconv.Atoi(vals[1])

		if spaces[x][y] == "B" {
			spaces[x][y] = "X"
			hit_count = hit_count + 1
		} else {
			spaces[x][y] = "O"
		}
	}
	return hit_count
}

func get_output_file(filePath string) *os.File {

	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	return f
}

func write_grid(grid [][]string, out_file *os.File, grid_size int) {

	for i := 0; i < grid_size; i++ {
		for j := 0; j < grid_size; j++ {
			out_file.WriteString(grid[i][j])
		}
		out_file.WriteString("\n")
	}
}

func main() {

	lines := read_input("input.txt")
	grid_size, _ := strconv.Atoi(lines[0])

	space_size, _ := strconv.Atoi(lines[1])
	p1_spaces_str := lines[2]
	p2_spaces_str := lines[3]
	p1_spaces := parse_data(p1_spaces_str, grid_size)
	p2_spaces := parse_data(p2_spaces_str, grid_size)

	if space_size != len(p1_spaces) || space_size != len(p2_spaces) {
		println("Invalid SpaceShip data length")
	}

	missile_size, _ := strconv.Atoi(lines[4])
	p1_missiles_str := lines[5]
	p2_missiles_str := lines[6]
	p1_missiles := strings.Split(p1_missiles_str, ",")
	p2_missiles := strings.Split(p2_missiles_str, ",")

	if missile_size != len(p1_missiles) || missile_size != len(p2_missiles) {
		println("Invalid Missile data length")
	}

	p1_hits := calculate_missile_hit(p1_missiles, p2_spaces)
	p2_hits := calculate_missile_hit(p2_missiles, p1_spaces)

	out_file := get_output_file("out.txt")
	defer out_file.Close()

	out_file.WriteString("Player 1\n")
	write_grid(p1_spaces, out_file, grid_size)
	out_file.WriteString("Player 2\n")
	write_grid(p2_spaces, out_file, grid_size)

	out_file.WriteString(fmt.Sprintf("P1: %d\n", p1_hits))
	out_file.WriteString(fmt.Sprintf("P2: %d\n", p2_hits))
	if p1_hits > p2_hits {
		out_file.WriteString("Player 1 wins\n")
	} else if p1_hits < p2_hits {
		out_file.WriteString("Player 2 wins\n")
	} else {
		out_file.WriteString("It is a draw\n")
	}
}
