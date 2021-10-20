package epub

import (
	"testing"
)

func Test_extractTex(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			"simple",
			args{
				s: `The general principle is that the
change in the energy is the force times the distance that the force is
pushed, and that this is a change in energy in general:
\begin{equation}
\label{Eq:I:4:4}

\begin{pmatrix}
\text{change in}\\
\text{energy}
\end{pmatrix}=
(\text{force})\times

\begin{pmatrix}
\text{distance force}\\
\text{acts through}
\end{pmatrix}.
\end{equation}
We will return to many of these other kinds of energy as we continue the
course.
				`,
			},
			[][]byte{
				[]byte(`\begin{equation}
\label{Eq:I:4:4}

\begin{pmatrix}
\text{change in}\\
\text{energy}
\end{pmatrix}=
(\text{force})\times

\begin{pmatrix}
\text{distance force}\\
\text{acts through}
\end{pmatrix}.
\end{equation}`),
			},
		},
		{
			"two times",
			args{
				s: `The general principle is that the
change in the energy is the force times the distance that the force is
pushed, and that this is a change in energy in general:
\begin{equation}
\label{Eq:I:4:4}

\begin{pmatrix}
\text{change in}\\
\text{energy}
\end{pmatrix}=
(\text{force})\times

\begin{pmatrix}
\text{distance force}\\
\text{acts through}
\end{pmatrix}.
\end{equation}
We will return to many of these other kinds of energy as we continue the
course.
The general principle is that the
change in the energy is the force times the distance that the force is
pushed, and that this is a change in energy in general:
\begin{equation}
\label{Eq:I:4:5}

\begin{pmatrix}
\text{change in}\\
\text{energy}
\end{pmatrix}=
(\text{force})\times

\begin{pmatrix}
\text{distance force}\\
\text{acts through}
\end{pmatrix}.
\end{equation}
We will return to many of these other kinds of energy as we continue the
course.
				`,
			},
			[][]byte{
				[]byte(`\begin{equation}
\label{Eq:I:4:4}

\begin{pmatrix}
\text{change in}\\
\text{energy}
\end{pmatrix}=
(\text{force})\times

\begin{pmatrix}
\text{distance force}\\
\text{acts through}
\end{pmatrix}.
\end{equation}`),
				[]byte(`\begin{equation}
\label{Eq:I:4:5}

\begin{pmatrix}
\text{change in}\\
\text{energy}
\end{pmatrix}=
(\text{force})\times

\begin{pmatrix}
\text{distance force}\\
\text{acts through}
\end{pmatrix}.
\end{equation}`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractTex(tt.args.s)
			if len(got) != len(tt.want) {
				t.Errorf("extractTex() = %v, want %v", got, tt.want)
			}
			for i := 0; i < len(got); i++ {
				if string(got[i]) != string(tt.want[i]) {
					t.Errorf("extractTex() = %v, want %v", string(got[i]), string(tt.want[i]))

				}
			}
		})
	}
}
