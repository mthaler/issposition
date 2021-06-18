package tle

import (
	"bufio"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestReadTLE(t *testing.T) {
	s := `ISS (ZARYA)             
1 25544U 98067A   21168.87490640  .00001211  00000-0  30151-4 0  9995
2 25544  51.6445 339.3221 0003457 102.3260 342.4446 15.48994679288707`

	r := strings.NewReader(s)
	scanner := bufio.NewScanner(r)

	type args struct {
		scanner *bufio.Scanner
	}
	tests := []struct {
		name    string
		args    args
		want    TLE
		wantErr bool
	}{
		{name: "ISS",
			args:    args{scanner: scanner},
			want:    NewTLE("ISS (ZARYA)", "1 25544U 98067A   21168.87490640  .00001211  00000-0  30151-4 0  9995", "2 25544  51.6445 339.3221 0003457 102.3260 342.4446 15.48994679288707"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadTLE(tt.args.scanner)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadTLE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadTLE() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadTLEs(t *testing.T) {
	s := `ISS (ZARYA)             
1 25544U 98067A   21168.87490640  .00001211  00000-0  30151-4 0  9995
2 25544  51.6445 339.3221 0003457 102.3260 342.4446 15.48994679288707
KESTREL EYE IIM (KE2M)  
1 42982U 98067NE  21168.54480641  .00076049  13605-4  18146-3 0  9993
2 42982  51.6223 191.5090 0004178  31.1160 329.0095 15.96700404208771`

	r := strings.NewReader(s)

	m := make(map[string]TLE)
	m["ISS (ZARYA)"] = NewTLE("ISS (ZARYA)", "1 25544U 98067A   21168.87490640  .00001211  00000-0  30151-4 0  9995", "2 25544  51.6445 339.3221 0003457 102.3260 342.4446 15.48994679288707")
	m["KESTREL EYE IIM (KE2M)"] = NewTLE("KESTREL EYE IIM (KE2M)", "1 42982U 98067NE  21168.54480641  .00076049  13605-4  18146-3 0  9993", "2 42982  51.6223 191.5090 0004178  31.1160 329.0095 15.96700404208771")

	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]TLE
		wantErr bool
	}{
		{name: "ISS",
			args: args{r: r},
			want: m,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadTLEs(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadTLEs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadTLEs() got = %v, want %v", got, tt.want)
			}
		})
	}
}