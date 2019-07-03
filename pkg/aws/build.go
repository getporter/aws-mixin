package aws

import "fmt"

// Build will generate the necessary Dockerfile lines
// for an invocation image using this mixin
func (m *Mixin) Build() error {
	// TODO: This gets whatever the latest version of the cli is, there isn't a way for us to say what version we are using
	fmt.Fprintln(m.Out, `RUN apt-get update && apt-get install -y curl unzip python less groff`)
	fmt.Fprintln(m.Out, `RUN curl "https://s3.amazonaws.com/aws-cli/awscli-bundle.zip" -o "/tmp/awscli-bundle.zip"`)
	fmt.Fprintln(m.Out, `RUN unzip /tmp/awscli-bundle.zip -d /tmp`)
	fmt.Fprintln(m.Out, `RUN /tmp/awscli-bundle/install -i /usr/local/aws -b /usr/local/bin/aws`)
	fmt.Fprintln(m.Out, `RUN rm -fr /tmp/awscli-bundle.zip /tmp/awscli-bundle`)
	return nil
}
