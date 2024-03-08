Execute a program
=========================
    go run program.go


fmt
=========================

        https://pkg.go.dev/fmt

        fmt.Print(str), 
        fmt.Printf(format, str), 
        fmt.Println(str)  
                => Format string (default formatter or custom ) and output formatted string to standard output

        fmt.Sprint(str),
        fmt.Sprintf(format, str)
        fmt.Sprintln(str)
                => Format string (default formatter or custom) and return formatted string


        fmt.Fprint(writer, str),
        fmt.Fprintf(writer, format, str),
        fmt.Fprintln(writer, str)
                => Format string (default formatter or custom) and write specified writer
        