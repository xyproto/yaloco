package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

const versionString = "YALOCO 1.0.0"

func init() {
	color.NoColor = false
}

// colorWrite writes a colored string to the given strings.Builder
func colorWrite(sb *strings.Builder, s string, colorIndex int) {
	switch colorIndex % 4 {
	case 0:
		sb.WriteString(s)
	case 1:
		sb.WriteString(color.WhiteString(s))
	case 2:
		sb.WriteString(color.MagentaString(s))
	case 3:
		sb.WriteString(color.CyanString(s))
	}
}

// colorize comments in gray. Colorize (){}[] in an alternating way.
func colorize(line string) {
	var sb strings.Builder
	trimmed := strings.TrimSpace(line)
	if strings.HasPrefix(trimmed, "//") || strings.HasPrefix(trimmed, "#") || strings.Contains(trimmed, "Leaving directory") || strings.Contains(trimmed, "Entering directory") {
		fmt.Println(color.HiBlackString(line))
		return
	}
	if strings.Contains(trimmed, "In file included from") {
		fmt.Println(color.HiBlueString(line))
		return
	}
	if strings.HasPrefix(trimmed, "OPKG ") {
		fmt.Println(color.MagentaString(line))
		return
	}
	if strings.HasPrefix(trimmed, "STRIP ") {
		fmt.Println(color.CyanString(line))
		return
	}
	if (strings.HasPrefix(trimmed, "*") && strings.HasSuffix(trimmed, "*")) || (strings.HasPrefix(trimmed, "-") && strings.HasSuffix(trimmed, "-")) || (strings.HasPrefix(trimmed, "=") && strings.HasSuffix(trimmed, "=")) || strings.Contains(trimmed, "***") || strings.Contains(trimmed, "===") || strings.Contains(trimmed, ">>>") {
		fmt.Println(color.RedString(line))
		return
	}

	if strings.Contains(trimmed, ": In function ") || strings.Contains(trimmed, ": In member function ") {
		elements := strings.SplitN(trimmed, ":", 2)
		fn := elements[0]
		msg := elements[1]
		if strings.Count(msg, "'") >= 2 {
			parts := strings.SplitN(msg, "'", 3)
			a := parts[0]
			signature := parts[1]
			b := parts[2]
			fmt.Println(color.HiYellowString(fn) + ":" + color.HiWhiteString(a) + "'" + color.HiBlueString(signature) + "'" + color.HiWhiteString(b))
			return
		}
	}

	var (
		rainbowLine    strings.Builder
		curlyLevel     int
		bracketLevel   int
		parLevel       int
		closing        bool
		colorIndex     int
		word           string
		changed        bool
		singleToggle   bool
		quotingChanged bool
	)
	for _, c := range line {
		quotingChanged = false
		switch c {
		case '\'':
			if strings.Count(line, "'")%2 == 0 && !strings.Contains(line, "n't") {
				singleToggle = !singleToggle
				quotingChanged = true
				changed = true
				closing = false
			} else {
				changed = false
				closing = false
			}
		case '{':
			curlyLevel++
			colorIndex++
			changed = true
			closing = false
		case '[':
			bracketLevel++
			colorIndex++
			changed = true
			closing = false
		case '(':
			parLevel++
			colorIndex++
			changed = true
			closing = false
		case '}':
			curlyLevel--
			closing = true
			changed = true
		case ']':
			bracketLevel--
			closing = true
			changed = true
		case ')':
			parLevel--
			closing = true
			changed = true
		default:
			changed = false
			closing = false
		}
		// If the level changed, output the word we've got so far
		if changed {
			// THIS IS THE PLACE TO PROCESS THE THING BETWEEN THE BRACKETS
			colorWrite(&rainbowLine, word, colorIndex)
			// Then bump the color, if closing
			prevColor := colorIndex
			if closing && (colorIndex > 0) {
				colorIndex--
			}
			// Or bump the color, if the quoting changed
			if c == '\'' && quotingChanged {
				if singleToggle {
					colorIndex++
				} else {
					colorIndex--
				}
			}
			if c == '\'' && singleToggle {
				prevColor = colorIndex
			}
			// Then output the opening/closing thing
			if c == '\'' {
				colorWrite(&rainbowLine, string(c), prevColor)
			} else {
				colorWrite(&rainbowLine, string(c), prevColor)
			}
			// Then reset the word
			word = ""
		} else if c == ' ' {
			word += string(c)
			// Then output the opening/closing thing
			colorWrite(&rainbowLine, word, colorIndex)
			word = ""
		} else {
			// The level did not change, continue to collect the word
			word += string(c)
		}
	}
	if word != "" {
		// THIS IS THE SECOND PLACE TO PROCESS THE THING BETWEEN THE BRACKETS
		colorWrite(&rainbowLine, word, colorIndex)
	}

	line = rainbowLine.String()

	for i, word := range strings.Split(line, " ") {
		if i > 0 {
			sb.WriteString(" ")
		}
		switch word {
		case "\"GET", "\"POST":
			sb.WriteString(color.CyanString(string(word[0])) + color.BlueString(word[1:]))
			continue
		}
		switch strings.ToLower(word) {
		case "error", "error:", "abort", "quit":
			sb.WriteString(color.HiRedString(word))
		case "warning", "warning:", "removed", "deleted", "erased":
			sb.WriteString(color.HiYellowString(word))
		case "note", "note:":
			sb.WriteString(color.HiGreenString(word))
		case "cxx", "ld", "rm", "make", "strip", "ccgi", "opkg", "install", "run", "running", "move", "format", "upgrading":
			sb.WriteString(color.BlueString(word))
		case "upgraded", "installed", "moved", "ran", "formatted":
			sb.WriteString(color.MagentaString(word))
		case "=", "==", ":=":
			sb.WriteString(color.HiWhiteString(word))
		default:
			sb.WriteString(color.CyanString(word))
		}
	}

	fmt.Println(sb.String())
}

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] != "-" {
			fmt.Println(versionString)
			os.Exit(1)
		}
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		colorize(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
