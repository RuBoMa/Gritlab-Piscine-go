package main

import (
	"asciiart/ascii"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintAsciiArt(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "hello",
			input: "hello",
			want: ` _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
`,
		},
		{
			name:  "HELLO",
			input: "HELLO",
			want: ` _    _   ______   _        _         ____   
| |  | | |  ____| | |      | |       / __ \  
| |__| | | |__    | |      | |      | |  | | 
|  __  | |  __|   | |      | |      | |  | | 
| |  | | | |____  | |____  | |____  | |__| | 
|_|  |_| |______| |______| |______|  \____/  
                                             
                                             
`,
		},
		{
			name:  "1Hello 2There",
			input: "1Hello 2There",
			want: `     _    _          _   _                         _______   _                           
 _  | |  | |        | | | |                ____   |__   __| | |                          
/ | | |__| |   ___  | | | |   ___         |___ \     | |    | |__     ___   _ __    ___  
| | |  __  |  / _ \ | | | |  / _ \          __) |    | |    |  _ \   / _ \ | '__|  / _ \ 
| | | |  | | |  __/ | | | | | (_) |        / __/     | |    | | | | |  __/ | |    |  __/ 
|_| |_|  |_|  \___| |_| |_|  \___/        |_____|    |_|    |_| |_|  \___| |_|     \___| 
                                                                                         
                                                                                         
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ascii.PrintAsciiArt(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
