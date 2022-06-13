package main
//package sort 


import (
  "fmt"
  "time"
  "os"
  "text/tabwriter"
  "sort"
)


/*type Interface interface { 
  Len() int 
  Less(i, j int) bool // i, j are indices of sequence elements 
  Swap(i, j int) 
}
*/



type Track struct { 
  Title  string 
  Artist string 
  Album  string 
  Year int 
  Length time.Duration //represents the elapsed time between two instants as an int64 nanosecond count.
  } 
  
var tracks = []*Track{ 
  {"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")}, 
  {"Go", "Moby", "Moby", 1992, length("3m37s")}, 
  {"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")}, 
  {"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")}, 
}

func length(s string) time.Duration { 
  d, err := time.ParseDuration(s) // ParseDuration parses a duration string.
  if err != nil { 
    panic(s) 
    } 
  return d 
}

func printTracks(tracks []*Track) { 
  const format = "%v\t%v\t%v\t%v\t%v\t\n" //Format string
  tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2,' ',0) //translates tabbed columns in input into properly aligned text
  fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length") 
  fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------") 
  for _, t := range tracks { 
    fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length) 
  } 
  tw.Flush() // calculate column widths and print table 
}

type byArtist []*Track 

func (x byArtist) Len() int { return len(x) } 
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist } 
func (x byArtist) Swap(i, j int) {x[i], x[j] = x[j], x[i] }

type byTitle []*Track 

func (x byTitle) Len() int { return len(x) } 
func (x byTitle) Less(i, j int) bool { return x[i].Title < x[j].Title } 
func (x byTitle) Swap(i, j int) {x[i], x[j] = x[j], x[i] }

func main(){
  
  fmt.Println("Tracks available ..")
  printTracks(tracks)
  sort.Sort(byArtist(tracks))
  fmt.Println("\nTracks sorted by Artist ..")
  printTracks(tracks)
  sort.Sort(byTitle(tracks))
  fmt.Println("\nTracks sorted by Title ..")
  printTracks(tracks)
  
}