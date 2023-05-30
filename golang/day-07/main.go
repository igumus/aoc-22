package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const (
	ROOT = "/"

	PREFIX_CMD = "$"
	CMD_CD     = "cd"
	CMD_LS     = "ls"

	ALIAS_DIRNAME_CURR = "."  // alias for directory we are in
	ALIAS_DIRNAME_UPPR = ".." // alias for directory to parent

	TYPE_DIR = "dir"
)

type element struct {
	file     bool
	size     int64
	name     string
	parent   *element
	children map[string]*element
}

func (e *element) totalSize() int64 {
	if e.file {
		return e.size
	}
	size := int64(0)
	for _, d := range e.children {
		size += d.totalSize()
	}
	return size
}

func (e *element) cd(name string) *element {
	if name == ALIAS_DIRNAME_UPPR {
		if e.parent == nil {
			return e
		}
		return e.parent
	}
	return e.children[name]
}

func (e *element) processContent(contents ...string) []*element {
	if len(contents) == 0 {
		return nil
	}

	dirs := make([]*element, 0)

	for _, content := range contents {
		cnt := strings.Split(content, " ")
		typeID := strings.TrimSpace(cnt[0])
		name := strings.TrimSpace(cnt[1])
		ne := &element{
			name:     name,
			parent:   e,
			children: make(map[string]*element),
		}
		if typeID != TYPE_DIR {
			xsize, _ := strconv.Atoi(typeID)
			ne.file = true
			ne.size = int64(xsize)
		}
		if typeID == TYPE_DIR {
			dirs = append(dirs, ne)
		}
		e.children[name] = ne
	}
	return dirs
}

func parse(sc *bufio.Scanner) (*element, []*element) {
	root := &element{
		name:     ROOT,
		parent:   nil,
		children: make(map[string]*element),
	}
	dirs := make([]*element, 0)
	node := root
	isListMode := false
	contents := make([]string, 0)
	for sc.Scan() {
		input := strings.TrimSpace(sc.Text())
		if strings.HasPrefix(input, PREFIX_CMD) {
			if isListMode {
				xdirs := node.processContent(contents...)
				if len(xdirs) > 0 {
					dirs = append(dirs, xdirs...)
				}
				contents = make([]string, 0)
				isListMode = !isListMode
			}

			command := strings.TrimSpace(strings.Replace(input, PREFIX_CMD, "", 1))
			args := strings.Split(command, " ")
			switch args[0] {
			case CMD_CD:
				if ROOT == args[1] {
					node = root
					break
				}
				node = node.cd(args[1])
			case CMD_LS:
				isListMode = !isListMode
			}
		} else {
			contents = append(contents, input)
		}
	}
	if len(contents) > 0 {
		xdirs := node.processContent(contents...)
		if len(xdirs) > 0 {
			dirs = append(dirs, xdirs...)
		}
	}
	return root, dirs
}

func calculate(sc *bufio.Scanner) int64 {
	limit := int64(100000)
	_, dirs := parse(sc)
	total := int64(0)
	for _, val := range dirs {
		size := val.totalSize()
		if size <= limit {
			total += size
		}
	}
	return total
}

func calculatePartTwo(sc *bufio.Scanner) int64 {
	totalDiskSpace := int64(70000000)
	minRequiredFreeSpace := int64(30000000)
	node, dirs := parse(sc)
	toFree := minRequiredFreeSpace - (totalDiskSpace - node.totalSize())
	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].totalSize() < dirs[j].totalSize()
	})
	fmt.Println(toFree, len(dirs))
	fmt.Println(dirs[0].totalSize(), dirs[len(dirs)-1].totalSize())
	for _, dir := range dirs {
		if dir.totalSize() >= toFree {
			return dir.totalSize()
		}
	}
	return toFree
}
