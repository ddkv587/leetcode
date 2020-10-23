package isMatch

import "fmt"

type node struct {
    ch byte
    hasAsterisk bool
    leastCount int 
}

func formatScale( p string ) []node {
    scale := make( []node, 0 )
    if len( p ) == 0 {
        return scale
    }

    for i := 0; i < len(p) - 1;  {
        if p[i] == '*' {
            fmt.Println( "error when parse" )
            return scale
        }

        if p[ i + 1 ] == '*' {
            if len( scale ) > 0 && scale[ len( scale ) - 1 ].hasAsterisk && 
                ( scale[ len( scale ) - 1 ].ch == p[i] || scale[ len( scale ) - 1 ].ch == '.' ) {
                ;
            } else {
                scale = append( scale, node{ p[i], true, 0 } )
            }
            i += 2
        } else {
            if len( scale ) > 0 && scale[ len( scale ) - 1 ].hasAsterisk && scale[ len( scale ) - 1 ].ch == p[i] {
                scale[ len( scale ) - 1 ].leastCount++
            } else {
                scale = append( scale, node{ p[i], false, 1 } )
            }
            i++
        }
    }
    if p[len(p) - 1] != '*' {
        if len( scale ) > 0 && scale[ len( scale ) - 1 ].hasAsterisk && 
                scale[ len( scale ) - 1 ].ch == p[len(p) - 1] {
            scale[ len( scale ) - 1 ].leastCount++
        } else {
            scale = append( scale, node{ p[len(p) - 1], false, 1 } )
        }
    }

    return scale
}

func innerMatch( s []byte, scale []node ) bool {
    for ; len( s ) > 0 && len( scale ) > 0; {
        //fmt.Printf( "s: %s, scale: %v\n", s, scale )
        if !scale[0].hasAsterisk {
            if scale[0].ch == '.' || s[0] == scale[0].ch {
                scale = scale[1:]
                s = s[1:]
            } else {
                return false
            }
        } else {
            if scale[0].ch == '.' || scale[0].ch == s[0] {
                if scale[0].leastCount > 0 {
                    scale[0].leastCount--
                    s = s[1:]
                } else {
                    return innerMatch( s, append( []node(nil), scale[1:]... ) ) || 
                            innerMatch( append( []byte(nil), s[ 1: ]... ), scale )
                                
                }
            } else {
                if scale[0].leastCount > 0 {
                    return false
                } else {
                    scale = scale[1:]
                }
            }
        }
    }

    for _, node := range scale {
        if node.hasAsterisk && node.leastCount <= 0 {
            scale = scale[1:]
        } else {
            break
        }
    }

    //fmt.Printf( "111s: %v, scale: %v\n", s, scale )

    return ( len( scale ) == 0 ) && ( len( s ) == 0 )
}

func IsMatch(s string, p string) bool {
    scale := formatScale( p )

    return innerMatch( []byte( s ), scale )
}