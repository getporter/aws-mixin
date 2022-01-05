package aws

import "fmt"

// Build will generate the necessary Dockerfile lines
// for an invocation image using this mixin
func (m *Mixin) Build() error {
	// TODO: This gets whatever the latest version of the cli is, there isn't a way for us to say what version we are using
	fmt.Fprintln(m.Out, `RUN apt-get update && apt-get install -y --no-install-recommends curl unzip libc6 less groff`)
	fmt.Fprintln(m.Out, `RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-$(uname -m).zip" -o "awscliv2.zip"`)
	fmt.Fprintln(m.Out, `RUN unzip awscliv2.zip`)
	fmt.Fprintln(m.Out, `RUN ./aws/install`)
	fmt.Fprintln(m.Out, `RUN rm -fr awscliv2.zip ./aws`)
	return nil
}
