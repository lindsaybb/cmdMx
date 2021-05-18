package main

import (
        "fmt"
        "strings"
        "flag"
)

var (
        helpFlag = flag.Bool("h", false, "Show this help")
        iterFlag = flag.Int("n", 1, "Number of times to repeat")
        incrFlag = flag.Bool("i", false, "Increment the output")
        stepFlag = flag.Int("b", 1, "Increment by how many")
        startFlag = flag.Int("s", 1, "Increment Start Value")
)

const usage = `
'cmdMx' [options] <args>

Where <args> is the string that is to be repeated -n times.
If <args> contains %d they will be replaced with the -s value.
If the -i flag is thrown this -s value will be incremented each -n time.
The -b flag can be used to increment by a value other than 1.
`

func main() {
        flag.Parse()
        if *helpFlag || flag.NArg() < 1 {
                fmt.Println(usage)
                flag.PrintDefaults()
                return
        }
        args := flag.Args()             // []string
        cnt := *startFlag
        var found []int                 // holds the position of where matches are identified
        var cp []string                 // copies the args to prevent overwrite when replace

        for i, v := range args {
                // loop through args and check each one
                // match criteria restricted to just digits
                // this method allows multiple replacements
                // however each replacement will be incremented
                if strings.Contains(v, "%d") {
                        found = append(found, i)
                }
                cp = append(cp, v)
        }
        var line string
        for i := 0; i < *iterFlag; i++ {
        // loop -n times
                for _, c := range found {
                // loop through each occurance of a match criteria
                        // repl is the string which matched the criteria
                        // where %d is replaced by the counter value cnt
                        repl := strings.Replace(args[c], "%d", fmt.Sprintf("%d", cnt), -1)
                        args[c] = repl
                        if *incrFlag {
                                cnt += *stepFlag
                        }
                }
                // once the subsets of the line have been replaced
                // join all elements and print
                line = strings.Join(args, " ")
                fmt.Println(line)
                // this loop reverts the args placeholder back to original
                for n := range args {
                        args[n] = cp[n]
                }
        }
}
