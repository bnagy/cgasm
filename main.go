package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	verbose = flag.Bool("v", false, "Full documentation ( default is a truncated summary )")
	fuzzy   = flag.Bool("f", false, "Fuzzy search (search with go regexp syntax)")
)

func lookup(item string) (s string, ok bool) {
	s, ok = data[item]
	if !ok {
		return
	}
	if strings.HasPrefix(s, "-R:") {
		// Follow the redirect entry
		return lookup(strings.TrimPrefix(s, "-R:"))
	}
	return
}

func summary(item string) string {
	scanner := bufio.NewScanner(strings.NewReader(item))
	var output bytes.Buffer
	scanner.Scan()
	scanner.Scan()
	output.WriteString(strings.TrimSuffix(scanner.Text(), ":"))
	output.WriteString("\n\n")
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "Description:") {
			break
		}
	}
	if scanner.Text() == "" {
		return output.String()
	}
	output.WriteString(scanner.Text())
	output.WriteString("\n")
	for i := 15; i > 0; i-- {
		// Output up to the first 15 lines of the first paragraph of the
		// Description section
		if scanner.Scan() {
			// If we have hit the Operation section we have gone too far.
			// Should be impossible, because we would have broken on the blank
			// line.
			if strings.HasPrefix(scanner.Text(), "Operation:") || len(scanner.Text()) == 0 {
				break
			}
			output.WriteString(scanner.Text())
			output.WriteString("\n")
		}
		if i == 1 {
			output.WriteString("\n[... use -v to see full output ...]\n")
		}
	}
	return output.String()
}

func getHeader(item string) string {
	s, _ := lookup(item)
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Scan()
	scanner.Scan()
	header := strings.TrimSuffix(scanner.Text(), ":")
	if !strings.Contains(header, " - ") {
		fmt.Fprintf(os.Stderr, "[BUG] unexpected entry data for %s\n", header)
		os.Exit(1)
	}
	return header
}

func fuzzySearch(item string) []string {

	results := []string{}
	if !strings.HasPrefix(item, "(?") {
		item = "(?i)" + item
	}

	r, err := regexp.Compile(item)
	if err != nil {
		return results
	}
	for k, v := range data {
		if r.MatchString(k) {
			if strings.HasPrefix(v, "-R:") {
				results = append(results, fmt.Sprintf("%s -> %s", k, getHeader(k)))
				continue
			}
			results = append(results, getHeader(k))
		}
	}
	return results
}

func main() {

	flag.Parse()
	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	query := flag.Arg(0)

	out, ok := lookup(strings.ToUpper(query))
	if !ok {
		// No exact match, try a fuzzy search
		if *fuzzy {
			matches := fuzzySearch(query)
			if len(matches) == 0 {
				fmt.Printf("%s - no documentation found.\n", query)
				os.Exit(0)
			}
			fmt.Printf("Fuzzy matches for \"%s\" (%d):\n", query, len(matches))
			for _, match := range matches {
				fmt.Println(match)
			}
			os.Exit(0)
		}

		fmt.Printf("\n%s - no documentation found.\n", query)
		os.Exit(0)
	}

	if !strings.Contains(out, "\nDescription:") {
		fmt.Fprintf(os.Stderr, "[BUG] unexpected entry data for %s\n", query)
		os.Exit(1)
	}

	if *verbose {
		fmt.Println(out)
		os.Exit(0)
	}

	fmt.Println(summary(out))
}
